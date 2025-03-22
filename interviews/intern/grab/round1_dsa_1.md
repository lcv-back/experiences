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

```
bool solution(vector<int>& X, vector<int>& Y) {
    int N = X.size();

    int minX = *min_element(X.begin(), X.end());
    int maxX = *max_element(X.begin(), X.end());
    int minY = *min_element(Y.begin(), Y.end());
    int maxY = *max_element(Y.begin(), Y.end());

    int s1 = maxX - minX;
    int s2 = maxY - minY;
    int s = max(s1, s2);

    if (s == 0) return true;

    set<double> centerX_candidates, centerY_candidates;
    for (int i = 0; i < N; i++) {
        centerX_candidates.insert(X[i] + s / 2.0);
        centerX_candidates.insert(X[i] - s / 2.0);
        centerY_candidates.insert(Y[i] + s / 2.0);
        centerY_candidates.insert(Y[i] - s / 2.0);
    }

    for (double centerX : centerX_candidates) {
        for (double centerY : centerY_candidates) {
            double left = centerX - s / 2.0;
            double right = centerX + s / 2.0;
            double bottom = centerY - s / 2.0;
            double top = centerY + s / 2.0;

            bool allOnPerimeter = true;
            for (int i = 0; i < N; i++) {
                bool onPerimeter = false;
                if (abs(X[i] - left) < 1e-6 && Y[i] >= bottom && Y[i] <= top) onPerimeter = true;
                else if (abs(X[i] - right) < 1e-6 && Y[i] >= bottom && Y[i] <= top) onPerimeter = true;
                else if (abs(Y[i] - bottom) < 1e-6 && X[i] >= left && X[i] <= right) onPerimeter = true;
                else if (abs(Y[i] - top) < 1e-6 && X[i] >= left && X[i] <= right) onPerimeter = true;

                if (!onPerimeter) {
                    allOnPerimeter = false;
                    break;
                }
            }
            if (allOnPerimeter) return true;
        }
    }

    return false;
}
// accept
```
