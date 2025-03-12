### I. **Câu hỏi về lý thuyết cơ bản**

### 1. **Đa luồng là gì?**

**Đa luồng (Multithreading)** là kỹ thuật lập trình cho phép một chương trình thực hiện nhiều tác vụ song song bằng cách sử dụng nhiều luồng (thread) khác nhau. Mỗi thread có thể thực hiện một tác vụ riêng biệt, nhưng chúng chia sẻ tài nguyên chung như bộ nhớ và dữ liệu của ứng dụng.

**Tại sao nó quan trọng?**  
Đa luồng giúp tối ưu hóa hiệu suất, đặc biệt trong các hệ thống đa lõi (multi-core systems), khi các tác vụ có thể được xử lý đồng thời thay vì phải thực hiện lần lượt. Điều này làm giảm thời gian chờ và tăng tốc độ thực thi ứng dụng. Ngoài ra, nó còn giúp giảm độ trễ trong các ứng dụng yêu cầu xử lý nhiều tác vụ như trong game, web server, hoặc xử lý dữ liệu lớn.

---

### 2. **Process và Thread có gì khác nhau?**

- **Process** là một đơn vị thực thi độc lập, có không gian bộ nhớ riêng biệt. Mỗi process có một vùng nhớ riêng biệt và không thể chia sẻ bộ nhớ trực tiếp với các process khác, trừ khi sử dụng các cơ chế như IPC (Inter-Process Communication). Mỗi process có ít nhất một thread.
- **Thread** là một đơn vị nhỏ hơn của tiến trình, thực thi một phần công việc trong quá trình. Các thread trong cùng một process chia sẻ không gian bộ nhớ và tài nguyên của process đó. Thread có thể giao tiếp trực tiếp với nhau thông qua các biến chung.

**Sự khác biệt**:

- **Tài nguyên**: Process có bộ nhớ riêng biệt, còn các thread chia sẻ bộ nhớ của process.
- **Hiệu suất**: Thread nhẹ hơn process vì không cần không gian bộ nhớ riêng biệt.
- **Giao tiếp**: Process giao tiếp thông qua cơ chế IPC, còn các thread trong cùng process có thể giao tiếp trực tiếp.

---

### 3. **Thread có thể chia sẻ tài nguyên như thế nào?**

Các thread trong cùng một process chia sẻ tài nguyên hệ thống như bộ nhớ, file và socket. Các tài nguyên này được quản lý bởi hệ điều hành và thread có thể truy cập chúng mà không cần sao chép. Cách thức chia sẻ tài nguyên cụ thể:

- **Bộ nhớ**: Các thread có thể truy cập trực tiếp vào bộ nhớ chung của process, bao gồm các biến toàn cục hoặc heap. Điều này giúp tiết kiệm bộ nhớ nhưng cũng tạo ra vấn đề cần đồng bộ hóa.
- **File**: Các thread trong một process có thể truy cập và thao tác với các file được mở chung, miễn là chúng không xung đột trong quá trình sử dụng tài nguyên này.

- **Socket**: Các thread có thể chia sẻ kết nối socket nếu chúng cùng làm việc với một dịch vụ mạng.

**Lưu ý quan trọng**: Mặc dù các thread chia sẻ tài nguyên chung, việc truy cập đồng thời vào những tài nguyên này cần được đồng bộ hóa để tránh các vấn đề như race condition.

---

### 4. **Thread-safe là gì?**

**Thread-safe** là khái niệm mô tả một đoạn mã (hoặc hàm) có thể được thực thi đồng thời bởi nhiều thread mà không gây ra lỗi hoặc kết quả không xác định. Một hàm hoặc đoạn mã thread-safe đảm bảo rằng các thread sẽ không làm hỏng dữ liệu hoặc gây xung đột khi truy cập tài nguyên chung.

**Đảm bảo một hàm là thread-safe**:

- **Sử dụng đồng bộ hóa**: Dùng các cơ chế đồng bộ hóa như mutex, semaphore, hoặc lock để đảm bảo rằng chỉ một thread tại một thời điểm có thể truy cập vào tài nguyên chung.
- **Tránh thay đổi trạng thái toàn cục**: Nếu có thể, hạn chế sử dụng các biến toàn cục và đảm bảo rằng mỗi thread chỉ sử dụng các dữ liệu riêng biệt.

- **Atomic operations**: Sử dụng các phép toán nguyên tử (atomic) để tránh các tình huống tranh chấp tài nguyên, nơi một thread có thể thay đổi giá trị trong khi thread khác đang sử dụng nó.

- **Không sử dụng tài nguyên chung mà không đồng bộ**: Nếu nhiều thread cần truy cập vào tài nguyên chung (ví dụ: danh sách, bộ đếm), cần phải có cơ chế khóa (lock) để ngăn các thread khác thay đổi tài nguyên khi một thread đang sử dụng.

---
