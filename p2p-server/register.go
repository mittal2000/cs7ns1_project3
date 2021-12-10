package p2pserver
// Yue Yu
import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fishjump/cs7ns1_project3/p2p-server/entities"
)

func (s *P2PServer) register(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error(err)
	}

	request := &entities.RegisterRequest{}
	json.Unmarshal(body, request)

	token := entities.GenToken(request.Name)
	s.Record.Add(token, request.Name)

	response := &entities.RegisterResponse{
		Status: true,
		Reason: "",
		Token:  token,
	}

	jsonStr, err := json.Marshal(response)
	if err != nil {
		logger.Error(err)
	}

	rw.Write(jsonStr)
}
