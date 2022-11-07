package com.siteOl.services.platform.service.impl;

import com.siteOl.services.platform.entity.RolePermission;
import com.siteOl.services.platform.mapper.RolePermissionMapper;
import com.siteOl.services.platform.service.IRolePermissionService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

/**
 * <p>
 * 授权关系，角色与权限的关系 服务实现类
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-11-07
 */
@Service
public class RolePermissionServiceImpl extends ServiceImpl<RolePermissionMapper, RolePermission> implements IRolePermissionService {

}
