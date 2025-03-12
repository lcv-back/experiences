### II. **Câu hỏi về đồng bộ hóa (Synchronization)**

### 1. **Mutex là gì? Làm sao để sử dụng mutex trong C/C++ hoặc Golang?**

**Mutex** (Mutual Exclusion) là một cơ chế đồng bộ hóa được sử dụng để đảm bảo rằng chỉ một thread tại một thời điểm có thể truy cập vào một tài nguyên chung. Khi một thread chiếm được mutex, các thread khác phải đợi đến khi mutex được giải phóng trước khi chúng có thể truy cập tài nguyên đó.

**Cách sử dụng mutex**:

- **Trong C++**: Thường sử dụng thư viện `<mutex>` để sử dụng `std::mutex`.

  Ví dụ:

  ```cpp
  #include <iostream>
  #include <thread>
  #include <mutex>

  std::mutex mtx; // Tạo một mutex toàn cục

  void print_hello(int i) {
      mtx.lock(); // Khóa mutex
      std::cout << "Hello from thread " << i << std::endl;
      mtx.unlock(); // Giải phóng mutex
  }

  int main() {
      std::thread t1(print_hello, 1);
      std::thread t2(print_hello, 2);

      t1.join();
      t2.join();

      return 0;
  }
  ```

  - `mtx.lock()` sẽ khóa mutex, và chỉ một thread có thể vào trong vùng mã này.
  - `mtx.unlock()` sẽ giải phóng mutex để các thread khác có thể sử dụng.

- **Trong Golang**: Sử dụng `sync.Mutex` để làm việc với mutex.

  Ví dụ:

  ```go
  package main

  import (
      "fmt"
      "sync"
  )

  var mtx sync.Mutex // Tạo một mutex toàn cục

  func printHello(i int) {
      mtx.Lock() // Khóa mutex
      fmt.Println("Hello from thread", i)
      mtx.Unlock() // Giải phóng mutex
  }

  func main() {
      go printHello(1)
      go printHello(2)

      // Đợi cho các goroutine chạy xong
      var input string
      fmt.Scanln(&input)
  }
  ```

  - `mtx.Lock()` khóa mutex.
  - `mtx.Unlock()` giải phóng mutex.

---

### 2. **Semaphore là gì?**

**Semaphore** là một cơ chế đồng bộ hóa khác, được dùng để kiểm soát số lượng thread có thể truy cập tài nguyên đồng thời. Một semaphore có giá trị ban đầu, và mỗi lần thread truy cập tài nguyên, giá trị này giảm đi. Khi giá trị semaphore bằng 0, các thread khác sẽ phải đợi cho đến khi có tài nguyên.

**Khác biệt giữa semaphore và mutex**:

- **Mutex**: Chỉ cho phép một thread truy cập tài nguyên tại một thời điểm.
- **Semaphore**: Cho phép nhiều thread truy cập tài nguyên đồng thời (tùy thuộc vào giá trị của semaphore). Semaphore có thể được sử dụng để giới hạn số lượng thread truy cập tài nguyên.

Ví dụ về **semaphore** trong Golang:

```go
package main

import (
	"fmt"
	"sync"
)

var semaphore = make(chan struct{}, 3) // Cho phép tối đa 3 goroutine truy cập tài nguyên

func accessResource(i int) {
	semaphore <- struct{}{} // Đặt semaphore (tương đương với sem.Wait())
	fmt.Println("Thread", i, "accessing resource")
	<-semaphore // Giải phóng semaphore (tương đương với sem.Signal())
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			accessResource(i)
		}(i)
	}
	wg.Wait()
}
```

Ở ví dụ trên, `semaphore` giới hạn số lượng goroutine có thể truy cập tài nguyên đến 3.

---

### 3. **Lock là gì và làm sao để sử dụng lock trong lập trình đa luồng?**

**Lock** là một cơ chế đồng bộ hóa giúp các thread điều khiển việc truy cập vào tài nguyên chung. Khi một thread "lock" một tài nguyên, các thread khác không thể truy cập tài nguyên đó cho đến khi thread đầu tiên giải phóng lock.

**Cách sử dụng lock**:

- **Trong C++**: `std::lock_guard` hoặc `std::unique_lock` được sử dụng để lock tự động và giải phóng lock khi ra khỏi phạm vi.

  Ví dụ với `std::lock_guard`:

  ```cpp
  std::mutex mtx;

  void print_hello(int i) {
      std::lock_guard<std::mutex> guard(mtx); // Khóa tự động khi vào hàm và giải phóng khi ra khỏi
      std::cout << "Hello from thread " << i << std::endl;
  }
  ```

- **Trong Golang**: Sử dụng `sync.Mutex.Lock()` và `sync.Mutex.Unlock()` để lock và unlock thủ công.

  Ví dụ:

  ```go
  var mtx sync.Mutex

  func printHello(i int) {
      mtx.Lock() // Khóa
      fmt.Println("Hello from thread", i)
      mtx.Unlock() // Giải phóng
  }
  ```

**Tình huống cần sử dụng lock**:

- Khi nhiều thread truy cập hoặc thay đổi cùng một tài nguyên (như biến toàn cục, file, hay cơ sở dữ liệu) và cần đảm bảo rằng chỉ một thread có thể thao tác tại một thời điểm.
- Khi bạn muốn tránh **race condition** hoặc các tình huống không xác định khác do đồng thời truy cập tài nguyên.

---

### 4. **Điều gì sẽ xảy ra nếu hai thread cùng cố gắng truy cập tài nguyên đồng thời mà không sử dụng đồng bộ hóa?**

Khi hai thread cùng cố gắng truy cập tài nguyên đồng thời mà không có cơ chế đồng bộ hóa, sẽ xảy ra **race condition**. Đây là tình huống mà kết quả của chương trình phụ thuộc vào thứ tự mà các thread thực thi, và có thể dẫn đến các lỗi không thể dự đoán được.

**Ví dụ về race condition**:
Giả sử có một biến toàn cục `counter` được các thread tăng lên, nhưng không có cơ chế đồng bộ hóa:

```cpp
#include <iostream>
#include <thread>

int counter = 0; // Biến toàn cục chia sẻ

void increment() {
    for (int i = 0; i < 1000; ++i) {
        counter++; // Không đồng bộ hóa
    }
}

int main() {
    std::thread t1(increment);
    std::thread t2(increment);

    t1.join();
    t2.join();

    std::cout << "Counter value: " << counter << std::endl; // Kết quả không chính xác
    return 0;
}
```

**Vấn đề**: Nếu không sử dụng đồng bộ hóa, có thể một thread đọc giá trị của `counter` trong khi thread khác thay đổi giá trị đó, dẫn đến kết quả không chính xác (counter có thể không tăng đúng 2000).

**Giải pháp**: Sử dụng mutex hoặc lock để đảm bảo rằng chỉ một thread có thể truy cập và thay đổi giá trị của `counter` tại một thời điểm.

Race condition có thể gây ra các lỗi nghiêm trọng như dữ liệu bị hỏng, chương trình bị treo hoặc crash, và khó debug vì các lỗi này không xuất hiện theo cách nhất quán.
