curl --include \
  --no-buffer \
  --header "Connection: Upgrade" \
  --header "Upgrade: websocket" \
  --header "Host: trading.raceoption.com" \
  --header "Origin: https://trading.raceoption.com" \
  --header "Sec-WebSocket-Key: NVwjmQUcWCenfWu98asDmg==" \
  --header "Sec-WebSocket-Version: 13" \
  https://trading.raceoption.com/api/ws/price
#
#
#curl --include \
#     --no-buffer \
#     --header "Connection: Upgrade" \
#     --header "Upgrade: websocket" \
#     --header "Host: huobiapi.iqoptionproo.com" \
#     --header "Origin: https://huobiapi.iqoptionproo.com" \
#     --header "Sec-WebSocket-Key: NVwjmQUcWCenfWu98asDmg==" \
#     --header "Sec-WebSocket-Version: 13" \
#     https://huobiapi.iqoptionproo.com/v1/api/kline/ws_index

#curl --include \
#     --no-buffer \
#     --header "Connection: Upgrade" \
#     --header "Upgrade: websocket" \
#     --header "Host: 8.215.38.97:8080" \
#     --header "Origin: http://8.215.38.97:8080" \
#     --header "Sec-WebSocket-Key: NVwjmQUcWCenfWu98asDmg==" \
#     --header "Sec-WebSocket-Version: 13" \
#     http://8.215.38.97:8080/v1/api/kline/ws_index
