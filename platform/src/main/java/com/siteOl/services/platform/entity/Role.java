package com.siteOl.services.platform.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import java.io.Serializable;
import java.time.LocalDateTime;
import lombok.Getter;
import lombok.Setter;

/**
 * <p>
 * 内置角色表（超管专用）- 为各租户类型配置默认角色
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-13
 */
@Getter
@Setter
public class Role implements Serializable {

    private static final long serialVersionUID = 1L;

    /**
     * 角色ID
     */
    @TableId(value = "id", type = IdType.AUTO)
    private Long id;

    /**
     * 角色名称（内置基础角色）
     */
    private String name;

    /**
     * 角色类型（归属租户类型） 0 超管（无需配置） 1 代理机构 2 企业
     */
    private Integer type;

    /**
     * 角色状态 0正常 1禁用 2封存
     */
    private Integer status;

    /**
     * 变更标识 0可变更 1禁止变更（超管角色无需变更）
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
