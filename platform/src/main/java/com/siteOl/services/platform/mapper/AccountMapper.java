package com.siteOl.services.platform.mapper;

import com.siteOl.services.platform.entity.Account;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Mapper;

/**
 * <p>
 * 基础登录账号 Mapper 接口
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-13
 */
@Mapper
public interface AccountMapper extends BaseMapper<Account> {

}
