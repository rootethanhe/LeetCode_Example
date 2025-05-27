package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total%2 == 1 { // 总长度为奇数，找中间一个数
		return float64(findKth(nums1, nums2, (total+1)/2))
	} else {
		return float64(findKth(nums1, nums2, total/2)+findKth(nums1, nums2, total/2+1)) / 2.0
	}
}

// 寻找两个数组中第k小的元素
func findKth(nums1 []int, nums2 []int, k int) int {
	m, n := len(nums1), len(nums2)
	// 总假设nums1是较短的数组，若不符合则交换(优化时间复杂度)
	if m > n {
		return findKth(nums2, nums1, k)
	}

	// 递归终止条件：当nums1为空或k=1时直接处理
	if m == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}

	// 计算两个数组中间位置
	i := min(m, k/2)
	j := min(n, k/2)

	// 比较中间元素，排除较小者的前j个元素
	if nums1[i-1] < nums2[j-1] {
		return findKth(nums1[i:], nums2, k-i)
	} else {
		return findKth(nums1, nums2[j:], k-j)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 示例测试
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Printf("示例1输出: %.5f\n", findMedianSortedArrays(nums1, nums2)) // 输出: 2.00000

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Printf("示例2输出: %.5f\n", findMedianSortedArrays(nums1, nums2)) // 输出: 2.50000

	// 边界测试：一个数组为空
	nums1 = []int{}
	nums2 = []int{2}
	fmt.Printf("边界测试1输出: %.5f\n", findMedianSortedArrays(nums1, nums2)) // 输出: 2.00000

	// 边界测试：所有元素相同
	nums1 = []int{2, 2}
	nums2 = []int{2, 2}
	fmt.Printf("边界测试2输出: %.5f\n", findMedianSortedArrays(nums1, nums2)) // 输出: 2.00000
}
