# AWS Budget Setup

### AWS Budget là gì?

AWS Budgets là một công cụ trong AWS Management Console giúp bạn theo dõi và quản lý chi phí sử dụng dịch vụ AWS. Bạn có thể thiết lập ngân sách để đặt giới hạn chi tiêu, nhận cảnh báo khi chi phí vượt ngưỡng, và tối ưu hóa việc sử dụng tài nguyên.

### Các bước thiết lập AWS Budget

Dưới đây là hướng dẫn từng bước để thiết lập ngân sách trên AWS:

#### 1. **Truy cập AWS Budgets**

- Đăng nhập vào AWS Management Console.
- Tìm **AWS Budgets** trong thanh tìm kiếm hoặc vào menu **Billing and Cost Management** → chọn **Budgets**.

#### 2. **Tạo ngân sách mới**

- Nhấn nút **Create budget**.
- Chọn loại ngân sách phù hợp:
  - **Cost Budget**: Theo dõi tổng chi phí (phổ biến nhất).
  - **Usage Budget**: Theo dõi mức sử dụng của một dịch vụ cụ thể (ví dụ: giờ sử dụng EC2).
  - **Reservation Budget**: Theo dõi việc sử dụng Reserved Instances hoặc Savings Plans.
  - **Savings Plans Budget**: Tập trung vào Savings Plans.

#### 3. **Đặt thông số ngân sách**

- **Tên ngân sách**: Đặt tên dễ nhận biết (ví dụ: "WebServer_Monthly_Budget").
- **Thời gian**: Chọn khoảng thời gian áp dụng ngân sách:
  - Monthly (hàng tháng).
  - Quarterly (hàng quý).
  - Annually (hàng năm).
- **Số tiền ngân sách**: Nhập số tiền tối đa bạn muốn chi (tính bằng USD). Ví dụ: $100/tháng.
- **Phạm vi áp dụng**: Bạn có thể áp dụng ngân sách cho toàn bộ tài khoản hoặc lọc theo:
  - Dịch vụ cụ thể (EC2, S3, Lambda, v.v.).
  - Tags (ví dụ: chi phí của một dự án cụ thể).
  - Region (khu vực).

#### 4. **Thiết lập cảnh báo (Alerts)**

- Nhấn **Configure alerts** để thêm ngưỡng cảnh báo.
- Ví dụ:
  - Cảnh báo 1: Gửi email khi chi phí đạt 80% ngân sách ($80 nếu ngân sách là $100).
  - Cảnh báo 2: Gửi email khi chi phí vượt 100% ($100).
- Nhập địa chỉ email hoặc kết nối với SNS (Simple Notification Service) để nhận thông báo.

#### 5. **Xem lại và hoàn tất**

- Kiểm tra lại các thông số (số tiền, ngưỡng cảnh báo, email).
- Nhấn **Create** để hoàn tất.

### Một số mẹo khi dùng AWS Budgets

- **Kiểm tra thường xuyên**: AWS Budgets không tự động dừng dịch vụ khi vượt ngân sách, chỉ gửi cảnh báo. Bạn cần chủ động kiểm tra và điều chỉnh.
- **Kết hợp với Cost Explorer**: Dùng Cost Explorer để phân tích chi tiết chi phí trước khi đặt ngân sách.
- **Dùng Tags**: Gắn thẻ cho tài nguyên để quản lý chi phí theo dự án hoặc nhóm.

### Ví dụ thực tế

Giả sử bạn muốn giới hạn chi phí EC2 ở mức $50/tháng:

1. Tạo Cost Budget.
2. Đặt tên: "EC2_March_2025".
3. Chọn phạm vi: Lọc theo dịch vụ "EC2".
4. Đặt ngân sách: $50.
5. Thêm cảnh báo: 80% ($40) và 100% ($50).
6. Nhập email cá nhân để nhận thông báo.

Nếu bạn cần thêm chi tiết hoặc muốn tôi phân tích một kịch bản cụ thể, cứ hỏi nhé!
