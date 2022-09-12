use cosmwasm_std::Addr;
use cw_storage_plus::Item;
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Clone, PartialEq, JsonSchema, Debug)]
pub struct ContractInfo {
    pub name: String,
    pub symbol: String,
    pub minter: Addr,
    pub owner: Addr
}

pub const CONTRACT_INFO: Item<ContractInfo> = Item::new("ship_nft");
