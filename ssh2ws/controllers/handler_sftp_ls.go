package controllers

import (
	"github.com/dejavuzhou/felix/flx"
	"github.com/dejavuzhou/felix/models"
	"github.com/dejavuzhou/felix/ssh2ws/utils"
	"net/http"
	"path"
)

type Ls struct {
	Name  string `json:"name"`
	Path  string `json:"path"` // including Name
	IsDir bool   `json:"is_dir"`
}

func SftpLs(w http.ResponseWriter, r *http.Request) {
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
		if files, err := client.ReadDir(fullPath); err != nil {
			response.Message = "no such path"
		} else {
			response.HasError = false
			fileList := make([]Ls, 0) // this will not be converted to null if slice is empty.
			for _, file := range files {
				fileList = append(fileList, Ls{Name: file.Name(), IsDir: file.IsDir(), Path: path.Join(relativePath, file.Name())})
			}
			response.Message = fileList
		}
	} else {
		response.Message = "no such path"
	}
	utils.ServeJSON(w, response)
}
