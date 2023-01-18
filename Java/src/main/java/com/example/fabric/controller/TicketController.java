package com.example.fabric.controller;


import com.example.fabric.bean.*;
import com.example.fabric.call.ChaincodeCaller;
import com.example.fabric.constant.SearchType;
import com.example.fabric.constant.Tickets;
import com.example.fabric.constant.Submit;
import com.example.fabric.service.ChaincodeService;
import com.example.fabric.httpdto.ObjectCollectionResponse;
import com.example.fabric.httpdto.ObjectDataResponse;
import com.example.fabric.tools.Tools;
import io.swagger.annotations.ApiOperation;
import org.springframework.web.bind.annotation.*;
import io.swagger.annotations.Api;

import javax.validation.Valid;
import java.util.*;

@RequestMapping("/console")
@RestController
@CrossOrigin(origins = "*", maxAge = 3600)
@Api("")
public class TicketController {
    ChaincodeService chaincodeService;

    {
        chaincodeService = new ChaincodeService();
    }

    @PostMapping("submit")
    @ApiOperation("登录")
    public ObjectDataResponse<String> createTicket(@RequestBody() @Valid User user) {
        if (!"123456".equals(user.getPassword())) {
            return new ObjectDataResponse<>("-1");
        }

        Submit.currentUser = user.getOrgName();
        ChaincodeCaller.init(Submit.currentUser, Submit.channelName, Submit.contractName);

        if (Submit.currentUser != "") {
            if (Submit.coreUsers.contains(Submit.currentUser)) {
                return new ObjectDataResponse<>("0");
            } else {
                return new ObjectDataResponse<>("1");
            }
        } else {
            return new ObjectDataResponse<>("-1");
        }
    }

    @GetMapping("queryBlockInfoList")
    @ApiOperation("查询区块信息列表")
    public ObjectCollectionResponse<BlockListInfo> queryBlockInfoList() {
        List<BlockListInfo> blockListInfos = chaincodeService.queryBlockInfoList();
        return new ObjectCollectionResponse<>(blockListInfos);
    }


    @GetMapping("queryChaincodeInfo")
    @ApiOperation("查询区块链总览")
    public ObjectDataResponse<ChaincodeInfo> queryChaincodeInfo() {
        ChaincodeInfo chaincodeInfo = chaincodeService.queryChaincodeInfo();
        return new ObjectDataResponse<>(chaincodeInfo);
    }

    @GetMapping("queryTicketsInfo")
    @ApiOperation("查询单据总览")
    public ObjectDataResponse<TicketsInfo> queryTicketsInfo() {
        // 查询票据
        List<Ticket> list = chaincodeService.queryTickets();

        // 总开立票据
        long ticketsNumber = list.size();
        // 总到期票据
        long expiredTicketsNumber = 0;
        // 总开立金额
        long price = 0;
        // 总到期金额
        long expiredPrice = 0;
        // 今日开立票据
        long todayTicketsNumber = 0;
        // 今日到期票据
        long todayExpiredTicketsNumber = 0;
        // 今日开立金额
        long todayPrice = 0;
        // 今日到期金额
        long todayExpiredPrice = 0;


        String currentDate = Tools.getCurrentDate();
        for (Ticket ticket : list) {
            // 判断已到期票据
            if (currentDate.compareTo(ticket.getDuedate()) > 0) {
                expiredTicketsNumber += 1;
                expiredPrice += Integer.valueOf(ticket.getPrice());
            }
            // 判断今日到期票据
            if (currentDate.compareTo(ticket.getDuedate()) == 0) {
                todayExpiredTicketsNumber += 1;
                todayExpiredPrice += Integer.valueOf(ticket.getPrice());
            }
            // 判断今日创建票据，获得今日金额
            if (currentDate.compareTo(ticket.getCreateTime()) == 0) {
                todayTicketsNumber += 1;
                try {
                    todayPrice += Integer.valueOf(ticket.getPrice());
                } catch (Exception e) {
                }
            }
            // 查询总金额
            try {
                price += Integer.valueOf(ticket.getPrice());
            } catch (Exception e) {
            }
        }


        TicketsInfo ticketsInfo = new TicketsInfo();
        ticketsInfo.setTicketsNumber(ticketsNumber);
        ticketsInfo.setExpiredTicketsNumber(expiredTicketsNumber);
        ticketsInfo.setPrice(price);
        ticketsInfo.setExpiredPrice((expiredPrice));
        ticketsInfo.setTodayTicketsNumber(todayTicketsNumber);
        ticketsInfo.setTodayExpiredTicketsNumber(todayExpiredTicketsNumber);
        ticketsInfo.setTodayPrice(todayPrice);
        ticketsInfo.setTodayExpiredPrice((todayExpiredPrice));

        return new ObjectDataResponse<>(ticketsInfo);
    }

