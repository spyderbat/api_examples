#!/usr/bin/env python3

import sys
import json
from typing import Dict, List
import argparse
import fnmatch
from rich.console import Console
from rich.table import Table

def get_cluster(pod):
    return pod.get('cluster', 'cblatt-test-cluster1')

def get_deployments(pods, args):
    """ given set of pods, returns apps of interest, and
    the pods associated with them.
    We currently filter for pods in the cvtocloud
    namespace managed by replicasets. We infer the deployment
    name from the pod name and replicaset name.
    """
    arg_deployments=args.deployment.split(',')
    deployments = {}
    for pod in pods:
        metadata = pod['metadata']
        if (args.namespace in [metadata['namespace'], "all"] and
            metadata['ownerReferences'][0]['kind'] == 'ReplicaSet'):
            deployment = "-".join(metadata['name'].split('-')[:-2])
            if args.deployment == 'all' or deployment in arg_deployments:
                cluster_name = get_cluster(pod)
                if deployments.get(deployment) is None:
                    deployments[deployment] = {
                        cluster_name: [pod_summary(pod)]
                    }
                else:
                    deployments[deployment][cluster_name] = \
                        deployments[deployment].setdefault(cluster_name, []) +\
                        [pod_summary(pod)]
    return deployments


def pod_summary(pod):
    return {
        'pod_name': pod['metadata']['name'],
        'image_specs': [cont['image'] for cont in pod['spec']['containers']],
        'image_status': [cont['image'] for cont in pod['status']['containerStatuses']],
        'pod_status': pod['status']['phase'],
        'container_statuses': [{
            'ready': cont['ready'],
            'started': cont['started'],
            'state': cont['state']
            } for cont in pod['status']['containerStatuses']]
    }


def check_health(pod):
    pod_healthy = pod['pod_status'] == "Running"
    conts_healthy = [cs['ready'] and cs['started']for cs in pod['container_statuses']]
    return pod_healthy, conts_healthy


def check_image(pod, pattern):
    return [fnmatch.fnmatch(image, pattern) for image in pod['image_status']]

def check_deployments(deployments, image_pattern):
    for deploy in deployments.values():
        for cluster in deploy.values():
            for pod in cluster:
                pod['healthy'], pod['conts_healthy'] = check_health(pod)
                pod['image_matches'] = check_image(pod, image_pattern)
    return deployments

def process(pods_in: Dict, args) -> Dict:
    if type(pods_in) == dict:
        pods = pods_in['items']
    else: #list
        pods = []
        for entry in pods_in:
            pods.extend(entry["items"])
    deployments = get_deployments(pods, args)
    deployments = check_deployments(deployments, args.check_image)
    return deployments


def render_result(result, args):

    if args.out_format == 'json':
        print(json.dumps(result, indent=2))
        return

    if args.out_format == 'table':
        render_table(result, args.errors_only)

    if args.out_format == 'txt':
        for deployment_name, deployment in result.items():
            print(f'Deployment {deployment_name}')
            render_deploy(deployment)
            print('')

def render_deploy(deploy):
    print(f'    is deployed on clusters:')
    for cluster, pods in deploy.items():
        print(f'    - Cluster {cluster}')
        if (pods):
            print(f'        On pods:')
            for pod in pods:
                cont_st = str([list(st['state'].keys())[0] for st in pod["container_statuses"]])
                print(f"        - {pod['pod_name']}")
                print(f"            Running image(s): {pod['image_status']}")
                print(f"            Pod status: {pod['pod_status']}")
                print(f"            Container status(es): {cont_st}")
        else:
            print(f'     No pods found')


def emphasize(text):
    return f"[bold red]{text}[/bold red]"


def render_table(result, errors_only):
    table = Table()
    table.add_column("Deployment")
    table.add_column("Cluster")
    table.add_column("Pod")
    table.add_column("Image")
    table.add_column("PodStatus")
    table.add_column("ContainerStatus")

    previous = {
        "deployment_name": "None",
        "cluster": "None",
        }
    error_count = 0
    for deployment_name, deployment in result.items():
        for cluster, pods in deployment.items():
            for pod in pods:
                for image_st, cont_st, cont_healthy, image_matches in zip(
                        pod['image_status'],
                        pod['container_statuses'],
                        pod['conts_healthy'],
                        pod['image_matches']
                        ):
                    error = not (pod['healthy'] and cont_healthy and image_matches)
                    if error:
                        error_count += 1
                    if errors_only and not error:
                        continue

                    cont_ready = 'ready' if cont_st['ready'] else 'not ready'
                    cont_started = 'started' if cont_st['started'] else 'not started'
                    cont_info = "|".join([cont_started, cont_ready])
                    deploy_render = deployment_name if deployment_name != previous['deployment_name'] else ""
                    cluster_render = cluster if (deployment_name+cluster !=
                                                 previous['deployment_name']+previous['cluster']) else ""
                    pod_render = pod['pod_name']
                    style = "red" if error else "green"
                    if not pod['healthy']:
                        pod_render = emphasize(pod_render)
                    if not cont_healthy:
                        cont_info = emphasize(cont_info)
                    if not image_matches:
                        image_st = emphasize(image_st)
                    table.add_row(deploy_render, cluster_render, pod_render, image_st, pod['pod_status'], cont_info, style=style)
                    previous['deployment_name'] = deployment_name
                    previous['cluster'] = cluster

    if errors_only and error_count == 0:
        print('no problems encountered')
    else:
        console = Console()
        console.print(table)
        print(f"{error_count} problems encountered")


def _parse_cmd_line():
    parser = argparse.ArgumentParser(
        description='Check status of deployment rollouts across clusters',
        formatter_class=argparse.RawDescriptionHelpFormatter
    )
    parser.add_argument('-i', '--input_file', help='report input, as collected from running "spyctl get pods -o json"')
    parser.add_argument('-d', '--deployment', default='all', help='filter for specific deployments (delimit with comma if more than one). Default is "all"')
    parser.add_argument('-n', '--namespace', default='all', help='filter for a specific namespace, use "all" for all namespaces. Default is "cvtocloud"')
    parser.add_argument('-o', '--out_format', choices=['txt', 'json', 'table'], default='table')
    parser.add_argument('-c', '--check_image', default='*', help='image pattern to check for (support * and ? glob-style pattern matching). Default is *')
    parser.add_argument('-e', '--errors_only', dest='errors_only', action='store_true', help='show errors only')

    args = parser.parse_args()
    return args

if __name__ == "__main__":
    if len(sys.argv) == 1:
        infile = sys.stdin
        args = argparse.Namespace(
            deployment= "all",
            namespace="all",
            out_format="table",
            check_image="*",
            errors_only=False
            )
    else:
        args = _parse_cmd_line()
        infile = open(args.input_file, 'rt') if args.input_file else sys.stdin

    with infile:
        pods = json.load(infile)
        report = process(pods, args)
        render_result(report, args)
