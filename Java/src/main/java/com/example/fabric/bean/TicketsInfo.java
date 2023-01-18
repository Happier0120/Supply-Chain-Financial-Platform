package com.example.fabric.bean;

import lombok.Data;
import lombok.experimental.Accessors;

@Data
@Accessors(chain = true)
public class TicketsInfo {
    // 总开立票据
    long ticketsNumber;
    // 总到期票据
    long expiredTicketsNumber;
    // 总开立金额
    long price;
    // 总到期金额
    long expiredPrice = 0;
    // 今日开立票据
    long todayTicketsNumber;
    // 今日到期票据
    long todayExpiredTicketsNumber;
    // 今日开立金额
    long todayPrice;
    // 今日到期金额
    long todayExpiredPrice = 0;
}
