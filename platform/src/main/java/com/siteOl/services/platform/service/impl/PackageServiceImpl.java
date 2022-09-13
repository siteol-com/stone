package com.siteOl.services.platform.service.impl;

import com.siteOl.services.platform.entity.Package;
import com.siteOl.services.platform.mapper.PackageMapper;
import com.siteOl.services.platform.service.IPackageService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

/**
 * <p>
 * 功能套餐（租户可自主订购）- 代理可订购全部，机构订购代理已开通的功能（可拓展收费能力，字段预留）	 服务实现类
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-13
 */
@Service
public class PackageServiceImpl extends ServiceImpl<PackageMapper, Package> implements IPackageService {

}
