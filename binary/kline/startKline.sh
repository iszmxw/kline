#!/bin/bash
./iqoption-go-binance --APP_PORT=01 >/dev/null 2>error/binance.log &
./iqoption-go-hsb --APP_PORT=02 >/dev/null 2>error/hsb.log &
./iqoption-go-huobi --APP_PORT=03 >/dev/null 2>error/huobi.log &
./iqoption-go-fxcm --APP_PORT=06 >/dev/null 2>error/fxcm.log &
./iqoption-go-iqoption --APP_PORT=07 >/dev/null 2>error/iqoption.log &
./iqoption-go-monitoring --APP_PORT=08 >/dev/null 2>error/monitoring.log &
wait
