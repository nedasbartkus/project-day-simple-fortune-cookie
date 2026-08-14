#!/bin/bash                                                                                                                
  echo "$KUBECONFIG" | base64 -d > kubeconfig
  kubectl --kubeconfig kubeconfig apply --dry-run=client --validate=false -f k8s/