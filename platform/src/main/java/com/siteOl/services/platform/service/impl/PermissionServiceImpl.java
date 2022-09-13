package com.siteOl.services.platform.service.impl;

import com.siteOl.services.platform.entity.Permission;
import com.siteOl.services.platform.mapper.PermissionMapper;
import com.siteOl.services.platform.service.IPermissionService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

/**
 * <p>
 * 权限表（基础权限为角色赋予/业务权限由套餐处理） 服务实现类
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-13
 */
@Service
public class PermissionServiceImpl extends ServiceImpl<PermissionMapper, Permission> implements IPermissionService {

}
