# Tối ưu cơ sở dữ liệu tại công ty

## Trong mysql, có bao nhiêu loại chỉ mục ?

### Normal Index:

- Bình thường nhất, cơ bản
- Không có tác dụng ràng buộc nào, mục đích của nó là cải thiện hiệu quả truy vấn

### Unique Index:

- Hay sử dụng để làm email, primary key, không có tính trùng lặp trong database

### Primary Index (phổ biến nhất):

- Chỉ mục khóa chính

### Full-text Index:

- Chỉ mục toàn văn

### Composite Index (chỉ mục tổng hợp):

- Chỉ mục tổng hợp, chỉ mục chung, bao gồm nhiều cột
- Ví dụ trong bảng user: name, gender,...
- Khá quan trọng
- Trong thực tế, có nhiều câu lệnh phức tạp để query chứ không phải primary key index

## Nguyên tắc sử dụng index ngoài cùng bên trái ?

## Partition SQL

## 1. Đánh index
