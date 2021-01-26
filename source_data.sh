#!/bin/bash
#
### INTRODUCTION
#
# This shell script demonstrates how to send data for external data sources to
# spyderbat
#
### SOURCES 
#
# To send data to spyderbat a source must be specified, the source of data has
# an associated source UID. A source UID can be specified by the user and is 
# used as a way to group data as coming from a particular source. 
# 
# For example a single alert source, this way the source can be monitored to 
# make sure it's actively sending data. Data supplied by a source may be 
# utilized to enrich data for other sources. So for example an IDS system 
# represents a source which may supply data about multiple targets or 
# machines, and might be used to enrich data for those machines. 
# 
### VARIABLES 
#
# SOURCE_UID is the UID associated with a source, this can be user supplied
# and is unique identifier for the source of data, the source UID cannot
# contain '/' 
SOURCE_UID="src_001"
#
# JWT is the JWT API Token for a specific organization
JWT="YOUR API JWT"
#
# ORG_UID is the UID associated with the organization to create the source
# in and to send data to
ORG_UID="YOUR ORG"
#
# DATATYPE is the datatype of data being sent to spyderbat, which can be
# used for routing (and validation of) the data through the analytics 
# system. At the moment the only supported data type is 'sb-agent'. 
DATATYPE="sb-agent"
#
# API_URL for the API endpoint
API_URL=api.tigerbat.com

### USING THE API
#
# You can reference the API documentation here: https://$API_URL/openapi.pdf
# 
# For the API to work you will need to supply a JWT and make sure your API_URL 
# is correctly specified. Each API call requires the JWT to be specified in an 
# authorization header, and the correct content type must be specified. 
# 
### CREATING THE SOURCE
#
# To create a source we need to use the "Create source" API, you can reference
# the api documentation to read the full specification. This call is designed
# to be idempotent, so that no check is required to see if a source exists 
# before creating it. So for example if a process needs to send data for a specific
# source the source can be created each time without needing to verify that the
# source exists. 
# 
# If this call is successful a json document is returned with the UID of the created
# source, or a msg indicating that the source already exists is returned
cat <<EOF | curl -H "Content-Type:application/json" -H "Authorization: Bearer $JWT" https://$API_URL/api/v1/org/$ORG_UID/source/ --data-binary @-
{
	"uid":"$SOURCE_UID",
	"name":"Example source",
	"org_uid":"$ORG_UID"
}
EOF

#
#### SENDING DATA TO THE SOURCE
# 
# To send data to a source, we first need to create the source , which we did 
# above. Data sent to a source should be in ndjson format and must be compressed 
# using gzip compression, and the 'Content-Encoding' header must be set to gzip 
# to indicate that the data is compressed. Uncompressed data is not supported. 
#  
# Data sent to a source will then be forwarded to analytics. For the data to 
# be processed it must be understood by analytics. The basic requirements here 
# are that each JSON object has a 'schema' which can be user specified, a 'time'
# property which must be an EPOCH time, which may have a fractional component,
# and a unique 'id' which must be a UUID. 
# 
# We're going to construct these values
SCHEMA='example_data:1.0'
ID=`head -1000 /dev/urandom | md5sum | cut -f 1 -d " "`
TIME=`date +%s`
cat <<EOF > /tmp/source_data.ndjson
{ "schema":"$SCHEMA", "id":"$ID.0", "time":$TIME.000 }
{ "schema":"$SCHEMA", "id":"$ID.1", "time":$TIME.001 }
{ "schema":"$SCHEMA", "id":"$ID.2", "time":$TIME.002 }
EOF

echo "Sending"
cat /tmp/source_data.ndjson

# When sending source data, it must be compressed using gzip. Notice that 
# gzip is included in the pipe that sends data to curl, and that a 
# 'Content-Encoding' header is specified. The reason we don't use the 
# --compressed flag for curl is because that flag asks for the request 
# response to be compressed, not for the request body to be compressed. 
# Curl doesn't support request body compression.
cat /tmp/source_data.ndjson | gzip | curl -v -v \
	-H "Content-Encoding: gzip" \
	-H "Content-Type:application/ndjson" \
	-H "Authorization: Bearer $JWT" \
	https://$API_URL/api/v1/org/$ORG_UID/source/$SOURCE_UID/data/$DATATYPE \
	--data-binary @-
