@ECHO off

SET WORKDIR=%~dp0

SET OPERATOR_ENGINE=%WORKDIR%operator
rem SET OPERATOR_UI=%WORKDIR%ui

ECHO.
ECHO --- WORKDIR: %WORKDIR%
ECHO --- OPERATOR_ENGINE: %OPERATOR_ENGINE%
rem ECHO --- OPERATOR_UI: %OPERATOR_UI%
ECHO.

REM ### Build Operator ENGINE
ECHO "Build Operator ENGINE"
ECHO.

cd %OPERATOR_ENGINE%
docker build -t jnnkrdb/vaultrdb-engine:latest .
docker push jnnkrdb/vaultrdb-engine:latest

set /p enginetag_release=Set the ReleaseTag of the Engine-Container:

if "%enginetag_release%" == "" goto END

docker tag jnnkrdb/vaultrdb-engine:latest jnnkrdb/vaultrdb-engine:%enginetag_release%
docker push jnnkrdb/vaultrdb-engine:%enginetag_release%

goto END

REM ### Build Operator UI
ECHO "Build Operator UI"
ECHO.

cd %OPERATOR_UI%
docker build -t jnnkrdb/vaultrdb-ui:latest .
docker push jnnkrdb/vaultrdb-ui:latest

set /p uitag_release=Set the ReleaseTag of the UI-Container:

docker tag jnnkrdb/vaultrdb-ui:latest jnnkrdb/vaultrdb-ui:%uitag_release%
docker push jnnkrdb/vaultrdb-ui:%uitag_release%

REM ### back to original workdir
:END
cd %WORKDIR%