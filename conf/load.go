package conf

import (
	"fmt"
	"os"

	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/zzzzer91/gopkg/logx"
	"gopkg.in/yaml.v3"
)

func readLocalFile(configFile string) []byte {
	bs, err := os.ReadFile(*configLocalPath + "/" + configFile)
	if err != nil {
		panic(err)
	}
	return bs
}

// LoadFromLocal 从本地文件载入 yaml 配置
func LoadFromLocal(configFile string, out interface{}) {
	bs := readLocalFile(configFile)
	err := yaml.Unmarshal(bs, out)
	if err != nil {
		panic(err)
	}
}

// LoadFromNacos 从 nacos 载入配置
func LoadFromNacos(dataId string, out interface{}) {
	configContent, err := nacosClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  nacosStaticConfigGroupName,
	})
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(configContent), out)
	if err != nil {
		panic(err)
	}
}

// LoadConf 根据环境决定从本地还是 nacos 载入配置
func LoadConf(configName string, out interface{}) {
	if IsOnline() {
		LoadFromNacos(configName, out)
	} else {
		LoadFromLocal(configName, out)
	}
}

func LoadAppConf(out interface{}) {
	LoadConf(fmt.Sprintf("app.%s.yaml", GetEnv()), out)
}

// ListenDynamicConfigChange 动态监听 nacos 配置
func ListenDynamicConfigChange(configName string, callback func(content []byte) error) {
	if IsOnline() {
		// 先初始加载一遍，因为下面的动态加载 onChange 是异步的，
		// 可能会导致配置还没加载完就获取了
		configContent, err := nacosClient.GetConfig(vo.ConfigParam{
			DataId: configName,
			Group:  nacosDynamicConfigGroupName,
		})
		if err != nil {
			panic(err)
		}
		err = callback([]byte(configContent))
		if err != nil {
			panic(err)
		}

		// 动态监听配置更改
		err = nacosClient.ListenConfig(vo.ConfigParam{
			DataId: configName,
			Group:  nacosDynamicConfigGroupName,
			OnChange: func(namespace, group, dataId, data string) {
				err = callback([]byte(data))
				if err != nil {
					logx.Error(err)
				}
			},
		})
		if err != nil {
			panic(err)
		}
	} else {
		bs := readLocalFile(configName)
		err := callback(bs)
		if err != nil {
			panic(err)
		}
	}
}
