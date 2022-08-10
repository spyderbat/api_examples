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
pub struct DaoAgentClass {
    #[serde(rename = "name")]
    pub name: String,
    #[serde(rename = "rbac_roles", skip_serializing_if = "Option::is_none")]
    pub rbac_roles: Option<Vec<String>>,
}

impl DaoAgentClass {
    pub fn new(name: String) -> DaoAgentClass {
        DaoAgentClass {
            name,
            rbac_roles: None,
        }
    }
}


