package merge

func CreateMaxLengthSliceFilledWithFirst(a, b []interface{}) []interface{} {
	// 确定新 slice 的长度
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}

	// 创建新的 slice
	newSlice := make([]interface{}, maxLen)

	// 检查 A 是否为空
	if len(a) == 0 {
		return newSlice // 返回一个全是 nil 的 slice
	}

	// 用 A 的第一个元素填充整个新 slice
	for i := range newSlice {
		newSlice[i] = DeepCopy(a[0])
	}

	return newSlice
}
