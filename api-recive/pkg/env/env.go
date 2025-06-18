package env

import (
	"fmt"
	"os"
	"reflect"
	"sync"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// This constant list should be in sync with the .env file
type env struct {
	PRODUCT_CODE string
	MODULE_NAME  string
	BUILD        string
	RELEASE      string
	PORT         string

	DB_HOST     string
	DB_PORT     string
	DB_DATABASE string
	DB_FEDAUTH  string

	APP_INSIGHTS_KEY               string
	APP_INSIGHTS_ROLE              string
	APP_INSIGHTS_SEND_LIVE_METRICS string

	RedisHost string
	RedisPort string
	RedisKey  string

	AAD_SP_AUTHORITY     string
	AAD_SP_TENANT_ID     string
	AAD_SP_CLIENT_ID     string
	AAD_SP_CLIENT_SECRET string
	AAD_SP_CLIENT_SCOPE  string

	AZURE_ACCOUNT_NAME        string
	AZURE_ACCOUNT_KEY         string
	AZURE_BLOB_CONTAINER      string
	AZURE_MAIL_BLOB_CONTAINER string

	USER_PROFILE_BASE_URL string

	PATH_PREFIX string

	SB_HOST                    string
	SB_TOPIC                   string
	SB_QUEUE                   string
	SB_SUBSCRIPTION            string
	SB_RECEIVER_TIMEOUT_IN_SEC string
	SB_RECEIVER_BATCH_SIZE     string
}

// The single instance of env struct
var instance *env

// To make sure one goroutine access this at a time
var lock = &sync.Mutex{}

// Returns the same instance of env struct
func Env() *env {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			log.Debug().Msg("No 'env' instance, create one and initialize its values")
			newInst := load()
			instance = &newInst
		}
	}
	return instance
}

// Loads and initialize values from .env file
func load() env {
	// Read .env
	if _, err := os.Stat("/configs/.env"); err == nil {
		log.Debug().Msg("Loading config from /configs/.env")
		err := godotenv.Load("/configs/.env")
		if err != nil {
			log.Error().Err(err).Msg("")
			log.Fatal().Msg("Error loading .env file from /configs/.env\n")
		}
	} else {
		log.Debug().Msg("Loading config from default location .")
		fmt.Println("Hi")
		err := godotenv.Load()
		if err != nil {
			log.Debug().Msgf("Error loading .env from default location: %+v \n", err)
		}
	}
	newInstance := env{}
	fields := reflect.VisibleFields(reflect.TypeOf(newInstance))
	ps := reflect.ValueOf(&newInstance)
	for i := 0; i < len(fields); i++ {
		fieldname := fields[i].Name
		value := os.Getenv(fieldname)
		ps.Elem().FieldByName(fieldname).SetString(value)
	}
	return newInstance
}
