// module.exports = async ({ getNamedAccounts, deployments }) => {
//   const { deployer } = await getNamedAccounts();

//   const proxy = await deployments.get("DonationsManageContractProxy");
//   console.log("the proxy address is :",proxy.address)

//   const proxyOwnableUpgradeable = await ethers.getContractAt("OwnableUpgradeable", proxy.address);
//   const owner = await proxyOwnableUpgradeable.owner();
//   console.log("initalize owner:", owner);
//   console.log("current upgrade deployer:", deployer);
//   if (owner.toLowerCase() !== deployer.toLowerCase()) {
//     throw new Error("deployer is not the owner, cannot upgrade");
//   }

//   const DonationsManageContractV2Factory = await ethers.getContractFactory("DonationsManageContractV2");

//   console.log("upgrading")
//   // 部署代理合约并初始化
//   const donationsManageContractV2 = await hre.upgrades.upgradeProxy(
//     proxy.address,
//     DonationsManageContractV2Factory,
//     { kind: 'uups' ,
//       call: { fn: "initializeV2", args: [] } // 显式调用V2初始化
//     }
//   );
//   // 获取逻辑合约地址
//   const implAddress = await hre.upgrades.erc1967.getImplementationAddress(proxy.address);
//   console.log("upgraded Logic Contract Address:", implAddress);

//   const version = await donationsManageContractV2.getVersion();
//   console.log("Current version after upgrade:", version);
// };

// module.exports.tags = ["upgrade"];
