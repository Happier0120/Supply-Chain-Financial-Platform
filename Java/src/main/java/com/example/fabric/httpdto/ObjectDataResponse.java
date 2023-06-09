package com.example.fabric.httpdto;

import lombok.Data;
import lombok.EqualsAndHashCode;

@EqualsAndHashCode(callSuper = true)
@Data
public class ObjectDataResponse<T> extends BaseResponse {

    private T data;

    public ObjectDataResponse() {
        super(200, "OK");
        this.data = null;
    }

    public ObjectDataResponse(T data) {
        super(200, "OK");
        this.data = data;
    }

    public ObjectDataResponse(int code, String message, T data) {
        super(code, message);
        this.data = data;
    }

    @Deprecated
    public static <T> ObjectDataResponse<T> errorResponse(int code, String msg) {
        ObjectDataResponse<T> response = new ObjectDataResponse<>();
        response.setCode(code).setMessage(msg);
        response.data = null;
        return response;
    }

    @Deprecated
    public static <T> ObjectDataResponse<T> errorResponse(int code) {
        return errorResponse(code, MsgCodeUtils.getCodeMessage(code));
    }


}