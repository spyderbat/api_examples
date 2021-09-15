#!/usr/bin/env python3

import subprocess
from subprocess import PIPE
import shlex
import argparse
import os
import json
from ast import literal_eval
import requests

DESCRIPTION = """
Prototype for test-jig

Will launch a test command using the default information, 
track information to the spawned process, and make an api call
on the backend to pull/compare the results to see if they
match up with what we want.

"""

# -------------------------------------------------------------
# def _parse_cmd_line():
#     parser = argparse.ArgumentParser(
#         description=DESCRIPTION,
#         formatter_class=argparse.RawDescriptionHelpFormatter
#     )
#     parser.add_argument('-o', '--outfile', default=None)
#     parser.add_argument('-d', '--deployment', default='integration')
#     parser.add_argument('-O', '--org_uid', default=None, help="UID of org")
#     parser.add_argument('-m', '--muid', default=None, required=True, help="Machine UID")
#     parser.add_argument('-i', '--input', default=None, required=True, help="CommandToRUn")
#     parser.add_argument(
#         '-s', '--start', default=-3600,
#         help='Fetch chunks after this time; use negative for relative to now.  Default: -3600')
#     parser.add_argument(
#         '-e', '--end', default=None,
#         help='Fetch chunks up to this time; use negative for relative to now.  Default: now')
#     args = parser.parse_args()
#     return args

# def launch_command():
#     json.loads()
#     pass

class run_result:
    def __init__(self, PID, PPID, time_created, args) -> None:
        self.pid = PID
        self.ppid = PPID
        self.ctime = time_created
        self.args = args

def main():
    filename = "trace_out.txt"


    # ------ Run Commands -------
    commands = '''
        cp /usr/bin/wget ~
        cp ~/wget ~/Documents
    '''

    cmdstring = f"strace -f -ttt -o {filename} /bin/bash"
    cmd = shlex.split(cmdstring)
    proc = subprocess.run(cmd, input=commands, text=True)

    # ------ Parse Commands ------
    processes = []
    PPID = None
    with open(filename, 'r') as f:
        for line in f.readlines():
            if 'execve' in line and 'resumed' not in line:
                split = line.split(' ')
                PID = split[0]
                time_created = split[1]
                s = line.index('[')
                e = line.index(']') + 1
                args = literal_eval(line[s:e])
                if 'bin/bash' in line and PPID is None:
                    processes.append(run_result(PID, PPID, time_created, args))
                    PPID = PID
                else:
                    processes.append(run_result(PID, PPID, time_created, args))

    # ------ wait ------

    # ------ API Call -----
    
    api_fname = "apidata.json"

    org_uid = os.environ.get("I_ORG_UID"),
    api_key = os.environ.get("I_API_KEY"),
    with open('/opt/spyderbat/etc/muid', 'r') as f:
        muid = f.read()
    datatype = 'spydergraph'

    headers = {
        "Authorization": f"Bearer {api_key}",
        "Accept": "application/json"
    }

    with open(api_fname, 'w') as f:
        for p in processes:
            start = float(p.ctime) - 1
            end = start + 2
            url = f"https://api.kangaroobat.net/api/v1/org/{org_uid}/data/?src={muid}&st={start}&et={end}&dt={datatype}"
            resp = requests.get(url, headers=headers, stream=True)
            if resp.status_code == 200:
                for block in resp.iter_content(chunk_size=65536):
                    block = block.decode('ascii')
                    f.write(block)
            else:
                print ("Error fetching data: ", resp.status_code, resp.json()['msg'])


    # ------ Grab Results ------

    results = None
    with open(api_fname, 'r') as f:
        json.loads(f)

    # ------ Compare Results ------

    for proc in processes:
        print(proc.pid)


    
                

if __name__ == "__main__":
    main()