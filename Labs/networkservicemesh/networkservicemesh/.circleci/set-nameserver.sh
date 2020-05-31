#!/bin/bash
#
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

cp /etc/resolv.conf resolv.conf

if [ "$1" == "sudo" ]; then
    echo "echo \"nameserver 147.75.69.23\" > /etc/resolv.conf" | sudo sh
    echo "cat resolv.conf >> /etc/resolv.conf" | sudo sh
else
    echo "nameserver 147.75.69.23" > /etc/resolv.conf
    cat resolv.conf >> /etc/resolv.conf
fi
