package cmd

import (
	"context"
	"encoding/json"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/spf13/cobra"
	"log"
	"time"
)



var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Emit Issues related Events",
	Long:  `Emit Issues related CloudEvent`,

}

var(
	issueId         string
	issueTitle      string
	issueAuthor     string
	issueRepository string
	issueData       map[string]string
	issueComment    string
)

func init() {
	rootCmd.AddCommand(issueCmd)
	issueCmd.AddCommand(issueCreatedCmd)
	issueCmd.AddCommand(issueUpdatedCmd)
	issueCmd.AddCommand(issueCommentedCmd)
	issueCmd.AddCommand(issueClosedCmd)

	issueCmd.PersistentFlags().StringVarP(&issueId, "id", "i", "", "Issue's id")
	issueCmd.PersistentFlags().StringVarP(&issueTitle, "title", "t", "", "Issue's title")
	issueCmd.PersistentFlags().StringVarP(&issueAuthor, "author", "a", "", "Issue's author")
	issueCmd.PersistentFlags().StringVarP(&issueRepository, "repository", "r", "", "Issue's repository")
	issueCmd.PersistentFlags().StringToStringVarP(&issueData, "data", "d", map[string]string{}, "Issue's Data")

	issueCommentedCmd.Flags().StringVarP(&issueComment, "comment", "c", "", "Issue's Comment")
}





var issueCreatedCmd = &cobra.Command{
	Use:   "created",
	Short: "Emit Issue Created Event",
	Long:  `Emit Issue Created CloudEvent`,
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
		event.SetType("CDF.Issue.Created")
		event.SetTime(time.Now())
		setExtensionForIssueEvents(event)

		event.SetData(cloudevents.ApplicationJSON, issueData) //

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

func setExtensionForIssueEvents(event cloudevents.Event ) {
	event.SetExtension("cdfissueid", issueId)
	event.SetExtension("cdfissuerepo", issueRepository)
	event.SetExtension("cdfissuetitle", issueTitle)
	event.SetExtension("cdfissueauthor", issueAuthor)

	var extension = map[string]string{
		"cdfissueid":     issueId,
		"cdfissuerepo":   issueRepository,
		"cdfissuetitle":  issueTitle,
		"cdfissueauthor": issueAuthor,
	}

	bytes, err := json.Marshal(extension)
	if err != nil {
		log.Fatalf("failed to marsha extension, %v", err)
	}
	event.SetExtension("cdf", bytes)
}

var issueUpdatedCmd = &cobra.Command{
	Use:   "updated",
	Short: "Emit Issue Update Event",
	Long:  `Emit Issue Update CloudEvent`,
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
		event.SetType("CDF.Issue.Updated")
		event.SetTime(time.Now())

		setExtensionForIssueEvents(event)

		event.SetData(cloudevents.ApplicationJSON, issueData) //

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

var issueClosedCmd = &cobra.Command{
	Use:   "closed",
	Short: "Emit Issue Closed Event",
	Long:  `Emit Issue Closed CloudEvent`,
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
		event.SetType("CDF.Issue.Closed")
		event.SetTime(time.Now())

		setExtensionForIssueEvents(event)

		event.SetData(cloudevents.ApplicationJSON, issueData) //

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
var issueCommentedCmd = &cobra.Command{
	Use:   "commented",
	Short: "Emit Issue Commented Event",
	Long:  `Emit Issue Commented CloudEvent`,
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
		event.SetType("CDF.Issue.Commented")
		event.SetTime(time.Now())

		setExtensionForIssueEvents(event)

		event.SetData(cloudevents.ApplicationJSON, map[string]string{
			"comment": issueComment,
		})//

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