#!/usr/bin/env bash
echo=echo
for cmd in echo /bin/echo; do
    $cmd >/dev/null 2>&1 || continue

    if ! $cmd -e "" | grep -qE '^-e'; then
        echo=$cmd
        break
    fi
done

CSI=$($echo -e "\033[")
CEND="${CSI}0m"
CDGREEN="${CSI}32m"
CRED="${CSI}1;31m"
CGREEN="${CSI}1;32m"
CYELLOW="${CSI}1;33m"
CBLUE="${CSI}1;34m"
CMAGENTA="${CSI}1;35m"
CCYAN="${CSI}1;36m"

OUT_ALERT() {
    echo -e "${CYELLOW}$1${CEND}"
}

OUT_ERROR() {
    echo -e "${CRED}$1${CEND}"

    exit $?
}

OUT_INFO() {
    echo -e "${CCYAN}$1${CEND}"
}

if [[ -f /etc/redhat-release ]]; then
    release="centos"
elif cat /etc/issue | grep -q -E -i "debian"; then
    release="debian"
elif cat /etc/issue | grep -q -E -i "ubuntu"; then
    release="ubuntu"
elif cat /etc/issue | grep -q -E -i "centos|red hat|redhat"; then
    release="centos"
elif cat /proc/version | grep -q -E -i "raspbian|debian"; then
    release="debian"
elif cat /proc/version | grep -q -E -i "ubuntu"; then
    release="ubuntu"
elif cat /proc/version | grep -q -E -i "centos|red hat|redhat"; then
    release="centos"
else
    OUT_ERROR "[错误] 不支持的操作系统！"
fi

if [[ x"${release}" == x"centos" ]]; then
    yum install epel-release -y
    yum install wget curl unzip zip -y
else
    apt update -y
    apt install wget curl unzip zip -y
fi

cd ~

OUT_ALERT "[提示] 部署 SmartDNS 中"
(curl -fsSL https://raw.githubusercontent.com/aiocloud/stream/master/scripts/smartdns.sh | bash) || OUT_ERROR "部署 SmartDNS 失败！"

OUT_ALERT "[提示] 生成密钥中"
SECRET=$(openssl rand -hex 12)

OUT_ALERT "[提示] 获取地址中"
CURRENT=$(curl -fsSL -4 https://www.cloudflare.com/cdn-cgi/trace | grep ip | tr -d 'ip=')
if [[ "$CURRENT" == "" ]]; then
    OUT_ERROR "[错误] 获取地址失败！"
fi

OUT_ALERT "[信息] 下载程序中"
rm -fr release
wget -O release.zip https://github.com/aiocloud/stream/releases/latest/download/release.zip || OUT_ERROR "下载失败！"

OUT_ALERT "[信息] 解压程序中"
unzip release.zip && rm -f release.zip && cd release

OUT_ALERT "[信息] 设置权限中"
chmod +x stream

OUT_ALERT "[提示] 复制配置中"
cp -f example.json /etc/stream.json
sed -i "s/ccd6c0fe-c4f0-4d36-8dbc-73cd1674dab7/$SECRET/" /etc/stream.json

OUT_ALERT "[提示] 复制程序中"
cp -f stream /usr/bin

OUT_ALERT "[提示] 配置服务中"
cat >/etc/systemd/system/stream.service <<EOF
[Unit]
Description=Stream Unlock Service
After=network.target

[Service]
Type=simple
ExecStart=/usr/bin/stream -c /etc/stream.json
Restart=always
RestartSec=4

[Install]
WantedBy=multi-user.target
EOF

OUT_ALERT "[提示] 重载服务中"
systemctl daemon-reload

OUT_ALERT "[提示] 启动服务中"
systemctl enable --now stream
systemctl enable --now smartdns

OUT_INFO  "[信息] 部署完毕！"
OUT_ALERT "[提示] 您的 DNS 地址 $CURRENT:53"
OUT_ALERT "[提示] 您的 API 密钥 $SECRET"
OUT_ALERT "[提示] 您的 API 地址 http://$CURRENT:8888/aio?secret=$SECRET"
cd ~ && rm -fr release
exit 0
