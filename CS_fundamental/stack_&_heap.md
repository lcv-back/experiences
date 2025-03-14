# Stack và Heap

Trong lập trình máy tính, **stack** và **heap** là hai vùng bộ nhớ chính được sử dụng để quản lý dữ liệu, nhưng chúng khác nhau về cách hoạt động, mục đích sử dụng và cách quản lý. Dưới đây là sự phân biệt chi tiết:

### 1. **Định nghĩa**

- **Stack (Ngăn xếp)**: Là một vùng bộ nhớ hoạt động theo cơ chế LIFO (Last In, First Out - Vào sau, ra trước). Nó được sử dụng để lưu trữ các biến cục bộ và quản lý luồng thực thi của chương trình (như các hàm gọi đệ quy).
- **Heap (Đống)**: Là một vùng bộ nhớ động, cho phép cấp phát và giải phóng bộ nhớ linh hoạt trong suốt thời gian chạy chương trình. Nó được dùng để lưu trữ các đối tượng hoặc dữ liệu có kích thước không cố định hoặc cần tồn tại lâu hơn phạm vi của một hàm.

---

### 2. **Cách hoạt động**

- **Stack**:
  - Bộ nhớ được cấp phát và giải phóng tự động bởi trình biên dịch.
  - Khi một hàm được gọi, các biến cục bộ và thông tin ngữ cảnh (như địa chỉ trả về) được đẩy vào stack. Khi hàm kết thúc, phần bộ nhớ đó tự động được giải phóng.
  - Kích thước stack thường bị giới hạn (ví dụ: 1MB hoặc 8MB tùy hệ điều hành).
- **Heap**:
  - Bộ nhớ phải được cấp phát và giải phóng thủ công (trong các ngôn ngữ như C/C++) hoặc thông qua garbage collector (trong các ngôn ngữ như Java, Python).
  - Người lập trình quyết định khi nào cấp phát (dùng `malloc`, `new`) và giải phóng (dùng `free`, `delete`) hoặc để hệ thống tự quản lý (garbage collection).
  - Kích thước heap lớn hơn nhiều so với stack và chỉ bị giới hạn bởi bộ nhớ vật lý của máy.

---

### 3. **Ưu điểm và nhược điểm**

- **Stack**:
  - **Ưu điểm**: Nhanh, dễ quản lý, không cần lo về việc giải phóng bộ nhớ thủ công.
  - **Nhược điểm**: Giới hạn kích thước, không phù hợp với dữ liệu lớn hoặc dữ liệu cần tồn tại lâu dài.
- **Heap**:
  - **Ưu điểm**: Linh hoạt, có thể chứa dữ liệu lớn và tồn tại lâu dài.
  - **Nhược điểm**: Chậm hơn stack, dễ dẫn đến rò rỉ bộ nhớ (memory leak) nếu không quản lý tốt.

---

### 4. **Ví dụ minh họa**

Giả sử bạn viết một chương trình trong C:

```c
#include <stdio.h>
#include <stdlib.h>

void ham_vi_du() {
    int a = 10;          // Lưu trên stack
    int *b = malloc(sizeof(int)); // Lưu trên heap
    *b = 20;
    printf("a = %d, b = %d\n", a, *b);
    free(b);             // Giải phóng heap
}

int main() {
    ham_vi_du();
    return 0;
}
```

- Biến `a` được lưu trên stack, tự động bị xóa khi hàm `ham_vi_du()` kết thúc.
- Biến `b` được cấp phát trên heap, cần gọi `free()` để giải phóng, nếu không sẽ gây rò rỉ bộ nhớ.

---

### 5. **So sánh tóm tắt**

| Tiêu chí       | Stack                     | Heap                            |
| -------------- | ------------------------- | ------------------------------- |
| **Cơ chế**     | LIFO                      | Tự do, động                     |
| **Quản lý**    | Tự động (trình biên dịch) | Thủ công hoặc garbage collector |
| **Tốc độ**     | Nhanh                     | Chậm hơn                        |
| **Kích thước** | Giới hạn nhỏ              | Lớn, tùy bộ nhớ hệ thống        |
| **Ứng dụng**   | Biến cục bộ, gọi hàm      | Đối tượng động, dữ liệu lớn     |
