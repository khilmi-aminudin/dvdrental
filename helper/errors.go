package helper

import "github.com/sirupsen/logrus"

func LogErrorAndPanic(err error) {
	if err != nil {
		Logger().Panic(err.Error())
	}
}

func LogErrorWithFields(err error, fieldName string, value interface{}) {
	if err != nil {
		Logger().WithFields(logrus.Fields{
			fieldName: value,
		}).Error(err.Error())
	}
}

func LogError(err error) {
	if err != nil {
		Logger().Error(err.Error())
	}
}
