use spyderbat_api::{
    apis::{agent_api, agent_registration_api, agent_work_api},
    models::{
        AgentRegistrationCreateInput, AgentRegistrationUpdateInput, AgentSetAgentWorkInput,
        AgentSetOrgWorkInput, DaoAgentClass, DaoAgentConfig, OrcApiAgentWork,
    },
};
mod common;

/// Test the agent loading and individual agent work API
#[async_std::test]
async fn test_agent_api() -> Result<(), Box<dyn std::error::Error>> {
    let (configuration, org_uid) = common::init();

    // test the agent list endpoint
    let mut agents =
        agent_api::agent_list(configuration, &org_uid, None, None, None, None, None).await?;
    println!("Received {} agents with no filters", agents.len());

    // get the first agent uid for later use
    let agent_uid = agents[0]
        .uid
        .take()
        .ok_or(common::TestError("No agent uid".to_owned()))?;
    // get the agent registration uid
    let agent_registration_uid = agents[0]
        .agent_registration_uid
        .take()
        .ok_or(common::TestError("No agent registration id".to_owned()))?;

    // test agent list endpoint with args
    let agents = agent_api::agent_list(
        configuration,
        &org_uid,
        Some(&agent_registration_uid),
        None,
        None,
        None,
        None,
    )
    .await?;
    println!(
        "Received {} agents with registration uid {}",
        agents.len(),
        agent_registration_uid
    );

    // test that an incorrect org_uid fails
    let failure = agent_api::agent_list(
        &configuration,
        "invalid_org_uid",
        None,
        None,
        None,
        None,
        None,
    )
    .await;
    assert!(failure.is_err());

    // test the agent load endpoint
    println!("Loading agent {}", agent_uid);
    let agent = agent_api::agent_load(configuration, &agent_uid, &org_uid).await?;
    println!("Loaded agent {:?}", agent.uid);

    // test that an incorrect agent_uid fails
    let failure = agent_api::agent_load(configuration, "invalid_agent_uid", &org_uid).await;
    assert!(failure.is_err());

    // test agent work because it needs the agent registration uid
    agent_work_api::agent_set_agent_work(
        configuration,
        &agent_uid,
        &org_uid,
        Some(AgentSetAgentWorkInput {
            work: Some(Box::new(OrcApiAgentWork { work: Some(vec![]) })),
            tags: Some(vec!["test_agent".to_owned()]),
        }),
    )
    .await?;

    let work = agent_work_api::agent_get_agent_work(configuration, &agent_uid, &org_uid).await?;
    println!("Work tags: {:?}", dbg!(work).tags);

    agent_work_api::agent_delete_agent_work(configuration, &agent_uid, &org_uid).await?;

    Ok(())
}

/// Test the agent registration API
#[async_std::test]
async fn test_agent_registration_api() -> Result<(), Box<dyn std::error::Error>> {
    let (configuration, org_uid) = common::init();

    // there is no delete endpoint, but these registrations seem to be automatically deleted
    let mut result = agent_registration_api::agent_registration_create(
        configuration,
        &org_uid,
        Some(AgentRegistrationCreateInput {
            created_by: Some("openapi-tests".to_owned()),
            description: Some("An agent registration created by openapi-tests".to_owned()),
            name: Some("Openapi Test Registration".to_owned()),
            ..Default::default()
        }),
    )
    .await?;
    println!("Created agent registration {:?}", result.uid);

    // get the OpenAPI test agent registration uid
    let reg_uid = result
        .uid
        .take()
        .ok_or(common::TestError("No agent registration uid".to_owned()))?;

    let response =
        agent_registration_api::agent_registration_list(configuration, &org_uid, None, None)
            .await?;
    println!("Received {} agent registrations", dbg!(&response).len());

    // this currently fails because the classes may be null in the response
    // let response =
    //     agent_registration_api::agent_registration_load(&configuration, ORG_UID, &reg_uid).await;
    // dbg!(response)?;

    let _ = agent_registration_api::agent_registration_update(
        configuration,
        &org_uid,
        &reg_uid,
        Some(AgentRegistrationUpdateInput {
            created_by: Some("openapi-tests".to_owned()),
            description: Some("An agent registration updated by openapi-tests".to_owned()),
            config: Some(Box::new(DaoAgentConfig {
                classes: vec![DaoAgentClass {
                    name: "Test Class".to_owned(),
                    rbac_roles: None,
                }],
                ..Default::default()
            })),
            ..Default::default()
        }),
    )
    .await?;

    let reg =
        agent_registration_api::agent_registration_load(configuration, &org_uid, &reg_uid).await?;
    println!("Loaded agent registration {:?}", reg.uid);

    let response =
        agent_registration_api::agent_registration_download_link(configuration, &org_uid, &reg_uid)
            .await?;
    println!("Received download link: {:?}", response.url);

    let response = agent_registration_api::agent_registration_get_log(
        configuration,
        &org_uid,
        &reg_uid,
        None,
        None,
    )
    .await?;
    println!("Received log of length: {}", response.len());

    Ok(())
}

/// Test the agent org work API
#[async_std::test]
async fn test_agent_work_api() -> Result<(), Box<dyn std::error::Error>> {
    let (configuration, org_uid) = common::init();

    agent_work_api::agent_set_org_work(
        configuration,
        &org_uid,
        Some(AgentSetOrgWorkInput {
            work: Some(Box::new(OrcApiAgentWork { work: Some(vec![]) })),
            tags: Some(vec!["test".to_owned()]),
        }),
    )
    .await?;

    let work = agent_work_api::agent_get_org_work(configuration, &org_uid).await?;
    println!("Work tags: {:?}", work.tags);

    agent_work_api::agent_delete_org_work(configuration, &org_uid).await?;

    Ok(())
}
