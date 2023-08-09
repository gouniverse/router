# router

<a href="https://gitpod.io/#https://github.com/gouniverse/router" style="float:right;"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

A declarative router running on top of Chi.

## Why another router?

Just to avoid several frustrations with the existing routers:

- Routing declaraion should be clear, easy and concise without the extra bloat.

- Routing should not be "obfuscated" behind the implementation. Declaration should be simple enough to understand without furter explanations.

- The standard handlers in Go do not expect anything to be returned. As a result quite often an error is handled but a return statement is simply forgotten, and errors fall through. To avoid this common scenario, here the handlers expect a string to be returned.

- Simplify declaring routes. Writing the HTTP method each time and is very repetitive, and usually not needed for most endpoints. Implicitly all routes respond to all HTTP verbs. Unless explicitly specified.

- Middlewares are defined for each route explicitly.

## Example Routes
```
routes = router.Route{
    {
        Path: "/",
        Handler: responses.HTMLHandler(func(w http.ResponseWriter, r *http.Request) string {
            return "Hello world"
        }),
    },
    {
        Path: "/*",
        Handler: responses.HTMLHandler(func(w http.ResponseWriter, r *http.Request) string {
            return "Page not found"
        }),
    },
}
```

## Example with Chi
```
globalMiddlewares := []func(http.Handler) http.Handler{
    middleware.Compress(5, "text/html", "text/css"),
    middleware.GetHead,
    middleware.CleanPath,
    middleware.RedirectSlashes,
    middleware.Timeout(time.Second * 30),
    httprate.LimitByIP(20, 1*time.Second),  // 20 req per second
    httprate.LimitByIP(180, 1*time.Minute), // 180 req per minute
    httprate.LimitByIP(12000, 1*time.Hour), // 12000 req hour
}

routes := []router.Route{}
routes = append(routes, adminControllers.Routes()...)
routes = append(routes, userControllers.Routes()...)
routes = append(routes, websiteControllers.Routes()...)


chiRouter := chi.NewRouter()
router.AddMiddlewaresToChiRouter(chiRouter, globalMiddlewares)
router.AddRoutesToChiRouter(chiRouter, routes)
return chiRouter
```

