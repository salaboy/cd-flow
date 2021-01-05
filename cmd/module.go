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
	projectCmd.PersistentFlags().StringVarP(&moduleName, "name", "n", "", "Module's Name")
	projectCmd.PersistentFlags().StringVarP(&moduleProjectName, "project", "p", "", "Module Project's Name")
	projectCmd.PersistentFlags().StringToStringVarP(&projectData, "data", "d", map[string]string{}, "Project's Data")
}

var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Emit Module related Events",
	Long:  `Emit Module related CloudEvent`,

}


var(
	moduleName  string
	moduleProjectName  string
	moduleData map[string]string
)

var moduleCreatedCmd = &cobra.Command{
	Use:   "created",
	Short: "Emit Module Created Event",
	Long:  `Emit Module Created CloudEvent`,
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
		event.SetType("CDF.Module.Created")
		event.SetTime(time.Now())

		setExtensionForModuleEvents(event)

		event.SetData(cloudevents.ApplicationJSON, moduleData)

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

var moduleDeletedCmd = &cobra.Command{
	Use:   "deleted",
	Short: "Emit Module Deleted Event",
	Long:  `Emit Module Deleted CloudEvent`,
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
		event.SetType("CDF.Module.Deleted")
		event.SetTime(time.Now())

		setExtensionForModuleEvents(event)

		event.SetData(cloudevents.ApplicationJSON, moduleData)//

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

func setExtensionForModuleEvents(event cloudevents.Event ) {
	event.SetExtension("cdfmodulename", moduleName)
	event.SetExtension("cdfmoduleprojectname", moduleProjectName)


	var extension = map[string]string{
		"cdfmodulename": moduleName,
		"cdfmoduleprojectname":   moduleProjectName,
	}

	bytes, err := json.Marshal(extension)
	if err != nil {
		log.Fatalf("failed to marshal extension, %v", err)
	}
	event.SetExtension("cdf", bytes)
}

