package service

import (
	"net/http"
	"fmt"
	"github.com/unrolled/render"
	// "html/template"
)

func loginHander(fomatformatter *render.Render ) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("method:", req.Method) //获取请求的方法
		if req.Method == "GET" {
			fomatformatter.HTML(w, http.StatusOK, "login", "")
			// t, _ := template.ParseFiles("login.")
			// log.Println(t.Execute(w, nil))
		} else {
			//请求的是登录数据，那么执行登录的逻辑判断
			req.ParseForm()
			fomatformatter.HTML(w, http.StatusOK, "userInfo", struct{
				Username string
				Password string
				List []struct{
					Id string
					Work string
				}
			}{
				Username: req.Form["username"][0],
				Password: req.Form["password"][0],
				List: []struct{
					Id string
					Work string
				} {
					{"1", "programmer"},
					{"2", "doctor"},
				},
			})
			fmt.Println("username:", req.Form["username"][0])
			fmt.Println("password:", req.Form["password"][0])
		}
	}
}

func apiTestHandler(formatter *render.Render) http.HandlerFunc {
	
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		}{ID: "8675309", Content: "Hello from Go!"})
	}
}

func unimplementHander(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "501 not implemented", http.StatusNotImplemented)
}