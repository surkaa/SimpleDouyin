CREATE TABLE `user`
(
    `id`        bigint(20)                                              NOT NULL AUTO_INCREMENT COMMENT '用户主键',
    `username`  varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '账号',
    `password`  varchar(512) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
    `nickname`  varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '昵称',
    `gender`    varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '性别',
    `phone`     varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '电话号码',
    `email`     varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
    `age`       int(11)                                                 NULL DEFAULT NULL COMMENT '年龄',
    `status`    tinyint(4)                                              NULL DEFAULT 0 COMMENT '状态',
    `role`      tinyint(4)                                              NULL DEFAULT 0 COMMENT '角色',
    `is_delete` tinyint(4)                                              NULL DEFAULT 0 COMMENT '逻辑删除',
    `create_at` datetime                                                NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` datetime                                                NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `username` (`username`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci COMMENT = '用户表'
  ROW_FORMAT = Dynamic;