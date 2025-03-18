Dưới đây là các câu hỏi phổ biến về **React** dành cho **fresher** (mới bắt đầu):

### 1. **React là gì?**

- **Câu trả lời**: React là một thư viện JavaScript mã nguồn mở do Facebook phát triển, được sử dụng để xây dựng giao diện người dùng, đặc biệt là các ứng dụng đơn trang (SPA). React cho phép phát triển giao diện hiệu quả, dễ dàng quản lý và tối ưu hóa các thay đổi giao diện mà không cần phải tải lại trang web.

### 2. **JSX trong React là gì?**

- **Câu trả lời**: JSX (JavaScript XML) là một cú pháp mở rộng cho JavaScript, cho phép bạn viết các thẻ HTML trong mã JavaScript. JSX không phải là một phần của JavaScript chuẩn, nhưng React sử dụng nó để mô tả giao diện người dùng. Nó được biên dịch thành mã JavaScript thuần trước khi chạy.

### 3. **Sự khác biệt giữa `class` component và `function` component là gì?**

- **Câu trả lời**:
  - **Class components** là các component trong React được tạo ra bằng cách sử dụng lớp (class). Chúng có thể sử dụng **state** và **lifecycle methods** (ví dụ: `componentDidMount`, `shouldComponentUpdate`).
  - **Function components** là các component đơn giản hơn và không sử dụng class. Trước khi có **Hooks**, function components chỉ có thể nhận props và trả về JSX. Tuy nhiên, hiện tại, function components có thể sử dụng **state** và **lifecycle methods** thông qua **React Hooks** (như `useState`, `useEffect`).

### 4. **State trong React là gì và tại sao nó quan trọng?**

- **Câu trả lời**: **State** là một đối tượng trong React dùng để lưu trữ dữ liệu có thể thay đổi và ảnh hưởng đến việc render lại giao diện. Khi state thay đổi, React sẽ tự động re-render component để phản ánh các thay đổi đó trên giao diện.

### 5. **Props trong React là gì?**

- **Câu trả lời**: **Props** (viết tắt của "properties") là đối số được truyền vào một component từ component cha. Props giúp các component có thể tái sử dụng và chia sẻ dữ liệu giữa các component khác nhau trong ứng dụng.

### 6. **Cách xử lý sự kiện trong React là gì?**

- **Câu trả lời**: Trong React, bạn có thể xử lý sự kiện giống như trong HTML, nhưng có một số điểm khác biệt:
  - Các tên sự kiện trong React được viết theo cú pháp camelCase (ví dụ: `onClick` thay vì `onclick`).
  - Để gán một hàm xử lý sự kiện, bạn phải truyền hàm vào thuộc tính sự kiện (ví dụ: `<button onClick={this.handleClick}>Click me</button>`).
  - Không cần phải sử dụng `this` khi gọi hàm xử lý sự kiện trong function component.

### 7. **React Hooks là gì và tại sao chúng lại quan trọng?**

- **Câu trả lời**: **React Hooks** là các hàm đặc biệt cho phép bạn "hook vào" các tính năng của React như **state** và **lifecycle methods** mà không cần phải viết class. Các hook phổ biến bao gồm:
  - `useState`: Quản lý state trong function component.
  - `useEffect`: Xử lý các side effects (như API calls, cập nhật DOM).
  - `useContext`: Dễ dàng truy cập vào context.

### 8. **Component Lifecycle trong React là gì?**

- **Câu trả lời**: **Lifecycle** là các giai đoạn mà một component trải qua từ khi được tạo ra cho đến khi bị xóa khỏi DOM. Các phương thức lifecycle thường gặp trong class component bao gồm:
  - `componentDidMount()`: Được gọi sau khi component đã được render lần đầu tiên.
  - `componentDidUpdate()`: Được gọi mỗi khi component được cập nhật.
  - `componentWillUnmount()`: Được gọi trước khi component bị xóa khỏi DOM.
- Trong function component, bạn có thể thay thế lifecycle methods bằng `useEffect`.

### 9. **Virtual DOM là gì?**

- **Câu trả lời**: **Virtual DOM** là một bản sao của DOM thực tế được lưu trữ trong bộ nhớ. Khi state hoặc props thay đổi, React sẽ so sánh Virtual DOM mới với Virtual DOM cũ và chỉ cập nhật những phần của DOM thực tế mà đã thay đổi. Điều này giúp tối ưu hóa hiệu suất và giảm thiểu việc render lại toàn bộ giao diện.

### 10. **React Context là gì?**

- **Câu trả lời**: **React Context** là một cách để chia sẻ dữ liệu giữa các component mà không cần phải truyền props qua từng cấp. Đây là một giải pháp hữu ích khi bạn cần truyền dữ liệu toàn cục như theme, ngôn ngữ, hoặc thông tin người dùng giữa các component trong ứng dụng.

### 11. **React Router là gì và tại sao chúng ta cần nó?**

- **Câu trả lời**: **React Router** là thư viện cho phép quản lý điều hướng trong ứng dụng React. Nó cho phép bạn chuyển đổi giữa các trang hoặc component khác nhau mà không cần tải lại toàn bộ trang, tạo ra một trải nghiệm người dùng mượt mà hơn cho ứng dụng đơn trang (SPA).

### 12. **Cách tối ưu hóa hiệu suất trong React là gì?**

- **Câu trả lời**: Để tối ưu hóa hiệu suất trong React, bạn có thể thực hiện một số bước như:
  - **React.memo**: Hạn chế việc render lại không cần thiết của component.
  - **useMemo** và **useCallback**: Dùng để nhớ các giá trị hoặc hàm giúp tránh tính toán lại không cần thiết.
  - **Lazy loading**: Chỉ tải các phần của ứng dụng khi người dùng cần.

### 13. **Key trong React là gì và tại sao nó quan trọng?**

- **Câu trả lời**: **Key** là một thuộc tính đặc biệt được sử dụng trong React khi tạo ra danh sách các element (ví dụ: trong `map()`). Key giúp React nhận diện các phần tử đã thay đổi, thêm mới hoặc bị xóa, giúp tối ưu hóa hiệu suất khi render lại các phần tử của danh sách.

### 14. **Sự khác biệt giữa `useState()` và `useEffect()` là gì?**

- **Câu trả lời**:
  - **`useState()`** được sử dụng để quản lý trạng thái (state) trong function component.
  - **`useEffect()`** được sử dụng để xử lý các side effects như gọi API, cập nhật DOM, hoặc các tác vụ bất đồng bộ khác. `useEffect()` có thể thay thế các lifecycle methods trong class component.

### 15. **Component trong React là gì?**

- **Câu trả lời**: Component trong React là một đơn vị cơ bản để xây dựng giao diện người dùng. Mỗi component có thể có trạng thái riêng và nhận dữ liệu thông qua props. React hỗ trợ việc tái sử dụng component, giúp mã nguồn trở nên dễ bảo trì và mở rộng.

Những câu hỏi này giúp kiểm tra kiến thức cơ bản và hiểu biết của bạn về React khi bắt đầu học và phát triển ứng dụng với React.
