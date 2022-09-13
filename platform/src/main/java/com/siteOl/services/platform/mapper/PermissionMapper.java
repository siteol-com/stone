package com.siteOl.services.platform.mapper;

import com.siteOl.services.platform.entity.Permission;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Mapper;

/**
 * <p>
 * 权限表（基础权限为角色赋予/业务权限由套餐处理） Mapper 接口
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-13
 */
@Mapper
public interface PermissionMapper extends BaseMapper<Permission> {

}
