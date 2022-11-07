package com.siteOl.services.platform.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import java.io.Serializable;
import lombok.Getter;
import lombok.Setter;

/**
 * <p>
 * 权限表，可分配的基础权限结构体
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-11-07
 */
@Getter
@Setter
public class Permission implements Serializable {

    private static final long serialVersionUID = 1L;

    /**
     * 权限ID
     */
    @TableId(value = "id", type = IdType.AUTO)
    private Long id;

    /**
     * 权限名
     */
    private String name;

    /**
     * 权限别名
     */
    private String alias;

    /**
     * 权限类型（0平台/1业务）
     */
    private Integer type;

    /**
     * 权限等级 0 系统 1 业务 2 页面 3 操作 （2/3绑定路由）
     */
    private Integer level;

    /**
     * 权限父级ID（顶级为0）
     */
    private Long pid;

    /**
     * 权限状态 0正常 1锁定 2封存
     */
    private Integer status;


}
