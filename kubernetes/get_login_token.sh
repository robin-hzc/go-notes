#!/bin/bash
# shellcheck disable=SC2046
kubectl describe secret -n=kube-system $(kubectl -n kube-system get secret | grep kube-proxy |awk '{print $1}')
