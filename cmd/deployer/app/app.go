package app

import (
	"strings"

	"os"

	"github.com/golang/glog"
	"github.com/knative-sample/deployer/cmd/deployer/app/options"
	"github.com/knative-sample/deployer/pkg/deployer"
	"github.com/spf13/cobra"
)

// start edas api
func NewCommandStartServer(stopCh <-chan struct{}) *cobra.Command {
	ops := &options.Options{}
	mainCmd := &cobra.Command{
		Short: "Knative Deployer",
		Long:  "Knative Deployer ",
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
	if ops.Image == "" {
		glog.Fatalf("--image is empty")
	}
	if ops.ServiceName == "" {
		glog.Fatalf("--service-name is empty")
	}

	ns := ops.Namespace
	if ns == "" {
		ns = os.Getenv("NAMESPACE")
		if ns == "" {
			glog.Fatalf("--namespace and NAMESPACE ENV is empty")
		}
	}

	dp := deployer.Deployer{
		Namespace:   ns,
		ServiceName: ops.ServiceName,
		Image:       ops.Image,
		Port:        ops.Port,
	}

	// run deployer
	dp.Run()

	<-stopCh
}
