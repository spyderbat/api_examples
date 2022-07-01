# Ec2bat user guide

### Description

The EC2 bat runs using the picoagent, and fetches metadata tags from amazon's ec2 service.
It then outputs the tags as a series of json records with the following format:
`{instanceID: [{key1:val1}, {key2:val2}, ...]}`
For example, `{HI233JSfba: [{"name": "Kit's Bat Tester"}]}`

## Dependencies

To run the EC2 bat, you must have the following:
Python 3.6 or higher
Packages:
    boto3

## Running the bat

To start up the ec2bat, fill out the required configuration fields as shown below and the run the following command: 'python3 picoagent.py -c ec2config.json'.
Optionally, use -v to output the results of the bat to stdout. By default, nothing will be outputted to stdout - this is normal, and the bat is still running correctly.

## Configuration Fields (Required)

To set the following fields, edit ec2config.json

`API Key`: The API key \
`Org ID`: The organization ID \
`URL`: The API url to use (default: kangaroobat.net) \

Under `bats` enter your AWS access key to replace YOURKEY in `-k=YOURKEY` and your secret key to replace YOURSECRETKEY in `-sk=YOURSECRETKEY`