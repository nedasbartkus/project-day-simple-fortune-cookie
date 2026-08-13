#!/bin/bash                                                                                                                
  echo "$KUBECONFIG" | base64 -d > kubeconfig
  kubectl --kubeconfig kubeconfig apply -f k8s/ 