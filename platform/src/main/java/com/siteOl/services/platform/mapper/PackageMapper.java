package com.siteOl.services.platform.mapper;

import com.siteOl.services.platform.entity.Package;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Mapper;

/**
 * <p>
 * 功能套餐（租户可自主订购）- 代理可订购全部，机构订购代理已开通的功能（可拓展收费能力，字段预留）	 Mapper 接口
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-13
 */
@Mapper
public interface PackageMapper extends BaseMapper<Package> {

}
