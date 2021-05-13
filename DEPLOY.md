# 部署
### 准备工作
- 干净的机器，推荐 Debian 10 系统

### 部署程序
```bash
curl -fsSL https://cdn.jsdelivr.net/gh/aiocloud/smartdns/scripts/kickstart.sh | bash
```

### 修改配置
```bash
nano /etc/stream.json
```

### 开启服务
```bash
systemctl enable --now stream
systemctl enable --now smartdns
```
