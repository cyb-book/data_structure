package merge

func DeepMerge(v1, v2 map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range v1 {
		result[k] = v
	}
	for k, v := range v2 {
		if v2Value, ok := v.(map[string]interface{}); ok {
			if resultValue, exists := result[k]; exists {
				if resultMap, ok := resultValue.(map[string]interface{}); ok {
					result[k] = DeepMerge(resultMap, v2Value)
					continue
				}
			}
		} else if v2Value, ok := v.([]interface{}); ok {
			if resultValue, exists := result[k]; exists {
				if resultSlice, ok := resultValue.([]interface{}); ok {
					result[k] = MergeSlices(resultSlice, v2Value)
					continue
				}
			}
		}
		result[k] = v
	}
	return result
}

func MergeSlices(s1, s2 []interface{}) []interface{} {
	result := CreateMaxLengthSliceFilledWithFirst(s1, s2)
	for i, v := range s2 {
		if i < len(result) {
			if m1, ok := result[i].(map[string]interface{}); ok {
				if m2, ok := v.(map[string]interface{}); ok {
					result[i] = DeepMerge(m1, m2)
					continue
				}
			}
		}
		if i < len(result) {
			result[i] = v
		} else {
			result = append(result, v)
		}
	}
	return result
}
