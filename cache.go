/* Package cache provides cache capabilities with rich support of options and rules.

Example code:


		 import (
		 	"time"

		 	"github.com/get-ion/ion"
		 	"github.com/get-ion/ion/context"
		 	"github.com/get-ion/cache"
		 )

		 func main(){
		 	app := ion.Default()
		 	cachedHandler := cache.WrapHandler(h, 2 *time.Minute)
		 	app.Get("/hello", cachedHandler)
		 	app.Run(ion.Addr(":8080"))
		 }

		 func h(ctx context.Context) {
		 	ctx.HTML("<h1> Hello, this should be cached. Every 2 minutes it will be refreshed, check your browser's inspector</h1>")
		 }
*/

package cache

import (
	"time"

	"github.com/get-ion/cache/client"
	"github.com/get-ion/ion/context"
)

// Cache accepts two parameters
// first is the context.Handler which you want to cache its result
// the second is, optional, the cache Entry's expiration duration
// if the expiration <=2 seconds then expiration is taken by the "cache-control's maxage" header
// returns context.Handler, which you can use as your default router or per-route handler
//
// All types of response can be cached, templates, json, text, anything.
//
// You can add validators with this function.
func Cache(bodyHandler context.Handler, expiration time.Duration) *client.Handler {
	return client.NewHandler(bodyHandler, expiration)
}

// WrapHandler accepts two parameters
// first is the context.Handler which you want to cache its result
// the second is, optional, the cache Entry's expiration duration
// if the expiration <=2 seconds then expiration is taken by the "cache-control's maxage" header
// returns context.Handler, which you can use as your default router or per-route handler
//
// All types of response can be cached, templates, json, text, anything.
//
// it returns a context.Handler, for more options use the `Cache`
func WrapHandler(bodyHandler context.Handler, expiration time.Duration) context.Handler {
	return Cache(bodyHandler, expiration).ServeHTTP
}

// Handler accepts one single parameter:
// the cache Entry's expiration duration
// if the expiration <=2 seconds then expiration is taken by the "cache-control's maxage" header
// returns context.Handler.
//
// It's the same as Cache and WrapHandler but it sets the "bodyHandler" to the next handler in the chain.
//
// All types of response can be cached, templates, json, text, anything.
//
// it returns a context.Handler, for more options use the `Cache`.
func Handler(expiration time.Duration) context.Handler {
	h := WrapHandler(nil, expiration)
	return h
}

var (
	// NoCache disables the cache for a particular request,
	// can be used as a middleware or called manually from the handler.
	NoCache = client.NoCache
)
