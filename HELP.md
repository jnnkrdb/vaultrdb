# VaultRDB

```bash
operator-sdk init --domain jnnkrdb.de --repo github.com/jnnkrdb/vaultrdb
```


```bash
operator-sdk create api --version v1 --kind VaultRequest --resource --controller
```
--group vault 
```bash
docker build docker.io/jnnkrdb/vaultrdb:v0.0.1
docker push docker.io/jnnkrdb/vaultrdb:v0.0.1
```