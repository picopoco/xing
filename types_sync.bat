@echo off

REM *********** 
REM *  32Bit  *
REM ***********
call %GOPATH%\dep\batch_scripts\32.bat

cls
cd %PROJECT_ROOT%\xing

REM copy types_c.h %PROJECT_ROOT%\xing_C32\internal\
copy types_c.orig types_1.go
go tool cgo -godefs types_1.go > types_2.go
sed -e 's/uint8/byte/g' types_2.go > types_3.go
sed -e 's/int8/byte/g' types_3.go > types_c.go
del types_1.go
del types_2.go
del types_3.go