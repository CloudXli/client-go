/*
Copyright 2016 The Kubernetes Authors.
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

/*
cloudxli
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	c "kubernetes-client/client"
	MC "kubernetes-client/manage"
)

func main() {
	// get kubernetes client
	clientset := c.Getk8sclient()

	// list pods
	MC.ListPods(clientset)
	// start action
	deploymentNamespace := "liang-namespace01"
	// deploymentName := "use-client-go-demo-deployment"
	MC.ListReplicas(deploymentNamespace, clientset)
	// MC.CreateDeployment(deploymentNamespace)
	// DeleteDeploy(deploymentNamespace, deploymentName)

}
