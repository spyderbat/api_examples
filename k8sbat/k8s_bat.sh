#!/bin/bash

config=$(kubectl config view -o json)
contexts=$(echo $config | jq -r .contexts[].name)
#clusternames=$(echo $config | jq .contexts[].context.cluster)

schema="k8s:bat:1.0.0"

while True
    do
        for context in $contexts;
        do
            kubectl config use-context "$context" > /dev/null
            clusterid=$(echo $context | tr -d '/')
            k8s_version=$(kubectl version --short | grep 'Server Version' | cut -f 2 -d : | tr -d ' ')
            namespaces=$(kubectl get namespaces -A -o json | tr -d '\n' | tr -d ' ')
            nodes=$(kubectl get nodes -A -o json | tr -d '\n' | tr -d ' ')
            pods=$(kubectl get pods -A -o json | tr -d '\n' | tr -d ' ')
            replicasets=$(kubectl get replicasets -A -o json | tr -d '\n' | tr -d ' ')
            deployments=$(kubectl get deployments -A -o json | tr -d '\n' | tr -d ' ')
            daemonsets=$(kubectl get daemonsets -A -o json | tr -d '\n' | tr -d ' ')
            services=$(kubectl get services -A -o json | tr -d '\n' | tr -d ' ')
            endpoints=$(kubectl get endpoints -A -o json | tr -d '\n' | tr -d ' ')
            time=$(date +%s)
            echo "{\"time\": $time, \"orc_time\": $time,\"schema\": \"$schema\",\"name\": \"$context\", \"k8s_version\": \"$k8s_version\", \"clusterid\": \"$clusterid\", \"namespaces\": $namespaces, \"nodes\": $nodes, \"pods\": $pods, \"replicasets\": $replicasets, \"deployments\": $deployments, \"daemonsets\": $daemonsets, \"services\": $services, \"endpoints\": $endpoints}"
        done
    sleep 60
done
