package main
import (
   "context"
   "fmt"
   "log"
   "net/http"
   "golang.ngrok.com/ngrok"
   "golang.ngrok.com/ngrok/config"
)

func main() {
   ctx := context.Background()
   l, err := ngrok.Listen(ctx,
       config.HTTPEndpoint(),
       ngrok.WithAuthtokenFromEnv(),
   )
   if err != nil {
       log.Fatal(err)
   }
   fmt.Println("ngrok ingress url: ", l.Addr())
   http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       fmt.Fprintln(w, "Hello from your ngrok-delivered Go app!")
   }))
}
