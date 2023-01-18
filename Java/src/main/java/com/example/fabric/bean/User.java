package com.example.fabric.bean;


import lombok.Data;
import lombok.experimental.Accessors;

import javax.validation.constraints.Pattern;

@Data
@Accessors(chain = true)
public class User {
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String orgName;
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String password;
}
