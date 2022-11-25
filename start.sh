#!/bin/bash
echo "一、清理环境，重启网络并创建通道mychannel"
./network.sh down
./network.sh up createChannel

echo "二、安装链码"
./network.sh deployCC -ccn chaincode01 -ccp ../../chaincode/ -ccl go -ccep "OR('Org1MSP.peer','Org2MSP.peer')"

echo "三、使用or1环境"
export PATH=${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=localhost:7051

echo "四、初始化链码'"
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n chaincode01 -c '{"function":"Init","Args":[]}'
sleep 5
echo "当前账户列表"
peer chaincode query -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n chaincode01 -c '{"function":"QueryAccountInfor","Args":[""]}'

echo "五、创建一个资产"
export Ticket_PROPERTIES=$(echo -n "{\"price\":\"101\",\"fromOrder\":\"ccccit\"}" | base64 | tr -d \\n)
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n chaincode01 -c '{"function":"CreateTicket","Args":["ticket01", "A new ticket for Org1MSP"]}' --transient "{\"ticket_properties\":\"$Ticket_PROPERTIES\"}"
sleep 5

echo "六、查询票据"
echo "6.1 票据公开数据"
peer chaincode query -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n chaincode01 -c '{"function":"QueryTicket","Args":["ticket01"]}'
echo "6.2 票据隐私数据"
peer chaincode query -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n chaincode01 -c '{"function":"ReadTicketPrivateProperties","Args":["ticket01"]}'

echo "七、 更改票据备注"
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n chaincode01 -c '{"function":"ChangeDescription","Args":["ticket01","new01"]}'
sleep 5

echo "八、 转让票据"
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n chaincode01 -c '{"function":"TransferTicket","Args":["ticket01","Org2MSP"]}' --transient "{\"ticket_properties\":\"$Ticket_PROPERTIES\"}" --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt"
sleep 5
echo "转让后票据公开信息依然可查看"
peer chaincode query -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n chaincode01 -c '{"function":"QueryTicket","Args":["ticket01"]}'
echo "转让后票据隐私信息无法再查看"
peer chaincode query -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n chaincode01 -c '{"function":"ReadTicketPrivateProperties","Args":["ticket01"]}'

echo "九、切换org2环境"
export PATH=${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=localhost:9051
echo "只有coreOrg可创建票据，org2无法创建票据"
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n chaincode01 -c '{"function":"CreateTicket","Args":["ticket01", "A new ticket for Org1MSP"]}' --transient "{\"ticket_properties\":\"$Ticket_PROPERTIES\"}"
echo "org2可查看接收到票据的隐私信息"
peer chaincode query -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n chaincode01 -c '{"function":"ReadTicketPrivateProperties","Args":["ticket01"]}'

echo "十、对票据交易全流程溯源"
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n chaincode01 -c '{"function":"QueryTicketHistory","Args":["ticket01"]}'


# echo "清理环境"
# ./network.sh down
