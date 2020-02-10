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

 Date: 10/02/2020 17:47:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for taro_enum
-- ----------------------------
DROP TABLE IF EXISTS `taro_enum`;
CREATE TABLE `taro_enum`  (
  `enum_id` int(1) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT,
  `enum_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `enum_value` varchar(8192) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`enum_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of taro_enum
-- ----------------------------
INSERT INTO `taro_enum` VALUES (1, 'user_role', '管理员##经理##员工##第三方团队##发布者##订阅者');
INSERT INTO `taro_enum` VALUES (2, 'resource_type', '主题##报表##工具##态势图##Peer节点##Order节点');
INSERT INTO `taro_enum` VALUES (3, 'policy_act', 'read##write##exec##upload##sub##pub');
INSERT INTO `taro_enum` VALUES (4, 'user_organization', '[{\"orgId\":35},{\"id\":24,\"label\":\"顶级部门1\",\"value\":\"顶级部门1\",\"isEdit\":false,\"children\":[{\"id\":27,\"label\":\"子部门1-1\",\"value\":\"子部门1-1\",\"isEdit\":false,\"children\":[{\"id\":28,\"label\":\"子部门1-1-1\",\"value\":\"子部门1-1-1\",\"isEdit\":false,\"children\":[{\"id\":33,\"label\":\"子部门1-1-1-1\",\"value\":\"子部门1-1-1-1\",\"isEdit\":false}]}]}]},{\"id\":25,\"label\":\"顶级部门2\",\"value\":\"顶级部门2\",\"isEdit\":false,\"children\":[{\"id\":29,\"label\":\"子部门2-1\",\"value\":\"子部门2-1\",\"isEdit\":false}]},{\"id\":26,\"label\":\"顶级部门3\",\"value\":\"顶级部门3\",\"isEdit\":false,\"children\":[{\"id\":30,\"label\":\"子部门3-1\",\"value\":\"子部门3-1\",\"isEdit\":false,\"children\":[{\"id\":32,\"label\":\"子部门3-1-1\",\"value\":\"子部门3-1-1\",\"isEdit\":false}]}]},{\"id\":27,\"label\":\"顶级部门4\",\"value\":\"顶级部门4\",\"isEdit\":false,\"children\":[{\"id\":31,\"label\":\"子部门4-1\",\"value\":\"子部门4-1\",\"isEdit\":false}]},{\"id\":35,\"label\":\"顶级部门5\",\"value\":\"顶级部门5\",\"isEdit\":false,\"children\":[]}]');
INSERT INTO `taro_enum` VALUES (5, 'policy_tree', '[{\"policyTreeId\":9},{\"id\":1,\"label\":\"策略1\",\"value\":\"策略1\",\"isEdit\":false,\"children\":[]},{\"id\":3,\"label\":\"策略2\",\"value\":\"策略2\",\"isEdit\":false,\"children\":[]},{\"id\":4,\"label\":\"策略3\",\"value\":\"策略3\",\"isEdit\":false,\"children\":[{\"id\":4,\"label\":\"子策略3-1\",\"value\":\"子策略3-1\",\"isEdit\":false}]},{\"id\":8,\"label\":\"父策略4\",\"value\":\"父策略4\",\"isEdit\":false,\"children\":[]},{\"id\":9,\"label\":\"父策略5\",\"value\":\"父策略5\",\"isEdit\":false,\"children\":[]}]');
INSERT INTO `taro_enum` VALUES (6, 'policy_model', '[{\"policy_name\":\"策略1\",\"model_type\":\"ACL\"},{\"policy_name\":\"策略2\",\"model_type\":\"RBAC\"},{\"policy_name\":\"策略3\",\"model_type\":\"RBAC\"},{\"policy_name\":\"父策略4\",\"model_type\":\"ACL\"},{\"policy_name\":\"父策略5\",\"model_type\":\"ACL\"}]');
INSERT INTO `taro_enum` VALUES (7, 'policy_resource', '[{\"policyResourceId\":null},{\"id\":1,\"label\":\"资源1\",\"value\":\"资源1\",\"isEdit\":false,\"children\":[]},{\"id\":3,\"label\":\"资源2\",\"value\":\"资源2\",\"isEdit\":false,\"children\":[]},{\"id\":4,\"label\":\"资源3\",\"value\":\"资源3\",\"isEdit\":false,\"children\":[{\"id\":4,\"label\":\"子资源3-1\",\"value\":\"子资源3-1\",\"isEdit\":false}]},{\"id\":6,\"label\":\"资源4\",\"value\":\"资源4\",\"isEdit\":false,\"children\":[]},{\"id\":null,\"label\":\"资源5\",\"value\":\"资源5\",\"isEdit\":false,\"children\":[]}]');
INSERT INTO `taro_enum` VALUES (8, 'identity_organization', '[{\"orgId\":37},{\"id\":35,\"label\":\"Org1\",\"value\":\"Org1\",\"isEdit\":false,\"children\":[]},{\"id\":36,\"label\":\"Org2\",\"value\":\"Org2\",\"isEdit\":false,\"children\":[]},{\"id\":37,\"label\":\"Org3\",\"value\":\"Org3\",\"isEdit\":false,\"children\":[]}]');

-- ----------------------------
-- Table structure for taro_identity
-- ----------------------------
DROP TABLE IF EXISTS `taro_identity`;
CREATE TABLE `taro_identity`  (
  `identity_id` int(1) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT 'Fabric 注册Id',
  `identity_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Fabric 参与者名',
  `identity_secret` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Fabric 参与者密码',
  `identity_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Fabric 参与者类型',
  `identity_affiliation` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Fabric 参与者从属关系',
  `identity_attrs` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Fabric 参与者属性',
  `identity_ctime` datetime(0) NULL DEFAULT NULL COMMENT 'Fabric 参与者创建时间',
  `identity_status` int(1) NULL DEFAULT NULL COMMENT 'Fabric 参与者状态',
  `identity_ip` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Fabric 参与者主机IP',
  `identity_user` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Fabric 参与者主机名',
  `identity_pw` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Fabric 参与者主机密码',
  `identity_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Fabric 参与者主机路径',
  PRIMARY KEY (`identity_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of taro_identity
-- ----------------------------
INSERT INTO `taro_identity` VALUES (1, 'peer1', 'peer1pw', 'peer', 'org1', 'email=peer1@gmail.com', '2019-11-21 16:58:40', 2, '211.159.147.194', 'wayne', 'wayne941001', '/home/wayne/kong');
INSERT INTO `taro_identity` VALUES (3, 'order1', 'order1pw', 'order', 'org1', '', '2019-11-21 16:59:19', 1, '211.159.147.194', 'wayne', 'wayne941001', NULL);
INSERT INTO `taro_identity` VALUES (4, 'peer2', 'peer2pw', 'peer', 'org1', '', '2019-11-21 16:59:45', 0, '211.159.147.194', 'wayne', 'wayne941001', NULL);
INSERT INTO `taro_identity` VALUES (5, 'order2', 'order2pw', 'order', 'org2', '', '2019-11-22 08:30:53', 0, '211.159.147.194', 'wayne', 'wayne941001', NULL);
INSERT INTO `taro_identity` VALUES (6, 'user1', 'user1pw', 'client', 'org2', 'app1Admin=true:ecert,email=user1@gmail.com', '2019-11-21 15:00:43', 0, '211.159.147.194', 'wayne', 'wayne941001', NULL);

-- ----------------------------
-- Table structure for taro_policy
-- ----------------------------
DROP TABLE IF EXISTS `taro_policy`;
CREATE TABLE `taro_policy`  (
  `policy_id` int(1) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT '策略id',
  `policy_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '策略名称',
  `policy_sub` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '策略规则主体',
  `policy_obj` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '策略规则资源',
  `policy_act` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '策略规则动作',
  `policy_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '策略类型',
  `policy_ctime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`policy_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 42 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of taro_policy
-- ----------------------------
INSERT INTO `taro_policy` VALUES (2, '策略1', 'kong', 'data2', 'write', 'IOT策略', '2019-10-11 21:24:22');
INSERT INTO `taro_policy` VALUES (3, '策略1', 'kong', 'data2', 'exec', 'IOT策略', '2019-10-11 21:25:02');
INSERT INTO `taro_policy` VALUES (4, '策略1', 'zhao', 'data1', 'read', 'IOT策略', '2019-10-11 21:25:54');
INSERT INTO `taro_policy` VALUES (6, '策略1', '管理员', 'data3', 'upload', 'IOT策略', '2019-10-11 21:26:32');
INSERT INTO `taro_policy` VALUES (7, '策略2', 'Org1MSP.admin', 'Org1MSP', 'read', 'IOT策略', '2020-01-01 15:01:02');
INSERT INTO `taro_policy` VALUES (8, '策略2', 'Org1MSP.peer', 'Org1MSP', 'read', 'IOT策略', '2020-01-01 15:01:45');
INSERT INTO `taro_policy` VALUES (15, '策略1', '管理员', 'data4', 'exec', 'IOT策略', '2019-10-14 10:43:45');
INSERT INTO `taro_policy` VALUES (20, '策略1', 'kong', 'start', 'sub', 'IOT策略', '2020-01-01 09:19:06');
INSERT INTO `taro_policy` VALUES (21, '策略1', 'yang', 'start', 'pub', 'IOT策略', '2020-01-01 09:19:11');
INSERT INTO `taro_policy` VALUES (23, '策略3', 'zhao', '444444', 'read', 'Fabric策略', '2020-02-06 17:10:03');
INSERT INTO `taro_policy` VALUES (24, '策略3', 'Org3/order1', '资源1', 'read', 'Fabric策略', '2020-02-06 17:10:24');
INSERT INTO `taro_policy` VALUES (27, '策略3', 'Org1/peer1', '资源5', 'read', 'Fabric策略', '2020-02-08 16:08:37');
INSERT INTO `taro_policy` VALUES (28, '策略3', 'zhao', '12121212', 'read', 'Fabric策略', '2020-02-08 16:09:48');
INSERT INTO `taro_policy` VALUES (31, '策略3#子策略3-1', 'Org3/peer1', '资源3/子资源3-1', 'read', 'Fabric策略', '2020-02-08 17:47:47');
INSERT INTO `taro_policy` VALUES (33, '策略3', '顶级部门1/yang', '3333333', 'read', 'Fabric策略', '2020-02-08 18:02:11');
INSERT INTO `taro_policy` VALUES (34, '策略3#子策略3-1', 'zzzz', '99999', 'read', 'Fabric策略', '2020-02-08 18:12:58');
INSERT INTO `taro_policy` VALUES (38, '策略3', 'user1', '资源3/子资源3-1', 'read', 'Fabric策略', '2020-02-10 16:12:13');
INSERT INTO `taro_policy` VALUES (39, '策略3', 'Org1/peer1', '资源5', 'read', 'Fabric策略', '2020-02-10 16:34:45');
INSERT INTO `taro_policy` VALUES (40, '策略2', '顶级部门5/yang', '资源5', 'read', 'IOT策略', '2020-02-10 16:42:42');
INSERT INTO `taro_policy` VALUES (41, '父策略4', '顶级部门5/zzzz', '资源5', 'pub', 'IOT策略', '2020-02-10 17:33:29');

-- ----------------------------
-- Table structure for taro_resource
-- ----------------------------
DROP TABLE IF EXISTS `taro_resource`;
CREATE TABLE `taro_resource`  (
  `resource_id` int(1) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT COMMENT '资源id',
  `resource_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '资源名',
  `resource_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '资源类型',
  `resource_ctime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`resource_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 45 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of taro_resource
-- ----------------------------
INSERT INTO `taro_resource` VALUES (9, 'peer1节点', 'Peer节点', '2019-10-10 10:42:50');
INSERT INTO `taro_resource` VALUES (10, '流量态势图', '态势图', '2019-10-10 10:43:18');
INSERT INTO `taro_resource` VALUES (11, '安全状态态势图', '态势图', '2019-10-10 11:05:11');
INSERT INTO `taro_resource` VALUES (12, '月度报表', '报表', '2019-10-10 11:05:19');
INSERT INTO `taro_resource` VALUES (13, '业务流程绘制工具', '工具', '2019-10-10 11:05:41');
INSERT INTO `taro_resource` VALUES (14, '季度报表', '报表', '2019-10-10 11:05:50');
INSERT INTO `taro_resource` VALUES (15, '工具1', '工具', '2019-10-10 15:51:13');
INSERT INTO `taro_resource` VALUES (16, '时间单元11111', '时间单元', '2019-10-10 15:51:15');
INSERT INTO `taro_resource` VALUES (18, '工具4', '工具', '2019-10-10 15:51:22');
INSERT INTO `taro_resource` VALUES (19, '态势图1', '态势图', '2019-10-10 15:52:47');
INSERT INTO `taro_resource` VALUES (20, '态势图2', '态势图', '2019-10-10 15:52:50');
INSERT INTO `taro_resource` VALUES (21, '态势图3', '态势图', '2019-10-10 15:52:52');
INSERT INTO `taro_resource` VALUES (22, '态势图4', '态势图', '2019-10-10 15:52:55');
INSERT INTO `taro_resource` VALUES (23, '态势图5', '态势图', '2019-10-10 15:52:57');
INSERT INTO `taro_resource` VALUES (25, '好用的工具', '工具', '2019-10-10 20:26:21');
INSERT INTO `taro_resource` VALUES (26, '好用的工具2', '工具', '2019-10-10 20:59:27');
INSERT INTO `taro_resource` VALUES (27, '好用的工具2', '工具', '2019-10-10 21:06:01');
INSERT INTO `taro_resource` VALUES (28, '好用的工具2', '工具', '2019-10-10 21:07:11');
INSERT INTO `taro_resource` VALUES (29, '好用的工具2', '工具', '2019-10-10 21:07:34');
INSERT INTO `taro_resource` VALUES (30, '好用的工具2', '工具', '2019-10-10 21:07:48');
INSERT INTO `taro_resource` VALUES (31, '好用的工具2', '工具', '2019-10-10 21:12:38');
INSERT INTO `taro_resource` VALUES (32, '好用的工具2', '工具', '2019-10-10 21:28:29');
INSERT INTO `taro_resource` VALUES (33, '好用的工具5', '工具', '2019-10-10 21:34:11');
INSERT INTO `taro_resource` VALUES (34, '好用的工具5', '工具', '2019-10-10 21:35:09');
INSERT INTO `taro_resource` VALUES (35, '好用的工具5', '工具', '2019-10-10 21:39:45');
INSERT INTO `taro_resource` VALUES (36, '好用的工具5', '工具', '2019-10-10 21:46:46');
INSERT INTO `taro_resource` VALUES (37, '好用的工具3', '工具', '2019-10-10 21:48:17');
INSERT INTO `taro_resource` VALUES (38, '好用的工具3', '工具', '2019-10-10 21:55:38');
INSERT INTO `taro_resource` VALUES (39, '好用的工具4', '工具', '2019-10-10 21:57:02');
INSERT INTO `taro_resource` VALUES (41, '好看的态势图', '态势图', '2019-10-10 22:08:45');
INSERT INTO `taro_resource` VALUES (42, '好看的态势图', '态势图', '2019-10-10 22:12:28');
INSERT INTO `taro_resource` VALUES (43, '贼好用的工具', '工具', '2019-10-17 21:03:40');
INSERT INTO `taro_resource` VALUES (44, '微波单元1111', '微波单元', '2019-10-18 15:06:43');

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
  `user_email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '电子邮箱',
  `user_phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系方式',
  `user_status` int(1) NULL DEFAULT NULL COMMENT '用户状态',
  `user_hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户哈希',
  `user_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户证书存储路径',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of taro_user
-- ----------------------------
INSERT INTO `taro_user` VALUES (1, 'zhao', '管理员', '顶级部门1/子部门1-1/子部门1-1-1/子部门1-1-1-1', '111111', '6555@163.com', '123123123', 1, '3d5557a4b09c67bbc0019cc809d26da6', 'D:\\goProjects\\tarobackend\\card\\zhao');
INSERT INTO `taro_user` VALUES (2, 'kong', '经理,第三方团队', '顶级部门3/子部门3-1', '123131231', 'df@qq.com', '12312312', 1, '575c776ac08615a4949a6e35d4b73896', 'D:\\goProjects\\tarobackend\\card\\kong');
INSERT INTO `taro_user` VALUES (3, 'yang', '经理', '顶级部门2/子部门2-1', '1231', 'dfdf@qq.com', '12312312', 1, '', 'D:\\goProjects\\tarobackend\\card\\yang');
INSERT INTO `taro_user` VALUES (4, 'zzzz', '员工', '顶级部门4/子部门4-1', '1231adfasdf', 'dfdf@qq.com', '12312312', 0, '', 'D:\\goProjects\\tarobackend\\card\\');
INSERT INTO `taro_user` VALUES (5, 'adsfasdf', '第三方团队', '顶级部门3/子部门3-1/子部门3-1-1', '1231adfasdf', 'dfdf@qq.com', '12312312', 0, '', 'D:\\goProjects\\tarobackend\\card\\yang');
INSERT INTO `taro_user` VALUES (6, 'ggggggg', '经理', '顶级部门2/子部门2-1', '1231adfasdf', 'dfdf@qq.com', '12312312', 0, '', 'D:\\goProjects\\tarobackend\\card\\yang');
INSERT INTO `taro_user` VALUES (14, 'abddd', '员工', '顶级部门1/子部门1-1/子部门1-1-1', 'adfad', '11233434322@qq.com', '123124ddd', 0, '', 'D:\\goProjects\\tarobackend\\card\\yang');
INSERT INTO `taro_user` VALUES (15, 'abdddaaaa', '员工', '顶级部门2/子部门2-1', 'adfad', '11233434322@qq.com', '123124ddd', 0, '', 'D:\\goProjects\\tarobackend\\card\\yang');
INSERT INTO `taro_user` VALUES (16, 'adfasdfd', '员工', '顶级部门4/子部门4-1', 'dfdfdf', '1123131@123.com', '544342341', 0, '', 'D:\\goProjects\\tarobackend\\card\\yang');
INSERT INTO `taro_user` VALUES (17, 'test9', '员工', '顶级部门4/子部门4-1', 'dfadf', '123124323@qq.com', '1231243432', 0, '75cded9e84623b344159aaab0aa8e7fd', 'D:\\goProjects\\tarobackend\\card\\yang');
INSERT INTO `taro_user` VALUES (18, 'test1', '经理', '顶级部门1/子部门1-1/子部门1-1-1/子部门1-1-1-1', 'bupt', '123456@qq.com', '123456', 0, '', 'D:\\goProjects\\tarobackend\\card\\yang');
INSERT INTO `taro_user` VALUES (19, 'eeeeeee', '管理员', '顶级部门1/子部门1-1', 'erererere', '55566@qq.com', '1231230434231', 0, '', 'D:\\goProjects\\tarobackend\\card\\yang');

SET FOREIGN_KEY_CHECKS = 1;
