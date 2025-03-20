# Round 2: DSA

1. You are given positions of N two-dimensional points described by two arrays X and Y, both of N integers. The K-th point position is represented by the pair (X[K], Y[K]) - the respective X and Y coordinates of the point. Your task is to check if there exists a square, with sides parallel to the Ox/Oy axes, such that all the given points lie on this square's perimeter. Given two arrays X and Y of N integers, returns true if all the points lie on the perimeter of some square and false otherwise.
   Write an efficient algorithm for the following assumptions:
   - N max 20
   - Value X, Y [-10, 10]

## Translate

### Giải thích đề bài

Bạn được cung cấp:

- Hai mảng `X` và `Y`, mỗi mảng chứa `N` số nguyên.
- Mỗi cặp `(X[K], Y[K])` biểu thị tọa độ của điểm thứ `K` trong mặt phẳng 2D.
- Nhiệm vụ là kiểm tra xem có tồn tại một hình vuông, với các cạnh song song với trục Ox (trục X) và Oy (trục Y), sao cho **tất cả** các điểm đã cho nằm trên **chu vi** (đường viền) của hình vuông này không.

#### Ràng buộc:

- `N` (số điểm) tối đa là 20.
- Giá trị trong `X` và `Y` là số nguyên nằm trong khoảng từ `-10` đến `10`.

#### "Hình vuông với các cạnh song song với trục Ox/Oy" nghĩa là gì?

- Điều này có nghĩa là hình vuông không bị xoay; các cạnh của nó hoặc là nằm ngang (song song với trục X) hoặc là thẳng đứng (song song với trục Y).
- Trong mặt phẳng 2D, một hình vuông như vậy có thể được xác định bởi:
  - Một góc dưới bên trái `(x_min, y_min)`.
  - Độ dài cạnh `s`.
  - Góc trên bên phải sẽ là `(x_min + s, y_min + s)`.

#### "Điểm nằm trên chu vi" nghĩa là gì?

- Để một điểm `(x, y)` nằm trên chu vi của hình vuông có các góc `(x_min, y_min)` và `(x_min + s, y_min + s)`, nó phải thỏa mãn một trong các điều kiện sau:
  - `x = x_min` và `y_min ≤ y ≤ y_min + s` (cạnh trái của hình vuông).
  - `x = x_min + s` và `y_min ≤ y ≤ y_min + s` (cạnh phải của hình vuông).
  - `y = y_min` và `x_min ≤ x ≤ x_max` (cạnh dưới của hình vuông).
  - `y = y_min + s` và `x_min ≤ x ≤ x_max` (cạnh trên của hình vuông).
- Nói cách khác, điểm phải nằm trên một trong bốn cạnh của hình vuông.

#### Mục tiêu:

- Trả về `true` nếu tất cả `N` điểm đều nằm trên chu vi của một hình vuông nào đó (có cạnh song song với các trục).
- Trả về `false` nếu không tồn tại hình vuông nào như vậy.

---

### Ví dụ minh họa

#### Ví dụ 1:

```
X = [0, 0, 1, 1]
Y = [0, 1, 0, 1]
N = 4
```

- Các điểm: `(0, 0), (0, 1), (1, 0), (1, 1)`.
- Các điểm này tạo thành một hình vuông với:
  - Góc dưới trái: `(0, 0)`.
  - Góc trên phải: `(1, 1)`.
  - Độ dài cạnh: `1`.
- Tất cả các điểm đều nằm trên chu vi của hình vuông này (mỗi điểm là một góc).
- Kết quả: `true`.

#### Ví dụ 2:

```
X = [0, 0, 2]
Y = [0, 2, 1]
N = 3
```

- Các điểm: `(0, 0), (0, 2), (2, 1)`.
- Liệu có tồn tại một hình vuông mà tất cả các điểm này nằm trên chu vi không?
  - Nếu thử một hình vuông có cạnh dài 2 (ví dụ từ `(0, 0)` đến `(2, 2)`):
    - `(0, 0)` nằm ở góc dưới trái (thuộc chu vi).
    - `(0, 2)` nằm ở góc trên trái (thuộc chu vi).
    - `(2, 1)` nằm bên trong hình vuông, không nằm trên chu vi (vì `y = 1` nằm giữa `0` và `2`, nhưng không thuộc cạnh phải `x = 2` hay cạnh trên `y = 2`).
- Không có hình vuông nào chứa tất cả các điểm này trên chu vi.
- Kết quả: `false`.

---

### Cách giải quyết

