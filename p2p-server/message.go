package p2pserver
// Yue Yu
import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/fishjump/cs7ns1_project3/p2p-server/entities"
)

func (s *P2PServer) message(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error(err)
		return
	}

	request := &entities.MessageRequest{}
	json.Unmarshal(body, request)

	name, find := s.Record.FindByToken(request.Token)
	if !find {
		logger.Error(errors.New("not find token"))
		return
	}

	if s.msgCbk != nil {
		s.msgCbk(name, request)
	}

	response := entities.MessageResponse{
		Status: true,
		Reason: "",
	}

	jsonStr, err := json.Marshal(response)
	if err != nil {
		logger.Error(err)
		return
	}

	rw.Write(jsonStr)
}
