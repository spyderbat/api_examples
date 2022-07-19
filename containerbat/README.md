# Containerbat

## Description

The container bat is a collection prototype implementation that fetches docker information about 
docker container instances known to the docker API. 

This includes instances in all states, including terminated containers that have not been removed. 
This proptotype polls the information on a regular interval (currently 60 minutes)

## Dependencies

To run the containerbat, you must have the following:
- Python 3.7 or higher and the dependencies in requirements.txt (docker)
- The spyderbat picoagent installed and configured (as described in "Running the bat")

## Running the bat

First, download the picoagent from https://github.com/spyderbat/api_examples/tree/main/picoagent. Find the picoagent.py file and move it into the same directory as this file.
Assuming python is installed, run 'pip install requirements.txt'

To start up the containerbat, fill out the required configuration fields as shown below and the run the following command: 'python3 picoagent.py -c containerbat-config.json'. 

Optionally, use -v to output the results of the bat to stdout. By default, nothing will be outputted to stdout - this is normal, and the bat is still running correctly.

## Configuration Fields (Required)

To set the following fields, edit ec2config.json. Do not edit any other fields or the program may not function correctly.

`API Key`: Your Spyderbat API key (see below for where to find this) \
`Org ID`: Your Spyderbat organization ID (see below for where to find this) \

## How to find your Spyderbat API Key

In the spyderbat console, click your profile picture, then click "API keys". 
If none are listed, click "Create API key", then click the copy icon under "Key" and paste it into ec2config.json
If you already have one or more API keys, select one that's listed as "Active" and copy and paste the key into ec2config.json as listed above

## How to find your Spyderbat organization ID

In the spyderbat console, go to the sources tab. The url should be listed as "https://spyderbat.net/app/org/XXXXXXXX/sources".
The value listed after /org/ is your Spyderbat organization ID