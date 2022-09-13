package com.siteOl.services.platform.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import java.io.Serializable;
import java.math.BigDecimal;
import java.time.LocalDateTime;
import lombok.Getter;
import lombok.Setter;

/**
 * <p>
 * 功能套餐（租户可自主订购）- 代理可订购全部，机构订购代理已开通的功能（可拓展收费能力，字段预留）	
 * </p>
 *
 * @author 米虫@mebugs.com
 * @since 2022-09-13
 */
@Getter
@Setter
public class Package implements Serializable {

    private static final long serialVersionUID = 1L;

    /**
     * 功能套餐包ID
     */
    @TableId(value = "id", type = IdType.AUTO)
    private Long id;

    /**
     * 套餐包名称
     */
    private String name;

    /**
     * 套餐状态 0开放订购 1订购下线 2封存（存在订购不可封存）
     */
    private Integer status;

    /**
     * 参考定价（代理订购后可以自行修改，预留字段，用于自主订购）
     */
    private BigDecimal price;

    /**
     * 创建时间
     */
    private LocalDateTime createTime;

    /**
     * 更新时间
     */
    private LocalDateTime updateTime;


}
