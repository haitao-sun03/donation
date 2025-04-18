module.exports = async ({ getNamedAccounts, deployments }) => {
    
    const proxyAddress = '0xB218f8A4Bc926cF1cA7b3423c154a0D627Bdb7E5';
    // 获取逻辑合约地址
    const implAddress = await hre.upgrades.erc1967.getImplementationAddress(proxyAddress);
  
    // 打印关键信息
    console.log("================ L1 Infos ================");
    console.log("Proxy Contract Address:", proxyAddress);
    console.log("Logic Contract Address:", implAddress);
    console.log("=================================================");
  
  };
  
  module.exports.tags = ["temp"];
  