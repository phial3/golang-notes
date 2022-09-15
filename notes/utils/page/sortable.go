package result

import (
	"container/list"
	"sort"
	"strings"
)

type SortableSlice[T any] struct {
	Slice       []T
	CompareFunc func(T, T) bool
}

func (s SortableSlice[T]) Len() int { return len(s.Slice) }

func (s SortableSlice[T]) Less(i, j int) bool { return s.CompareFunc(s.Slice[i], s.Slice[j]) }

func (s SortableSlice[T]) Swap(i, j int) { s.Slice[i], s.Slice[j] = s.Slice[j], s.Slice[i] }

func PageWrapper[T any](dataSlice []T, cmpFunc func(a, b T) bool, pageNo, pageSize int32,
	orderASC bool, orderField ...string) PageResult {

	if dataSlice == nil {
		dataSlice = make([]T, 0)
	}

	// 排序
	sortSlice := SortableSlice[T]{
		Slice:       dataSlice,
		CompareFunc: cmpFunc,
	}

	var order = ASC
	if orderASC {
		sort.Sort(sortSlice)
	} else {
		order = DESC
		sort.Sort(sort.Reverse(sortSlice))
	}

	// 分页
	total, start, end, pageCount := Page(int32(len(dataSlice)), pageNo, pageSize)

	return PageResult{
		PageNo:     pageNo,
		PageSize:   pageSize,
		PageCount:  pageCount,
		Total:      total,
		Order:      order,
		OrderField: strings.Join(orderField, ","),
		List:       dataSlice[start:end],
	}
}

// Slice2List convert slice to list
func Slice2List[T any](slc []T) list.List {
	var retList list.List
	for _, v := range slc {
		retList.PushBack(v)
	}
	return retList
}

// List2Slice convert list to slice
func List2Slice[T any](lst list.List) []T {
	retSlice := make([]T, 0, lst.Len())
	for i := lst.Front(); i != nil; i = i.Next() {
		retSlice = append(retSlice, i.Value.(T))
	}
	return retSlice
}

func ReverseString(str string) string {
	r := []rune(str)
	j := len(r) - 1
	for i := 0; i < len(r)/2; i++ {
		r[i], r[j] = r[j], r[i]
		j--
	}
	return string(r)
}

func ReverseSlice[T any](slice []T) []T {
	if len(slice) > 0 {
		for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	return slice
}

func ReserveMap(moMap map[string]string) map[string][]string {
	// 建立一个 resMap 与 moMap 容量相同
	// 由于对调可能存在多个值对应一个Key
	// string 需转为 切片[]string
	resMap := make(map[string][]string, len(moMap))

	// 通过for range 遍历 moMap
	// k 即为 Key v 即为 Value
	for k, v := range moMap {
		// 由于现在对应为 切片[]string
		// 使用 append 达到添加多个的效果
		resMap[v] = append(resMap[v], k)
	}

	return resMap
}

// NumberFormat 格式化数值 1,234,567,898.55
func NumberFormat(str string) string {
	length := len(str)
	if length < 4 {
		return str
	}
	arr := strings.Split(str, ".") //用小数点符号分割字符串,为数组接收
	length1 := len(arr[0])
	if length1 < 4 {
		return str
	}
	count := (length1 - 1) / 3
	for i := 0; i < count; i++ {
		arr[0] = arr[0][:length1-(i+1)*3] + "," + arr[0][length1-(i+1)*3:]
	}
	return strings.Join(arr, ".") //将一系列字符串连接为一个字符串，之间用sep来分隔。
}
