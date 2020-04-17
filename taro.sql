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

 Date: 17/04/2020 18:20:38
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
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of taro_enum
-- ----------------------------
INSERT INTO `taro_enum` VALUES (1, 'user_role', '管理员##经理##员工##发布者##订阅者##值班调度人员##设备运维人员');
INSERT INTO `taro_enum` VALUES (2, 'resource_type', '主题##报表##工具##态势图##Peer节点##Order节点');
INSERT INTO `taro_enum` VALUES (3, 'policy_act', 'read##write##exec##upload##sub##pub##查看##修改##检查');
INSERT INTO `taro_enum` VALUES (4, 'user_organization', '[{\"orgId\":49},{\"id\":44,\"label\":\"领导小组\",\"value\":\"领导小组\",\"isEdit\":false,\"children\":[]},{\"id\":45,\"label\":\"指挥大厅\",\"value\":\"指挥大厅\",\"isEdit\":false,\"children\":[]},{\"id\":46,\"label\":\"值班调度部\",\"value\":\"值班调度部\",\"isEdit\":false,\"children\":[]},{\"id\":47,\"label\":\"设备运维部\",\"value\":\"设备运维部\",\"isEdit\":false,\"children\":[]},{\"id\":48,\"label\":\"光纤运维部\",\"value\":\"光纤运维部\",\"isEdit\":false,\"children\":[]},{\"id\":49,\"label\":\"后台运维部\",\"value\":\"后台运维部\",\"isEdit\":false,\"children\":[]}]');
INSERT INTO `taro_enum` VALUES (5, 'policy_tree', '[{\"policyTreeId\":12},{\"id\":1,\"label\":\"策略1\",\"value\":\"策略1\",\"isEdit\":false,\"children\":[]},{\"id\":3,\"label\":\"策略2\",\"value\":\"策略2\",\"isEdit\":false,\"children\":[]},{\"id\":10,\"label\":\"区块链访问控制策略\",\"value\":\"区块链访问控制策略\",\"isEdit\":false,\"children\":[]},{\"id\":11,\"label\":\"跨物联网服务和区块链的业务流程策略\",\"value\":\"跨物联网服务和区块链的业务流程策略\",\"isEdit\":false,\"children\":[]},{\"id\":12,\"label\":\"运控分系统访问控制策略集合\",\"value\":\"运控分系统访问控制策略集合\",\"isEdit\":false,\"children\":[]}]');
INSERT INTO `taro_enum` VALUES (6, 'policy_model', '[{\"policy_name\":\"策略1\",\"model_type\":\"ACL\"},{\"policy_name\":\"策略2\",\"model_type\":\"ABAC\"},{\"policy_name\":\"区块链访问控制策略\",\"model_type\":\"ACL\"},{\"policy_name\":\"跨物联网服务和区块链的业务流程策略\",\"model_type\":\"RBAC\"},{\"policy_name\":\"运控分系统访问控制策略集合\",\"model_type\":\"ABAC\"}]');
INSERT INTO `taro_enum` VALUES (7, 'policy_resource', '[{\"policyResourceId\":58},{\"id\":1,\"label\":\"区块链资源\",\"value\":\"区块链资源\",\"isEdit\":false,\"children\":[{\"id\":17,\"label\":\"lscc\",\"value\":\"lscc\",\"isEdit\":false,\"children\":[{\"id\":20,\"label\":\"ChaincodeExists\",\"value\":\"ChaincodeExists\",\"isEdit\":false},{\"id\":21,\"label\":\"GetDeploymentSpec\",\"value\":\"GetDeploymentSpec\",\"isEdit\":false},{\"id\":22,\"label\":\"GetChaincodeData\",\"value\":\"GetChaincodeData\",\"isEdit\":false},{\"id\":23,\"label\":\"GetInstantiatedChaincodes\",\"value\":\"GetInstantiatedChaincodes\",\"isEdit\":false}]},{\"id\":18,\"label\":\"qscc\",\"value\":\"qscc\",\"isEdit\":false,\"children\":[{\"id\":24,\"label\":\"GetChainInfo\",\"value\":\"GetChainInfo\",\"isEdit\":false},{\"id\":25,\"label\":\"GetBlockByNumber\",\"value\":\"GetBlockByNumber\",\"isEdit\":false},{\"id\":26,\"label\":\"GetBlockByHash\",\"value\":\"GetBlockByHash\",\"isEdit\":false},{\"id\":27,\"label\":\"GetTransactionByID\",\"value\":\"GetTransactionByID\",\"isEdit\":false},{\"id\":28,\"label\":\"GetBlockByTxID\",\"value\":\"GetBlockByTxID\",\"isEdit\":false}]},{\"id\":19,\"label\":\"cscc\",\"value\":\"cscc\",\"isEdit\":false,\"children\":[{\"id\":29,\"label\":\"GetConfigBlock\",\"value\":\"GetConfigBlock\",\"isEdit\":false},{\"id\":30,\"label\":\"GetConfigTree\",\"value\":\"GetConfigTree\",\"isEdit\":false},{\"id\":31,\"label\":\"SimulateConfigTreeUpdate\",\"value\":\"SimulateConfigTreeUpdate\",\"isEdit\":false}]},{\"id\":32,\"label\":\"peer\",\"value\":\"peer\",\"isEdit\":false,\"children\":[{\"id\":33,\"label\":\"Propose\",\"value\":\"Propose\",\"isEdit\":false},{\"id\":34,\"label\":\"ChaincodeToChaincode\",\"value\":\"ChaincodeToChaincode\",\"isEdit\":false}]},{\"id\":35,\"label\":\"event\",\"value\":\"event\",\"isEdit\":false,\"children\":[{\"id\":36,\"label\":\"Block\",\"value\":\"Block\",\"isEdit\":false},{\"id\":37,\"label\":\"FilteredBlock\",\"value\":\"FilteredBlock\",\"isEdit\":false}]}]},{\"id\":6,\"label\":\"物联网资源\",\"value\":\"物联网资源\",\"isEdit\":false,\"children\":[{\"id\":11,\"label\":\"表\",\"value\":\"表\",\"isEdit\":false,\"children\":[{\"id\":38,\"label\":\"稳定度指标\",\"value\":\"稳定度指标\",\"isEdit\":false},{\"id\":39,\"label\":\"设备运行状态\",\"value\":\"设备运行状态\",\"isEdit\":false},{\"id\":40,\"label\":\"参数状态表\",\"value\":\"参数状态表\",\"isEdit\":false},{\"id\":41,\"label\":\"设备属性\",\"value\":\"设备属性\",\"isEdit\":false},{\"id\":42,\"label\":\"光纤属性\",\"value\":\"光纤属性\",\"isEdit\":false},{\"id\":43,\"label\":\"故障报表\",\"value\":\"故障报表\",\"isEdit\":false},{\"id\":44,\"label\":\"日报\",\"value\":\"日报\",\"isEdit\":false},{\"id\":45,\"label\":\"月报\",\"value\":\"月报\",\"isEdit\":false},{\"id\":46,\"label\":\"年报\",\"value\":\"年报\",\"isEdit\":false}]},{\"id\":12,\"label\":\"工具\",\"value\":\"工具\",\"isEdit\":false},{\"id\":13,\"label\":\"图\",\"value\":\"图\",\"isEdit\":false,\"children\":[{\"id\":47,\"label\":\"中国地图\",\"value\":\"中国地图\",\"isEdit\":false},{\"id\":48,\"label\":\"阿伦方差曲线\",\"value\":\"阿伦方差曲线\",\"isEdit\":false},{\"id\":49,\"label\":\"相位噪声曲线\",\"value\":\"相位噪声曲线\",\"isEdit\":false},{\"id\":50,\"label\":\"线路地图\",\"value\":\"线路地图\",\"isEdit\":false},{\"id\":51,\"label\":\"设备关系图\",\"value\":\"设备关系图\",\"isEdit\":false},{\"id\":52,\"label\":\"故障树\",\"value\":\"故障树\",\"isEdit\":false},{\"id\":53,\"label\":\"参数信息态势\",\"value\":\"参数信息态势\",\"isEdit\":false},{\"id\":54,\"label\":\"光纤布局图\",\"value\":\"光纤布局图\",\"isEdit\":false},{\"id\":55,\"label\":\"OTDR图\",\"value\":\"OTDR图\",\"isEdit\":false}]},{\"id\":56,\"label\":\"业务流程\",\"value\":\"业务流程\",\"isEdit\":false,\"children\":[{\"id\":57,\"label\":\"设备故障0715\",\"value\":\"设备故障0715\",\"isEdit\":false}]}]}]');
INSERT INTO `taro_enum` VALUES (8, 'identity_organization', '[{\"orgId\":41},{\"id\":35,\"label\":\"org1\",\"value\":\"org1\",\"isEdit\":false,\"children\":[]},{\"id\":36,\"label\":\"org2\",\"value\":\"org2\",\"isEdit\":false,\"children\":[]},{\"id\":37,\"label\":\"org3\",\"value\":\"org3\",\"isEdit\":false,\"children\":[]},{\"id\":38,\"label\":\"Channel\",\"value\":\"Channel\",\"isEdit\":false,\"children\":[{\"id\":38,\"label\":\"Application\",\"value\":\"Application\",\"isEdit\":false}]}]');
INSERT INTO `taro_enum` VALUES (9, 'mutex_role', '[{\"user_role1\":\"值班调度部/员工\",\"user_role2\":\"设备运维部/员工\"},{\"user_role1\":\"设备运维部/员工\",\"user_role2\":\"光纤运维部/员工\"},{\"user_role1\":\"后台运维部/员工\",\"user_role2\":\"值班调度部/员工\"}]');

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
  `identity_hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Fabric 参与者Hash',
  PRIMARY KEY (`identity_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of taro_identity
