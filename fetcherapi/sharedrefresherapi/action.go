package sharedrefresherapi

import (
	"bytes"
	"io"
	"net/http"

	"github.com/herb-go/herbdata/datautil/sharedrefresher"
)

//ServeRefreshShared serve refresh shared with given http response writer,request and refresher.
//Return if shared data changed
func ServeRefreshShared(w http.ResponseWriter, r *http.Request, refresher sharedrefresher.SharedRefresher) bool {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return false
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	new, err := refresher.RefreshShared(data)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(new)
	if err != nil {
		panic(err)
	}
	return bytes.Compare(data, new) == 0
}

//NewSharedRefresherAction create action for giver shared refresher.
func NewSharedRefresherAction(refresher sharedrefresher.SharedRefresher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ServeRefreshShared(w, r, refresher)
	}
}
