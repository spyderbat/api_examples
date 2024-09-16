use spyderbat_api::{
    apis::{configuration::Configuration, org_api, org_type_api},
    models::{
        NotificationPolicy, OrgAssignRoleInput, OrgInviteUsersInput, OrgUnassignRoleInput,
        OrgUpdateInput,
    },
};

mod common;

/// Test the org update API
#[async_std::test]
async fn test_org_update_api() -> Result<(), Box<dyn std::error::Error>> {
    let (configuration, org_uid) = common::init();
    let org = org_api::org_load(configuration, &org_uid).await?;
    println!("Loaded org {:?}", org.name);

    // update the org
    org_api::org_update(
        configuration,
        &org_uid,
        Some(OrgUpdateInput {
            name: org.name,
            tags: org.tags.or(Some(vec![])).map(|mut x| {
                x.push("test_tag".to_owned());
                x
            }),
            org_type_uid: org.org_type_uid,
            owner_email: org.owner_email,
            owner_uid: org.owner_uid,
            resource_name: org.resource_name,
            resource_policy: org.resource_policy,
            valid_from: org.valid_from,
            valid_to: org.valid_to,
            active_sources: org.active_sources,
            active_users: org.active_users,
            total_sources: org.total_sources,
            total_users: org.total_users,
        }),
    )
    .await?;
    let org = org_api::org_load(configuration, &org_uid).await?;
    println!("Loaded updated org {:?}", org.uid);
    assert!(org.tags.clone().unwrap().contains(&"test_tag".to_owned()));

    // revert the update
    org_api::org_update(
        configuration,
        &org_uid,
        Some(OrgUpdateInput {
            name: org.name,
            tags: org.tags.map(|mut x| {
                x.retain(|x| x != &"test_tag".to_owned());
                x
            }),
            org_type_uid: org.org_type_uid,
            owner_email: org.owner_email,
            owner_uid: org.owner_uid,
            resource_name: org.resource_name,
            resource_policy: org.resource_policy,
            valid_from: org.valid_from,
            valid_to: org.valid_to,
            active_sources: org.active_sources,
            active_users: org.active_users,
            total_sources: org.total_sources,
            total_users: org.total_users,
        }),
    )
    .await?;

    Ok(())
}

/// Disables notifications for the org
async fn disable_org_notifications(
    configuration: &Configuration,
    org_uid: &str,
) -> Result<NotificationPolicy, Box<dyn std::error::Error>> {
    let policy = org_api::org_load_notification_policy(configuration, org_uid).await;
    let policy = policy.unwrap_or(NotificationPolicy::default());

    org_api::org_update_notification_policy(configuration, &org_uid, NotificationPolicy::default())
        .await?;

    Ok(policy)
}

/// Enables notifications for the org
async fn enable_org_notifications(
    configuration: &Configuration,
    org_uid: &str,
    policy: NotificationPolicy,
) -> Result<(), Box<dyn std::error::Error>> {
    org_api::org_update_notification_policy(configuration, org_uid, policy).await?;

    Ok(())
}

/// Test the org roles and notifications API
/// Note: This produces email notifications for the org
#[async_std::test]
async fn test_org_roles_and_notifications() -> Result<(), Box<dyn std::error::Error>> {
    let (configuration, org_uid) = common::init();

    // disable notifications
    let policy = disable_org_notifications(configuration, &org_uid).await?;

    // invite a user to the org
    org_api::org_invite_users(
        configuration,
        &org_uid,
        Some(OrgInviteUsersInput {
            emails: Some(vec!["test@example.com".to_owned()]),
            roles: Some(vec!["Org/UserReadOnly".to_owned()]),
        }),
    )
    .await?;

    let roles = org_api::org_list_role(configuration, &org_uid, None, None, None).await?;
    println!("Loaded org roles {:?}", roles.len());
    // there is at least one user, so there should be at least one role
    assert!(roles.len() > 0);

    let role = roles
        .into_iter()
        .find(|x| x.user_email == Some("test@example.com".to_owned()))
        .expect("Could not find added test user");

    // assign them the Admin role
    org_api::org_assign_role(
        configuration,
        &org_uid,
        Some(OrgAssignRoleInput {
            role_uid: Some("Org/Admin".to_owned()),
            user_uid: role.user_uid.clone(),
        }),
    )
    .await?;

    // unassign the user from the Admin role
    org_api::org_unassign_role(
        configuration,
        &org_uid,
        Some(OrgUnassignRoleInput {
            role_uid: Some("Org/Admin".to_owned()),
            user_uid: role.user_uid.clone(),
        }),
    )
    .await?;

    // unassign the user from the UserReadOnly role
    org_api::org_unassign_role(
        configuration,
        &org_uid,
        Some(OrgUnassignRoleInput {
            role_uid: Some("Org/UserReadOnly".to_owned()),
            user_uid: role.user_uid.clone(),
        }),
    )
    .await?;

    // enable notifications
    enable_org_notifications(configuration, &org_uid, policy).await?;

    Ok(())
}

/// Test the org types API
#[async_std::test]
async fn test_org_type() -> Result<(), Box<dyn std::error::Error>> {
    let (configuration, org_uid) = common::init();

    let org_type = org_type_api::org_type_load(configuration, &org_uid).await?;
    dbg!(&org_type);
    println!("Loaded org type {:?}", org_type.description);

    let org_limits = org_type_api::org_type_limit_active_sources(configuration, &org_uid).await?;
    dbg!(&org_limits);
    println!(
        "Loaded org type active source limits {:?}",
        org_limits.description
    );

    let org_limits = org_type_api::org_type_limit_org_roles(configuration, &org_uid).await?;
    dbg!(&org_limits);
    println!(
        "Loaded org type org roles limits {:?}",
        org_limits.description
    );

    Ok(())
}
