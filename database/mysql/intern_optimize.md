# Các câu hỏi phỏng vấn ở level intern, fresher

### 1. **Làm thế nào để tối ưu hóa một câu truy vấn trong MySQL? Bạn có thể chỉ ra các chỉ số quan trọng nào để tối ưu hóa?**

Tối ưu hóa câu truy vấn trong MySQL giúp giảm thiểu thời gian xử lý và tài nguyên hệ thống. Một số cách để tối ưu hóa truy vấn bao gồm:

- **Sử dụng chỉ mục (index)**: Đảm bảo các cột thường xuyên được truy vấn (ví dụ: trong `WHERE`, `JOIN`, `ORDER BY`,...) có chỉ mục.
- **Tránh sử dụng `SELECT *`**: Chỉ chọn các cột cần thiết thay vì lấy tất cả các cột, giúp giảm lượng dữ liệu cần xử lý.
- **Tránh sử dụng các hàm trong câu lệnh `WHERE`**: Các hàm như `LOWER()`, `NOW()`,... có thể làm mất hiệu quả của chỉ mục.
- **Sử dụng `JOIN` thay vì `SUBQUERY`**: Câu lệnh `JOIN` thường nhanh hơn `SUBQUERY`, vì nó tối ưu tốt hơn trong nhiều trường hợp.
- **Giới hạn số lượng bản ghi trả về**: Sử dụng `LIMIT` nếu không cần tất cả kết quả.
- **Cấu trúc dữ liệu hợp lý**: Đảm bảo bảng có cấu trúc chuẩn (normalization), tránh dữ liệu dư thừa.

**Các chỉ số quan trọng để tối ưu hóa**:

- **Sử dụng chỉ mục**: Kiểm tra xem câu truy vấn có sử dụng chỉ mục không.
- **Sử dụng bộ nhớ (buffer pool)**: Kiểm tra chỉ số về bộ nhớ sử dụng như `key_buffer_size`, `innodb_buffer_pool_size`.
- **Lượng dữ liệu quét**: Xem số lượng bản ghi được quét và số lượng bản ghi trả về.
- **Tổng thời gian thực thi**: Sử dụng `EXPLAIN` để xác định thời gian và cách MySQL thực thi câu truy vấn.

### 2. **Giải thích về chỉ mục (index) trong MySQL và khi nào thì nên sử dụng chúng.**

- **Chỉ mục (Index)** là cấu trúc dữ liệu giúp MySQL tìm kiếm dữ liệu nhanh hơn. Chỉ mục tương tự như một danh mục trong sách, giúp truy xuất dữ liệu nhanh hơn mà không phải duyệt qua tất cả các bản ghi.

**Khi nào sử dụng chỉ mục**:

- **Tăng tốc độ tìm kiếm**: Khi bạn thực hiện truy vấn `SELECT` với điều kiện `WHERE`, `JOIN`, `ORDER BY`, chỉ mục sẽ giúp tìm kiếm nhanh chóng.
- **Truy vấn với cột thường xuyên được lọc**: Chỉ mục rất hữu ích khi bạn thường xuyên truy vấn các cột như `ID`, `email`, `name`, v.v.
- **Sử dụng trong các câu truy vấn phức tạp**: Đặc biệt là các câu truy vấn có nhiều `JOIN`, `GROUP BY` hoặc `ORDER BY`.

**Lưu ý khi sử dụng chỉ mục**:

- Chỉ mục không phải lúc nào cũng giúp cải thiện hiệu suất. Khi bảng có ít bản ghi hoặc truy vấn chỉ lấy ít dữ liệu, MySQL có thể không cần chỉ mục.
- Việc tạo và duy trì chỉ mục cũng tốn chi phí (thời gian và bộ nhớ), vì vậy chỉ nên tạo chỉ mục trên các cột truy vấn thường xuyên.

### 3. **Phân biệt giữa `MyISAM` và `InnoDB` trong MySQL. Khi nào thì chọn mỗi loại?**

- **MyISAM** và **InnoDB** là hai kiểu lưu trữ (storage engines) phổ biến trong MySQL, mỗi loại có ưu điểm và nhược điểm riêng.

**MyISAM**:

- **Ưu điểm**:
  - Lưu trữ và truy xuất dữ liệu nhanh hơn trong các truy vấn đọc (SELECT).
  - Nhẹ nhàng hơn về tài nguyên hệ thống.
