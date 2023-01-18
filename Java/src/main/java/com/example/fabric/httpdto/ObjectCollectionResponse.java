package com.example.fabric.httpdto;

import com.baomidou.mybatisplus.core.metadata.IPage;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.Collection;

@Data
@NoArgsConstructor
public class ObjectCollectionResponse<T> extends BaseResponse {
    private Collection<T> data;
    private long total;
    private long size;

    public ObjectCollectionResponse(Collection<T> data, long total, long size) {
        super(200, "OK");
        this.data = data;
        this.total = total;
        this.size = size;
    }

    @SuppressWarnings("unchecked")
    public ObjectCollectionResponse(IPage page) {
        super(200, "OK");
        this.data = page.getRecords();
        this.total = page.getTotal();
        this.size = page.getSize();
    }

    public ObjectCollectionResponse(Collection<T> data) {
        super(200, "OK");
        this.data = data;
        if (data != null) {
            this.total = data.size();
            try {
                this.size = Integer.valueOf(String.valueOf(this.total));
            } catch (Exception e) {
            }
        }
    }

    @Deprecated
    public static <T> ObjectCollectionResponse<T> errorResponse(int code, String msg) {

        ObjectCollectionResponse<T> response = new ObjectCollectionResponse<>();
        response.setCode(code);
        response.setMessage(msg);
        response.data = null;
        response.total = 0L;
        return response;
    }

    @Deprecated
    public static <T> ObjectCollectionResponse<T> errorResponse(int code) {
        return errorResponse(code, MsgCodeUtils.getCodeMessage(code));
    }


}
