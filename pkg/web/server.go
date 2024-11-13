package web

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/imaah/go-calculator/pkg/parser"

	"github.com/imaah/go-calculator/pkg/operation"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "HTTP ", log.Flags())
}

func StartServer(address, port string) {
	mux := http.NewServeMux()

	staticFolder := http.FileServer(http.Dir("./static"))

	mux.Handle("/", staticFolder)
	mux.HandleFunc("/api/calculate", handleApiCalculate)

	middleware := logMiddleware(mux)

	addr := fmt.Sprintf("%s:%s", address, port)

	logger.Printf("Starting server listening to port %s...\n", addr)
	err := http.ListenAndServe(addr, middleware)

	if err != nil {
		errFormat := fmt.Errorf("Failed to start the server : %w\n", err)
		logger.Fatalln(errFormat)
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

		ope, err := readCalculation(req.Body)

		if err != nil {
			sendJson(errorRes{Error: err.Error()}, res)
			return
		}

		result := fmt.Sprintf("%f", ope.Eval())
		sendJson(resultRes{result}, res)

		return
	}

	res.WriteHeader(400)
	sendJson(errorRes{"Bad Request"}, res)

}

func readCalculation(reader io.Reader) (operation.Operation, error) {
	content, err := io.ReadAll(reader)
	var ope operation.Operation

	if err != nil {
		return nil, err
	}

	var req calculationReq

	err = json.Unmarshal(content, &req)

	if err != nil {
		return nil, err
	}

	ope, err = parser.ParseV2(req.Calculation)

	if err != nil {
		return nil, err
	}

	return ope, nil
}

func sendJson(elem interface{}, res http.ResponseWriter) {
	jsonElem, _ := json.Marshal(elem)
	_, err := res.Write(jsonElem)

	if err != nil {
		errFormat := fmt.Errorf("Failed to send data to client : %w\n", err)
		logger.Println(errFormat)
	}
}
