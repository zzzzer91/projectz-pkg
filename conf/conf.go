package conf

type (
	// GinConf 定义了 gin 的配置
	GinConf struct {
		HttpAddr        string `yaml:"httpAddr"`
		ShutdownTimeout int64  `yaml:"shutdownTimeout"` // 单位秒
		SecretKey       string `yaml:"secretKey"`
		ResourcePath    string `yaml:"resourcePath"`
		CorsAllowOrigin string `yaml:"corsAllowOrigin"`
	}

	KitexConf struct {
		Addr string `yaml:"addr"`
	}

	EtcdConf struct {
		Addresses []string `yaml:"addresses"`
	}

	// OtelConf 是 opentelemetry collector 配置
	OtelConf struct {
		Endpoint string `yaml:"endpoint"` // OTLP gRPC receiver
	}

	EsConf struct {
		Addresses              []string `yaml:"addresses"`
		Username               string   `yaml:"username"`
		Password               string   `yaml:"password"`
		CertificateFingerprint string   `yaml:"certificateFingerprint"`
	}

	RedisConf struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	}

	NacosConf struct {
		Host      string
		Port      uint64
		Namespace string
	}

	RpcClientConf struct {
		ConnPoolSize     uint32 `yaml:"connPoolSize"`
		ConnectTimeout   int64  `yaml:"connectTimeout"`
		RpcTimeout       int64  `yaml:"rpcTimeout"`
		KeepAlive        bool   `yaml:"keepAlive"`
		KeepAliveTime    int64  `yaml:"keepAliveTime"`    // KeepAlive 是 false 时，会被忽略
		KeepAliveTimeout int64  `yaml:"keepAliveTimeout"` // KeepAlive 是 false 时，会被忽略
	}
)
