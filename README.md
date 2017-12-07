tiny-redir is a tiny Go daemon for redirecting HTTP requests. It can be used to redirect any URL to one destination or to 'remap' a domain name by redirecting to a URL while preserving the original request's URI.

**options**

```
Usage of tiny-redir:
  -append-uri
        append URI? (default true)
  -dest string
        destination URL
  -listen-addr string
        network address to liste on (default "0.0.0.0:8000")
  -permanent
        use 301 permanent redirect
```