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
    echo -e "$CYELLOW$1$CEND"
}

OUT_ERROR() {
    echo -e "$CRED$1$CEND"

    exit 1
}

OUT_INFO() {
    echo -e "$CCYAN$1$CEND"
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

OUT_ALERT "[提示] 安装软件中"
if [[ "$release" == "centos" ]]; then
    yum install openssl-devel pkgconfig make gcc git -y || exit $?
else
    apt update || exit $?
    apt install build-essential pkg-config libssl-dev make git -y || exit $?
fi

OUT_ALERT "[提示] 复制仓库中"
cd /opt && rm -fr smartdns
git clone https://github.com/pymumu/smartdns || exit $?

OUT_ALERT "[提示] 编译程序中"
cd smartdns
make -j$(nproc) || exit $?

OUT_ALERT "[提示] 复制程序中"
cd src && cp -f smartdns /usr/bin/smartdns

OUT_ALERT "[提示] 清理配置中"
rm -fr /etc/smartdns && mkdir /etc/smartdns

OUT_ALERT "[提示] 获取地址中"
CURRENT=$(curl -fsSL -4 https://www.cloudflare.com/cdn-cgi/trace | grep ip | tr -d 'ip=')
if [[ "$CURRENT" == "" ]]; then
    OUT_ERROR "[错误] 获取地址失败！"
fi

OUT_ALERT "[提示] 写入配置中"
wget -O /etc/smartdns/smartdns.conf          https://cdn.jsdelivr.net/gh/aiocloud/stream/smartdns/smartdns.conf    || exit $?
wget -O /etc/smartdns/stream.conf            https://cdn.jsdelivr.net/gh/aiocloud/stream/smartdns/stream.conf      || exit $?
wget -O /etc/systemd/system/smartdns.service https://cdn.jsdelivr.net/gh/aiocloud/stream/smartdns/smartdns.service || exit $?
sed -i "s/1.1.1.1/$CURRENT/" /etc/smartdns/stream.conf

OUT_ALERT "[提示] 重载服务中"
systemctl daemon-reload

OUT_ALERT "[提示] 清理垃圾中"
cd /opt && rm -fr smartdns

OUT_INFO "[信息] 部署完毕！"
exit 0
