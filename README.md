# Stream Unlock
[![](https://img.shields.io/badge/Telegram-Group-blue)](https://t.me/aioCloud) [![](https://img.shields.io/badge/Telegram-Channel-green)](https://t.me/aioCloud_channel) 

简单的流媒体解锁后端

## 食用方法
```bash
# 从 GitHub Actions 或者 Releases 中下载 release.zip 并上传至 VPS 上
rm -fr release
mkdir release
cd release
cp ../release.zip .
unzip release.zip
rm -f release.zip
./deploy.sh
```

## 配置文件
```bash
nano /etc/stream.json
```

```jsonc
{
    "api": 8888, // 看下文，好好想想
    "secret": "114514", // 其他机器上 Crontab 下这个 curl -fsSL http://11.4.5.14:8888/aio?secret=114514 就行
    "domains": [ // 解锁域名列表
        "netflix.com"
    ],
    "allowedips": [ // 设定两条自带的允许 IP
        "114.114.114.114",
        "114.114.115.115"
    ],
    "address": "11.4.5.14", // 当前机器的 IP 地址，用于 DNS 回复的
    "upstream": "1.1.1.1:53" // 上游 DNS 地址
}
```

## 控制命令
```bash
systemctl enable --now stream # 开启自启 并 启动服务
systemctl disable --now stream # 关闭自启 兵 停止服务
```