    @GetMapping("queryTickets")
    @ApiOperation("查询票据列表")
    public ObjectCollectionResponse<Ticket> queryTickets(@Valid Search search) {
        List<Ticket> list = chaincodeService.queryTickets();

        // 根据查询类型筛选
        List<Ticket> typeList = new ArrayList<>();
        if (SearchType.notExpired.equals(search.getType())) {
            String currentDate = Tools.getCurrentDate();
            for (Ticket ticket : list) {
                if (currentDate.compareTo(ticket.getDuedate()) < 0) {
                    typeList.add(ticket);
                }
            }
        } else if (SearchType.expired.equals(search.getType())) {
            String currentDate = Tools.getCurrentDate();
            for (Ticket ticket : list) {
                if (currentDate.compareTo(ticket.getDuedate()) >= 0) {
                    typeList.add(ticket);
                }
            }

        } else if (SearchType.transferred.equals(search.getType())) {
            for (Ticket ticket : list) {
                int size = chaincodeService.queryTicketHistoryById(ticket.getTicketID()).size();
                if (size > 1) {
                    ticket.setType(Tickets.transferred);
                    typeList.add(ticket);
                }
            }
        } else if (SearchType.notTransferred.equals(search.getType())) {
            for (Ticket ticket : list) {
                int size = chaincodeService.queryTicketHistoryById(ticket.getTicketID()).size();
                if (size == 1) {
                    ticket.setType(Tickets.create);
                    typeList.add(ticket);
                }
            }
        }

        // 根据查询条件筛选
        Iterator<Ticket> iterator = typeList.iterator();
        if (search.getTicketID() != null && !"".equals(search.getTicketID())) {
            while (iterator.hasNext()) {
                Ticket ticket = iterator.next();
                if (!search.getTicketID().equals(ticket.getTicketID())) {
                    iterator.remove();
                }
            }
        }
        if (search.getCreateTime() != null && !"".equals(search.getCreateTime())) {
            while (iterator.hasNext()) {
                Ticket ticket = iterator.next();
                if (!search.getCreateTime().equals(ticket.getCreateTime())) {
                    iterator.remove();
                }
            }
        }
        if (search.getDuedate() != null && !"".equals(search.getDuedate())) {
            while (iterator.hasNext()) {
                Ticket ticket = iterator.next();
                if (!search.getDuedate().equals(ticket.getDuedate())) {
                    iterator.remove();
                }
            }
        }

        return new ObjectCollectionResponse<>(typeList);
    }

    @GetMapping("queryNotExpiredTickets")
    @ApiOperation("查询未到期票据列表")
    public ObjectCollectionResponse<Ticket> queryNotExpiredTickets() {
        List<Ticket> list = chaincodeService.queryTickets();

        List<Ticket> notExpiredTicketsList = new ArrayList<>();
        String currentDate = Tools.getCurrentDate();
        for (Ticket ticket : list) {
            if (currentDate.compareTo(ticket.getDuedate()) < 0) {
                notExpiredTicketsList.add(ticket);
            }
        }

        return new ObjectCollectionResponse<>(notExpiredTicketsList);
    }

    @GetMapping("queryExpiredTickets")
    @ApiOperation("查询已到期票据列表")
    public ObjectCollectionResponse<Ticket> queryExpiredTickets() {
        List<Ticket> list = chaincodeService.queryTickets();

        List<Ticket> expiredTicketsList = new ArrayList<>();
        String currentDate = Tools.getCurrentDate();
        for (Ticket ticket : list) {
            if (currentDate.compareTo(ticket.getDuedate()) >= 0) {
                expiredTicketsList.add(ticket);
            }
        }

        return new ObjectCollectionResponse<>(expiredTicketsList);
    }

