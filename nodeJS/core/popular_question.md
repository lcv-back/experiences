Dưới đây là các câu hỏi cốt lõi về **Node.js**, cùng với các giải thích cơ bản:

### 1. **Node.js là gì?**

- **Câu trả lời**: Node.js là một môi trường thực thi JavaScript phía máy chủ (server-side) dựa trên engine V8 của Google Chrome. Nó cho phép bạn viết mã JavaScript để xử lý các tác vụ server như tạo API, tương tác với cơ sở dữ liệu, và quản lý các kết nối mạng. Node.js nổi bật với khả năng xử lý bất đồng bộ và song song, phù hợp với các ứng dụng thời gian thực.

### 2. **Node.js hoạt động như thế nào?**

- **Câu trả lời**: Node.js sử dụng mô hình **non-blocking, event-driven**. Điều này có nghĩa là thay vì chờ một tác vụ hoàn thành, Node.js sẽ tiếp tục thực hiện các tác vụ khác và chỉ quay lại khi có dữ liệu từ tác vụ trước đó. Điều này giúp tối ưu hiệu suất và tài nguyên, đặc biệt trong các ứng dụng yêu cầu xử lý nhiều kết nối đồng thời.

### 3. **Callback là gì trong Node.js?**

- **Câu trả lời**: Callback là một hàm được truyền vào như một đối số và sẽ được gọi khi tác vụ hoàn thành. Callback giúp xử lý bất đồng bộ trong Node.js. Ví dụ, sau khi đọc tệp, một callback có thể xử lý dữ liệu tệp khi tệp đã được tải lên.

### 4. **Event Loop trong Node.js là gì?**

- **Câu trả lời**: Event loop là cơ chế giúp Node.js xử lý các tác vụ bất đồng bộ. Khi một tác vụ bất đồng bộ được gọi (ví dụ: đọc tệp, truy vấn cơ sở dữ liệu), Node.js sẽ đưa tác vụ đó vào một hàng đợi, và khi tác vụ này hoàn thành, nó sẽ kích hoạt một callback để xử lý tiếp. Event loop giúp Node.js duy trì hiệu suất cao khi làm việc với số lượng lớn các kết nối đồng thời.

### 5. **Sự khác biệt giữa `require()` và `import` trong Node.js là gì?**

- **Câu trả lời**: `require()` là cú pháp sử dụng trong Node.js để import các module, trong khi `import` là cú pháp ES6 được hỗ trợ trong các phiên bản Node.js mới hơn (sử dụng `type: module` trong `package.json`). `require()` có thể được sử dụng trong mọi phiên bản của Node.js, trong khi `import` chỉ khả dụng trong các phiên bản hỗ trợ ES6 module.

### 6. **Sự khác biệt giữa `process.nextTick()` và `setImmediate()` là gì?**

- **Câu trả lời**:
  - **`process.nextTick()`**: Được sử dụng để thực thi một callback ngay sau khi hiện tại gọi hàm hoàn tất, trước khi event loop tiếp tục.
  - **`setImmediate()`**: Được sử dụng để thực thi một callback ngay sau khi tất cả các I/O events trong vòng lặp hiện tại đã hoàn tất.

### 7. **EventEmitter là gì trong Node.js?**

- **Câu trả lời**: **EventEmitter** là một lớp trong Node.js, cho phép phát và lắng nghe các sự kiện. Nó rất hữu ích trong việc phát tín hiệu (emit events) và xử lý các sự kiện trong ứng dụng, ví dụ như khi người dùng gửi yêu cầu HTTP.

### 8. **Stream trong Node.js là gì?**

- **Câu trả lời**: **Stream** là một đối tượng trong Node.js cho phép đọc hoặc ghi dữ liệu theo từng phần mà không cần phải tải toàn bộ vào bộ nhớ. Có 4 loại stream chính:
  - Readable stream (đọc dữ liệu)
  - Writable stream (ghi dữ liệu)
  - Duplex stream (cả đọc và ghi)
  - Transform stream (có thể thay đổi dữ liệu trong khi đọc/ghi)

### 9. **`npm` là gì và cách sử dụng nó như thế nào?**

- **Câu trả lời**: **npm** (Node Package Manager) là công cụ quản lý gói cho Node.js, cho phép bạn tải xuống và quản lý các thư viện, framework, và module mà bạn sử dụng trong dự án của mình. Bạn có thể sử dụng `npm install <package_name>` để cài đặt gói, và `npm start` để chạy ứng dụng.

### 10. **Callback Hell là gì và cách giải quyết vấn đề này trong Node.js?**

- **Câu trả lời**: **Callback hell** là hiện tượng khi các callback được lồng vào nhau, gây khó khăn trong việc đọc và bảo trì mã nguồn. Để tránh callback hell, có thể sử dụng **Promises** hoặc **Async/Await** để làm cho mã nguồn trở nên sạch hơn và dễ quản lý hơn.

### 11. **Sự khác biệt giữa `fs.readFile()` và `fs.createReadStream()` là gì?**

- **Câu trả lời**:
  - **`fs.readFile()`** đọc toàn bộ nội dung của tệp vào bộ nhớ trước khi trả về.
  - **`fs.createReadStream()`** tạo một luồng đọc, cho phép bạn đọc tệp theo từng phần, giúp tiết kiệm bộ nhớ khi làm việc với các tệp lớn.

### 12. **Các loại module trong Node.js là gì?**

- **Câu trả lời**: Node.js hỗ trợ 3 loại module:
  1.  **Core modules**: Các module có sẵn trong Node.js (ví dụ: `http`, `fs`, `path`).
  2.  **Third-party modules**: Các module được tải xuống từ npm (ví dụ: `express`, `lodash`).
  3.  **Custom modules**: Các module do người phát triển tạo ra trong dự án của mình.

### 13. **Cách xử lý lỗi trong Node.js là gì?**

- **Câu trả lời**: Trong Node.js, lỗi có thể được xử lý bằng cách sử dụng các callback để kiểm tra lỗi hoặc thông qua cơ chế **try-catch** nếu dùng `async/await`. Ngoài ra, việc xử lý lỗi bất đồng bộ thông qua event listeners hoặc trả về các thông báo lỗi qua các hàm callback là rất quan trọng.

### 14. **Những điểm cần chú ý khi làm việc với các ứng dụng Node.js sản xuất?**

- **Câu trả lời**:
  - Quản lý bộ nhớ và tài nguyên hệ thống
  - Sử dụng **cluster** để xử lý nhiều yêu cầu đồng thời
  - Xử lý các lỗi nghiêm túc và log lỗi
  - Giảm thiểu độ trễ và tối ưu hóa hiệu suất

Các câu hỏi này giúp kiểm tra kiến thức cốt lõi về Node.js và có thể giúp bạn chuẩn bị cho các cuộc phỏng vấn liên quan đến Node.js.
