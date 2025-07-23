package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	"k8s.io/kubernetes/pkg/scheduler/framework"

	_ "github.com/gocrane/crane-scheduler/pkg/plugins/apis/config/scheme"

	"github.com/gocrane/crane-scheduler/pkg/plugins/dynamic"
	"github.com/gocrane/crane-scheduler/pkg/plugins/noderesourcetopology"
)

// Adapter function to convert our plugin factory to the expected type
func pluginFactoryAdapter(factoryFn func(runtime.Object, framework.Handle) (framework.Plugin, error)) func(context.Context, runtime.Object, framework.Handle) (framework.Plugin, error) {
	return func(ctx context.Context, args runtime.Object, handle framework.Handle) (framework.Plugin, error) {
		return factoryFn(args, handle)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	cmd := app.NewSchedulerCommand(
		app.WithPlugin(dynamic.Name, pluginFactoryAdapter(dynamic.NewDynamicScheduler)),
		app.WithPlugin(noderesourcetopology.Name, pluginFactoryAdapter(noderesourcetopology.New)),
	)

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
