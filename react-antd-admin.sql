/*
 Navicat Premium Data Transfer

 Source Server         : 阿里云
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : 139.224.11.85:3306
 Source Schema         : react-antd-admin

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 24/05/2021 16:28:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_permissions
-- ----------------------------
DROP TABLE IF EXISTS `admin_permissions`;
CREATE TABLE `admin_permissions`  (
  `permission_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
  `code` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '唯一标识',
  `parent_id` int(11) NOT NULL DEFAULT 0,
  `is_help` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否辅助信息',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '修改时间',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`permission_id`) USING BTREE,
  UNIQUE INDEX `ix_admin_permissions_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_permissions
-- ----------------------------
INSERT INTO `admin_permissions` VALUES (1, '用户管理', '', 0, 1, NULL, '2021-05-15 17:26:45');
INSERT INTO `admin_permissions` VALUES (2, '用户列表', 'adminuser:list:view', 1, 0, NULL, '2021-05-15 18:19:56');
INSERT INTO `admin_permissions` VALUES (3, '权限管理', '', 0, 1, NULL, '2021-05-15 18:45:05');
INSERT INTO `admin_permissions` VALUES (4, '权限列表', 'adminpermission:list:view', 3, 0, NULL, '2021-05-15 18:45:19');
INSERT INTO `admin_permissions` VALUES (5, '新增用户', 'adminuser:list:add', 1, 0, NULL, '2021-05-15 18:46:13');
INSERT INTO `admin_permissions` VALUES (6, '编辑用户', 'adminuser:list:edit', 1, 0, NULL, '2021-05-15 18:46:25');
INSERT INTO `admin_permissions` VALUES (7, '新增权限', 'adminpermission:list:add', 3, 0, NULL, '2021-05-15 18:47:25');
INSERT INTO `admin_permissions` VALUES (9, '编辑权限', 'adminpermission:list:edit', 3, 0, '2021-05-23 18:10:45', '2021-05-23 17:59:02');
INSERT INTO `admin_permissions` VALUES (10, '删除权限', 'adminpermission:list:del', 3, 0, '2021-05-23 18:05:09', '2021-05-23 17:59:21');
INSERT INTO `admin_permissions` VALUES (14, '删除用户', 'adminuser:list:del', 1, 0, NULL, '2021-05-23 18:11:16');

-- ----------------------------
-- Table structure for admin_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_users`;
CREATE TABLE `admin_users`  (
  `admin_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
  `password` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码 32位md5',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '状态',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`admin_id`) USING BTREE,
  INDEX `ix_admin_users_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_users
-- ----------------------------
INSERT INTO `admin_users` VALUES (1, 'admin', '82790085228cf8a1e3bac41f45271e5f', 1, '2021-05-24 13:54:53', '2021-05-10 17:17:13');

-- ----------------------------
-- Table structure for admin_users_permissions
-- ----------------------------
DROP TABLE IF EXISTS `admin_users_permissions`;
CREATE TABLE `admin_users_permissions`  (
  `admin_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL,
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0)
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_users_permissions
-- ----------------------------
INSERT INTO `admin_users_permissions` VALUES (1, 3, '2021-05-22 18:33:12');
INSERT INTO `admin_users_permissions` VALUES (1, 7, '2021-05-22 18:33:12');
INSERT INTO `admin_users_permissions` VALUES (1, 4, '2021-05-22 18:33:13');
INSERT INTO `admin_users_permissions` VALUES (1, 1, '2021-05-22 18:33:13');
INSERT INTO `admin_users_permissions` VALUES (1, 6, '2021-05-22 18:33:13');
INSERT INTO `admin_users_permissions` VALUES (1, 5, '2021-05-22 18:33:13');
INSERT INTO `admin_users_permissions` VALUES (1, 2, '2021-05-22 18:33:13');
INSERT INTO `admin_users_permissions` VALUES (1, 14, '2021-05-24 13:54:52');
INSERT INTO `admin_users_permissions` VALUES (1, 9, '2021-05-24 13:54:52');
INSERT INTO `admin_users_permissions` VALUES (1, 10, '2021-05-24 13:54:53');

SET FOREIGN_KEY_CHECKS = 1;
