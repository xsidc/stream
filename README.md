# Stream Unlock
[![](https://img.shields.io/badge/Telegram-Group-blue)](https://t.me/aioCloud)
[![](https://img.shields.io/badge/Telegram-Channel-green)](https://t.me/aioCloud_channel) 

流媒体解锁后端

## 推荐系统
- Debian 10
- Ubuntu 20.04
- CentOS 8 Stream

## 部署命令
```bash
# 部署
curl -fsSL https://git.io/JkMeC | bash

# 升级
curl -fsSL https://git.io/JkMel | bash

# 卸载
curl -fsSL https://git.io/JkMeR | bash
```

## 配置文件
```bash
nano /etc/stream.json
```
```jsonc
/*  不要复制这里的配置，这里的配置仅作解释用
    不要复制这里的配置，这里的配置仅作解释用
    不要复制这里的配置，这里的配置仅作解释用 */
{
    // API
    "api": {
        // 绑定地址（为空则不启用 API 模块）
        "addr": ":8888",

        // 访问密钥
        "secret": "ccd6c0fe-c4f0-4d36-8dbc-73cd1674dab7"

        /*
            DDNS API
            curl -fsSL http://1.1.1.1:8888/aio?secret=ccd6c0fe-c4f0-4d36-8dbc-73cd1674dab7

            注意替换 IP 和端口，写入 crontab 即可
        */
    },

    // DNS
    "dns": {
        // 绑定地址（为空则不启用 DNS 模块）
        "addr": ":53",

        // 当前机器的 IP 地址，回复需要解锁的域名使用
        "myip": "1.1.1.1",

        // 上游 DNS 地址，解析其他域名使用
        "upstream": "1.1.1.1:53"
    },

    // 不懂不要动这里
    "mitm": {
        "http": [ ":80" ],
        "tls": [ ":443" ]
    },

    // 允许的地址（可以把 IP 填写到这里，用 API 就可以不用写这里了）
    "allowed": [
        "1.1.1.0/24",
        "1.0.0.1"
    ],

    // 需要解锁的域名（不懂不要改这里，默认配置就会包含绝大多数域名）
    "domains": []
}
```

## 控制命令
```bash
# 启动服务并开启自启
systemctl enable --now stream

# 停止服务并关闭自启
systemctl disable --now stream

# 查看服务状态
systemctl status stream

# 获取实时日志
journalctl -f -u stream
```
