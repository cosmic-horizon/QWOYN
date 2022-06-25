# Introduction

This document contains the json schema for the Avatars, Ships and Planets.  Avatars and Ships will be true NFTs which means that users will be able to remove them from within the game and put them up for sale on stargaze or similar NFT marketplaces.  The planets in-game will not be able to leave the game at the moment but their properties are still described using the json schema.

# Example Schema

Some of these schema was taken from stargaze which has based their attributes on the opensea model.

```
{
  "nftType": {
    "attributes": [
    {
      "trait_type": "hat",
      "value": "bandana"
    },
    {
      "trait_type": "glasses",
      "value": "sunglasses"
    },
    {
      "trait_type": "personality",
      "value": "chill"
    },
    {
      "trait_type": "shirt_color",
      "value": "purple"
    },
    {
      "display_type": "number",
      "trait_type": "generation",
      "value": 1
    }
  ],
  "description": "Just some guy that likes to code and listen to Stargaze Trooprs music.",
  "external_url": "https://example.com/?token_id=1",
  "image": "ipfs://bafybeih3ykpa42eipgtzcrfkeo5nvazcdqhj3oh3ztju44tcoipzsdaauy/images/1.png",
  "animation_url": "ipfs://bafybeia5r3hwyou3iggzfvakjkxu2zy5pt3kjil6nyqzvrqwrrtkwe6xrm/images/Genesis.m4a"
  "name": "Shane Stargaze"
  }
)

# Ships

Ship attributes are broken up into 4 different categories:

- General Attributes
- Ship Specs
- Ship Upgrades
- Ship Items

**JSON Schema**

{
  "ship": {
    "type": 1,
    "owner": Avatar_TokenID,
    "location": "Location of Ship at any given time",
    "battle": "true or false",
    "docked": "true or false",
    "specs": [
      "maxFuel": 100,
      "actFuel": 50,
      "maxCargo": 100,
      "cargo": [
        {
          resourceA: 1
        },
        {
          resourceB: 1
        },
        {
          resourceC: 1
        },
        {
          resourceD: 1
        },
        {
          resourceE: 1
        },
        {
          resourceF: 1
        },
      ],
      "attackRatio": 1,
      "defenseRatio": 1,
      "weaponSystemType": 1,
      "weaponsAccuracy": 0.9,
      "defenseSystemType": 1,
      "engineType": 1,
      "engineEfficiency": 0.9,
      "escapeOdds": 0.99,
      "upgrades": [
        {
        "scannerType": 1,
        },
        {
        "transporter": true,
        },
        {
        "refiner": true,
        },
        {
        "cloaking": true,
        },
        {
        "jumpDrive": true,
        }
      ],
      "items": [
        {
          "maxMines": 10,
          "actMines": 1,
        },
        {
          "maxProbes": 10,
          "actProbes": 1,
        }
      ],
    "external_url": "https://example.com/?token_id=1",
    "image": "ipfs://bafybeih3ykpa42eipgtzcrfkeo5nvazcdqhj3oh3ztju44tcoipzsdaauy/images/1.png",
    "animation_url": "ipfs://bafybeia5r3hwyou3iggzfvakjkxu2zy5pt3kjil6nyqzvrqwrrtkwe6xrm/images/Genesis.m4a"
    "name": "Vedic Cruiser"
}
```

# Avatars

```
{}
```    
