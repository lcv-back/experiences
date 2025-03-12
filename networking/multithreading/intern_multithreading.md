### III. **Câu hỏi về vấn đề đa luồng**

### 1. **Race Condition là gì? Làm sao để tránh nó?**

**Race Condition** xảy ra khi hai hoặc nhiều thread cùng truy cập và thay đổi dữ liệu chung một cách đồng thời, dẫn đến kết quả không xác định. Điều này thường xảy ra khi các thread không được đồng bộ hóa đúng cách, và kết quả của chương trình phụ thuộc vào thứ tự thực thi của các thread.

**Cách tránh race condition**:

- **Sử dụng Mutex**: Một mutex (mutual exclusion) cho phép chỉ một thread tại một thời điểm truy cập tài nguyên chung. Các thread khác phải đợi cho đến khi mutex được giải phóng.

  Ví dụ trong C++:

  ```cpp
  std::mutex mtx;
  void increment() {
      std::lock_guard<std::mutex> guard(mtx); // Mutex tự động khóa và giải phóng
      counter++; // Thực hiện thay đổi dữ liệu chung
  }
  ```

- **Sử dụng Semaphore**: Semaphore có thể giới hạn số lượng thread có thể truy cập tài nguyên đồng thời. Với semaphore, ta có thể đảm bảo chỉ có một số lượng thread nhất định được phép thao tác với tài nguyên.
- **Sử dụng Atomic operations**: Các phép toán nguyên tử đảm bảo rằng quá trình thay đổi dữ liệu chung là không thể bị gián đoạn và thread khác không thể can thiệp vào giữa quá trình này.

---

### 2. **Deadlock là gì? Làm sao để tránh deadlock trong chương trình đa luồng?**

**Deadlock** là tình huống xảy ra khi hai hoặc nhiều thread đang chờ nhau giải phóng tài nguyên mà chúng cần để tiếp tục thực thi, dẫn đến một vòng lặp vô hạn mà không có thread nào có thể tiến hành được nữa.

**Cách tránh deadlock**:

1. **Thứ tự khóa hợp lý**: Đảm bảo rằng tất cả các thread yêu cầu các tài nguyên trong một thứ tự cố định. Nếu mọi thread đều cố gắng khóa các tài nguyên theo cùng một thứ tự, sẽ không có tình huống deadlock xảy ra.

   Ví dụ: Nếu thread A khóa tài nguyên X và cần tài nguyên Y, thread B phải khóa tài nguyên Y trước khi yêu cầu tài nguyên X. Nếu không, sẽ xảy ra deadlock.

2. **Sử dụng Timeout**: Nếu một thread không thể lấy được tài nguyên trong một khoảng thời gian nhất định, nó sẽ từ bỏ và thử lại sau hoặc giải phóng tài nguyên đã chiếm giữ trước đó.

3. **Tránh chặn tài nguyên**: Đảm bảo rằng một thread không giữ tài nguyên và đồng thời yêu cầu tài nguyên khác mà một thread khác đang giữ.

4. **Sử dụng Try-Lock**: Một phương pháp khác là sử dụng "try-lock" thay vì "lock". Nếu không thể chiếm tài nguyên ngay lập tức, thread có thể thử lại sau thay vì chờ vô thời hạn.

---

### 3. **Livelock là gì và khác gì so với deadlock?**

**Livelock** xảy ra khi hai hoặc nhiều thread liên tục thay đổi trạng thái của chúng để giải quyết một vấn đề, nhưng không tiến triển hơn. Các thread không bị chặn lại hoàn toàn như trong deadlock, nhưng chúng vẫn không thể hoàn thành công việc do liên tục phải thay đổi trạng thái hoặc hành động mà không tiến xa hơn.

**Khác biệt với deadlock**:

- **Deadlock**: Các thread bị chặn lại vĩnh viễn vì chúng đang đợi nhau giải phóng tài nguyên mà chúng cần.
- **Livelock**: Các thread không bị chặn hoàn toàn, nhưng chúng liên tục thay đổi trạng thái mà không tiến triển vào kết quả cuối cùng.

**Ví dụ về livelock**:
Một ví dụ về livelock là khi hai thread cố gắng tránh gây tắc nghẽn trong khi cùng truy cập tài nguyên, nhưng cả hai thread cứ di chuyển qua lại mà không đạt được mục tiêu cuối cùng.

```cpp
#include <iostream>
#include <thread>
#include <atomic>

std::atomic<bool> flagA(false), flagB(false);

void threadA() {
    while (true) {
        if (!flagB) {
            flagA = true;
            if (!flagB) {
                std::cout << "Thread A performing work." << std::endl;
                return;
            }
            flagA = false; // Giải phóng và thử lại
        }
    }
}

void threadB() {
    while (true) {
        if (!flagA) {
            flagB = true;
            if (!flagA) {
                std::cout << "Thread B performing work." << std::endl;
                return;
            }
            flagB = false; // Giải phóng và thử lại
        }
    }
}

int main() {
    std::thread t1(threadA);
    std::thread t2(threadB);

    t1.join();
    t2.join();
}
```

Trong ví dụ này, cả hai thread liên tục thử kiểm tra điều kiện nhưng không bao giờ hoàn thành công việc của mình, tạo ra livelock.

---

### 4. **Priority Inversion là gì? Làm sao để giải quyết vấn đề này trong hệ thống đa luồng?**

**Priority Inversion** là vấn đề xảy ra khi một thread có mức ưu tiên thấp hơn lại giữ tài nguyên mà một thread có mức ưu tiên cao hơn cần. Điều này khiến cho thread ưu tiên cao bị "đảo ngược" mức độ ưu tiên và không thể thực hiện công việc vì phải chờ thread ưu tiên thấp hoàn thành công việc.

**Ví dụ**: Một thread có mức ưu tiên cao (A) cần tài nguyên mà thread có mức ưu tiên thấp (C) đang giữ. Nếu giữa chúng có một thread mức ưu tiên trung bình (B), thread B sẽ có thể tiếp tục thực thi trong khi thread A bị chặn lại, dù thread A cần tài nguyên đó.

**Cách giải quyết**:

1. **Priority Inheritance**: Một kỹ thuật giải quyết priority inversion bằng cách cho phép thread có mức ưu tiên thấp (giữ tài nguyên) mượn mức ưu tiên của thread có mức ưu tiên cao cho đến khi nó giải phóng tài nguyên. Điều này giúp thread cao ưu tiên có thể hoàn thành công việc mà không bị chặn bởi thread có mức ưu tiên thấp.

2. **Priority Ceiling Protocol**: Trong giao thức này, tất cả các thread có thể yêu cầu tài nguyên sẽ phải thực thi với mức ưu tiên cao nhất mà chúng có thể đạt được khi chiếm tài nguyên đó. Điều này giúp tránh tình huống thread có mức ưu tiên thấp giữ tài nguyên khi thread có mức ưu tiên cao cần.

**Ví dụ về Priority Inheritance**:
Giả sử thread A có mức ưu tiên cao nhất và thread C có mức ưu tiên thấp. Nếu thread A cần tài nguyên mà thread C đang giữ, thread C sẽ tạm thời "mượn" mức ưu tiên của thread A để thực thi nhanh chóng và giải phóng tài nguyên.

---

Những vấn đề trên thường xuyên xuất hiện trong các ứng dụng đa luồng và có thể gây ra các lỗi nghiêm trọng nếu không được xử lý đúng cách. Việc hiểu rõ và áp dụng các giải pháp phòng tránh giúp xây dựng phần mềm ổn định và hiệu quả.
