import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'

// Contract addresses
export const contractAddresses = {
  token: '0xeeCfE85EdcD43288858C56763dD8971333f9256B',
  nft: '0x5102579fdA461CAEf1cf4F4aa8C4fFb4639e34fC',
  donations: '0x7cd3A9c4A78345450aE7e6FC0c467322109e7661'
}

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
)
