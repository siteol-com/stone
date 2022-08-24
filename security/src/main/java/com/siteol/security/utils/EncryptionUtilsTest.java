package com.siteol.security.utils;

import org.junit.Test;


public class EncryptionUtilsTest {

    /**
     * 测试加密
     */
    @Test
    public void encryptTest() {
        System.out.println(EncryptionUtils.encrypt("AdminAdminAdmin1","KEY_SiteOL_Stone"));
    }
    /**
     * 测试解密
     */
    @Test
    public void decryptTest() {
        System.out.println(EncryptionUtils.decrypt("VEmHXzPPPKzYZtY2w0tYHM3OGMA7jPZn50BermJ37Jc=","KEY_SiteOL_Stone"));
    }
}
