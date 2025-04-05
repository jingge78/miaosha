## 秒杀项目结构

## 1. 项目架构

```shell
├── api
│   └── client
│   └── handler
│       └── router
│   └── request
│   └── response
├── common
│   └── config
│   └── glbal
│   └── GRPC
│   └── initialize
│   └── middleware
│   └── model
│   └── pkg
│   └── proto
│       └── order
│       └── product
│       └── register
│       └── store_cart
│       └── user_enter
│       └── user
│   └── utils
├── docs
├── log
└── service
    └── order_service
    └── product_service
    └── register
    └── store_cart_service
    └── user_enter_service
    └── user_service
```

| 文件夹                    | 说明            | 描述                                                |
|------------------------|---------------|---------------------------------------------------|
| `api`                  | api层          | api层                                              |
| `--client`             | gRPC用户客户端     | 创建和使用gRPC用户客户端                                    |
| `--handler`            | 处理请求          | 处理api业务逻辑                                         |
| `----router`           | 路由层           | 路由层                                               |
| `--request`            | 入参结构体          | 接收前端发送到后端的数据                                            |
| `--response`           | 出参结构体         | 返回给前端的数据结构体                                            |
| `common`               | common层       | common层                                           |
| `--config`             | 配置包           | config.yaml对应的配置结构体                               |
| `--global`             | 全局对象          | 全局对象                                            |
| `--GRPC`               | 模型层           | 模型对应数据表                                           |
| `--initialize`         | 初始化           | 负责存放静态文件                                            |
| `--middleware`         | 中间件层          | 中间鉴权                                            |
| `--model`              | 模型层           | 模型对应数据表                                            |
| `--pkg`                | 第三方接口         | 封装第三方接口                                           |
| `--proto`              | proto文件       | 封装proto生成文件                                       |
| `--utils`              | 工具包           | 工具函数封装                                            |
| `docs`                 | swagger文档目录           | swagger文档目录                                            |
| `log`                  | 日志包           | 日志包                                            |
| `service`              | service层      | 存放业务逻辑问题                                          |
| `--order_service`      | 订单业务           | 存放订单业务逻辑问题                                        |
| `--product_service`    | 商品业务           | 存放商品业务逻辑问题                                        |
| `--register`           | 工具函数封装           | 工具函数封装                                            |
| `--store_cart_service` | 购物车业           | 存放购物车业务逻辑问题                                       |
| `--user_enter_service` | 商户业务           | 存放商户业务逻辑问题                                        |
| `--user_service`       | 用户业务           | 存放用户业务逻辑问题                                        |


## 2. 项目地址

github:

https://github.com/jingge78/miaosha/tree/main

1、clone项目源代码

`git clone https://github.com/jingge78/miaosha/tree/main`

2、新建数据库名（如：seckill_project_db） 、导入目录 ./common/config/dev.yaml

根据实际环境修改 ./common/config/dev.yaml

```
    tables:          ""
    tablesEx:        ""
    removePrefix:    "gin"
    descriptionTag:  true
    noModelComment:  true
    path: "./common/config/dev.yaml"
    
    # Mysql 配置示例
    mysql:
    user: ""
    password: ""
    host: ""
    port: ""
    database: "seckill_project_db"
    
    # Redis 配置示例
    redis:
    host: ""
    password: ""
    
    # Elasticsearch 配置示例
    es:
    address: ""
```



## 3. 技术选型

- 前端：用基于 [Vue](https://vuejs.org) 的 [Element](https://github.com/ElemeFE/element) 构建基础页面。
- 后端：用 [Gin](https://gin-gonic.com/) 快速搭建基础restful风格API，[Gin](https://gin-gonic.com/) 是一个go语言编写的Web框架。
- 数据库：采用`MySql`  (8.4.4) 版本 数据库引擎 InnoDB，使用 [gorm](http://gorm.cn) 实现对数据库的基本操作。
- 缓存：采用`Redis`  (7.4.2) 版本实现记录当前活跃用户的`jwt`令牌并实现多点登录限制。
- 搜索：采用`Elasticsearch`  (7.17.6) 版本实现。
- API文档：使用`Swagger`构建自动化文档。
- 配置文件：使用 [fsnotify](https://github.com/fsnotify/fsnotify) 和 [viper](https://github.com/spf13/viper) 实现`yaml`格式的配置文件。
- 日志：使用 [zap](https://github.com/uber-go/zap) 实现日志记录。




