package com.siteOl.security.utils;

import sun.misc.BASE64Decoder;
import sun.misc.BASE64Encoder;

import javax.crypto.Cipher;
import javax.crypto.spec.IvParameterSpec;
import javax.crypto.spec.SecretKeySpec;

/**
 *
 * AES加解密工具类
 * 使用AES-128-CBC加密模式，key需要为16位。
 * key 由方法传入(生成估计16位的随机数）
 *
 * @author 米虫@mebugs.com
 * @since 2022-08-24
 */
public class EncryptionUtils {

    private static String iv = "AES_SiteOL_Stone";

    // 加密
    public static String encrypt(String origData,String key) {
        try {
            Cipher cipher = Cipher.getInstance("AES/CBC/PKCS5Padding");
            byte[] raw = key.getBytes();
            SecretKeySpec keySpec = new SecretKeySpec(raw, "AES");
            IvParameterSpec ivSpec = new IvParameterSpec(iv.getBytes());//使用CBC模式，需要一个向量iv，可增加加密算法的强度
            cipher.init(Cipher.ENCRYPT_MODE, keySpec, ivSpec);
            byte[] encrypted = cipher.doFinal(origData.getBytes("utf-8"));
            // Base64 密文
            return new BASE64Encoder().encode(encrypted);
        } catch (Exception ex) {
            return null;
        }
    }

    // 解密
    public static String decrypt(String cryData,String key) {
        try {
            byte[] raw = key.getBytes("ASCII");
            SecretKeySpec keySpec = new SecretKeySpec(raw, "AES");
            Cipher cipher = Cipher.getInstance("AES/CBC/PKCS5Padding");
            IvParameterSpec ivSpec = new IvParameterSpec(iv.getBytes());
            cipher.init(Cipher.DECRYPT_MODE, keySpec, ivSpec);
            // Base64 解密
            byte[] encrypted1 = new BASE64Decoder().decodeBuffer(cryData);
            byte[] original = cipher.doFinal(encrypted1);
            String originalString = new String(original, "utf-8");
            return originalString;
        } catch (Exception ex) {
            return null;
        }
    }
}
