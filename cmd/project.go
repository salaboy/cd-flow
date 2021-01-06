package cmd

import (
	"context"
	"encoding/json"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/spf13/cobra"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectCreatedCmd)
	projectCmd.AddCommand(projectDeletedCmd)

	projectCmd.PersistentFlags().StringVarP(&projectName, "name", "n", "", "Project's Name")
	projectCmd.PersistentFlags().StringToStringVarP(&projectData, "data", "d", map[string]string{}, "Project's Data")
}

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Emit Project related Events",
	Long:  `Emit Project related CloudEvent`,

}


var(
	projectName  string
	projectData map[string]string
)

var projectCreatedCmd = &cobra.Command{
	Use:   "created",
	Short: "Emit Project Created Event",
	Long:  `Emit Project Created CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error{
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event :=  cloudevents.NewEvent()
		event.SetID(uuid.NewV4().String())
		event.SetSource("cdf-events")
		event.SetType("CDF.Project.Created")
		event.SetTime(time.Now())

		setExtensionForProjectEvents(event)

		event.SetData(cloudevents.ApplicationJSON, projectData)

		// Set a target.
		ctx := cloudevents.ContextWithTarget(context.Background(), CDF_SINK)

		// Send that Event.
		log.Println("sending event %s", event)

		if result := c.Send(ctx, event); !cloudevents.IsACK(result) {
			log.Fatalf("failed to send, %v", result)
			return result
		}

		return nil
	},

}

var projectDeletedCmd = &cobra.Command{
	Use:   "deleted",
	Short: "Emit Project Deleted Event",
	Long:  `Emit Project Deleted CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error{
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event :=  cloudevents.NewEvent()
		event.SetID(uuid.NewV4().String())
		event.SetSource("cdf-events")
		event.SetType("CDF.Project.Deleted")
		event.SetTime(time.Now())

		setExtensionForProjectEvents(event)

		event.SetData(cloudevents.ApplicationJSON, projectData)//

		// Set a target.
		ctx := cloudevents.ContextWithTarget(context.Background(), CDF_SINK)

		// Send that Event.
		log.Println("sending event %s", event)

		if result := c.Send(ctx, event); !cloudevents.IsACK(result) {
			log.Fatalf("failed to send, %v", result)
			return result
		}

		return nil
	},

}

func setExtensionForProjectEvents(event cloudevents.Event ) {
	event.SetExtension("cdfprojectname", projectName)


	var extension = map[string]string{
		"cdfprojectname":   projectName,
	}

	bytes, err := json.Marshal(extension)
	if err != nil {
		log.Fatalf("failed to marshal extension, %v", err)
	}
	event.SetExtension("cdf", bytes)
}

