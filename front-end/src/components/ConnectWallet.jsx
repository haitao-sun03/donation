import React, { useState, useEffect } from 'react'
import { ethers } from 'ethers'
import { Button } from '@mui/material'

export default function ConnectWallet({ onConnect }) {
  const [error, setError] = useState(null)
  const [account, setAccount] = useState(null)

  const connect = async () => {
    console.log("click connect")
    try {
      if (!window.ethereum) {
        throw new Error('Please install MetaMask!')
      }

      const provider = new ethers.BrowserProvider(window.ethereum)
      const accounts = await provider.send('eth_requestAccounts', [])
      const signer = await provider.getSigner()
      
      // Clear local state to ensure proper updates
      // setAccount(null)
      // await new Promise(resolve => setTimeout(resolve, 0))
      
      // // Force state update
      // onConnect(null)
      // await new Promise(resolve => setTimeout(resolve, 0))
      
      onConnect({
        address: accounts[0],
        provider,
        signer
      })
      
      setError(null)
    } catch (err) {
      setError(err.message)
      onConnect(null)
    }
  }

  useEffect(() => {
    console.log('@@@@@@@@@@@@@@@@@')
    if (!window.ethereum) return;

    const handleAccountsChanged = async (accounts) => {
      try {
        if (accounts.length > 0) {
          const provider = new ethers.BrowserProvider(window.ethereum);
          const signer = await provider.getSigner();
          const currentAccount = await signer.getAddress();
          
          setAccount(currentAccount);
          onConnect({
            address: currentAccount,
            provider,
            signer
          });
        } else {
          setAccount(null);
          onConnect(null);
        }
      } catch (error) {
        console.error('Error handling account change:', error);
      }
    };

    // Add event listener
    window.ethereum.on('accountsChanged', handleAccountsChanged);
    
    // Initial check
    const checkAccount = async () => {
      try {
        const provider = new ethers.BrowserProvider(window.ethereum);
        const accounts = await provider.send('eth_accounts', []);
        handleAccountsChanged(accounts);
      } catch (error) {
        console.error('Error checking initial account:', error);
      }
    };
    checkAccount();
    
    // Cleanup
    return () => {
      window.ethereum.removeListener('accountsChanged', handleAccountsChanged);
    };
  }, [onConnect]);

  return (
    <div className="connect-wallet">
      <Button
        variant="contained"
        onClick={connect}
        sx={{
          background: 'linear-gradient(135deg, #6366f1 0%, #4f46e5 100%)',
          color: 'white',
          padding: '12px 24px',
          borderRadius: '12px',
          fontWeight: 600,
          '&:hover': {
            background: 'linear-gradient(135deg, #4f46e5 0%, #4338ca 100%)',
            transform: 'translateY(-1px)',
            boxShadow: '0 4px 12px rgba(99, 102, 241, 0.2)',
          },
        }}
      >
        Connect Wallet
      </Button>
      {error && <p className="error">{error}</p>}
    </div>
  )
}
