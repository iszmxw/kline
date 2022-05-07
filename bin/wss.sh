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



curl 'https://trading.raceoption.com/api/login' \
  -H 'authority: trading.raceoption.com' \
  -H 'accept: */*' \
  -H 'accept-language: zh-CN,zh;q=0.9,en;q=0.8' \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json; charset=utf-8' \
  -H 'cookie: _gcl_au=1.1.1873918188.1650598252; _ga=GA1.2.1639315437.1650598253; __cf_bm=.qUbqMPc1YVdCEg1a0FF5S_nFHlHz4q6zEv.CabA0QI-1651909777-0-Aa22MKMf9rhvvXGIAGWRDTSVFM8jn4rwFiH6Qkrs1x+F1y55n/2f1wTo5VY13K4tyoKD/gAz0M014m1kwS7ZjP20eln7ixlvXd/R3AuIn8Gkcizse0jC1Qk5X2MzlbXzjw==' \
  -H 'origin: https://trading.raceoption.com' \
  -H 'pragma: no-cache' \
  -H 'referer: https://trading.raceoption.com/' \
  -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Windows"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36' \
  --data '{"login":"mail@54zm.com","password":"wap.54zm.com"}' \
  --compressed