- **Nhược điểm**:
  - Không hỗ trợ các giao dịch (transactions).
  - Không hỗ trợ `FOREIGN KEY` (không có tính toàn vẹn tham chiếu).
  - Không hỗ trợ khóa dòng (row-level locking), chỉ hỗ trợ khóa bảng.

**InnoDB**:

- **Ưu điểm**:
  - Hỗ trợ giao dịch (transactions) với tính năng ACID (Atomicity, Consistency, Isolation, Durability).
  - Hỗ trợ `FOREIGN KEY` và tính toàn vẹn tham chiếu.
  - Hỗ trợ khóa dòng (row-level locking) giúp cải thiện hiệu suất trong các tình huống truy cập đồng thời.
- **Nhược điểm**:
  - Có chi phí tài nguyên cao hơn so với MyISAM trong trường hợp không sử dụng các tính năng như giao dịch hoặc khóa dòng.

**Khi nào chọn mỗi loại**:

- **Chọn MyISAM**: Nếu ứng dụng của bạn chủ yếu thực hiện các truy vấn đọc (SELECT) và không cần tính năng giao dịch hoặc tính toàn vẹn tham chiếu.
- **Chọn InnoDB**: Nếu bạn cần hỗ trợ giao dịch, tính toàn vẹn tham chiếu, và các truy vấn đồng thời, hoặc ứng dụng của bạn cần sự chính xác cao trong các thay đổi dữ liệu (INSERT, UPDATE, DELETE).

### 4. **Cách sử dụng `EXPLAIN` trong MySQL để phân tích câu truy vấn.**

- **`EXPLAIN`** là một công cụ mạnh mẽ giúp bạn phân tích cách MySQL thực thi một câu truy vấn, bao gồm cách các bảng được truy vấn, loại join nào được sử dụng, và có sử dụng chỉ mục hay không.

**Cách sử dụng `EXPLAIN`**:

- Đặt từ khóa `EXPLAIN` trước câu truy vấn SQL của bạn.

  Ví dụ:

  ```sql
  EXPLAIN SELECT * FROM employees WHERE department = 'HR';
  ```

**Các cột thông tin quan trọng trong kết quả của `EXPLAIN`**:

- **id**: Thứ tự của câu truy vấn (có thể có nhiều câu truy vấn nếu có `JOIN`).
- **select_type**: Loại câu truy vấn, ví dụ: `SIMPLE` (truy vấn đơn), `PRIMARY` (truy vấn chính), `SUBQUERY` (truy vấn phụ).
- **table**: Tên bảng đang được truy vấn.
- **type**: Loại kết nối bảng, có thể là `ALL` (quét toàn bộ bảng), `index` (sử dụng chỉ mục), `range` (sử dụng phạm vi chỉ mục).
- **possible_keys**: Các chỉ mục có thể được sử dụng trong câu truy vấn.
- **key**: Chỉ mục thực tế được sử dụng trong câu truy vấn.
- **key_len**: Chiều dài của chỉ mục.
- **ref**: Cột hoặc hằng số được so sánh với chỉ mục.
- **rows**: Số lượng bản ghi mà MySQL phải quét.
- **Extra**: Thông tin bổ sung về cách truy vấn được thực thi, ví dụ: `Using where`, `Using index`.

**Ví dụ phân tích**:

```sql
EXPLAIN SELECT * FROM employees WHERE department = 'HR';
```

Kết quả có thể trông như sau:

```
+----+-------------+------------+-------+-------------------+---------+---------+-------+------+-------------+
| id | select_type | table      | type  | possible_keys     | key     | key_len | ref   | rows | Extra       |
+----+-------------+------------+-------+-------------------+---------+---------+-------+------+-------------+
| 1  | SIMPLE      | employees  | ref   | department_idx    | department_idx | 5   | const | 100  | Using where |
+----+-------------+------------+-------+-------------------+---------+---------+-------+------+-------------+
```

Trong ví dụ trên:

- `type: ref` cho thấy rằng MySQL sử dụng chỉ mục `department_idx` để truy vấn.
- `rows: 100` cho biết MySQL phải quét 100 bản ghi từ bảng `employees`.
- `Extra: Using where` cho biết MySQL sử dụng điều kiện `WHERE` để lọc kết quả sau khi truy vấn.
