const { usePlugin } = require('@nomiclabs/buidler/config')

usePlugin("@nomiclabs/buidler-ganache")
usePlugin('@nomiclabs/buidler-truffle5')

module.exports = {
  networks: {
    // Development network is just left as truffle's default settings
    ganache: {
      url: 'http://localhost:8545',
      gasLimit: 5000000,
      gasPrice: 1000000000,  // 1 gwei (in wei)
      defaultBalanceEther: 100
    },
    titan: {
      url: 'http://localhost:8545',
      gasLimit: 5000000,     // Gas sent with each transaction
      gasPrice: 100000000000,  // 100 gwei (in wei)
    },
  },
  solc: {
    version: '0.4.26',
    optimizer: {
      enabled: true,
      runs: 10000,
    },
  },
}
