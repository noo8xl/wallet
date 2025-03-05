package database

import "strings"

func (s *DatabaseService) deleteTestWalletItem(tableName string) error {
	s.db = dbConnect()
	defer s.db.Close()
	sql := strings.Join([]string{"DELETE FROM ", tableName, " WHERE customerId=90990"}, "")

	err := s.db.QueryRow(sql).Err()
	return err
}
