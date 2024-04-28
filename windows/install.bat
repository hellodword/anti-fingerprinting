@echo off
cd /d %~dp0

powershell.exe -ExecutionPolicy Bypass -File common.ps1
powershell.exe -ExecutionPolicy Bypass -File run-custom.ps1
