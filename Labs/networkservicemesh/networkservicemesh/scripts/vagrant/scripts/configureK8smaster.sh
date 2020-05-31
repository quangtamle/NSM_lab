#!/bin/bash
# Get the IP address that VirtualBox has given this VM
IPADDR=$(ifconfig ens3 | grep -i Mask | awk '{print $2}'| cut -f2 -d:) #Change the name of the interface depend on your hosts
echo This VM has IP address "$IPADDR"

# Setup Hugepages
#echo "Copying /vagrant/10-kubeadm.conf to /etc/systemd/system/kubelet.service.d/10-kubeadm.conf"
#cp /vagrant/10-kubeadm.conf /etc/systemd/system/kubelet.service.d/10-kubeadm.conf

# Set up Kubernetes
NODENAME=$(hostname -s)
kubeadm init --apiserver-cert-extra-sans="$IPADDR" --apiserver-advertise-address="$IPADDR" --node-name "$NODENAME" --pod-network-cidr="10.32.0.0/12" #Change IP range according to your network

echo "KUBELET_EXTRA_ARGS= --node-ip=${IPADDR}" > /etc/default/kubelet
sudo service kubelet restart

# Set up admin creds for the vagrant user
echo Copying credentials to /home/vagrant...
sudo --user=vagrant mkdir -p /home/vagrant/.kube
cp -i /etc/kubernetes/admin.conf /home/vagrant/.kube/config
chown "$(id -u vagrant):$(id -g vagrant)" /home/vagrant/.kube/config

# Set up admin creds for the root user
echo Copying credentials to /root
mkdir -p /root/.kube
cp -i /etc/kubernetes/admin.conf /root/.kube/config

# Make credentials available outside of vagrant
echo Copying credentials out of vagrant
mkdir -p /vagrant/.kube/
cp /etc/kubernetes/admin.conf /vagrant/.kube/config

echo "Attempting kubectl version"
kubectl version

# Install networking
    kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')"
    
# kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/62e44c867a2846fefb68bd5f178daf4da3095ccb/Documentation/kube-flannel.yml
# kubectl apply -f https://docs.projectcalico.org/v3.7/manifests/calico.yaml

# Untaint master
echo "Untainting Master"
kubectl taint nodes --all node-role.kubernetes.io/master-

# Save the kubeadm join command with token
echo '#!/bin/sh' > /vagrant/scripts/kubeadm_join_cmd.sh
kubeadm token create --print-join-command >> /vagrant/scripts/kubeadm_join_cmd.sh
