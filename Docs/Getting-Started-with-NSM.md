To get started, follow the [Quick Start](https://github.com/networkservicemesh/networkservicemesh/blob/master/docs/guide-quickstart.md) guide.

If you want a more detailed look, you can follow the [Build page](https://github.com/networkservicemesh/networkservicemesh/blob/master/docs/guide-build.md).

In this part, There is a detailed instruction of deploying NSM infrastructure in OpenStack.

# Guide to install Network Service Mesh on OpenStack:

You need at least one OpenStack instance (VM), two is recommended. The operation system on these instances can be any Linux distributions.

**Attention !!!** Before you run any scripts in the below, you need to take a look into that because there are some variables varies from instance to instance.
I put all the detailed comments into those scripts, you just need to change them to adapt to your system.

Step 1: Install Docker on these instances via [install_docker script](https://github.com/quangtamle/NSM_lab/blob/master/Labs/networkservicemesh/networkservicemesh/scripts/openstack/install_docker.sh)

Step 2: Install Kubernetes on these instances via [install_kubernetes script](https://github.com/quangtamle/NSM_lab/blob/master/Labs/networkservicemesh/networkservicemesh/scripts/openstack/install_kubernetes.sh)

Step 3: You need to decide which instance is the master node and the other will be the worker ones.
* For the master node, you run [configureK8smaster script](https://github.com/quangtamle/NSM_lab/blob/master/Labs/networkservicemesh/networkservicemesh/scripts/openstack/configureK8smaster.sh). You need to change the interface network name to match to your instance and you can choose whatever Container Network Interfaces (CNI which is mandatory when install K8s) such as calico, weave, flannel,...
* For the worker nodes, you run [configureK8sworker script](https://github.com/quangtamle/NSM_lab/blob/master/Labs/networkservicemesh/networkservicemesh/scripts/openstack/configureK8sworker.sh).

Now you have a K8s system. The next step is to deployed NSM infrastructure on this K8s system.
Step 4: NSM infrastructure is deployed by the make machinery or by Helm (you need to install [Helm](https://helm.sh/docs/using_helm/#installing-helm) before)
* "make k8s-infra-deploy" will deploy the NSM infrastructure such as NSMgr, NSM dataplane, NSM monitor tool (Skydive, jaeger,...)
  
Now you have NSM system installed in your K8s system


