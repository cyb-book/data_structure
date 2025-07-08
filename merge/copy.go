package merge

import "reflect"

func DeepCopy(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	// 使用反射获取值的类型和种类
	value := reflect.ValueOf(v)
	kind := value.Kind()

	switch kind {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		// 基本类型可以直接返回，因为它们是值类型
		return v

	case reflect.Slice:
		newSlice := reflect.MakeSlice(value.Type(), value.Len(), value.Cap())
		for i := 0; i < value.Len(); i++ {
			newSlice.Index(i).Set(reflect.ValueOf(DeepCopy(value.Index(i).Interface())))
		}
		return newSlice.Interface()

	case reflect.Map:
		newMap := reflect.MakeMap(value.Type())
		for _, key := range value.MapKeys() {
			newValue := DeepCopy(value.MapIndex(key).Interface())
			newMap.SetMapIndex(key, reflect.ValueOf(newValue))
		}
		return newMap.Interface()

	case reflect.Ptr:
		if value.IsNil() {
			return nil
		}
		newPtr := reflect.New(value.Elem().Type())
		newPtr.Elem().Set(reflect.ValueOf(DeepCopy(value.Elem().Interface())))
		return newPtr.Interface()

	case reflect.Struct:
		newStruct := reflect.New(value.Type()).Elem()
		for i := 0; i < value.NumField(); i++ {
			newField := DeepCopy(value.Field(i).Interface())
			newStruct.Field(i).Set(reflect.ValueOf(newField))
		}
		return newStruct.Interface()

	default:
		// 对于其他类型，我们简单地返回原值
		// 在实际应用中，你可能需要为特定类型添加更多的处理逻辑
		return v
	}
}
