package commands

import (
	"fmt"
	"log"

	"github.com/Skarlso/go-furnace/utils"
	"github.com/Yitsushi/go-commander"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// Status command.
type Status struct {
}

// Execute defines what this command does.
func (c *Status) Execute(opts *commander.CommandHelper) {
	stackname := opts.Arg(0)
	if len(stackname) < 1 {
		log.Fatalln("A stackname to look for must be provided.")
	}

	sess := session.New(&aws.Config{Region: aws.String("eu-central-1")})
	cfClient := cloudformation.New(sess, nil)
	descResp, err := cfClient.DescribeStacks(&cloudformation.DescribeStacksInput{StackName: aws.String(stackname)})
	utils.CheckError(err)
	fmt.Println()
	log.Println("Stack state is: ", descResp.Stacks[0].GoString())

}

// NewStatus Creates a new Status command.
func NewStatus(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &Status{},
		Help: &commander.CommandDescriptor{
			Name:             "status",
			ShortDescription: "Status of a stack.",
			LongDescription:  `Get detailed status of the stack.`,
			Arguments:        "name",
			Examples:         []string{"status FurnaceStack"},
		},
	}
}