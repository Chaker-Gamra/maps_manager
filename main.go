package main

import (
	"net/http"
	"os"

	_ "github.com/Chaker-Gamra/maps_manager/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/Chaker-Gamra/maps_manager/router"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Maps Services Manager APIs
// @version 1.0
// @description This is a microservice that handles communication with different maps services providers. It offers a simple API that abstracts away the complexity of dealing directly with the provider’s api.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey jwtTokenKey
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /
func main() {

	// load .env file
	godotenv.Load(".env")

	//Set the port from env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	//Init New Mux Router
	r := router.InitRouter()

	// Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	//Start The Server
	http.ListenAndServe(":"+port, r)
}
