## Exam 3

There is an array A of N integers and three tiles. Each tile can cover two neighboring numbers from the array but cannot intersect with another tile. It also cannot be placed outside the array, even partially. Given an array A of N integers, returns the maximum sum of numbers that can be covered using at most three tiles.

**Examples:**

- Given A = [2, 3, 5, 2, 3, 4, 6, 4, 1], the function should return 25. There is only one optimal placement of tiles: (3, 5), (3, 4), (6, 4).
- Given A = [1, 5, 3, 2, 6, 6, 10, 4, 7, 2, 1], the function should return 35. One of the three optimal placements of tiles is (5, 3), (6, 10), (4, 7).
- Given A = [1, 2, 3, 2, 3], the function should return 10. There is one optimal placement of tiles: (2, 3), (2, 3). Only two tiles can be used because A is too small to contain another one.
- Given A = [5, 10, 3], the function should return 15. Only one tile can be used.

Write an efficient algorithm for the following assumptions:

- N is an integer within the range [2..100,000];
- Each element of array A is an integer within the range [0..1,000,000].

---

### Giải thích bài toán

Bài toán yêu cầu tìm tổng lớn nhất có thể đạt được khi sử dụng tối đa 3 tấm lát (tiles), mỗi tấm có thể che 2 số liền kề trong mảng, với các điều kiện:

- Các tấm lát không được giao nhau (không có phần tử nào được che bởi nhiều hơn một tấm).
- Tấm lát phải nằm hoàn toàn trong mảng (không được vượt ra ngoài).

Mục tiêu là chọn tối đa 3 cặp số liền kề sao cho tổng của chúng là lớn nhất, đồng thời đảm bảo các cặp này không chồng lấn.

Ví dụ:

1. Với mảng `[2, 3, 5, 2, 3, 4, 6, 4, 1]`:

   - Chọn các cặp: `(3, 5)` (vị trí 1-2), `(3, 4)` (vị trí 4-5), `(6, 4)` (vị trí 6-7).
   - Tổng = 3 + 5 + 3 + 4 + 6 + 4 = 25.

2. Với mảng `[1, 5, 3, 2, 6, 6, 10, 4, 7, 2, 1]`:

   - Chọn các cặp: `(5, 3)` (vị trí 1-2), `(6, 10)` (vị trí 4-5), `(4, 7)` (vị trí 7-8).
   - Tổng = 5 + 3 + 6 + 10 + 4 + 7 = 35.

3. Với mảng `[1, 2, 3, 2, 3]`:

   - Chỉ chọn được 2 cặp: `(2, 3)` (vị trí 1-2), `(2, 3)` (vị trí 3-4).
   - Tổng = 2 + 3 + 2 + 3 = 10.

4. Với mảng `[5, 10, 3]`:
   - Chỉ chọn được 1 cặp: `(5, 10)` (vị trí 0-1).
   - Tổng = 5 + 10 = 15.

### Ý tưởng giải pháp

- Sử dụng quy hoạch động (Dynamic Programming - DP) để giải quyết.
- Tại mỗi vị trí trong mảng, ta có thể:
  1. Không chọn cặp nào (bỏ qua).
  2. Chọn cặp hiện tại (2 số liền kề) và bỏ qua số tiếp theo để tránh giao nhau.
- Vì ta được dùng tối đa 3 tấm lát, ta cần theo dõi số tấm đã sử dụng.
- DP[i][k] sẽ biểu thị tổng lớn nhất có thể đạt được từ vị trí i trở đi khi sử dụng k tấm lát.

### Thuật toán

- **Trạng thái**: `dp[i][k]` là tổng lớn nhất từ vị trí i đến cuối mảng khi dùng k tấm lát.
- **Chuyển trạng thái**:
  - Nếu không chọn cặp tại i: `dp[i][k] = dp[i+1][k]`.
  - Nếu chọn cặp tại i (i và i+1): `dp[i][k] = A[i] + A[i+1] + dp[i+2][k-1]` (nếu k > 0 và i+1 < N).
- **Kết quả**: `dp[0][3]` (tổng lớn nhất khi bắt đầu từ 0 và dùng tối đa 3 tấm).

### Độ phức tạp

- Thời gian: O(N), vì ta điền bảng DP với N trạng thái và mỗi trạng thái xem xét tối đa 4 giá trị k (0, 1, 2, 3).
- Không gian: O(N), dùng mảng 2D `dp[N][4]`.

### Code C++

```cpp
#include <vector>
#include <algorithm>
#include<iostream>

using namespace std;

int solution(vector<int> &arr) {
    int n = arr.size();
    // dp[i][k]: maximum sum from i to last when use k tiles
    vector<vector<long long>> dp(n + 1, vector<long long>(4, 0));

    // right -> left
    for (int i = n - 1; i >= 0; i--) {
        for (int k = 0; k <= 3; k++) {
            // not choose
            dp[i][k] = dp[i + 1][k];

            // i and i+1
            if (k > 0 && i + 1 < n) dp[i][k] = max(dp[i][k], (long long)arr[i] + arr[i + 1] + dp[i + 2][k - 1]);
        }
    }

    return dp[0][3];
}

int main() {
    vector<int> arr1 = {2, 3, 5, 2, 3, 4, 6, 4, 1};
    vector<int> arr2 = {1, 5, 3, 2, 6, 6, 10, 4, 7, 2, 1};
    vector<int> arr3 = {1, 2, 3, 2, 3};
    vector<int> arr4 = {5, 10, 3};

    cout << solution(arr1) << endl;
    cout << solution(arr2) << endl;
    cout << solution(arr3) << endl;
    cout << solution(arr4) << endl;
    return 0;
}
```

### Giải thích code

- Dùng `long long` để tránh tràn số vì giá trị tối đa của mỗi phần tử là 1,000,000 và tổng có thể lớn.
- Mảng `dp` có kích thước `(N+1) x 4`, trong đó `dp[N][k] = 0` (khi vượt quá mảng).
- Duyệt từ cuối mảng về đầu, tại mỗi vị trí i, tính toán cho k = 0, 1, 2, 3.
- So sánh giữa không chọn và chọn cặp để lấy giá trị tối ưu.

Hy vọng giải thích và code trên rõ ràng! Nếu bạn cần thêm chi tiết, cứ hỏi nhé!
