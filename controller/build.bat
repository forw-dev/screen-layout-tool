@echo off

cd .\cmd
::rsrc -manifest app.manifest -o rsrc.syso
::go build -o ../bin/controller.exe
go build -ldflags="-s -w" -o ../bin/controller.exe
cd ..
if errorlevel 1 goto :pauseOnError :: 有错误就停下来

::copy ..\layouts\Default.json .\bin\
::if errorlevel 1 goto :pauseOnError

goto :eof :: 返回父脚本调用点，不退出命令行

:pauseOnError
pause
