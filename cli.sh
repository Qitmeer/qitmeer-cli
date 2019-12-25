#!/bin/bash

dir=$(pwd)

${dir}/qitmeer-cli \
--rpcserver="192.168.3.10:1236" \
--rpcuser="admin" \
--rpcpassword="123" \
--rpccert="" \
--notls=true \
--tlsskipverify=true \
--proxy="" \
--proxyuser="" \
--proxypass="" \
--debug=false \
--jq \
--timeout="30s" \
$*