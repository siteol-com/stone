package com.siteOl.services.platform.service.impl;

import com.siteOl.services.platform.entity.Permission;
import com.siteOl.services.platform.mapper.PermissionMapper;
import com.siteOl.services.platform.service.IPermissionService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

/**
 * <p>
 * 权限表，可分配的基础权限结构体 服务实现类
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-11-07
 */
@Service
public class PermissionServiceImpl extends ServiceImpl<PermissionMapper, Permission> implements IPermissionService {

}
