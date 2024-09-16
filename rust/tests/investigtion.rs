use std::collections::HashMap;

use spyderbat_api::{
    apis::{configuration::Configuration, investigation_api},
    models::{InvestigationCreateInput, InvestigationUpdateInput},
};

mod common;

/// Test the investigation list API
#[async_std::test]
async fn test_investigation_list() -> Result<(), Box<dyn std::error::Error>> {
    let (configuration, org_uid) = common::init();
    let investigations = investigation_api::investigation_list(&configuration, &org_uid).await?;
    println!("Received {} investigations", investigations.len());
    Ok(())
}

/// A full test of the investigation APIs.
/// This test creates a investigation, updates it, and then deletes it,
/// for a default investigation and a custom investigation.
#[async_std::test]
async fn test_investigation_full() -> Result<(), Box<dyn std::error::Error>> {
    // default investigation
    let (configuration, org_uid) = common::init();
    let investigation = investigation_api::investigation_create(
        &configuration,
        &org_uid,
        Some(InvestigationCreateInput::default()),
    )
    .await?;
    println!("Created empty investigation {:?}", investigation.uid);
    test_with_id(&configuration, &org_uid, &investigation.uid.unwrap()).await?;

    // investigation with data
    let investigation = investigation_api::investigation_create(
        &configuration,
        &org_uid,
        Some(InvestigationCreateInput {
            data: Some(HashMap::new()),
            name: Some("Test Investigation".to_owned()),
            tags: Some(vec!["test_investigation".to_owned()]),
            ..Default::default()
        }),
    )
    .await?;
    println!("Created investigation {:?}", investigation.uid);
    test_with_id(&configuration, &org_uid, &investigation.uid.unwrap()).await?;
    Ok(())
}

/// Test the investigation APIs with a given investigation UID.
/// This test updates the investigation, tests versions,
/// and then deletes it.
async fn test_with_id(
    configuration: &Configuration,
    org_uid: &str,
    investigation_uid: &str,
) -> Result<(), Box<dyn std::error::Error>> {
    // fetch the investigation
    let investigation =
        investigation_api::investigation_load(&configuration, &investigation_uid, &org_uid).await?;
    println!("Received investigation {:?}", investigation.uid);

    // get all versions of the investigation
    let versions =
        investigation_api::investigation_list_versions(configuration, investigation_uid, org_uid)
            .await?;

    println!("Received {} versions", versions.len());

    // update the investigation
    investigation_api::investigation_update(
        configuration,
        investigation_uid,
        org_uid,
        Some(InvestigationUpdateInput {
            name: Some("Test Investigation Updated".to_owned()),
            version: Some(1),
            ..Default::default()
        }),
    )
    .await?;

    // get the first version
    let v0 =
        investigation_api::investigation_load_version(configuration, investigation_uid, org_uid, 0)
            .await?;
    println!("Loaded version 0 {:?}", v0.uid);

    // delete the investigation
    investigation_api::investigation_delete(configuration, investigation_uid, org_uid).await?;

    Ok(())
}
