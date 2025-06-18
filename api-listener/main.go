package main

import (
	"go-stater-listener/pkg/health"
	"go-stater-listener/pkg/utils"
	"go-stater-listener/servicebus"

	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsight"

	bpLogCenter "gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/log_center/logx"
)

func main() {
	run()
}

func run() {
	appinsight := appinsight.NewAppinsights()
	logT := utils.TerminalLogger()
	logM := bpLogCenter.NewLogCenter(appinsight, logT)

	s := servicebus.NewServiceBus(logM)

	go health.Health(logM)
	s.SubscriptionSuccess()
}
