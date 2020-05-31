#!/bin/bash

# Copyright (c) 2016-2017 Bitnami
# Copyright (c) 2018 Cisco and/or its affiliates.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -xe

./scripts/install-kubectl.sh

kubectl get nodes -o wide
kubectl version
kubectl api-versions
kubectl label --overwrite --all=true nodes app=nsmgr-daemonset

kubectl apply -f k8s/conf/crd-networkservices.yaml
kubectl apply -f k8s/conf/crd-networkserviceendpoints.yaml
kubectl apply -f k8s/conf/crd-networkservicemanagers.yaml

# vim: sw=4 ts=4 et si
