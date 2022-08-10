/*
 * Spyderbat API UI & Public APIs
 *
 * Restful APIs for use by UI & customers.
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: apisupport@spyderbat.com
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Default, Serialize, Deserialize)]
pub struct ApiSourceCreateHandlerOutput {
    #[serde(rename = "uid", skip_serializing_if = "Option::is_none")]
    pub uid: Option<String>,
}

impl ApiSourceCreateHandlerOutput {
    pub fn new() -> ApiSourceCreateHandlerOutput {
        ApiSourceCreateHandlerOutput {
            uid: None,
        }
    }
}


