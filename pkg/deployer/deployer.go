package deployer

import (
	"github.com/golang/glog"
	"github.com/knative-sample/deployer/pkg/utils/kube"
	"github.com/knative/serving/pkg/apis/serving/v1alpha1"
	"github.com/knative/serving/pkg/apis/serving/v1beta1"
	servingclientset "github.com/knative/serving/pkg/client/clientset/versioned"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
	"fmt"
)

type Deployer struct {
	Image       string
	Namespace   string
	ServiceName string
	Port        string
}

func (dp *Deployer) Run() error {
	cfg, err := kube.GetKubeconfig()
	if err != nil {
		glog.Errorf("get kubeconfig error:%s ", err)
		return err
	}

	servingClient, err := servingclientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building Serving clientset: %v", err)
	}

	if svc, err := servingClient.ServingV1alpha1().Services(dp.Namespace).Get(dp.ServiceName, metav1.GetOptions{}); err != nil {
		// The Build resource may not exist.
		if !errors.IsNotFound(err) {
			glog.Errorf("get Serving %s/%s error:%s ", dp.Namespace, dp.ServiceName, err.Error())
			return err
		}

		// create Serving
		newSvc := &v1alpha1.Service{}
		newSvc.Namespace = dp.Namespace
		newSvc.Name = dp.ServiceName
		newSvc.Spec.Template = &v1alpha1.RevisionTemplateSpec{
			Spec: v1alpha1.RevisionSpec{
				RevisionSpec: v1beta1.RevisionSpec{
					PodSpec: v1beta1.PodSpec{
						Containers: []corev1.Container{
							{
								Image:           dp.Image,
								ImagePullPolicy: corev1.PullAlways,
							},
						},
					},
				},
			},
		}
		if _, err := servingClient.ServingV1alpha1().Services(dp.Namespace).Create(newSvc); err != nil {
			glog.Errorf("create serving: %s/%s error:%s", dp.Namespace, dp.ServiceName, err.Error())
			return err
		}
	} else {
		// Update Serving
		svc.Spec.Template.Annotations["updated"] = fmt.Sprintf("%v", time.Now().Unix())
		svc.Spec.Template.Spec.Containers[0].Image = dp.Image
		if _, err := servingClient.ServingV1alpha1().Services(dp.Namespace).Update(svc); err != nil {
			glog.Errorf("create serving: %s/%s error:%s", dp.Namespace, dp.ServiceName, err.Error())
			return err
		}
	}

	return nil
}
