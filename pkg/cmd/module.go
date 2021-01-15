package cmd

import (
	"context"
	"encoding/json"
	"log"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(moduleCmd)
	moduleCmd.AddCommand(moduleCreatedCmd)
	moduleCmd.AddCommand(moduleCreatedCmd)
	moduleCmd.PersistentFlags().StringVarP(&moduleName, "name", "n", "", "Module's Name")
	moduleCmd.PersistentFlags().StringVarP(&moduleRepository, "repo", "r", "", "Module's Repository URL")
	moduleCmd.PersistentFlags().StringVarP(&projectName, "project", "p", "", "Module Project's Name")
	moduleCmd.PersistentFlags().StringToStringVarP(&moduleData, "data", "d", map[string]string{}, "Module's Data")
}

var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Emit Module related Events",
	Long:  `Emit Module related CloudEvent`,
}

var (
	moduleName       string
	moduleRepository string
	moduleData       map[string]string
)

var moduleCreatedCmd = &cobra.Command{
	Use:   "created",
	Short: "Emit Module Created Event",
	Long:  `Emit Module Created CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event := cloudevents.NewEvent()
		event.SetID(uuid.NewV4().String())
		event.SetSource("cdf-events")
		event.SetType("CDF.Module.Created")
		event.SetTime(time.Now())

		setExtensionForModuleEvents(event)

		event.SetData(cloudevents.ApplicationJSON, moduleData)

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

var moduleDeletedCmd = &cobra.Command{
	Use:   "deleted",
	Short: "Emit Module Deleted Event",
	Long:  `Emit Module Deleted CloudEvent`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cloudevents.NewDefaultClient()
		if err != nil {
			log.Fatalf("failed to create client, %v", err)
			return err
		}

		// Create an Event.
		event := cloudevents.NewEvent()
		event.SetID(uuid.NewV4().String())
		event.SetSource("cdf-events")
		event.SetType("CDF.Module.Deleted")
		event.SetTime(time.Now())

		setExtensionForModuleEvents(event)

		event.SetData(cloudevents.ApplicationJSON, moduleData) //

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

func setExtensionForModuleEvents(event cloudevents.Event) {
	event.SetExtension("cdfmodulename", moduleName)
	event.SetExtension("cdfmodulerepo", moduleRepository)
	event.SetExtension("cdfprojectname", projectName)

	var extension = map[string]string{
		"cdfmodulename":  moduleName,
		"cdfmodulerepo":  moduleRepository,
		"cdfprojectname": projectName,
	}

	bytes, err := json.Marshal(extension)
	if err != nil {
		log.Fatalf("failed to marshal extension, %v", err)
	}
	event.SetExtension("cdf", bytes)
}
