package router

import (
	helper "ShorUrl/Helper"
	"fmt"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	fmt.Println("Inside Router")
	// router.HandleFunc("/", helper.ShowSomethingOnBrowser)
	router.HandleFunc("/convertIntoshortUrl", helper.ConvertIntoShortUrl)
	router.HandleFunc(`/shortUrl/{id:[0-9]*}`, helper.RedirectToOriginalUrl)
	return router
}
