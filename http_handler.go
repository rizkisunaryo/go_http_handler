package go_http_handler
import (
"github.com/rizkisunaryo/go_recover"
"net/http"
)

func HttpHandler(header string, handler func()(func(string, http.ResponseWriter, *http.Request), func(string, http.ResponseWriter, *http.Request), func(string, http.ResponseWriter, *http.Request), func(string, http.ResponseWriter, *http.Request))) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header = r.RemoteAddr + ": "+ header +": "

		defer go_recover.Recover(header)

		r.Header.Set("Connection", "close")
		w.Header().Set("Connection", "close")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		getHandler, postHandler, putHandler, deleteHandler := handler()
		switch r.Method {
			case "GET":
				getHandler(header, w, r)
			case "POST":
				postHandler(header, w, r)
			case "PUT":
				putHandler(header, w, r)
			case "DELETE":
				deleteHandler(header, w, r)
		}
	}
}