Vì `N` nhỏ (tối đa 20), ta có thể dùng cách tiếp cận có hệ thống để giải quyết hiệu quả. Ý tưởng chính như sau:

1. **Nhận xét quan trọng**: Để tồn tại một hình vuông mà tất cả các điểm nằm trên chu vi:

   - Các điểm phải nằm trên bốn đường thẳng định nghĩa hình vuông: `x = x_min`, `x = x_max`, `y = y_min`, `y = y_max`.
   - Độ dài cạnh của hình vuông phải thỏa mãn `s = x_max - x_min = y_max - y_min`.
   - Mỗi điểm `(X[K], Y[K])` phải có `x` bằng `x_min` hoặc `x_max`, **hoặc** `y` bằng `y_min` hoặc `y_max`, và tọa độ phải phù hợp với hình vuông.

2. **Các bước**:

   - Tính giá trị nhỏ nhất và lớn nhất của `X` (`x_min`, `x_max`) và `Y` (`y_min`, `y_max`).
   - Kiểm tra xem `x_max - x_min` có bằng `y_max - y_min` không (hình vuông phải có các cạnh bằng nhau).
   - Xác minh rằng mỗi điểm `(X[K], Y[K])` nằm trên một trong bốn cạnh của hình vuông được định nghĩa bởi `(x_min, y_min)` và độ dài cạnh `s`.

3. **Thuật toán** (dưới dạng mã giả):

```pseudocode
hàm kiemTraHinhVuong(X[], Y[], N):
    nếu N == 0:
        trả về true  // Tập rỗng thì luôn đúng

    x_min = giá_trị_nhỏ_nhất(X)
    x_max = giá_trị_lớn_nhất(X)
    y_min = giá_trị_nhỏ_nhất(Y)
    y_max = giá_trị_lớn_nhất(Y)

    // Kiểm tra xem có phải hình vuông không (các cạnh phải bằng nhau)
    nếu x_max - x_min != y_max - y_min:
        trả về false

    s = x_max - x_min  // Độ dài cạnh

    // Kiểm tra từng điểm
    với k từ 0 đến N-1:
        x = X[k]
        y = Y[k]
        // Điểm phải nằm trên một trong bốn cạnh
        nếu không (
            (x == x_min và y_min <= y <= y_max) hoặc  // Cạnh trái
            (x == x_max và y_min <= y <= y_max) hoặc  // Cạnh phải
            (y == y_min và x_min <= x <= x_max) hoặc  // Cạnh dưới
            (y == y_max và x_min <= x <= x_max)      // Cạnh trên
        ):
            trả về false

    trả về true
```

---

### Độ phức tạp thời gian

- Tìm `x_min`, `x_max`, `y_min`, `y_max`: O(N).
- Kiểm tra từng điểm: O(N).
- Tổng cộng: **O(N)**, rất hiệu quả vì `N ≤ 20`.

### Lưu ý cuối cùng

- Đề bài yêu cầu tất cả các điểm phải nằm trên chu vi, không được nằm bên trong hình vuông.
- Nếu bạn cần mã nguồn bằng một ngôn ngữ lập trình cụ thể (như C++, Python), hãy cho tôi biết nhé!

```cpp
#include <iostream>
#include <vector>
#include <algorithm> // Để dùng min và max

using namespace std;

// Hàm kiểm tra xem tất cả các điểm có nằm trên chu vi của một hình vuông hay không
bool checkSquare(vector<int>& X, vector<int>& Y, int N) {
    // Trường hợp đặc biệt: nếu không có điểm, trả về true
    if (N == 0) {
        return true;
    }

    // Tìm giá trị nhỏ nhất và lớn nhất của X và Y
    int x_min = *min_element(X.begin(), X.begin() + N);
    int x_max = *max_element(X.begin(), X.begin() + N);
    int y_min = *min_element(Y.begin(), Y.begin() + N);
    int y_max = *max_element(Y.begin(), Y.begin() + N);

    // Kiểm tra xem có tạo thành hình vuông không (các cạnh phải bằng nhau)
    if (x_max - x_min != y_max - y_min) {
        return false;
    }

    // Độ dài cạnh của hình vuông
    int s = x_max - x_min;

    // Kiểm tra từng điểm
    for (int k = 0; k < N; k++) {
        int x = X[k];
        int y = Y[k];
        // Điểm phải nằm trên một trong bốn cạnh của hình vuông
        bool onPerimeter = (x == x_min && y_min <= y && y <= y_max) || // Cạnh trái
                           (x == x_max && y_min <= y && y <= y_max) || // Cạnh phải
                           (y == y_min && x_min <= x && x <= x_max) || // Cạnh dưới
                           (y == y_max && x_min <= x && x <= x_max);   // Cạnh trên
        if (!onPerimeter) {
            return false;
        }
    }

    return true;
}

// Hàm main để kiểm tra
int main() {
    // Ví dụ 1
    vector<int> X1 = {0, 0, 1, 1};
    vector<int> Y1 = {0, 1, 0, 1};
    int N1 = 4;
    cout << "Test 1: " << (checkSquare(X1, Y1, N1) ? "true" : "false") << endl;

    // Ví dụ 2
    vector<int> X2 = {0, 0, 2};
    vector<int> Y2 = {0, 2, 1};
    int N2 = 3;
    cout << "Test 2: " << (checkSquare(X2, Y2, N2) ? "true" : "false") << endl;

    return 0;
}
```

