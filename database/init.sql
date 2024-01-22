-- 用户表
CREATE TABLE `user_basic` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `uid` INT UNSIGNED NOT NULL COMMENT '用户唯一标识 ID',
    `account` VARCHAR(50) NOT NULL COMMENT '用户名',
    `password` VARCHAR(255) NOT NULL COMMENT '密码',
    `nickname` VARCHAR(50) NOT NULL COMMENT '昵称',
    `email` VARCHAR(100) NULL COMMENT '邮箱',
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `deleted_at` BIGINT NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
);

-- 房间信息表
CREATE TABLE `room_basic` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `uid` VARCHAR(50) NOT NULL COMMENT '房间唯一标识 ID',
    `user_uid` INT UNSIGNED NOT NULL COMMENT '创建者ID',
    `name` VARCHAR(100) NOT NULL COMMENT '房间名称',
    `info` VARCHAR(255) NULL COMMENT '房间简介',
    `salt` VARCHAR(50) NULL COMMENT '聊天室密码盐',
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `deleted_at` BIGINT NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
);

-- 消息表
CREATE TABLE `message_basic` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_uid` INT UNSIGNED NOT NULL COMMENT '用户ID',
  `room_uid` VARCHAR(50) NOT NULL COMMENT '房间ID',
  `content` LONGTEXT NOT NULL COMMENT '聊天内容',
  `created_at` BIGINT NOT NULL,
  `updated_at` BIGINT NOT NULL,
  `deleted_at` BIGINT NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ROW_FORMAT = DYNAMIC;


-- 用户房间关联表
CREATE TABLE `user_room` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_uid` INT UNSIGNED NOT NULL COMMENT '用户ID',
    `room_uid` VARCHAR(50) NOT NULL COMMENT '房间ID',
    `room_type` TINYINT NOT NULL COMMENT '房间类型 1:群聊房间 2:私聊房间',
    `joined_at` BIGINT NOT NULL,
    `created_at` BIGINT NOT NULL,
    `updated_at` BIGINT NOT NULL,
    `deleted_at` BIGINT NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
);
