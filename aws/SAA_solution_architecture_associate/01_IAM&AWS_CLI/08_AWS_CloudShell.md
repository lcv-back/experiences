[AWS CloudShell Docs](https://docs.aws.amazon.com/cloudshell/latest/userguide/supported-aws-regions.html)

---

### **AWS CloudShell Overview**

#### **What is AWS CloudShell?**

AWS CloudShell is a pre-authenticated, browser-based command-line interface integrated into the AWS Management Console. It allows you to run AWS CLI commands, scripts, and other tools without needing to install the AWS CLI locally or manage credentials on your machine.

#### **Key Features**

- **Pre-Configured**: Comes with the AWS CLI (version 2), common tools (e.g., `bash`, `python`, `jq`), and SDKs like Boto3 pre-installed.
- **Authentication**: Automatically uses your AWS Management Console credentials (IAM user or role), supporting MFA if enabled.
- **Persistent Storage**: `Provides 1 GB` of free storage per region for scripts and files, persisted across sessions.
- **Accessibility**: Available in supported regions directly from the browser, no local setup required.

#### **Use Cases**

- Quick AWS resource management (e.g., listing S3 buckets, starting EC2 instances).
- Testing scripts or commands without a local environment.
- Managing AWS resources on the go from any device with a browser.

#### **Limitations**

- 1 GB storage limit per region.
- Session timeout after 20 minutes of inactivity.
- Not all regions support CloudShell (check AWS documentation for availability).

---

### **Hands-On with AWS CloudShell**

#### **Step 1: Access AWS CloudShell**

1. **Log in to AWS Management Console**:
   - Go to `console.aws.amazon.com` and sign in with your IAM user or root account.
2. **Launch CloudShell**:
   - In the top-right corner of the console, click the **CloudShell icon** (a terminal symbol: `>_`).
   - If it’s your first time, it may take a minute to initialize. A terminal window will open at the bottom of the screen.
3. **Verify Environment**:
   - Check the AWS CLI version:
     ```bash
     aws --version
     ```
   - Expected output: `aws-cli/2.x.x ...`.
   - Confirm your identity:
     ```bash
     aws sts get-caller-identity
     ```
   - Output shows your AWS account ID, user ARN, etc.

#### **Step 2: Basic Commands**

1. **List S3 Buckets**:
   - Run:
     ```bash
     aws s3 ls
     ```
   - Output: Lists all buckets your IAM user has permission to see.
2. **Create a File and Upload to S3**:
   - Create a test file:
     ```bash
     echo "Hello from CloudShell" > testfile.txt
     ```
   - Upload to an existing bucket (replace `my-bucket`):
     ```bash
     aws s3 cp testfile.txt s3://my-bucket/
     ```
   - Verify:
     ```bash
     aws s3 ls s3://my-bucket/
     ```
3. **Download a File**:
   - Download the file back:
     ```bash
     aws s3 cp s3://my-bucket/testfile.txt downloaded.txt
     ```
   - View it:
     ```bash
     cat downloaded.txt
     ```

#### **Step 3: Work with Persistent Storage**

1. **Check Storage**:
   - List files in your home directory:
     ```bash
     ls -la
     ```
   - Files you create (e.g., `testfile.txt`) persist across sessions in the `/home/cloudshell-user` directory.
2. **Create a Script**:
   - Use the built-in editor:
     ```bash
     nano myscript.sh
     ```
   - Add a simple script:
     ```bash
     #!/bin/bash
     echo "Listing S3 buckets:"
     aws s3 ls
     ```
   - Save and exit (`Ctrl+O`, `Enter`, `Ctrl+X`).
3. **Run the Script**:
   - Make it executable:
     ```bash
     chmod +x myscript.sh
     ```
   - Execute:
     ```bash
     ./myscript.sh
     ```

#### **Step 4: Use CloudShell with MFA (If Enabled)**

- If your IAM user has MFA enabled and requires it for CLI access:
  - CloudShell uses your console session credentials. If MFA was provided during console login, no additional steps are needed.
  - If temporary credentials are required (e.g., for a role), you may need to assume a role manually:
    ```bash
    aws sts assume-role --role-arn arn:aws:iam::<account-id>:role/<role-name> --role-session-name test
    ```
    - Update environment variables with the returned `AccessKeyId`, `SecretAccessKey`, and `SessionToken`.

#### **Step 5: Explore Additional Features**

1. **Change Region**:
   - Check current region:
     ```bash
     aws configure get region
     ```
   - Set a new region (e.g., `us-west-2`):
     ```bash
     aws configure set region us-west-2
     ```
2. **Run Python with Boto3**:
   - Test Boto3 (pre-installed):
     ```bash
     python3
     ```
     ```python
     import boto3
     s3 = boto3.client('s3')
     buckets = s3.list_buckets()
     for bucket in buckets['Buckets']:
         print(bucket['Name'])
     exit()
     ```
3. **Customize the Shell**:
   - Full-screen mode: Click the **Actions** menu > **Full screen**.
   - Upload/Download files: Use the **Actions** menu to upload local files or download from CloudShell.

---

### **Additional Tips**

- **Permissions**: Commands fail if your IAM user lacks necessary permissions (e.g., `s3:ListBucket`). Update your IAM policy if needed.
- **Storage Cleanup**: Remove unused files with `rm <filename>` to manage the 1 GB limit.
- **Documentation**: Access AWS CLI help with `aws help` or specific service help like `aws s3 help`.
