@echo off
setlocal enabledelayedexpansion

for /f "tokens=*" %%a in (aws_acc_list.txt) do (
  echo Checking IP: %%a
  netguru --csp=aws --id=%%a --output=table >> "aws_results_!date:/=-!_!time::=-!.txt"
)

echo Netguru scanning completed for all
endlocal

REM netguru --csp=aws --id=657169126488 --output=table 