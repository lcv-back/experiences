# Các câu hỏi phỏng vấn ở level intern, fresher

### 1. **Làm thế nào để kết nối và thao tác với MySQL trong Golang? Hãy mô tả quy trình kết nối một cơ sở dữ liệu MySQL trong Go.**

Quy trình kết nối một cơ sở dữ liệu MySQL trong Golang thường bao gồm các bước sau:

1. **Cài đặt thư viện MySQL**: Đầu tiên, bạn cần cài đặt thư viện MySQL cho Golang. Thư viện phổ biến là `github.com/go-sql-driver/mysql`.

   Cài đặt thư viện:

   ```bash
   go get -u github.com/go-sql-driver/mysql
   ```

2. **Nhập thư viện**: Bạn cần import thư viện MySQL vào chương trình Go.

   ```go
   import (
       "database/sql"
       "fmt"
       "log"

       _ "github.com/go-sql-driver/mysql" // Import package để sử dụng driver
   )
   ```

3. **Mở kết nối**: Bạn cần sử dụng hàm `sql.Open()` để tạo kết nối đến cơ sở dữ liệu MySQL. Lưu ý rằng `sql.Open()` không thực sự mở kết nối mà chỉ chuẩn bị kết nối. Để thực sự kết nối, bạn cần gọi hàm `db.Ping()` để kiểm tra kết nối.

   ```go
   func main() {
       // Chuỗi kết nối MySQL
       dsn := "user:password@tcp(localhost:3306)/dbname"

       // Mở kết nối
       db, err := sql.Open("mysql", dsn)
       if err != nil {
           log.Fatal(err)
       }

       // Kiểm tra kết nối
       err = db.Ping()
       if err != nil {
           log.Fatal(err)
       }
       fmt.Println("Connected to the database successfully!")
   }
   ```

4. **Thực hiện các thao tác CRUD**: Sau khi kết nối thành công, bạn có thể thực hiện các câu lệnh SQL như `SELECT`, `INSERT`, `UPDATE`, `DELETE`.

5. **Đóng kết nối**: Khi hoàn tất các thao tác, đừng quên đóng kết nối cơ sở dữ liệu bằng cách gọi `db.Close()`.

### 2. **Bạn sử dụng thư viện nào trong Go để giao tiếp với MySQL? Giải thích lý do tại sao bạn chọn thư viện đó.**

- **Thư viện: `github.com/go-sql-driver/mysql`**.
- **Lý do chọn thư viện này**:
  - Đây là thư viện chính thức và phổ biến nhất để kết nối Golang với MySQL.
  - Hỗ trợ đầy đủ các tính năng của MySQL, bao gồm giao dịch, khóa ngoại, và chuẩn chuẩn kết nối.
  - Cung cấp API đơn giản và dễ sử dụng.
  - Được bảo trì và cập nhật thường xuyên, hỗ trợ cả MySQL và MariaDB.
  - Thư viện này cũng hỗ trợ sử dụng context và cung cấp các tùy chọn tối ưu hóa hiệu suất.

### 3. **Thực hiện một câu lệnh `SELECT` trong Golang sử dụng MySQL, và xử lý kết quả trả về như thế nào?**

Giả sử bạn muốn thực hiện một câu lệnh `SELECT` và xử lý kết quả trả về. Dưới đây là ví dụ cách thực hiện:

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

    // Thực hiện câu lệnh SELECT
    rows, err := db.Query("SELECT id, name FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Lặp qua các kết quả
    for rows.Next() {
        var id int
        var name string

        // Đọc kết quả từ rows
        if err := rows.Scan(&id, &name); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d, Name: %s\n", id, name)
    }

    // Kiểm tra lỗi sau khi lặp qua rows
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
}
```

**Giải thích**:

- `db.Query()` thực hiện câu lệnh `SELECT` và trả về một `rows` chứa tất cả các kết quả.
- `rows.Next()` được sử dụng để duyệt qua tất cả các bản ghi.
- `rows.Scan()` lấy dữ liệu từ mỗi bản ghi và gán vào các biến.
- Đảm bảo sử dụng `defer` để đóng `rows` và kết nối sau khi hoàn tất.

### 4. **Bạn làm thế nào để tránh SQL injection khi sử dụng Go với MySQL? Ví dụ về việc sử dụng `prepared statements`.**

Để tránh SQL injection, **`prepared statements`** là một kỹ thuật rất hiệu quả. Kỹ thuật này giúp tách biệt câu lệnh SQL và dữ liệu đầu vào của người dùng, từ đó ngăn chặn SQL injection.

Ví dụ sử dụng `prepared statements` trong Go:

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

    // Sử dụng prepared statement để tránh SQL injection
    stmt, err := db.Prepare("SELECT id, name FROM users WHERE age = ?")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    // Thực thi câu lệnh với tham số
    rows, err := stmt.Query(30) // Truy vấn người dùng có độ tuổi 30
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Xử lý kết quả trả về
    for rows.Next() {
        var id int
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d, Name: %s\n", id, name)
    }

    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
}
```

**Giải thích**:

- `db.Prepare()` tạo ra một `prepared statement` với dấu hỏi (`?`) thay vì tham số trực tiếp.
- Các giá trị tham số được truyền vào khi gọi `stmt.Query()`, thay vì chèn trực tiếp vào câu lệnh SQL.
- Kỹ thuật này giúp tránh SQL injection vì giá trị đầu vào không thể bị chèn vào câu lệnh SQL dưới dạng mã độc.

### Tóm tắt:

- **Kết nối MySQL trong Go**: Sử dụng `sql.Open()` để mở kết nối, `db.Ping()` để kiểm tra kết nối.
- **Thư viện MySQL trong Go**: Sử dụng `github.com/go-sql-driver/mysql`, vì nó phổ biến và được hỗ trợ tốt.
- **Thực hiện câu lệnh `SELECT`**: Dùng `db.Query()` để lấy kết quả và xử lý với `rows.Scan()`.
- **Tránh SQL injection**: Dùng `prepared statements` với dấu hỏi (`?`) để ngăn chặn SQL injection.
