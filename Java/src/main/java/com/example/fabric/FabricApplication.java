package com.example.fabric;

import com.example.fabric.call.ChaincodeCaller;
import com.example.fabric.constant.Submit;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.jdbc.DataSourceAutoConfiguration;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

@SpringBootApplication(exclude = {DataSourceAutoConfiguration.class})
/*@MapperScan("com.example.**.mapper")*/
@EnableSwagger2
public class FabricApplication {

    public static void main(String[] args) {
        SpringApplication.run(FabricApplication.class, args);
        System.out.println("init start");
        if (args.length >= 1) {
            ChaincodeCaller.init(Submit.currentUser, Submit.channelName, args[0]);
        } else {
            ChaincodeCaller.init(Submit.currentUser, Submit.channelName, Submit.contractName);
        }
        System.out.println("init end");
    }

}
