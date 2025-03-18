Thuật toán Kadane là một thuật toán nổi tiếng để giải quyết bài toán **tìm dãy con có tổng lớn nhất** (Maximum Subarray Problem) trong một mảng số. Đây là một thuật toán tối ưu với độ phức tạp thời gian **O(n)**, rất hiệu quả và được sử dụng rộng rãi trong lập trình thi đấu cũng như các ứng dụng thực tế. Dưới đây là tất tần tật về thuật toán Kadane:

---

### **1. Bài toán**

- **Đầu vào**: Một mảng số nguyên (có thể chứa số âm, số dương hoặc số 0).
- **Đầu ra**: Tổng lớn nhất của một dãy con liên tục (subarray) trong mảng đó.
- **Ví dụ**:
  - Mảng: `[-2, 1, -3, 4, -1, 2, 1, -5, 4]`
  - Dãy con có tổng lớn nhất: `[4, -1, 2, 1]` → Tổng = `6`.

---

### **2. Ý tưởng cơ bản**

Thuật toán Kadane sử dụng **quy hoạch động** (Dynamic Programming) theo cách tiếp cận tham lam (greedy). Ý tưởng chính là:

- Duy trì hai biến:
  - **maxSoFar**: Lưu tổng lớn nhất của dãy con tìm được cho đến hiện tại.
  - **maxEndingHere**: Lưu tổng lớn nhất của dãy con kết thúc tại phần tử hiện tại.
- Tại mỗi phần tử, quyết định:
  - Hoặc bắt đầu một dãy con mới từ phần tử đó.
  - Hoặc thêm phần tử đó vào dãy con hiện tại nếu tổng vẫn dương.
- Nếu tổng tạm thời (`maxEndingHere`) trở thành âm, reset nó về 0 vì một dãy con bắt đầu từ đây không thể đóng góp vào tổng lớn hơn.

---

### **3. Cách hoạt động**

Dưới đây là thuật toán Kadane cơ bản:

1. Khởi tạo:
   - `maxSoFar = mảng[0]` (hoặc một giá trị nhỏ nhất nếu mảng không rỗng).
   - `maxEndingHere = 0`.
2. Duyệt qua từng phần tử của mảng:
   - Cập nhật `maxEndingHere = max(0, maxEndingHere + mảng[i])`.
   - Cập nhật `maxSoFar = max(maxSoFar, maxEndingHere)`.
3. Kết quả cuối cùng là `maxSoFar`.

---

### **4. Ví dụ minh họa**

Mảng: `[-2, 1, -3, 4, -1, 2, 1, -5, 4]`

| Bước | Phần tử | maxEndingHere (trước) | maxEndingHere (sau) | maxSoFar |
| ---- | ------- | --------------------- | ------------------- | -------- |
| 0    | -2      | 0                     | 0                   | -2       |
| 1    | 1       | 0                     | 1                   | 1        |
| 2    | -3      | 1                     | 0                   | 1        |
| 3    | 4       | 0                     | 4                   | 4        |
| 4    | -1      | 4                     | 3                   | 4        |
| 5    | 2       | 3                     | 5                   | 5        |
| 6    | 1       | 5                     | 6                   | 6        |
| 7    | -5      | 6                     | 1                   | 6        |
| 8    | 4       | 1                     | 5                   | 6        |

- Kết quả: `maxSoFar = 6`, dãy con tương ứng là `[4, -1, 2, 1]`.

---

### **5. Code triển khai**

#### **Phiên bản cơ bản (chỉ tìm tổng lớn nhất)**

