package cmd

import (
	"github.com/cyber-republic/develap/cmd/serve"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Setup route to different node containers",
	Long:  `Setup route to different node containers`,
	Run: func(c *cobra.Command, args []string) {
		router := mux.NewRouter()

		// Default route
		router.HandleFunc("/", serve.HandleIndex)

		// Node routes
		serve.HandleNodeEndpoints(router)

		// Status routes
		router.HandleFunc("/status/nodes", serve.HandleStatusAllNodesEndpoint)
		router.HandleFunc("/status/nodes/running", serve.HandleStatusRunningNodesEndpoint)
		router.HandleFunc("/status/nodes/stopped", serve.HandleStatusStoppedNodesEndpoint)

		// Handle CORS
		router.Use(mux.CORSMethodMiddleware(router))

		// set timeouts so that a slow or malicious client doesn't
		// hold resources forever
		httpSrv := &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
			Handler:      router,
			Addr:	":5000",
		}

		// Launch HTTP server
		log.Println("Starting server http://localhost:5000")

		log.Fatal(httpSrv.ListenAndServe())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
