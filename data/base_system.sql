-- name: create-admin
INSERT INTO `admin` (`id`, `login_name`, `real_name`, `password`, `role_ids`, `phone`, `email`, `avatar`, `salt`, `last_ip`, `last_login`, `status`, `create_id`, `update_id`, `created_at`, `updated_at`, `level`, `deleted_at`)
VALUES
	(1,'admin','admin','8ebf088c456f925ce22b881eff76bec7','0','','','/static/images/users/avatar-1.png','BtBFNwXINC','::1','2020-10-15 17:07:29',1,0,0,'2020-09-27 17:44:35','2020-10-15 17:07:29',99,NULL),
	(2,'ceshi','测试账号','fa3fb5825c2e64bc764f29245dd1ec7a','1,2','13988009988','abc@188.com',NULL,'i8Nf','','2020-01-01 15:01:01',1,1,0,NULL,'2020-10-09 11:38:42',1,NULL);

-- name: create-auth
INSERT INTO `auth` (`id`, `auth_name`, `auth_url`, `user_id`, `pid`, `sort`, `icon`, `is_show`, `status`, `create_id`, `update_id`, `created_at`, `updated_at`)
VALUES
	(1,'所有权限','/',1,0,1,'',0,1,1,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(2,'权限管理','/',1,1,999,'mdi-graph-outline',1,1,0,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(3,'管理员','/system/admin/list',1,2,1,'fa-user-o',1,1,1,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(4,'角色管理','/system/role/list',1,2,2,'fa-user-circle-o',1,1,0,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(5,'新增','/system/admin/add',1,3,1,'',0,1,0,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(6,'修改','/system/admin/edit',1,3,2,'',0,1,0,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(7,'删除','/system/admin/delete',1,3,3,'',0,1,1,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(8,'新增','/system/role/add',1,4,1,'',1,1,0,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(9,'修改','/system/role/edit',1,4,2,'',0,1,1,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(10,'删除','/system/role/delete',1,4,3,'',0,1,1,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(11,'权限因子','/system/auth/list',1,2,3,'fa-list',1,1,1,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(12,'新增','/system/auth/edit',1,11,1,'',0,1,1,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(13,'修改','/system/auth/edit',1,11,2,'',0,1,1,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(14,'删除','/system/auth/delete',1,11,3,'',0,1,1,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(15,'个人中心','profile/edit',1,1,1001,'mdi-guy-fawkes-mask',1,1,0,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(20,'基础设置','/',1,1,2,'mdi-reddit',1,1,1,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(24,'资料修改','/system/user/edit',1,15,1,'fa-edit',1,1,0,1,'2020-10-14 17:04:30','2020-10-14 17:04:30'),
	(45,'权限树','/system/auth/get_nodes',0,11,4,'',0,1,0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(46,'单个权限获取','/system/auth/get_node',0,11,5,'',0,1,0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(47,'管理员列表','/system/admin/list_json',0,3,4,'',0,1,0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(48,'角色列表','/system/role/list_json',0,4,4,'',0,1,0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(49,'站点设置','/system/site/list',0,20,1,'',1,1,0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(50,'上传图片','/system/upload/def_upload',0,49,1,'',0,1,0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(51,'更新站点设置','/system/site/edit',0,49,2,'',0,1,0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(52,'邮件配置页面','/system/mail/list',0,49,2,'',0,1,0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(53,'更新邮件配置','/system/mail/edit',0,49,4,'',0,1,0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(54,'测试邮件发送','/system/mail/test',0,49,5,'',0,1,0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(55,'更新头像','/system/user/avatar',0,24,2,'',0,1,0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00'),
	(56,'修改密码','/system/user/pwd',0,15,2,'',1,1,0,0,'0000-00-00 00:00:00','0000-00-00 00:00:00');

-- name: create-role
INSERT INTO `role` (`id`, `role_name`, `detail`, `status`, `create_id`, `update_id`, `created_at`, `updated_at`)
VALUES
	(1,'API管理员','拥有API所有权限',1,0,2,'2020-09-28 13:59:53','2020-09-28 14:00:03'),
	(2,'系统管理员','系统管理员',1,0,0,'2020-09-28 14:00:10','2020-09-28 14:00:06');

-- name: create-role-auth
INSERT INTO `role_auth` (`role_id`, `auth_id`)
VALUES
	(1,16),
	(1,17),
	(1,18),
	(1,19),
	(2,0),
	(2,1),
	(2,15),
	(2,20),
	(2,21),
	(2,22),
	(2,23),
	(2,24);
