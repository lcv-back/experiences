package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// User đại diện cho dữ liệu người dùng
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Response đại diện cho phản hồi JSON
type Response struct {
	Message string `json:"message"`
}

var db *sql.DB

func main() {
	// Kết nối tới SQLite
	var err error
	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Tạo bảng users nếu chưa có
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE,
			email TEXT UNIQUE,
			password TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Khởi tạo router
	router := mux.NewRouter()
	router.HandleFunc("/register", registerHandler).Methods("POST")

	// Khởi động server
	log.Println("Server running at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

// Xử lý yêu cầu đăng ký với transaction
func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, `{"message": "Dữ liệu không hợp lệ!"}`, http.StatusBadRequest)
		return
	}

	// Kiểm tra dữ liệu đầu vào
	if user.Username == "" || user.Email == "" || user.Password == "" {
		json.NewEncoder(w).Encode(Response{Message: "Vui lòng điền đầy đủ thông tin!"})
		return
	}

	// Bắt đầu transaction
	tx, err := db.Begin()
	if err != nil {
		http.Error(w, `{"message": "Lỗi server khi khởi tạo giao dịch!"}`, http.StatusInternalServerError)
		return
	}

	// Đảm bảo rollback nếu có lỗi
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
		if err != nil {
			log.Printf("Error committing transaction: %v", err)
		}
	}()

	// Kiểm tra xem username hoặc email đã tồn tại chưa
	var existingUser int
	err = tx.QueryRow("SELECT COUNT(*) FROM users WHERE username = ? OR email = ?", user.Username, user.Email).Scan(&existingUser)
	if err != nil {
		http.Error(w, `{"message": "Lỗi server khi kiểm tra dữ liệu!"}`, http.StatusInternalServerError)
		return
	}
	if existingUser > 0 {
		json.NewEncoder(w).Encode(Response{Message: "Tên người dùng hoặc email đã tồn tại!"})
		return
	}

	// Mã hóa mật khẩu
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, `{"message": "Lỗi server khi mã hóa mật khẩu!"}`, http.StatusInternalServerError)
		return
	}

	// Chèn dữ liệu vào cơ sở dữ liệu
	_, err = tx.Exec(
		"INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		user.Username, user.Email, hashedPassword,
	)
	if err != nil {
		http.Error(w, `{"message": "Lỗi server khi lưu dữ liệu!"}`, http.StatusInternalServerError)
		return
	}

	// Trả về phản hồi thành công
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{Message: "Đăng ký thành công!"})
}
