package com.example.fabric.service;

import com.alibaba.fastjson.JSONObject;
import com.example.fabric.bean.*;
import com.example.fabric.call.ChaincodeCaller;
import com.example.fabric.constant.Submit;
import com.example.fabric.exception.BusinessException;
import com.google.protobuf.InvalidProtocolBufferException;
import org.hyperledger.fabric.gateway.ContractException;
import org.hyperledger.fabric.gateway.Transaction;
import org.hyperledger.fabric.sdk.BlockInfo;
import org.hyperledger.fabric.sdk.Channel;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.hyperledger.fabric.sdk.exception.ProposalException;

import java.nio.charset.StandardCharsets;
import java.util.*;
import java.util.concurrent.TimeoutException;

public class ChaincodeService {
    public static void main(String[] args) {
        new ChaincodeService().queryTickets();
        System.out.println();
    }

    // 查询区块链总览
    public ChaincodeInfo queryChaincodeInfo() {
        Channel channel = ChaincodeCaller.getNetwork().getChannel();

        // 获得区块数
        long blockNumber = 0;
        // 获得节点数
        int nodesNumber = 0;
        // 获得交易数
        long transactionCount = 0;

        try {
            blockNumber = channel.queryBlockchainInfo().getHeight();
            nodesNumber = channel.getPeers().size();
            byte[] currentBlockHash = channel.queryBlockchainInfo().getCurrentBlockHash();
            long currentBlockNumber = channel.queryBlockByHash(currentBlockHash).getBlockNumber();
            for (int i = 0; i <= currentBlockNumber; i++) {
                transactionCount += channel.queryBlockByNumber(i).getTransactionCount();
            }
        } catch (ProposalException | InvalidArgumentException e) {
            e.printStackTrace();
        }

        ChaincodeInfo chaincodeInfo = new ChaincodeInfo();
        chaincodeInfo.setBlockNumber(String.valueOf(blockNumber));
        chaincodeInfo.setNodesNumber(String.valueOf(nodesNumber));
        chaincodeInfo.setTransactionCount(String.valueOf(transactionCount));

        return chaincodeInfo;
    }

    // 查询区块信息
    public List<BlockListInfo> queryBlockInfoList() {
        Channel channel = ChaincodeCaller.getNetwork().getChannel();
        List<BlockListInfo> list = new ArrayList<>();
        try {
            byte[] currentBlockHash = channel.queryBlockchainInfo().getCurrentBlockHash();
            long currentBlockNumber = channel.queryBlockByHash(currentBlockHash).getBlockNumber();
            for (int i = 0; i <= currentBlockNumber; i++) {
                BlockListInfo info = new BlockListInfo();
                info.setTransactionCount(String.valueOf(channel.queryBlockByNumber(i).getTransactionCount()));
                info.setBlockNumber(String.valueOf(i));
                info.setDataHash(Base64.getEncoder().encodeToString(channel.queryBlockByNumber(i).getDataHash()));
                list.add(info);
            }
        } catch (ProposalException | InvalidArgumentException e) {
            e.printStackTrace();
        }
        return list;
    }

    // 查询单据列表
    public List<Ticket> queryTickets() {
        // 查询公开票据列表
        List<Ticket> list = queryTicketsPublic();
        if (list == null) {
            throw new BusinessException(10001, "查询无效");
        }

        List<Ticket> relist = new ArrayList<>();

        for (Ticket ticket : list) {
            // 如果不是核心企业，并且非当前企业，则不返回对应票据
            if (!Submit.coreUsers.contains(Submit.currentUser) && !ticket.getOwnerOrg().equals(Submit.currentUser)) {
                continue;
            }

            // 查询票据对应私有信息
            Ticket ticketPrivate = queryTicketPrivateById(ticket.getTicketID());
            ticket.setPrice(ticketPrivate.getPrice());

            // 查询票据对应时间
            String timestamp = queryTicketHistoryById(ticket.getTicketID()).get(0).getTimestamp();
            ticket.setCreateTime(timestamp.split("T")[0]);

            relist.add(ticket);
        }

        return relist;
    }

    // 查询公开单据列表
    public List<Ticket> queryTicketsPublic() {
        String resultString = "";
        try {
            byte[] resultBytes = ChaincodeCaller.getInstance().evaluateTransaction("GetAllTicket");
            resultString = new String(resultBytes);
        } catch (ContractException e) {
            System.err.println(e);
        }

        resultString = resultString.replaceAll("\\\\", "");
        resultString = resultString.replaceAll("\"\\{", "{");
        resultString = resultString.replaceAll("\\}\"", "}");
        List<Ticket> list = JSONObject.parseArray(resultString, Ticket.class);
        return list;
    }

    public Ticket queryTicketPublicById(String ticketID) {
        String resultString = "";
        try {
            byte[] resultBytes = ChaincodeCaller.getInstance().evaluateTransaction("QueryTicket", ticketID);
            resultString = new String(resultBytes);
        } catch (ContractException e) {
            System.err.println(e);
        }

        List<Ticket> list = JSONObject.parseArray(resultString, Ticket.class);
        Ticket ticket = list.get(0);
        return ticket;
    }

