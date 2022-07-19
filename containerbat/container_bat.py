#! /usr/bin/env python3
#############################################################################
# Bat to pull docker containerinformation to support populating a model_container
# abstraction in the spyderbat engine and ui
# Outputs a record with the list of docker containers present on the system
# and the key salient attributes available
##
# Author: Guy Duchatelet <guy.duchatelet@spyderbat.com>
# Dependencies:
# Copyright Spyderbat 2022
#############################################################################

import argparse
import json
import sys
import time
import uuid
import docker
from docker.api import container

CONTAINERBAT_SCHEMA = 'container:bat:1.0.0'
DELAY = 60


##
# Helpers
##
def get_uuid():
    """Get a compliant UUID, we prefer type 4 as they are easier to see and less likely for collisions"""
    return str(uuid.uuid4())


def get_monotonic_time():
    """Get the monotonic time,  this doesn't work correctly on macOS"""
    m= time.monotonic()
    return int(m * 1000000000)


def __eprint(*args, **kwargs):
    """Spit out data to stderr"""
    print(*args, file=sys.stderr, **kwargs)


def __emit(*args, **kwargs):
    """Emit a normal healthy record"""
    print(*args, file=sys.stdout, **kwargs)
    # print(*args, **kwargs)

def emit_record(record):
    """Emit a record"""
    __emit(json.dumps(record, sort_keys=True), flush=True)


def create_record(containers: list) -> dict:
    """ creates the list of containers reported by the docker api
        and their key properties, and returns in a single record """
    cs: list = []
    for container in containers:
        # Container identity
        c={}
        c["container_id"]= container.id
        c["container_short_id"] = container.short_id
        c["container_name"] = container.attrs["Name"]

        # Container runtime info
        c["platform"] = container.attrs["Platform"] 
        c["created"] = container.attrs['Created']
        c["path"] = container.attrs['Path']
        c["args"] = container.attrs['Args']
        c["cmd"] = container.attrs['Config']['Cmd']
        c["entrypoint"] = container.attrs['Config']['Entrypoint']
        c["container_state"]= container.attrs["State"]["Status"]
        c["container_detail_state"] = container.attrs["State"]
        c["env"] = container.attrs["Config"]["Env"]

        #Container image info
        c["image_id"] = container.image.id
        c["image_short_id"] = container.image.short_id
        c["image_name"] = container.attrs["Config"]["Image"]
        c["image_tags"] = container.image.tags

        #Resources used and exposed by the container
        c["exposed_ports"] = container.attrs["Config"].get("ExposedPorts", []) 
        c["port_bindings"] = container.ports  #provides bindings of container ports to host ports
        c["mounts"] = container.attrs["Mounts"]
        c["volumes"] = container.attrs["Config"]["Volumes"]
        c["hostname"] = container.attrs["Config"]["Hostname"]
        c["Domainname"] = container.attrs["Config"]["Domainname"]
        c["User"] = container.attrs["Config"]["User"]
        c["networks"] = container.attrs["NetworkSettings"]["Networks"]
        cs.append(c)

    rv= {}
    # Spyderbat fields and envelope
    rv["schema"]= CONTAINERBAT_SCHEMA
    rv["monotonic_time"]= get_monotonic_time()
    rv["time"]= time.time()
    rv["orc_time"]= time.time()
    rv["containers"]= cs
    return rv


def main(args):
    try:
        client= docker.from_env()
    except:
        print("Could not connect to docker api using environment sittings - check whether your docker environment variables are set")

    while True:
        clist= client.containers.list(all=True)
        record = create_record(clist)
        emit_record(record)
        time.sleep(DELAY)



# MAIN #
if __name__ == '__main__':
    parser= argparse.ArgumentParser(description='Spyderbat Container Bat')
    args= parser.parse_args()
    main(args)
