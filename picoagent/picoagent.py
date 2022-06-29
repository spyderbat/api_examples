from genericpath import exists
import os
import json
import argparse
import subprocess as sp
from threading import Thread
import requests
import gzip
from uuid import uuid4
from base64 import urlsafe_b64encode

DESCRIPTION = "Runs bats like the nano agent, but even smaller!"
APIKEY = "API Key"
ORGID = "Org ID"
APIURL = "URL"
SOURCE = "Source"
SRCDESC = "Source Desc"
PREFIX = "Prefix"
BATS = "Bats"
MUID = "MUID"

REQ_CONFIG = [APIKEY, ORGID, MUID]
REQ_FIELDS = ["schema", "time"]

global config

def parse_args():
    parser = argparse.ArgumentParser(description=DESCRIPTION)
    parser.add_argument('-c', '--config', required=True, help="The configuration file location to use")
    parser.add_argument('-u', '--url', help="A custom API url to use")
    parser.add_argument('-p', '--prefix', help="A custom prefix to use for generating the source ID")
    parser.add_argument('-m', '--muid', help="A custom machine ID to use for analytics")
    parser.add_argument('-i', '--input', action=argparse.BooleanOptionalAction, help="Read bat data from standard input instead of sepcified bats")
    args = parser.parse_args()
    return args

def read_config(path):
    if not exists(path):
        raise ValueError("Config file did not exist")
    with open(path, 'r') as f:
        return json.loads(f.read())

def write_config(path):
    with open(path, 'w') as f:
        f.write(json.dumps(config))

def set_config(args):
    if args.url:
        config[APIURL] = args.url
    if args.prefix:
        config[PREFIX] = args.prefix
    if args.muid:
        config[MUID] = args.muid
    write_config(args.config)

def get_sources():
    headers = {"Authorization": f"Bearer {config[APIKEY]}"}
    url = f"{config[APIURL]}/api/v1/org/{config[ORGID]}/source/"
    r = requests.get(url, headers=headers)
    if r.status_code == 200:
        msg = r.json()
        return [src["uid"] for src in msg]
    else:
        raise ValueError(f"API request failed, status code {r.status_code}")

def rand_id():
    base = uuid4()
    bstr = urlsafe_b64encode(base.bytes[0:8]).strip(b'=')
    return bstr.decode('utf-8')

def make_source(sources, path):
    src = config[PREFIX] + rand_id()
    while src in sources:
        src = config[PREFIX] + rand_id()
    config[SOURCE] = src
    write_config(path)
    
    headers = {"Authorization": f"Bearer {config[APIKEY]}"}
    content = {
        "uid": config[SOURCE],
        "description": config[SRCDESC],
        "org_uid": config[ORGID]
    }
    url = f"{config[APIURL]}/api/v1/org/{config[ORGID]}/source/"
    r = requests.post(url, headers=headers, json=content)
    if r.status_code != 200:
        raise ValueError(f"API request failed, status code {r.status_code}")

def validate_source(sources):
    if config[SOURCE] not in sources:
        raise ValueError("Source specified does not exist")
    if not config[SOURCE].startswith(f"{config[PREFIX]}:"):
        raise ValueError("Source prefix did not match source")

def validate_config(run_bats, path):
    for req in REQ_CONFIG:
        if not config[req] or not config[req]:
            raise ValueError(f"Config was missing required field: {req}")
    # also validates API config immediately since it makes a query
    sources = get_sources()
    if not config[SOURCE]:
        make_source(sources, path)
    else:
        validate_source(sources)
    if run_bats and not config[BATS]:
        raise ValueError("No bats were specified to be run")

def send_data(data):
    try:
        data = json.loads(data)
    except json.decoder.JSONDecodeError as e:
        raise ValueError("Received data was not valid JSON") from e
    for field in REQ_FIELDS:
        if not field in data or not data[field]:
            raise ValueError(f"Received data was missing field {field}")
    if not "id" in data or not data["id"]:
        data["id"] = str(uuid4())
    if not "muid" in data or not data["muid"]:
        data["muid"] = config[MUID]
    # could add more data validation here
    data = json.dumps(data).encode('ascii')
    # TODO: batch data sending
    data = gzip.compress(data)
    headers = {
        "Authorization": f"Bearer {config[APIKEY]}",
        "Content-Encoding": "gzip",
        "Content-Type": "application/ndjson"
    }
    url = f"{config[APIURL]}/api/v1/org/{config[ORGID]}/source/{config[SOURCE]}/data/sb-agent"
    r = requests.post(url, headers=headers, data=data)
    if r.status_code != 200:
        raise RuntimeError(f"Failed to send data to source, status code: {r.status_code}")

def run_bat(bat):
    proc = sp.Popen(f"exec {bat}", shell=True, stdout=sp.PIPE)
    try:
        for out in proc.stdout:
            out = out.decode("utf-8")
            send_data(out.strip())
    except Exception as e:
        proc.kill()
        raise e
    except KeyboardInterrupt:
        exit()

if __name__ == '__main__':
    os.environ["PYTHONUNBUFFERED"] = "1"
    args = parse_args()
    config = read_config(args.config)
    set_config(args)
    validate_config(not args.input, args.config)
    try:
        if args.input:
            while True:
                send_data(input())
        else:
            threads = []
            for bat in config[BATS]:
                threads.append(Thread(target=run_bat, args=(bat,)))
            for thread in threads:
                thread.start()
            for thread in threads:
                thread.join()
    except KeyboardInterrupt:
        exit()
    