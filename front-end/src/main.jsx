import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'

// Contract addresses
export const contractAddresses = {
  token: '0x8D192A2f68700AEAeA4F5D3ADd666198e2047A81',
  nft: '0x0A7Ceb2B9707123ceD34B0f6e444Cc5562bEA3DD',
  donations: '0x6Cf23549b2678027E4A1dDdC34c59c6255d05D13'
}

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
)
