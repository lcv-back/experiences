# Data type

Trong Go (Golang), có nhiều kiểu dữ liệu khác nhau được chia thành các nhóm chính như kiểu số, kiểu chuỗi, kiểu bool, kiểu cấu trúc (struct), v.v. Dưới đây là danh sách các kiểu dữ liệu trong Go:

### 1. **Kiểu dữ liệu cơ bản:**

- **int**: Kiểu số nguyên. Kích thước của `int` tùy thuộc vào kiến trúc hệ thống (32-bit hoặc 64-bit).
  - Ví dụ: `var a int = 42`
- **int8, int16, int32, int64**: Kiểu số nguyên có độ dài cố định (8, 16, 32, 64 bit).
  - Ví dụ: `var b int8 = -128`
- **uint**: Kiểu số nguyên không dấu, có kích thước tùy thuộc vào hệ thống.
  - Ví dụ: `var c uint = 42`
- **uint8, uint16, uint32, uint64**: Kiểu số nguyên không dấu có độ dài cố định (8, 16, 32, 64 bit).
  - Ví dụ: `var d uint8 = 255`
- **float32, float64**: Kiểu số thực (số thập phân) với độ chính xác khác nhau.
  - Ví dụ: `var e float64 = 3.14`
- **complex64, complex128**: Kiểu số phức, với phần thực và phần ảo có thể là `float32` hoặc `float64`.
  - Ví dụ: `var f complex128 = 1 + 2i`

### 2. **Kiểu chuỗi:**

- **string**: Kiểu chuỗi, dùng để lưu trữ chuỗi văn bản.
  - Ví dụ: `var g string = "Hello, Go!"`

### 3. **Kiểu bool:**

- **bool**: Kiểu giá trị Boolean, chỉ có hai giá trị `true` hoặc `false`.
  - Ví dụ: `var h bool = true`

### 4. **Kiểu mảng (Array):**

- **array**: Mảng có độ dài cố định. Mảng trong Go có kiểu dữ liệu và số lượng phần tử cố định.
  - Ví dụ: `var i [5]int = [5]int{1, 2, 3, 4, 5}`

### 5. **Kiểu slice:**

- **slice**: Là một mảng động, có thể thay đổi kích thước và không có độ dài cố định.
  - Ví dụ: `var j []int = []int{1, 2, 3, 4}`

### 6. **Kiểu bản đồ (Map):**

- **map**: Kiểu dữ liệu ánh xạ (key-value pair). Mỗi phần tử trong `map` bao gồm một khóa (key) và giá trị (value).
  - Ví dụ: `var k map[string]int = map[string]int{"a": 1, "b": 2}`

### 7. **Kiểu con trỏ (Pointer):**

- **pointer**: Go hỗ trợ con trỏ, cho phép tham chiếu đến các giá trị thay vì sao chép chúng.
  - Ví dụ: `var l *int = &a` (con trỏ đến biến `a`)

### 8. **Kiểu cấu trúc (Struct):**

- **struct**: Cấu trúc là kiểu dữ liệu có thể chứa nhiều kiểu dữ liệu khác nhau. Dùng để định nghĩa đối tượng với các thuộc tính.
  - Ví dụ:
    ```go
    type Person struct {
        Name string
        Age  int
    }
    var m Person = Person{"John", 30}
    ```

### 9. **Kiểu giao diện (Interface):**

- **interface**: Kiểu interface là một tập hợp các phương thức mà một đối tượng phải thực hiện. Giao diện giúp Go hỗ trợ lập trình hướng đối tượng.
  - Ví dụ:
    ```go
    type Speaker interface {
        Speak() string
    }
    ```

### 10. **Kiểu hàm (Function):**

- **function**: Hàm là kiểu dữ liệu trong Go, có thể được gán vào biến và truyền như tham số.
  - Ví dụ:
    ```go
    var fn func(int, int) int = func(a, b int) int {
        return a + b
    }
    ```

### 11. **Kiểu kênh (Channel):**

- **channel**: Kênh được sử dụng để giao tiếp giữa các goroutines trong Go.
  - Ví dụ: `var n chan int = make(chan int)`

### 12. **Kiểu loại (Type alias):**

- **type alias**: Go cho phép định nghĩa tên mới cho các kiểu dữ liệu có sẵn.
  - Ví dụ: `type Age int` (định nghĩa kiểu mới `Age` là alias của `int`)

### Tổng kết:

- Các kiểu dữ liệu trong Go rất linh hoạt và đa dạng, bao gồm cả kiểu cơ bản (như int, float), kiểu phức tạp (struct, map), kiểu động (slice), và các kiểu đặc biệt như con trỏ, hàm, giao diện và kênh.
