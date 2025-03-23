package initialition

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"miaosha-jjl/common/global"
	"os"
)

func Init() {
	InitZap()
	InitConfig()
	InitDB()
	InitRdb()
	InitEs()
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
	viper.SetConfigFile("common/config/dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.Unmarshal(&global.GlobalConfig)
	fmt.Println("配置参数", global.GlobalConfig)
}
func InitDB() {
	var err error
	config := global.GlobalConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", config.User, config.Password, config.Host, config.Port, config.Database)
	global.DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	log.Println("MySQL connect successful!!")
}
func InitRdb() {
	config := global.GlobalConfig.Redis
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
	})
	err := global.Rdb.Ping().Err()
	if err != nil {
		panic(err)
	}
	log.Println("Redis connect successful!!")
}
func InitEs() {
	var err error
	cfg := elasticsearch.Config{
		Addresses: []string{
			global.GlobalConfig.Es.Address,
		},
		// ...
	}
	global.Es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	log.Println("elasticsearch connect successful!!")
}
