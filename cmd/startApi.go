package cmd

import (
	"log"

	"github.com/nachoconques0/quizzer/internal/api"
	"github.com/nachoconques0/quizzer/internal/db"
	"github.com/spf13/cobra"
)

// startApiCmd represents the startApi command
var startApiCmd = &cobra.Command{
	Use:   "startApi",
	Short: "this command will start the API. NEEDED in order to use the quiz",
	Run:   startApi,
}

func init() {
	rootCmd.AddCommand(startApiCmd)
}

func startApi(cmd *cobra.Command, args []string) {
	// init internal DB
	internalDB := db.NewDB()
	internalDB.SeedQuiz()

	// create quiz service
	quizSvc := api.NewService(internalDB)
	server, err := api.NewServer("8080", quizSvc)
	if err != nil {
		log.Fatalf("could not create API server err: %v", err)
	}
	err = server.Run()
	if err != nil {
		log.Fatalf("could not start API server err: %v", err)
	}
}
