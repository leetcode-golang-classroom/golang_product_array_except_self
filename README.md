# golang_product_array_except_self

Given an integer array `nums`, return *an array* `answer` *such that* `answer[i]` *is equal to the product of all the elements of* `nums` *except* `nums[i]`.

The product of any prefix or suffix of `nums` is **guaranteed** to fit in a **32-bit** integer.

You must write an algorithm that runs in `O(n)` time and without using the division operation.

## Examples

**Example 1:**

```
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

```

**Example 2:**

```
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

```

**Constraints:**

- `2 <= nums.length <= 105`
- `30 <= nums[i] <= 30`
- The product of any prefix or suffix of `nums` is **guaranteed** to fit in a **32-bit** integer.

**Follow up:** Can you solve the problem in `O(1)` extra space complexity? (The output array **does not** count as extra space for space complexity analysis.)

## 解析

給定一個整數陣列 nums 

要求回傳一個陣列 res 每個 res[i] = 除了 nums[i] 外其他數值的 product

對每個 i

ans[i] = nums[0]*..nums[i-1] * nums[i+1]*nums[i+2]..*num[n-1], where n = len(nums)

其中可以分為兩個部份

第一個部份 nums[0]*…nums[i-1]

第二個部份 nums[i+1]*…nums[n-1]


![](https://i.imgur.com/Myh4ddA.png)

這兩個部份可以分別由兩個陣列陣 lower, upper 兩個

lower[0] = 1, upper[n-1]= 1

lower[i] = lower[i-1] * nums[i-1]

upper[n-1-i] = upper[n-i]*nums[n-i]  for  1≤i ≤ n-1

最後 ans[i] = lower[i] * upper[i] for  0≤i ≤ n-1

用這個演算法 時間複雜度是 O(n)

空間複雜度是 O(n)

然而可以發現每次 upper, lower 做連乘其使只會使用一次

所以給可以改成

初始化 ans[i] = 1

lower 跟 upper 每次個別從不同方向做連乘

![](https://i.imgur.com/FtNsGLt.png)
![](https://i.imgur.com/DB8BMLx.png)

初始化 lower = 1, upper = 1

if pos > 0 {

  lower *= nums[pos-1]

  upper *= nums[n-1-pos]

}

ans[pos] *= lower

ans[n-1-pos] *= upper

用這個作法可以省去儲存中間計算的暫存數

時間複雜度是 O(n)

空間複雜度可以優化到 O(1) 

## 程式碼
```go
package sol

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	lower, upper := 1, 1
	for pos := 0; pos < n; pos++ {
		ans[pos] = 1
	}
	for pos := 0; pos < n; pos++ {
		if pos > 0 {
			lower *= nums[pos-1]
			upper *= nums[n-pos]
		}
		ans[pos] *= lower
		ans[n-1-pos] *= upper
	}
	return ans
}

```
## 困難點

1. 需要看出每個連乘位置的關係

## Solve Point

- [x]  初始化 lower = 1, upper = 1, n = len(nums)
- [x]  建立一個陣列 ans 長度為 n 初始化每個值 = 1
- [x]  從 pos = 0..n-1 做以下操作
- [x]  if pos > 0  更新 lower *= nums[pos-1] , upper *= nums[len-pos]
- [x]  更新 ans[pos] *= lower , ans[len-1-pos] *= upper
- [x]  回傳 ans