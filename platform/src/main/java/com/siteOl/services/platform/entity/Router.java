package com.siteOl.services.platform.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import java.io.Serializable;
import lombok.Getter;
import lombok.Setter;

/**
 * <p>
 * 系统路由表
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-13
 */
@Getter
@Setter
public class Router implements Serializable {

    private static final long serialVersionUID = 1L;

    /**
     * 路由ID
     */
    @TableId(value = "id", type = IdType.AUTO)
    private Long id;

    /**
     * 路由名称
     */
    private String name;

    /**
     * 路由路径
     */
    private String path;

    /**
     * 账号状态 0正常 1锁定 2封存
     */
    private Integer status;

    /**
     * 父级路由ID
     */
    private Long pid;


}