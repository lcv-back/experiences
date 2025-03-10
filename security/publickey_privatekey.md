# Public key và private key

## 1. Authentication

- Trong bước kiểm tra người dùng đã từng đăng ký tài khoản chưa, có thể là chưa được ghi xuống cơ sở dữ liệu hoặc chưa được đăng ký thật
- Giải pháp tối ưu hơn: Tìm kiếm trong cache trước, vì cache được lưu trong RAM nếu không có thì sẽ kiểm tra trong database
