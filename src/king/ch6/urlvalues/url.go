package main

//Values 映射字符串到字符串列表
type Values map[string][]string

//Get返回第一个具有给定Key的值
// 如不存在, 则返回空字符串
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// Add 添加一个键值到对应 key 列表中。
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}
