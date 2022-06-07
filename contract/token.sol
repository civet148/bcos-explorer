

pragma solidity ^0.4.24;
//pragma experimental ABIEncoderV2; //ABIEncoderV2支持合约方法返回结构体

//代币合约
contract Token {

    address _owner; //合约拥有者地址
    mapping(address => uint64) _tokens; //用户token存储对象

    //构造函数(部署合约时初始化拥有者代币总数，也就是市面上常说的代币发行总数)
    constructor(uint64 amount) public {
        _owner = msg.sender; //部署合约时上链消息的发送者设置为合约owner (msg对象为solidity合约上链消息上下文对象)
        _tokens[msg.sender] = amount; //部署合约时设置拥有者的token数, token数为了简单不设置小数点仅以整数表示(可转账给别的地址)
    }

    //查询: 合约owner地址（无需签名）
    function getOwner() public view returns (address) {
        return _owner;
    }

    //查询: 获取某地址的token余额（无需签名）
    function balanceOf(address addr) public view returns (uint64) {
        return  _tokens[addr];
    }

    //操作:    token转账（需私钥签名）
    //from    签名人地址
    //to      目标地址
    //amount  要转移的数量
    function transferFrom(address from, address to, uint64 amount) public returns (bool) {
        //检查from地址是否为签名人合法地址
        require(msg.sender==from, "from address is not msg sender"); //检查不通过则终止执行并记录错误日志
        //检查to地址是否合法
        require(to != address(0), "to address is 0x0");//检查不通过则终止执行并记录错误日志
        //检查amount是否为0或大于from地址的余额
        require(amount != 0, "transfer amount is 0");//检查不通过则终止执行并记录错误日志
        //检查from地址余额是否足够
        uint256 amt = _tokens[from];
        require(amt >= amount, "from address no enough balance");//检查不通过则终止执行并记录错误日志
        //将from地址的amount资金转给to地址
        _tokens[from] -= amount;
        _tokens[to] += amount;
        return true;
    }
}