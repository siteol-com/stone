package com.siteOl.services.platform.mapper;

import com.siteOl.services.platform.entity.RolePermission;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Mapper;

/**
 * <p>
 * 授权关系，角色与权限的关系 Mapper 接口
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-11-07
 */
@Mapper
public interface RolePermissionMapper extends BaseMapper<RolePermission> {

}
