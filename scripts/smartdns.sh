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
    exit 1
fi

CURRENT=$(curl -fsSL -4 https://www.cloudflare.com/cdn-cgi/trace | grep ip | tr -d 'ip=')
if [[ "$CURRENT" == "" ]]; then
    exit 1
fi

if [[ "$release" == "centos" ]]; then
    yum install openssl-devel pkgconfig make gcc git -y || exit $?
else
    apt update || exit $?
    apt install build-essential pkg-config libssl-dev make git -y || exit $?
fi

cd /opt && rm -fr smartdns
git clone https://github.com/pymumu/smartdns || exit $?

cd smartdns && make -j$(nproc) || exit $?
cd src && cp -f smartdns /usr/bin/smartdns

rm -fr /etc/smartdns && mkdir /etc/smartdns
wget -O /etc/smartdns/smartdns.conf          https://raw.githubusercontent.com/aiocloud/stream/master/smartdns/smartdns.conf    || exit $?
wget -O /etc/smartdns/stream.conf            https://raw.githubusercontent.com/aiocloud/stream/master/smartdns/stream.conf      || exit $?
wget -O /etc/systemd/system/smartdns.service https://raw.githubusercontent.com/aiocloud/stream/master/smartdns/smartdns.service || exit $?
sed -i "s/1.1.1.1/$CURRENT/" /etc/smartdns/stream.conf

cd /opt && rm -fr smartdns
exit 0
