-- User
CREATE USER 'postar' @'localhost' IDENTIFIED BY '123456';

GRANT ALL PRIVILEGES ON `postar`.* TO 'postar' @'localhost';

FLUSH PRIVILEGES;

-- Postar database
CREATE DATABASE IF NOT EXISTS `postar` DEFAULT CHARACTER SET utf8mb4;

-- Space table
CREATE TABLE IF NOT EXISTS `postar`.`spaces` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name` VARCHAR(64) NOT NULL COMMENT '名称',
    `token` VARCHAR(512) NOT NULL DEFAULT '' COMMENT '令牌',
    `state` TINYINT NOT NULL DEFAULT 0 COMMENT '状态',
    `create_time` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` BIGINT NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`) COMMENT '主键',
    UNIQUE `name_index` (`name`) COMMENT '名称的查询索引'
) AUTO_INCREMENT = 100 ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '业务空间表';

-- Account table
CREATE TABLE IF NOT EXISTS `postar`.`accounts` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT '主键',
    `space_id` INT NOT NULL COMMENT '业务空间主键',
    `host` VARCHAR(128) NOT NULL COMMENT '主机地址',
    `port` SMALLINT NOT NULL DEFAULT 0 COMMENT '端口',
    `username` VARCHAR(128) NOT NULL COMMENT '用户名',
    `password` VARCHAR(128) NOT NULL COMMENT '密码',
    `smtp_auth` TINYINT NOT NULL DEFAULT 0 COMMENT 'SMTP 认证方式',
    `state` TINYINT NOT NULL DEFAULT 0 COMMENT '状态',
    `create_time` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` BIGINT NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`) COMMENT '主键',
    UNIQUE `space_id_username_index` (`space_id`, `username`) COMMENT '业务空间和用户名的查询索引',
    INDEX `space_id_host_index` (`space_id`, `host`) COMMENT '业务空间和主机的查询索引',
    INDEX `space_id_state_index` (`space_id`, `state`) COMMENT '业务空间和状态的查询索引'
) AUTO_INCREMENT = 10000 ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '账号表';

-- Template table
CREATE TABLE IF NOT EXISTS `postar`.`templates` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键',
    `space_id` INT NOT NULL COMMENT '业务空间主键',
    `account_id` INT NOT NULL COMMENT '账号主键',
    `name` VARCHAR(128) NOT NULL COMMENT '模板名称',
    `description` VARCHAR(512) NOT NULL COMMENT '描述',
    `email_subject` VARCHAR(256) NOT NULL COMMENT '邮件主题',
    `email_to` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '邮件收件人',
    `email_cc` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '邮件抄送人',
    `email_bcc` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '邮件秘密抄送人',
    `email_content_type` INT NOT NULL DEFAULT 0 COMMENT '账号主键',
    `email_content` VARCHAR(4096) NOT NULL COMMENT '邮件内容',
    `state` TINYINT NOT NULL DEFAULT 0 COMMENT '状态',
    `create_time` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` BIGINT NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`) COMMENT '主键',
    INDEX `space_id_account_id_index` (`space_id`, `account_id`) COMMENT '业务空间和账号的查询索引',
    UNIQUE `space_id_name_index` (`space_id`, `name`) COMMENT '业务空间和模板名称的查询索引',
    INDEX `space_id_state_index` (`space_id`, `state`) COMMENT '业务空间和状态的查询索引'
) AUTO_INCREMENT = 1000000 ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '邮件模板表';