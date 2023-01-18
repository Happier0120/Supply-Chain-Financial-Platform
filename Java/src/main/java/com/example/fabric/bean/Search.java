package com.example.fabric.bean;

import lombok.Data;
import lombok.experimental.Accessors;

import javax.validation.constraints.Pattern;

@Data
@Accessors(chain = true)
public class Search {
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String type;
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String ticketID;
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String createTime;
    @Pattern(regexp = "^[A-Za-z0-9_]*$", message = "只能使用字母、数字、下划线")
    String duedate;
}
