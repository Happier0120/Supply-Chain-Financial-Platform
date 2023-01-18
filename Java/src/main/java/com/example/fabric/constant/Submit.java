package com.example.fabric.constant;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Set;

public class Submit {
    public static String currentUser = "org1";

    public static HashMap<String, String> users = new HashMap<>();
    public static Set<String> coreUsers = new HashSet<>();

    public static String channelName = "channel1";
    public static String contractName = "scf03";
    // public static String contractName = "tickets02";

    static {
        users.put("org1", "org1");
        users.put("org2", "org2");
        users.put("org3", "org3");

        coreUsers.add("org1");
    }
}
