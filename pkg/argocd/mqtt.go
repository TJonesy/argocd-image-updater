package argocd

import (
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"github.com/eclipse/paho.golang/paho/log"
	"github.com/eclipse/paho.golang/paho/session/state"
	storefile "github.com/eclipse/paho.golang/paho/store/file"
)

func notifyChangesMqtt(app *v1alpha1.Application, wbc *WriteBackConfig, changeList []ChangeEntry) error {
	autopaho.ClientConfig{}
	return nil
}
