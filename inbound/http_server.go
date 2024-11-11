package inbound

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/Bilbottom/ecom-application/config"
	"github.com/Bilbottom/ecom-application/domain"
	"github.com/Bilbottom/ecom-application/inbound/rest"
)

type Server struct {
	appConfig          config.AppConfig
	authService        domain.AuthService
	healthChecker      rest.HealthChecker
	authVerifier       rest.AuthVerifier
	userAddressGetter  rest.UserAddressGetter
	userAddressUpdater rest.UserAddressUpdater
	userLoginner       rest.UserLoginner
	userRegisterer     rest.UserRegisterer
	productGetter      rest.ProductGetter
	productUpdater     rest.ProductUpdater
	cartCheckouter     rest.CartCheckouter
}

func NewServer(
	appConfig config.AppConfig,
	authService domain.AuthService,
	healthChecker rest.HealthChecker,
	authVerifier rest.AuthVerifier,
	userAddressGetter rest.UserAddressGetter,
	userAddressUpdater rest.UserAddressUpdater,
	userLoginer rest.UserLoginner,
	userRegisterer rest.UserRegisterer,
	productGetter rest.ProductGetter,
	productUpdater rest.ProductUpdater,
	cartCheckouter rest.CartCheckouter,
) *Server {
	return &Server{
		appConfig:          appConfig,
		authService:        authService,
		healthChecker:      healthChecker,
		authVerifier:       authVerifier,
		userAddressGetter:  userAddressGetter,
		userAddressUpdater: userAddressUpdater,
		userLoginner:       userLoginer,
		userRegisterer:     userRegisterer,
		productGetter:      productGetter,
		productUpdater:     productUpdater,
		cartCheckouter:     cartCheckouter,
	}
}

func (s *Server) ListenAndServe() error {
	log.Println("Listening on", s.appConfig.Port)

	server := http.Server{
		Addr:         ":" + s.appConfig.Port,
		Handler:      s.createRouter(),
		IdleTimeout:  65 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server.ListenAndServe()
}

func (s *Server) createRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(requestLogger)

	healthcheckHandler := rest.NewHealthCheck(s.healthChecker).ServeHTTP
	router.Get("/healthcheck", healthcheckHandler)

	getProductHandler := rest.NewGetProductHandler(s.productGetter).ServeHTTP
	postProductHandler := rest.NewPostProductHandler(s.productUpdater).ServeHTTP
	getUserAddressHandler := rest.NewGetUserAddressHandler(s.userAddressGetter).ServeHTTP
	postUserAddressHandler := rest.NewPostUserAddressHandler(s.userAddressUpdater).ServeHTTP
	postUserLoginHandler := rest.NewPostUserLoginHandler(s.appConfig, s.authService, s.userLoginner).ServeHTTP
	postUserRegisterHandler := rest.NewPostUserRegisterHandler(s.userRegisterer, s.userLoginner).ServeHTTP
	postCartCheckoutHandler := rest.NewPostCartCheckoutHandler(s.cartCheckouter, s.productGetter).ServeHTTP

	router.Route("/api/v1", func(subRouter chi.Router) {
		subRouter.Post("/login", postUserLoginHandler)
		subRouter.Post("/register", postUserRegisterHandler)

		subRouter.Get("/address", s.authVerifier.WithJWTAuth(getUserAddressHandler))
		subRouter.Post("/address", s.authVerifier.WithJWTAuth(postUserAddressHandler))
		subRouter.Post("/cart/checkout", s.authVerifier.WithJWTAuth(postCartCheckoutHandler))
		subRouter.Get("/products", s.authVerifier.WithJWTAuth(getProductHandler))
		subRouter.Post("/products", s.authVerifier.WithJWTAuth(postProductHandler))
	})

	return router
}

// requestLogger is a custom middleware which prints the incoming request method
// and URL
func requestLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
