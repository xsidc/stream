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
