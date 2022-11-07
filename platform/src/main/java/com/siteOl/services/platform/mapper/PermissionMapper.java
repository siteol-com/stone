package com.siteOl.services.platform.mapper;

import com.siteOl.services.platform.entity.Permission;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Mapper;

/**
 * <p>
 * 权限表，可分配的基础权限结构体 Mapper 接口
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-11-07
 */
@Mapper
public interface PermissionMapper extends BaseMapper<Permission> {

}
