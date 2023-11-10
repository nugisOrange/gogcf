package petbackPost

import (
    "fmt"
    "net/http"

    "github.com/GoogleCloudPlatform/functions-framework-go/functions"
    "github.com/nugisOrange/petback"
)

func init() {
    functions.HTTP("petbackPost", petbackPost)
}

func petbackPost(w http.ResponseWriter, r *http.Request) {
    // Set CORS headers for the preflight request
    if r.Method == http.MethodOptions {
        w.Header().Set("Access-Control-Allow-Origin", "https://nugisorange.github.io")
        w.Header().Set("Access-Control-Allow-Methods", "POST")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
        w.Header().Set("Access-Control-Max-Age", "3600")
        w.WriteHeader(http.StatusNoContent)
        return
    }
    // Set CORS headers for the main request.
    w.Header().Set("Access-Control-Allow-Origin", "https://nugisorange.github.io")
    fmt.Fprintf(w, petback.GCFPostHandler("PASETOPRIVATEKEYENV" ,"MONGOSTRING", "petback", "tesgis", r))

}