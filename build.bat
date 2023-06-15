@echo off
git add .

set /p commitgmsg=CommitMessage:

git commit -m "%commitgmsg%"

git push origin master

git tag

docker build . -t docker.io/jnnkrdb/vaultrdb:latest

docker push docker.io/jnnkrdb/vaultrdb:latest

set /p tag=NewTag:

if "%tag%" == "" goto END

docker build . -t docker.io/jnnkrdb/vaultrdb:%tag%

docker push docker.io/jnnkrdb/vaultrdb:%tag%

:END