package com.example.fabric.bean;

import lombok.Data;
import lombok.experimental.Accessors;

@Data
@Accessors(chain = true)
public class Account {
    String accountId;
    String userName;
    String balance;
}
