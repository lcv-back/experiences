Below is a hands-on guide for working with **IAM Roles for AWS Services**, focusing on creating and using an IAM role with an AWS service (specifically an EC2 instance). This guide is in English and walks you through the process step-by-step using the AWS Management Console and testing the setup.

---

### **Hands-On with IAM Roles for AWS Services**

#### **Objective**

Create an IAM role that allows an EC2 instance to access an S3 bucket with read-only permissions, attach it to an EC2 instance, and test the functionality.

---

#### **Step 1: Create an IAM Role for EC2**

1. **Log in to AWS Management Console**:
   - Go to `console.aws.amazon.com` and sign in with an account that has IAM and EC2 permissions.
2. **Navigate to IAM**:
   - In the search bar, type "IAM" and select it from the Services menu.
3. **Create a New Role**:
   - Click **Roles** in the left sidebar, then **Create role**.
4. **Select Trusted Entity**:
   - Choose **AWS service** as the trusted entity type.
   - Under "Common use cases," select **EC2** (this allows EC2 instances to assume the role).
   - Click **Next: Permissions**.
5. **Attach Permissions Policy**:
   - Search for `AmazonS3ReadOnlyAccess` in the policy list.
   - Check the box next to it and click **Next: Tags** (optional, you can skip tags for now).
   - Click **Next: Review**.
6. **Name and Create the Role**:
   - Enter a role name (e.g., `EC2S3ReadRole`).
   - Optionally, add a description (e.g., "Allows EC2 instances to read from S3").
   - Click **Create role**.
7. **Verify the Role**:
   - Find `EC2S3ReadRole` in the Roles list. Click it to view:
     - **Trust relationship**: Should include `ec2.amazonaws.com`.
     - **Permissions**: Should show `AmazonS3ReadOnlyAccess`.

---

#### **Step 2: Launch an EC2 Instance with the Role**

1. **Navigate to EC2**:
   - In the AWS Console, search for "EC2" and select it.
2. **Launch an Instance**:
   - Click **Launch instance** > **Launch instance**.
   - **Name**: Enter a name (e.g., `S3ReadInstance`).
   - **AMI**: Choose "Amazon Linux 2 AMI" (free tier eligible).
   - **Instance Type**: Select `t2.micro` (free tier eligible).
   - **Key Pair**: Create or select an existing key pair for SSH access (e.g., `my-key.pem`).
3. **Configure IAM Role**:
   - Under **Advanced details** (expand if needed), find **IAM instance profile**.
   - Select `EC2S3ReadRole` from the dropdown.
4. **Complete Launch**:
   - Leave other settings as default (or configure a VPC/security group if needed).
   - Click **Launch instance**.
5. **Verify Instance**:
   - Go to **Instances**, find `S3ReadInstance`, and wait for it to enter the "Running" state.

---

#### **Step 3: Test the Role on the EC2 Instance**

1. **SSH into the EC2 Instance**:
   - Get the public IP from the EC2 Instances page.
   - Open a terminal and connect:
     ```bash
     ssh -i my-key.pem ec2-user@<public-ip>
     ```
     - Replace `my-key.pem` with your key file and `<public-ip>` with the instance’s IP.
2. **Check Role Credentials**:
   - Run:
     ```bash
     curl http://169.254.169.254/latest/meta-data/iam/security-credentials/EC2S3ReadRole
     ```
   - Output: JSON with temporary credentials (e.g., `AccessKeyId`, `SecretAccessKey`, `Token`).
   - This confirms the instance has assumed the role via the EC2 metadata service.
3. **Install AWS CLI (if needed)**:
   - On Amazon Linux 2, install the AWS CLI:
     ```bash
     sudo yum install awscli -y
     ```
   - Verify:
     ```bash
     aws --version
     ```
4. **Test S3 Access**:
   - List all S3 buckets:
     ```bash
     aws s3 ls
     ```
     - Should succeed if your account has buckets and the role’s permissions are correct.
   - Access a specific bucket (replace `my-bucket` with an existing bucket you have):
     ```bash
     aws s3 ls s3://my-bucket/
     ```
     - Should list objects if the bucket exists and has objects.
   - Try writing (should fail):
     ```bash
     echo "Test" > testfile.txt
     aws s3 cp testfile.txt s3://my-bucket/
     ```
     - Expected error: `AccessDenied` (since the role only allows read access).

---

#### **Step 4: Modify the Role (Optional)**

1. **Add Write Permissions**:
   - In IAM, go to **Roles** > `EC2S3ReadRole` > **Permissions** tab.
   - Click **Attach policies**, search for `AmazonS3FullAccess`, select it, and attach.
2. **Retest on EC2**:
   - SSH back into the instance and retry the upload:
     ```bash
     aws s3 cp testfile.txt s3://my-bucket/
     ```
   - Should now succeed.
3. **Revert (Best Practice)**:
   - Detach `AmazonS3FullAccess` and keep `AmazonS3ReadOnlyAccess` to enforce least privilege.

---

#### **Step 5: Cleanup**

1. **Terminate the Instance**:
   - In EC2, select `S3ReadInstance`, click **Instance state** > **Terminate instance**.
2. **Delete the Role (Optional)**:
   - In IAM, go to **Roles**, select `EC2S3ReadRole`, and click **Delete** if no longer needed.

---

### **Additional Notes**

- **Trust Policy**: The default EC2 trust policy looks like:
  ```json
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Principal": {
          "Service": "ec2.amazonaws.com"
        },
        "Action": "sts:AssumeRole"
      }
    ]
  }
  ```
- **Permissions Troubleshooting**: If `aws s3 ls` fails with `AccessDenied`, ensure the bucket exists and the role’s policy matches the bucket ARN.
- **Alternative Services**: Use the same process for Lambda (select `lambda.amazonaws.com` as the trusted entity) or other services.
