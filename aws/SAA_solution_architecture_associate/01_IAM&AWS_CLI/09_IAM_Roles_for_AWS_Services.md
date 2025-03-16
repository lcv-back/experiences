Below is an explanation of **IAM Roles for AWS Services**, including their purpose, how they work, and a hands-on example. I'll provide versions in both **English** and **Vietnamese**.

---

### **English Version: IAM Roles for AWS Services**

#### **What Are IAM Roles for AWS Services?**

IAM roles are identities in AWS that you can assume to grant permissions to AWS services, applications, or users without using long-term credentials like access keys. When used with AWS services, roles allow those services (e.g., EC2, Lambda, ECS) to securely interact with other AWS resources on your behalf.

#### **Key Concepts**

- **Trust Relationship**: Defines who or what can assume the role (e.g., an AWS service like `lambda.amazonaws.com`).
- **Permissions Policy**: Specifies what the role can do (e.g., read from S3, write to DynamoDB).
- **Temporary Credentials**: When a service assumes a role, AWS provides temporary security credentials (Access Key ID, Secret Access Key, Session Token) via the Security Token Service (STS).

#### **Why Use Roles for Services?**

- **Security**: Avoid hardcoding access keys in applications or services.
- **Flexibility**: Easily update permissions without redeploying code.
- **Least Privilege**: Grant only the permissions a service needs.

#### **Common Use Cases**

- **EC2 Instance**: An EC2 instance assumes a role to access S3 buckets.
- **Lambda Function**: A Lambda function uses a role to write logs to CloudWatch or query an RDS database.
- **ECS Tasks**: ECS tasks assume roles to interact with other AWS services.

#### **How It Works**

1. Create a role with a trust policy (e.g., allowing `ec2.amazonaws.com` to assume it).
2. Attach a permissions policy to the role (e.g., `AmazonS3ReadOnlyAccess`).
3. Assign the role to an AWS service (e.g., attach it to an EC2 instance).
4. The service uses STS to assume the role and receive temporary credentials.

---

### **Hands-On Example: IAM Role for an EC2 Instance**

#### **Step 1: Create an IAM Role**

1. **Log in to AWS Management Console**:
   - Go to `console.aws.amazon.com` and sign in.
2. **Navigate to IAM**:
   - Search for "IAM" and select it.
3. **Create a Role**:
   - Click **Roles** > **Create role**.
   - **Select Trusted Entity**: Choose **AWS service** > **EC2**.
   - Click **Next: Permissions**.
4. **Attach Permissions**:
   - Search for `AmazonS3ReadOnlyAccess`, select it, and click **Next: Tags** (optional).
   - Click **Next: Review**.
5. **Name the Role**:
   - Enter a name (e.g., `EC2S3ReadRole`) and click **Create role**.

#### **Step 2: Attach the Role to an EC2 Instance**

1. **Launch or Select an EC2 Instance**:
   - Go to **EC2** > **Instances**, launch a new instance, or select an existing one.
2. **Attach the Role**:
   - For a new instance: During launch, under **IAM role**, select `EC2S3ReadRole`.
   - For an existing instance:
     - Stop the instance (if running).
     - Right-click the instance > **Instance settings** > **Attach/Replace IAM role**.
     - Select `EC2S3ReadRole` and apply.
3. **Start the Instance**:
   - Start or reboot the instance if modified.

#### **Step 3: Test the Role**

1. **SSH into the EC2 Instance**:
   - Connect using SSH (e.g., `ssh -i key.pem ec2-user@<instance-ip>`).
2. **Verify Role Credentials**:
   - Run:
     ```bash
     curl http://169.254.169.254/latest/meta-data/iam/security-credentials/EC2S3ReadRole
     ```
   - Output shows temporary credentials (Access Key ID, Secret Access Key, Token).
3. **Test S3 Access**:
   - Install AWS CLI (if not present):
     ```bash
     sudo yum install awscli -y  # Amazon Linux
     ```
   - List S3 buckets:
     ```bash
     aws s3 ls
     ```
   - Should work without manual credential configuration due to the role.

---

### **Vietnamese Version: Vai Trò IAM cho Dịch Vụ AWS**

#### **Vai Trò IAM cho Dịch Vụ AWS Là Gì?**

Vai trò IAM là các danh tính trong AWS mà bạn có thể đảm nhận để cấp quyền cho các dịch vụ AWS, ứng dụng hoặc người dùng mà không cần sử dụng thông tin xác thực dài hạn như khóa truy cập. Khi dùng với dịch vụ AWS, vai trò cho phép các dịch vụ (ví dụ: EC2, Lambda, ECS) tương tác an toàn với các tài nguyên AWS khác thay mặt bạn.

#### **Khái Niệm Chính**

