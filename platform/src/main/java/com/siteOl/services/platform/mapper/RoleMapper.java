package com.siteOl.services.platform.mapper;

import com.siteOl.services.platform.entity.Role;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Mapper;

/**
 * <p>
 * 角色表，各租户下的内置或自定义角色 Mapper 接口
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-11-07
 */
@Mapper
public interface RoleMapper extends BaseMapper<Role> {

}
