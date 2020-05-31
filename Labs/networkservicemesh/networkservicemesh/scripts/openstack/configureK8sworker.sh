#!/usr/bin/env bash

IPADDR=$(ifconfig eth1 | grep -i Mask | awk '{print $2}'| cut -f2 -d:) #Change the name of the interface depend on your hosts
echo This VM has IP address "$IPADDR"

# Joining K8s
bash /openstack/kubeadm_join_cmd.sh

echo "KUBELET_EXTRA_ARGS= --node-ip=${IPADDR}" > /etc/default/kubelet
service kubelet restart


