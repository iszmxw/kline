#!/bin/bash
pwd &
ls -la &
./iqoption-go-binance --APP_PORT=01 &
./iqoption-go-hsb --APP_PORT=02 &
./iqoption-go-huobi --APP_PORT=03 &
./iqoption-go-fxcm --APP_PORT=06 &
./iqoption-go-iqoption --APP_PORT=07 &
./iqoption-go-monitoring --APP_PORT=08 &
wait
