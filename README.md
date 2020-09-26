# Stream Unlock
[![](https://img.shields.io/badge/Telegram-Group-blue)](https://t.me/aioCloud)
[![](https://img.shields.io/badge/Telegram-Channel-green)](https://t.me/aioCloud_channel) 

流媒体解锁后端

## 推荐系统
- CentOS 8
- Debian 9
- Debian 10
- Ubuntu 18.04
- Ubuntu 20.04

## 部署命令
```bash
# 部署
curl -fsSL https://git.io/JUVOa | bash

# 升级
curl -fsSL https://git.io/JUVOK | bash

# 卸载
curl -fsSL https://git.io/JUVO6 | bash
```

## 配置文件
```bash
nano /etc/stream.json
```

```jsonc
// 这个带注释的配置文件，不要写进去 ... 会报错的
{
    // 设置 API 端口
    "api": 8888,

    // 设置 API 密钥
    "secret": "aioCloud",

    /*
        使用 API 接口

        * * * * * curl -fsSL http://11.4.5.14:8888/aio?secret=aioCloud > /dev/null 2>&1 &

        把上面的命令加入到需要解锁的机器的 Cron 里即可，注意修改 地址、端口 和 密钥
    */

    // 设置 DNS 端口，请在清楚明白的情况下修改此值
    "dnsport": 53,

    // 设置解锁域名
    "domains": [
        "netflix.com"
    ],

    // 设置允许的 IP 地址（如果机器多，并且经常变更 IP 地址，建议使用上面的 API 接口）
    "allowedips": [
        "1.14.5.14",
        "1.145.1.4",
    ],

    // 这里填写解锁机的 IP 地址（用于 DNS 回复）
    "address": "11.4.5.14",

    // 上游 DNS 地址（必须带端口）
    "upstream": "1.1.1.1:53"
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