-- ----------------------------
INSERT INTO `taro_identity` VALUES (1, 'peer1', 'peer1pw', 'peer', 'org1', 'email=peer1@gmail.com', '2019-11-21 16:58:40', 2, '175.24.88.50', 'ubuntu', 'zky0825ASP', '/home/ubuntu/store', NULL);
INSERT INTO `taro_identity` VALUES (3, 'order1', 'order1pw', 'order', 'org1', 'email=peer1@gmail.com', '2019-11-21 16:59:19', 0, '211.159.147.194', 'wayne', 'wayne941001', '/home/wayne/kong', NULL);
INSERT INTO `taro_identity` VALUES (4, 'peer2', 'peer2pw', 'peer', 'org2', 'app1Admin=true:ecert,email=user1@gmail.com', '2019-11-21 16:59:45', 0, '211.159.147.194', 'wayne', 'wayne941001', '/home/wayne/kong', NULL);
INSERT INTO `taro_identity` VALUES (5, 'order2', 'order2pw', 'order', 'org2', '', '2019-11-22 08:30:53', 1, '211.159.147.194', 'wayne', 'wayne941001', '/home/wayne/kong', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 193 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of taro_policy
-- ----------------------------
INSERT INTO `taro_policy` VALUES (4, '策略1', 'zhao', '物联网资源/态势图', 'read', '物联网用户', '2019-10-11 21:25:54');
INSERT INTO `taro_policy` VALUES (6, '策略1', '管理员', '物联网资源/工具', 'upload', '物联网用户', '2019-10-11 21:26:32');
INSERT INTO `taro_policy` VALUES (15, '策略1', '管理员', '物联网资源/态势图', 'exec', '物联网用户', '2019-10-14 10:43:45');
INSERT INTO `taro_policy` VALUES (20, '策略1', '总裁办/kong', '物联网资源/态势图', 'sub', '物联网用户', '2020-01-01 09:19:06');
INSERT INTO `taro_policy` VALUES (21, '策略1', '总裁办/yang', '资源1', 'exec', '物联网用户', '2020-01-01 09:19:11');
INSERT INTO `taro_policy` VALUES (51, '区块链访问控制策略', 'Channel/Application/Readers', '区块链资源/lscc/ChaincodeExists', 'permit', '区块链参与者', '2020-02-13 18:05:49');
INSERT INTO `taro_policy` VALUES (52, '区块链访问控制策略', 'Channel/Application/Readers', '区块链资源/lscc/GetDeploymentSpec', 'permit', '区块链参与者', '2020-02-13 18:07:49');
INSERT INTO `taro_policy` VALUES (53, '区块链访问控制策略', 'Channel/Application/Readers', '区块链资源/lscc/GetChaincodeData', 'permit', '区块链参与者', '2020-02-13 18:08:07');
INSERT INTO `taro_policy` VALUES (54, '区块链访问控制策略', 'Channel/Application/Readers', '区块链资源/lscc/GetInstantiatedChaincodes', 'permit', '区块链参与者', '2020-02-13 18:08:22');
INSERT INTO `taro_policy` VALUES (55, '区块链访问控制策略', 'Channel/Application/Readers', '区块链资源/qscc/GetChainInfo', 'permit', '区块链参与者', '2020-02-13 18:08:51');
INSERT INTO `taro_policy` VALUES (56, '区块链访问控制策略', 'Channel/Application/Readers', '区块链资源/qscc/GetBlockByNumber', 'permit', '区块链参与者', '2020-02-13 18:09:14');
INSERT INTO `taro_policy` VALUES (57, '区块链访问控制策略', 'Channel/Application/Readers', '区块链资源/qscc/GetBlockByHash', 'permit', '区块链参与者', '2020-02-13 18:09:33');
INSERT INTO `taro_policy` VALUES (58, '区块链访问控制策略', 'Channel/Application/Readers', '区块链资源/qscc/GetTransactionByID', 'permit', '区块链参与者', '2020-02-13 18:09:46');
INSERT INTO `taro_policy` VALUES (59, '区块链访问控制策略', 'Channel/Application/Readers', '区块链资源/qscc/GetBlockByTxID', 'permit', '区块链参与者', '2020-02-13 18:10:12');
INSERT INTO `taro_policy` VALUES (60, '区块链访问控制策略', 'Channel/Application/Readers', '区块链资源/cscc/GetConfigBlock', 'permit', '区块链参与者', '2020-02-13 18:10:39');
INSERT INTO `taro_policy` VALUES (61, '区块链访问控制策略', 'Channel/Application/Readers', '区块链资源/cscc/GetConfigTree', 'permit', '区块链参与者', '2020-02-13 18:10:59');
INSERT INTO `taro_policy` VALUES (62, '区块链访问控制策略', 'Channel/Application/Readers', '区块链资源/cscc/SimulateConfigTreeUpdate', 'permit', '区块链参与者', '2020-02-13 18:11:17');
INSERT INTO `taro_policy` VALUES (63, '区块链访问控制策略', 'Channel/Application/Writers', '区块链资源/peer/Propose', 'permit', '区块链参与者', '2020-02-13 18:11:48');
INSERT INTO `taro_policy` VALUES (64, '区块链访问控制策略', 'Channel/Application/Readers', '区块链资源/peer/ChaincodeToChaincode', 'permit', '区块链参与者', '2020-02-13 18:12:07');
INSERT INTO `taro_policy` VALUES (83, '区块链访问控制策略', 'Channel/Application/Readers', '区块链资源/event/FilteredBlock', 'permit', '区块链参与者', '2020-02-14 19:50:49');
INSERT INTO `taro_policy` VALUES (84, '区块链访问控制策略', 'Channel/Application/Writers', '区块链资源/event/Block', 'permit', '区块链参与者', '2020-02-14 19:51:23');
INSERT INTO `taro_policy` VALUES (95, '跨物联网服务和区块链的业务流程策略', 'org1/peer1', '区块链资源/peer/Propose', 'read', '区块链参与者', '2020-02-25 16:55:49');
INSERT INTO `taro_policy` VALUES (96, '跨物联网服务和区块链的业务流程策略', '值班调度部/员工', '物联网资源/表/日报', '查看', '物联网用户', '2020-02-25 16:57:38');
INSERT INTO `taro_policy` VALUES (97, '跨物联网服务和区块链的业务流程策略', 'org2/order2', '区块链资源/event/Block', 'read', '区块链参与者', '2020-02-25 16:58:53');
INSERT INTO `taro_policy` VALUES (98, '跨物联网服务和区块链的业务流程策略', '后台运维部/员工', '物联网资源/表/日报', '修改', '物联网用户', '2020-02-25 17:04:07');
INSERT INTO `taro_policy` VALUES (99, '跨物联网服务和区块链的业务流程策略', 'org1/peer1', '区块链资源/event/FilteredBlock', 'write', '区块链参与者', '2020-02-25 17:04:35');
INSERT INTO `taro_policy` VALUES (106, '策略1', '销售部/电商部/经理', '物联网资源/工具', 'write', '物联网用户', '2020-03-09 14:15:44');
INSERT INTO `taro_policy` VALUES (121, '策略2', '员工', '物联网资源/报表', 'write', '物联网用户', '2020-03-09 14:15:44');
INSERT INTO `taro_policy` VALUES (122, '策略2', '员工', '物联网资源/工具', 'write', '物联网用户', '2020-03-09 14:15:44');
INSERT INTO `taro_policy` VALUES (123, '策略2', '员工', '资源1', 'write', '物联网用户', '2020-03-09 14:15:44');
INSERT INTO `taro_policy` VALUES (155, '策略2', '管理员', '物联网资源/报表', 'write', '物联网用户', '2020-03-09 14:15:44');
INSERT INTO `taro_policy` VALUES (156, '策略2', '管理员', '物联网资源/工具', 'write', '物联网用户', '2020-03-09 14:15:44');
INSERT INTO `taro_policy` VALUES (157, '策略2', '管理员', '物联网资源/态势图', 'write', '物联网用户', '2020-03-09 14:15:44');
INSERT INTO `taro_policy` VALUES (158, '运控分系统访问控制策略集合', '指挥大厅/员工', '物联网资源/图/中国地图', '查看', '物联网用户', '2020-04-16 18:06:55');
INSERT INTO `taro_policy` VALUES (159, '运控分系统访问控制策略集合', '指挥大厅/员工', '物联网资源/图/阿伦方差曲线', '查看', '物联网用户', '2020-04-16 18:06:55');
INSERT INTO `taro_policy` VALUES (160, '运控分系统访问控制策略集合', '指挥大厅/员工', '物联网资源/图/相位噪声曲线', '查看', '物联网用户', '2020-04-16 18:06:55');
INSERT INTO `taro_policy` VALUES (161, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/图/中国地图', '查看', '物联网用户', '2020-04-16 18:07:11');
INSERT INTO `taro_policy` VALUES (162, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/图/阿伦方差曲线', '查看', '物联网用户', '2020-04-16 18:07:15');
INSERT INTO `taro_policy` VALUES (163, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/图/相位噪声曲线', '查看', '物联网用户', '2020-04-16 18:07:17');
INSERT INTO `taro_policy` VALUES (164, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/图/线路地图', '查看', '物联网用户', '2020-04-16 18:07:20');
INSERT INTO `taro_policy` VALUES (165, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/图/设备关系图', '查看', '物联网用户', '2020-04-16 18:07:22');
INSERT INTO `taro_policy` VALUES (166, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/图/故障树', '查看', '物联网用户', '2020-04-16 18:07:24');
INSERT INTO `taro_policy` VALUES (167, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/图/参数信息态势', '查看', '物联网用户', '2020-04-16 18:07:28');
INSERT INTO `taro_policy` VALUES (168, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/图/光纤布局图', '查看', '物联网用户', '2020-04-16 18:07:31');
INSERT INTO `taro_policy` VALUES (169, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/图/OTDR图', '查看', '物联网用户', '2020-04-16 18:07:33');
INSERT INTO `taro_policy` VALUES (170, '运控分系统访问控制策略集合', '设备运维部/员工', '物联网资源/图/故障树', '修改', '物联网用户', '2020-04-16 18:07:35');
INSERT INTO `taro_policy` VALUES (171, '运控分系统访问控制策略集合', '光纤运维部/员工', '物联网资源/图/光纤布局图', '修改', '物联网用户', '2020-04-16 18:07:38');
INSERT INTO `taro_policy` VALUES (172, '运控分系统访问控制策略集合', '光纤运维部/员工', '物联网资源/图/OTDR图', '修改', '物联网用户', '2020-04-16 18:07:41');
INSERT INTO `taro_policy` VALUES (173, '运控分系统访问控制策略集合', '指挥大厅/员工', '物联网资源/表/稳定度指标', '查看', '物联网用户', '2020-04-16 18:08:20');
INSERT INTO `taro_policy` VALUES (174, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/表/稳定度指标', '查看', '物联网用户', '2020-04-16 18:08:22');
INSERT INTO `taro_policy` VALUES (175, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/表/设备运行状态', '查看', '物联网用户', '2020-04-16 18:08:23');
INSERT INTO `taro_policy` VALUES (176, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/表/参数状态表', '查看', '物联网用户', '2020-04-16 18:08:26');
INSERT INTO `taro_policy` VALUES (177, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/表/设备属性', '查看', '物联网用户', '2020-04-16 18:08:29');
INSERT INTO `taro_policy` VALUES (178, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/表/光纤属性', '查看', '物联网用户', '2020-04-16 18:08:32');
INSERT INTO `taro_policy` VALUES (179, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/表/日报', '查看', '物联网用户', '2020-04-16 18:08:34');
INSERT INTO `taro_policy` VALUES (180, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/表/月报', '查看', '物联网用户', '2020-04-16 18:08:37');
INSERT INTO `taro_policy` VALUES (181, '运控分系统访问控制策略集合', '值班调度部/员工', '物联网资源/表/年报', '查看', '物联网用户', '2020-04-16 18:08:39');
INSERT INTO `taro_policy` VALUES (182, '运控分系统访问控制策略集合', '设备运维部/员工', '物联网资源/表/设备属性', '修改', '物联网用户', '2020-04-16 18:08:40');
INSERT INTO `taro_policy` VALUES (183, '运控分系统访问控制策略集合', '设备运维部/员工', '物联网资源/表/故障报表', '修改', '物联网用户', '2020-04-16 18:08:42');
INSERT INTO `taro_policy` VALUES (184, '运控分系统访问控制策略集合', '光纤运维部/员工', '物联网资源/表/光纤属性', '修改', '物联网用户', '2020-04-16 18:08:45');
INSERT INTO `taro_policy` VALUES (185, '跨物联网服务和区块链的业务流程策略', '后台运维部/员工', '物联网资源/表/稳定度指标', '查看', '物联网用户', NULL);
INSERT INTO `taro_policy` VALUES (186, '跨物联网服务和区块链的业务流程策略', '后台运维部/员工', '物联网资源/表/设备运行状态', '查看', '物联网用户', NULL);
INSERT INTO `taro_policy` VALUES (187, '跨物联网服务和区块链的业务流程策略', '后台运维部/员工', '物联网资源/表/参数状态表', '查看', '物联网用户', NULL);
INSERT INTO `taro_policy` VALUES (188, '跨物联网服务和区块链的业务流程策略', '后台运维部/员工', '物联网资源/图/中国地图', '查看', '物联网用户', NULL);
INSERT INTO `taro_policy` VALUES (189, '跨物联网服务和区块链的业务流程策略', '后台运维部/员工', '物联网资源/图/阿伦方差曲线', '查看', '物联网用户', NULL);
INSERT INTO `taro_policy` VALUES (190, '跨物联网服务和区块链的业务流程策略', '后台运维部/员工', '物联网资源/图/相位噪声曲线', '查看', '物联网用户', NULL);
INSERT INTO `taro_policy` VALUES (192, '运控分系统访问控制策略集合', '值班调度人员', '物联网资源/业务流程/设备故障0715', '修改', '物联网用户', NULL);

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
  `user_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户地址',
  `user_email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '电子邮箱',
  `user_phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系方式',
  `user_status` int(1) NULL DEFAULT NULL COMMENT '用户状态',
  `user_hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户哈希',
  `user_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户证书存储路径',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of taro_user
-- ----------------------------
INSERT INTO `taro_user` VALUES (1, 'zhao', '管理员#', '西土城路9号', 'zky0825@163.com', '13912345678', 1, '', 'D:/goProjects/tarobackend/test/');
INSERT INTO `taro_user` VALUES (2, 'kong', '后台运维部/员工#', '西土城路9号', '470967263@qq.com', '13912334678', 0, '', 'E:/');
INSERT INTO `taro_user` VALUES (3, 'yang', '值班调度部/管理员#领导小组/员工', '西土城路9号', 'dfdf@qq.com', '13913545678', 0, '', 'D:/goProjects/tarobackend/test/');
INSERT INTO `taro_user` VALUES (22, '刘杰', '光纤运维部/管理员#', '西安', '470967263@qq.com', '18000000000', 0, '', 'E:/');
INSERT INTO `taro_user` VALUES (23, '刘涛', '值班调度人员#', '西安', 'zky0825@qq.com', '18111111111', 1, '123bb05baf635d586df3d37ab5bbb9d6', 'D:/');

SET FOREIGN_KEY_CHECKS = 1;
