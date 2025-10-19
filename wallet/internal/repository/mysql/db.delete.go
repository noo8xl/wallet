package mysql

import "strings"

// DeleteTestWalletItem -> using to delete all data after testing
func (s *DatabaseService) DeleteTestWalletItem(tableName string) error {

	sql1 := strings.Join([]string{"DELETE FROM ", tableName, " WHERE userId=90990"}, "")
	sql2 := strings.Join([]string{"DELETE FROM ", tableName, " WHERE userId=832918"}, "")

	err := s.db.QueryRow(sql1).Err()
	if err != nil {
		return err
	}
	err = s.db.QueryRow(sql2).Err()
	return err
}
