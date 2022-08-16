# TimuShow 微服务项目 Srv

本项目使用 NaCos 作为配置中心，Consul作为服务中心，Raids 为手机验证码存储，Mysql8为用户等数据存储，接入aliyun短信服务，实现简易的分布式微服务项目，各子项目配置信息如下:

## NaCos 存储的配置文件信息
user_srv 配置文件
```json
{
    "name": "user-srv",
    "host":"192.168.0.162",
    "tags":["user","srv"],
    "mysql": {
        "host": "127.0.0.1",
        "port": 3306,
        "user": "root",
        "password": "123456",
        "db": "shop_user_srv"
    },
    "consul": {
        "host": "127.0.0.1",
        "port": 8500
    }
}
```

goods_srv 配置文件
```json
{
    "name": "goods-srv",
    "host":"192.168.0.162",
    "tags":["goods","srv"],
    "mysql": {
        "host": "127.0.0.1",
        "port": 3306,
        "user": "root",
        "password": "123456",
        "db": "shop_goods_srv"
    },
    "consul": {
        "host": "127.0.0.1",
        "port": 8500
    }
}
```