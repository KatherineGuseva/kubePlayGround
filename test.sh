#!/bin/bash
# -----------------------------------------------------------------------------
# IBM Confidential
# OCO Source Materials
# IBM Business Platform: sample-service
# (C) Copyright IBM Corp. 2017, 2018
#
# The source code for this program is not published or otherwise divested of
# its trade secrets, irrespective of what has been deposited with the U.S.
# Copyright Office.
# -----------------------------------------------------------------------------
set -e


echo "Checking up if cluster exists"
export cluster=$(kind get clusters | grep katya-cluster)
echo ${cluster}
if [ -z ${cluster} ];
then
    echo "Looks like cluster doesn't exist, creating..."
    kind create cluster --name=katya-cluster 
fi
kubectl config use-context kind-katya-cluster
kubectl get ns
#docker build -t katya-test -f  Dockerfile.dev .  