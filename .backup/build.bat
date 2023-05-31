@ECHO off

REM ### Build Operator ENGINE
ECHO "Build Operator ENGINE"
ECHO.

docker build -t jnnkrdb/vaultrdb:latest .
docker push jnnkrdb/vaultrdb:latest

set /p enginetag_release=Set the ReleaseTag of the Engine-Container:

if "%enginetag_release%" == "" goto END

docker tag jnnkrdb/vaultrdb:latest jnnkrdb/vaultrdb:%enginetag_release%
docker push jnnkrdb/vaultrdb:%enginetag_release%

REM ### Finished
:END
ECHO "FINISHED BUILD"