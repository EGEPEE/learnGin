package logging

import (
	"fmt"
	"time"

	"github.com/EGEPEE/learnGin/models"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", models.AppSetting.RuntimeRootPath, models.AppSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		models.AppSetting.LogSaveName,
		time.Now().Format(models.AppSetting.TimeFormat),
		models.AppSetting.LogFileExt,
	)
}

//func openLogFile(fileName, filePath string) (*os.File, error) {
//	dir, err := os.Getwd()
//	if err != nil {
//		return nil, fmt.Errorf("os.Getwd err: %v", err)
//	}
//
//	src := dir + "/" + filePath
//	perm := file.CheckPermission(src)
//	if perm == true {
//		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
//	}
//
//	err = file.IsNotExistMkDir(src)
//	if err != nil {
//		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
//	}
//
//	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
//	}
//
//	return f, nil
//}
