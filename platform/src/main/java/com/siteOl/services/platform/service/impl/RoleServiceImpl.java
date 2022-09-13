package com.siteOl.services.platform.service.impl;

import com.siteOl.services.platform.entity.Role;
import com.siteOl.services.platform.mapper.RoleMapper;
import com.siteOl.services.platform.service.IRoleService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

/**
 * <p>
 * 内置角色表（超管专用）- 为各租户类型配置默认角色 服务实现类
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-13
 */
@Service
public class RoleServiceImpl extends ServiceImpl<RoleMapper, Role> implements IRoleService {

}
