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
	rootCmd.AddCommand(envCmd)
	envCmd.AddCommand(envCreatedCmd)
	envCmd.AddCommand(envDeletedCmd)

	envCmd.PersistentFlags().StringVarP(&envName, "name", "n", "", "Environment's Name")
	envCmd.PersistentFlags().StringToStringVarP(&envData, "data", "d", map[string]string{}, "Environment's Data")
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Emit Environment related Events",
	Long:  `Emit Environment related CloudEvent`,
}

var (
	envName string
	envRepoUrl string
	envData map[string]string
)

var envCreatedCmd = &cobra.Command{
	Use:   "created",
	Short: "Emit Env Created Event",
	Long:  `Emit Environment Created CloudEvent`,
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
		event.SetType("CDF.Environment.Created")
		event.SetTime(time.Now())

		setExtensionForEnvEvents(event)

		event.SetData(cloudevents.ApplicationJSON, projectData)

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

var envDeletedCmd = &cobra.Command{
	Use:   "deleted",
	Short: "Emit Environment Deleted Event",
	Long:  `Emit Environment Deleted CloudEvent`,
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
		event.SetType("CDF.Project.Deleted")
		event.SetTime(time.Now())

		setExtensionForEnvEvents(event)

		event.SetData(cloudevents.ApplicationJSON, projectData) //

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

func setExtensionForEnvEvents(event cloudevents.Event) {
	event.SetExtension("cdfenvname", envName)
	event.SetExtension("cdfenvrepo", envRepoUrl)

	var extension = map[string]string{
		"cdfenvname": envName,
		"cdfenvrepo": envRepoUrl,
	}

	bytes, err := json.Marshal(extension)
	if err != nil {
		log.Fatalf("failed to marshal extension, %v", err)
	}
	event.SetExtension("cdf", bytes)
}
