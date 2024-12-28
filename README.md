# quitable-docker-selenium

[dockerhub](https://hub.docker.com/r/nakanishi123/quitable-standalone-chrome)

```bash
docker run --rm -p 4444:4444 -p 7900:7900 -p 11451:11451 --shm-size="2g" nakanishi123/quitable-standalone-chrome:131.0
```

```bash
curl http://localhost:11451/shutdown
```
