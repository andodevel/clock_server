package bootstrap

import (
	"image/color"
	"log"
	"os"

	"github.com/joho/godotenv"
	qrcode "github.com/skip2/go-qrcode"

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
	genQR()
}

func genQR() {
	// FIXME: QR should be saved in FS or better in DB.
	q, err := qrcode.New(GetClockURL(), qrcode.High)
	if err != nil {
		panic(err)
	}
	q.DisableBorder = true
	q.ForegroundColor = color.Black
	q.BackgroundColor = color.White
	err = q.WriteFile(512, "app/assets/images/clock_qr.png")
	if err != nil {
		panic(err)
	}
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

// PropDefault ...Get property from enviroment settings, if not exist return defaultValue
func PropDefault(prop string, defaultValue string) string {
	value := os.Getenv(prop)
	if "" != value {
		return value
	}

	return defaultValue
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

func GetHost() string {
	return "localhost"
}

func GetPort() string {
	return PropDefault(constants.EnvPort, "38080")
}

// GetClockURL ...
func GetClockURL() string {
	// FIXME: resolve current host and port
	return "http://" + GetHost() + ":" + GetPort() + "/clock"
}
