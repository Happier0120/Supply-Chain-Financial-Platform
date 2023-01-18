package com.example.fabric.exception;

import lombok.Getter;


public enum ExceptionEnum {
        LOGIN_FIRST(401, "请先登录");
        @Getter
        private int    code;
        @Getter
        private String message;

        ExceptionEnum(int code, String message) {
            this.code = code;
            this.message = message;
        }
}
