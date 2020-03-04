package bootstrap

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/andodevel/clock_server/constants"
)

// TODO: Consider to use github.com/spf13/viper and github.com/fsnotify/fsnotify instead of godotenv
// Export all configs in App struct

var isInDevMode = true
var isDebugEnabled = true
var loaded bool
var loadErr error

// developement profile as default
var profile = constants.ProfileDevelopement

// Init ...
func Init() {
	loadEnvs()
}

// TODO: OS ENV MUST HAVE HIGHEST PRIORITY!
// LoadEnvs ...
func loadEnvs() {
	var avaiEnvs = make([]string, 0, 10)
	if loaded {
		log.Println("Env has been already loaded!")
		return
	}

	envProfile := os.Getenv("PROFILE")
	if "" != envProfile {
		profile = envProfile
	}

	// Note: Order is import
	avaiEnvs = append(avaiEnvs, ".env")
	avaiEnvs = append(avaiEnvs, ".env."+profile)
	if "test" != profile {
		avaiEnvs = append(avaiEnvs, ".env.local")
	}
	avaiEnvs = append(avaiEnvs, ".env."+profile+".local")

	loadErr = godotenv.Overload(avaiEnvs...)

	loaded = true
}

// Prop ...Get property from enviroment settings
func Prop(prop string) string {
	return os.Getenv(prop)
}

// GetProfile ...
func GetProfile() string {
	return profile
}

// IsInDevMode ...
func IsInDevMode() bool {
	return constants.ProfileDevelopement == GetProfile()
}

// IsDebugEnabled ...
func IsDebugEnabled() bool {
	return isDebugEnabled
}

// EnableDebug ...
func EnableDebug() {
	isDebugEnabled = true
}

// DisableDebug ...
func DisableDebug() {
	isDebugEnabled = false
}
