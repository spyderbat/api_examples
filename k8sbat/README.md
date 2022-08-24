# Kubernetes bat

This is a very quick and dirty implementation of a kubernetes cluster bat. 
The bat needs to be installed on a host that has a kubectl installation, and has kubeconfigs
for all the clusters of interest. 

The bat will iterate over all the cluster contexts available in the kubectl config. 
(kubectl config get-contexts )

## Dependencies
The bat is dependent on 
- python3 (for the picoagent)
- bash
- jq  

## Usage
Enter your API and organization key in the k8s_bat-config.json

Run 
``` python picoagent.py -c k8s_bat-config-guy.json```

