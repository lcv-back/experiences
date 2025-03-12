# Các câu hỏi phỏng vấn ở level intern, fresher

### I. Câu hỏi cơ bản về MySQL

### 1. **MySQL là gì? Những tính năng cơ bản của MySQL là gì?**

- **MySQL** là một hệ quản trị cơ sở dữ liệu mã nguồn mở, sử dụng SQL (Structured Query Language) để truy vấn và quản lý dữ liệu. Nó được phát triển và duy trì bởi Oracle Corporation. MySQL là một phần của bộ công cụ LAMP (Linux, Apache, MySQL, PHP/Python/Perl) và là một lựa chọn phổ biến trong các ứng dụng web.
- **Những tính năng cơ bản của MySQL**:
  - **Mã nguồn mở**: Miễn phí và có thể tùy chỉnh.
  - **Quản lý cơ sở dữ liệu**: Hỗ trợ các thao tác CRUD (Create, Read, Update, Delete).
  - **Hỗ trợ ACID**: Đảm bảo tính toàn vẹn dữ liệu khi thực hiện các giao dịch.
  - **Hỗ trợ Indexing**: Cải thiện hiệu suất truy vấn.
  - **Quản lý kết nối**: Hỗ trợ nhiều kết nối đồng thời.
  - **Quản lý dữ liệu quan hệ**: Là hệ quản trị cơ sở dữ liệu quan hệ (RDBMS), dữ liệu được tổ chức trong bảng và có thể liên kết với nhau.

### 2. **Phân biệt giữa `INNER JOIN`, `LEFT JOIN`, và `RIGHT JOIN` trong MySQL.**

- **INNER JOIN**: Chỉ trả về các bản ghi có sự trùng khớp trong cả hai bảng. Nếu một bản ghi trong bảng A không có bản ghi tương ứng trong bảng B (và ngược lại), thì bản ghi đó không xuất hiện trong kết quả.

  Ví dụ:

  ```sql
  SELECT * FROM A INNER JOIN B ON A.id = B.id;
  ```

- **LEFT JOIN (hoặc LEFT OUTER JOIN)**: Trả về tất cả các bản ghi từ bảng bên trái (bảng A) và các bản ghi khớp từ bảng bên phải (bảng B). Nếu không có sự khớp, bảng bên phải sẽ có giá trị NULL.

  Ví dụ:

  ```sql
  SELECT * FROM A LEFT JOIN B ON A.id = B.id;
  ```

- **RIGHT JOIN (hoặc RIGHT OUTER JOIN)**: Tương tự như LEFT JOIN, nhưng trả về tất cả các bản ghi từ bảng bên phải (bảng B) và các bản ghi khớp từ bảng bên trái (bảng A). Nếu không có sự khớp, bảng bên trái sẽ có giá trị NULL.

  Ví dụ:

  ```sql
  SELECT * FROM A RIGHT JOIN B ON A.id = B.id;
  ```

### 3. **Giải thích sự khác nhau giữa `PRIMARY KEY` và `UNIQUE` trong MySQL.**

- **PRIMARY KEY**:
  - Là một ràng buộc (constraint) chỉ có thể có một trong mỗi bảng.
  - Mỗi giá trị của trường `PRIMARY KEY` phải là duy nhất và không được NULL.
  - Thường được sử dụng để xác định bản ghi duy nhất trong bảng.
- **UNIQUE**:
  - Cũng là một ràng buộc giúp đảm bảo rằng các giá trị trong một cột là duy nhất.
  - Một bảng có thể có nhiều ràng buộc `UNIQUE` nhưng chỉ có một `PRIMARY KEY`.
  - Khác với `PRIMARY KEY`, trường `UNIQUE` có thể chứa giá trị NULL (trong một số trường hợp).

### 4. **Câu lệnh `SELECT * FROM table_name` có hiệu quả như thế nào? Tại sao không nên sử dụng `SELECT *` trong các ứng dụng thực tế?**

- **Câu lệnh `SELECT * FROM table_name`**: Trả về tất cả các cột của tất cả các bản ghi trong bảng. Dễ dàng sử dụng trong việc thử nghiệm hoặc khi bạn muốn lấy toàn bộ dữ liệu từ bảng.
- **Vấn đề với `SELECT *`**:
  - **Hiệu suất kém**: Khi bảng có nhiều cột, việc lấy tất cả cột có thể làm giảm hiệu suất, đặc biệt khi chỉ cần một phần nhỏ dữ liệu.
  - **Tính bảo mật**: Nếu có các cột nhạy cảm như mật khẩu hoặc thông tin cá nhân, việc sử dụng `SELECT *` có thể vô tình lấy ra thông tin không cần thiết.
  - **Dễ gây lỗi trong phát triển**: Khi bảng thay đổi (thêm cột), `SELECT *` sẽ tự động lấy các cột mới mà bạn có thể không cần thiết.

### 5. **Cách sử dụng câu lệnh `GROUP BY` trong MySQL. Ví dụ về khi nào sử dụng `HAVING` thay vì `WHERE`.**

- **`GROUP BY`**: Dùng để nhóm các bản ghi có giá trị giống nhau ở một hoặc nhiều cột. Thường được sử dụng với các hàm tổng hợp như `COUNT()`, `SUM()`, `AVG()`, `MAX()`, `MIN()`.

  Ví dụ:

  ```sql
  SELECT department, COUNT(*) FROM employees GROUP BY department;
  ```

- **`HAVING` vs `WHERE`**:

  - **`WHERE`**: Dùng để lọc dữ liệu trước khi nhóm. Không thể sử dụng với các hàm tổng hợp.
  - **`HAVING`**: Dùng để lọc dữ liệu sau khi nhóm. Thường dùng khi bạn muốn áp dụng điều kiện cho các nhóm (ví dụ: lọc các nhóm có tổng số lớn hơn một giá trị nhất định).

  Ví dụ:

  ```sql
  SELECT department, COUNT(*) FROM employees
  GROUP BY department
  HAVING COUNT(*) > 5;
  ```

### 6. **Giải thích khái niệm `Normalization` trong cơ sở dữ liệu và các dạng chuẩn của nó (1NF, 2NF, 3NF).**

- **Normalization**: Là quá trình tổ chức dữ liệu trong cơ sở dữ liệu sao cho tránh được sự dư thừa và đảm bảo tính toàn vẹn của dữ liệu. Nó chia cơ sở dữ liệu thành các bảng nhỏ hơn và kết nối chúng lại với nhau thông qua các khóa ngoại.
- **1NF (First Normal Form)**:
  - Mỗi cột chứa các giá trị nguyên tử (không chia nhỏ hơn).
  - Mỗi bản ghi trong bảng phải có một giá trị duy nhất.
- **2NF (Second Normal Form)**:
  - Đảm bảo rằng bảng đã tuân thủ 1NF.
  - Mỗi cột không khóa phải phụ thuộc hoàn toàn vào khóa chính (không có phụ thuộc một phần).
- **3NF (Third Normal Form)**:
  - Đảm bảo rằng bảng đã tuân thủ 2NF.
  - Mỗi cột không khóa không được phụ thuộc vào các cột không khóa khác (không có phụ thuộc bắc cầu).
