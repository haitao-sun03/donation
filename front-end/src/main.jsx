import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'

// Contract addresses
export const contractAddresses = {
  token: '0xaAbdc95A2B717682FFF520Ca19cE562D25B3E478',
  nft: '0x8dffA18188198c03ee59352D44A78e10bC0F0241',
  donations: '0x136BB1424C6D733A0bA457c7ffC44CBbf0FB2153'
  // token: '0x3BC08475d2dD99d73d7dc1842c2E312029f6409d',
  // nft: '0x2A6751FB3d966FFd4DD623CcE03FEDa862cCc701',
  // donations: '0x3C549F7607bFA1Ccaf96fb777D06A1a7BF1B8cC5'
}

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
)
