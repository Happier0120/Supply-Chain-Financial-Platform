package com.example.fabric.bean;

import lombok.Data;
import lombok.experimental.Accessors;

@Data
@Accessors(chain = true)
public class BlockListInfo {
    String transactionCount;
    String dataHash;
    String blockNumber;
}
