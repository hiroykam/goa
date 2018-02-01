// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// calc gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design -o
// $(GOPATH)/src/goa.design/goa/examples/calc

package cli

import (
	"flag"
	"fmt"
	"os"

	goa "goa.design/goa"
	calcsvcc "goa.design/goa/examples/calc/gen/grpc/calc/client"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `calc add
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` --transport=grpc calc add --p '{
      "a": 546803495890724710,
      "b": 7837387407375911615
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		calcFlags = flag.NewFlagSet("calc", flag.ContinueOnError)

		calcAddFlags = flag.NewFlagSet("add", flag.ExitOnError)
		calcAddPFlag = calcAddFlags.String("p", "", "")
	)
	calcFlags.Usage = calcUsage
	calcAddFlags.Usage = calcAddUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if len(os.Args) < flag.NFlag()+3 {
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = os.Args[1+flag.NFlag()]
		switch svcn {
		case "calc":
			svcf = calcFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(os.Args[2+flag.NFlag():]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = os.Args[2+flag.NFlag()+svcf.NFlag()]
		switch svcn {
		case "calc":
			switch epn {
			case "add":
				epf = calcAddFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if len(os.Args) > 2+flag.NFlag()+svcf.NFlag() {
		if err := epf.Parse(os.Args[3+flag.NFlag()+svcf.NFlag():]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "calc":
			c := calcsvcc.NewClient(cc, opts...)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = calcsvcc.BuildAddPayload(*calcAddPFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// calcUsage displays the usage of the calc command and its subcommands.
func calcUsage() {
	fmt.Fprintf(os.Stderr, `The calc service performs operations on numbers
Usage:
    %s [globalflags] calc COMMAND [flags]

COMMAND:
    add: Adds the operands

Additional help:
    %s calc COMMAND --help
`, os.Args[0], os.Args[0])
}
func calcAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] calc add -p JSON

Adds the operands
    -p JSON: 

Example:
    `+os.Args[0]+` --transport=grpc calc add --p '{
      "a": 546803495890724710,
      "b": 7837387407375911615
   }'
`, os.Args[0])
}
