票据字段包括：

``` go
// 公开字段
ticket := &model.Ticket{
    Type:          ticketCollection,   
    ID:            ticketID,                // 随机生成10位16进制票据编号
    ChannelID:     channelID,               // 通道名称
    TxID:          txID,                    // 交易ID，Fabric是先执行再定序的，所以可以在交易成交之前获得交易ID
    TxTimestamp:   txTimestamp.String(),    // 交易时间戳，对应票据创建时间
    OwnerOrgMSPID: coreOrgMSPID,            // 所有者id，票据发行成功后自动对应所有为coreOrgMSPID
    Guarantor:     "ICBC",                  // 担保方
    CreateTime:    time.Now().Format("2006-01-02 15:04:05"),    // 票据创建时间，与TxTimestamp一致
    DueDate:       "2023-12-12",            // 到期时间，暂时写死
    Description:   description,             // 备注，唯一可更改字段
}

// 隐私字段: 现设计为value值为隐私字段，一个json格式，字段不限
```