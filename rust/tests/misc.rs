use spyderbat_api::{
    apis::{metadata_api_api, metrics_data_api, rbac_api},
    models::{CanUserPerformInput, MetricsDataQueryInput, RbacAction},
};

mod common;

/// Test the Schema API
#[async_std::test]
async fn test_meta_schema() -> Result<(), Box<dyn std::error::Error>> {
    let (configuration, _org_uid) = common::init();
    let schema = metadata_api_api::metadata_search_schema(configuration).await?;
    println!("Loaded meta schema {:?}", schema.len());
    Ok(())
}

/// Test the Metrics API
/// Note: this API is obsolete and will likely be removed in the future.
#[async_std::test]
async fn test_metrics_query() -> Result<(), Box<dyn std::error::Error>> {
    let (configuration, org_uid) = common::init();
    let metrics = metrics_data_api::metrics_data_query(
        configuration,
        Some(MetricsDataQueryInput {
            data_type: "metric".to_owned(),
            org_uid,
            ..Default::default()
        }),
    )
    .await;
    match metrics {
        Ok(metrics) => {
            println!("Loaded metrics {:?}", metrics);
        }
        Err(err) => {
            if err.to_string().contains("EOF") {
                println!("No metrics were loaded {:?}", err);
            } else {
                return Err(err.into());
            }
        }
    }
    Ok(())
}

/// Test the RBAC API
#[async_std::test]
async fn test_rbac() -> Result<(), Box<dyn std::error::Error>> {
    let (configuration, org_uid) = common::init();
    let rbac = rbac_api::can_user_perform(
        configuration,
        Some(CanUserPerformInput {
            actions: Some(vec![RbacAction {
                action: Some("org:ListOrgRoles".to_owned()),
                resource_name: Some(format!("srn:org::{0}:{0}", org_uid)),
                ..Default::default()
            }]),
        }),
    )
    .await?;
    println!("Loaded rbac {:?}", rbac);
    Ok(())
}
