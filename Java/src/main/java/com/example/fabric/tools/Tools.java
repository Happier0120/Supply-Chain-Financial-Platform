package com.example.fabric.tools;

import java.text.SimpleDateFormat;
import java.util.Date;

public class Tools {
    public static String getCurrentDate() {
        SimpleDateFormat sdf = new SimpleDateFormat();
        // sdf.applyPattern("yyyy-MM-dd HH:mm:ss");
        sdf.applyPattern("yyyy-MM-dd");
        Date date = new Date();
        String currentDate = sdf.format(date);
        return currentDate;
    }
}
