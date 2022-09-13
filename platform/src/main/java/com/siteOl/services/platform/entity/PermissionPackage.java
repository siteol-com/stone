package com.siteOl.services.platform.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import java.io.Serializable;
import lombok.Getter;
import lombok.Setter;

/**
 * <p>
 * 授权关系，权限与套餐关系
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-13
 */
@Getter
@Setter
@TableName("permission_package")
public class PermissionPackage implements Serializable {

    private static final long serialVersionUID = 1L;

    /**
     * 关联ID
     */
    @TableId(value = "id", type = IdType.AUTO)
    private Long id;

    /**
     * 套餐ID
     */
    private Long packageId;

    /**
     * 权限ID
     */
    private Long permissionId;


}
