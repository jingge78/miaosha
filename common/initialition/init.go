package initialition

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"github.com/hashicorp/consul/api"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"miaosha-jjl/common/global"
	"os"
	"time"
)

func init() {
	InitZap()
	InitConfig()
	InitNaCos()
	InitDB()
	InitRdb()
	InitEs()
	InitConsul()
}

func InitZap() {
	os.MkdirAll("../log", 0777)
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{
		"../log/dev.log",
	}
	build, _ := config.Build()
	zap.ReplaceGlobals(build)
}
func InitConfig() {
	viper.SetConfigFile("../common/config/dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&global.ConfigData)
	if err != nil {
		panic(err)
	}
}
func InitDB() {
	var err error
	Conf := global.NaCos.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", Conf.User, Conf.Pass, Conf.Host, Conf.Port, Conf.Database)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	sqlDB, err := global.DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic(err)
	}
	err = global.DB.AutoMigrate()
	if err != nil {
		panic(err)
	}

}
func InitRdb() {
	config := global.NaCos.Redis
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Pass, // no password set
		DB:       config.Db,   // use default DB
	})

	pong, err := global.Rdb.Ping(global.CTX).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong, err)
}

func InitEs() {
	var err error
	cfg := elasticsearch.Config{
		Addresses: []string{
			global.NaCos.Elasticsearch.Addr,
		},
		// ...
	}
	global.Es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
}

func InitNaCos() {
	Conf := global.ConfigData.NaCos
	clientConfig := constant.ClientConfig{
		NamespaceId:         Conf.NameSpace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		LogLevel:            "debug",
		Username:            Conf.UserName,
		Password:            Conf.Password,
	}
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      Conf.Host,
			ContextPath: "/nacos",
			Port:        Conf.Port,
			Scheme:      "http",
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
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: Conf.DataId,
		Group:  Conf.Group})
	//fmt.Println(content)
	err = json.Unmarshal([]byte(content), &global.NaCos)
	if err != nil {
		log.Println("配置文件解析失败")
	}
}

func InitConsul() {
	Conf := global.NaCos.Grpc
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", Conf.Host, Conf.Port)
	client, err := api.NewClient(config)
	if err != nil {
		zap.S().Error(err)
		return
	}
	global.ConsulClient = client
	//实例化consul对象
	// 注册之前 要先判断该服务是否已经被注册了  如果说已经被注册了 那就不注册了
	// 用哪些去判断    consul的服务过滤
	// 注册的 user服务

	// 因为客户端要在注册中心中与服务端建立连接 那么客户端和服务端必须都存在注册中心中
	// 如果服务端在注册中心中未进行注册  那么客户端肯定连接不到
	// 所以要先判断服务端是否已经注册到注册中心
	// 如果服务端已经存在注册中心，那么客户端连接服务端的时候，可以直接从注册中心获取服务端的连接信息
	// 如果服务端未注册到注册中心，返回提示语，服务端未进行注册

	//客户端注册注册中心，如果发现服务端未进行注册，则可以在客户端将服务端注册到注册中心  那么服务端就不用再次注册

	//consul := consul.NewConsul(Conf.Host, Conf.Port)
	////服务顾虑
	//filterConsul, err := consul.FilterConsul("user-service")
	//if err != nil {
	//	zap.S().Error("用户服务服务过滤失败")
	//}
	////如果未找到则注册
	//if len(filterConsul) == 0 {
	//	zap.S().Info("service not found consul register service")
	//	err := consul.RegisterConsul("user-service", Conf.Host, Conf.Port, []string{"user-service"})
	//	if err != nil {
	//		zap.S().Errorf("consul注册失败")
	//	} else {
	//		zap.S().Info("consul注册成功")
	//	}
	//} else {
	//	zap.S().Info("服务已经注册到consul")
	//}
}