    public Ticket queryTicketPrivateById(String ticketID) {
        String resultString = null;
        try {
            byte[] resultBytes = ChaincodeCaller.getInstance().evaluateTransaction("QueryTicketPrivate", ticketID);
            resultString = new String(resultBytes);
        } catch (ContractException e) {
            System.err.println(e);
        }

        if (resultString == null) {
            return new Ticket();
        } else {
            resultString = resultString.split("field: ")[1];
            Ticket ticket = JSONObject.parseObject(resultString, Ticket.class);
            return ticket;
        }
    }

    public List<Record> queryTicketHistoryById(String ticketID) {
        String resultString = "";
        try {
            byte[] resultBytes = ChaincodeCaller.getInstance().evaluateTransaction("QueryTicketHistory", ticketID);
            resultString = new String(resultBytes);
        } catch (ContractException e) {
            System.err.println(e);
        }

        List<Record> list = JSONObject.parseArray(resultString, Record.class);

        if (list == null) {
            list = new ArrayList<>();
        }

        for (Record record : list) {
            Record blockInfo = queryBlockInfoByTxId(record.getTxId());
            record.setChannelId(blockInfo.getChannelId());
            record.setHash(blockInfo.getHash());
        }
        return list;
    }

    public Record queryBlockInfoByTxId(String txId) {
        Channel channel = ChaincodeCaller.getNetwork().getChannel();
        BlockInfo blockInfo;
        String channelId = null;
        String hashString = null;
        try {
            blockInfo = channel.queryBlockByTransactionID(txId);
            channelId = blockInfo.getChannelId();
            hashString = Base64.getEncoder().encodeToString(blockInfo.getDataHash());
        } catch (InvalidArgumentException | ProposalException | InvalidProtocolBufferException e) {
            e.printStackTrace();
        }

        Record record = new Record();
        record.setChannelId(channelId);
        record.setHash(hashString);
        return record;
    }

    public String createTicket(Ticket ticket) {
        String resultString = null;
        try {
            Transaction transaction = ChaincodeCaller.getInstance().createTransaction("CreateTicket");

            // 插入隐私数据
            Map<String, String> transientMap = new HashMap<>();
            transientMap.put("price", ticket.getPrice());
//            transientMap.put("fromOrder", ticket.getFromOrder());
            String transientMapString = JSONObject.toJSONString(transientMap);
            HashMap<String, byte[]> transientMapFormat = new HashMap<>();
            transientMapFormat.put("ticket_properties", transientMapString.getBytes(StandardCharsets.UTF_8));
            transaction.setTransient(transientMapFormat);

            // 插入公开数据
            byte[] result = transaction.submit(
                    // ticket.getOwnerOrg(),
                    ticket.getDescription());
            resultString = new String(result);
        } catch (ContractException e) {
            System.err.println(e);
        } catch (InterruptedException | TimeoutException e) {
            e.printStackTrace();
        }

        return resultString;
    }

    public String editTicketPublic(Ticket ticket) {
        String resultString = "";
        try {
            Transaction transaction = ChaincodeCaller.getInstance().createTransaction("ChangeDescription");

            //修改公开数据
            byte[] result = transaction.submit(
                    ticket.getTicketID(),
                    ticket.getDescription());
            resultString = new String(result);
        } catch (ContractException e) {
            System.err.println(e);
        } catch (InterruptedException | TimeoutException e) {
            e.printStackTrace();
        }

        return resultString;
    }

    public String transferTicket(Ticket ticket) {
        String resultString = null;
        try {
            Transaction transaction = ChaincodeCaller.getInstance().createTransaction("TransferTicket");

            // 修改隐私数据
            Map<String, String> transientMap = new HashMap<>();
            transientMap.put("price", ticket.getPrice());
            String transientMapString = JSONObject.toJSONString(transientMap);
            HashMap<String, byte[]> transientMapFormat = new HashMap<>();
            transientMapFormat.put("ticket_properties", transientMapString.getBytes(StandardCharsets.UTF_8));
            transaction.setTransient(transientMapFormat);

            // 修改公开数据
            byte[] result = transaction.submit(
                    ticket.getTicketID(),
                    ticket.getOwnerOrg());
            resultString = new String(result);
        } catch (ContractException e) {
            System.err.println(e);
        } catch (InterruptedException | TimeoutException e) {
            e.printStackTrace();
        }

        return resultString;
    }

    public List<Account> queryAccounts() {
        String resultString = "";
        try {
            byte[] resultBytes = ChaincodeCaller.getInstance().evaluateTransaction("QueryAccountInfor");
            resultString = new String(resultBytes);
        } catch (ContractException e) {
            System.err.println(e);
        }

        List<Account> list = JSONObject.parseArray(resultString, Account.class);
        return list;
    }

    public Account queryAccountById(String accountId) {
        String resultString = "";
        try {
            byte[] resultBytes = ChaincodeCaller.getInstance().evaluateTransaction("QueryAccountInfor", accountId);
            resultString = new String(resultBytes);
        } catch (ContractException e) {
            System.err.println(e);
        }

        List<Account> list = JSONObject.parseArray(resultString, Account.class);
        Account account = list.get(0);
        return account;
    }
}
