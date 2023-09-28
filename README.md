# router

<a href="https://gitpod.io/#https://github.com/gouniverse/router" style="float:right;"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

A declarative router running on top of Chi.

## Why another router?

Just to avoid several frustrations and multiple shortcomings with the existing routers:

- Routing declaration should be clear, easy and concise without the extra bloat and cognitive load.

- Routing should not be "obfuscated" behind the implementation. Declaration should be simple enough to understand without additional explanations.

- The standard handlers in Go do not expect anything to be returned. As a result quite often an error is handled but a return statement is simply forgotten, and errors fall through. To avoid this common scenario, here the handlers expect a string to be returned.

- Simplify declaring routes. Writing the HTTP method each time is very repetitive, and usually not needed for most endpoints. Implicitly all routes respond to all HTTP verbs. Unless explicitly specified.

- Middlewares are defined for each route explicitly.

- Easy to test.

- Lightweight

- Fast

- List

## Installation

```sh
go get -u github.com/gouniverse/router
```

## Example Routes

```go
checkUserUnauthenticatedMiddleware := Middleware{
    Name: "Check if User is Unauthenticated"
    Handler: middleware.CheckUserAuthenticated,
}

routes = []router.Route{
    // Example of simple "Hello world" endpoint
    {
        Name: "Home",
        Path: "/",
        Handler: func(w http.ResponseWriter, r *http.Request) string {
            return "Hello world"
        },
    },
    // Example of POST route
    {
        Name: "Submit Form",
        Path: "/form-submit",
        Methods: []string{http.MethodPost],
        Handler: func(w http.ResponseWriter, r *http.Request) string {
            return "Form submitted"
        },
    },
    // Example of route with local middlewares
    {
        Name: "User Dashboard",
        Path: "/user/dashboard",
        Middlewares: []Middleware{
		    checkUserUnauthenticatedMiddleware,
        },
        Handler: func(w http.ResponseWriter, r *http.Request) string {
            return "Welcome to your dashboard"
        },
    },
    // Catch-all endpoint
    {
        Name: "Catch All. Page Not Found",
        Path: "/*",
        Handler: func(w http.ResponseWriter, r *http.Request) string {
            return "Page not found"
        },
    },
}
```

## Example with Chi

```go
// 1. Prepare your global middleware
globalMiddlewares := []Middleware{}
globalMiddlewares = append(globalMiddlewares, router.Middleware{
    Name:    "Compress",
    Handler: middleware.Compress(5, "text/html", "text/css"),
})

globalMiddlewares = append(globalMiddlewares, router.Middleware{
    Name:    "GetHead",
    Handler: middleware.GetHead,
})

globalMiddlewares = append(globalMiddlewares, router.Middleware{
    Name:    "CleanPath",
    Handler: middleware.CleanPath,
})

globalMiddlewares = append(globalMiddlewares, router.Middleware{
    Name:    "RedirectSlashes",
    Handler: middleware.RedirectSlashes,
})

globalMiddlewares = append(globalMiddlewares, router.Middleware{
    Name:    "Timeout",
    Handler: middleware.Timeout(time.Second * 30),
})

globalMiddlewares = append(globalMiddlewares, router.Middleware{
    Name:    "RateLimit 1/second",
    Handler: httprate.LimitByIP(20, 1*time.Second),
})

globalMiddlewares = append(globalMiddlewares, router.Middleware{
    Name:    "RateLimit 1/minute",
    Handler: httprate.LimitByIP(180, 1*time.Minute),
})

globalMiddlewares = append(globalMiddlewares, router.Middleware{
    Name:    "RateLimit 1/hour",
    Handler: httprate.LimitByIP(12000, 1*time.Hour),
})

globalMiddlewares = append(globalMiddlewares, router.Middleware{
    Name:    "Logger",
    Handler: middleware.Logger,
})

globalMiddlewares = append(globalMiddlewares, router.Middleware{
    Name:    "Recoverer",
    Handler: middleware.Recoverer,
})

// 2. Prepare your routes
routes := []router.Route{}
routes = append(routes, adminControllers.Routes()...)
routes = append(routes, userControllers.Routes()...)
routes = append(routes, websiteControllers.Routes()...)

// Get a Chi router
chiRouter := router.NewChiRouter(globalMiddlewares, routes)

// Now you can use it
http.ListenAndServe(":3000", chiRouter)
```

## Example Applying Path to Multiple Routes

RoutesPrependPath is a helper method allowing you to quickly add
a path to the beginning of a group of routes

```go
// Prepend /user to the path of the user routes
userRoutes = router.RoutesPrependPath(userRouters, "/user")

// Prepend /admin to the path of the admin routes
adminRoutes = router.RoutesPrependPath(adminRouters, "/admin")
```

## Example Applying Middleware to Multiple Routes


RoutesPrependMiddlewares is a helper method allowing you to quickly add
local middlewares to a group of routes. These middlewares are applied to
the beginning and will be called first, before the ones already defined

```go
router.RoutesPrependMiddlewares(userRouters, []func(http.Handler) http.Handler{
    middleware.CheckUserAuthenticated,
})
```

## Listing Routes

This router allows you to list routes for easy preview

```
+------------------------------------+
| GLOBAL MIDDLEWARE LIST (TOTAL: 2)  |
+---+--------------------------------+
| # | MIDDLEWARE NAME                |
+---+--------------------------------+
| 1 | Append JWT Token               |
| 2 | Append Session Cookies         |
+---+--------------------------------+
+-------------------------------------------------------------------------------------------------+
| ROUTES LIST (TOTAL: 5)                                                                          |
+---+-----------------+------------+---------------------------+----------------------------------+
| # | ROUTE PATH      | METHODS    | ROUTE NAME                | MIDDLEWARE NAME LIST             |
+---+-----------------+------------+---------------------------+----------------------------------+
| 1 | /               | [ALL]      | Home                      | [Web Middleware]                 |
| 2 | /example        | [GET POST] | Example                   | [Web Middleware]                 |
| 3 | /api/form-submit| [POST]     | Submit Form               | [API Middleware, Verify Form]    |
| 4 | /user/dashboard | [ALL]      | User Dashboard            | [Check if User is Authenticated] |
| 5 | /*              | [ALL]      | Catch All. Page Not Found | []                               |
+---+-----------------+------------+---------------------------+----------------------------------+
```
