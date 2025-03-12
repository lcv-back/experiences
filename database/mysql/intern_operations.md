# Các câu hỏi phỏng vấn ở level intern, fresher

Dưới đây là câu trả lời cho các câu hỏi liên quan đến thao tác dữ liệu trong MySQL và Golang:

### 1. **Làm thế nào để chèn một bản ghi mới vào cơ sở dữ liệu MySQL từ Golang? Ví dụ về việc sử dụng `INSERT INTO`.**

Để chèn một bản ghi mới vào cơ sở dữ liệu MySQL từ Golang, bạn có thể sử dụng câu lệnh `INSERT INTO`. Dưới đây là ví dụ:

#### Ví dụ:

```go
func main() {
    // Chuỗi kết nối MySQL
    dsn := "user:password@tcp(localhost:3306)/dbname"

    // Mở kết nối
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Câu lệnh INSERT INTO
    query := "INSERT INTO users (name, age) VALUES (?, ?)"
    result, err := db.Exec(query, "Alice", 25)
    if err != nil {
        log.Fatal(err)
    }

    // Kiểm tra số dòng bị ảnh hưởng
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Number of rows affected: %d\n", rowsAffected)
}
```

**Giải thích**:

- `db.Exec()` thực thi câu lệnh `INSERT INTO` và chèn dữ liệu vào bảng `users`.
- Các giá trị `"Alice"` và `25` được chèn vào bảng thông qua dấu hỏi (`?`), giúp tránh SQL injection.
- Hàm `result.RowsAffected()` trả về số dòng bị ảnh hưởng bởi câu lệnh.

### 2. **Làm thế nào để cập nhật dữ liệu trong MySQL từ Golang? Ví dụ về câu lệnh `UPDATE`.**

Để cập nhật dữ liệu trong MySQL từ Golang, bạn sử dụng câu lệnh `UPDATE`. Dưới đây là ví dụ:

#### Ví dụ:

```go
func main() {
    // Chuỗi kết nối MySQL
    dsn := "user:password@tcp(localhost:3306)/dbname"

    // Mở kết nối
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Câu lệnh UPDATE
    query := "UPDATE users SET age = ? WHERE name = ?"
    result, err := db.Exec(query, 30, "Alice")
    if err != nil {
        log.Fatal(err)
    }

    // Kiểm tra số dòng bị ảnh hưởng
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Number of rows affected: %d\n", rowsAffected)
}
```

**Giải thích**:

- `db.Exec()` thực thi câu lệnh `UPDATE` để thay đổi tuổi của người có tên `"Alice"` thành `30`.
- Hàm `result.RowsAffected()` trả về số dòng bị ảnh hưởng, giúp bạn biết có bao nhiêu bản ghi được cập nhật.

### 3. **Làm thế nào để xóa dữ liệu trong MySQL từ Golang?**

Để xóa dữ liệu trong MySQL từ Golang, bạn sử dụng câu lệnh `DELETE`. Dưới đây là ví dụ:

#### Ví dụ:

```go
func main() {
    // Chuỗi kết nối MySQL
    dsn := "user:password@tcp(localhost:3306)/dbname"

    // Mở kết nối
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Câu lệnh DELETE
    query := "DELETE FROM users WHERE name = ?"
    result, err := db.Exec(query, "Alice")
    if err != nil {
        log.Fatal(err)
    }

    // Kiểm tra số dòng bị ảnh hưởng
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Number of rows affected: %d\n", rowsAffected)
}
```

**Giải thích**:

- `db.Exec()` thực thi câu lệnh `DELETE` để xóa người dùng có tên `"Alice"` khỏi bảng `users`.
- Hàm `result.RowsAffected()` trả về số dòng bị ảnh hưởng, giúp bạn biết có bao nhiêu bản ghi bị xóa.

### Tóm tắt:

- **Chèn bản ghi**: Sử dụng câu lệnh `INSERT INTO` với `db.Exec()`.
- **Cập nhật dữ liệu**: Sử dụng câu lệnh `UPDATE` với `db.Exec()`.
- **Xóa dữ liệu**: Sử dụng câu lệnh `DELETE` với `db.Exec()`.

Các thao tác này đều sử dụng phương thức `db.Exec()` trong Go, và việc truyền tham số vào câu lệnh SQL thông qua dấu hỏi (`?`) giúp bảo vệ ứng dụng khỏi SQL injection.
