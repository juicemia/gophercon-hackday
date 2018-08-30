package api

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	*http.Server
	msgs chan string
}

func NewServer(addr string, msgs chan string) error {
	mux := http.NewServeMux()

	srv := &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: mux,
		},

		msgs: msgs,
	}

	mux.HandleFunc("/message", srv.handleMessages)

	return srv.ListenAndServe()
}

func (s *Server) handleMessages(wr http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		log.Infoln("GET /message")

		wr.WriteHeader(http.StatusOK)
		wr.Write([]byte("default msg"))
		return
	case http.MethodPost, http.MethodPut:
		log.Infof("%v /message", req.Method)

		buf, err := ioutil.ReadAll(req.Body)
		if err != nil {
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}

		select {
		case s.msgs <- string(buf):
		default:
			// send it
		}

		wr.WriteHeader(http.StatusAccepted)
		return
	default:
		log.Infof("%v unsupported method", req.Method)

		wr.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
