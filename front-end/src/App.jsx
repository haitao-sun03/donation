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

  // Initialize provider and setup event listeners
  useEffect(() => {
    let accountsChangedHandler
    let chainChangedHandler

    const init = async () => {
      if (window.ethereum) {
        try {
          const provider = new ethers.BrowserProvider(window.ethereum)
          setProvider(provider)
          
          const accounts = await provider.listAccounts()
          if (accounts.length > 0) {
            setAccount(accounts[0].address)
            setSigner(await provider.getSigner())
            setConnected(true)
          }
        } catch (error) {
          console.error('Error initializing wallet:', error)
        } finally {
          setIsLoading(false)
        }

        // Add event listeners
        accountsChangedHandler = async (accounts) => {
          if (accounts.length > 0) {
            setAccount(accounts[0])
            setSigner(await provider.getSigner())
            setConnected(true)
          } else {
            setAccount(null)
            setSigner(null)
            setConnected(false)
          }
        }

        chainChangedHandler = () => {
          window.location.reload()
        }

        window.ethereum.on('accountsChanged', accountsChangedHandler)
        window.ethereum.on('chainChanged', chainChangedHandler)
      }
    }
    
    init()

    // Cleanup event listeners
    return () => {
      if (window.ethereum) {
        if (accountsChangedHandler) {
          window.ethereum.removeListener('accountsChanged', accountsChangedHandler)
        }
        if (chainChangedHandler) {
          window.ethereum.removeListener('chainChanged', chainChangedHandler)
        }
      }
    }
  }, [])

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