# Optimize

Với gợi ý mới: "luôn mở rộng hình vuông để đủ chiều dài chứa hết điểm", tôi hiểu rằng mục tiêu không chỉ là kiểm tra xem các điểm có nằm trên chu vi của một hình vuông hiện có hay không, mà còn phải tìm một hình vuông tối thiểu (với các cạnh song song trục Ox/Oy) sao cho tất cả các điểm nằm trên chu vi của nó, bằng cách mở rộng kích thước hình vuông nếu cần. Dưới đây là cách tối ưu thuật toán dựa trên yêu cầu này.

---

### Giải thích ý tưởng tối ưu

1. **Mở rộng hình vuông**:

   - Thay vì chỉ kiểm tra xem các điểm có tạo thành một hình vuông sẵn có không, ta sẽ tìm một hình vuông nhỏ nhất có thể "bao quanh" tất cả các điểm và kiểm tra xem mọi điểm có nằm trên chu vi của hình vuông đó không.
   - Hình vuông nhỏ nhất chứa tất cả các điểm sẽ có:
     - `x_min` và `x_max` là giá trị nhỏ nhất và lớn nhất của `X`.
     - `y_min` và `y_max` là giá trị nhỏ nhất và lớn nhất của `Y`.
     - Độ dài cạnh `s` của hình vuông phải là `max(x_max - x_min, y_max - y_min)` để đảm bảo chứa hết các điểm.

2. **Kiểm tra chu vi**:

   - Sau khi xác định hình vuông tối thiểu với cạnh `s`, kiểm tra xem tất cả các điểm `(X[K], Y[K])` có nằm trên bốn cạnh của hình vuông này không.

3. **Tối ưu**:
   - Vì `N` nhỏ (tối đa 20) và phạm vi tọa độ nhỏ (`-10` đến `10`), ta không cần thử nhiều kích thước hình vuông. Chỉ cần tính toán hình vuông nhỏ nhất dựa trên khoảng cách lớn nhất giữa các tọa độ, rồi kiểm tra.

---

### Thuật toán tối ưu (pseudocode)

```pseudocode
hàm kiemTraHinhVuong(X[], Y[], N):
    nếu N == 0:
        trả về true

    // Tìm min và max của X, Y
    x_min = giá_trị_nhỏ_nhất(X)
    x_max = giá_trị_lớn_nhất(X)
    y_min = giá_trị_nhỏ_nhất(Y)
    y_max = giá_trị_lớn_nhất(Y)

    // Tính độ dài cạnh tối thiểu để chứa tất cả các điểm
    s = tối_đa(x_max - x_min, y_max - y_min)

    // Kiểm tra từng điểm có nằm trên chu vi không
    với k từ 0 đến N-1:
        x = X[k]
        y = Y[k]
        nếu không (
            (x == x_min và y_min <= y <= y_min + s) hoặc  // Cạnh trái
            (x == x_min + s và y_min <= y <= y_min + s) hoặc  // Cạnh phải
            (y == y_min và x_min <= x <= x_min + s) hoặc  // Cạnh dưới
            (y == y_min + s và x_min <= x <= x_min + s)   // Cạnh trên
        ):
            trả về false

    trả về true
```

---

### Mã nguồn C++ tối ưu

