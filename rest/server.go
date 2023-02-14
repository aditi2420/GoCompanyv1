package rest

import (
	"fmt"
	"go-company/configs"
	"net/http"
)

type ApiServer struct{
	 http.Server
}

func NewApiServer() *ApiServer{
	 server := ApiServer{
		 http.Server{
			Addr: (configs.ApiServer + ":" +  configs.APIServerPort),
		},
	}
	// add routers 
	router := NewRouter()
	server.Handler = router
	fmt.Println("initialising web server %s", configs.ApiServer)
	return &server

}

