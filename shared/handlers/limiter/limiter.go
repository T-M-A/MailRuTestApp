package limiter

import (
	"github.com/ulule/limiter/drivers/middleware/stdlib"
	"github.com/ulule/limiter"
	"github.com/ulule/limiter/drivers/store/memory"
	"net/http"
	"github.com/ruelephant/MailRuTestApp/shared/handlers/responce"
)

func GetLimitMiddleware() (*stdlib.Middleware, error) {
	resp := responce.JsonResponse{}

	// Define a limit rate to 1 requests per min.
	rate, err := limiter.NewRateFromFormatted("1-M")
	if err != nil {
		return nil, err
	}

	memoryStore := memory.NewStore()

	exceededHandlerOptions := stdlib.WithLimitReachedHandler(func(w http.ResponseWriter, r *http.Request) {
		resp.Error(w, 509, "Limit exceeded")
	})

	// Create a new middleware with the limiter instance.
	m := stdlib.NewMiddleware(limiter.New(memoryStore, rate), stdlib.WithForwardHeader(true), exceededHandlerOptions)
	return m, nil
}