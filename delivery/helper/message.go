package helper

import (
	"github.com/EGEPEE/learnGin/delivery/helper/logging"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST       = 10001
	ERROR_EXIST_FAIL  = 10002
	ERROR_NOT_EXIST   = 10003
	ERROR_GET_S_FAIL  = 10004
	ERROR_COUNT_FAIL  = 10005
	ERROR_ADD_FAIL    = 10006
	ERROR_EDIT_FAIL   = 10007
	ERROR_DELETE_FAIL = 10008
	ERROR_EXPORT_FAIL = 10009
	ERROR_IMPORT_FAIL = 10010

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)

var MsgFlags = map[int]string{
	SUCCESS:        "Ok",
	ERROR:          "Failed",
	INVALID_PARAMS: "Kesalahan parameter permintaan",

	ERROR_EXIST:       "Nama objek sudah ada",
	ERROR_EXIST_FAIL:  "Gagal mendapatkan objek yang ada",
	ERROR_NOT_EXIST:   "Objek tidak ada",
	ERROR_GET_S_FAIL:  "Gagal mendapatkan semua objek",
	ERROR_COUNT_FAIL:  "Objek statistik gagal",
	ERROR_ADD_FAIL:    "Gagal menambahkan objek",
	ERROR_EDIT_FAIL:   "Ubah objek gagal",
	ERROR_DELETE_FAIL: "Penghapusan objek gagal",
	ERROR_EXPORT_FAIL: "Ekspor objek gagal",
	ERROR_IMPORT_FAIL: "Gagal mengimpor objek",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Otentikasi token gagal",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token expired",
	ERROR_AUTH_TOKEN:               "Pembuatan token gagal",
	ERROR_AUTH:                     "Kesalahan token",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  GetMsg(errCode),
		"data": data,
	})

	return
}

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
}
