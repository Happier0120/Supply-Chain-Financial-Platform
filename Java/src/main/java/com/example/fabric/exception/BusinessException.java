package com.example.fabric.exception;

import com.example.fabric.httpdto.MsgCodeUtils;
import lombok.Data;

@Data
public class BusinessException extends RuntimeException {

    private int code;

    private String message;

    public BusinessException(int code, String message) {
        super();
        this.code = code;
        this.message = message;
    }

    public BusinessException(int code) {
        super();
        this.code = code;
        this.message = MsgCodeUtils.getCodeMessage(code);
    }

    public BusinessException(ExceptionEnum exceptionEnum) {
        super();
        this.code = exceptionEnum.getCode();
        this.message = exceptionEnum.getMessage();
    }


    public BusinessException(String message) {
        super();
        this.code = 20000;
        this.message = message;
    }







}