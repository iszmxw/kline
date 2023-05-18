#curl --include \
#  --no-buffer \
#  --header "Connection: Upgrade" \
#  --header "Upgrade: websocket" \
#  --header "Host: stream.binance.com:9443" \
#  --header "Origin: https://stream.binance.com:9443" \
#  --header "Sec-WebSocket-Key: NVwjmQUcWCenfWu98asDmg==" \
#  --header "Sec-WebSocket-Version: 13" \
#  https://stream.binance.com:9443/ws/ethusd@kline_1m
#
#
curl --include \
     --no-buffer \
     --header "Connection: Upgrade" \
     --header "Upgrade: websocket" \
     --header "Sec-WebSocket-Key: NVwjmQUcWCenfWu98asDmg==" \
     --header "Sec-WebSocket-Version: 13" \
     https://goclientio.smosignal.com/app/v3/trade/wsIndex.json?token=5D99F774-54C9-4515-8469-FC22C6A68FE3?isvest=0

#curl --include \
#     --no-buffer \
#     --header "Connection: Upgrade" \
#     --header "Upgrade: websocket" \
#     --header "Host: 8.215.38.97:8080" \
#     --header "Origin: http://8.215.38.97:8080" \
#     --header "Sec-WebSocket-Key: NVwjmQUcWCenfWu98asDmg==" \
#     --header "Sec-WebSocket-Version: 13" \
#     http://8.215.38.97:8080/v1/api/kline/ws_index



#curl --location --request POST 'https://api.adjust.com/public/v1/apps/yxs12pfewq/trackers'\
#--header 'Authorization: Token token={API_TOKEN}' \
#--header 'Content-Type: application/json' \
#--data-raw '{
#  "name": "Adroll"
#}'