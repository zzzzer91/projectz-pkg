package conf

import (
	"flag"
	"os"
	"path"
	"testing"
)

func Init() {
	if v := os.Getenv(envKeyTestMode); v == "true" {
		testing.Init()
		flag.Set("configEnv", "test")
		if testWorkSpaceDir := os.Getenv(envKeyVscodeTestWorkspaceDir); testWorkSpaceDir != "" {
			flag.Set("configLocalPath", path.Join(testWorkSpaceDir, "conf"))
		}
	}

	parseFlag()

	switch *configEnv {
	case EnvTest, EnvDev, EnvPre, EnvPro:
	default:
		panic("unknown configEnv: " + *configEnv)
	}

	if IsOnline() {
		if *configHost == "" {
			panic("config Host can not empty")
		}
		if *configNamespace == "" {
			panic("config Namespace can not empty")
		}
		client := newNacosClient(&NacosConf{
			Host:      *configHost,
			Port:      *configPort,
			Namespace: *configNamespace,
		})
		nacosClient = client
	}
}
