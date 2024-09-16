use std::ops::Deref;

use spyderbat_api::{
    apis::{configuration::Configuration, source_api, source_data_api},
    models::{SrcCreateInput, SrcUpdateInput},
};

mod common;

/// Test the source list API
#[async_std::test]
async fn test_source_list() -> Result<(), Box<dyn std::error::Error>> {
    let (configuration, org_uid) = common::init();
    let sources =
        source_api::src_list(&configuration, &org_uid, None, None, None, None, None, None).await?;
    println!("Received {} sources", sources.len());

    let soar_sources = source_api::integration_soar_src_list(
        &configuration,
        &org_uid,
        None,
        None,
        None,
        None,
        None,
        None,
        None,
    )
    .await?;
    println!("Received {} soar sources", soar_sources.len());

    Ok(())
}

/// A full test of the source APIs.
/// This test creates a source, updates it, and then deletes it,
#[async_std::test]
async fn test_sources_full() -> Result<(), Box<dyn std::error::Error>> {
    // default source
    let (configuration, org_uid) = common::init();
    let source =
        source_api::src_create(&configuration, &org_uid, Some(SrcCreateInput::default())).await?;
    println!("Created empty source {:?}", source.uid);
    test_with_id(&configuration, &org_uid, &source.uid.unwrap()).await?;

    // source with data
    let source = source_api::src_create(
        &configuration,
        &org_uid,
        Some(SrcCreateInput {
            name: Some("Test Source".to_owned()),
            tags: Some(vec!["test_source".to_owned()]),
            description: Some("Test Source".to_owned()),
            ..Default::default()
        }),
    )
    .await?;
    println!("Created source {:?}", source.uid);
    test_with_id(&configuration, &org_uid, &source.uid.unwrap()).await?;
    Ok(())
}

/// Test the source APIs with a given source UID.
/// This test updates the source, and then deletes it.
async fn test_with_id(
    configuration: &Configuration,
    org_uid: &str,
    source_uid: &str,
) -> Result<(), Box<dyn std::error::Error>> {
    // fetch the source
    let source = source_api::src_load(&configuration, &org_uid, &source_uid).await?;
    println!("Received source {:?}", source.uid);

    // update the source
    source_api::src_update(
        &configuration,
        &org_uid,
        &source_uid,
        Some(SrcUpdateInput {
            name: source.name,
            description: source.description,
            runtime_description: Some("Test Source description".to_owned()),
            tags: Some(vec!["test_source".to_owned()]),
            ..Default::default()
        }),
    )
    .await?;
    println!("Updated source {:?}", source.uid);

    // delete the source
    source_api::src_delete(&configuration, &org_uid, &source_uid).await?;
    println!("Deleted source {:?}", source.uid);

    Ok(())
}

/// Test the source data API
#[async_std::test]
async fn test_source_data() -> Result<(), Box<dyn std::error::Error>> {
    let (configuration, org_uid) = common::init();
    let source = SrcManager::new().await;
    println!("Created source {:?}", source);

    // Source send does not work yet

    // submit some data to the source
    // source_data_api::src_send_data(
    //     configuration,
    //     "gzip",
    //     "spydergraph",
    //     &org_uid,
    //     &source,
    //     &format!(
    //         r#"{{"schema":"example_data:0.1.0","id":"test_uuid{0}","time":"{0}"}}"#,
    //         SystemTime::now()
    //             .duration_since(UNIX_EPOCH)
    //             .unwrap()
    //             .as_secs_f64()
    //     ),
    // )
    // .await?;
    // println!("Sent data to source {:?}", source);

    // fetch the source data
    let source_data = source_data_api::src_data_query_v2(
        configuration,
        &org_uid,
        "spydergraph",
        None,
        None,
        None,
        None,
        None,
        None,
        None,
        None,
        Some(&source),
        None,
    )
    .await;
    // we may get an error if no data is available
    let source_data = match source_data {
        Ok(source_data) => source_data,
        Err(err) => {
            println!("Error: {:?}", err);
            return Ok(());
        }
    };
    dbg!(source_data);

    Ok(())
}

/// A source manager that creates a source and deletes it.
#[derive(Debug)]
struct SrcManager {
    source_id: String,
}

impl Drop for SrcManager {
    fn drop(&mut self) {
        let (configuration, org_uid) = common::init();
        async_std::task::block_on(source_api::src_delete(
            &configuration,
            &org_uid,
            &self.source_id,
        ))
        .unwrap();
    }
}

impl SrcManager {
    async fn new() -> Self {
        let (configuration, org_uid) = common::init();
        let source = source_api::src_create(
            &configuration,
            &org_uid,
            Some(SrcCreateInput {
                name: Some("Test Source".to_owned()),
                tags: Some(vec!["test_source".to_owned()]),
                description: Some("Test Source".to_owned()),
                _type: Some("spydergraph".to_owned()),
                ..Default::default()
            }),
        )
        .await
        .unwrap();
        SrcManager {
            source_id: source.uid.unwrap(),
        }
    }
}

impl Deref for SrcManager {
    type Target = String;

    fn deref(&self) -> &Self::Target {
        &self.source_id
    }
}
