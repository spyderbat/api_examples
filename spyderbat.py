# INTRODUCTION
#
# This python library is used to communicate with the Spyderbat API
#

import requests
import os
import json
import gzip

#
# API_KEY is the API Token for a specific organization
# Best practice is to store it in an environment variable.
#
API_KEY = os.environ.get("S_API_KEY")


#
# Get the ORG_UID from a URL in the spyderbat console.
# For instance, if you visit
#    https://www.tigerbat.com/app/org/taQ5dvsTc81oS8U0kQvi/sources
# the org is the compoenent after /org/
#
ORG_UID = "taQ5dvsTc81oS8U0kQvi"


#
# You can probably leave this alone.
#
API_URL = "https://api.tigerbat.com"


#
# If there's an error in the API call, an ApiError will be raised
#
class ApiError(Exception):
    pass


# -----------------------------------------------------------------------
def create_source(api_key, org_uid, source_uid, name, api_url=API_URL):
    """Creates a source in the specified organization.

    Args:
        api_key: a string containing your API key for spyderbat
        org_uid: a string containing the organization UID in spyderbat
        source_uid: a string containing the UID of the source to create.
        name: A descriptive string associated with the source
        api_url: an optional base URL that locates the Spyrderbat API

    Returns:
        A dictionary with a single key, "uid", mapping to source_uid
    """

    headers = {"Authorization": f"Bearer {api_key}"}
    content = {
        "uid": source_uid,
        "name": name,
        "org_uid": org_uid
    }
    url = f"{api_url}/api/v1/org/{org_uid}/source/"
    r = requests.post(url, headers=headers, json=content)
    msg = r.json()
    if r.status_code == 200:
        return msg
    else:
        raise ApiError(r.status_code, msg.get('msg', msg))


# -----------------------------------------------------------------------
def list_sources(api_key, org_uid, api_url=API_URL):
    """Returns a list of sources in the specified organization.

    Args:
        api_key: a string containing your API key for spyderbat
        org_uid: a string containing the organization UID in spyderbat
        api_url: an optional base URL that locates the Spyrderbat API

    Returns:
        A list of dictionary, one for each source
    """

    headers = {"Authorization": f"Bearer {api_key}"}
    url = f"{api_url}/api/v1/org/{org_uid}/source/"
    r = requests.get(url, headers=headers)
    msg = r.json()
    if r.status_code == 200:
        return msg
    else:
        raise ApiError(r.status_code, msg.get('msg', msg))


# -----------------------------------------------------------------------
def send_data(api_key, org_uid, source_uid, data,
              datatype="sb-agent", api_url=API_URL):
    """Sends data to Spyderbat analytic in the specified organization.

    Args:
        api_key:    a string containing your API key for spyderbat
        org_uid:    a string containing the organization UID in spyderbat
        source_uid: a string containing the UID of the data source
        data:       one or more JSON serializable objects to send
        datatype:   optional datatype.  Leave this alone unless you
                    know what you're doing.
        api_url:    an optional base URL that locates the Spyrderbat API

    Returns:
        None
    """

    if type(data) is not list:
        data = [data]
    data = [json.dumps(x) for x in data]
    data = '\n'.join(data).encode('ascii')
    data = gzip.compress(data)
    headers = {
        "Authorization": f"Bearer {api_key}",
        "Content-Encoding": "gzip",
        "Content-Type": "application/ndjson"
    }
    url = f"{api_url}/api/v1/org/{org_uid}/source/{source_uid}/data/{datatype}"
    r = requests.post(url, headers=headers, data=data)
    if r.status_code == 200:
        return None
    else:
        raise ApiError(r.status_code)


# -----------------------------------------------------------------------
if __name__ == "__main__":
    import time
    import uuid

    # Demo code
    src_uid = 'bs_src_02'
    # Make sure source exists
    sources = list_sources(API_KEY, ORG_UID)
    found = False
    for src in sources:
        if src['uid'] == src_uid:
            found = True
            break

    if not found:
        create_source(API_KEY, ORG_UID, src_uid, "Overwrite Source")

    intel = {
      "schema": "threat_intel",
      "name:": "FBI list",
      "time": time.time(),
      "data": {
        "24.54.161.13": {
          "severity": 0,
          "description": "Ian's Home Machine"
        },
        "68.203.10.2": {
          "severity": 0,
          "description": "Jon's Home Machine"
        },
        "68.187.47.233": {
          "severity": 0,
          "description": "Brian's Home Machine"
        }
      },
      "id": uuid.uuid4().hex
    }
    send_data(API_KEY, ORG_UID, src_uid, intel)
