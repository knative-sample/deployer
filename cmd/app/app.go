package app

import (
	"strings"

	"github.com/golang/glog"
	"github.com/knative-sample/deployer/cmd/app/options"
	"github.com/knative-sample/deployer/pkg/deployer"
	"github.com/spf13/cobra"
)

// start edas api
func NewCommandStartServer(stopCh <-chan struct{}) *cobra.Command {
	ops := &options.Options{}
	mainCmd := &cobra.Command{
		Short: "Golang Sandbox ",
		Long:  "Alibaba Cloud Container Service Sandbox for Golang ",
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
	dp := deployer.Deployer{}
	if err := dp.Run(); err != nil {
		glog.Errorf("dp.Run error:%s", err)
	}
	<-stopCh
}
