package datastoreviewer

import (
	"fmt"
        "net/http"
//      "appengine"
//      "appengine/datastore"
)

func datastoreViewerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "datastore viewer")
}
