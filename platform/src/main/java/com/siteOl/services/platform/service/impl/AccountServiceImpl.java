package com.siteOl.services.platform.service.impl;

import com.siteOl.services.platform.entity.Account;
import com.siteOl.services.platform.mapper.AccountMapper;
import com.siteOl.services.platform.service.IAccountService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

/**
 * <p>
 * 基础登录账号 服务实现类
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-13
 */
@Service
public class AccountServiceImpl extends ServiceImpl<AccountMapper, Account> implements IAccountService {

}
