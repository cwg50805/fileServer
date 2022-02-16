package v1

import (
	"CreativeServer/constant"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func ConcatentatePath(localSystemFilePath string) string {
	ex, _ := os.Executable()

	localSystemFilePath = path.Join(filepath.Dir(ex), localSystemFilePath)
	return localSystemFilePath
}

func GetFile(c *gin.Context) {
	localSystemFilePath := c.Param("localSystemFilePath")
	localSystemFilePath = ConcatentatePath(localSystemFilePath)

	fileInfo, err := os.Stat(localSystemFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			constant.ResponseWithData(c, http.StatusOK, constant.FILENOTEXIST, err)
		} else {
			constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err)
		}
		return
	}

	dirContent := []string{}
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(localSystemFilePath)
		if err != nil {
			constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err)
			return
		}
		for _, file := range files {
			dirContent = append(dirContent, file.Name())
		}
	} else {
		byteFile, err := ioutil.ReadFile(localSystemFilePath)
		if err != nil {
			constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err)
			return
		}

		c.Header("Content-Disposition", "attachment; filename=file-name.txt")
		c.Data(http.StatusOK, "application/octet-stream", byteFile)
	}
	constant.ResponseWithData(c, http.StatusOK, constant.SUCCESS, gin.H{"path": dirContent})
}

func CreateFile(c *gin.Context) {
	localSystemFilePath := c.Param("localSystemFilePath")
	localSystemFilePath = ConcatentatePath(localSystemFilePath)

	if _, err := os.Stat(localSystemFilePath); !errors.Is(err, os.ErrNotExist) {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err)
		return
	}

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err)
		return
	}
	src, err := file.Open()
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err)
		return
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(localSystemFilePath)
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err)
		return
	}
	defer dst.Close()
	io.Copy(dst, src)
	constant.ResponseWithData(c, http.StatusOK, constant.SUCCESS, gin.H{"path": localSystemFilePath})
}

func UpdateFile(c *gin.Context) {
	localSystemFilePath := c.Param("localSystemFilePath")
	localSystemFilePath = ConcatentatePath(localSystemFilePath)

	if _, err := os.Stat(localSystemFilePath); errors.Is(err, os.ErrNotExist) {
		constant.ResponseWithData(c, http.StatusOK, constant.FILENOTEXIST, err)
		return
	}

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err)
		return
	}
	src, err := file.Open()
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err)
		return
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(localSystemFilePath)
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err)
		return
	}
	defer dst.Close()

	// Copy
	io.Copy(dst, src)
	constant.ResponseWithData(c, http.StatusOK, constant.SUCCESS, gin.H{"path": localSystemFilePath})
}

func DeleteFile(c *gin.Context) {
	localSystemFilePath := c.Param("localSystemFilePath")
	localSystemFilePath = ConcatentatePath(localSystemFilePath)

	if _, err := os.Stat(localSystemFilePath); errors.Is(err, os.ErrNotExist) {
		constant.ResponseWithData(c, http.StatusOK, constant.FILENOTEXIST, err)
		return
	}

	if err := os.Remove(localSystemFilePath); err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err)
		return
	}

	constant.ResponseWithData(c, http.StatusOK, constant.SUCCESS, gin.H{"path": localSystemFilePath})
}
