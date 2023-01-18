package com.example.fabric.exception;

import com.baomidou.mybatisplus.core.toolkit.StringUtils;
import com.example.fabric.httpdto.ObjectDataResponse;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import java.util.Arrays;
import java.util.Enumeration;
import java.util.Map;

@ControllerAdvice
@Slf4j
public class GlobalExceptionHandler {



    @ExceptionHandler({BusinessException.class, Exception.class})
    @ResponseBody
    public ObjectDataResponse<Void> handleApiException(Exception e, HttpServletRequest request) {
        if (e instanceof BusinessException) {
            log.error("business error:{}", e.getMessage());
            BusinessException exception = (BusinessException) e;
            if (StringUtils.isNotEmpty(exception.getMessage())) {
                return new ObjectDataResponse<>(exception.getCode(), exception.getMessage(), null);
            }
            return ObjectDataResponse.errorResponse(exception.getCode());
        }
        log.error("internal error:", e);
        this.printSpecUrl(request);
        return new ObjectDataResponse<>(500, "internal error", null);
    }

    private void printSpecUrl(HttpServletRequest request) {
        if (request.getRequestURI().contains("undefined")) {
            Cookie[] cookies = request.getCookies();
            if (cookies != null) {
                Arrays.stream(cookies).forEach(c -> log.info("/meitiaolifestyle/api/index/tagArticleInfo/undefined cookie: " + c.getName() + "   " + c.getValue()));
            }
            request.getHeaderNames();
            Enumeration headerNames = request.getHeaderNames();
            while (headerNames.hasMoreElements()) {
                String key = (String) headerNames.nextElement();
                log.info(key + "=" + request.getHeader(key));
            }
            Map<String, String[]> map = request.getParameterMap();
            map.forEach((k, v) -> {
                log.info("/meitiaolifestyle/api/index/tagArticleInfo/undefined  param: " + k + "    " + Arrays.toString(v));
            });
        }
    }


}