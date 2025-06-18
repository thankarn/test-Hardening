package utils

import (
	"go-stater-listener/domain/model"
	"runtime"
)

func ErrorData(err error) model.ErrorProps {
	_, fileError, lineError, _ := runtime.Caller(1)
	return model.ErrorProps{
		Error:     err,
		FileError: fileError,
		LineError: lineError,
	}
}
