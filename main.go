/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/util/wait"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/klog"
	"time"
	"volcano.sh/noderestopo/pkg/args"
	"volcano.sh/noderestopo/pkg/numatopo"
	"volcano.sh/noderestopo/pkg/client/clientset/versioned"
)

var logFlushFreq = pflag.Duration("log-flush-frequency", 5*time.Second, "Maximum number of seconds between log flushes")


func getNumaTopoClient(argument *args.Argument) (*versioned.Clientset, error){
	config, err := args.BuildConfig(argument.KubeClientOptions)
	if err != nil {
		return nil, err
	}

	return versioned.NewForConfigOrDie(config), err
}


func main() {
	klog.InitFlags(nil)

	opt := args.NewArgument()
	opt.AddFlags(pflag.CommandLine)
	cliflag.InitFlags()

	go wait.Until(klog.Flush, *logFlushFreq, wait.NeverStop)
	defer klog.Flush()

	nodeInfoClient, err := getNumaTopoClient(opt)
	if err != nil {
		klog.Errorf("get numainfo client failed, err = %v", err)
		return
	}

	for {
		isChg := numatopo.NodeInfoRefresh(opt)
		if isChg {
			klog.V(4).Infof("node info changes.")
			numatopo.CreateOrUpdateNumatopo(nodeInfoClient)
		}

		time.Sleep(opt.CheckInterval)
	}
}
