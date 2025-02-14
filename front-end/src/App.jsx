import React, { useState, useEffect } from 'react'
import { ethers } from 'ethers'
import ConnectWallet from './components/ConnectWallet'
import DonationsManage from './components/DonationsManage'
import DonationHome from './components/DonationHome'
import './App.css'

function App() {
  const [isLoading, setIsLoading] = useState(true)
  const [connected, setConnected] = useState(false)
  const [account, setAccount] = useState(null)
  const [provider, setProvider] = useState(null)
  const [signer, setSigner] = useState(null)

  const handleConnect = async (wallet) => {

    console.log('handleConnect被调用了，Wallet connection:', wallet)
    if (!wallet) {
      console.log('Disconnecting wallet')
      setAccount(null)
      setProvider(null)
      setSigner(null)
      setConnected(false)
      return
    }

    try {
      const address = await wallet.signer.getAddress()
      console.log('Setting new account:', address)
      
      // Set new state
      const newProvider = new ethers.BrowserProvider(window.ethereum)
      const newSigner = await newProvider.getSigner()
      
      setAccount(address)
      setProvider(newProvider)
      setSigner(newSigner)
      setConnected(true)
    } catch (error) {
      console.error('Connection error:', error)
      setAccount(null)
      setProvider(null)
      setSigner(null)
      setConnected(false)
    }
  }

  // 修改 handleAccountsChanged 函数
  const handleAccountsChanged = async (accounts) => {
    try {
      if (accounts.length > 0) {
        // 等待 provider 准备就绪
        await new Promise(resolve => setTimeout(resolve, 100));
        
        // 只有当账户真的改变时才更新状态
        const web3Provider = new ethers.BrowserProvider(window.ethereum);
        const newSigner = await web3Provider.getSigner();
        const address = await newSigner.getAddress();
        
        // 检查账户是否真的改变了
        if (address !== account) {
          setAccount(address);
          setProvider(web3Provider);
          setSigner(newSigner);
          setConnected(true);
        }
      } else {
        // 断开连接时清除状态
        setAccount(null);
        setProvider(null);
        setSigner(null);
        setConnected(false);
      }
    } catch (error) {
      console.error('Error handling account change:', error);
      // 出错时也清除状态
      setAccount(null);
      setProvider(null);
      setSigner(null);
      setConnected(false);
    }
  };

  // 修改 init 函数
  useEffect(() => {
    const init = async () => {
      if (window.ethereum) {
        try {
          const provider = new ethers.BrowserProvider(window.ethereum);
          const accounts = await provider.listAccounts();
          
          if (accounts.length > 0) {
            const signer = await provider.getSigner();
            const address = await signer.getAddress();
            
            // 只在初始化时设置一次
            setAccount(address);
            setProvider(provider);
            setSigner(signer);
            setConnected(true);
          }
        } catch (error) {
          console.error('Error initializing wallet:', error);
        } finally {
          setIsLoading(false);
        }

        // 添加事件监听器
        window.ethereum.on('accountsChanged', handleAccountsChanged);
        window.ethereum.on('chainChanged', () => {
          window.location.reload();
        });
      } else {
        setIsLoading(false);
      }
    };
    
    init();

    // 清理事件监听器
    return () => {
      if (window.ethereum) {
        window.ethereum.removeListener('accountsChanged', handleAccountsChanged);
        window.ethereum.removeListener('chainChanged', () => {
          window.location.reload();
        });
      }
    };
  }, []); // 保持空依赖数组

  if (isLoading) {
    return (
      <div className="app-loading">
        <div className="loading-spinner"></div>
        <p>Checking wallet connection...</p>
      </div>
    )
  }

  return (
    <div className="app">
      <header>
        {connected ? (
          <div className="wallet-info">
            Connected: {account.slice(0, 6)}...{account.slice(-4)}
          </div>
        ) : (
          <ConnectWallet onConnect={handleConnect} />
        )}
      </header>

      <main>
        {connected ? (
          <DonationsManage 
            provider={provider}
            signer={signer}
          />
        ) : (
          <DonationHome />
        )}
      </main>
    </div>
  )
}

export default App
