Dưới đây là phần tổng hợp các câu hỏi bạn đưa ra, được chia theo **chủ đề** và có câu trả lời rõ ràng, dễ hiểu để bạn ôn tập phỏng vấn vị trí backend developer:

---

### I. **Ngôn ngữ Go (Golang)**

#### 1. **Gin Framework**

- **Khi nhận req đến response trong Gin như thế nào?**

  - `Gin` sử dụng cơ chế middleware. Khi một request đến:
    1. Request vào Router → tìm route phù hợp.
    2. Các middleware được thực thi theo thứ tự.
    3. Handler xử lý request.
    4. Response được trả về qua middleware theo chiều ngược lại (giống pipeline).

- **Logger trong Gin:**
  - Dùng `gin.Logger()` hoặc tự custom logger bằng cách sử dụng `logrus`, `zap`, `zerolog` và gắn vào middleware.

#### 2. **Struct và Interface**

- **Interface là gì? Khi nào dùng?**

  - Interface định nghĩa hành vi, không cài đặt cụ thể. Dùng để:
    - Tạo tính linh hoạt.
    - Hỗ trợ mock/test dễ dàng.
    - Đạt tính đa hình.

- **Struct là gì? Dùng khi nào?**

  - Struct là kiểu dữ liệu tự định nghĩa gồm nhiều field. Dùng khi cần biểu diễn dữ liệu có cấu trúc, như model (User, Product).

- **Tính kế thừa và đa hình trong Go?**
  - Go không hỗ trợ kế thừa OOP, nhưng hỗ trợ **composition**:
    ```go
    type Animal struct { Name string }
    type Dog struct { Animal }
    ```
  - Đa hình: thông qua interface:
    ```go
    type Speaker interface { Speak() }
    ```

#### 3. **Defer và xử lý lỗi**

- **Defer là gì?**

  - `defer` trì hoãn thực thi function đến khi function bao quanh kết thúc.
  - Thường dùng để:
    - Đóng file.
    - Đóng DB connection.
    - Ghi log.
    ```go
    defer file.Close()
    ```

- **Xử lý lỗi trong Go:**

  - Dùng `error` interface, thường kiểm tra theo kiểu:
    ```go
    if err != nil {
        return err
    }
    ```
  - Không có try/catch như Java.

- **So sánh với Java:**
  - Java có try/catch và throwable.
  - Go kiểm soát lỗi thủ công, nên dễ nhìn và tối ưu performance, nhưng code verbose.

#### 4. **Goroutine và Thread**

- **So sánh:**
  | Góc nhìn | Goroutine (Go) | Thread (Java, OS) |
  |-------------|------------------------|---------------------------|
  | Chi phí | Nhẹ, vài KB | Nặng, vài MB |
  | Scheduler | Do Go runtime quản lý | Do OS quản lý |
  | Số lượng | Hàng nghìn dễ dàng | Giới hạn (OS level) |

- **Override trong Go:**
  - Không có override như Java. Muốn override thì dùng interface:
    ```go
    type A interface { Do() }
    type B struct {}
    func (b B) Do() { fmt.Println("Override!") }
    ```

---

### II. **SQL và Database**

#### 1. **Các loại JOIN và khác nhau:**

| Loại JOIN  | Mô tả ngắn gọn                                            |
| ---------- | --------------------------------------------------------- |
| INNER JOIN | Chỉ lấy các bản ghi khớp ở cả 2 bảng                      |
| LEFT JOIN  | Lấy toàn bộ bảng trái + bản ghi khớp ở bảng phải (nếu có) |
| RIGHT JOIN | Lấy toàn bộ bảng phải + bản ghi khớp ở bảng trái (nếu có) |
| FULL JOIN  | Lấy tất cả bản ghi từ cả 2 bảng, nếu không khớp thì NULL  |
| CROSS JOIN | Nhân đề các dòng giữa 2 bảng (Cartesian product)          |

#### 2. **Tối ưu DB & Query**

- **Tối ưu database:**

  - Index đúng chỗ (các cột dùng trong WHERE, JOIN).
  - Tránh SELECT \*.
  - Chuẩn hóa dữ liệu (normalization).
  - Sử dụng LIMIT, phân trang.
  - Caching (Redis).

- **Giải thích vì sao index nhanh hơn:**
  - Index giống như mục lục trong sách. Thay vì đọc toàn bộ bảng (full scan), DB chỉ cần tra index để tìm dòng cần → nhanh hơn rất nhiều.

---

### III. **API và Web**

#### 1. **Thiết kế API gồm các thành phần:**

- **Endpoint**: đường dẫn định danh tài nguyên.
- **Method**: GET, POST, PUT, DELETE,...
- **Request/Response Format**: thường là JSON.
- **Status Code**: 200, 400, 401, 500,...
- **Middleware**: auth, log, error handler, rate-limit,...

#### 2. **WebSocket vs REST**

| Tiêu chí     | WebSocket                          | REST (HTTP API)              |
| ------------ | ---------------------------------- | ---------------------------- |
| Kết nối      | Liên tục, hai chiều                | Mỗi request là một kết nối   |
| Giao tiếp    | Realtime, push/pull dữ liệu        | Client phải chủ động request |
| Dùng khi nào | Chat app, game, realtime dashboard | CRUD thông thường, đơn giản  |
| Ưu điểm      | Giảm độ trễ, không cần polling     | Đơn giản, dễ tích hợp        |

- **Dùng WebSocket chưa?**  
  → Có thể nói: "Em từng dùng WebSocket để làm chức năng chat realtime, sử dụng thư viện như Gorilla Websocket (Go)."

---

### IV. **Các công nghệ liên quan**

#### 1. **gRPC**

- RPC framework nhanh, nhẹ, dùng protobuf.
- Ưu điểm:

  - Performace cao, nhẹ hơn REST.
  - Tự generate code từ `.proto`.
  - Hỗ trợ streaming, auth, load balance tốt.

- **Dùng khi:**
  - Microservices cần tốc độ cao.
  - Nội bộ hệ thống, không public.

#### 2. **Kết nối database trong Go dùng gì?**

- `database/sql` + driver (Postgres, MySQL).
- ORM phổ biến: `gorm`, `sqlx`.

---

### V. **Kinh nghiệm & Dự án**

- **Dự án từng làm qua là gì?**

  - Nêu một dự án backend có REST API, database, authentication, pagination,...

- **Các DB đã từng làm:**
  - MySQL, PostgreSQL, MongoDB, Redis,...

---
