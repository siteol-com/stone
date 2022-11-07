package com.siteOl.services.platform.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import java.io.Serializable;
import java.time.LocalDateTime;
import lombok.Getter;
import lombok.Setter;

/**
 * <p>
 * 租户表
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-11-07
 */
@Getter
@Setter
public class Tenant implements Serializable {

    private static final long serialVersionUID = 1L;

    /**
     * 租户ID
     */
    @TableId(value = "id", type = IdType.AUTO)
    private Long id;

    /**
     * 租户名称
     */
    private String name;

    /**
     * 租户英文别名（访问URL）
     */
    private String alias;

    /**
     * 租户类型 0 超管 1 代理机构 2 企业 3 其他
     */
    private Integer type;

    /**
     * 租户主题（预留字段）
     */
    private Integer theme;

    /**
     * 租户Logo
     */
    private String logo;

    /**
     * 租户背景CSS（图片或颜色）
     */
    private String background;

    /**
     * 账号状态 0正常 1锁定 2封存
     */
    private Integer status;

    /**
     * 变更标识 0可变更 1禁止变更
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
