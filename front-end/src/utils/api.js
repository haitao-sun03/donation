// src/utils/api.js
import { ethers } from 'ethers';

export async function getCurrentUser() {
  if (window.ethereum) {
    const web3Provider = new ethers.BrowserProvider(window.ethereum);
    const accounts = await web3Provider.listAccounts();
    return accounts[0] || null; // 返回当前连接的钱包地址
  }
  return null;
}


export async function request(url, options = {}) {
    const currentUser = await getCurrentUser();
    const jwt = localStorage.getItem(`jwt_${ currentUser.address.toLowerCase()}`);
    const headers = {
      ...(jwt ? { 'Authorization': `${jwt}` } : {}),
      ...options.headers,
    };
  
    const response = await fetch(`/api${url}`, {
      method: 'GET', // 默认 GET
      ...options,
      headers,
    });
  
    if (!response.ok) {
      if (response.status === 401) {
        console.log('Token expired, redirecting to login...');
      }
      throw new Error(`Request failed with status ${response.status}`);
    }
  
    return response.json();
  }