```go
func maxSubArray(arr []int) int {
    if len(arr) == 0 {
        return 0
    }
    maxSoFar := arr[0]
    maxEndingHere := 0

    for _, num := range arr {
        maxEndingHere = max(maxEndingHere+num, 0)
        maxSoFar = max(maxSoFar, maxEndingHere)
    }
    return maxSoFar
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

#### **Phiên bản mở rộng (tìm cả dãy con)**

Nếu bạn muốn lấy cả chỉ số bắt đầu và kết thúc của dãy con:

```go
func maxSubArrayWithIndices(arr []int) (int, int, int) {
    if len(arr) == 0 {
        return 0, 0, 0
    }
    maxSoFar := arr[0]
    maxEndingHere := 0
    start := 0
    end := 0
    tempStart := 0

    for i, num := range arr {
        maxEndingHere += num
        if maxEndingHere > maxSoFar {
            maxSoFar = maxEndingHere
            start = tempStart
            end = i
        }
        if maxEndingHere < 0 {
            maxEndingHere = 0
            tempStart = i + 1
        }
    }
    return maxSoFar, start, end
}
```

---

### **6. Độ phức tạp**

- **Thời gian**: **O(n)** – chỉ cần duyệt qua mảng một lần.
- **Không gian**: **O(1)** – chỉ dùng một vài biến, không cần mảng phụ.

---

### **7. Ưu điểm**

- Hiệu quả cao, không cần đệ quy hay lưu trữ trạng thái phức tạp.
- Dễ triển khai và mở rộng (ví dụ: tìm chỉ số, xử lý mảng 2D).
- Hoạt động tốt với cả mảng chứa số âm và dương.

---

### **8. Nhược điểm**

- **Không tìm tất cả dãy con**: Chỉ tìm một dãy con có tổng lớn nhất, không liệt kê tất cả các khả năng (nếu cần thì phải sửa đổi thuật toán).
- **Không tối ưu cho trường hợp đặc biệt**: Nếu mảng toàn số âm, phiên bản cơ bản cần điều chỉnh để trả về số lớn nhất trong mảng.

#### **Điều chỉnh khi mảng toàn số âm**

Nếu mảng chỉ chứa số âm, khởi tạo `maxSoFar` bằng phần tử đầu tiên và bỏ logic reset `maxEndingHere` về 0:

```go
func maxSubArrayAllNegative(arr []int) int {
    if len(arr) == 0 {
        return 0
    }
    maxSoFar := arr[0]
    maxEndingHere := arr[0]

    for i := 1; i < len(arr); i++ {
        maxEndingHere = max(arr[i], maxEndingHere+arr[i])
        maxSoFar = max(maxSoFar, maxEndingHere)
    }
    return maxSoFar
}
```

---

### **9. Ứng dụng thực tế**

- **Phân tích tài chính**: Tìm khoảng thời gian có lợi nhuận lớn nhất từ dữ liệu giá cổ phiếu.
- **Xử lý tín hiệu**: Xác định đoạn tín hiệu có năng lượng cao nhất.
- **Trò chơi và đồ họa**: Tối ưu hóa các bài toán liên quan đến điểm số hoặc vùng dữ liệu.

---

### **10. Mở rộng: Kadane cho mảng 2D**

Kadane cũng có thể được áp dụng cho mảng 2D để tìm hình chữ nhật con có tổng lớn nhất. Ý tưởng là:

1. Nén mảng 2D thành mảng 1D bằng cách cộng các cột hoặc hàng trong một phạm vi.
2. Áp dụng Kadane 1D lên mảng tạm thời.

- Độ phức tạp: **O(n³)** cho ma trận kích thước `n x n`.

---

### **Kết luận**

Thuật toán Kadane là một công cụ mạnh mẽ và đơn giản để giải quyết bài toán tìm dãy con có tổng lớn nhất. Với hiệu suất cao và tính ứng dụng rộng rãi, nó là một trong những thuật toán cơ bản mà mọi lập trình viên nên nắm vững. Nếu bạn cần thêm ví dụ, giải thích chi tiết hơn hoặc muốn áp dụng nó vào một bài toán cụ thể, hãy cho tôi biết nhé!

### Example Problem

[121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/?envType=problem-list-v2&envId=array)

```plaintext
explain solution:
- key insight: kadane algorithm
- in kadane, have maxSoFar is sum so far and maxEndingHere is sum at here
- in stock, have buy and profit
- buy initialize is the first element on array, profit is 0
- for each iterator in array, check buy encountained the lower price, update buy to lower price
- concurrency check if price in current time *minus* buy greater than current profit, update the current profit to prices at current *minus* buy (buy prices at before time)
- profit is what we need return
```

```cpp
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int profit = 0;
        int buy = prices[0];

        for(int i=0; i<prices.size(); ++i){
            if(prices[i] < buy) buy = prices[i];
            else if(prices[i] - buy - profit > 0) profit = prices[i] - buy;
        }

        return profit;
    }
};
```
