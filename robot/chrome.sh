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

$PABOT_CMD --processes 4 --variable browser:$BROWSER $GOPATH/src/github.com/guitarpawat/wsp-ecommerce/robot/*.robot

CODE=$?
if [[ "$CODE" == 0 ]]; then
        echo -e "\e[32mTest Passed\e[0m"
        exit 0
fi

for i in 2 3 4 5; do
        echo ""
        echo -e "\e[33mTest Failed, Attempt $i of 5\e[0m"
        $PABOT_CMD --processes 4 --rerunfailed $GOPATH/src/github.com/guitarpawat/wsp-ecommerce/robot/output.xml --variable browser:$BROWSER $GOPATH/src/github.com/guitarpawat/wsp-ecommerce/robot/*.robot
        
        CODE=$?
        if [[ "$CODE" == 0 ]] || [[ "$CODE" == 252 ]]; then
                echo -e "\e[32mTest Passed\e[0m"
                exit 0
        fi
done

echo -e "\e[31mTest Failed\e[0m"
echo ""
echo -e "\e[101mSummary Logs\e[0m"
for f in pabot_results/*/output.xml; do 
        echo -e "\e[96m$f\e[0m"
        grep -hE --color=auto "FAIL" "$f"
        echo ""
done
exit 1