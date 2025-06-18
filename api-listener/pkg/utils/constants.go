package utils

var EventTypeRequestStruct = []string{"STARTER_EVENT_REQUEST"}
var EventTypeResponseStruct = []string{"STARTER_EVENT_RESPONSE"}

const (
	PROCESS_HEALTH_CHECK      = "Health check"
	PROCESS_APP_LOG           = "Appinsight and terminal log"
	PROCESS_PANIC_LOG         = "Panic log"
	PROCESS_COMPLETE_ERROR    = "Complete message error"
	PROCESS_START_PROJECT     = "Start go starter listener"
)

var TAGS_SERVICE_NEW_SERVICE = []string{"service", "New service"}
var TAGS_SERVICE_CALLBACK = []string{"service", "callback"}
var TAGS_PANIC = []string{"service", "panic"}
var TAGS_LOG = []string{"service", "log"}
var TAGS_HEALTH_CHECK = []string{"main", "health"}
