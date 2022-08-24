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

            namespaces=$(kubectl get namespaces -A -o json | tr -d '\n')
            nodes=$(kubectl get nodes -A -o json | tr -d '\n')
            pods=$(kubectl get pods -A -o json | tr -d '\n')
            services=$(kubectl get services -A -o json | tr -d '\n')
            endpoints=$(kubectl get endpoints -A -o json | tr -d '\n')
            time=$(date +%s)
            echo "{\"time\": $time, \"orc_time\": $time,\"schema\": \"$schema\",\"name\": \"$context\", \"clusterid\": \"$context\", \"namespaces\": $namespaces, \"nodes\": $nodes, \"pods\": $pods, \"services\": $services, \"endpoints\": $endpoints}"
        done
    sleep 60
done
