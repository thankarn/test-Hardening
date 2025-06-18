package health

import (
	"fingw-listener-req/pkg/utils"
	"fmt"
	"os"
	"time"

	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsightsx"
)

func Health(ai appinsightsx.Appinsightsx) {
	f, err := os.Create(fmt.Sprintf("health.txt"))
	if err != nil {
		ai.Error(appinsightsx.LoggerRequest{
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
			ai.Error(appinsightsx.LoggerRequest{
				Tags:    utils.TAGS_HEALTH_CHECK,
				Process: utils.PROCESS_HEALTH_CHECK,
				Error:   fmt.Sprintf("Error Health: %v", err),
			})
		}
		time.Sleep(10 * time.Second)
	}
}
