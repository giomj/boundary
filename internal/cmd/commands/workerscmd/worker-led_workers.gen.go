// Code generated by "make cli"; DO NOT EDIT.
package workerscmd

import (
	"errors"
	"fmt"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/workers"
	"github.com/hashicorp/boundary/internal/cmd/base"
	"github.com/hashicorp/boundary/internal/cmd/common"
	"github.com/hashicorp/go-secure-stdlib/strutil"
	"github.com/mitchellh/cli"
	"github.com/posener/complete"
)

func initWorkerLedFlags() {
	flagsOnce.Do(func() {
		extraFlags := extraWorkerLedActionsFlagsMapFunc()
		for k, v := range extraFlags {
			flagsWorkerLedMap[k] = append(flagsWorkerLedMap[k], v...)
		}
	})
}

var (
	_ cli.Command             = (*WorkerLedCommand)(nil)
	_ cli.CommandAutocomplete = (*WorkerLedCommand)(nil)
)

type WorkerLedCommand struct {
	*base.Command

	Func string

	plural string

	extraWorkerLedCmdVars
}

func (c *WorkerLedCommand) AutocompleteArgs() complete.Predictor {
	initWorkerLedFlags()
	return complete.PredictAnything
}

func (c *WorkerLedCommand) AutocompleteFlags() complete.Flags {
	initWorkerLedFlags()
	return c.Flags().Completions()
}

func (c *WorkerLedCommand) Synopsis() string {
	if extra := extraWorkerLedSynopsisFunc(c); extra != "" {
		return extra
	}

	synopsisStr := "worker"

	synopsisStr = fmt.Sprintf("%s %s", "worker-led-type", synopsisStr)

	return common.SynopsisFunc(c.Func, synopsisStr)
}

func (c *WorkerLedCommand) Help() string {
	initWorkerLedFlags()

	var helpStr string
	helpMap := common.HelpMap("worker")

	switch c.Func {

	default:

		helpStr = c.extraWorkerLedHelpFunc(helpMap)

	}

	// Keep linter from complaining if we don't actually generate code using it
	_ = helpMap
	return helpStr
}

var flagsWorkerLedMap = map[string][]string{

	"create": {"scope-id", "name", "description"},
}

func (c *WorkerLedCommand) Flags() *base.FlagSets {
	if len(flagsWorkerLedMap[c.Func]) == 0 {
		return c.FlagSet(base.FlagSetNone)
	}

	set := c.FlagSet(base.FlagSetHTTP | base.FlagSetClient | base.FlagSetOutputFormat)
	f := set.NewFlagSet("Command Options")
	common.PopulateCommonFlags(c.Command, f, "worker-led-type worker", flagsWorkerLedMap, c.Func)

	extraWorkerLedFlagsFunc(c, set, f)

	return set
}

func (c *WorkerLedCommand) Run(args []string) int {
	initWorkerLedFlags()

	switch c.Func {
	case "":
		return cli.RunResultHelp

	case "update":
		return cli.RunResultHelp

	}

	c.plural = "worker-led-type worker"
	switch c.Func {
	case "list":
		c.plural = "worker-led-type workers"
	}

	f := c.Flags()

	if err := f.Parse(args); err != nil {
		c.PrintCliError(err)
		return base.CommandUserError
	}

	if strutil.StrListContains(flagsWorkerLedMap[c.Func], "id") && c.FlagId == "" {
		c.PrintCliError(errors.New("ID is required but not passed in via -id"))
		return base.CommandUserError
	}

	var opts []workers.Option

	if strutil.StrListContains(flagsWorkerLedMap[c.Func], "scope-id") {
		switch c.Func {

		case "create":
			if c.FlagScopeId == "" {
				c.PrintCliError(errors.New("Scope ID must be passed in via -scope-id or BOUNDARY_SCOPE_ID"))
				return base.CommandUserError
			}

		}
	}

	client, err := c.Client()
	if c.WrapperCleanupFunc != nil {
		defer func() {
			if err := c.WrapperCleanupFunc(); err != nil {
				c.PrintCliError(fmt.Errorf("Error cleaning kms wrapper: %w", err))
			}
		}()
	}
	if err != nil {
		c.PrintCliError(fmt.Errorf("Error creating API client: %w", err))
		return base.CommandCliError
	}
	workersClient := workers.NewClient(client)

	switch c.FlagName {
	case "":
	case "null":
		opts = append(opts, workers.DefaultName())
	default:
		opts = append(opts, workers.WithName(c.FlagName))
	}

	switch c.FlagDescription {
	case "":
	case "null":
		opts = append(opts, workers.DefaultDescription())
	default:
		opts = append(opts, workers.WithDescription(c.FlagDescription))
	}

	switch c.FlagRecursive {
	case true:
		opts = append(opts, workers.WithRecursive(true))
	}

	if c.FlagFilter != "" {
		opts = append(opts, workers.WithFilter(c.FlagFilter))
	}

	var version uint32

	if ok := extraWorkerLedFlagsHandlingFunc(c, f, &opts); !ok {
		return base.CommandUserError
	}

	var result api.GenericResult

	switch c.Func {

	}

	result, err = executeExtraWorkerLedActions(c, result, err, workersClient, version, opts)

	if err != nil {
		if apiErr := api.AsServerError(err); apiErr != nil {
			var opts []base.Option

			c.PrintApiError(apiErr, fmt.Sprintf("Error from controller when performing %s on %s", c.Func, c.plural), opts...)
			return base.CommandApiError
		}
		c.PrintCliError(fmt.Errorf("Error trying to %s %s: %s", c.Func, c.plural, err.Error()))
		return base.CommandCliError
	}

	output, err := printCustomWorkerLedActionOutput(c)
	if err != nil {
		c.PrintCliError(err)
		return base.CommandUserError
	}
	if output {
		return base.CommandSuccess
	}

	switch c.Func {

	}

	switch base.Format(c.UI) {
	case "table":
		c.UI.Output(printItemTable(result))

	case "json":
		if ok := c.PrintJsonItem(result); !ok {
			return base.CommandCliError
		}
	}

	return base.CommandSuccess
}

var (
	extraWorkerLedActionsFlagsMapFunc = func() map[string][]string { return nil }
	extraWorkerLedSynopsisFunc        = func(*WorkerLedCommand) string { return "" }
	extraWorkerLedFlagsFunc           = func(*WorkerLedCommand, *base.FlagSets, *base.FlagSet) {}
	extraWorkerLedFlagsHandlingFunc   = func(*WorkerLedCommand, *base.FlagSets, *[]workers.Option) bool { return true }
	executeExtraWorkerLedActions      = func(_ *WorkerLedCommand, inResult api.GenericResult, inErr error, _ *workers.Client, _ uint32, _ []workers.Option) (api.GenericResult, error) {
		return inResult, inErr
	}
	printCustomWorkerLedActionOutput = func(*WorkerLedCommand) (bool, error) { return false, nil }
)