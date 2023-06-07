#!/bin/bash
./iqoption-go-binance --APP_PORT=01 >/dev/null 2>error_binance.log &
./iqoption-go-hsb --APP_PORT=02 >/dev/null 2>error_hsb.log &
./iqoption-go-huobi --APP_PORT=03 >/dev/null 2>error_huobi.log &
./iqoption-go-fxcm --APP_PORT=06 >/dev/null 2>error_fxcm.log &
./iqoption-go-iqoption --APP_PORT=07 >/dev/null 2>error_iqoption.log &
./iqoption-go-monitoring --APP_PORT=08 >/dev/null 2>error_monitoring.log &
wait
