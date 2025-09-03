# Traefik Plugin - Timestamp Injector

A Traefik middleware plugin that adds an 'epoch-seconds' header to HTTP responses, containing the current Unix time in `seconds.milliseconds` format (similar to Nginx `$msec`).
NOTE: Headername is configurable or else the default is set as epoch-seconds.

## Static Configuration

Enable experimental plugins and specify this plugin:

```yaml
experimental:
  plugins:
    timestamp-injector:
      moduleName: github.com/helpshift/timestamp-injector
      version: v1.0.0
```

## Dynamic Configuration

Create the middleware:

```yaml
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
    name: mytimestamp-middleware
    namespace: my-namespace
spec:
    plugin:
        timestamp-injector: {}
        # OR
        #timestamp-injector:
        #  HeaderName: mytimestamp-header
```

## How It Works

This plugin injects a header into each response with the current Unix time (seconds.milliseconds).

