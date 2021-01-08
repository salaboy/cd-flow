package cmd

import (
	"context"
	"log"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(envCmd)
	envCmd.AddCommand(envUpdatedCmd)
	envCmd.AddCommand(envCreatedCmd)
}

var (
	envId       string
	envTitle    string
	envRepoName string
	envAuthor   string
	envData     map[string]string
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Emit Environment related Events",
	Long:  `Emit Environment related CloudEvent`,
}

var envUpdatedCmd = &cobra.Command{
	Use:   "updated",
	Short: "Emit Environment Updated Event",
	Long:  `Emit Environment Updated CloudEvent`,
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
		event.SetType("CDF.Environment.Updated")
		event.SetTime(time.Now())

		event.SetData(cloudevents.ApplicationJSON, envData) //

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

var envCreatedCmd = &cobra.Command{
	Use:   "created",
	Short: "Emit Environment Created Event",
	Long:  `Emit Environment Created CloudEvent`,
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
		event.SetType("CDF.Environment.Created")
		event.SetTime(time.Now())

		event.SetData(cloudevents.ApplicationJSON, envData) //

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
