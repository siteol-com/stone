package com.siteOl.services.platform.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import java.io.Serializable;
import lombok.Getter;
import lombok.Setter;

/**
 * <p>
 * 租户类型下默认具备的权限
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-11-07
 */
@Getter
@Setter
@TableName("tenant_permission")
public class TenantPermission implements Serializable {

    private static final long serialVersionUID = 1L;

    /**
     * 数据ID
     */
    @TableId(value = "id", type = IdType.AUTO)
    private Long id;

    /**
     * 租户类型 0 超管（全局） 1 代理机构 2 企业 3 其他
     */
    private Long tenantType;

    /**
     * 权限ID
     */
    private Long permissionId;


}
