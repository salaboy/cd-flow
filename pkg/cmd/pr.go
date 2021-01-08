package cmd

import (
	"context"
	"encoding/json"
	"log"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(prCmd)
	prCmd.AddCommand(prCreatedCmd)
	prCmd.AddCommand(prMergedCmd)
	prCmd.AddCommand(prClosedCmd)

	prCmd.PersistentFlags().StringVarP(&prId, "id", "i", "", "PR's Id")
	prCmd.PersistentFlags().StringVarP(&prTitle, "title", "t", "", "PR's Title")
	prCmd.PersistentFlags().StringVarP(&prRepoName, "repository", "r", "", "PR's Repository")
	prCmd.PersistentFlags().StringVarP(&prAuthor, "author", "a", "", "PR's Author")
	prCmd.PersistentFlags().StringToStringVarP(&prData, "data", "d", map[string]string{}, "PR's Data")
}

var (
	prId       string
	prTitle    string
	prRepoName string
	prAuthor   string
	prData     map[string]string
)

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Pull Request Events",
	Long:  `Emit PR related CloudEvents`,
}

var prCreatedCmd = &cobra.Command{
	Use:   "created",
	Short: "Emit Pull Request Created Event",
	Long:  `Emit Pull Request Created CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event := cloudevents.NewEvent()
		event.SetID("abc-123") //Generate with UUID
		event.SetSource("cdf-events")
		event.SetType("CDF.PR.Created")
		event.SetTime(time.Now())

		setExtensionForPREvents(event)

		event.SetData(cloudevents.ApplicationJSON, prData) //

		// Set a target.
		ctx := cloudevents.ContextWithTarget(context.Background(), CDF_SINK)

		// Send that Event.
		log.Printf("sending event %s\n", event)

		if result := c.Send(ctx, event); !cloudevents.IsACK(result) {
			log.Fatalf("failed to send, %v", result)
			return result
		}

		return nil
	},
}

var prMergedCmd = &cobra.Command{
	Use:   "merged",
	Short: "Emit Pull Request Merged Event",
	Long:  `Emit Pull Request Merged CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event := cloudevents.NewEvent()
		event.SetID("abc-123") //Generate with UUID
		event.SetSource("cdf-events")
		event.SetType("CDF.PR.Merged")
		event.SetTime(time.Now())

		setExtensionForPREvents(event)

		event.SetData(cloudevents.ApplicationJSON, prData) //

		// Set a target.
		ctx := cloudevents.ContextWithTarget(context.Background(), CDF_SINK)

		// Send that Event.
		log.Printf("sending event %s\n", event)

		if result := c.Send(ctx, event); !cloudevents.IsACK(result) {
			log.Fatalf("failed to send, %v", result)
			return result
		}

		return nil
	},
}

var prClosedCmd = &cobra.Command{
	Use:   "closed",
	Short: "Emit Pull Request Closed Event",
	Long:  `Emit Pull Request Closed CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event := cloudevents.NewEvent()
		event.SetID("abc-123") //Generate with UUID
		event.SetSource("cdf-events")
		event.SetType("CDF.PR.Closed")
		event.SetTime(time.Now())

		setExtensionForPREvents(event)

		event.SetData(cloudevents.ApplicationJSON, prData) //

		// Set a target.
		ctx := cloudevents.ContextWithTarget(context.Background(), CDF_SINK)

		// Send that Event.
		log.Printf("sending event %s\n", event)

		if result := c.Send(ctx, event); !cloudevents.IsACK(result) {
			log.Fatalf("failed to send, %v", result)
			return result
		}

		return nil
	},
}

func setExtensionForPREvents(event cloudevents.Event) {
	event.SetExtension("cdfprid", prId)
	event.SetExtension("cdfprrepository", prRepoName)
	event.SetExtension("cdfprauthor", prAuthor)
	event.SetExtension("cdfprtitle", prTitle)

	var extension = map[string]string{
		"cdfprid":         prId,
		"cdfprrepository": prRepoName,
		"cdfprauthor":     prAuthor,
		"cdfprtitle":      prTitle,
	}

	bytes, err := json.Marshal(extension)
	if err != nil {
		log.Fatalf("failed to marsha extension, %v", err)
	}
	event.SetExtension("cdf", bytes)
}
