package com.siteOl.services.platform.service.impl;

import com.siteOl.services.platform.entity.TenantPermission;
import com.siteOl.services.platform.mapper.TenantPermissionMapper;
import com.siteOl.services.platform.service.ITenantPermissionService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

/**
 * <p>
 * 租户类型下默认具备的权限 服务实现类
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-11-07
 */
@Service
public class TenantPermissionServiceImpl extends ServiceImpl<TenantPermissionMapper, TenantPermission> implements ITenantPermissionService {

}
