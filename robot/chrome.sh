#!/bin/bash

if [[ "$SOLID_ENV" == "CI" ]]; then
        BROWSER=headlesschrome
else
        BROWSER=chrome
        curl -s --ssl --insecure https://localhost:4433/mock/ > /dev/null
fi

if [[ "$OSTYPE" == "msys" ]] || [[ "$OSTYPE" == "win32" ]]; then
        PABOT_CMD=pabot.bat
else
        if [[ "$GOPATH" == "" ]]; then
                export GOPATH=~/go
        fi
        PABOT_CMD=pabot
fi
$PABOT_CMD --processes 8 --variable browser:$BROWSER $GOPATH/src/github.com/guitarpawat/wsp-ecommerce/robot/*.robot

CODE=$?
if [[ "$CODE" == 1 ]]; then
        exit 1
else
        exit 0
fi