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
	rootCmd.AddCommand(repoCmd)
	repoCmd.AddCommand(repoCreatedCmd)
	repoCmd.AddCommand(repoDeletedCmd)

	repoCmd.PersistentFlags().StringVarP(&repoName, "name", "n", "", "Repository's name")
	repoCmd.PersistentFlags().StringVarP(&issueTitle, "url", "u", "", "Repository's URL")
	repoCmd.PersistentFlags().StringToStringVarP(&repoData, "data", "d", map[string]string{}, "Repository's Data")
}

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Emit Repository related Events",
	Long:  `Emit Repository related CloudEvent`,
}

var (
	repoName string
	repoUrl  string
	repoData map[string]string
)

var repoCreatedCmd = &cobra.Command{
	Use:   "created",
	Short: "Emit Repository Created Event",
	Long:  `Emit Repository Created CloudEvent`,
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
		event.SetType("CDF.Repository.Created")
		event.SetTime(time.Now())

		setExtensionForRepoEvents(event)

		event.SetData(cloudevents.ApplicationJSON, repoData) //

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

var repoDeletedCmd = &cobra.Command{
	Use:   "deleted",
	Short: "Emit Repository Deleted Event",
	Long:  `Emit Repository Deleted CloudEvent`,
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
		event.SetType("CDF.Repository.Deleted")
		event.SetTime(time.Now())

		setExtensionForRepoEvents(event)

		event.SetData(cloudevents.ApplicationJSON, repoData) //

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

func setExtensionForRepoEvents(event cloudevents.Event) {
	event.SetExtension("cdfreponame", repoName)
	event.SetExtension("cdfrepourl", repoUrl)

	var extension = map[string]string{
		"cdfreponame": repoName,
		"cdfrepourl":  repoUrl,
	}

	bytes, err := json.Marshal(extension)
	if err != nil {
		log.Fatalf("failed to marsha extension, %v", err)
	}
	event.SetExtension("cdf", bytes)
}
