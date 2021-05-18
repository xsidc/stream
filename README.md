# Stream Unlock
[![](https://img.shields.io/badge/Telegram-Group-blue)](https://t.me/aioCloud)
[![](https://img.shields.io/badge/Telegram-Channel-green)](https://t.me/aioCloud_channel) 

流媒体解锁后端

## 推荐系统
- Debian 10
- Ubuntu 20.04
- CentOS 8 Stream

## 安装Unzip解压程序
```bash
yum install -y unzip zip
```

## 安装Yum安装
```bash
yum -y install wget
```

## 一键部署
```bash
部署
curl -fsSL https://raw.githubusercontent.com/aiocloud/stream/master/scripts/kickstart.sh | bash

升级
curl -fsSL https://git.io/JkMel | bash

卸载
curl -fsSL https://git.io/JkMeR | bash
```

## 配置文件
Stream
```
/etc/stream.json
# 访问端口 "addr": ":8888",
# 访问秘钥 "secret": "weiguanyun"
```
SmartDNS
```
/etc/smartdns/smartdns.conf

请自行增加解锁内容
twitter.com 全国
api.twitter.com 全国
syosetu.com 日本
rakuten.co.jp 日本

设置好记得重启主机reboot
```

## DDNSAPI
```
curl -fsSL http://DNSIP:8888/aio?secret=weiguanyun
注意替换 IP 和端口，写入 crontab 即可
```

## 控制命令
```
# 启动服务并开启自启
systemctl enable --now stream

# 停止服务并关闭自启
systemctl disable --now stream

# 查看启动服务状态
systemctl status stream

# 查看DNS服务状态
systemctl status smartdns

# 获取实时日志
journalctl -f -u stream
```