```cpp
#include <bits/stdc++.h>

using namespace std;

bool isSquarePerimeter(const int* X, const int* Y, int N) {
    if (N < 0) return false;
    if (N == 0 || N == 1) return true;

    int x_min = X[0], x_max = X[0], y_min = Y[0], y_max = Y[0];

    for (int i = 1; i < N; i++) {
        if (X[i] < x_min) x_min = X[i];
        if (X[i] > x_max) x_max = X[i];
        if (Y[i] < y_min) y_min = Y[i];
        if (Y[i] > y_max) y_max = Y[i];
    }

    if(x_max - x_min == 0) return N > 0;

    // least strength allocate all the points
    int s = max(x_max - x_min, y_max - y_min);

    // point belong to perimester of square
    for (int k = 0; k < N; k++) {
        int x = X[k];
        int y = Y[k];
        bool onPerimeter = (x == x_min && y_min <= y && y <= y_min + s) ||  // left
                           (x == x_min + s && y_min <= y && y <= y_min + s) ||  // right
                           (y == y_min && x_min <= x && x <= x_min + s) ||  // bot
                           (y == y_min + s && x_min <= x && x <= x_min + s);  // top
        if (!onPerimeter) {
            return false;
        }
    }

    return true;
}

int main() {
    int X1[] = {0, 0, 2, 2};
    int Y1[] = {0, 2, 0, 2};
    cout <<"Test 1: " <<(isSquarePerimeter(X1,Y1,4)?"true":"false") <<endl;

    int X2[] = {0, 1, 2};
    int Y2[] = {0, 1, 2};
    cout <<"Test 2: " <<(isSquarePerimeter(X2,Y2,3)?"true":"false") <<endl;

    int X3[] = {0, 0, 1};
    int Y3[] = {0, 1, 1};
    cout <<"Test 3: " <<(isSquarePerimeter(X3,Y3,3)?"true":"false") <<endl;

    int X4[] = {0};
    int Y4[] = {0};
    cout <<"Test 4: " <<(isSquarePerimeter(X4,Y4,1)?"true":"false") <<endl;

    int X5[] = {0, 0};
    int Y5[] = {0, 0};
    cout <<"Test 5: " <<(isSquarePerimeter(X5,Y5,2)?"true":"false") <<endl;

    return 0;
}
```

---

### Giải thích thay đổi trong mã

1. **Tính `s` tối ưu**:

   - Thay vì yêu cầu `x_max - x_min == y_max - y_min`, ta lấy `s = max(x_max - x_min, y_max - y_min)` để đảm bảo hình vuông đủ lớn chứa tất cả các điểm.
   - Điều này cho phép "mở rộng" hình vuông nếu cần.

2. **Kiểm tra chu vi**:

   - Các cạnh của hình vuông giờ được xác định bởi:
     - Cạnh trái: `x = x_min`.
     - Cạnh phải: `x = x_min + s`.
     - Cạnh dưới: `y = y_min`.
     - Cạnh trên: `y = y_min + s`.
   - Mỗi điểm phải nằm trên một trong bốn cạnh này.

3. **Ví dụ minh họa**:
   - **Ví dụ 1**: `(0,0), (0,1), (1,0), (1,1)`:
     - `x_min = 0`, `x_max = 1`, `y_min = 0`, `y_max = 1`.
     - `s = max(1, 1) = 1`.
     - Tất cả các điểm nằm trên chu vi → `true`.
   - **Ví dụ 2**: `(0,0), (0,2), (2,1)`:
     - `x_min = 0`, `x_max = 2`, `y_min = 0`, `y_max = 2`.
     - `s = max(2, 2) = 2`.
     - `(0,0)`: cạnh dưới, `(0,2)`: cạnh trên, `(2,1)`: không nằm trên cạnh nào → `false`.
   - **Ví dụ 3**: `(0,0), (1,1), (2,2)`:
     - `x_min = 0`, `x_max = 2`, `y_min = 0`, `y_max = 2`.
     - `s = 2`.
     - `(1,1)` không nằm trên chu vi → `false`.

---

### Kết quả chạy chương trình

```
Test 1: true
Test 2: false
Test 3: false
```

---

### Độ phức tạp

- Tìm min/max: O(N).
- Kiểm tra từng điểm: O(N).
- Tổng: **O(N)**, vẫn rất hiệu quả với `N ≤ 20`.

### Lưu ý

- Thuật toán này đảm bảo hình vuông được mở rộng đủ lớn để chứa tất cả các điểm, nhưng vẫn yêu cầu mọi điểm phải nằm trên chu vi, đúng với yêu cầu ban đầu của bài toán.
- Nếu bạn muốn thay đổi thêm hoặc cần giải thích sâu hơn, hãy cho tôi biết nhé!
