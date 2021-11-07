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
API_KEY = os.environ.get("API_KEY")


#
# Get the ORG_UID from a URL in the spyderbat console.
# For instance, if you visit
#    https://www.tigerbat.com/app/org/taQ5dvsTc81oS8U0kQvi/sources
# the org is the compoenent after /org/
#
ORG_UID = os.environ.get("ORG_UID")


#
# You can probably leave this alone.
#
API_URL = os.environ.get("API_URL", "https://api.spyderbat.com")


#
# If there's an error in the API call, an ApiError will be raised
#
class ApiError(Exception):
    pass


# -----------------------------------------------------------------------
def create_source(api_key, org_uid, source_uid, description, api_url=API_URL):
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
        "description": description,
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
def query(api_key, org_uid, source_uid, start, end,
              datatype="spydergraph", api_url=API_URL):
    """Retrieve data from Spyderbat in the specified organization.

    Args:
        api_key:    a string containing your API key for spyderbat
        org_uid:    a string containing the organization UID in spyderbat
        source_uid: a string containing the UID of the data source
        start:      start time (epoch time) for the fetch
        end:        end time (epoch time) for the fetch
        datatype:   optional datatype.  Leave this alone unless you
                    know what you're doing.
        api_url:    an optional base URL that locates the Spyrderbat API

    Returns:
        Data fetched from API
    """

    headers = {
        "Authorization": f"Bearer {api_key}",
        "Accept": "application/json"
    }
    url = f"{api_url}/api/v1/org/{org_uid}/data/?src={source_uid}&st={start}&et={end}&dt={datatype}"
    r = requests.get(url, headers=headers)
    if r.status_code == 200:
        data = r.text.split('\n')
        return [json.loads(x) for x in data]
    else:
        raise ApiError(r.status_code)

# -----------------------------------------------------------------------
def send_data(api_key, org_uid, source_uid, data, api_url=API_URL):
    """Sends data to Spyderbat analytic in the specified organization.
    Args:
        api_key:    a string containing your API key for spyderbat
        org_uid:    a string containing the organization UID in spyderbat
        source_uid: a string containing the UID of the data source
        data:       one or more JSON serializable objects to send
                    use a list to send multiple objects
        api_url:    an optional base URL that locates the Spyrderbat API
    Returns:
        None
    """

    if type(data) is not list:
        data = [data]
    data = [json.dumps(x) if type(x) is not str else x for x in data]
    data = '\n'.join(data).encode('ascii')
    data = gzip.compress(data)
    headers = {
        "Authorization": f"Bearer {api_key}",
        "Content-Encoding": "gzip",
        "Content-Type": "application/ndjson"
    }
    url = f"{api_url}/api/v1/org/{org_uid}/source/{source_uid}/data/sb-agent"
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
    src_uid = 'bs_src_03'
    # Make sure source exists
    sources = list_sources(API_KEY, ORG_UID)
    found = False
    print("%20s %10s %s" % ('uid', 'name', 'description'))
    print("----------------------------------------------")
    for src in sources:
        print("%20s %10s %s" % (src['uid'], src['name'], src['description']))
        if src['uid'] == src_uid:
            found = True

    if not found:
        print("Creating source %s" % src_uid)
        create_source(API_KEY, ORG_UID, src_uid, "My Third Source")
    else:
        print("Source %s already exists" % src_uid)

    intel = {
      "schema": "threat_intel",
      "id": uuid.uuid4().hex,
      "time": time.time(),
      "list_name": "Dynamic C&C",
      "key_field": "remote_ip",
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
    }
    send_data(API_KEY, ORG_UID, src_uid, intel)
