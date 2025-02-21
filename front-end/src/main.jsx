import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'

// Contract addresses
export const contractAddresses = {
  token: '0x2df7E7A787057f0FF66fA41bAa91502ce96f1349',
  nft: '0xEF55F840e37E3C1f24AF005a0505B2B9E42C4Fe5',
  donations: '0xF604d56825F2f09ed8001cCFBf3C9a0caf8C7882'
}

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
)