- **Quan Hệ Tin Cậy**: Xác định ai hoặc cái gì có thể đảm nhận vai trò (ví dụ: dịch vụ AWS như `lambda.amazonaws.com`).
- **Chính Sách Quyền**: Chỉ định vai trò có thể làm gì (ví dụ: đọc từ S3, ghi vào DynamoDB).
- **Thông Tin Tạm Thời**: Khi dịch vụ đảm nhận vai trò, AWS cung cấp thông tin bảo mật tạm thời (ID Khóa Truy Cập, Khóa Bí Mật, Mã Thông Báo Phiên) qua Dịch Vụ Token Bảo Mật (STS).

#### **Tại Sao Dùng Vai Trò cho Dịch Vụ?**

- **Bảo Mật**: Tránh mã hóa cứng khóa truy cập trong ứng dụng hoặc dịch vụ.
- **Linh Hoạt**: Dễ dàng cập nhật quyền mà không cần triển khai lại mã.
- **Quyền Tối Thiểu**: Chỉ cấp quyền mà dịch vụ cần.

#### **Trường Hợp Sử Dụng Phổ Biến**

- **Phiên Bản EC2**: Phiên bản EC2 đảm nhận vai trò để truy cập bucket S3.
- **Hàm Lambda**: Hàm Lambda dùng vai trò để ghi log vào CloudWatch hoặc truy vấn cơ sở dữ liệu RDS.
- **Tác Vụ ECS**: Tác vụ ECS đảm nhận vai trò để tương tác với dịch vụ AWS khác.

#### **Cách Hoạt Động**

1. Tạo vai trò với chính sách tin cậy (ví dụ: cho phép `ec2.amazonaws.com` đảm nhận).
2. Gắn chính sách quyền vào vai trò (ví dụ: `AmazonS3ReadOnlyAccess`).
3. Gán vai trò cho dịch vụ AWS (ví dụ: gắn vào phiên bản EC2).
4. Dịch vụ dùng STS để đảm nhận vai trò và nhận thông tin tạm thời.

---

### **Thực Hành: Vai Trò IAM cho Phiên Bản EC2**

#### **Bước 1: Tạo Vai Trò IAM**

1. **Đăng Nhập vào AWS Management Console**:
   - Truy cập `console.aws.amazon.com` và đăng nhập.
2. **Điều Hướng đến IAM**:
   - Tìm kiếm "IAM" và chọn nó.
3. **Tạo Vai Trò**:
   - Nhấp **Roles** > **Create role**.
   - **Chọn Thực Thể Tin Cậy**: Chọn **AWS service** > **EC2**.
   - Nhấp **Next: Permissions**.
4. **Gắn Quyền**:
   - Tìm `AmazonS3ReadOnlyAccess`, chọn nó, nhấp **Next: Tags** (tùy chọn).
   - Nhấp **Next: Review**.
5. **Đặt Tên Vai Trò**:
   - Nhập tên (ví dụ: `EC2S3ReadRole`) và nhấp **Create role**.

#### **Bước 2: Gắn Vai Trò vào Phiên Bản EC2**

1. **Khởi Tạo hoặc Chọn Phiên Bản EC2**:
   - Đi đến **EC2** > **Instances**, khởi tạo phiên bản mới hoặc chọn phiên bản hiện có.
2. **Gắn Vai Trò**:
   - Với phiên bản mới: Trong lúc khởi tạo, dưới **IAM role**, chọn `EC2S3ReadRole`.
   - Với phiên bản hiện có:
     - Dừng phiên bản (nếu đang chạy).
     - Nhấp chuột phải vào phiên bản > **Instance settings** > **Attach/Replace IAM role**.
     - Chọn `EC2S3ReadRole` và áp dụng.
3. **Khởi Động Phiên Bản**:
   - Khởi động hoặc khởi động lại phiên bản nếu đã sửa đổi.

#### **Bước 3: Kiểm Tra Vai Trò**

1. **SSH vào Phiên Bản EC2**:
   - Kết nối bằng SSH (ví dụ: `ssh -i key.pem ec2-user@<instance-ip>`).
2. **Xác Minh Thông Tin Vai Trò**:
   - Chạy:
     ```bash
     curl http://169.254.169.254/latest/meta-data/iam/security-credentials/EC2S3ReadRole
     ```
   - Đầu ra hiển thị thông tin tạm thời (ID Khóa Truy Cập, Khóa Bí Mật, Mã Thông Báo).
3. **Kiểm Tra Truy Cập S3**:
   - Cài đặt AWS CLI (nếu chưa có):
     ```bash
     sudo yum install awscli -y  # Amazon Linux
     ```
   - Liệt kê bucket S3:
     ```bash
     aws s3 ls
     ```
   - Sẽ hoạt động mà không cần cấu hình thông tin thủ công nhờ vai trò.

---
