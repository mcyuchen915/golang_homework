package homework

import "sort"

// 1. 只出现一次的数字
/*
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
*/
func SingleNumber(nums []int) int {
	// 异或性质：a^a = 0，a^0 = a，且异或满足交换律/结合律。
	// 成对出现的元素异或后互相抵消为 0，全部异或完剩下的就是只出现一次的元素。
	// 时间 O(n)，额外空间 O(1)。
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

// 非法数组判断并不全面
// func SingleNumber(nums []int) (int) {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	flagMap := make(map[int]int)
// 	for _, num := range nums {
// 		if flagMap[num] != 0 {
// 			flagMap[num]++
// 		} else {
// 			flagMap[num] = 1
// 		}
// 	}
// 	for key, value := range flagMap {
// 		if value == 1 {
// 			return key
// 		}
// 	}
// 	return 0
// }

// 2. 回文数
/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。-121也不是
需注意反转后数字的溢出问题
*/
func IsPalindrome(x int) bool {
	// 负数不是回文数；末尾为0的非零数也不是回文数*/
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	if x < 10 {
		return true
	}
	if x < 100 {
		return x%10 == x/10
	}
	/*
		反转一半数字,把 x 整体反转，得到 rev，然后判断 rev 和 x 是否相等。
		但实际上，对于回文数来说，只需判断 x 的左半部分是否等于右半部分就行
		rev=x，说明 x 是（长度为偶数的）回文数,rev=x/10，说明 x 是（长度为奇数的）回文数。
		同时,只反转一半,rev 的量级就不会超过x,就不会溢出。
	*/
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	return x == revertedNumber || x == revertedNumber/10
}

// 3. 有效的括号
/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
*/
func IsValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s)%2 != 0 {
		return false
	}
	/*
		使用栈来存储左括号，右括号来匹配左括号
	*/
	stack := make([]rune, 0, len(s)/2)
	/*
		使用map来存储右括号和左括号的映射关系
	*/
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	// 遍历字符串，如果遇到左括号，则压入栈，如果遇到右括号，则弹出栈顶元素，如果栈顶元素不等于右括号的映射关系，则返回false
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 4. 最长公共前缀
/**
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*/
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	// 取第一个字符串作为公共前缀
	prefix := strs[0]
	/*
		遍历字符串数组，比较每个字符串的公共前缀
		如果当前字符串长度小于公共前缀，则截取当前字符串长度
		如果当前字符串长度大于公共前缀，则截取公共前缀长度
		如果当前字符串长度等于公共前缀，则不截取
		如果当前字符串长度不等于公共前缀，则截取公共前缀长度
	*/
	for _, str := range strs {
		if len(str) < len(prefix) {
			prefix = prefix[:len(str)]
		}
		for i := 0; i < len(prefix); i++ {
			if str[i] != prefix[i] {
				prefix = prefix[:i]
				break
			}
		}
	}
	return prefix
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	// 如果数组为空，则返回空数组
	if len(digits) == 0 {
		return nil
	}
	for i := len(digits) - 1; i >= 0; i-- {
        if digits[i] != 9 {
            digits[i]++
            return digits
        }
        digits[i] = 0
    }
    // 全部是 9，数组长度加 1
	digits = append(digits, 0)
    digits[0] = 1
	return digits
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
/*
指引：可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，
	  当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
*/
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		// 如果当前元素和前一个元素不相等，则将当前元素复制到前一个元素的下一个位置
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	// 返回去重后的数组长度
	nums = nums[:i+1]
	return len(nums)
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
/*
指引：
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
func Merge(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return nil
	}
	// 先对区间数组按照区间的起始位置进行排序（从小到大）,sort.Slice是Go标准库的接受一个比较函数作为参数
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 直接将第一个区间加入结果集
	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		curr := intervals[i]

		//重叠条件：当前起点 <= 上一个的终点，需合并
		if curr[0] <= last[1] {
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
		} else {
			// 不重叠直接加入
			result = append(result, curr)
		}
	}
	return result
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
