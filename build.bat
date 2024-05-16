rem Build dockerfiles

cd base

docker build . --tag vaultrdbbase:latest

cd ../vrdb 

docker build . -f Operator.Dockerfile --tag operator:latest
docker build . -f StorageAPI.Dockerfile --tag storageapi:latest
docker build . -f UiServer.Dockerfile --tag uiserver:latest