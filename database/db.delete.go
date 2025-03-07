package database

import "strings"

// deleteTestWalletItem -> using to delete all data after testing
func (s *DatabaseService) deleteTestWalletItem(tableName string) error {
	s.db = dbConnect()
	defer s.db.Close()
	sql1 := strings.Join([]string{"DELETE FROM ", tableName, " WHERE userId=90990"}, "")
	sql2 := strings.Join([]string{"DELETE FROM ", tableName, " WHERE userId=832918"}, "")

	err := s.db.QueryRow(sql1).Err()
	if err != nil {
		return err
	}
	err = s.db.QueryRow(sql2).Err()
	return err
}
