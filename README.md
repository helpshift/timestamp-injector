# Traefik Plugin - Timestamp Injector

A Traefik middleware plugin that adds an 'epoch-seconds' header to HTTP responses, containing the current Unix time in `seconds.milliseconds` format (similar to Nginx `$msec`).

## Static Configuration

Enable experimental plugins and specify this plugin:

```yaml
experimental:
  plugins:
    timestampinjector:
      moduleName: github.com/helpshift/timestamp-injector
      version: v1.0.0
```

## Dynamic Configuration

Attach the middleware to a router:

```yaml
http:
  middlewares:
    add-timestamp-header:
      plugin:
        timestampinjector: {}

  routers:
    my-router:
      rule: "Host(`example.com`)"
      service: my-service
      middlewares:
        - add-timestamp-header
```

## How It Works

This plugin injects a header into each response with the current Unix time (seconds.milliseconds).

## How to install plugin

```
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
    name: add-timestamp-header
    namespace: my-namespace
spec:
    plugin:
        timestampinjector: {}
        # OR add custome header name
        # Headers:
        #   HeaderName: mytimestamp

```
