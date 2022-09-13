package com.siteOl.services.platform.mapper;

import com.siteOl.services.platform.entity.PermissionPackage;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Mapper;

/**
 * <p>
 * 授权关系，权限与套餐关系 Mapper 接口
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-13
 */
@Mapper
public interface PermissionPackageMapper extends BaseMapper<PermissionPackage> {

}
