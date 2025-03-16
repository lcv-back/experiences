Below is an explanation of **AWS Access Keys**, **AWS CLI**, and **AWS SDK**, including their roles, usage, and how they integrate for managing AWS resources. I'll provide versions in both **English** and **Vietnamese**.

---

### **English Version: AWS Access Keys, CLI, and SDK**

#### **AWS Access Keys**

- **What Are They?**: Access keys are long-term credentials used for programmatic access to AWS services. They consist of:
  - **Access Key ID**: A 20-character alphanumeric string (e.g., `AKIAIOSFODNN7EXAMPLE`).
  - **Secret Access Key**: A 40-character string (e.g., `wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`).
- **Purpose**: Authenticate API requests via the AWS Command Line Interface (CLI), Software Development Kits (SDKs), or custom applications.
- **Creation**:
  - Generated for IAM users in the AWS Management Console under **Users** > **Security credentials** > **Create access key**.
  - Download the `.csv` file containing the keys and store it securely (you can’t retrieve the secret key later).
- **Best Practices**:
  - Rotate keys regularly (e.g., every 90 days).
  - Use IAM roles with temporary credentials instead of access keys for EC2 instances or Lambda functions.
  - Never share or hardcode keys in code.

#### **AWS CLI (Command Line Interface)**

- **What Is It?**: A tool to manage AWS services from a terminal using commands.
- **Installation**:
  - Download from `aws.amazon.com/cli` and install (available for Windows, macOS, Linux).
  - Verify with `aws --version`.
- **Configuration**:
  - Run `aws configure` and enter:
    - Access Key ID
    - Secret Access Key
    - Default region (e.g., `us-east-1`)
    - Output format (e.g., `json`)
  - Credentials are stored in `~/.aws/credentials` and config in `~/.aws/config`.
- **Usage Example**:
  - List S3 buckets: `aws s3 ls`
  - Upload a file: `aws s3 cp myfile.txt s3://my-bucket/`
- **MFA Support**: For MFA-protected accounts, use `aws sts get-session-token` with an MFA code to generate temporary credentials.

#### **AWS SDK (Software Development Kit)**

- **What Is It?**: Libraries for programming languages (e.g., Python, Java, Node.js) to interact with AWS services programmatically.
- **Installation**:
  - Example for Python (Boto3): `pip install boto3`
- **Configuration**:
  - Uses the same credentials as the CLI (from `~/.aws/credentials`) or environment variables (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`).
- **Usage Example (Python)**:
  ```python
  import boto3
  s3 = boto3.client('s3')
  buckets = s3.list_buckets()
  for bucket in buckets['Buckets']:
      print(bucket['Name'])
  ```
- **Benefits**:
  - Simplifies API calls with native language support.
  - Supports IAM roles and temporary credentials for secure access.

#### **Integration**

- **Access Keys**: Provide the authentication for CLI and SDK.
- **CLI**: Ideal for quick tasks, scripting, or automation.
- **SDK**: Best for building applications with AWS integration.
- Example: Use CLI to test a command (`aws s3 ls`), then embed it in an SDK-based app for scalability.

---

### **Vietnamese Version: Khóa Truy Cập AWS, CLI và SDK**

#### **Khóa Truy Cập AWS (AWS Access Keys)**

- **Chúng Là Gì?**: Khóa truy cập là thông tin xác thực dài hạn dùng để truy cập lập trình vào các dịch vụ AWS. Bao gồm:
  - **ID Khóa Truy Cập**: Chuỗi 20 ký tự chữ và số (ví dụ: `AKIAIOSFODNN7EXAMPLE`).
  - **Khóa Bí Mật**: Chuỗi 40 ký tự (ví dụ: `wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`).
- **Mục Đích**: Xác thực các yêu cầu API qua AWS CLI, SDK hoặc ứng dụng tùy chỉnh.
- **Tạo Khóa**:
  - Tạo cho người dùng IAM trong AWS Management Console tại **Users** > **Security credentials** > **Create access key**.
  - Tải xuống tệp `.csv` chứa khóa và lưu trữ an toàn (không thể lấy lại khóa bí mật sau đó).
- **Thực Hành Tốt Nhất**:
  - Thay đổi khóa định kỳ (ví dụ: mỗi 90 ngày).
  - Sử dụng vai trò IAM với thông tin tạm thời thay vì khóa truy cập cho EC2 hoặc Lambda.
  - Không chia sẻ hoặc mã hóa cứng khóa trong mã nguồn.

#### **AWS CLI (Giao Diện Dòng Lệnh)**

- **Nó Là Gì?**: Công cụ để quản lý dịch vụ AWS từ terminal bằng lệnh.
- **Cài Đặt**:
  - Tải từ `aws.amazon.com/cli` và cài đặt (hỗ trợ Windows, macOS, Linux).
  - Kiểm tra bằng `aws --version`.
- **Cấu Hình**:
  - Chạy `aws configure` và nhập:
    - ID Khóa Truy Cập
    - Khóa Bí Mật
    - Vùng mặc định (ví dụ: `us-east-1`)
    - Định dạng đầu ra (ví dụ: `json`)
  - Thông tin xác thực được lưu trong `~/.aws/credentials` và cấu hình trong `~/.aws/config`.
- **Ví Dụ Sử Dụng**:
  - Liệt kê bucket S3: `aws s3 ls`
  - Tải lên tệp: `aws s3 cp myfile.txt s3://my-bucket/`
- **Hỗ Trợ MFA**: Với tài khoản bảo vệ bằng MFA, dùng `aws sts get-session-token` với mã MFA để tạo thông tin tạm thời.

#### **AWS SDK (Bộ Công Cụ Phát Triển Phần Mềm)**

- **Nó Là Gì?**: Thư viện cho các ngôn ngữ lập trình (ví dụ: Python, Java, Node.js) để tương tác lập trình với dịch vụ AWS.
- **Cài Đặt**:
  - Ví dụ cho Python (Boto3): `pip install boto3`
- **Cấu Hình**:
  - Sử dụng cùng thông tin xác thực với CLI (từ `~/.aws/credentials`) hoặc biến môi trường (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`).
- **Ví Dụ Sử Dụng (Python)**:
  ```python
  import boto3
  s3 = boto3.client('s3')
  buckets = s3.list_buckets()
  for bucket in buckets['Buckets']:
      print(bucket['Name'])
  ```
- **Lợi Ích**:
  - Đơn giản hóa gọi API với hỗ trợ ngôn ngữ tự nhiên.
  - Hỗ trợ vai trò IAM và thông tin tạm thời để truy cập an toàn.

#### **Tích Hợp**

- **Khóa Truy Cập**: Cung cấp xác thực cho CLI và SDK.
- **CLI**: Phù hợp cho tác vụ nhanh, viết script hoặc tự động hóa.
- **SDK**: Tốt nhất để xây dựng ứng dụng tích hợp AWS.
- Ví dụ: Dùng CLI để thử lệnh (`aws s3 ls`), sau đó nhúng vào ứng dụng dựa trên SDK để mở rộng.

---
