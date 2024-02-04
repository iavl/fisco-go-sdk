// SPDX-License-Identifier: MIT
pragma solidity 0.6.10;

contract FileHub {
    mapping(address => bool) public registeredUsers;
    mapping(bytes32 => address) public fileOwners;
    mapping(bytes32 => uint256) public fileUploadTime;

    mapping(address => mapping(address => mapping(address => mapping(bytes32 => bool))))
        public fileForwardingRecords;

    event Registered(address indexed account);
    event FileUploaded(address indexed account, bytes32 indexed fileHash);
    event FileForwarded(
        address indexed from,
        address indexed to,
        address indexed proxy,
        bytes32 fileHash
    );
    event FileShared(address indexed from, address indexed to, bytes32 indexed fileHash);

    function register(address account) external {
        registeredUsers[account] = true;

        emit Registered(account);
    }

    function uploadFile(address account, bytes32 fileHash) external {
        require(msg.sender == account, "FileHub: sender not authorized");
//        require(registeredUsers[account], "FileHub: account not registered");
//        require(fileOwners[fileHash] == address(0), "FileHub: file already exists");

        fileOwners[fileHash] = account;
        fileUploadTime[fileHash] = block.timestamp;

        emit FileUploaded(account, fileHash);
    }

    function shareFile(address from, address to, bytes32 fileHash) external {
        require(msg.sender == from, "FileHub: sender not authorized");
//        require(registeredUsers[from], "FileHub: from account not registered");
//        require(registeredUsers[to], "FileHub: to account not registered");
        require(fileOwners[fileHash] == from, "FileHub: file not owned by from account");

        emit FileShared(from, to, fileHash);
    }

    function forwardFile(address from, address to, address proxy, bytes32 fileHash) external {
        require(msg.sender == proxy, "FileHub: sender not authorized");

        fileForwardingRecords[from][to][proxy][fileHash] = true;

        emit FileForwarded(from, to, proxy, fileHash);
    }
}
