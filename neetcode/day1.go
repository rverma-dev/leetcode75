package neetcode

import (
	"sort"

	"leetcode75/ds"
)

// in: []int{1, 7, 3, 6, 5, 6}
// in: []int{1, 7, 3, 6, 5, 6, 2, 7}
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	return nums
}

func pivotIndex(nums []int) int {
	var suml, sumr int
	for i := 0; i < len(nums); i++ {
		sumr += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		sumr -= nums[i]
		if sumr == suml {
			return i
		}
		suml += nums[i]

	}
	return -1
}

func minPartitions(n string) int {
	var max int
	for i := 0; i < len(n); i++ {
		if int(n[i])-48 > max {
			max = int(n[i]) - 48
		}
	}
	return max
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1 //nums2 is small
	}

	var maxLeft, minRight int
	left, right, half := 0, len(nums1), (len(nums1)+len(nums2)+1)/2 //+1 to make compatible with both even and odd no of elements
	var i, j int
	for left <= right {
		i = (left + right) / 2
		j = half - i
		if i < len(nums1) && nums2[j-1] > nums1[i] {
			left = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			right = i - 1
		} else {
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else if nums1[i-1] > nums2[j-1] {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = nums2[j-1]
			}

			if (len(nums1)+len(nums2))%2 != 0 {
				return float64(maxLeft)
			}

			if i == len(nums1) {
				minRight = nums2[j]
			} else if j == len(nums2) {
				minRight = nums1[i]
			} else if nums1[i] < nums2[j] {
				minRight = nums1[i]
			} else {
				minRight = nums2[j]
			}
			return float64(maxLeft+minRight) / 2
		}
	}
	return 0
}

func isIsomorphic(s string, t string) bool {
	st := make(map[byte]byte)
	ts := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		target, ok := st[s[i]]
		source, sok := ts[t[i]]
		if !ok && !sok {
			st[s[i]] = t[i]
			ts[t[i]] = s[i]
		} else if !(source == s[i] && target == t[i]) {
			return false
		}

	}
	return true
}

func hash(s string) string {
	freq := [26]byte{}
	for _, v := range s {
		freq[v-'a']++
	}
	return string(freq[:])
}

func sortStr(s string) string {
	runeSlice := []rune(s)
	sort.Slice(runeSlice, func(i, j int) bool {
		return runeSlice[i] < runeSlice[j]
	})
	return string(runeSlice)
}

func groupAnagrams(strs []string) [][]string {
	// since the input range to 10^4, int64 is not sufficient to store hash value as prime number multiplication is used
	set := make(map[string][]string)
	for _, str := range strs {
		sortstr := hash(str)
		if lis, ok := set[sortstr]; ok {
			set[sortstr] = append(lis, str)
		} else {
			set[sortstr] = []string{str}
		}
	}
	result := make([][]string, 0, len(set))
	for _, v := range set {
		result = append(result, v)
	}
	return result
}

func topKFrequent(nums []int, k int) []int {
	hm := make(map[int]int)
	for _, v := range nums {
		hm[v]++
	}
	freq := ds.NewHeap(func(a, b []int) bool { return a[1] < b[1] })
	for k, v := range hm {
		freq.Push([]int{v, k})
	}
	result := []int{}
	for i := 0; i < k; i++ {
		val := freq.Pop()
		result = append(result, val[1])
	}
	return result
}

func productExceptSelf(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	// iterate over from left to right
	result[0] = 1
	for i := 1; i < length; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	// iterate over from right to left and multiply with result[i-1]
	right := 1
	for i := length - 1; i >= 0; i-- {
		result[i] = result[i] * right
		right *= nums[i]
	}
	return result
}

// func Test(i int) {
// 	tests := []interface {}{
// 		runningSum([]int{1, 7, 3, 6, 5, 6}),
// 		pivotIndex([]int{1, 7, 3, 6, 5, 6}),
// 		minPartitions("32"),
// 		findMedianSortedArrays([]int{},[]int{1}),
// 		isIsomorphic("paper", "title"),
// 		groupAnagrams([]string{"stop", "pots", "reed", "", "tops", "deer", "opts", ""}),
// 		sortStr("aaaabbbbddddcccc"),
// 		topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2),
// 		productExceptSelf([]int{1, 2, 3, 4}),
// 	}

// 	fmt.Println(tests[i])
// }
