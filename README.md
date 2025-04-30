# dis-router-nginx-poc

---

:warning: This repository was archived in May 2025 and is no longer in development. :warning:

---

Several POCS for exploring the concept of an NGINX router configured from an external source

## Getting started

### Dependencies

- colima
- nvm

### Local running

#### Openresty

This stack is for running:

- [OpenResty](https://openresty.org/en/) for proxying / redirects
- http-echo as an upstream server
- router-api to fake an API serving redirect configuration

Currently this is configured just for redirects with a starting set held in ./pocs/router-api/data/db.json

To start the stack you can run (from the `./pocs/openresty` directory)

```sh
    make up
```

Or to have it reload on nginx / lua change:

```sh
    make watch
```

The openresty server will load redirects from the router-api. When router-api's "database" is updated it currently fires a webhook at the openresty server, causing it to reload it's configuration. You can do this by sending a POST request to localhost:7777/redirects with a body like so:

```json
    {
        "from": "/origin/carrots",
        "to": "/destination/carrots"
    }
```

Alternatively you can trigger the reload by sending a request to localhost:8080/webhook/redirect-config

Requests to be proxied / redirected can be send to localhost:8080/${path}

This POC does not currently have any facility for wildcard redirecting.

#### Nginx sidecar

This stack is for running:

- nginx for proxying / redirects
- http-echo as an upstream server(s)
- router-api to fake an API serving redirect configuration

This is configured for both redirects and path based routing with a starting set held in ./pocs/router-api/data/db.json

To start the stack you can run (from the `./pocs/nginx-sidecar` directory)

```sh
    make up
```

The sidecar server will load redirects from the router-api. When router-api's "database" is updated it currently fires a webhook at the sidecar server, causing it to write the nginx configurationa and then reload nginx. You can do this by sending a POST request to localhost:7777/redirects with a body like so:

```json
    // redirect
    {
        "from": "/origin/carrots",
        "to": "/destination/carrots"
    }

    // route
    {
        "path": "/sausages",
        "service": "http-echo-server-1:5678"
    }
```

Alternatively you can trigger the reload by sending a request to localhost:5000/update-config

Requests to be proxied / redirected can be send to localhost:8080/${path}

This service will accept wildcard routing inside the `from` directive for redirects or the `path` directive for path based routing.

#### Nginx with retry

This POC runs:

- nginx as proxy
- http-echo server for successful requests
- not-found-server for 404s

If you request `/business` it will be returned not found by the `not-found-server` and then proxy the request through to the http-echo server.

If you request `/economy/*` the request will be returned by the `not-found-server`.

#### Redis

This POC combines several of the above concepts as well as that looked at by dis-routing-go-poc

This [POC](./redis) runs:

- nginx as initial proxy
- http-echo server to represent legacy (babbage)
- not-found-server to represent new service (wagtail)
- redirector as a redirecting service for legacy
- redirect-api as an api to interact with redis
- redis for configuration storage

The command to run is:

```sh
    make watch
```

On initial setup you can do the following on localhost:8080 (nginx):

- `/consumer-price-inflation/bulletin` -> wagtail (proxied by nginx)
- `/economy/cpi/bulletin` -> babbage (proxied by nginx and the redirector)
- `/releases/babbagerelease` -> babbage (proxied by nginx and the redirector)[^1]
- `/releases/wagtailrelease` -> wagtail (proxied by nginx)

[^1]: For paths in `/releases/` nginx will try wagtail first to see if it has a page, then try the redirector for proxying onwards.

The redirector will evaluate if there is a redirect that matches the path provided. If it matches it will redirect, otherwise proxy onwards.

To create a new redirect you can:

```json
    // POST http://localhost:3003/redirects
    [
        {
            "path": "/economy/cpi/bulletin",
            "redirect": "/consumer-price-inflation/bulletin",
            "type": "permenant"
        }
    ]
```

Now when you request `/economy/cpi/bulletin` it will redirect to `/consumer-price-inflation/bulletin`

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

## License

Copyright © 2025, Office for National Statistics <https://www.ons.gov.uk>

Released under MIT license, see [LICENSE](LICENSE.md) for details.
