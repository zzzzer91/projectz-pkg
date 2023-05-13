package conf

import (
	"flag"
	"os"
	"path"
)

var (
	configEnv       = flag.String("configEnv", "", "application env")
	configLocalPath = flag.String("configLocalPath", "", "local config file path")
	configHost      = flag.String("configHost", "", "the config center host")
	configPort      = flag.Uint64("configPort", 8848, "the config center port")
	configNamespace = flag.String("configNamespace", "", "the config center namespace")
)

func parseFlag() {
	flag.Parse()
	if *configEnv == "" {
		panic("App's env can not be empty")
	}
	if *configLocalPath == "" {
		pwd, _ := os.Getwd()
		*configLocalPath = path.Join(pwd, "conf")
	}
}

func GetEnv() string {
	return *configEnv
}

func IsTestEnv() bool {
	return GetEnv() == EnvTest
}

func IsDevEnv() bool {
	return GetEnv() == EnvDev
}

func IsProEnv() bool {
	return GetEnv() == EnvPro
}

func IsOnline() bool {
	switch GetEnv() {
	case EnvPro:
		return true
	default:
		return false
	}
}
