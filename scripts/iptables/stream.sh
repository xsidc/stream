#!/usr/bin/env bash
HOST=$(wg | grep endpoint | head -n 1 | sed "s/  endpoint: //" | grep -Eo --color=never '[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+')
if [[ $HOST == "" ]]; then
    echo "WireGuard Not Connected"
    exit 0
fi

LAST=""
if [[ -f /opt/stream/stream.last ]]; then
    LAST=$(cat /opt/stream/stream.last)
fi

if [[ $LAST == $HOST ]]; then
    exit 0
fi
echo $HOST > /opt/stream/stream.last

echo "$HOST"
sed "s/upstream/${HOST}/" /opt/stream/origin.conf > /etc/smartdns/stream.conf

echo "systemctl restart smartdns"
systemctl restart smartdns
exit 0