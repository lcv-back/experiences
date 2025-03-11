# Cơ bản về cloud và aws

## What is on-premise ?

- Đây là cách truyền thống mà các công ty vận hành công nghệ của họ. Thay vì dùng dịch vụ từ internet, họ mua máy chủ, thiết bị lưu trữ, và phần mềm, rồi đặt chúng trong văn phòng hoặc trung tâm dữ liệu của chính mình. Họ phải tự quản lý mọi thứ.
- Trong AWS: AWS là một giải pháp thay thế cho on-premise. Với AWS, bạn không cần mua máy móc – bạn “thuê” chúng từ AWS qua internet. Ví dụ, thay vì mua máy chủ vật lý, bạn có thể dùng `Amazon EC2` (một dịch vụ của AWS) để chạy máy ảo.

### High upfront cost

- Khi bạn xây dựng hệ thống tại chỗ `(on-premise)`, bạn phải trả rất nhiều tiền ngay từ đầu để mua máy chủ, phần mềm, và thuê người cài đặt. Nếu nhu cầu thay đổi, bạn vẫn bị kẹt với những thứ đã mua.
- Trong AWS: AWS hoạt động theo kiểu `“trả tiền khi dùng” (pay-as-you-go)`. Bạn không cần trả trước nhiều – chỉ trả cho những gì bạn sử dụng, như thuê máy tính hoặc lưu trữ từ AWS. Ví dụ, với Amazon S3 (dịch vụ lưu trữ), bạn chỉ trả theo dung lượng bạn dùng.

### Capacity planing and over provisioning

- `Capacity Planning`: Là việc đoán xem bạn cần bao nhiêu máy móc hoặc tài nguyên để hệ thống chạy tốt. Nếu đoán sai, bạn có thể thiếu hoặc thừa.
- `Over-Provisioning`: Để an toàn, nhiều công ty mua nhiều tài nguyên hơn mức cần, dẫn đến lãng phí tiền bạc.
- Trong AWS: AWS giúp bạn tránh việc này. Với dịch vụ như Auto Scaling, hệ thống tự động tăng hoặc giảm tài nguyên (như máy chủ EC2) dựa trên nhu cầu thực tế. Bạn không cần đoán trước hay mua thừa nữa.

### Skill sds requirement

- Với hệ thống tại chỗ, bạn cần người biết lắp ráp máy chủ, sửa mạng, và quản lý phần mềm – những kỹ năng này tốn thời gian để học. Khi chuyển sang đám mây, bạn cần kỹ năng mới, như cách dùng các công cụ trên internet.
- Trong AWS: Để dùng AWS, bạn không cần biết sửa máy móc, nhưng bạn cần học cách sử dụng các dịch vụ của họ, như tạo máy ảo với EC2, lưu file trên S3.

## What is cloud ?

### Public cloud

- Là dịch vụ đám mây mà nhiều người hoặc công ty cùng dùng chung, nhưng dữ liệu của mỗi người được giữ riêng biệt. Nhà cung cấp (như AWS) quản lý mọi thứ.
- Trong AWS: AWS là một public cloud. Bất kỳ ai có tài khoản AWS đều có thể dùng các dịch vụ như S3 hoặc EC2. Dù là “công cộng”, dữ liệu của bạn vẫn an toàn nhờ các công cụ bảo mật như IAM.

### Private cloud

- Là một hệ thống đám mây chỉ dành riêng cho một công ty, không chia sẻ với ai. Nó giống như có một đám mây riêng tư, nhưng thường đắt hơn.
- Trong AWS: AWS hỗ trợ private cloud qua dịch vụ như `AWS Outposts` (đưa máy chủ AWS đến chỗ bạn) hoặc `VPC (Virtual Private Cloud)`, cho phép bạn tạo mạng riêng trong đám mây công cộng của AWS.

### Hybrid

- Là sự kết hợp giữa hệ thống tại chỗ (on-premise) và đám mây. Ví dụ, bạn giữ dữ liệu quan trọng trong văn phòng nhưng dùng đám mây để xử lý công việc lớn.
- Trong AWS: AWS hỗ trợ mô hình này bằng cách kết nối hệ thống tại chỗ của bạn với đám mây của họ. Ví dụ, AWS Direct Connect giúp bạn nối mạng nội bộ với AWS, hoặc Storage Gateway cho phép lưu file cả ở chỗ bạn và trên S3.
