package com.example.fabric.call;

import com.example.fabric.constant.Submit;
import org.hyperledger.fabric.gateway.Contract;
import org.hyperledger.fabric.gateway.Gateway;
import org.hyperledger.fabric.gateway.Network;

public class ChaincodeCaller {
    private static Contract contract;
    private static Network network;

    public static Contract getInstance() {
        return contract;
    }

    public static Network getNetwork() {
        return network;
    }


    public static void init(String orgName, String channelPara, String contractPara) {
        Submit.currentUser = orgName;
        Gateway gateway = null;
        try {
            gateway = ChaincodeConnect.connect();
        } catch (Exception e) {
            e.printStackTrace();
        }

        // get the network and contract
        network = gateway.getNetwork(channelPara);
        contract = network.getContract(contractPara);
    }
}
