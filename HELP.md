# VaultRDB

## Using the MAKE Commandline for operator-sdk

```bash
operator-sdk init --domain jnnkrdb.de --repo github.com/jnnkrdb/vaultrdb
```

```bash
operator-sdk create api --version v1 --kind VaultRequest --resource --controller
operator-sdk create api --group="core" --version v1 --kind Namespace --resource=false --controller=true
```

## Building the Container

```bash
docker build docker.io/jnnkrdb/vaultrdb:v0.0.1
docker push docker.io/jnnkrdb/vaultrdb:v0.0.1
```

