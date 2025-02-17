import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'

// Contract addresses
export const contractAddresses = {
  token: '0xeeCfE85EdcD43288858C56763dD8971333f9256B',
  nft: '0x812f60E4e7822A2439504f5806410244136BFbd2',
  donations: '0xDF832146EF0988567B1eCA64B78fCdDC90EF27F1'
}

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
)
