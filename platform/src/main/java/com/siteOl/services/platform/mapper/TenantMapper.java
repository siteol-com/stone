package com.siteOl.services.platform.mapper;

import com.siteOl.services.platform.entity.Tenant;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Mapper;

/**
 * <p>
 * 租户表 Mapper 接口
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-13
 */
@Mapper
public interface TenantMapper extends BaseMapper<Tenant> {

}
