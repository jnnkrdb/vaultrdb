@echo off
git add .

set /p commitgmsg=CommitMessage:

git commit -m "%commitgmsg%"

git push origin master

git tag

set /p tag=NewTag:

if "%tag%" == "" goto END

docker build docker.io/jnnkrdb/vaultrdb:%tag%

docker push docker.io/jnnkrdb/vaultrdb:%tag%

:END