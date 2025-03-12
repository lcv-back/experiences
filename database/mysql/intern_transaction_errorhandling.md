# Các câu hỏi phỏng vấn ở level intern, fresher

### 1. **Giải thích khái niệm về transaction trong MySQL và cách sử dụng chúng trong Golang.**

- **Transaction** trong MySQL là một nhóm các câu lệnh SQL mà được thực thi như một khối duy nhất. Một transaction đảm bảo rằng các thay đổi dữ liệu được thực hiện một cách nhất quán. Nếu có một lỗi xảy ra trong khi thực hiện transaction, bạn có thể **rollback** để hủy bỏ các thay đổi đã thực hiện, đảm bảo tính toàn vẹn của dữ liệu.

- **Các bước của một transaction**:
  - **START TRANSACTION**: Bắt đầu một transaction.
  - **COMMIT**: Xác nhận và lưu các thay đổi trong database.
  - **ROLLBACK**: Hủy bỏ các thay đổi nếu có lỗi xảy ra.

Trong Golang, bạn có thể sử dụng thư viện `github.com/go-sql-driver/mysql` để thực hiện các transaction như sau:

#### Ví dụ sử dụng transaction trong Golang:

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

    // Bắt đầu một transaction
    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }

    // Thực hiện các câu lệnh SQL trong transaction
    _, err = tx.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 25)
    if err != nil {
        // Nếu có lỗi, rollback transaction
        tx.Rollback()
        log.Fatal(err)
    }

    // Thực hiện thêm một câu lệnh SQL trong transaction
    _, err = tx.Exec("INSERT INTO orders (user_id, amount) VALUES (?, ?)", 1, 100)
    if err != nil {
        // Nếu có lỗi, rollback transaction
        tx.Rollback()
        log.Fatal(err)
    }

    // Nếu không có lỗi, commit transaction
    err = tx.Commit()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Transaction completed successfully!")
}
```

**Giải thích**:

- **`db.Begin()`**: Mở một transaction.
- **`tx.Exec()`**: Thực thi các câu lệnh SQL trong transaction.
- **`tx.Rollback()`**: Nếu có lỗi xảy ra, rollback transaction để hủy bỏ các thay đổi.
- **`tx.Commit()`**: Nếu không có lỗi, commit transaction để lưu các thay đổi.

### 2. **Làm thế nào để quản lý lỗi trong quá trình thực hiện truy vấn MySQL khi sử dụng Golang?**

Quản lý lỗi trong quá trình thực hiện truy vấn MySQL trong Golang có thể được thực hiện qua các phương pháp sau:

1. **Kiểm tra lỗi ngay sau mỗi câu lệnh SQL**: Sau khi thực hiện bất kỳ câu lệnh SQL nào, bạn nên kiểm tra lỗi ngay lập tức. Điều này giúp phát hiện các vấn đề kịp thời và xử lý lỗi hiệu quả.

2. **Sử dụng `defer` để đảm bảo đóng kết nối hoặc đóng `rows`**: Các tài nguyên như kết nối database hoặc các kết quả truy vấn (`rows`) cần phải được đóng sau khi sử dụng xong, ngay cả khi có lỗi xảy ra.

3. **Rollback transaction khi có lỗi**: Nếu bạn đang sử dụng transaction, bạn phải đảm bảo rollback transaction nếu có lỗi xảy ra, để tránh thay đổi không mong muốn được lưu vào database.

#### Ví dụ về cách xử lý lỗi khi thực hiện truy vấn:

```go
func main() {
    // Chuỗi kết nối MySQL
    dsn := "user:password@tcp(localhost:3306)/dbname"

    // Mở kết nối
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err) // Nếu kết nối gặp lỗi, dừng chương trình
    }
    defer db.Close() // Đảm bảo đóng kết nối sau khi hoàn tất

    // Thực hiện một câu lệnh SQL
    result, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "John", 30)
    if err != nil {
        log.Fatal(err) // Nếu có lỗi, dừng chương trình và in lỗi
    }

    // Kiểm tra số lượng dòng bị ảnh hưởng
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Fatal(err) // Nếu không thể lấy được số dòng bị ảnh hưởng, in lỗi
    }

    fmt.Printf("Number of rows affected: %d\n", rowsAffected)
}
```

**Giải thích**:

- **Kiểm tra lỗi sau mỗi câu lệnh SQL**: Mỗi lần gọi hàm như `db.Exec()` hoặc `tx.Exec()`, chúng ta cần kiểm tra lỗi ngay lập tức để phát hiện vấn đề sớm.
- **Sử dụng `defer`**: Chúng ta sử dụng `defer db.Close()` để đảm bảo kết nối được đóng sau khi hoàn tất công việc.
- **Lấy thông tin về số dòng bị ảnh hưởng**: Sau khi thực thi một câu lệnh như `INSERT`, chúng ta có thể sử dụng `result.RowsAffected()` để xem số dòng bị ảnh hưởng.

### Tóm tắt:

- **Transactions** trong MySQL là một cách nhóm các câu lệnh SQL lại với nhau để thực hiện chúng một cách nhất quán, có thể rollback nếu có lỗi.
- **Cách sử dụng transactions trong Golang**: Mở transaction bằng `db.Begin()`, thực hiện các câu lệnh SQL bằng `tx.Exec()`, và cuối cùng commit hoặc rollback transaction dựa trên kết quả.
- **Quản lý lỗi**: Kiểm tra lỗi ngay sau mỗi câu lệnh SQL, sử dụng `defer` để đóng tài nguyên, và rollback transaction nếu có lỗi.
