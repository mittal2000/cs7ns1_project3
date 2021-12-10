package main
// Swati Poojary

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	p2pserver "github.com/fishjump/cs7ns1_project3/p2p-server"
	"github.com/fishjump/cs7ns1_project3/p2p-server/entities"

	"github.com/withmandala/go-log"
)

var (
	dir              string
	externalHostName string
	initialIndexHost string
	externalPort     int

	clientToken map[string]string

	external *p2pserver.P2PServer
	client   *http.Client

	wg sync.WaitGroup

	logger *log.Logger
)

func runBackground(fn func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fn()
	}()
}

func init() {
	clientToken = make(map[string]string)

	logger = log.New(os.Stderr)

	flag.StringVar(&dir, "dir", ".", "directory to save data")
	flag.StringVar(&externalHostName, "host", "rasp-019.scss.tcd.ie", "")
	flag.IntVar(&externalPort, "port", 33000, "")
	flag.StringVar(&initialIndexHost, "index", "rasp-019.scss.tcd.ie", "")
	flag.Parse()

	external = p2pserver.NewServer(externalHostName, externalPort,
		dir+"/bundled.key",
		dir+"/bundled.crt",
		dir+"/ca.crt",
		nil)

	external.Record.Add(entities.GenToken(initialIndexHost), initialIndexHost)

	certPair, err := tls.LoadX509KeyPair(dir+"/bundled.crt", dir+"/bundled.key")
	if err != nil {
		logger.Error(err)
	}

	caCert, err := ioutil.ReadFile(dir + "/ca.crt")
	if err != nil {
		logger.Error(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{certPair},
			},
		},
	}
}

func main() {
	runBackground(external.RunTLS)
	runBackground(c.Start)

	wg.Wait()
}
