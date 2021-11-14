package files

import (
	"fmt"
	"os"
	"time"
)

type FileCreateError struct {
	created time.Time
	err     error
}

func (e FileCreateError) Error() string {
	return fmt.Sprintf("Ошибка создания файла в %v: %v", e.created, e.err)
}

func CreateFile(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return WrapFileCreateError(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	return nil
}

func WrapFileCreateError(err error) *FileCreateError {
	return &FileCreateError{
		created: time.Now(),
		err:     err,
	}
}
