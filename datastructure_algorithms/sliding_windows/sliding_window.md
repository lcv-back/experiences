# Basic về Sliding Windows

### Sliding Window là gì?

Sliding Window là một kỹ thuật thường được sử dụng trong các bài toán xử lý mảng, chuỗi (string), hoặc danh sách để tối ưu hóa việc tìm kiếm hoặc tính toán trên một tập hợp con liên tục của dữ liệu. Thay vì duyệt toàn bộ dữ liệu nhiều lần, kỹ thuật này sử dụng một "cửa sổ" có kích thước cố định hoặc thay đổi, di chuyển qua dữ liệu để giải quyết vấn đề một cách hiệu quả.

Kỹ thuật này đặc biệt hữu ích trong các bài toán như:

- Tìm `chuỗi con dài nhất/nhỏ nhất` thỏa mãn điều kiện.
- Tính `tổng/tích của các phần tử trong một đoạn liên tục.`
- Giải các bài toán `tối ưu hóa thời gian từ O(n²) xuống O(n).`

Có hai loại chính:

1. **Fixed-size Sliding Window**: Cửa sổ có kích thước cố định, di chuyển qua dữ liệu.
2. **Variable-size Sliding Window**: Cửa sổ thay đổi kích thước linh hoạt dựa trên điều kiện bài toán.

---

### Ví dụ: Fixed-size Sliding Window

Giả sử bạn cần tìm tổng lớn nhất của một đoạn con có độ dài `k` trong một mảng số nguyên.

#### Code Golang:

```go
package main

import (
	"fmt"
)

func maxSumSubarray(arr []int, k int) int {
	// Kiểm tra điều kiện đầu vào
	if len(arr) < k || k <= 0 {
		return 0
	}

	// Tính tổng của cửa sổ đầu tiên
	maxSum := 0
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	maxSum = windowSum

	// Di chuyển cửa sổ và cập nhật tổng lớn nhất
	for i := k; i < len(arr); i++ {
		// Bỏ phần tử đầu của cửa sổ trước đó, thêm phần tử mới
		windowSum = windowSum - arr[i-k] + arr[i]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

func main() {
	arr := []int{1, 4, 2, 10, 23, 3, 1, 0, 20}
	k := 4
	result := maxSumSubarray(arr, k)
	fmt.Printf("Tổng lớn nhất của đoạn con độ dài %d là: %d\n", k, result)
}
```

#### Giải thích:

1. **Khởi tạo**: Tính tổng của `k` phần tử đầu tiên (cửa sổ đầu tiên).
2. **Di chuyển cửa sổ**: Với mỗi bước, loại bỏ phần tử đầu tiên của cửa sổ cũ (`arr[i-k]`) và thêm phần tử mới (`arr[i]`).
3. **Cập nhật kết quả**: So sánh và giữ tổng lớn nhất.

**Đầu ra**:

```
Tổng lớn nhất của đoạn con độ dài 4 là: 36
```

(36 đến từ đoạn `[4, 2, 10, 23]`).

**Độ phức tạp**: O(n), với n là độ dài mảng.

---

### Ví dụ: Variable-size Sliding Window

Giả sử bạn cần tìm độ dài chuỗi con dài nhất trong một chuỗi ký tự mà không có ký tự lặp lại.

#### Code Golang:

```go
package main

import (
	"fmt"
)

func longestSubstringWithoutRepeating(s string) int {
	// Map để lưu vị trí cuối cùng của ký tự
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0

	// Duyệt qua chuỗi
	for i := 0; i < len(s); i++ {
		// Nếu ký tự đã xuất hiện, cập nhật điểm bắt đầu của cửa sổ
		if lastIndex, found := charIndex[s[i]]; found && lastIndex >= start {
			start = lastIndex + 1
		} else {
			// Nếu không lặp, kiểm tra độ dài cửa sổ hiện tại
			currentLength := i - start + 1
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
		// Cập nhật vị trí cuối cùng của ký tự hiện tại
		charIndex[s[i]] = i
	}

	return maxLength
}

func main() {
	s := "abcabcbb"
	result := longestSubstringWithoutRepeating(s)
	fmt.Printf("Độ dài chuỗi con không lặp dài nhất: %d\n", result)
}
```

#### Giải thích:

1. **Map lưu vị trí**: Dùng `map` để ghi lại vị trí cuối cùng của mỗi ký tự.
2. **Cửa sổ linh hoạt**: Nếu gặp ký tự lặp, di chuyển điểm bắt đầu (`start`) đến sau vị trí lặp gần nhất.
3. **Tính độ dài**: Cập nhật `maxLength` khi cửa sổ mở rộng mà không có ký tự lặp.

**Đầu ra**:

```
Độ dài chuỗi con không lặp dài nhất: 3
```

(3 đến từ chuỗi con `"abc"`).

**Độ phức tạp**: O(n), với n là độ dài chuỗi.

---

### Khi nào dùng Sliding Window?

- Khi bài toán yêu cầu xử lý trên các đoạn con liên tục.
- Khi có thể tránh lặp lại tính toán bằng cách tận dụng kết quả của cửa sổ trước đó.
- Khi cần tối ưu hóa thời gian từ duyệt lồng nhau (O(n²)) xuống tuyến tính (O(n)).
