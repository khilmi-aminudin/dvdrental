package helper

func LogErrorAndPanic(err error) {
	if err != nil {
		Logger().Panic(err.Error())
	}
}
