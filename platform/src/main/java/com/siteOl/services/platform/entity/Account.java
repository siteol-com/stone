package com.siteOl.services.platform.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import java.io.Serializable;
import java.time.LocalDateTime;
import lombok.Getter;
import lombok.Setter;

/**
 * <p>
 * 基础登录账号
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-13
 */
@Getter
@Setter
public class Account implements Serializable {

    private static final long serialVersionUID = 1L;

    /**
     * 账号ID
     */
    @TableId(value = "id", type = IdType.AUTO)
    private Long id;

    /**
     * 账号
     */
    private String account;

    /**
     * 密文密码
     */
    private String encryption;

    /**
     * 盐值秘钥
     */
    private String saltKey;

    /**
     * 密码超期时间（修改后的90天）
     */
    private LocalDateTime pwdExpTime;

    /**
     * 租户ID
     */
    private Long tenantId;

    /**
     * 账号状态 0正常 1锁定 2封存
     */
    private Integer status;

    /**
     * 变更标识 0可变更 1禁止变更（除密码）
     */
    private Integer mark;

    /**
     * 创建时间
     */
    private LocalDateTime createTime;

    /**
     * 更新时间
     */
    private LocalDateTime updateTime;


}
