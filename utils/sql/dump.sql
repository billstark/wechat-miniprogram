-- reset id sequence
ALTER SEQUENCE muser_id_seq  RESTART WITH 1;
ALTER SEQUENCE mgroup_id_seq RESTART WITH 1;

-- insert users
INSERT INTO MUser (w_name, w_id, created_at) VALUES ('方头', 'w00001', '2019-07-17T10:23:31+00:00');
INSERT INTO MUser (w_name, w_id, created_at) VALUES ('风神', 'w00002', '2019-07-17T10:23:31+00:00');
INSERT INTO MUser (w_name, w_id, created_at) VALUES ('灰灰', 'w00004', '2019-07-17T10:23:31+00:00');
INSERT INTO MUser (w_name, w_id, created_at) VALUES ('老白', 'w00003', '2019-07-17T10:23:31+00:00');
INSERT INTO MUser (w_name, w_id, created_at) VALUES ('丁樊', 'w00005', '2019-07-17T10:23:31+00:00');
INSERT INTO MUser (w_name, w_id, created_at) VALUES ('王府狗', 'w00006', '2019-07-17T10:23:31+00:00');
INSERT INTO MUser (w_name, w_id, created_at) VALUES ('租一居', 'w00007', '2019-07-17T10:23:31+00:00');
INSERT INTO MUser (w_name, w_id, created_at) VALUES ('杨大爷', 'w00008', '2019-07-17T10:23:31+00:00');
INSERT INTO MUser (w_name, w_id, created_at) VALUES ('裤裤', 'w00009', '2019-07-17T10:23:31+00:00');
INSERT INTO MUser (w_name, w_id, created_at) VALUES ('撒谜题', 'w00010', '2019-07-17T10:23:31+00:00');

-- insert groups
INSERT INTO MGroup (name, description, updated_at) VALUES('N\A', 'N\A', '2019-07-17T10:23:31+00:00');
INSERT INTO MGroup (name, description, updated_at) VALUES('黄狗', '母婴用品', '2019-07-17T10:23:31+00:00');
INSERT INTO MGroup (name, description, updated_at) VALUES('292 Pasir Panjang', '吃喝玩乐', '2019-07-17T10:23:31+00:00');
INSERT INTO MGroup (name, description, updated_at) VALUES('ti9 约起来', '看Ti', '2019-07-17T10:23:31+00:00');

-- relations
INSERT INTO BelongTo (u_id, g_id) VALUES('w00001', 2);
INSERT INTO BelongTo (u_id, g_id) VALUES('w00001', 3);
INSERT INTO BelongTo (u_id, g_id) VALUES('w00001', 4);
INSERT INTO BelongTo (u_id, g_id) VALUES('w00002', 2);
INSERT INTO BelongTo (u_id, g_id) VALUES('w00003', 2);
INSERT INTO BelongTo (u_id, g_id) VALUES('w00004', 2);
INSERT INTO BelongTo (u_id, g_id) VALUES('w00005', 2);
INSERT INTO BelongTo (u_id, g_id) VALUES('w00006', 3);
INSERT INTO BelongTo (u_id, g_id) VALUES('w00007', 3);
INSERT INTO BelongTo (u_id, g_id) VALUES('w00008', 3);
INSERT INTO BelongTo (u_id, g_id) VALUES('w00009', 3);
INSERT INTO BelongTo (u_id, g_id) VALUES('w00010', 3);
INSERT INTO BelongTo (u_id, g_id) VALUES('w00006', 4);
INSERT INTO BelongTo (u_id, g_id) VALUES('w00007', 4);

-- dummy records
-- personal records
INSERT INTO Record (g_id, day, payer, spliters, pay_amount, updated_at, description) VALUES (1, '2019-07-15', 'w00001', '{"w00002"}', 6, '2019-07-16T10:02:35+00:00', '带饭');
INSERT INTO Record (g_id, day, payer, spliters, pay_amount, updated_at, description) VALUES (1, '2019-07-15', 'w00001', '{"w00003"}', 5, '2019-07-16T10:02:35+00:00', '带饭');
INSERT INTO Record (g_id, day, payer, spliters, pay_amount, updated_at, description) VALUES (1, '2019-07-16', 'w00001', '{"w00002"}', 6, '2019-07-16T10:02:35+00:00', '带饭');
INSERT INTO Record (g_id, day, payer, spliters, pay_amount, updated_at, description) VALUES (2, '2019-07-13', 'w00001', '{"w00002", "w00003", "w00004", "w000005"}', 48, '2019-07-16T10:02:35+00:00', '午饭外卖');
INSERT INTO Record (g_id, day, payer, spliters, pay_amount, updated_at, description) VALUES (2, '2019-07-14', 'w00001', '{"w00002", "w00003", "w00004", "w000005"}', 52, '2019-07-16T10:02:35+00:00', '晚饭外卖');
INSERT INTO Record (g_id, day, payer, spliters, pay_amount, updated_at, description) VALUES (4, '2019-07-16', 'w00006', '{"w00006", "w00001", "w00007"}', 2492, '2019-07-16T10:02:35+00:00', '订房间');
INSERT INTO Record (g_id, day, payer, spliters, pay_amount, updated_at, description) VALUES (4, '2019-07-16', 'w00006', '{"w00006", "w00001", "w00007"}', 599, '2019-07-16T10:02:35+00:00', '观战费');