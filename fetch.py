#!/usr/bin/env python3

import sys
import os
import argparse
import time
import gzip
import math
import requests



DESCRIPTION = """
Download records from API service

The following environment variables are useful:
    I_ORG_UID -- your org_uid in integration
    I_API_KEY -- your api key in integration
    P_ORG_UID -- your org_uid in production
    P_API_KEY -- your api key in production
    D_ORG_UID -- your org_uid in dev
    D_API_KEY -- your api key in dev
    S_ORG_UID -- your org_uid in staging
    S_API_KEY -- your api key in staging

Start and end time can be a negative number, in seconds.
For example, '-s -300' means a start time 5 minutes ago.

If end time has a '+' prefix, it is added to start time.
For example, '-e +300' means 5 minutes after start time

"""

# -------------------------------------------------------------
def _parse_cmd_line():
    parser = argparse.ArgumentParser(
        description=DESCRIPTION,
        formatter_class=argparse.RawDescriptionHelpFormatter
    )
    parser.add_argument('-o', '--outfile', default=None)
    parser.add_argument('-d', '--deployment', default='integration')
    parser.add_argument('-O', '--org_uid', default=None, help="UID of org")
    parser.add_argument('-m', '--muid', default=None, required=True, help="Machine UID")
    parser.add_argument(
        '-s', '--start', default=-3600,
        help='Fetch chunks after this time; use negative for relative to now.  Default: -3600')
    parser.add_argument(
        '-e', '--end', default=None,
        help='Fetch chunks up to this time; use negative for relative to now.  Default: now')
    args = parser.parse_args()
    return args

DEPLOYMENTS = {
    'integration': {
       "ORG_UID": os.environ.get("I_ORG_UID"),
       "API_KEY": os.environ.get("I_API_KEY"),
       "API_URL": "https://api.kangaroobat.net",
    },
    'production': {
       "ORG_UID": os.environ.get("P_ORG_UID"),
       "API_KEY": os.environ.get("P_API_KEY"),
       "API_URL": "https://api.spyderbat.com",
    },
    'dev': {
       "ORG_UID": os.environ.get("D_ORG_UID"),
       "API_KEY": os.environ.get("D_API_KEY"),
       "API_URL": "https://api.gatorbat.net",
    },
    'staging': {
       "ORG_UID": os.environ.get("S_ORG_UID"),
       "API_KEY": os.environ.get("S_API_KEY"),
       "API_URL": "https://api.tigerbat.com",
    }
}
options = _parse_cmd_line()

if options.deployment not in DEPLOYMENTS:
    print(f"Unknown deployment {options.deployment}")
    print("Known deployments:")
    for d in DEPLOYMENTS:
        print("    ", d)
    raise SystemExit(1)

org_uid = options.org_uid or DEPLOYMENTS[options.deployment]["ORG_UID"]
api_key = DEPLOYMENTS[options.deployment]["API_KEY"]
api_url = DEPLOYMENTS[options.deployment]["API_URL"]
muid = options.muid

datatype="spydergraph"

if options.end:
    end = float(options.end)
    if end < 0:
        end = time.time() + end
else:
    end = time.time()

if options.start:
    start = float(options.start)
    if start < 0:
        start = time.time() + start
else:
    start = end - 3600          # 1 hour

start = math.floor(start)
end = math.ceil(end)

if options.end is not None and len(options.end) > 0 and options.end[0] == '+':
    end += start

if end <= start:
    sys.stderr.write("invalid start/end times\n")
    sys.exit(1)


if options.outfile is None:
    outfile = sys.stdout
elif options.outfile.endswith('gz'):
    outfile = gzip.open(options.outfile, 'wt')
else:
    outfile = open(options.outfile, 'wt')

headers = {
    "Authorization": f"Bearer {api_key}",
    "Accept": "application/json"
}

url = f"{api_url}/api/v1/org/{org_uid}/data/?src={muid}&st={start}&et={end}&dt={datatype}"
resp = requests.get(url, headers=headers, stream=True)
if resp.status_code == 200:
    for block in resp.iter_content(chunk_size=65536):
        block = block.decode('ascii')
        outfile.write(block)
else:
    print ("Error fetching data: ", resp.status_code, resp.json()['msg'])
