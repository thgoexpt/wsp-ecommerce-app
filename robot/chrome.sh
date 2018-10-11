#!/bin/bash

if [[ "$OSTYPE" == "msys" ]] || [[ "$OSTYPE" == "win32" ]]; then
        cmd //c "robot --variable browser:chrome %GOPATH%\src\github.com\guitarpawat\wsp-ecommerce\robot\tests.robot"
else
        robot --variable browser:chrome $GOPATH/src/github.com/guitarpawat/wsp-ecommerce/robot/tests.robot
fi