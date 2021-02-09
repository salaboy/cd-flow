package cmd

import (
	"context"
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(artifactCmd)
	artifactCmd.AddCommand(artifactBuiltCmd)
	artifactCmd.AddCommand(artifactFailedCmd)
	artifactCmd.AddCommand(artifactTestsStartedCmd)
	artifactCmd.AddCommand(artifactTestsEndedCmd)
	artifactCmd.AddCommand(artifactReleasedCmd)

	artifactCmd.PersistentFlags().StringVarP(&artifactId, "id", "i", "", "Artifact's ID")
	artifactCmd.PersistentFlags().StringVarP(&projectName, "project", "p", "", "Artifact's Project's Name")
	artifactCmd.PersistentFlags().StringVarP(&pipelineId, "pipelineId", "j", "", "Artifact's Pipeline ID")
	artifactCmd.PersistentFlags().StringVarP(&artifactVersion, "version", "v", "", "Artifact's Version")
	artifactCmd.PersistentFlags().StringVarP(&moduleName, "module", "m", "", "Artifact's Module Name")
	artifactCmd.PersistentFlags().StringToStringVarP(&artifactData, "data", "d", map[string]string{}, "Artifact's Data")

}

var artifactCmd = &cobra.Command{
	Use:   "artifact",
	Short: "Emit Artifact related Events",
	Long:  `Emit Artifact related CloudEvent`,
}

var (
	artifactId         string
	artifactVersion	   string
	artifactData       map[string]string
)

var artifactBuiltCmd = &cobra.Command{
	Use:   "built",
	Short: "Emit Artifact Built Event",
	Long:  `Emit Artifact Built CloudEvent`,
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
		event.SetType("CDF.Artifact.Built")
		event.SetTime(time.Now())

		setExtensionForArtifactEvents(event)

		event.SetData(cloudevents.ApplicationJSON, artifactData)

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

var artifactTestsStartedCmd = &cobra.Command{
	Use:   "test-started",
	Short: "Emit Artifact Tests Started Event",
	Long:  `Emit Artifact Tests Started CloudEvent`,
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
		event.SetType("CDF.Artifact.TestsStarted")
		event.SetTime(time.Now())

		setExtensionForArtifactEvents(event)

		event.SetData(cloudevents.ApplicationJSON, artifactData) //

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

var artifactTestsEndedCmd = &cobra.Command{
	Use:   "test-ended",
	Short: "Emit Artifact Tests Ended Event",
	Long:  `Emit Artifact Tests Ended CloudEvent`,
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
		event.SetType("CDF.Artifact.TestsEnded")
		event.SetTime(time.Now())

		setExtensionForArtifactEvents(event)

		event.SetData(cloudevents.ApplicationJSON, artifactData) //

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

var artifactReleasedCmd = &cobra.Command{
	Use:   "released",
	Short: "Emit Artifact Released Event",
	Long:  `Emit Artifact Released CloudEvent`,
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
		event.SetType("CDF.Artifact.Released")
		event.SetTime(time.Now())

		setExtensionForArtifactEvents(event)

		event.SetData(cloudevents.ApplicationJSON, artifactData) //

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

var artifactFailedCmd = &cobra.Command{
	Use:   "failed",
	Short: "Emit Artifact Failed Event",
	Long:  `Emit Artifact Failed CloudEvent`,
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
		event.SetType("CDF.Artifact.Failed")
		event.SetTime(time.Now())

		setExtensionForArtifactEvents(event)

		event.SetData(cloudevents.ApplicationJSON, artifactData) //

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

func setExtensionForArtifactEvents(event cloudevents.Event) {
	event.SetExtension("cdfartifactid", artifactId)
	event.SetExtension("cdfartifactversion", artifactVersion)
	if moduleName == "" {
		moduleName = viper.GetString("cdf.module.name")
	}

	if projectName == "" {
		projectName = viper.GetString("cdf.project.name");
	}

	event.SetExtension("cdfmodulename", moduleName)
	event.SetExtension("cdfpipeid", pipelineId)
	event.SetExtension("cdfprojectname", projectName)

	var extension = map[string]string{
		"cdfartifactid":     artifactId,
		"cdfartifactversion": artifactVersion,
		"cdfmodulename": moduleName,
		"cdfpipeid": pipelineId,
		"cdfprojectname": projectName,
	}

	bytes, err := json.Marshal(extension)
	if err != nil {
		log.Fatalf("failed to marshal extension, %v", err)
	}
	event.SetExtension("cdf", bytes)
}
