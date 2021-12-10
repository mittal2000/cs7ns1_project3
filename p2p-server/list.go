package p2pserver
// Yue Yu
import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fishjump/cs7ns1_project3/p2p-server/entities"
)

func (s *P2PServer) list(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error(err)
	}

	request := &entities.ListRequest{}
	json.Unmarshal(body, request)

	if _, find := s.Record.FindByToken(request.Token); !find {
		rw.WriteHeader(http.StatusBadRequest)

		response := entities.ListResponse{
			Status: false,
			Reason: "Wrong Token",
		}

		jsonStr, err := json.Marshal(response)
		if err != nil {
			logger.Error(err)
		}

		rw.Write(jsonStr)
	} else {
		rw.WriteHeader(http.StatusOK)

		response := entities.ListResponse{
			Status: true,
			Reason: "",
			Nodes:  s.Record.GetNodes(),
		}

		jsonStr, err := json.Marshal(response)
		if err != nil {
			logger.Error(err)
		}

		rw.Write(jsonStr)
	}
}
