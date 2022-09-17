package helper

import "reflect"

func Contains(list interface{}, elem interface{}) bool {
	listV := reflect.ValueOf(list)

	if listV.Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < listV.Len(); i++ {
		item := listV.Index(i).Interface()
		if elem == nil {
			itemV := reflect.ValueOf(item)
			if itemV.Kind() != reflect.Slice {
				if IsNilOrEmpty(item) {
					return true
				}
			}
			continue
		}

		// 型変換可能か確認する
		if !reflect.TypeOf(elem).ConvertibleTo(reflect.TypeOf(item)) {
			continue
		}

		// 型変換する
		target := reflect.ValueOf(elem).Convert(reflect.TypeOf(item)).Interface()
		// 等価判定をする
		if ok := reflect.DeepEqual(item, target); ok {
			return true
		}
	}
	return false
}
