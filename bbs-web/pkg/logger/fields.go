package logger

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-01 17:48

func String(key, val string) Field {
	return Field{
		Key:   key,
		Value: val,
	}
}

func Int32(key string, val int32) Field {
	return Field{
		Key:   key,
		Value: val,
	}
}

func Int64(key string, val int64) Field {
	return Field{
		Key:   key,
		Value: val,
	}
}

func Error(err error) Field {
	return Field{
		Key:   "error",
		Value: err,
	}
}
