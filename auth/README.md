# Simple Auth sample

- Use htpasswd in [apache2-utils](https://packages.debian.org/stretch/apache2-utils) to create username and password.

```
htpasswd -cB -b auth.htpasswd myuser mypass
```

- Start there registry 

```
docker run -it --rm -p 5000:5000 \
    -v $(pwd)/auth.htpasswd:/etc/docker/registry/auth.htpasswd \
    -e REGISTRY_AUTH="{htpasswd: {realm: localhost:5000, path: /etc/docker/registry/auth.htpasswd}}" \
    registry
```

```
docker pull busybox:latest
docker tag busybox:latest 
docker push localhost:5000/busybox:latest
```

### Build and Run the sample 

```
$ go build && go run .
sha256:205a121ea8a7a142e5f1fdb9ad72c70ffc8e4a56efec5b70b78a93ebfdddae87 
```

The following enables setting up the credentials provider for a repository client - 

https://github.com/sajayantony/oras-samples/blob/aaecba31d3cb43927eea0ba514df7ad1c417bbad/auth/main.go#L19-L29
