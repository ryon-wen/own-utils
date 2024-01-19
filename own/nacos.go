package own

import (
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type Nacos struct {
	IpAddr      string // 127.0.0.1
	Port        uint64 // 8848
	NamespaceId string
	DataId      string
	Group       string
}

func InitNacos(c *Nacos) (string, error) {
	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         c.NamespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: c.IpAddr,
			Port:   c.Port,
		},
	}
	// Another way of create config client for dynamic configuration (recommend)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return "", err
	}
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: c.DataId,
		Group:  c.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	if err != nil {
		return "", err
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: c.DataId,
		Group:  c.Group})
	if err != nil {
		return "", err
	}
	fmt.Println("nacos content:" + content)
	return content, nil
}
