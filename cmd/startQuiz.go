package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/nachoconques0/quizzer/internal/helpers"
	"github.com/nachoconques0/quizzer/internal/model"
)

// startQuizCmd represents the startQuiz command
var startQuizCmd = &cobra.Command{
	Use:   "startQuiz",
	Short: "This command will start the quiz! REMEMBER. API MUST BE ACTIVE",
	Run: func(cmd *cobra.Command, args []string) {
		startQuiz()
	},
}

func init() {
	rootCmd.AddCommand(startQuizCmd)
}

func startQuiz() {
	// We initiate the prompts to the User
	questions := getQuestionToPrompt()
	runPrompts(questions)
	lb, err := submitAnswers(questions)
	if err != nil {
		helpers.HandleError(err, "Failed to submit answers. Please try again.")
	}
	displayLeaderboard(*lb)
}

func submitAnswers(items []*model.Item) (*model.Leaderboard, error) {
	// Prepare the answers payload
	var answers []model.Answer
	for _, q := range items {
		answers = append(answers, model.Answer{
			QuestionID: uuid.MustParse(q.ID),
			Value:      q.Answer,
		})
	}

	// Make the API request
	responseBytes, err := helpers.MakeAPIRequest(http.MethodPost, "answer", model.SubmitRequest{Answers: answers})
	if err != nil {
		helpers.HandleError(err, "Failed to submit answers. Please try again.")
		return nil, err
	}

	var leaderboard model.Leaderboard
	if err := json.Unmarshal(responseBytes, &leaderboard); err != nil {
		helpers.HandleError(err, "Failed to parse leaderboard response.")
		return nil, err
	}
	return &leaderboard, nil
}

func getQuestionToPrompt() []*model.Item {
	responseBytes, err := helpers.MakeAPIRequest(http.MethodGet, "question", nil)
	if err != nil {
		helpers.HandleError(err, "Error fetching questions")
		return nil
	}

	var questionModel []model.Question
	if err := json.Unmarshal(responseBytes, &questionModel); err != nil {
		helpers.HandleError(err, "Error unmarshalling questions")
		return nil
	}

	var questionsToPrompt []*model.Item
	for _, q := range questionModel {
		questionsToPrompt = append(questionsToPrompt, &model.Item{
			ID:            q.ID.String(),
			Label:         q.Text,
			SelectOptions: q.Answers,
			Answer:        -1,
		})
	}
	return questionsToPrompt
}

func runPrompts(items []*model.Item) {
	doneID := "Done"
	exitID := "Exit"

	items = append([]*model.Item{{ID: doneID, Label: "Submit Answers"}}, items...)
	items = append([]*model.Item{{ID: exitID, Label: "Exit Quiz"}}, items...)

	for {
		templates := &promptui.SelectTemplates{
			Label:    "Questions:",
			Active:   "\U0001F355 {{ .Label | red }}",
			Inactive: "{{ .Label | cyan }}",
			Selected: "\U0001F355 {{ .Label | red }}",
		}

		prompt := promptui.Select{
			Items:        items,
			Templates:    templates,
			Size:         len(items),
			HideSelected: true,
			CursorPos:    0,
		}

		idx, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt error: %v\n", err)
			break
		}

		if items[idx].ID == doneID {
			fmt.Println("Quiz completed.")
			break
		}

		if items[idx].ID == exitID {
			os.Exit(1)
		}

		items[idx].Answer, err = promptSelect(*items[idx])
		if err != nil {
			fmt.Printf("Error answering question: %v\n", err)
		}

		items = append(items[:idx], items[idx+1:]...)
	}
}

func promptSelect(q model.Item) (int, error) {
	prompt := promptui.Select{
		Label:        q.Label,
		Items:        q.SelectOptions,
		HideSelected: true,
	}

	index, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1, err
	}
	return index, nil
}

func displayLeaderboard(leaderboard model.Leaderboard) {
	fmt.Println("")
	fmt.Println("Final result:", leaderboard.Result)
	fmt.Println("############ Leaderboard ############")
	fmt.Println("             Name & Score            ")
	fmt.Println("#####################################")
	fmt.Println("")

	for _, user := range leaderboard.Users {
		fmt.Printf("%v scored: %v \n", user.Name, user.Points)
	}
	fmt.Println("")
}
