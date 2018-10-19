package main

// basename 移除路径部分和. 后缀

func basename(s string) {
	// 将最后一个'/'和之前的部分全都舍弃
	for i := len(s) -1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// 保留最后一个 '.' 之前的所有内容
	for i := len(s) -1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s

}
