This is an example openfaas image that logs a lot for the sake of a test suite for the [of-watchdog](github.com/openfaas-incubator/of-watchdog).

Testing process

```
docker build -t chatty-fn:latest .
docker run -p 8080:8080 chatty-fn:latest
```

In other terminal you can call to trigger heavy logging
```
curl localhost:8080
```

Or can run in `-d` mode and use logging functionality to read the output after
