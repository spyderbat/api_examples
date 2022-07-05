# Ec2bat user guide

### Description

The EC2 bat fetches metadata tags from amazon's ec2 service and sends them to the analytics engine, where they're automatically paired with machines displayed on the spydertrace
It outputs the tags to analytics as a series of json records with the following format:
`{instanceID: [{key1:val1}, {key2:val2}, ...]}`
For example, `{HI233JSfba: [{"name": "Kit's Bat Tester"}]}`

## Dependencies

To run the EC2 bat, you must have the following:
Python 3.8 or higher and the dependencies in requirements.txt
The spyderbat picoagent installed and configured (as described in "Running the bat")

## Running the bat

First, download the picoagent from https://github.com/spyderbat/api_examples/tree/main/picoagent. Find the picoagent.py file and move it into the same directory as this file.
Assuming python is installed, run 'pip install requirements.txt' To start up the ec2bat, fill out the required configuration fields as shown below and the run the following command: 'python3 picoagent.py -c ec2config.json'. Make certain that you are either running this on an ec2 instance with permission to read the tags of other instances, or your environment variables are set as follows: 
AWS_ACCESS_KEY_ID as the access key for the account the bat should fetch tags from
AWS_SECRET_ACCESS_KEY as the secret key for the account the bat should fetch tags from
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