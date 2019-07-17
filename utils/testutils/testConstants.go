package testutils

const (
	// CreateUserTable creates user table
	CreateUserTable = `CREATE TABLE MUser (
		id        SERIAL,
		w_id      varchar(255) PRIMARY KEY
	);`
	// CreateGroupTable creates group table
	CreateGroupTable = `CREATE TABLE MGroup (
		id          SERIAL PRIMARY KEY,
		name        varchar(255) NOT NULL,
		description text
	);`
	// CreateBelongToTable creates belong to table
	CreateBelongToTable = `CREATE TABLE BelongTo (
		id              SERIAL,
		u_id            varchar(255) NOT NULL,
		g_id            int NOT NULL,
		FOREIGN KEY (u_id) REFERENCES MUser(w_id) ON DELETE CASCADE,
		FOREIGN KEY (g_id) REFERENCES MGroup(id) ON DELETE CASCADE,
		PRIMARY KEY (id, u_id, g_id)
	);`
	// CreateRecordTable creates record table
	CreateRecordTable = `CREATE TABLE Record (
		id          SERIAL,
		g_id        int,
		day         date NOT NULL,
		payer       varchar(255) NOT NULL,
		spliters    varchar[],
		pay_amount  float NOT NULL,
		description text NOT NULL,
		updated_at  timestamp,
		deleted     boolean DEFAULT false,
		FOREIGN KEY (g_id) REFERENCES MGroup(id) ON DELETE CASCADE,
		FOREIGN KEY (payer) REFERENCES MUser(w_id) ON DELETE CASCADE,
		PRIMARY KEY (id, g_id)
	);`
	// CreateHistoryTable creates history table
	CreateHistoryTable = `CREATE TABLE OpHistory (
		id         SERIAL,
		u_id       varchar(255) NOT NULL,
		g_id       int,
		message    text NOT NULL,
		created_at timestamp,
		FOREIGN KEY (u_id) REFERENCES MUser(w_id) ON DELETE CASCADE,
		FOREIGN KEY (g_id) REFERENCES MGroup(id) ON DELETE CASCADE,
		PRIMARY KEY (id, u_id, g_id)
	);`

	// DropUserTable drops user talbe
	DropUserTable = `DROP TABLE muser`
	// DropGroupTable drops group talbe
	DropGroupTable = `DROP TABLE mgroup`
	// DropBelongToTable drops belong to talbe
	DropBelongToTable = `DROP TABLE belongto`
	// DropRecordTable drops record talbe
	DropRecordTable = `DROP TABLE record`
	// DropHistoryTable drops history talbe
	DropHistoryTable = `DROP TABLE history`
	// RecordTableClear clears record table
	RecordTableClear = `DELETE FROM Record`
	// UserTableClear clears user table
	UserTableClear = `DELETE FROM MUser`
	// GroupTableClear clears group table
	GroupTableClear = `DELETE FROM MGroup`
	// InsertUserA inserts a sample user A
	InsertUserA = `INSERT INTO MUser (w_name, w_id) VALUES ('test user1', '12345u')`
	// InsertUserB inserts a sample user B
	InsertUserB = `INSERT INTO MUser (w_name, w_id) VALUES ('test user2', '23456u')`
	// InsertUserC inserts a sample user C
	InsertUserC = `INSERT INTO MUser (w_name, w_id) VALUES('test user3', '34567u')`
	// InsertGroupA inserts a sample group A
	InsertGroupA = `INSERT INTO MGroup (name, description) VALUES('292 Pasir Panjang Road', 'This is just a test group')`
	// InsertRecordOne inserts a sample record
	InsertRecordOne = `INSERT INTO Record (g_id, day, payer, spliters, pay_amount, description, updated_at) VALUES(1, '2018-07-14', '12345u', '{"12345u", "23456u"}', 100, 'dinner', '2018-07-14 20:38:40')`
	// InsertRecordTwo inserts a sample record
	InsertRecordTwo = `INSERT INTO Record (g_id, day, payer, spliters, pay_amount, description, updated_at) VALUES(1, '2016-07-14', '12345u', '{"12345u", "34567u"}', 100, 'dinner', '2018-07-14 20:38:40')`
	// InsertRecordThree inserts a sample record
	InsertRecordThree = `INSERT INTO Record (g_id, day, payer, spliters, pay_amount, description, updated_at) VALUES(1, '2016-07-14', '23456u', '{"12345u"}', 30, 'settle', '2018-07-14 20:38:40')`
	// ResetUserSeq resets user sequence
	ResetUserSeq = `ALTER SEQUENCE muser_id_seq RESTART;`
	// ResetGroupSeq resets group sequence
	ResetGroupSeq = `ALTER SEQUENCE mgroup_id_seq RESTART;`
	// ResetRecordSeq resets group sequence
	ResetRecordSeq = `ALTER SEQUENCE record_id_seq RESTART;`
)
