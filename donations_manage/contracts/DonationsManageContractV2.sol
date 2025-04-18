// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./DonationsManageContract.sol";

/**
 * @title donates contract
 * @author Zhang haitao
 */
contract DonationsManageContractV2 is DonationsManageContract {
    string private constant version = "2.0.0";

    function getVersion() external pure override returns(string memory) {
        return version;
    }

     function initializeV2() public reinitializer(2) onlyOwner {
        // 可添加V2特有的初始化逻辑
        // 如果没有新初始化需求，函数体可为空
    }
}
