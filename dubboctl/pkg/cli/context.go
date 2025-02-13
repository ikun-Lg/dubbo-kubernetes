package cli

import (
	"github.com/apache/dubbo-kubernetes/operator/pkg/util/pointer"
	"github.com/apache/dubbo-kubernetes/pkg/kube"
	"k8s.io/client-go/rest"
)

type instance struct {
	clients map[string]kube.CLIClient
	RootFlags
}

type Context interface {
	CLIClient() (kube.CLIClient, error)
	CLIClientWithRevision(rev string) (kube.CLIClient, error)
	DubboNamespace() string
}

func NewCLIContext(rootFlags *RootFlags) Context {
	if rootFlags == nil {
		rootFlags = &RootFlags{
			kubeconfig:     pointer.Of[string](""),
			Context:        pointer.Of[string](""),
			namespace:      pointer.Of[string](""),
			dubboNamespace: pointer.Of[string](""),
		}
	}
	return &instance{
		RootFlags: *rootFlags,
	}
}

func (i *instance) CLIClient() (kube.CLIClient, error) {
	return i.CLIClientWithRevision("")
}

func (i *instance) CLIClientWithRevision(rev string) (kube.CLIClient, error) {
	if i.clients == nil {
		i.clients = make(map[string]kube.CLIClient)
	}
	if i.clients[rev] == nil {
		impersonationConfig := rest.ImpersonationConfig{}
		client, err := newKubeClientWithRevision(*i.kubeconfig, *i.Context, rev, impersonationConfig)
		if err != nil {
			return nil, err
		}
		i.clients[rev] = client
	}
	return i.clients[rev], nil
}

func newKubeClientWithRevision(kubeconfig, context, revision string, impersonationConfig rest.ImpersonationConfig) (kube.CLIClient, error) {
	drc, err := kube.DefaultRestConfig(kubeconfig, context, func(config *rest.Config) {
		config.QPS = 55
		config.Burst = 95
		config.Impersonate = impersonationConfig
	})
	if err != nil {
		return nil, err
	}
	return kube.NewCLIClient(kube.NewClientConfigForRestConfig(drc), kube.WithRevision(revision))
}
