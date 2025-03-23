### EC2 là gì?

Amazon EC2 là một dịch vụ của AWS cung cấp khả năng tính toán (compute) trên đám mây. Nói đơn giản, nó cho phép bạn thuê máy chủ ảo (gọi là **instance**) để chạy ứng dụng, lưu trữ dữ liệu hoặc thực hiện các tác vụ tính toán mà không cần mua phần cứng vật lý.

### Các khái niệm cơ bản

1. **Instance**: Một máy chủ ảo trong EC2. Bạn có thể tưởng tượng nó như một máy tính chạy trên đám mây.
2. **Instance Type**: Loại cấu hình của instance, quyết định CPU, RAM, dung lượng lưu trữ, v.v. Ví dụ:
   - `t2.micro`: Nhẹ, rẻ, phù hợp cho thử nghiệm.
   - `m5.large`: Cân bằng, dùng cho ứng dụng vừa.
   - `c5.xlarge`: Tối ưu cho tính toán nặng.
3. **AMI (Amazon Machine Image)**: Hình ảnh hệ điều hành và phần mềm cài sẵn (ví dụ: Ubuntu, Windows, Amazon Linux).
4. **Region và Availability Zone**: Khu vực địa lý và trung tâm dữ liệu nơi instance được chạy (ví dụ: us-east-1, ap-southeast-1).
5. **EBS (Elastic Block Store)**: Ổ cứng ảo gắn vào instance để lưu dữ liệu.

### Cách hoạt động cơ bản

1. **Khởi tạo instance**:
   - Chọn AMI (hệ điều hành).
   - Chọn loại instance (t2.micro miễn phí trong Free Tier).
   - Cấu hình mạng (VPC, subnet, security group).
   - Thêm key pair (file .pem) để truy cập qua SSH (Linux) hoặc RDP (Windows).
2. **Truy cập instance**:
   - Dùng SSH (Linux): `ssh -i key.pem ec2-user@<public-ip>`.
   - Dùng Remote Desktop (Windows).
3. **Sử dụng**: Cài ứng dụng, chạy web server, hoặc bất kỳ tác vụ nào bạn muốn.
4. **Dừng/Kết thúc**: Dừng instance để tiết kiệm chi phí hoặc xóa khi không cần nữa.

### Chi phí

- EC2 tính phí theo giờ sử dụng (hoặc giây, tùy loại instance).
- **Free Tier**: 750 giờ/tháng cho t2.micro trong năm đầu tiên.
- Giá thay đổi theo loại instance, khu vực, và cách thanh toán (On-Demand, Reserved, Spot).

### Ví dụ đơn giản

Bạn muốn chạy một website nhỏ:

1. Khởi tạo instance `t2.micro` với AMI Ubuntu.
2. Cài web server (Apache/Nginx) qua SSH.
3. Trỏ tên miền đến Public IP của instance.
4. Theo dõi chi phí qua AWS Budgets.

### Lợi ích của EC2

- **Linh hoạt**: Tăng giảm tài nguyên dễ dàng.
- **Tiết kiệm**: Chỉ trả cho những gì bạn dùng.
- **Mở rộng**: Thêm instance khi tải tăng (dùng cùng Auto Scaling).
