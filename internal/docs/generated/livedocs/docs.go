// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "mdtogo"; DO NOT EDIT.
package livedocs

var READMEShort = `Reconcile configuration files with the live state`
var READMELong = `
Tools to safely apply and remove packages from clusters.

| Command   | Description                                       |
|-----------|---------------------------------------------------|
| [apply]   | apply a package to the cluster                    |
| [preview] | preview the operations that apply will perform    |
| [destroy] | remove the package from the cluster               |

**Data Flow**: local configuration or stdin -> kpt live -> apiserver (Kubernetes cluster)

| Configuration Read From | Configuration Written To |
|-------------------------|--------------------------|
| local files or stdin    | apiserver                |
| apiserver               | stdout                   |

#### Pruning
kpt live apply will automatically delete resources which have been
previously applied, but which are no longer included. This clean-up
functionality is called pruning. For example, consider a package
which has been applied with the following three resources:

	service-1 (Service)
	deployment-1 (Deployment)
	config-map-1 (ConfigMap)

Then imagine the package is updated to contain the following resources,
including a new ConfigMap named ` + "`" + `config-map-2` + "`" + ` (Notice that ` + "`" + `config-map-1` + "`" + `
is not part of the updated package):

	service-1 (Service)
	deployment-1 (Deployment)
	config-map-2 (ConfigMap)

When the updated package is applied, ` + "`" + `config-map-1` + "`" + ` is automatically
deleted (pruned) since it is omitted.


In order to take advantage of this automatic clean-up, a package must contain
a **grouping object template**, which is a ConfigMap with a special label. An example is:

	apiVersion: v1
	kind: ConfigMap
	metadata:
	  name: test-grouping-object
	  labels:
	    cli-utils.sigs.k8s.io/inventory-id: test-group

And the special label is:

	cli-utils.sigs.k8s.io/inventory-id: *group-name*

` + "`" + `kpt live apply` + "`" + ` recognizes this template from the special label, and based
on this kpt will create new grouping object with the metadata of all applied
objects in the ConfigMap's data field. Subsequent ` + "`" + `kpt live apply` + "`" + ` commands can
then query the grouping object, and calculate the omitted objects, cleaning up
accordingly. When a grouping object is created in the cluster, a hash suffix
is added to the name. Example:

	test-grouping-object-17b4dba8

#### Status
kpt live apply also has support for computing status for resources. This is 
useful during apply for making sure that not only are the set of resources applied
into the cluster, but also that the desired state expressed in the resource are
fully reconciled in the cluster. An example of this could be applying a deployment. Without
looking at the status, the operation would be reported as successful as soon as the
deployment resource has been created in the apiserver. With status, kpt live apply will
wait until the desired number of pods have been created and become available.

Status is computed through a set of rules for specific types, and
functionality for polling a set of resources and computing the aggregate status
for the set. For CRDs, there is a set of recommendations that if followed, will allow
kpt live apply to correctly compute status.

###
[tutorial-script]: ../gifs/live.sh
[apply]: apply.md
[preview]: preview.md
[destroy]: destroy.md`

var ApplyShort = `apply a package to the cluster`
var ApplyLong = `
    kpt live apply [FILENAME... | DIRECTORY] [flags]

The apply command creates, updates or deletes any resources in the cluster to
make the state of resources in the cluster match the desired state as specified
through the set of manifests. This command is similar to the apply command
available in kubectl, but also has support for pruning and waiting until all
resources has been fully reconciled.

Args:
  NONE:
    Input will be read from StdIn. Exactly one ConfigMap manifest
    with the grouping object annotation must be present.

  FILENAME:
    A set of files that contains k8s manifests. Exactly one of them
    needs to be a ConfigMap with the grouping object annotation.
    
  DIRECTORY:
    One or more directories that contain k8s manifests. The directories 
    must contain exactly one ConfigMap with the grouping object annotation.
    
Flags:
  no-prune:
    If true, previously applied objects will not be pruned.
    
  wait-for-reconcile:
    If true, after all resources have been applied, the cluster will
    be polled until either all resources have been fully reconciled
    or the timeout is reached.
    
  wait-polling-period:
    The frequency with which the cluster will be polled to determine 
    the status of the applied resources. The default value is every 2 seconds.
    
  wait-timeout:
    The threshold for how long to wait for all resources to reconcile before
    giving up. The default value is 1 minute.`

var DestroyShort = `remove a package from the cluster`
var DestroyLong = `
    kpt live destroy [FILENAME... | DIRECTORY] [flags]

The destroy command removes all files belonging to a package from
the cluster.

Args:
  NONE:
    Input will be read from StdIn. Exactly one ConfigMap manifest
    with the grouping object annotation must be present.

  FILENAME:
    A set of files that contains k8s manifests. Exactly one of them
    needs to be a ConfigMap with the grouping object annotation.
    
  DIRECTORY:
    One or more directories that contain k8s manifests. The directories 
    must contain exactly one ConfigMap with the grouping object annotation.`

var PreviewShort = `preview shows the changes apply will make against the live state of the cluster`
var PreviewLong = `
    kpt live preview [FILENAME... | DIRECTORY] [flags]

The preview command will run through the same steps as apply, but 
it will only print what would happen when running apply against the current
live cluster state. 

Args:
  NONE:
    Input will be read from StdIn. Exactly one ConfigMap manifest
    with the grouping object annotation must be present.

  FILENAME:
    A set of files that contains k8s manifests. Exactly one of them
    needs to be a ConfigMap with the grouping object annotation.
    
  DIRECTORY:
    One or more directories that contain k8s manifests. The directories 
    must contain exactly one ConfigMap with the grouping object annotation.
    
Flags:
  no-prune:
    If true, previously applied objects will not be pruned.`
