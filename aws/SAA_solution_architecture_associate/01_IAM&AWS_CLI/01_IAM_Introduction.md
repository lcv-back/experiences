# Users, Groups, Policies

Dưới đây là hai phiên bản giới thiệu về IAM (users, groups, policies) bằng tiếng Anh và tiếng Việt:

---

### **English Version**

In AWS Identity and Access Management (IAM), the concepts of users, groups, and policies are foundational to managing access to AWS resources securely. Here's a breakdown of each:

#### **IAM Users**

An IAM user represents an individual or entity (like an application) that interacts with AWS services. Each user has a unique identity and can be assigned specific permissions to perform actions (e.g., launching an EC2 instance or accessing an S3 bucket).

- **Key Features**:
  - Users are assigned credentials, such as a password for console access or access keys for programmatic access (e.g., via CLI or SDK).
  - Best practice: Avoid using the root account for daily tasks; create IAM users instead.
  - Example: A developer named "Alice" might have an IAM user with credentials to access specific AWS resources.

#### **IAM Groups**

An IAM group is a collection of IAM users. Groups make it easier to manage permissions for multiple users with similar roles or responsibilities.

- **Key Features**:
  - Permissions are assigned to the group via policies, and all users in the group inherit those permissions.
  - Users can belong to multiple groups, allowing flexible permission management.
  - Example: A "Developers" group might include Alice and Bob, granting them access to development-related resources.

#### **IAM Policies**

IAM policies are JSON documents that define permissions—what actions are allowed or denied, and on which AWS resources. Policies are attached to users, groups, or roles.

- **Key Features**:
  - Policies can be managed (AWS-managed or customer-managed) or inline (embedded directly in a user, group, or role).
  - They follow a least-privilege principle: grant only the permissions needed.
  - Example: A policy might allow a user to list objects in an S3 bucket named "example-bucket."

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "s3:ListBucket",
      "Resource": "arn:aws:s3:::example-bucket"
    }
  ]
}
```

---

### **Vietnamese Version**

#### **Người Dùng IAM (IAM Users)**

Người dùng IAM đại diện cho một cá nhân hoặc thực thể (như một ứng dụng) tương tác với các dịch vụ AWS. Mỗi người dùng có một danh tính duy nhất và có thể được cấp quyền cụ thể để thực hiện các hành động (ví dụ: khởi chạy một phiên bản EC2 hoặc truy cập vào một bucket S3).

- **Đặc Điểm Chính**:
  - Người dùng được cấp thông tin đăng nhập, chẳng hạn như mật khẩu để truy cập giao diện điều khiển hoặc khóa truy cập để truy cập lập trình (ví dụ: qua CLI hoặc SDK).
  - Thực hành tốt nhất: Tránh sử dụng tài khoản root cho các tác vụ hàng ngày; thay vào đó, hãy tạo người dùng IAM.
  - Ví dụ: Một lập trình viên tên "Alice" có thể có một người dùng IAM với thông tin đăng nhập để truy cập các tài nguyên AWS cụ thể.

#### **Nhóm IAM (IAM Groups)**

Nhóm IAM là tập hợp của các người dùng IAM. Nhóm giúp việc quản lý quyền cho nhiều người dùng có vai trò hoặc trách nhiệm tương tự trở nên dễ dàng hơn.

- **Đặc Điểm Chính**:
  - Quyền được gán cho nhóm thông qua các chính sách, và tất cả người dùng trong nhóm sẽ kế thừa những quyền đó.
  - Người dùng có thể thuộc nhiều nhóm, cho phép quản lý quyền linh hoạt.
  - Ví dụ: Nhóm "Developers" (Lập trình viên) có thể bao gồm Alice và Bob, cấp cho họ quyền truy cập vào các tài nguyên liên quan đến phát triển.

#### **Chính Sách IAM (IAM Policies)**

Chính sách IAM là các tài liệu JSON định nghĩa quyền—những hành động nào được phép hoặc bị từ chối, và trên các tài nguyên AWS nào. Chính sách được gắn vào người dùng, nhóm hoặc vai trò.

- **Đặc Điểm Chính**:
  - Chính sách có thể được quản lý (do AWS quản lý hoặc do khách hàng quản lý) hoặc nội tuyến (được nhúng trực tiếp vào người dùng, nhóm hoặc vai trò).
  - Chúng tuân theo nguyên tắc quyền tối thiểu: chỉ cấp những quyền cần thiết.
  - Ví dụ: Một chính sách có thể cho phép người dùng liệt kê các đối tượng trong một bucket S3 có tên "example-bucket."

---
