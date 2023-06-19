package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	fmt.Println("服务启动成功")
	http.HandleFunc("/upload/", uploadHandle)    // 上传
	http.HandleFunc("/uploaded/", showPicHandle) //显示图片
	err := http.ListenAndServe(":7771", nil)
	fmt.Println(err)
}

// 上传图像接口
func uploadHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("接受到请求")
	setupCORS(&w)
	w.Header().Set("Content-Type", "text/html")

	req.ParseForm()

	// 接收图片
	uploadFile, handle, err := req.FormFile("image")
	if err != nil {
		errorHandle(errors.New("接收图片出错"), w)
	}

	// 检查图片后缀
	ext := strings.ToLower(path.Ext(handle.Filename))
	if ext != ".jpg" && ext != ".png" {
		errorHandle(errors.New("只支持jpg/png图片上传"), w)
		return
		//defer os.Exit(2)
	}

	// 保存图片
	os.Mkdir("./uploaded/", 0777)
	saveFile, err := os.OpenFile("./uploaded/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	errorHandle(err, w)
	io.Copy(saveFile, uploadFile)

	defer uploadFile.Close()
	defer saveFile.Close()
	// 上传图片成功
	w.Write(([]byte("/uploaded/" + handle.Filename)))

}

// 显示图片接口
func showPicHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("接受到请求")
	setupCORS(&w)

	file, err := os.Open("." + req.URL.Path)
	errorHandle(err, w)

	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	errorHandle(err, w)
	w.Write(buff)
}

// 统一错误输出接口
func errorHandle(err error, w http.ResponseWriter) {

	if err != nil {
		w.Write([]byte(errors.New("请求出错").Error()))
	}
}

func setupCORS(w *http.ResponseWriter) {
	(*w).Header().Add("Access-Control-Allow-Origin", "*")
	(*w).Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
