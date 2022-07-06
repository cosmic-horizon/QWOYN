#[cfg(not(feature = "library"))]
use cosmwasm_std::entry_point;
use cosmwasm_std::{
    to_binary, Binary, Deps, DepsMut, Empty, Env, MessageInfo, Response, StdResult,
};
use cw721_base::{
    msg::{ExecuteMsg, InstantiateMsg as Cw721InstantiateMsg, QueryMsg},
    ContractError, Cw721Contract,
};
use coho_nft::token::{ContractInfoResponse, InstantiateMsg};

use crate::state::{ContractInfo, CONTRACT_INFO};
use crate::token::{TokenInfoExtension};

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn instantiate(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: InstantiateMsg,
) -> StdResult<Response> {
    let contract_info = ContractInfo {
        name: msg.name.clone(),
        symbol: msg.symbol.clone(),
        minter: deps.api.addr_validate(&msg.minter)?,
    };
    CONTRACT_INFO.save(deps.storage, &contract_info)?;

    let tract = Cw721Contract::<TokenInfoExtension, Empty>::default();
    tract.instantiate(
        deps,
        env,
        info,
        Cw721InstantiateMsg {
            name: msg.name,
            symbol: msg.symbol,
            minter: msg.minter,
        },
    )
}

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn execute(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: ExecuteMsg<TokenInfoExtension>,
) -> Result<Response, ContractError> {
    let tract = Cw721Contract::<TokenInfoExtension, Empty>::default();
    tract.execute(deps, env, info, msg)
}

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn query(deps: Deps, env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::ContractInfo {} => to_binary(&query_contract_info(deps)?),
        _ => {
            let tract = Cw721Contract::<TokenInfoExtension, Empty>::default();
            tract.query(deps, env, msg)
        }
    }
}
fn query_contract_info(deps: Deps) -> StdResult<ContractInfoResponse> {
    let contract_info = CONTRACT_INFO.load(deps.storage)?;

    Ok(ContractInfoResponse {
        name: contract_info.name,
        symbol: contract_info.symbol,
        minter: contract_info.minter.to_string(),
    })
}
