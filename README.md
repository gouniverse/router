# router

<a href="https://gitpod.io/#https://github.com/gouniverse/router" style="float:right;"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

A declarative router running on top of Chi.

## Why another router?

After checking the existing routers, all of them came short from what we expect a router to be:

- Routing declaration should be clear, easy and concise without the extra bloat and cognitive load.

- Routing should not be "obfuscated" behind the implementation. Declaration should be simple enough to understand without additional explanations.

- The standard handlers in Go do not expect anything to be returned. As a result quite often an error is handled but a return statement is simply forgotten, and errors fall through. To avoid this common scenario, here the handlers expect a string to be returned.

- Simplify declaring routes. Writing the HTTP method each time is very repetitive, and usually not needed for most endpoints. Implicitly all routes respond to all HTTP verbs. Unless explicitly specified.

- Middlewares are defined for each route explicitly.

- List routes in table format, with path, name and middlewares applicable

- Easy to test. By returning a string, most routes are easy to test on their own by calling them directly and inspecting the output.

- Lightweight, should not add extra fluff

- Fast, should not slow down the application unnecessarily

## Installation

```sh
go get -u github.com/gouniverse/router
```

## Listing Routes

This router allows you to list routes for easy preview

```golang
router.List(globalMiddlewares, routes)
```

```sh
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


## Example Routes

```go
checkUserAuthenticatedMiddleware := Middleware{
    Name: "Check if User is Authenticated"
    Handler: middleware.CheckUserAuthenticated,
}

routes = []router.RouteInterface{
    // Example of simple "Hello world" endpoint
    &router.Route{
        Name: "Home",
        Path: "/",
        HTMLHandler: func(w http.ResponseWriter, r *http.Request) string {
            return "Hello world"
        },
    },
    // Example of POST route
    &router.Route{
        Name: "Submit Form",
        Path: "/form-submit",
        Methods: []string{http.MethodPost],
        JSONHandler: func(w http.ResponseWriter, r *http.Request) string {
            return api.Success("Form submitted")
        },
    },
    // Example of route with local middlewares
    &router.Route{
        Name: "User Dashboard",
        Path: "/user/dashboard",
        Middlewares: []Middleware{
			checkUserAuthenticatedMiddleware,
        },
        HTMLHandler: func(w http.ResponseWriter, r *http.Request) string {
            return "Welcome to your dashboard"
        },
    },
    // Catch-all endpoint
    &router.Route{
        Name: "Catch All. Page Not Found",
        Path: "/*",
        HTMLHandler: func(w http.ResponseWriter, r *http.Request) string {
            return "Page not found"
        },
    },
}
```

## Example with Chi

```go
// 1. Prepare your global middleware
globalMiddlewares := []Middleware{
    NewCompressMiddleware(5, "text/html", "text/css"),
    NewGetHeadMiddleware(),
    NewCleanPathMiddleware(),
    NewRedirectSlashesMiddleware(),
    NewTimeoutMiddleware(30),
    NewLimitByIpMiddleware(20, 1),       // 20 req per second
	NewLimitByIpMiddleware(180, 60),     // 180 req per minute
	NewLimitByIpMiddleware(12000, 3600), // 12000 req hour
}

// 1.1. Example skipping middlewares while testing
if config.AppEnvironment != config.APP_ENVIRONMENT_TESTING {
	globalMiddlewares = append(globalMiddlewares, NewLoggerMiddleware())
	globalMiddlewares = append(globalMiddlewares, NewRecovererMiddleware())
}

// 1.2. Example of declaring custom middleware (on the fly)
globalMiddlewares = append(globalMiddlewares, router.Middleware{
    Name:    "My Custom Middleware",
    Handler: func (next http.Handler) http.Handler {
        // My custom implementation here
    },
})

// 2. Prepare your routes
routes := []router.RouteInterface{}
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
userRoutes = router.RoutesPrependPath(userRoutes, "/user")

// Prepend /admin to the path of the admin routes
adminRoutes = router.RoutesPrependPath(adminRoutes, "/admin")
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



## Using Controllers

The MVC paradigm is quite popular in software development. This router supports easily supports controllers

```go
routes = []router.RouteInterface{
    // Example of an HTML controller
    &router.Route{
        Name: "HTML Endpoint",
        Path: "/html-endpoint",
        HTMLHandler: (&htmlController{}).Handler,
    },
    // Example of a JSON controller
    &router.Route{
        Name: "JSON Endpoint",
        Path: "/",
        JSONHandler: (&jsonController{}).Handler,
    },
    // Example of an HTML controller
    &router.Route{
        Name: "Idiomatic Endpoint",
        Path: "/",
        HTMLHandler: (&idiomaticController{}).Handler,
    },
}
```

- Definition of a HTML Controller

The HTML controller extends the HTMLControllerInterface.
The Handler method of this controller, returns an HTML string.

```go
type homeController struct{}

var _ router.HTMLControllerInterface = (*homeController)(nil)

func (controller *homeController) Handler(w http.ResponseWriter, r *http.Request) string {
	return "Hello world"
}
```

- Definition of a JSON Controller

The JSON controller extends the JSONControllerInterface.
The Handler method of this controller returns a JSON string

```go
type homeController struct{}

var _ router.JSONControllerInterface = (*homeController)(nil)

func (controller *homeController) Handler(w http.ResponseWriter, r *http.Request) string {
    return api.Success("Hello world")
}
```

- Definition of an Idiomatic Controller

The idiomatic controller extends the ControllerInterface.
The Handler method is a standard Go handler, which does not return anything.

```go
type homeController struct{}

var _ router.ControllerInterface = (*homeController)(nil)

func (controller *homeController) Handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello world"))
}
```