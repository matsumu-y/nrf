/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Management

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	AddService(router)
	return router
}

func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/nnrf-nfm/v1")

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		case "PATCH":
			group.PATCH(route.Pattern, route.HandlerFunc)
		}
	}

	return group
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"DeregisterNFInstance",
		strings.ToUpper("Delete"),
		"/nf-instances/:nfInstanceID",
		DeregisterNFInstance,
	},

	{
		"GetNFInstance",
		strings.ToUpper("Get"),
		"/nf-instances/:nfInstanceID",
		GetNFInstance,
	},

	{
		"RegisterNFInstance",
		strings.ToUpper("Put"),
		"/nf-instances/:nfInstanceID",
		RegisterNFInstance,
	},

	{
		"UpdateNFInstance",
		strings.ToUpper("Patch"),
		"/nf-instances/:nfInstanceID",
		UpdateNFInstance,
	},

	{
		"GetNFInstances",
		strings.ToUpper("Get"),
		"/nf-instances",
		GetNFInstances,
	},

	{
		"RemoveSubscription",
		strings.ToUpper("Delete"),
		"/subscriptions/:subscriptionID",
		RemoveSubscription,
	},

	{
		"UpdateSubscription",
		strings.ToUpper("Patch"),
		"/subscriptions/:subscriptionID",
		UpdateSubscription,
	},

	{
		"CreateSubscription",
		strings.ToUpper("Post"),
		"/subscriptions",
		CreateSubscription,
	},
}
