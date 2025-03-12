### IV. **Câu hỏi về lập trình bất đồng bộ**

Dưới đây là các câu trả lời cho các câu hỏi về lập trình bất đồng bộ, chỉ áp dụng trong **Golang**:

### 1. **Lập trình bất đồng bộ (Asynchronous Programming) là gì?**

Trong **Golang**, lập trình bất đồng bộ chủ yếu được thực hiện thông qua **goroutines**. Một goroutine là một đơn vị thực thi nhẹ mà Golang cung cấp để thực thi các tác vụ đồng thời mà không tạo ra overhead như thread hệ điều hành. Khi tạo một goroutine, chương trình không cần phải chờ đợi hoàn thành của nó mà có thể tiếp tục thực thi các tác vụ khác. Đây chính là cách thực hiện lập trình bất đồng bộ trong Golang.

- **Khác biệt với lập trình đồng bộ**: Trong lập trình đồng bộ, mỗi tác vụ sẽ chờ hoàn thành trước khi thực hiện tác vụ tiếp theo. Trong khi đó, lập trình bất đồng bộ cho phép nhiều tác vụ chạy song song mà không cần chờ đợi lẫn nhau.

**Ví dụ về lập trình bất đồng bộ trong Golang**:

```go
package main

import (
	"fmt"
	"time"
)

func doTask() {
	fmt.Println("Task started")
	time.Sleep(2 * time.Second) // Giả lập một tác vụ tốn thời gian
	fmt.Println("Task finished")
}

func main() {
	go doTask() // Tạo một goroutine bất đồng bộ
	fmt.Println("Main function continues executing...")
	time.Sleep(3 * time.Second) // Đợi goroutine hoàn thành
}
```

Ở ví dụ trên, `doTask` được chạy trong một goroutine và không làm chương trình chính bị chặn lại.

---

### 2. **Promise và Future là gì?**

Trong Golang, **Promise** và **Future** không có sẵn như trong một số ngôn ngữ khác. Tuy nhiên, bạn có thể sử dụng **channel** để mô phỏng chức năng này.

- **Promise**: Bạn có thể tạo một **channel** để gửi dữ liệu từ một goroutine trở lại hàm chính hoặc các goroutine khác, tương tự như cách Promise hoạt động.
- **Future**: Tương tự như Promise, Future trong Golang có thể được mô phỏng bằng cách sử dụng channel để nhận kết quả sau khi một tác vụ bất đồng bộ hoàn thành.

**Ví dụ với channel mô phỏng Future trong Golang**:

```go
package main

import (
	"fmt"
	"time"
)

func computeResult(ch chan int) {
	time.Sleep(2 * time.Second)
	ch <- 42 // Gửi kết quả vào channel
}

func main() {
	resultCh := make(chan int) // Channel để nhận kết quả

	go computeResult(resultCh) // Goroutine thực hiện tác vụ bất đồng bộ

	result := <-resultCh // Nhận kết quả từ channel (tương tự Future.get())
	fmt.Println("Result:", result)
}
```

---

### 3. **Trong Golang, goroutine là gì và cách hoạt động của nó?**

**Goroutine** là một đơn vị thực thi nhẹ trong Golang, tương tự như một thread nhưng ít tốn tài nguyên hơn. Golang sử dụng **Go runtime** để quản lý và thực thi các goroutine. Điều đặc biệt là bạn có thể tạo hàng ngàn hoặc hàng triệu goroutine mà không tốn quá nhiều tài nguyên.

**Cách tạo và hoạt động của goroutine**:

- Để tạo một goroutine, chỉ cần thêm từ khóa `go` trước hàm hoặc biểu thức mà bạn muốn thực thi bất đồng bộ.
- Các goroutine có thể giao tiếp với nhau thông qua **channel** để đồng bộ hóa hoặc chia sẻ dữ liệu.

**Ví dụ về goroutine trong Golang**:

```go
package main

import "fmt"

func task(id int) {
	fmt.Printf("Task %d is running\n", id)
}

func main() {
	for i := 1; i <= 5; i++ {
		go task(i) // Tạo một goroutine cho mỗi task
	}
	// Để chương trình không kết thúc quá sớm
	var input string
	fmt.Scanln(&input)
}
```

Ở ví dụ trên, mỗi lần lặp tạo ra một **goroutine** mới để thực thi hàm `task`.

**Đồng bộ hóa giữa các goroutine**: Bạn có thể sử dụng **channel** hoặc **WaitGroup** để đồng bộ hóa và đợi các goroutine hoàn thành.

**Ví dụ với WaitGroup**:

```go
package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Đảm bảo WaitGroup được giảm khi goroutine hoàn thành
	fmt.Printf("Task %d is running\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Thêm 1 vào WaitGroup
		go task(i, &wg)
	}

	wg.Wait() // Chờ tất cả các goroutine hoàn thành
	fmt.Println("All tasks completed")
}
```

---

### 4. **Cách sử dụng callback trong lập trình bất đồng bộ là gì?**

Trong Golang, **callback** có thể được sử dụng thông qua các hàm được truyền vào các goroutine để xử lý kết quả khi một tác vụ bất đồng bộ hoàn thành. Hàm callback này có thể thực thi khi một tác vụ bất đồng bộ hoàn thành hoặc khi có sự kiện xảy ra.

**Ví dụ về callback trong Golang**:

```go
package main

import "fmt"

func performTask(callback func(string)) {
	// Giả lập một tác vụ bất đồng bộ
	callback("Task completed successfully!")
}

func main() {
	// Truyền một hàm callback vào performTask
	performTask(func(result string) {
		fmt.Println(result) // Xử lý kết quả khi tác vụ hoàn thành
	})
}
```

Trong ví dụ này, hàm `performTask` nhận một callback và thực thi callback đó khi tác vụ hoàn thành.

Ngoài ra, callback cũng có thể được sử dụng trong các trường hợp như xử lý kết quả của nhiều goroutine hoặc các tác vụ I/O không đồng bộ.

---

Tóm lại, trong Golang, lập trình bất đồng bộ chủ yếu dựa vào **goroutine** và **channel** để thực hiện các tác vụ song song, không đồng bộ. Các khái niệm như **Promise**, **Future**, và **callback** có thể được mô phỏng hiệu quả bằng cách sử dụng các công cụ tích hợp của Golang.
