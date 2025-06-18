package main

import (
	"fingw-listener-req/api"
	"fingw-listener-req/pkg/env"
	"fingw-listener-req/pkg/health"
	"fingw-listener-req/servicebus"
	"fmt"

	"gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/appinsightsx"
	httpClient "gitlab.com/banpugroup/banpucoth/itsddev/library/golang/go-azure-sdk.git/http_client"
)

func main() {
	run()
}

func run() {
	ai := appinsightsx.NewAppinsightsx(appinsightsx.InitProperties{
		InstrumentationKey: env.Env().APP_INSIGHTS_KEY,
		RoleName:           env.Env().APP_INSIGHTS_ROLE,
		EnableZlog:         true,
	})
	ai.Defer()

	go health.Health(ai)

	// http-client
	acquireToken, _ := httpClient.NewBanpuAcquireToken(false)
	http := httpClient.NewHttpClient(acquireToken)

	// http-base
	//httpBase := httpBase.NewHttpClient(acquireToken)

	finGwApi := api.NewFinGwApi(http)
	s := servicebus.NewServiceBus(ai, finGwApi)

	s.SubscriptionSuccess()

	ai.Info(appinsightsx.LoggerRequest{
		Process: fmt.Sprintf("Start Project | Listen Port:%v", env.Env().PORT),
	})

}
