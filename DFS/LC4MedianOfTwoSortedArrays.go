package main

import "math"

// link: https://leetcode.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totLen := len(nums1) + len(nums2)
	if totLen % 2 == 0 {
		return (float64(findKth(nums1, 0, nums2, 0, totLen / 2)) +
			float64(findKth(nums1, 0, nums2, 0, totLen / 2 + 1))) / 2.0
	} else {
		return float64(findKth(nums1, 0, nums2, 0, (totLen + 1) / 2.0))
	}
}

/*
寻找nums1,nums2两个数组合在一起的第k个元素
每次扔掉 k/2 个元素
 */
func findKth(nums1 []int, idx1 int,
			nums2 []int, idx2 int,
			k int) int {
	// 第x个数的下标是x-1

	/*
	递归边界1:
	某个数组的数都扔完了 直接返回另一个数组的第start_idx + k - 1个元素
	 */
	if idx1 >= len(nums1) {
		return nums2[idx2 + k - 1]
	}
	if idx2 >= len(nums2) {
		return nums1[idx1 + k - 1]
	}

	/*
	递归边界2:
	只需要取一个数 k == 1 直接返回nums1和nums2里较小的数
	经过边界条件1 两个数组保证都有元素可取
	 */
	if k == 1 {
		return Min(nums1[idx1], nums2[idx2])
	}

	/*
	下面的两个比较 比较的是A和B的start_idx + k/2 - 1的位置哪个数字小 谁小就把谁的前k/2个数字扔掉 因为这k/2个数肯定不会包含结果
	注意: 从start_idx开始的第k/2个数的下标为: start + k/2 - 1
	如果start_idx + k/2 - 1 < length 为什么返回正最大？-> 让另一个数组扔数
	比如: A有2个数 B有101个数 中位数是两数组合并后第52个数 每次应该从A或者B里面挑52/2 = 26个数字扔掉
	但是A没有这么多数可以扔 那就扔B的26个数就可以
	为什么可以这么做？因为求的是第52个数字 B扔掉26个 也不会扔掉可能的结果 所以当一个数组数字不够需要扔的数时 选择扔对面的数组即可
	会不会出现都不够的情况？不会 因为k是合法的 在都不够扔之前肯定就找到第k个数了
	*/
	halfNums1 := math.MaxInt32
	halfNums2 := math.MaxInt32
	if idx1 + k / 2 - 1 < len(nums1) {
		halfNums1 = nums1[idx1 + k / 2 - 1]
	}
	if idx2 + k / 2 - 1 < len(nums2) {
		halfNums2 = nums2[idx2 + k / 2 - 1]
	}

	if halfNums1 <= halfNums2 {
		return findKth(nums1, idx1 + k / 2,  nums2, idx2, k - k / 2)
	} else {
		return findKth(nums1, idx1, nums2, idx2+k/2, k-k/2)
	}
}

func Min (a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}