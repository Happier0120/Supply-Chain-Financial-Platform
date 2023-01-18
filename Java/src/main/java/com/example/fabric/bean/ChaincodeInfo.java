package com.example.fabric.bean;


import lombok.Data;
import lombok.experimental.Accessors;

@Data
@Accessors(chain = true)
public class ChaincodeInfo {
    String blockNumber;
    String nodesNumber;
    String transactionCount;
}
