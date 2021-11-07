#! /usr/bin/env python3.7
import os
import sys
import boto3
import json
from datetime import datetime

import spyderbat

"""
Expected environment variables:
    API_URL - optional; base url of spyderbat API (default is "https://api.spyderbat.com")
    API_KEY - Spyderbat API key
    ORG_UID - Spyderbat ORG Uid that receives the data
    SOURCE_UID - A uid for this source
    SOURCE_DESCRIPTION - A description for this source
"""

class DatetimeEncoder(json.JSONEncoder):
  def default(self, obj):
    if isinstance(obj, datetime):
        return obj.timestamp()
    elif isinstance(obj, date):
        return datetime.combine(obj, datetime.min.time()).timestamp()
    # Let the base class default method raise the TypeError
    return json.JSONEncoder.default(self, obj)    


def get_env(name):
    if name not in os.environ:
        sys.stderr.print(f"Error: {name} environment variable is not set.  Cannot continue")
        sys.exit(-1)
    return os.environ[name]

# Get all the env variables
api_key = get_env("API_KEY")
api_url = get_env("API_URL")
org_uid = get_env("ORG_UID")
source_description = get_env("SOURCE_DESCRIPTION")
source_uid = get_env("SOURCE_UID")

# Make sure the source exists
try:
    spyderbat.create_source(api_key, org_uid, source_uid, source_description, api_url=api_url)
except spyderbat.ApiError as exc:
    if exc.args[0] != 400 or exc.args[1] != 'resource with same uid or values already exists':
        print("Error creating source.  API returned ", exc)
        sys.exit(-1)

client = boto3.client('ec2')
resp = {'NextToken': ''}
while 'NextToken' in resp:
    resp = client.describe_instances(NextToken=resp['NextToken'])
    data = []
    for r in resp["Reservations"]:
        data.extend([json.dumps(i, cls=DatetimeEncoder) for i in r["Instances"]])
    import pdb; pdb.set_trace()
    spyderbat.send_data(api_key, org_uid, source_uid, data, api_url=api_url)
