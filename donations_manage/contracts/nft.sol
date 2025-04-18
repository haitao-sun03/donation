// SPDX-License-Identifier: MIT
// Compatible with OpenZeppelin Contracts ^5.0.0
pragma solidity ^0.8.22;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract NFT is ERC721, ERC721Enumerable, ERC721URIStorage, ERC721Burnable, Ownable,ReentrancyGuard  {
    uint256 private _nextTokenId;

    mapping(string=>string) public levelToMetaDataURI;

    // 定义一个结构体来存储键值对
    struct LevelMetaDataURI {
        string level;
        string uri;
    }

    // 事件定义，使用数组来传递映射内容
    event NFTDeployed(LevelMetaDataURI[] levelToMetaDataURIs);
    event NFTMinted(address to,uint256 tokenId,string level);

    constructor(string memory tokenName,string memory tokenSymol)
        ERC721(tokenName, tokenSymol)
        Ownable(msg.sender) {
            levelToMetaDataURI["Diamond"] = "ipfs://QmXHntjAgjthxXGmtQ33QPPLr3zh5Xy1r7Dnd5y3rQh9Ma";
            levelToMetaDataURI["Gold"] = "ipfs://QmRysa1LAFq2yBVsQfztVavdJCQC9EHSe2KD9xHm2p8xkt";
            levelToMetaDataURI["Silver"] = "ipfs://QmWuxajrCBXD8QmS41xbo9mkcMj8PDCMUfBQZsFpP1TWZe";
        }
    
    // 添加一个函数来触发NFTDeployed事件
    function triggerNFTDeployedEvent() public onlyOwner {
         // 将mapping转换为数组
        LevelMetaDataURI[] memory uris = new LevelMetaDataURI[](3);
        uris[0] = LevelMetaDataURI("Diamond", levelToMetaDataURI["Diamond"]);
        uris[1] = LevelMetaDataURI("Gold", levelToMetaDataURI["Gold"]);
        uris[2] = LevelMetaDataURI("Silver", levelToMetaDataURI["Silver"]);
        // 发送事件
        emit NFTDeployed(uris);
    }

    function safeMint(address to,string memory level) public nonReentrant  {
        uint256 tokenId = _nextTokenId++;
        _setTokenURI(tokenId,  levelToMetaDataURI[level]);
        _safeMint(to, tokenId);
        emit NFTMinted(to, tokenId, level);
    }

    // The following functions are overrides required by Solidity.

    function _update(address to, uint256 tokenId, address auth)
        internal
        override(ERC721, ERC721Enumerable)
        returns (address) {
        return super._update(to, tokenId, auth);
    }

    function _increaseBalance(address account, uint128 value)
        internal
        override(ERC721, ERC721Enumerable) {
        super._increaseBalance(account, value);
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (string memory) {
        return super.tokenURI(tokenId);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, ERC721Enumerable, ERC721URIStorage)
        returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}