package argocd

import (
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/eclipse-paho/paho/autopaho"
)

func notifyChangesMqtt(app *v1alpha1.Application, wbc *WriteBackConfig, changeList []ChangeEntry) error {
	autopaho.ClientConfig{}
	return nil
}
