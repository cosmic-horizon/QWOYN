use cosmwasm_std::Uint128;
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct InstantiateMsg {
    // nft name
    pub nft_name: String,
    // nft symbol
    pub nft_symbol: String,
    // nft code id
    pub nft_ci: u64,
    // Royalty % based on 10000
    pub royalty_bp: u64,
    // Royalty receiving address
    pub royalty_addr: String,
    // purchase denom amount
    pub nft_price_amount: Uint128,
    // max limit count to mint for whitelisted users
    pub whitelist_mint_limit: u64,
    // mint period for whitelisted users
    pub whitelist_mint_period: u64,
    // max limit count to mint for public users
    pub public_mint_limit: u64,
    // public mint period
    pub public_mint_period: u64,
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum ExecuteMsg {
    StartMint {},
    Mint {},
    Whitelist {
        addrs: Vec<String>,
    },
    PreMint {
        tier_index: u64,
        token_uri: String,
        number: u64,
    },
    AddTierInfo {
        token1_addr: String,
        token1_amount: Uint128,
        token2_addr: Option<String>,
        token2_amount: Option<Uint128>,
        vesting_period: u64,
        max_supply: u64,
    },
    WithdrawFund {},
}

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
#[serde(rename_all = "snake_case")]
pub enum QueryMsg {
    Config {},
    TierInfo { index: u64 },
    TierInfos {},
    IsWhitelisted { addr: String },
    SupplyByAddress { addr: String },
}

// We define a custom struct for each query response
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct ConfigResponse {
    pub nft_addr: String,
    pub nft_symbol: String,
    pub royalty_bp: u64,
    pub royalty_addr: String,
    pub nft_max_supply: u64,
    pub nft_current_supply: u64,
    pub nft_price_amount: Uint128,
    pub mint_start_time: u64,
    pub whitelist_mint_period: u64,
    pub whitelist_mint_limit: u64,
    pub public_mint_period: u64,
    pub public_mint_limit: u64,
}

/// We currently take no arguments for migrations
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, JsonSchema)]
pub struct MigrateMsg {}
