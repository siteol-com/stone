package com.siteOl.utils.generator;

import com.baomidou.mybatisplus.generator.AutoGenerator;
import com.baomidou.mybatisplus.generator.config.*;

import java.io.File;
import java.util.Collections;

/**
 *
 * 代码生成器
 * 调用版/界面版（部分方法支持调用）
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-06
 */
public class CodeMakeUtils {

    /**
     * 入口调用方式
     * 用于不启动前端快速生成
     * @param args
     */
    public static void main(String[] args) {
        // JDBC Config
        String url = "jdbc:mysql://localhost:3306/platform?useUnicode=true&characterEncoding=UTF-8&zeroDateTimeBehavior=convertToNull&useJDBCCompliantTimezoneShift=true&useLegacyDatetimeCode=false&serverTimezone=Asia/Shanghai";
        String userName = "root";
        String password = "123456";
        String projectName = "platform";
        String basePackage = "com.siteOl.services";
        String module = "platform";
        String author = "米虫@mebugs.com";
        String[] tables =  new String[]{
                "account",
                "permission",
                "permission_router",
                "role",
                "role_permission",
                "router",
                "tenant",
                "tenant_permission"
        };
        // 调用代码生成
        codeMaker(url,userName,password,projectName,basePackage,module,author,tables);
    }

    /**
     * 代码生成器
     * @param url           数据库地址
     * @param userName      数据库账号
     * @param password      数据库密码
     * @param projectName   工程名
     * @param basePackage   基础包名
     * @param module        模块名
     * @param author        作者
     * @param tables        表
     */
    public static void codeMaker(String url,String userName,String password,String projectName,String basePackage,String module,String author,String[] tables) {
        // 数据库基本连接信息配置
        DataSourceConfig dataSourceConfig = new DataSourceConfig.Builder(url,userName,password).build();
        // 读取目标工程目标路径
        String projectPath = System.getProperty("user.dir") + File.separator + projectName + File.separator + "src"+ File.separator + "main"+ File.separator ;
        // 公共参数配置
        String codePath = projectPath+"java";
        GlobalConfig globalConfig = new GlobalConfig.Builder().outputDir(codePath).author(author).build();
        // 父包名（未传取默认值）
        basePackage = (basePackage == null || ("").equals(basePackage)) ? "com.siteOl.services" : basePackage;
        // 包配置
        String mapperPath = projectPath+"resources"+ File.separator + "mapper" + File.separator + module;
        PackageConfig packageConfig = new PackageConfig.Builder().parent(basePackage).moduleName(module).pathInfo(Collections.singletonMap(OutputFile.xml, mapperPath)).build();
        // 模板配置（使用MyBatis-Plus自带的）
        // 代码配置
        StrategyConfig strategyConfig = new StrategyConfig.Builder().addInclude(tables)// 指明生码表
                .entityBuilder().enableLombok()// 实体类配置：开启lombok
                .controllerBuilder().enableRestStyle()// 控制层配置：开启Rest注解
                .serviceBuilder()// 业务层配置：
                .mapperBuilder().enableMapperAnnotation()// Mapper层配置：开启Mapper注解
                .build();
        // 组装配置
        AutoGenerator autoGenerator = new AutoGenerator(dataSourceConfig).global(globalConfig).packageInfo(packageConfig).strategy(strategyConfig);
        // 执行生成
        autoGenerator.execute();
    }
}
