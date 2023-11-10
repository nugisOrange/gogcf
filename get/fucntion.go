package petasal

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/nugisOrange/petback"
)

func init() {
	functions.HTTP("petback", petaSalGet)
}

func petaSalGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://nugisorange.github.io/")
	fmt.Fprintf(w, petback.GCFHandler("MONGOJAMBE", "petasalget", "petback"))
}