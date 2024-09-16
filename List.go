package router

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"
)

func List(globalMiddlewares []Middleware, routes []RouteInterface) {
	tableRouteRows := []table.Row{}
	for index, route := range routes {
		methods := route.GetMethods()
		middlewares := route.GetMiddlewares()
		if len(methods) < 1 {
			methods = []string{"ALL"}
		}
		if len(methods) == 6 {
			methods = []string{"ALL"}
		}
		path := route.GetPath()
		name := route.GetName()
		middlewareNames := []string{}
		for _, middleware := range middlewares {
			middlewareName := middleware.Name
			if middlewareName == "" {
				middlewareName = "unnamed"
			}
			middlewareNames = append(middlewareNames, middlewareName)
		}
		row := table.Row{index + 1, path, methods, name, middlewareNames}
		tableRouteRows = append(tableRouteRows, row)
	}

	tableMiddlewareRows := []table.Row{}
	for index, middleware := range globalMiddlewares {
		middlewareName := middleware.Name
		if middlewareName == "" {
			middlewareName = "unnamed"
		}
		tableMiddlewareRow := table.Row{index + 1, middlewareName}
		tableMiddlewareRows = append(tableMiddlewareRows, tableMiddlewareRow)
	}

	tableMiddleware := table.NewWriter()
	tableMiddleware.AppendHeader(table.Row{"#", "Middleware Name"})
	tableMiddleware.AppendRows(tableMiddlewareRows)
	tableMiddleware.SetIndexColumn(1)
	tableMiddleware.SetTitle("GLOBAL MIDDLEWARE LIST (TOTAL: " + fmt.Sprint(len(globalMiddlewares)) + ")")
	fmt.Println(tableMiddleware.Render())

	tableRoutes := table.NewWriter()
	tableRoutes.AppendHeader(table.Row{"#", "Route Path", "Methods", "Route Name", "Middleware Name List"})
	tableRoutes.AppendRows(tableRouteRows)
	tableRoutes.SetIndexColumn(1)
	tableRoutes.SetTitle("ROUTES LIST (TOTAL: " + fmt.Sprint(len(routes)) + ")")
	fmt.Println(tableRoutes.Render())
}
