SOA:
	1.long to short
		完整的url 经过算法 变成 几个字母
		返回完整的短url, 前缀放在数据库里
		判断
			a.是否存在
	2.short to long

	3.建表语句
		CREATE TABLE `base_redirect` (
			`redirect_id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '短网址唯一id,自增长',
			`long_url` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '原始url',
			`short_url` CHAR(25) NOT NULL DEFAULT '' COMMENT '短url',
			`long_crc` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '原始url crc',
			`short_crc` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '短url crc',
			`status` TINYINT(4) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态 0:删除 1:正常',
			`created_by_ip` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者ip',
			`updated_by_ip` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者ip',
			`created_at` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间timestamp',
			`updated_at` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间timestamp',
			PRIMARY KEY (`redirect_id`),
			KEY `long_crc` (`long_crc`),
			KEY `short_url` (`short_url`)
		) ENGINE = INNODB DEFAULT CHARSET=utf8 COMMENT='短网址表';

	4.redis
		全部走redis

技术疑问:
	0.路由					  			 √
		namespace		  				√
		疑问								√
		NSBefore 
		NSAfter
		路由域名						√
	1.参数以及验证
		验证json的合法性 						√
		多维的meta data 					  √
		json相关知识 				 		 √
		dataParams为数组: 定义struct  		 √
		接受的参数							√
		参数的验证							√
		映射到struct 							√
		struct里的验证 							√
		继承数组验证 								√
		错误机制的封装 						√
		Content-Type验证					√
2.baseConrtoller						√
3.filter
签名验证
goworker
5.mysql
6.redis
7.日志
	基于队列 RabbitMQ
	MongoDB
8.错误页面
9.代码提示 							√
10.封装返回 						√
11.ffjson							√
