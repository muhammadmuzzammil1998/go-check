package check

func Error(e error, f func(v ...interface{})) {
	if e != nil {
		f(e)
	}
}
