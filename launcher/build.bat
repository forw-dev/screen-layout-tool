@echo off

cd .\src
set exe="D:\Program Files\AutoHotkey\Compiler\Ahk2Exe.exe"
set in=".\launcher.ahk"
set out="..\bin\launcher.exe"
set icon=".\launcher.ico"
%exe% /in %in% /out %out% /icon %icon%
cd ..
if errorlevel 1 goto :pauseOnError

::copy ..\layouts\*.json .\bin\layouts\
::if errorlevel 1 goto :pauseOnError

::copy ..\controller\bin\controller.exe .\bin\
::if errorlevel 1 goto :pauseOnError

goto :eof :: 返回父脚本调用点，不退出命令行。

:pauseOnError
pause
