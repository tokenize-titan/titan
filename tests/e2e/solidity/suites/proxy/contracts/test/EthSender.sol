pragma solidity 0.4.26;


contract EthSender {
    function sendEth(address to) external payable {
        to.transfer(msg.value);
    }
}
