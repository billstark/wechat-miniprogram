package testutils

const (
	// RecordTableClear clears record table
	RecordTableClear = `DELETE FROM Record`
	// UserTableClear clears user table
	UserTableClear = `DELETE FROM MUser`
	// GroupTableClear clears group table
	GroupTableClear = `DELETE FROM MGroup`
	// InsertUserA inserts a sample user A
	InsertUserA = `INSERT INTO MUser VALUES('12345u')`
	// InsertUserB inserts a sample user B
	InsertUserB = `INSERT INTO MUser VALUES('23456u')`
	// InsertUserC inserts a sample user C
	InsertUserC = `INSERT INTO MUser VALUES('34567u')`
	// InsertGroupA inserts a sample group A
	InsertGroupA = `INSERT INTO MGroup VALUES('292 Pasir Panjang Road', 'This is just a test group')`
	// InsertRecordOne inserts a sample record
	InsertRecordOne = `INSERT INTO Record (g_id, day, payer, spliters, pay_amount, description, updated_at) VALUES(1, '2018-07-14', '12345u', '{"12345u", "23456u"}', 100, 'dinner', '2018-07-14 20:38:40')`
	// InsertRecordTwo inserts a sample record
	InsertRecordTwo = `INSERT INTO Record (g_id, day, payer, spliters, pay_amount, description, updated_at) VALUES(1, '2016-07-14', '12345u', '{"12345u", "34567u"}', 100, 'dinner', '2018-07-14 20:38:40')`
	// InsertRecordThree inserts a sample record
	InsertRecordThree = `INSERT INTO Record (g_id, day, payer, spliters, pay_amount, description, updated_at) VALUES(1, '2016-07-14', '23456u', '{"12345u"}', 30, 'settle', '2018-07-14 20:38:40')`
)
