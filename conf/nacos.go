package conf

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var (
	nacosClient config_client.IConfigClient
)

func newNacosClient(nacosConf *NacosConf) config_client.IConfigClient {
	clientConfig := constant.NewClientConfig(
		constant.WithNamespaceId(nacosConf.Namespace),
	)
	client, err := clients.NewConfigClient(vo.NacosClientParam{
		ServerConfigs: []constant.ServerConfig{
			{
				IpAddr: nacosConf.Host,
				Port:   nacosConf.Port,
			},
		},
		ClientConfig: clientConfig,
	})
	if err != nil {
		panic(err)
	}
	return client
}
