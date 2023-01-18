package com.example.fabric.bean;

import lombok.Data;
import lombok.experimental.Accessors;

@Data
@Accessors(chain = true)
public class Record {
    String txId;
    String timestamp;
    Ticket record;

    String channelId;
    String hash;
}