package testfsnotify

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func WatchFile(fileName string) {
	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// 要监控的文件名
	filePath := filepath.Join(currentDir, fileName)

	// 创建文件监视器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// 启动一个goroutine来处理文件系统事件
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Name == filePath {
					log.Println("event:", event)
					if event.Has(fsnotify.Write) {
						log.Println("modified file:", event.Name)
					}
					if event.Has(fsnotify.Create) {
						log.Println("created file:", event.Name)
					}
					if event.Has(fsnotify.Remove) {
						log.Println("deleted file:", event.Name)
					}
					if event.Has(fsnotify.Rename) {
						log.Println("renamed file:", event.Name)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// 添加文件到监控列表
	err = watcher.Add(currentDir)
	if err != nil {
		log.Fatal(err)
	}

	// 阻塞主goroutine，使程序保持运行状态
	<-make(chan struct{})
}
