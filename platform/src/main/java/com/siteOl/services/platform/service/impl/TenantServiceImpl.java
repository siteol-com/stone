package com.siteOl.services.platform.service.impl;

import com.siteOl.services.platform.entity.Tenant;
import com.siteOl.services.platform.mapper.TenantMapper;
import com.siteOl.services.platform.service.ITenantService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

/**
 * <p>
 * 租户表 服务实现类
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-11-07
 */
@Service
public class TenantServiceImpl extends ServiceImpl<TenantMapper, Tenant> implements ITenantService {

}
