package com.siteOl.services.platform.mapper;

import com.siteOl.services.platform.entity.PermissionRouter;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Mapper;

/**
 * <p>
 * 授权关系，权限与路由的关系 Mapper 接口
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-11-07
 */
@Mapper
public interface PermissionRouterMapper extends BaseMapper<PermissionRouter> {

}
