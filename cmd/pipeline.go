package cmd

import (
	"context"
	"encoding/json"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/spf13/cobra"
	"log"
	"time"
)

func init() {
	rootCmd.AddCommand(pipelineCmd)
	pipelineCmd.AddCommand(pipelineStartedCmd)
	pipelineCmd.AddCommand(pipelineFinishedCmd)
	pipelineCmd.AddCommand(pipelineFailedCmd)

	pipelineCmd.PersistentFlags().StringVarP(&pipelineId, "id", "i", "", "Pipeline's Id")
	pipelineCmd.PersistentFlags().StringVarP(&pipelineName, "name", "n", "", "Pipeline's Name")
	pipelineCmd.PersistentFlags().StringVarP(&pipelineRepoName, "repository", "r", "", "Pipeline's Repository")
	pipelineCmd.PersistentFlags().StringToStringVarP(&pipelineData, "data", "d", map[string]string{}, "Pipeline's Data")


}

var(
	pipelineId string
	pipelineName string
	pipelineRepoName string
	pipelineData map[string]string
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
		event.SetType("CDF.Pipeline.Started")
		event.SetTime(time.Now())

		setExtensionForPipelineEvents(event)

		event.SetData(cloudevents.ApplicationJSON, pipelineData)//

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


var pipelineFinishedCmd = &cobra.Command{
	Use:   "finished",
	Short: "Emit Pipeline Finished Event",
	Long:  `Emit Pipeline Finished CloudEvent`,
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
		event.SetType("CDF.Pipeline.Finished")
		event.SetTime(time.Now())

		setExtensionForPipelineEvents(event)

		event.SetData(cloudevents.ApplicationJSON, prData)//

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

var pipelineFailedCmd = &cobra.Command{
	Use:   "failed",
	Short: "Emit Pipeline Failed Event",
	Long:  `Emit Pipeline Failed CloudEvent`,
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
		event.SetType("CDF.Pipeline.Failed")
		event.SetTime(time.Now())

		setExtensionForPipelineEvents(event)

		event.SetData(cloudevents.ApplicationJSON, prData)//

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


func setExtensionForPipelineEvents(event cloudevents.Event ) {
	event.SetExtension("cdfpipelineid", pipelineId)
	event.SetExtension("cdfpipelinename", pipelineName)
	event.SetExtension("cdfpipelinerepository", pipelineRepoName)




	var extension = map[string]string{
		"cdfpipelineid": pipelineId,
		"cdfpipelinename":  pipelineName,
		"cdfpipelinerepository":  pipelineRepoName,

	}

	bytes, err := json.Marshal(extension)
	if err != nil {
		log.Fatalf("failed to marsha extension, %v", err)
	}
	event.SetExtension("cdf", bytes)
}