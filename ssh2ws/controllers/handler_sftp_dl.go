package controllers

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/models"
	"github.com/dejavuzhou/felix/ssh2ws/utils"
	"io"
	"net/http"
	"path"
)

func SftpDl(w http.ResponseWriter, r *http.Request) {
	response := models.JsonResponse{HasError: true}
	idx := utils.GetQueryInt(r, "id", 0)
	mc, err := models.MachineFind(uint(idx))
	if err != nil {
		response.Message = err
		utils.ServeJSON(w, response)
		return
	}
	client := flx.NewSftpClient(mc)
	if wd, err := client.Getwd(); err == nil {
		relativePath := r.URL.Query().Get("path") // get path.
		fullPath := path.Join(wd, relativePath)
		if fileInfo, err := client.Stat(fullPath); err == nil && !fileInfo.IsDir() {
			if file, err := client.Open(fullPath); err == nil {
				defer file.Close()
				w.Header().Add("Content-Disposition", "attachment;filename="+fileInfo.Name())
				w.Header().Add("Content-Type", "application/octet-stream")
				io.Copy(w, file)
				return
			}
		}
	}
	utils.Abort(w, "no such file", 400)
	return
}
