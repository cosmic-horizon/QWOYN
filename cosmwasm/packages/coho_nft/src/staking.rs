use cosmwasm_std::Uint128;
use cw721::Cw721ReceiveMsg;
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InstantiateMsg {}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum ExecuteMsg {
    UpdateConfig {
        owner: String,
    },
    AddWhitelistCollection {
        contract_addr: String,
    },
    RemoveWhitelistCollection {
        contract_addr: String,
    },
    ReceiveNft(Cw721ReceiveMsg),
    Unstake {
        nft_contract_addr: String,
        nft_token_id: String,
    },
    Claim {
        nft_contract_addr: String,
        nft_token_id: String,
    },
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum Cw721HookMsg {
    Stake {},
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum QueryMsg {
    Config {},
    WhitelistedCollections {},
    StakeInfo {
        nft_contract_addr: String,
        nft_token_id: String,
    },
    NftDistributionReward {
        nft_contract_addr: String,
        nft_token_id: String,
    },
    StakerInfo {
        addr: String,
    },
    TokenIds {
        addr: String,
        nft_contract_addr: String,
    },
}

// We define a custom struct for each query response
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct ConfigResponse {
    pub owner: String,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct StakeInfoResponse {
    pub claimed_days: u64,
    pub total_staked_time: u64,
    pub is_staking: bool,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct StakerInfoResponse {
    pub token: String,
    pub daily_reward: Uint128,
    pub total_reward: Uint128,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct RewardResponse {
    pub token: String,
    pub reward: Uint128,
}

/// We currently take no arguments for migrations
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct MigrateMsg {}
