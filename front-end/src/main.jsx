import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'

// Contract addresses
export const contractAddresses = {
  token: '0xAC6909951eD3c2cA4a4DdC58C0db8dAD92772Dd2',
  nft: '0xb8055c600A54587E6A0f324A2e2ac996c39D9558',
  donations: '0x4c142a072B494a47A4bfe8e1e49427c36c643F34'
}

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
)
