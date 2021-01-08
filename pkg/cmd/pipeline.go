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
	rootCmd.AddCommand(pipelineCmd)
	pipelineCmd.AddCommand(pipelineStartedCmd)
	pipelineCmd.AddCommand(pipelineFinishedCmd)
	pipelineCmd.AddCommand(pipelineFailedCmd)

	pipelineCmd.PersistentFlags().StringVarP(&pipelineId, "id", "i", "", "Pipeline's Id")
	pipelineCmd.PersistentFlags().StringVarP(&pipelineName, "name", "n", "", "Pipeline's Name")
	pipelineCmd.PersistentFlags().StringVarP(&pipelineBranch, "branch", "b", "", "Pipeline's Branch")
	pipelineCmd.PersistentFlags().StringVarP(&pipelineType, "type", "t", "", "Pipeline's Type: Module / Environment")
	pipelineCmd.PersistentFlags().StringVarP(&pipelineModuleName, "module", "p", "", "Pipeline's Module Name")
	pipelineCmd.PersistentFlags().StringVarP(&pipelineEnvName, "env", "e", "", "Pipeline's Environment Name")
	pipelineCmd.PersistentFlags().StringToStringVarP(&pipelineData, "data", "d", map[string]string{}, "Pipeline's Data")

}

var (
	pipelineId         string
	pipelineName       string
	pipelineBranch     string
	pipelineType       string //module/environment
	pipelineModuleName string
	pipelineEnvName    string
	pipelineData       map[string]string
)

var pipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Pipeline Events",
	Long:  `Emit Pipeline related CloudEvents`,
}

var pipelineStartedCmd = &cobra.Command{
	Use:   "started",
	Short: "Emit Pipeline Started Event",
	Long:  `Emit Pipeline Started CloudEvent`,
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
		event.SetType("CDF.Pipeline.Started")
		event.SetTime(time.Now())

		setExtensionForPipelineEvents(event)

		event.SetData(cloudevents.ApplicationJSON, pipelineData) //

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

var pipelineFinishedCmd = &cobra.Command{
	Use:   "finished",
	Short: "Emit Pipeline Finished Event",
	Long:  `Emit Pipeline Finished CloudEvent`,
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
		event.SetType("CDF.Pipeline.Finished")
		event.SetTime(time.Now())

		setExtensionForPipelineEvents(event)

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

var pipelineFailedCmd = &cobra.Command{
	Use:   "failed",
	Short: "Emit Pipeline Failed Event",
	Long:  `Emit Pipeline Failed CloudEvent`,
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
		event.SetType("CDF.Pipeline.Failed")
		event.SetTime(time.Now())

		setExtensionForPipelineEvents(event)

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

func setExtensionForPipelineEvents(event cloudevents.Event) {
	event.SetExtension("cdfpipeid", pipelineId)
	event.SetExtension("cdfpipename", pipelineName)
	event.SetExtension("cdfpipetype", pipelineType)
	event.SetExtension("cdfpipemodulename", pipelineModuleName)
	event.SetExtension("cdfpipeenvname", pipelineEnvName)
	event.SetExtension("cdfpipebranch", pipelineBranch)

	var extension = map[string]string{
		"cdfpipeid":         pipelineId,
		"cdfpipename":       pipelineName,
		"cdfpipetype":       pipelineType,
		"cdfpipemodulename": pipelineModuleName,
		"cdfpipeenvname":    pipelineEnvName,
		"cdfpipebranch":     pipelineBranch,
	}

	bytes, err := json.Marshal(extension)
	if err != nil {
		log.Fatalf("failed to marsha extension, %v", err)
	}
	event.SetExtension("cdf", bytes)
}
