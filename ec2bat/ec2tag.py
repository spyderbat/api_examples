#! /usr/bin/env python3
#############################################################################
## Bat to pull the tags from ec2 and list them by instance id.
## 
## Outputs a record with the field "tags", a dict where the key is the instance
## id and the value is a dict of tags.
##
## Author: Kit Smith Hanna <kit.smith.hanna@spyderbat.com>
## Dependencies: boto3
## Copyright Spyderbat May 24, 2022
#############################################################################

import argparse
from http.client import NON_AUTHORITATIVE_INFORMATION
import json
import sys
import time
import uuid
import boto3

BAT_SCHEMA='ec2metadata:bat:1.0.0'
frequency = 10

##
## Helpers
##
def get_uuid():
    """Get a compliant UUID, we prefer type 4 as they are easier to see and less likely for collisions"""
    return str(uuid.uuid4())


def get_monotonic_time():
    """Get the monotonic time,  this doesn't work correctly on macOS"""
    m = time.monotonic()
    return int(m * 1000000000)


def create_dummy_output(schema, muid):
    ret = {
        "schema": schema,
        "muid": muid,
        "id": get_uuid(),
        "monotonic_time": get_monotonic_time(),
        "time": time.time(),
        "orc_time": time.time(),
    }
    return ret


def __eprint(*args, **kwargs):
    """Spit out data to stderr"""
    print(*args, file=sys.stderr, **kwargs)


def __emit(*args, **kwargs):
    """Emit a normal healthy record"""
    print(*args, file=sys.stdout, **kwargs)


def emit_record(record):
    """Emit a record"""
    __emit(json.dumps(record, sort_keys=True), flush=True)



# MAIN #
if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Spyderbat Skeleton Bat')
    parser.add_argument('-m', '--muid', help="MUID to use", action="store", default='nomuid')
    parser.add_argument('-r', '--region', default='us-east-1')
    parser.add_argument('-k', '--key', help="Your AWS Access Key. Required")
    parser.add_argument('-sk', '--secretkey', help="Your AWS Secret Access Key. Required")
    parser.add_argument('-version', '--version', help="Outputs the version", action="store_true")
    args = parser.parse_args()

    # Main loop
    while True:
        ## Grab data from ec2
        if args.version:
            print("\nVersion: v1.0.0\n")
            sys.exit()
        if args.key and args.secretkey:
            ec2 = boto3.client(
                "ec2",
                aws_access_key_id = args.key,
                aws_secret_access_key =  args.secretkey,
                region_name=args.region)
        ## If a key pair was not provided in args, attempt to connect using environment variables or
        ## pulling from (~/.aws/credentials) 
        else:
            ec2 = boto3.client("ec2", region_name=args.region)
        rawdata = ec2.describe_instances()            

        ## Extract the tags and store in a dict by UUID
        tagsByUUID = {}
        for reservation in rawdata['Reservations']:
            instances = reservation['Instances']
            for instance in instances:
                instanceID = instance["InstanceId"]
                taglist = {}
                for tag in instance['Tags']:
                    taglist[tag["Key"]]=tag["Value"]
                tagsByUUID[instanceID] = taglist

        ## Create an output record
        record = create_dummy_output(BAT_SCHEMA, args.muid)

        ## Attach tags
        record["tags"] = tagsByUUID

        ## emit the output, then wait
        emit_record(record)
        time.sleep(frequency)