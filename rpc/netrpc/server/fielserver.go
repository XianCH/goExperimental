package server

import "os"

type FileArgs struct {
	FileName string
}

type FileServer int

func (t *FileServer) ReadTextFile(args *FileArgs, reply *[]byte) error {
	_, err := os.Stat(args.FileName)
	if err != nil {
		return err
	}
	return nil
}
