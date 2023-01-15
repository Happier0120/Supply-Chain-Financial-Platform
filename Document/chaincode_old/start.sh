#!/bin/bash
echo "一、清理环境，重启网络并创建通道mychannel"
./network.sh down
./network.sh up createChannel

echo "二、安装链码"
./network.sh deployCC -ccn cc_scf_3 -ccp ../chaincode -ccl go -ccep "OR('Org1MSP.peer','Org2MSP.peer')" -cccg ../chaincode/api/collections_config_1.json

echo "三、使用or1环境"
export PATH=${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=localhost:7051

echo "四、初始化链码'"
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n cc_scf_3 -c '{"function":"Init","Args":[]}' 
sleep 5
echo "当前账户列表"
peer chaincode query -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n cc_scf_3 -c '{"function":"QueryAccountaById","Args":[""]}'

echo "五、创建一个资产"
export Ticket_PROPERTIES=$(echo -n "{\"price\":\"101\",\"fromOrder\":\"ccccit\"}" | base64 | tr -d \\n)
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n cc_scf_3 -c '{"function":"CreateTicket","Args":["Org2MSP", "A new ticket for Org2MSP"]}' --transient "{\"ticket_properties\":\"$Ticket_PROPERTIES\"}"
sleep 5
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n cc_scf_3 -c '{"function":"CreateTicket","Args":["Org1MSP", "A new ticket for Org1MSP"]}' --transient "{\"ticket_properties\":\"$Ticket_PROPERTIES\"}"
sleep 5

"{\"objectType\":\"ticketCollection\",\"ticketID\":\"5a84jjJkwz\",\"ownerOrg\":\"Org1MSP\",\"channelID\":\"mychannel\",\"txID\":\"0e9f693f122527ce4cfbca7c04be861227a2b138499ca4aaa565c8fa8b2629f9\",\"txTimestamp\":\"seconds:1673526253 nanos:736482000\",\"createTime\":\"2023-01-12 12:24:13\",\"duedate\":\"2023-12-12\",\"guarantor\":\"ICBC\",\"description\":\"A new ticket for Org1MSP\"}",
"{\"objectType\":\"ticketCollection\",\"ticketID\":\"BpLnfgDsc2\",\"ownerOrg\":\"Org1MSP\",\"channelID\":\"mychannel\",\"txID\":\"e34350ea0ec1ad6701f09e12b596ae775662a367457c2de8bf0efe63290686f4\",\"txTimestamp\":\"seconds:1673525863 nanos:977599000\",\"createTime\":\"2023-01-12 12:17:44\",\"duedate\":\"2023-12-12\",\"guarantor\":\"ICBC\",\"description\":\"A 200 ticket for Org1MSP\"}",
"{\"objectType\":\"ticketCollection\",\"ticketID\":\"VuS9jZ8uVb\",\"ownerOrg\":\"Org1MSP\",\"channelID\":\"mychannel\",\"txID\":\"be36021618f400b144e41babf4370eef104c07a0982ff39b9d51f77bee076082\",\"txTimestamp\":\"seconds:1673526524 nanos:344196000\",\"createTime\":\"2023-01-12 12:28:44\",\"duedate\":\"2023-12-12\",\"guarantor\":\"ICBC\",\"description\":\"A new 002 ticket for Org1MSP\"}"

echo "六、查询票据"  WD8F2qNfHK   BpLnfgDsc2
echo "6.1 票据公开数据"
peer chaincode query -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n cc_scf_3 -c '{"function":"QueryTicket","Args":["WD8F2qNfHK"]}'
echo "6.2 票据隐私数据"
peer chaincode query -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n cc_scf_3 -c '{"function":"ReadTicketPrivateProperties","Args":["WD8F2qNfHK"]}'
echo "6.3 查询所有数据"
peer chaincode query -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n cc_scf_3 -c '{"function":"GetAllTicket","Args":[""]}'

echo "七、 更改票据备注"
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n cc_scf_3 -c '{"function":"ChangeDescription","Args":["WD8F2qNfHK","new01"]}'
sleep 5

echo "八、 转让票据"
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n cc_scf_3 -c '{"function":"TransferTicket","Args":["WD8F2qNfHK","Org2MSP"]}' --transient "{\"ticket_properties\":\"$Ticket_PROPERTIES\"}" --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"
sleep 5
echo "转让后票据公开信息依然可查看"
peer chaincode query -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n cc_scf_3 -c '{"function":"QueryTicket","Args":["ticket01"]}'
echo "转让后票据隐私信息无法再查看"
peer chaincode query -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n cc_scf_3 -c '{"function":"ReadTicketPrivateProperties","Args":["ticket01"]}'

echo "九、切换org2环境"
export PATH=${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=localhost:9051
echo "只有coreOrg可创建票据，org2无法创建票据"
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n cc_scf_3 -c '{"function":"CreateTicket","Args":["ticket01", "A new ticket for Org1MSP"]}' --transient "{\"ticket_properties\":\"$Ticket_PROPERTIES\"}"
echo "org2可查看接收到票据的隐私信息"
peer chaincode query -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n cc_scf_3 -c '{"function":"ReadTicketPrivateProperties","Args":["ticket01"]}'

echo "十、对票据交易全流程溯源"
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n cc_scf_3 -c '{"function":"QueryTicketHistory","Args":["ticket01"]}'


# echo "清理环境"
# ./network.sh down

{\"objectType\":\"ticketCollection\",\"ticketID\":\"ticket01\",\"ownerOrg\":\"Org1MSP\",\"channelID\":\"mychannel\",\"txID\":\"c94ccc9554ef06f8a3c7ca4ec01746af1e4a35c54ada70e927fff483ca86b9ad\",\"txTimestamp\":\"seconds:1670338031 nanos:941528000\",\"createTime\":\"6066-126-126\",\"duedate\":\"2023-12-12\",\"guarantor\":\"ICBC\",\"description\":\"A new ticket for Org1MSP\"}","{\"objectType\":\"ticketCollection\",\"ticketID\":\"ticket02\",\"ownerOrg\":\"Org1MSP\",\"channelID\":\"mychannel\",\"txID\":\"c24467e727904fcc4e884bfcde103dae4bbb0990a705e41ad5c80ef9dcc89a3b\",\"txTimestamp\":\"seconds:1670338037 nanos:121729000\",\"createTime\":\"6066-126-126\",\"duedate\":\"2023-12-12\",\"guarantor\":\"ICBC\",\"description\":\"A new ticket for Org1MSP\"}