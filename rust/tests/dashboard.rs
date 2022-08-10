use std::collections::HashMap;

use spyderbat_api::{
    apis::{configuration::Configuration, dashboard_search_api},
    models::{DashboardSearchCreateInput, DashboardSearchUpdateInput},
};

#[macro_use]
mod common;

/// Test the dashboard list API
#[async_std::test]
async fn dashboard_list_test() -> Result<(), Box<dyn std::error::Error>> {
    let (configuration, org_uid) = common::init();
    let dashboard_list =
        dashboard_search_api::dashboard_search_list(&configuration, &org_uid).await?;
    println!("Received {} dashboards", dashboard_list.len());
    Ok(())
}

/// A full test of the dashboard APIs.
/// This test creates a dashboard, updates it, and then deletes it,
/// for a default dashboard card and a custom dashboard card.
#[async_std::test]
async fn full_dashboard_test() -> Result<(), Box<dyn std::error::Error>> {
    let (configuration, org_uid) = common::init();

    // default dashboard card
    let dashboard_uid = dashboard_search_api::dashboard_search_create(
        &configuration,
        &org_uid,
        Some(DashboardSearchCreateInput {
            search: Some("schema:event_red_flag".to_owned()),
            ..Default::default()
        }),
    )
    .await?;

    println!("Created empty dashboard {:?}", dashboard_uid);
    test_with_id(&configuration, &org_uid, &dashboard_uid).await?;

    // custom dashboard card
    let dashboard_uid = dashboard_search_api::dashboard_search_create(
        configuration,
        &org_uid,
        Some(DashboardSearchCreateInput {
            data: Some(HashMap::new()),
            description: Some("A test dashboard card".to_owned()),
            notify: Some(false),
            notify_frequency: Some(86400),
            search: Some("schema:event_red_flag".to_owned()),
            tags: Some(vec!["test_dashboard".to_owned()]),
        }),
    )
    .await?;

    println!("Created dashboard {:?}", dashboard_uid);
    test_with_id(&configuration, &org_uid, &dashboard_uid).await?;

    Ok(())
}

/// Test the dashboard search API with a specific id.
/// This handles updating and deleting the dashboard.
async fn test_with_id(
    configuration: &Configuration,
    org_uid: &str,
    dashboard_uid: &str,
) -> Result<(), Box<dyn std::error::Error>> {
    let dashboard =
        dashboard_search_api::dashboard_search_get(configuration, dashboard_uid, org_uid).await?;
    println!("Loaded dashboard {:?}", dashboard);

    dashboard_search_api::dashboard_search_update(
        configuration,
        dashboard_uid,
        org_uid,
        Some(DashboardSearchUpdateInput {
            data: None,
            tags: dashboard.tags.map(|x| {
                x.into_iter()
                    .chain(vec!["updated".to_owned()].into_iter())
                    .collect()
            }),
            description: dashboard.description,
            notify: dashboard.notify,
            notify_frequency: dashboard.notify_frequency,
            search: dashboard.search,
        }),
    )
    .await?;
    let updated_dashboard =
        dashboard_search_api::dashboard_search_get(configuration, dashboard_uid, org_uid).await?;
    assert!(updated_dashboard
        .tags
        .map(|x| x.contains(&"updated".to_owned()))
        .unwrap_or(true));

    dashboard_search_api::dashboard_search_delete(configuration, dashboard_uid, org_uid).await?;

    print!("Deleted dashboard {:?}", dashboard_uid);

    Ok(())
}
