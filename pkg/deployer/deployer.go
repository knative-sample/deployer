package deployer

import (
	"context"

	"encoding/json"

	"github.com/cloudevents/sdk-go"
	"github.com/golang/glog"
	"github.com/knative/eventing-sources/pkg/kncloudevents"
)

type Deployer struct {
}

func (dp *Deployer) Run() error {
	c, err := kncloudevents.NewDefaultClient()
	if err != nil {
		glog.Error("Failed to create client, ", err)
		return err
	}

	glog.Fatal(c.StartReceiver(context.Background(), dp.deploy))
	return nil
}

func (dp *Deployer) deploy(event cloudevents.Event) error {
	str, _ := json.Marshal(event)
	glog.Infof("cloudevents.Event\n%s", str)
	return nil
}
