package utils

import (
	"encoding/json"
	"errors"
	"runtime"

	ai "gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsight"
)

func LogErr(appinsight ai.Appinsight, param map[string]any, err error) error {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	if err != nil {
		errMsg := map[string]any{
			"File":  file,
			"Line":  line,
			"Param": param,
			"Error": err.Error(),
		}

		msg, _ := json.Marshal(errMsg)

		go appinsight.Error(errors.New(string(msg)).Error())
		return err
	}

	return nil
}
