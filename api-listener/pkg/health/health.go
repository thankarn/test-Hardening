package health

import (
	"fmt"
	"go-stater-listener/pkg/env"
	"go-stater-listener/pkg/utils"
	"os"
	"time"

	bpLogCenter "gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/log_center/logx"
	bpLogCenterModel "gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/log_center/model"
)

func Health(logM bpLogCenter.LogCenter) {
	f, err := os.Create(fmt.Sprintf("%s/health.txt", env.Env().TEMP_PATH))
	if err != nil {
		logM.Error(bpLogCenterModel.LoggerRequest{
			Tags:    utils.TAGS_HEALTH_CHECK,
			Process: utils.PROCESS_HEALTH_CHECK,
			Error:   fmt.Sprintf("Error Health: %v", err),
		})
	}
	defer f.Close()

	for {
		fmt.Println("Health: ", time.Now().String())
		_, err = f.WriteAt([]byte(time.Now().String()), 0)
		if err != nil {
			logM.Error(bpLogCenterModel.LoggerRequest{
				Tags:    utils.TAGS_HEALTH_CHECK,
				Process: utils.PROCESS_HEALTH_CHECK,
				Error:   fmt.Sprintf("Error Health: %v", err),
			})
		}
		time.Sleep(10 * time.Second)
	}
}
