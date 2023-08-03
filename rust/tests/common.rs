//! A set of common types and utilities for testing.

use reqwest::Client;
use spyderbat_api::apis::configuration::Configuration;
use spyderbat_api::apis::org_api;
use std;
use std::error::Error;
use std::fmt::Display;

use std::sync::Once;

static INIT: Once = Once::new();
static mut CONFIGURATION: Option<Configuration> = None;
static mut ORG_UID: Option<String> = None;

/// A custom error type for testing.
#[derive(Debug)]
pub struct TestError(pub String);

impl Display for TestError {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "{}", self.0)
    }
}

impl Error for TestError {}

/// Returns the configuration and org_uid for testing.
/// Initialization is done once.
pub fn init() -> (&'static Configuration, String) {
    INIT.call_once(|| {
        let client = Client::builder().build().unwrap();
        let configuration = Configuration {
            base_path: format!("https://{}", env!("SPYDERBAT_API_URL")),
            client,
            bearer_access_token: Some(env!("SPYDERBAT_API_BEARER_ACCESS_TOKEN").to_owned()),
            ..Default::default()
        };

        let response = async_std::task::block_on(org_api::org_list(
            &configuration,
            None,
            None,
            None,
            None,
            None,
            None,
            None,
            None,
        ))
        .unwrap_or_else(|x| panic!("Failed to get org list; {:?}", x));

        let mut org_uid = response[0].uid.clone();
        if org_uid == Some("defend_the_flag".to_owned()) {
            org_uid = response[1].uid.clone();
        }

        // Safety: we know that the configuration is valid, and the org uid is valid.
        // if another thread is accessing the configuration, it may beat this thread,
        // but should read a "None" value, panicking.
        // Essentially, if this causes bad behavior, it *should* fail tests pretty clearly.
        unsafe {
            ORG_UID = org_uid;
            CONFIGURATION = Some(configuration);
        }
    });
    // Safety: we know that the initialization has been done, and the configuration is valid.
    // the same data races as above apply.
    unsafe {
        if let Some(config) = &CONFIGURATION {
            return (config, ORG_UID.clone().unwrap());
        } else {
            panic!("Configuration not initialized");
        }
    }
}
