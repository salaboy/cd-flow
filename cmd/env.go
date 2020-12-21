package cmd

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(envCmd)
	envCmd.AddCommand(envUpdatedCmd)
	envCmd.AddCommand(envCreatedCmd)
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Emit Environment related Events",
	Long:  `Emit Environment related CloudEvent`,


}

var envUpdatedCmd = &cobra.Command{
	Use:   "update",
	Short: "Emit Environment Update Event",
	Long:  `Emit Environment Update CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error{
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event :=  cloudevents.NewEvent()
		event.SetID("abc-123")//Generate with UUID
		event.SetSource("cdf-events")
		event.SetType("CDF.Environment.Updated")
		event.SetData(cloudevents.ApplicationJSON, map[string]string{"hello": "world"})//

		// Set a target.
		ctx := cloudevents.ContextWithTarget(context.Background(), "http://localhost:8080/")

		// Send that Event.
		log.Println("sending event %s", event)

		if result := c.Send(ctx, event); !cloudevents.IsACK(result) {
			log.Fatalf("failed to send, %v", result)
			return result
		}

		return nil
	},

}

var envCreatedCmd = &cobra.Command{
	Use:   "create",
	Short: "Emit Environment Created Event",
	Long:  `Emit Environment Created CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error{
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event :=  cloudevents.NewEvent()
		event.SetID("abc-123")//Generate with UUID
		event.SetSource("cdf-events")
		event.SetType("CDF.Environment.Created")
		event.SetData(cloudevents.ApplicationJSON, map[string]string{"hello": "world"})//

		// Set a target.
		ctx := cloudevents.ContextWithTarget(context.Background(), "http://localhost:8080/")

		// Send that Event.
		log.Println("sending event %s", event)

		if result := c.Send(ctx, event); !cloudevents.IsACK(result) {
			log.Fatalf("failed to send, %v", result)
			return result
		}

		return nil
	},

}