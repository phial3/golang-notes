package utils

// 并集：以属于A或属于B的元素为元素的集合成为A与B的并（集）
// 交集： 以属于A且属于B的元素为元素的集合成为A与B的交（集）
// 差集：以属于A而不属于B的元素为元素的集合成为A与B的差（集）
// 子集：以属于A且被包含于B的元素为元素的集合成为A与B的子（集）

// 求并集
func Union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// 求交集
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// 是否有交集
func HasIntersect(first, second []string) bool {
	m := make(map[string]int)
	for _, v := range first {
		m[v]++
	}

	for _, v := range second {
		times, _ := m[v]
		if times == 1 {
			return true
		}
	}
	return false
}

// ContainsValue slice是否包含target
func ContainsValue(slice []string, target string) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}

// int slice 是否包含target
func ContainsIntValue(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// ForEqual 判断相等，需要明确知道切片的类型
func ForEqual(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	if (x == nil) != (y == nil) {
		return false
	}

	for i, v := range x {
		if v != y[i] {
			return false
		}
	}

	return true
}

// 求差集
func Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

func DifferenceV2(a, b []string) []string {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	itr := Intersect(a, b)
	m := map[string]int{}
	for _, v := range itr {
		m[v]++
	}

	var df []string
	for _, v := range a {
		if m[v] == 0 {
			df = append(df, v)
		}
	}

	for _, v := range b {
		if m[v] == 0 {
			df = append(df, v)
		}
	}
	return df
}

func DifferenceV3(slice1, slice2 []string) []string {
	m := make(map[string]string)
	for _, v := range slice1 {
		m[v] = v
	}

	for _, v := range slice2 {
		if m[v] != "" {
			delete(m, v)
		}
	}

	var str []string
	for _, s2 := range m {
		str = append(str, s2)
	}
	return str
}

// subset returns true if the first array is completely
// contained in the second array. There must be at least
// the same number of duplicate values in second as there
// are in first.
// Subset 判断是否子集
func Subset(first, second []string) bool {
	set := make(map[string]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

// 切片去重实现
func RemoveRepeat(arr []string) []string {
	result := make([]string, 0, len(arr))
	temp := map[string]struct{}{}
	for i := 0; i < len(arr); i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}
