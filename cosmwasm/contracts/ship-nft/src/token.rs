use cosmwasm_std::Uint256;
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct TokenInfoExtension {
    /// ship type
    pub ship_type: u64,

    /// avatar token id
    pub owner: Uint256,
}
