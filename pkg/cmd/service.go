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
	rootCmd.AddCommand(serviceCmd)
	serviceCmd.AddCommand(serviceCreatedCmd)
	serviceCmd.AddCommand(serviceUpdatedCmd)
	serviceCmd.AddCommand(serviceDeletedCmd)

	serviceCmd.PersistentFlags().StringVarP(&envName, "env", "e", "", "Environment where the Service is running")
	serviceCmd.PersistentFlags().StringVarP(&serviceName, "name", "n", "", "Service's Name")
	serviceCmd.PersistentFlags().StringVarP(&serviceVersion, "version", "v", "", "Service's Version")
	serviceCmd.PersistentFlags().StringVarP(&serviceArtifact, "artifact", "a", "", "Service's Artifact")
	serviceCmd.PersistentFlags().StringToStringVarP(&serviceData, "data", "d", map[string]string{}, "Service's Data")
}

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Emit Environment related Events",
	Long:  `Emit Environment related CloudEvent`,
}

var (
	serviceName string
	serviceVersion string
	serviceArtifact string
	serviceData map[string]string
)

var serviceCreatedCmd = &cobra.Command{
	Use:   "created",
	Short: "Emit Service Created Event",
	Long:  `Emit Service Created CloudEvent`,
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
		event.SetType("CDF.Service.Created")
		event.SetTime(time.Now())

		setExtensionForServiceEvents(event)

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

var serviceUpdatedCmd = &cobra.Command{
	Use:   "updated",
	Short: "Emit Service Updated Event",
	Long:  `Emit Service Updated CloudEvent`,
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
		event.SetType("CDF.Service.Updated")
		event.SetTime(time.Now())

		setExtensionForServiceEvents(event)

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

var serviceDeletedCmd = &cobra.Command{
	Use:   "deleted",
	Short: "Emit Service Deleted Event",
	Long:  `Emit Service Deleted CloudEvent`,
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
		event.SetType("CDF.Service.Deleted")
		event.SetTime(time.Now())

		setExtensionForServiceEvents(event)

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

func setExtensionForServiceEvents(event cloudevents.Event) {
	event.SetExtension("cdfenvname", envName)
	event.SetExtension("cdfservicename", serviceName)
	event.SetExtension("cdfserviceversion", serviceVersion)
	event.SetExtension("cdfserviceartifact", serviceArtifact)

	var extension = map[string]string{
		"cdfenvname": envName,
		"cdfservicename": serviceName,
		"cdfserviceversion": serviceVersion,
		"cdfserviceartifact": serviceArtifact,
	}

	bytes, err := json.Marshal(extension)
	if err != nil {
		log.Fatalf("failed to marshal extension, %v", err)
	}
	event.SetExtension("cdf", bytes)
}
