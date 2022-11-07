package com.siteOl.services.platform.mapper;

import com.siteOl.services.platform.entity.TenantPermission;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Mapper;

/**
 * <p>
 * 租户类型下默认具备的权限 Mapper 接口
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-11-07
 */
@Mapper
public interface TenantPermissionMapper extends BaseMapper<TenantPermission> {

}
