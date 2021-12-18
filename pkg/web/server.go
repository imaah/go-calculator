package web

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"emorisse.fr/go-calculator/pkg/operation"
	"emorisse.fr/go-calculator/pkg/parser"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "HTTP ", log.Flags())
}

func StartServer(address, port string) {
	var mux = http.NewServeMux()

	var staticFolder = http.FileServer(http.Dir("./static"))

	mux.Handle("/", staticFolder)
	mux.HandleFunc("/api/calculate", handleApiCalculate)

	var middleware = logMiddleware(mux)

	var addr = fmt.Sprintf("%s:%s", address, port)

	logger.Printf("Starting server listening to port %s...\n", addr)
	err := http.ListenAndServe(addr, middleware)

	if err != nil {
		logger.Fatalln(err)
	}
}

func logMiddleware(next http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger.Printf("%s %s %s\n", req.Method, req.URL, req.Proto)
		next.ServeHTTP(res, req)
	}
}

func handleApiCalculate(res http.ResponseWriter, req *http.Request) {
	// TODO : add more security, and input checking

	res.Header().Set("Content-Type", "application/json")

	if req.Method == "POST" {
		res.WriteHeader(200)

		var ope, err = readCalculation(req.Body)

		if err != nil {
			sendJson(errorRes{Error: err.Error()}, res)
		}

		var result = fmt.Sprintf("%s", ope.Eval().GetString())
		sendJson(resultRes{result}, res)

		return
	}

	res.WriteHeader(400)
	sendJson(errorRes{"Bad Request"}, res)

}

func readCalculation(reader io.Reader) (operation.Operation, error) {
	var content, err = ioutil.ReadAll(reader)
	var ope operation.Operation

	if err != nil {
		return nil, err
	}

	var req calculationReq

	err = json.Unmarshal(content, &req)

	if err != nil {
		return nil, err
	}

	ope, err = parser.Parse(req.Calculation)

	if err != nil {
		return nil, err
	}

	return ope, nil
}

func sendJson(elem interface{}, res http.ResponseWriter) {
	var jsonElem, _ = json.Marshal(elem)
	var _, err = res.Write(jsonElem)

	if err != nil {
		fmt.Printf("Failed to send data to client : %s\n", err)
	}
}
