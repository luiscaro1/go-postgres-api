package error_handlers

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
