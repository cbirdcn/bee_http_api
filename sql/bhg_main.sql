/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50734
 Source Host           : localhost:3306
 Source Schema         : main

 Target Server Type    : MySQL
 Target Server Version : 50734
 File Encoding         : 65001

 Date: 18/11/2021 14:34:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for compensation_items
-- ----------------------------
DROP TABLE IF EXISTS `compensation_items`;
CREATE TABLE `compensation_items` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `compensation_id` int(10) unsigned NOT NULL,
  `items` text NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_comp_id` (`compensation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for compensations
-- ----------------------------
DROP TABLE IF EXISTS `compensations`;
CREATE TABLE `compensations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `server_id` int(10) unsigned NOT NULL,
  `title` varchar(255) NOT NULL,
  `content` varchar(5000) NOT NULL,
  `note` varchar(255) NOT NULL DEFAULT '' COMMENT '面向使用者的备忘录',
  `executed_at` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for server
-- ----------------------------
DROP TABLE IF EXISTS `server`;
CREATE TABLE `server` (
  `id` int(10) unsigned NOT NULL,
  `capacity` int(11) unsigned NOT NULL DEFAULT '10000' COMMENT '单服最大创建角色数量',
  `show_status` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '状态，仅仅用于显示的flag：0-关服，1-开服正常，2-维护',
  `real_status` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '状态，实际状态，增加3-测试，4-提审等',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for server_config
-- ----------------------------
DROP TABLE IF EXISTS `server_config`;
CREATE TABLE `server_config` (
  `id` int(10) unsigned NOT NULL,
  `api_url` varchar(255) NOT NULL DEFAULT '' COMMENT '客户端获取api的地址',
  `master_url` varchar(255) NOT NULL DEFAULT '' COMMENT '客户端获取配置表的地址',
  `master_version` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '使用的配置表版本号',
  `websocket_url` varchar(255) NOT NULL DEFAULT '' COMMENT '长连接地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user_account
-- ----------------------------
DROP TABLE IF EXISTS `user_account`;
CREATE TABLE `user_account` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `device_id` varchar(255) NOT NULL DEFAULT '',
  `use_flg` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '使用状态，robot=0，player=1',
  `password_hash` varchar(255) NOT NULL DEFAULT '',
  `device_name` varchar(255) DEFAULT '',
  `is_banned` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '封号状态：0-正常，1-封号',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for user_server
-- ----------------------------
DROP TABLE IF EXISTS `user_server`;
CREATE TABLE `user_server` (
  `user_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `account_id` int(10) unsigned NOT NULL,
  `server_id` int(10) unsigned NOT NULL,
  `last_login` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_logout` datetime NOT NULL DEFAULT '1000-01-01 00:00:00',
  PRIMARY KEY (`user_id`),
  KEY `account_id` (`account_id`),
  KEY `server_id` (`server_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
