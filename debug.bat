rem Start non-kubernetes dockerfiles

rem running ui
docker stop uiserver
docker rm uiserver
docker run --detach --name uiserver -p 8888:80 -d uiserver:latest vaultrdb-ui --swagger --storageapi-address=192.168.178.25:8080

rem running storageapi
docker stop storageapi
docker rm storageapi
docker run --detach --name storageapi -p 8080:80 --mount src="C:\Users\Rodenburger\Documents\GitHub\vaultrdb\_tmp",target="/opt/vaultrdb/data",type=bind -d storageapi:latest 