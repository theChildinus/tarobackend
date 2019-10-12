/*
 Navicat Premium Data Transfer

 Source Server         : mysql_localhost
 Source Server Type    : MySQL
 Source Server Version : 50727
 Source Host           : localhost:3306
 Source Schema         : taro

 Target Server Type    : MySQL
 Target Server Version : 50727
 File Encoding         : 65001

 Date: 12/10/2019 17:10:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for taro_policy
-- ----------------------------
DROP TABLE IF EXISTS `taro_policy`;
CREATE TABLE `taro_policy`  (
  `policy_id` int(1) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT '策略id',
  `policy_sub` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '策略主体',
  `policy_obj` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '策略资源',
  `policy_act` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '策略动作',
  `policy_ctime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`policy_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of taro_policy
-- ----------------------------
INSERT INTO `taro_policy` VALUES (2, 'bob', 'data2', 'write', '2019-10-11 21:24:22');
INSERT INTO `taro_policy` VALUES (3, 'alice', 'data1', 'read', '2019-10-11 21:25:02');
INSERT INTO `taro_policy` VALUES (4, 'data2_admin', 'data2', 'read', '2019-10-11 21:25:54');
INSERT INTO `taro_policy` VALUES (5, 'data2_admin', 'data2', 'write', '2019-10-11 21:26:07');
INSERT INTO `taro_policy` VALUES (6, 'cora', 'data2', 'exec', '2019-10-11 21:26:32');

-- ----------------------------
-- Table structure for taro_resource
-- ----------------------------
DROP TABLE IF EXISTS `taro_resource`;
CREATE TABLE `taro_resource`  (
  `resource_id` int(1) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT '资源id',
  `resource_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '资源名',
  `resource_type` int(10) NULL DEFAULT NULL COMMENT '资源类型',
  `resource_ctime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`resource_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 44 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of taro_resource
-- ----------------------------
INSERT INTO `taro_resource` VALUES (8, '点点图', 2, '2019-10-10 10:41:56');
INSERT INTO `taro_resource` VALUES (9, '绘图工具', 1, '2019-10-10 10:42:50');
INSERT INTO `taro_resource` VALUES (10, '流量态势图', 2, '2019-10-10 10:43:18');
INSERT INTO `taro_resource` VALUES (11, '安全状态态势图', 2, '2019-10-10 11:05:11');
INSERT INTO `taro_resource` VALUES (12, '月度报表', 0, '2019-10-10 11:05:19');
INSERT INTO `taro_resource` VALUES (13, '制表工具', 1, '2019-10-10 11:05:41');
INSERT INTO `taro_resource` VALUES (14, '季度报表', 0, '2019-10-10 11:05:50');
INSERT INTO `taro_resource` VALUES (15, '工具1', 1, '2019-10-10 15:51:13');
INSERT INTO `taro_resource` VALUES (16, '工具2', 1, '2019-10-10 15:51:15');
INSERT INTO `taro_resource` VALUES (18, '工具4', 1, '2019-10-10 15:51:22');
INSERT INTO `taro_resource` VALUES (19, '态势图1', 2, '2019-10-10 15:52:47');
INSERT INTO `taro_resource` VALUES (20, '态势图2', 2, '2019-10-10 15:52:50');
INSERT INTO `taro_resource` VALUES (21, '态势图3', 2, '2019-10-10 15:52:52');
INSERT INTO `taro_resource` VALUES (22, '态势图4', 2, '2019-10-10 15:52:55');
INSERT INTO `taro_resource` VALUES (23, '态势图5', 2, '2019-10-10 15:52:57');
INSERT INTO `taro_resource` VALUES (25, '好用的工具', 1, '2019-10-10 20:26:21');
INSERT INTO `taro_resource` VALUES (26, '好用的工具2', 1, '2019-10-10 20:59:27');
INSERT INTO `taro_resource` VALUES (27, '好用的工具2', 1, '2019-10-10 21:06:01');
INSERT INTO `taro_resource` VALUES (28, '好用的工具2', 1, '2019-10-10 21:07:11');
INSERT INTO `taro_resource` VALUES (29, '好用的工具2', 1, '2019-10-10 21:07:34');
INSERT INTO `taro_resource` VALUES (30, '好用的工具2', 1, '2019-10-10 21:07:48');
INSERT INTO `taro_resource` VALUES (31, '好用的工具2', 1, '2019-10-10 21:12:38');
INSERT INTO `taro_resource` VALUES (32, '好用的工具2', 1, '2019-10-10 21:28:29');
INSERT INTO `taro_resource` VALUES (33, '好用的工具5', 1, '2019-10-10 21:34:11');
INSERT INTO `taro_resource` VALUES (34, '好用的工具5', 1, '2019-10-10 21:35:09');
INSERT INTO `taro_resource` VALUES (35, '好用的工具5', 1, '2019-10-10 21:39:45');
INSERT INTO `taro_resource` VALUES (36, '好用的工具5', 1, '2019-10-10 21:46:46');
INSERT INTO `taro_resource` VALUES (37, '好用的工具3', 1, '2019-10-10 21:48:17');
INSERT INTO `taro_resource` VALUES (38, '好用的工具3', 1, '2019-10-10 21:55:38');
INSERT INTO `taro_resource` VALUES (39, '好用的工具4', 1, '2019-10-10 21:57:02');
INSERT INTO `taro_resource` VALUES (41, '好看的态势图', 2, '2019-10-10 22:08:45');
INSERT INTO `taro_resource` VALUES (42, '好看的态势图', 2, '2019-10-10 22:12:28');

-- ----------------------------
-- Table structure for taro_user
-- ----------------------------
DROP TABLE IF EXISTS `taro_user`;
CREATE TABLE `taro_user`  (
  `user_id` int(1) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `user_role` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户角色',
  `user_department` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户部门',
  `user_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户地址',
  `user_email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'email',
  `user_phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系方式',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
