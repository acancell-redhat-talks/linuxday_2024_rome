### Install `podman` and other container tools

```
sudo dnf install container-tools
```

### Check the `Containerfile`

```
cd demo01-container
cat Containerfile
```

### Build container image from `Containerfile`

```
podman build -t linuxday_2024_rome/demo-container:latest .
podman image ls
```

### Run the container and verify it replies

```
podman run --name demo-container --detach --publish 8080:8000 linuxday_2024_rome/demo-container:latest
curl -v http://127.0.0.1:8080/
```

### Analyze the container

```
podman ps
podman logs demo-container
```

### Verify the container isolation

```
ps -ax | grep http.server
podman exec -it demo-container ps -ax | grep http.server
```
```
sudo ss -tupln
podman exec -it demo-container ss -tupln
```

### Stop the container

```
podman kill demo-container; podman rm demo-container
```

### Containers are just processes ! 

```
podman run --name demo-container --detach --publish 8080:8000 --entrypoint date linuxday_2024_rome/demo-container:latest
podman exec -it demo-container ps -ax
podman ps -a; podman logs demo-container
```

### Containers are ephemeral

```
podman kill demo-container; podman rm demo-container
podman run --name demo-container --detach --publish 8080:8000 linuxday_2024_rome/demo-container:latest
podman cp second.html demo-container:/var/www/html/
curl -v http://127.0.0.1:8080/second.html
```
```
podman kill demo-container; podman rm demo-container
podman run --name demo-container --detach --publish 8080:8000 linuxday_2024_rome/demo-container:latest
curl -v http://127.0.0.1:8080/second.html
```

### Push the container image to a remote registry

```
podman login quay.io
podman image ls
podman tag localhost/linuxday_2024_rome/demo-container:latest quay.io/acancell-redhat-talks/linuxday_2024_rome/demo-container:latest
podman image ls
podman push quay.io/acancell-redhat-talks/linuxday_2024_rome/demo-container:latest
```

### Verify the image was pushed to the remote registry

Location: https://quay.io/repository/acancell-redhat-talks/linuxday_2024_rome/demo-container