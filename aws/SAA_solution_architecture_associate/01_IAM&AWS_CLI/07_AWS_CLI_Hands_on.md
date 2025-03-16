Below is a hands-on guide for working with the **AWS Command Line Interface (CLI)** in English. This guide covers installing the AWS CLI, configuring it with access keys, and performing common tasks like listing S3 buckets and uploading files.

---

### **Hands-On with AWS CLI**

#### **Step 1: Install the AWS CLI**

1. **Check Prerequisites**:
   - Ensure you have Python installed (version 3.6+ recommended) as the AWS CLI depends on it.
   - Verify with: `python --version` or `python3 --version`.
2. **Download and Install**:
   - **Windows**: Download the MSI installer from `aws.amazon.com/cli`, run it, and follow the prompts.
   - **macOS/Linux**: Use pip:
     ```bash
     pip3 install awscli
     ```
   - Alternative (Linux/macOS): Use the bundled installer:
     ```bash
     curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
     unzip awscliv2.zip
     sudo ./aws/install
     ```
3. **Verify Installation**:
   - Run:
     ```bash
     aws --version
     ```
   - Expected output: `aws-cli/2.x.x Python/3.x.x ...`.

#### **Step 2: Configure the AWS CLI**

1. **Obtain Access Keys**:
   - Log in to the AWS Management Console.
   - Go to **IAM** > **Users**, select a user (e.g., `Alice`), and under **Security credentials**, click **Create access key**.
   - Choose **CLI** as the use case, download the `.csv` file with the Access Key ID and Secret Access Key, and store it securely.
2. **Run Configuration**:
   - Open a terminal and run:
     ```bash
     aws configure
     ```
   - Enter the following when prompted:
     - **AWS Access Key ID**: (e.g., `AKIAIOSFODNN7EXAMPLE`)
     - **AWS Secret Access Key**: (e.g., `wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`)
     - **Default region name**: (e.g., `us-east-1`)
     - **Default output format**: (e.g., `json`, or leave blank for default).
3. **Verify Configuration**:
   - Check the credentials file: `~/.aws/credentials` (Linux/macOS) or `%USERPROFILE%\.aws\credentials` (Windows).
   - Check the config file: `~/.aws/config`.
   - Test with:
     ```bash
     aws sts get-caller-identity
     ```
   - Output shows your AWS account ID, user ARN, etc.

#### **Step 3: Perform Common AWS CLI Tasks**

1. **List S3 Buckets**:
   - Command:
     ```bash
     aws s3 ls
     ```
   - Output: Lists all S3 buckets in your account (e.g., `2023-01-01 12:00:00 my-bucket`).
   - Troubleshooting: If empty, ensure the user has `s3:ListAllMyBuckets` permission.
2. **Upload a File to S3**:
   - Create a test file:
     ```bash
     echo "Hello, AWS CLI!" > testfile.txt
     ```
   - Upload to an existing bucket (replace `my-bucket`):
     ```bash
     aws s3 cp testfile.txt s3://my-bucket/
     ```
   - Verify:
     ```bash
     aws s3 ls s3://my-bucket/
     ```
3. **Download a File from S3**:
   - Command:
     ```bash
     aws s3 cp s3://my-bucket/testfile.txt downloaded.txt
     ```
   - Check the downloaded file:
     ```bash
     cat downloaded.txt  # Linux/macOS
     type downloaded.txt  # Windows
     ```

#### **Step 4: Use AWS CLI with MFA (Optional)**

1. **Enable MFA**:
   - In IAM, assign an MFA device to the user (e.g., Google Authenticator).
2. **Get Temporary Credentials**:
   - Run:
     ```bash
     aws sts get-session-token --serial-number arn:aws:iam::<account-id>:mfa/<username> --token-code <6-digit-mfa-code>
     ```
   - Replace `<account-id>` with your AWS account ID, `<username>` with the IAM user (e.g., `Alice`), and `<6-digit-mfa-code>` with the code from your MFA app.
   - Output includes temporary `AccessKeyId`, `SecretAccessKey`, and `SessionToken`.
3. **Update CLI Credentials**:
   - Manually update `~/.aws/credentials` with the temporary keys or use:
     ```bash
     export AWS_ACCESS_KEY_ID=<temp-access-key>
     export AWS_SECRET_ACCESS_KEY=<temp-secret-key>
     export AWS_SESSION_TOKEN=<temp-session-token>
     ```
   - Test again: `aws sts get-caller-identity`.

#### **Step 5: Explore Additional Commands**

- **List EC2 Instances**:
  ```bash
  aws ec2 describe-instances --region us-east-1
  ```
- **Create an S3 Bucket**:
  ```bash
  aws s3 mb s3://my-new-bucket --region us-east-1
  ```
- **Get Help**:
  ```bash
  aws s3 help
  ```

---

### **Additional Tips**

- **Profiles**: Manage multiple accounts with `aws configure --profile myprofile` and use with `--profile myprofile`.
- **Permissions**: If a command fails (e.g., `AccessDenied`), check the IAM policy for the user.
- **Output Formats**: Change output to `text` or `table` with `--output text` or `--output table`.
