package incoming_event

import (
	"fingw-listener-req/pkg/env"
	"fmt"
)

type EVENT_TYPE struct {
	CASE_NAME string
	URL       string
}

var CASE_TYPE_REQUEST = []EVENT_TYPE{
	{
		CASE_NAME: "case_one",
		URL:       "/request/validate",
	},
}

// Generate HTTP methods and API endpoints for different event types
type HTTPMethod_TYPE struct {
	GET    string
	POST   string
	PUT    string
	DELETE string
}

var HTTP_METHODS = HTTPMethod_TYPE{
	GET:    "GET",
	POST:   "POST",
	PUT:    "PUT",
	DELETE: "DELETE",
}

// APIConfig represents the configuration for a single API endpoint
type APIConfig struct {
	EventType         string            `json:"eventType"`
	APIEndpoint       string            `json:"apiEndpoint"`
	HTTPMethod        string            `json:"httpMethod"`
	Headers           map[string]string `json:"headers,omitempty"`           // omitempty if headers are optional
	TransformJmesPath string            `json:"transformJmesPath,omitempty"` // JMESPath expression for payload transformation
}

// AppConfig holds the entire application configuration
type AppConfig struct {
	APIMappings []APIConfig `json:"apiMappings"`
}

// Registered APIs with their configurations
var (
	API_OrderCreated = APIConfig{
		EventType:   "OrderCreated",
		APIEndpoint: fmt.Sprintf("%s/%s", env.Env().FINGW_URL, "/orders/create"),
		HTTPMethod:  "POST",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-API-Key":    "your-order-api-key",
		},
	}
	API_CREATE = APIConfig{
		EventType:   "OrderCreated",
		APIEndpoint: fmt.Sprintf("%s/%s", env.Env().FINGW_URL, "/orders/create"),
		HTTPMethod:  "POST",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-API-Key":    "your-order-api-key",
		},
	}
	API_UPDATE   = APIConfig{}
	API_RESPONSE = APIConfig{}
)

// Registered_APIs is a slice of APIConfig that holds all registered APIs
var Registered_APIs = []APIConfig{
	API_OrderCreated,
}