    @GetMapping("queryTransferredTickets")
    @ApiOperation("查询已流转票据列表")
    public ObjectCollectionResponse<Ticket> queryTransferredTickets() {
        List<Ticket> list = chaincodeService.queryTickets();

        List<Ticket> transferredTicketsList = new ArrayList<>();
        for (Ticket ticket : list) {
            int size = chaincodeService.queryTicketHistoryById(ticket.getTicketID()).size();
            if (size > 1) {
                ticket.setType(Tickets.transferred);
                transferredTicketsList.add(ticket);
            }
        }

        return new ObjectCollectionResponse<>(transferredTicketsList);
    }

    @GetMapping("queryNotTransferredTickets")
    @ApiOperation("查询未流转票据列表")
    public ObjectCollectionResponse<Ticket> queryNotTransferredTickets() {
        List<Ticket> list = chaincodeService.queryTickets();

        List<Ticket> notTransferredTicketsList = new ArrayList<>();
        for (Ticket ticket : list) {
            int size = chaincodeService.queryTicketHistoryById(ticket.getTicketID()).size();
            if (size == 1) {
                ticket.setType(Tickets.create);
                notTransferredTicketsList.add(ticket);
            }
        }

        return new ObjectCollectionResponse<>(notTransferredTicketsList);
    }

    @PostMapping("createTicket")
    @ApiOperation("新增票据")
    public ObjectDataResponse<Boolean> createTicket(@RequestBody() @Valid Ticket ticket) {
        String re = chaincodeService.createTicket(ticket);

        if (re != null) {
            return new ObjectDataResponse<>(true);
        } else {
            return new ObjectDataResponse<>(false);
        }
    }

    @GetMapping("queryTicketPublicById")
    @ApiOperation("通过ID查询票据公开信息")
    public ObjectDataResponse<Ticket> queryTicketPublicById(@RequestParam() String ticketID) {
        Ticket ticket = chaincodeService.queryTicketPublicById(ticketID);
        return new ObjectDataResponse<>(ticket);
    }

    @GetMapping("queryTicketPrivateById")
    @ApiOperation("通过ID查询票据私有信息")
    public ObjectDataResponse<Ticket> queryTicketPrivateById(@RequestParam() String ticketID) {
        Ticket ticket = chaincodeService.queryTicketPrivateById(ticketID);
        return new ObjectDataResponse<>(ticket);
    }

    @GetMapping("queryTicketHistoryById")
    @ApiOperation("通过ID查询票据流转历史")
    public ObjectCollectionResponse<Record> queryTicketHistoryById(@RequestParam() String ticketID) {
        List<Record> list = chaincodeService.queryTicketHistoryById(ticketID);
        return new ObjectCollectionResponse<>(list);
    }

    @PostMapping("editTicketPublic")
    @ApiOperation("修改票据公开信息")
    public ObjectDataResponse<Boolean> editTicketPublic(@RequestBody() @Valid Ticket ticket) {
        String re = chaincodeService.editTicketPublic(ticket);

        if (re != null) {
            return new ObjectDataResponse<>(true);
        } else {
            return new ObjectDataResponse<>(false);
        }
    }

    @PostMapping("transferTicket")
    @ApiOperation("流转票据")
    public ObjectDataResponse<Boolean> transferTicket(@RequestBody() @Valid Ticket ticket) {
        String re = chaincodeService.transferTicket(ticket);

        if (re != null) {
            return new ObjectDataResponse<>(true);
        } else {
            return new ObjectDataResponse<>(false);
        }
    }

//    @GetMapping("queryAccounts")
//    @ApiOperation("查询账户列表")
//    public ObjectCollectionResponse<Account> queryAccounts() {
//        List<Account> list = chaincodeService.queryAccounts();
//        return new ObjectCollectionResponse<>(list);
//    }
//
//    @GetMapping("queryAccountById")
//    @ApiOperation("通过ID查询账户")
//    public ObjectDataResponse<Account> queryTicketById(@RequestParam() String accountId) {
//        Account account = chaincodeService.queryAccountById(accountId);
//        return new ObjectDataResponse<>(account);
//    }
}
