package app

import (
	"strings"

	"github.com/golang/glog"
	"github.com/knative-sample/deployer/cmd/trigger/app/options"
	"github.com/knative-sample/deployer/pkg/trigger"
	"github.com/spf13/cobra"
)

// start edas api
func NewCommandStartServer(stopCh <-chan struct{}) *cobra.Command {
	ops := &options.Options{}
	mainCmd := &cobra.Command{
		Short: "Knative github trigger ",
		Long:  "Knative github trigger ",
		RunE: func(c *cobra.Command, args []string) error {
			glog.V(2).Infof("NewCommandStartServer main:%s", strings.Join(args, " "))
			run(stopCh, ops)
			return nil
		},
	}

	ops.SetOps(mainCmd)
	return mainCmd
}

func run(stopCh <-chan struct{}, ops *options.Options) {
	if ops.TriggerConfig == "" {
		glog.Fatalf("--trigger-config is empty")
	}

	dp := trigger.Trigger{
		TriggerConfig: ops.TriggerConfig,
	}
	if err := dp.Run(); err != nil {
		glog.Errorf("dp.Run error:%s", err)
	}
	<-stopCh
}
