/*
 * Spyderbat API UI & Public APIs
 *
 * Restful APIs for use by UI & customers.
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: apisupport@spyderbat.com
 * Generated by: https://openapi-generator.tech
 */


use reqwest;

use crate::apis::ResponseContent;
use super::{Error, configuration};


/// struct for typed errors of method [`src_data_query`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum SrcDataQueryError {
    Status400(crate::models::ValidationError),
    Status403(),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`src_data_query_v2`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum SrcDataQueryV2Error {
    Status400(crate::models::ValidationError),
    Status403(),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`src_send_data`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum SrcSendDataError {
    Status400(crate::models::ValidationError),
    Status403(),
    UnknownValue(serde_json::Value),
}


///  Allows querying of the source data, data is stored as 'records' which are returned as json objects, in nd-json (see ndjson.org) format.   * Data is returned as it is matched, and no ordering guarentees exist.  * The call completes after it has finished searching for matching records.  * The query expression is limited to seaching a 24 hour period of time, it is the callers responsibility to construct an appropriate 24 hour range. * Documentation for the returned spydergraph datatype can be found at https://app.spyderbat.com/schema/spydergraph/index.html  * The user must have both the *org.ListSourceData* action on the org and *source_data.Query* action on the source * To get a count of results (up to 10K) but no data, use querySize: 0  
pub async fn src_data_query(configuration: &configuration::Configuration, src_data_query_input: Option<crate::models::SrcDataQueryInput>) -> Result<String, Error<SrcDataQueryError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/api/v1/source/query/", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = local_var_configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };
    local_var_req_builder = local_var_req_builder.json(&src_data_query_input);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<SrcDataQueryError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// Same as the post query above except results are cached
pub async fn src_data_query_v2(configuration: &configuration::Configuration, org_uid: &str, dt: &str, e: Option<&str>, et: Option<f64>, id: Option<Vec<String>>, pj: Option<Vec<String>>, q: Option<&str>, qf: Option<i32>, qs: Option<i32>, rr: Option<bool>, src: Option<&str>, st: Option<f64>) -> Result<String, Error<SrcDataQueryV2Error>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/api/v1/org/{orgUID}/data/", local_var_configuration.base_path, orgUID=crate::apis::urlencode(org_uid));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    local_var_req_builder = local_var_req_builder.query(&[("dt", &dt.to_string())]);
    if let Some(ref local_var_str) = e {
        local_var_req_builder = local_var_req_builder.query(&[("e", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = et {
        local_var_req_builder = local_var_req_builder.query(&[("et", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = id {
        local_var_req_builder = match "multi" {
            "multi" => local_var_req_builder.query(&local_var_str.into_iter().map(|p| ("id".to_owned(), p.to_string())).collect::<Vec<(std::string::String, std::string::String)>>()),
            _ => local_var_req_builder.query(&[("id", &local_var_str.into_iter().map(|p| p.to_string()).collect::<Vec<String>>().join(",").to_string())]),
        };
    }
    if let Some(ref local_var_str) = pj {
        local_var_req_builder = match "multi" {
            "multi" => local_var_req_builder.query(&local_var_str.into_iter().map(|p| ("pj".to_owned(), p.to_string())).collect::<Vec<(std::string::String, std::string::String)>>()),
            _ => local_var_req_builder.query(&[("pj", &local_var_str.into_iter().map(|p| p.to_string()).collect::<Vec<String>>().join(",").to_string())]),
        };
    }
    if let Some(ref local_var_str) = q {
        local_var_req_builder = local_var_req_builder.query(&[("q", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = qf {
        local_var_req_builder = local_var_req_builder.query(&[("qf", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = qs {
        local_var_req_builder = local_var_req_builder.query(&[("qs", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = rr {
        local_var_req_builder = local_var_req_builder.query(&[("rr", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = src {
        local_var_req_builder = local_var_req_builder.query(&[("src", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = st {
        local_var_req_builder = local_var_req_builder.query(&[("st", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = local_var_configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<SrcDataQueryV2Error> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

/// Sends data to a source
pub async fn src_send_data(configuration: &configuration::Configuration, data_type: &str, org_uid: &str, source_uid: &str, encoding: &str, file: std::path::PathBuf) -> Result<(), Error<SrcSendDataError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/api/v1/org/{orgUID}/source/{sourceUID}/data/{dataType}", local_var_configuration.base_path, dataType=crate::apis::urlencode(data_type), orgUID=crate::apis::urlencode(org_uid), sourceUID=crate::apis::urlencode(source_uid));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    if let Some(ref local_var_token) = local_var_configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };
    let mut local_var_form = reqwest::multipart::Form::new();
    local_var_form = local_var_form.text("encoding", encoding.to_string());
    // TODO: support file upload for 'file' parameter
    local_var_req_builder = local_var_req_builder.multipart(local_var_form);

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        Ok(())
    } else {
        let local_var_entity: Option<SrcSendDataError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

