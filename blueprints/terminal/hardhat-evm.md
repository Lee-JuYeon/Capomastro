---
ide: terminal
appName: Terminal
description: EVM Smart Contract (Solidity + Hardhat)
---

Hardhat scaffold with TypeScript, Ethers.js v6, OpenZeppelin contracts, and local node.
Deploy: npx hardhat run scripts/deploy.ts --network localhost

```json
[
  { "action": "shell", "cmd": "mkdir -p {name}", "cwd": "{outputDir}" },
  { "action": "shell", "cmd": "npm init -y && npm install --save-dev hardhat @nomicfoundation/hardhat-toolbox @openzeppelin/contracts && npx hardhat init --typescript --no-install-deps", "cwd": "{outputDir}/{name}", "timeout": 120000 }
]
```
