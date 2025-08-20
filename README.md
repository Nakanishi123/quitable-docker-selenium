# quitable-docker-selenium

[dockerhub](https://hub.docker.com/r/nakanishi123/quitable-standalone-chrome)

```bash
docker run --rm -p 4444:4444 -p 7900:7900 -p 11451:11451 --shm-size="2g" ghcr.io/nakanishi123/quitable-standalone-chrome:latest
```

```bash
curl http://localhost:11451/shutdown
```

動作確認 -> [http://localhost:7900/?autoconnect=1&resize=scale&password=secret](http://localhost:7900/?autoconnect=1&resize=scale&password=secret)