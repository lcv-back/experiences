Here’s the hands-on guide for working with **IAM Users** and **Groups** in AWS, in **English only**. This walks you through creating an IAM user, creating a group, adding the user to the group, and attaching a policy using the AWS Management Console.

---

### **Hands-On with IAM Users and Groups**

#### **Step 1: Create an IAM User**

1. **Log in to AWS Management Console**:
   - Go to `console.aws.amazon.com` and sign in with your root or admin account.
2. **Navigate to IAM**:
   - In the search bar, type "IAM" and select it from the Services menu.
3. **Create a New User**:
   - In the left sidebar, click **Users** > **Add users**.
   - Enter a username (e.g., `Alice`).
   - Select access type:
     - **AWS Management Console access** (for browser login): Provide a custom password or let AWS generate one.
     - **Programmatic access** (for CLI/SDK): This generates access keys.
   - Click **Next: Permissions**.
4. **Skip Permissions for Now**:
   - We'll add permissions via a group later. Click **Next: Tags** (optional), then **Next: Review**, and **Create user**.
5. **Save Credentials**:
   - Download the `.csv` file with the user’s credentials (password or access keys) and store it securely.

#### **Step 2: Create an IAM Group**

1. **Go to Groups**:
   - In the IAM dashboard, click **Groups** > **Create new group**.
2. **Name the Group**:
   - Enter a group name (e.g., `Developers`).
3. **Attach a Policy**:
   - In the "Attach permissions policies" section, search for a policy (e.g., `AmazonS3ReadOnlyAccess`).
   - Select it and click **Next: Review**, then **Create group**.

#### **Step 3: Add User to Group**

1. **Open the Group**:
   - Click on the newly created group (e.g., `Developers`).
2. **Add Users**:
   - Go to the **Users** tab, click **Add users to group**, select `Alice`, and click **Add users**.
   - Now, Alice inherits the group’s permissions (e.g., S3 read-only access).

#### **Step 4: Test the Setup**

1. **Log in as the IAM User**:
   - Open an incognito browser window, go to your AWS sign-in URL (e.g., `https://<account-id>.signin.aws.amazon.com/console`), and log in with Alice’s credentials.
2. **Verify Permissions**:
   - Try accessing S3 (e.g., list buckets). Alice should have read-only access, but no write permissions.

---

### **Notes**

- Replace `AmazonS3ReadOnlyAccess` with any policy relevant to your use case (e.g., `AmazonEC2FullAccess`).
- Always follow the **least privilege principle**: grant only the permissions needed.
- If you encounter issues, ensure your admin account has sufficient permissions to manage IAM.
