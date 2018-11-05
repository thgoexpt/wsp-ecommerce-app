#!/bin/bash

if [[ "$OSTYPE" == "msys" ]] || [[ "$OSTYPE" == "win32" ]]; then
        cmd //c "robot --variable browser:chrome %GOPATH%\src\github.com\guitarpawat\wsp-ecommerce\robot\tests2.robot"
else
        if [[ "$GOPATH" == "" ]]; then
                export GOPATH=~/go
        fi
        robot --variable browser:chrome $GOPATH/src/github.com/guitarpawat/wsp-ecommerce/robot/tests2.robot
fi
