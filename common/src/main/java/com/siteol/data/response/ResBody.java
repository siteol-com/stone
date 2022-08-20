package com.siteol.data.response;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;
import lombok.experimental.Accessors;

/**
 *
 * 统一数据JSON返回结构
 *
 *
 * @author 米虫@mebugs.com
 * @since 2022-08-16
 */
@Setter
@Getter
@ToString
@Accessors(chain = true)
public class ResBody {
    private static final long serialVersionUID = 1L;
    private int code = 200;
    private String msg;
    private Object data;

    // 404 路由获取失败
    public static ResBody routerErr() {
        return jsonResult(404, "Router Err", null);
    }

    // 403 鉴权失败
    public static ResBody AuthErr() {
        return jsonResult(403, "Auth Err", null);
    }

    // Json默认OK成功
    public static ResBody JsonOK() {
        return jsonResult(200, "OK", null);
    }

    // Json数据返回
    public static ResBody JsonSuccess(String msg, Object data) {
        return jsonResult(200, msg, data);
    }

    // Json失败返回
    public static ResBody JsonFail(String msg) {
        return jsonResult(500,msg,null);
    }

    // Json错误返回
    public static ResBody JsonError(Exception e) {
        return jsonResult(500,e.getMessage(),null);
    }

    // 公共调用
    public static ResBody jsonResult(int code, String msg, Object data) {
        ResBody res = new ResBody();
        res.setCode(code);
        res.setMsg(msg);
        res.setData(data);
        return res;
    }
}
