package slice

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-07 11:18

// Map[Src any, Dst any]
//
//	@Description: 辅助方法 用于将两个结构体的类型转换(函数式编程+泛型)
//	@param src
//	@param m
//	@return []Dst
func Map[Src any, Dst any](src []Src, m func(idx int, src Src) Dst) []Dst {
	dst := make([]Dst, len(src))
	for i, s := range src {
		dst[i] = m(i, s)
	}
	return dst
}
