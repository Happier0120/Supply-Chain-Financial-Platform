package com.example.fabric.bean;

import lombok.Data;
import lombok.experimental.Accessors;

import javax.validation.constraints.Pattern;

@Data
@Accessors(chain = true)
public class Ticket {
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String ticketID;
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String objectType;
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String ownerOrg;
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String channelID;
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String txID;
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String txTimestamp;
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String createTime;
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String duedate;
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String guarantor;
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String description;

    @Pattern(regexp = "^[0-9]*$", message = "只能使用数字")
    String price;
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String fromOrder;

    // 流转票据额外字段
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String type;
}