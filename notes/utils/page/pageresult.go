package result

import "math"

type PageResult struct {
	PageNo     int32       `json:"pageNo"`
	PageSize   int32       `json:"pageSize"`
	PageCount  int32       `json:"pageCount"`
	Total      int32       `json:"total"`
	Order      Order       `json:"order"`
	OrderField string      `json:"orderField"`
	List       interface{} `json:"list"`
	Extend     interface{} `json:"extend"`
}

func Page(dataCount, pageNo, pageSize int32) (total, start, end, pageCount int32) {
	// 页总数
	total = dataCount
	// 计算
	start = (pageNo - 1) * pageSize
	end = start + pageSize
	if end > total {
		end = dataCount
	}

	// 计算页总数,向上取整
	pageCount = int32(math.Ceil(float64(total) / float64(pageSize)))

	return total, start, end, pageCount
}
