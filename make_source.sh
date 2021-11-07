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
SOURCE_UID="bs_src_01"
#
# JWT is the JWT API Token for a specific organization
JWT=${I_API_KEY}
#
# ORG_UID is the UID associated with the organization to create the source
# in and to send data to
ORG_UID="spyderbatuid"
#
# DATATYPE is the datatype of data being sent to spyderbat, which can be
# used for routing (and validation of) the data through the analytics 
# system. At the moment the only supported data type is 'sb-agent'. 
DATATYPE="sb-agent"
#
# API_URL for the API endpoint
API_URL=api.kangaroobat.net

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
	"name":"Brian API SRC",
	"org_uid":"$ORG_UID"
}
EOF
