# Packaging

## DockerHub
### Build
```bash
docker build -f .\build\package\Dockerfile -t goldenhand/immortals .
```
### Push
```bash
docker push goldenhand/immortals:0.0.1
```

## Github Package Registry
### Login 
```bash
docker login ghcr.io -u goldenhand76 --password <YOUR_TOKEN>
```
### Build
```bash
docker build -f .\build\package\Dockerfile -t ghcr.io/goldenhand76/immortals:0.0.1 .
```
### Push
```bash
docker push ghcr.io/goldenhand76/immortals:0.0.1
```
