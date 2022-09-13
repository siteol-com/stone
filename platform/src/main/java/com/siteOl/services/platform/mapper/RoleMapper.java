package com.siteOl.services.platform.mapper;

import com.siteOl.services.platform.entity.Role;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Mapper;

/**
 * <p>
 * 内置角色表（超管专用）- 为各租户类型配置默认角色 Mapper 接口
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-13
 */
@Mapper
public interface RoleMapper extends BaseMapper<Role> {

}
