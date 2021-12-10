package p2pserver
// Yue Yu
import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fishjump/cs7ns1_project3/p2p-server/entities"

	"github.com/withmandala/go-log"
)

var (
	logger log.Logger
)

func init() {
	logger = *log.New(os.Stderr)
}

type (
	HttpHandler            func(http.ResponseWriter, *http.Request)
	MessageHandlerCallback func(string, *entities.MessageRequest)
	HttpHandlerMap         map[string]HttpHandler

	P2PServer struct {
		server     *http.Server
		Record     *entities.NodeTokenRecord
		privateKey string
		cert       string
		msgCbk     MessageHandlerCallback
	}
)

func NewServer(host string, port int, privateKey string, cert string, ca string, msgCbk MessageHandlerCallback) *P2PServer {
	caCert, err := ioutil.ReadFile(ca)
	if err != nil {
		fmt.Println(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	server := &P2PServer{
		server: &http.Server{
			Addr: fmt.Sprintf("%s:%d", host, port),
			TLSConfig: &tls.Config{
				ClientCAs:  caCertPool,
				ClientAuth: tls.RequireAndVerifyClientCert,
			},
		},
		cert:       cert,
		privateKey: privateKey,
		msgCbk:     msgCbk,
		Record:     entities.CreateNodeTokenRecord(),
	}

	handler := &http.ServeMux{}
	handler.HandleFunc("/register", server.register)
	handler.HandleFunc("/unregister", server.unregister)
	handler.HandleFunc("/healthz", server.healthz)
	handler.HandleFunc("/list", server.list)
	handler.HandleFunc("/message", server.message)
	server.server.Handler = handler

	return server
}

func (s *P2PServer) RunTLS() {
	if err := s.server.ListenAndServeTLS(s.cert, s.privateKey); err != nil {
		logger.Error(err)
	}
}
