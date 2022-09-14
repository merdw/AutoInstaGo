package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/valyala/fasthttp"
)

type Anas struct {
	Message                   string                    `json:"message"`
	TwoFactorRequired         bool                      `json:"two_factor_required"`
	TwoFactorInfo             TwoFactorInfo             `json:"two_factor_info"`
	PhoneVerificationSettings PhoneVerificationSettings `json:"phone_verification_settings"`
	Status                    string                    `json:"status"`
	ErrorType                 string                    `json:"error_type"`
	ExceptionName             string                    `json:"exception_name"`
}
type PhoneVerificationSettings struct {
	MaxSmsCount              int  `json:"max_sms_count"`
	ResendSmsDelaySec        int  `json:"resend_sms_delay_sec"`
	RobocallCountDownTimeSec int  `json:"robocall_count_down_time_sec"`
	RobocallAfterMaxSms      bool `json:"robocall_after_max_sms"`
}
type TwoFactorInfo struct {
	Pk                              int                       `json:"pk"`
	Username                        string                    `json:"username"`
	SmsTwoFactorOn                  bool                      `json:"sms_two_factor_on"`
	WhatsappTwoFactorOn             bool                      `json:"whatsapp_two_factor_on"`
	TotpTwoFactorOn                 bool                      `json:"totp_two_factor_on"`
	EligibleForMultipleTotp         bool                      `json:"eligible_for_multiple_totp"`
	ObfuscatedPhoneNumber           string                    `json:"obfuscated_phone_number"`
	TwoFactorIdentifier             string                    `json:"two_factor_identifier"`
	ShowMessengerCodeOption         bool                      `json:"show_messenger_code_option"`
	ShowNewLoginScreen              bool                      `json:"show_new_login_screen"`
	ShowTrustedDeviceOption         bool                      `json:"show_trusted_device_option"`
	ShouldOptInTrustedDeviceOption  bool                      `json:"should_opt_in_trusted_device_option"`
	PendingTrustedNotification      bool                      `json:"pending_trusted_notification"`
	SmsNotAllowedReason             interface{}               `json:"sms_not_allowed_reason"`
	TrustedNotificationPollingNonce interface{}               `json:"trusted_notification_polling_nonce"`
	IsTrustedDevice                 bool                      `json:"is_trusted_device"`
	PhoneVerificationSettings       PhoneVerificationSettings `json:"phone_verification_settings"`
}

type weblos struct {
	Message        string `json:"message"`
	CheckpointURL  string `json:"checkpoint_url"`
	Lock           bool   `json:"lock"`
	FlowRenderType int    `json:"flow_render_type"`
	Status         string `json:"status"`
}

type webtwo struct {
	FormData FormData `json:"form_data"`
}
type FormData struct {
	PhoneNumber              string `json:"phone_number"`
	TwoFactorEnabled         bool   `json:"two_factor_enabled"`
	EligibleForTotpTwoFactor bool   `json:"eligible_for_totp_two_factor"`
	TotpTwoFactorEnabled     bool   `json:"totp_two_factor_enabled"`
	ShowTrustedDevices       bool   `json:"show_trusted_devices"`
}

var passtwo int = 5
var webgiris int = 0
var fac int = 0
var identifier string = "bos"
var websuphe int = 0
var sessiondweb string
var csrf string
var zaman int64
var targetsessiond string
var exists bool
var sorun bool
var sessiond string
var checkurl string
var mid string

func LoginApi(username string, password string) string {
	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}
	req.SetRequestURI("https://i.instagram.com/api/v1/accounts/login/")
	req.Header.Add("Accept-Language", "en-US")
	req.Header.SetMethod("POST")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Instagram 76.0.0.15.395 Android (22/5.1.1; 240dpi; 720x1280; merd; A5010; A5010; intel; tr_TR; 138226758)")
	req.SetBody([]byte(fmt.Sprintf("username=%s&password=%s&device_id=android-72ca06793a4875e9&login_attempt_count=0", username, password)))
	println(username)
	println(password)

	err := client.Do(req, resp)
	if err != nil {
		println(err)
	}

	fasthttp.ReleaseRequest(req)
	buffer := string(resp.Body())
	println(req)
	println(resp)
	println(string(buffer))
	midmerd := resp.Header
	rer := regexp.MustCompile(`mid=(.*); Domain=.instagram.com;`)
	matchh := rer.FindStringSubmatch(midmerd.String())

	fmt.Println(matchh[1])
	mid = matchh[1]
	println(mid)
	if strings.Contains(string(buffer), "challenge_required") {
		fmt.Println("Suphe Dusen hesap: " + username)

	} else if strings.Contains(string(buffer), "logged_in_user") {

		deneme := resp.Header
		re := regexp.MustCompile(`sessionid=(.*); Domain=.instagram.com;`)
		match := re.FindStringSubmatch(deneme.String())

		sessiond = match[1]

		fmt.Println("api giris basarili")
		return sessiond

	} else if strings.Contains(string(buffer), "two_factor_required\":true") {
		twofa := buffer
		var identifier string
		fac = 1
		var Anas Anas
		json.Unmarshal([]byte(twofa), &Anas)
		identifier = Anas.TwoFactorInfo.TwoFactorIdentifier
		println(identifier)
		identifierr = identifier

	} else if strings.Contains(string(buffer), "too many request") {

	} else {
		fmt.Println("SORUN VAR: " + username + string(buffer))
	}

	return sessiond
}

func FactorLogin(identifier, otpinput, username string) string {
	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}
	req.SetRequestURI("https://i.instagram.com/api/v1/accounts/two_factor_login/")
	req.Header.Add("Accept-Language", "en-US")
	req.Header.SetMethod("POST")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Instagram 98.0.0.15.119 Android (22/5.1.1; 240dpi; 720x1280; OnePlus; A5010; A5010; intel; tr_TR; 159526791)")

	req.SetBody([]byte(fmt.Sprintf("signed_body=.{\"verification_code\":\"%s\",\"two_factor_identifier\":\"%s\",\"username\":\"%s\",\"trust_this_device\":\"0\",\"device_id\":\"android-72ca06792a4875e9\"}", otpinput, identifier, username)))
	println(username)
	println(identifier)
	println(otpinput)
	err := client.Do(req, resp)
	if err != nil {
		println(err)
	}

	fasthttp.ReleaseRequest(req)
	buffer := string(resp.Body())
	println(string(buffer))
	if strings.Contains(string(buffer), "wrong") {
		fmt.Println("Suphe Dusen hesap: " + username)

	} else if strings.Contains(string(buffer), "logged_in_user") {

		deneme := resp.Header
		re := regexp.MustCompile(`sessionid=(.*); Domain=.instagram.com;`)
		match := re.FindStringSubmatch(deneme.String())

		sessiond = match[1]
		passtwo = 1
		fmt.Println("api giris basarili")
		return sessiond

	} else if strings.Contains(string(buffer), "This code is no longer valid") {
		fmt.Println("mnakoyim")
	} else if strings.Contains(string(buffer), "sms_code_validation_code_invalid") {
		passtwo = 0

	} else {
		fmt.Println("SORUN VAR: " + username + string(buffer))
	}

	return sessiond
}

func LoginWeb(username string, password string) string {

	csrf = randomdata.RandStringRunes(33)
	now := time.Now()
	zaman = now.Unix()
	fmt.Println(zaman)
	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}
	req.SetRequestURI("https://instagram.com/accounts/login/ajax/")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8")
	req.Header.SetMethod("POST")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edge/101.0.1210.32")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("x-csrftoken", fmt.Sprintf("%s", csrf))
	req.SetBody([]byte(fmt.Sprintf("username=%s&enc_password=#PWD_INSTAGRAM_BROWSER:0:%d:%s", username, zaman, password)))

	err := client.Do(req, resp)
	if err != nil {
		println(err)
	}
	fasthttp.ReleaseRequest(req)
	buffer := string(resp.Body())
	midmerd := resp.Header
	rer := regexp.MustCompile(`mid=(.*); Domain=.instagram.com;`)
	matchh := rer.FindStringSubmatch(midmerd.String())

	fmt.Println(matchh[1])
	mid = matchh[1]
	println(mid)
	if strings.Contains(string(buffer), "two_factor_required") {

		twofa := buffer
		var identifier string
		fac = 1
		var Anas Anas
		json.Unmarshal([]byte(twofa), &Anas)
		identifier = Anas.TwoFactorInfo.TwoFactorIdentifier
		println(identifier)
		identifierr = identifier

	} else if strings.Contains(string(buffer), "checkpoint_required") {

		twofa := buffer

		var weblos weblos
		json.Unmarshal([]byte(twofa), &weblos)
		checkurl = weblos.CheckpointURL
		println(checkurl)
		getsuphe(checkurl)

		postsuphe(checkurl)

	} else if strings.Contains(string(buffer), "\"oneTapPrompt\":true") {

		deneme := resp.Header
		re := regexp.MustCompile(`sessionid=(.*); Domain=.instagram.com;`)
		match := re.FindStringSubmatch(deneme.String())

		fmt.Println(match[1])
		sessiondweb = match[1]
		fmt.Println("web giris basarili")
		exists = true
		return sessiondweb

	} else {
		fmt.Println("SORUN VAR")
		fmt.Println(string(buffer))
		os.Exit(1)
	}
	return sessiondweb
}
func webFactorLogin(identifier, otpinput, username string) string {
	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}
	req.SetRequestURI("https://www.instagram.com/accounts/login/ajax/two_factor/")
	req.Header.Add("Accept-Language", "en-US")
	req.Header.SetMethod("POST")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edge/101.0.1210.32")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("x-csrftoken", fmt.Sprintf("%s", csrf))
	req.Header.Add("Cookie", fmt.Sprintf("mid=%s", mid))

	req.SetBody([]byte(fmt.Sprintf("identifier=%s&trust_signal=true&username=%s&verificationCode=%s&verification_method=3&queryParams={\"next\":\"/\"}", identifier, username, otpinput)))
	println(username)
	println(identifier)

	err := client.Do(req, resp)
	if err != nil {
		println(err)
	}

	fasthttp.ReleaseRequest(req)
	buffer := string(resp.Body())
	println(string(buffer))
	if strings.Contains(string(buffer), "wrong") {
		fmt.Println("Suphe Dusen hesap: " + username)

	} else if strings.Contains(string(buffer), "authenticated\":true") {

		deneme := resp.Header
		re := regexp.MustCompile(`sessionid=(.*); Domain=.instagram.com;`)
		match := re.FindStringSubmatch(deneme.String())
		passtwo = 1
		sessiondweb = match[1]
		println(sessiondweb)
		fmt.Println("web giris basarili")
		return sessiondweb

	} else if strings.Contains(string(buffer), "This code is no longer valid") {

	} else if strings.Contains(string(buffer), "sms_code_validation_code_invalid") {
		passtwo = 0
	} else {
		fmt.Println("SORUN VAR: " + username + string(buffer))
	}

	return sessiondweb
}

func getsuphe(checkurl string) {
	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}

	req.SetRequestURI(fmt.Sprintf("https://www.instagram.com%s", checkurl))
	req.Header.Add("Accept-Language", "en-US,en;q=0.8")
	req.Header.SetMethod("GET")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edge/101.0.1210.32")
	req.Header.Add("x-csrftoken", fmt.Sprintf("%s", csrf))
	req.Header.Add("Cookie", fmt.Sprintf("mid=%s", mid))

	err := client.DoRedirects(req, resp, 3)
	if err != nil {
		println(err)
	}
	fmt.Println("getlendi:", string(resp.Body()))

	websuphe = 1

	fasthttp.ReleaseRequest(req)

}
func postsuphe(checkurl string) {
	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}
	req.SetRequestURI(fmt.Sprintf("https://www.instagram.com%s", checkurl))
	req.Header.Add("Accept-Language", "en-US,en;q=0.8")
	req.Header.SetMethod("POST")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edge/101.0.1210.32")
	req.Header.Add("x-csrftoken", fmt.Sprintf("%s", csrf))
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("Cookie", fmt.Sprintf("mid=%s", mid))

	var choice string = "0"
	req.SetBody([]byte(fmt.Sprintf("choice=%s", choice)))

	err := client.Do(req, resp)
	if err != nil {
		println(err)
	}
	fmt.Println("son kod istendi:", string(resp.Body()))
	fasthttp.ReleaseRequest(req)
}
func postkodsuphe(checkurl, supheinput string) {
	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}
	req.SetRequestURI(fmt.Sprintf("https://instagram.com%s", checkurl))
	req.Header.Add("Accept-Language", "en-US,en;q=0.8")
	req.Header.SetMethod("POST")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edge/101.0.1210.32")
	req.Header.Add("x-csrftoken", fmt.Sprintf("%s", csrf))
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("Cookie", fmt.Sprintf("mid=%s", mid))

	req.SetBody([]byte(fmt.Sprintf("security_code=%s", supheinput)))

	err := client.DoRedirects(req, resp, 3)
	if err != nil {
		println(err)
	}
	fmt.Println(resp)
	bufferw := string(resp.Body())
	fmt.Println("kod instaya gonderildi:")
	if strings.Contains(string(bufferw), "status\":\"ok\"") {
		sessiondwebw := resp.Header
		re := regexp.MustCompile(`sessionid=(.*); Domain=.instagram.com;`)
		match := re.FindStringSubmatch(sessiondwebw.String())
		webgiris = 1
		fmt.Println("websessionid:", match[1])
		sessiondweb = match[1]
		fmt.Println("web giris basarili")

	} else {
		webgiris = 0
	}
	fasthttp.ReleaseRequest(req)

}

func gettwo(sessiondweb string) {
	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}

	req.SetRequestURI("https://www.instagram.com/accounts/two_factor_authentication/?__a=1&__d=dis")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8")
	req.Header.SetMethod("GET")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edge/101.0.1210.32")
	req.Header.Add("x-csrftoken", fmt.Sprintf("%s", csrf))
	req.Header.Add("Cookie", fmt.Sprintf("sessionid=%s", sessiondweb))

	err := client.DoRedirects(req, resp, 5)
	if err != nil {
		println(err)
	}
	fmt.Println("twofamethodlarigetlendi:", string(resp.Body()))
	fmt.Println(resp)
	fasthttp.ReleaseRequest(req)

	if strings.Contains(string(resp.Body()), "form_data") {
		buffer := string(resp.Body())

		var webtwo webtwo
		json.Unmarshal([]byte(buffer), &webtwo)
		duzotp := webtwo.FormData.TwoFactorEnabled
		totp := webtwo.FormData.TotpTwoFactorEnabled
		println(duzotp, totp)
		if totp == true {
			disabletwo(sessiondweb)
		}
	}

}

func disabletwo(sessiondweb string) {
TWODISABLE:
	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}
	req.SetRequestURI("https://www.instagram.com/accounts/two_factor_authentication/disable_totp/")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8")
	req.Header.SetMethod("POST")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edge/101.0.1210.32")
	req.Header.Add("x-csrftoken", fmt.Sprintf("%s", csrf))
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("Cookie", fmt.Sprintf("sessionid=%s", sessiondweb))

	err := client.Do(req, resp)
	if err != nil {
		println(err)
	}
	fmt.Println(resp)
	bufferw := string(resp.Body())

	fasthttp.ReleaseRequest(req)
	if strings.Contains(string(bufferw), "status\":\"ok\"") {
		fmt.Println("2faktor kaldirildi:")

		return

	} else {
		goto TWODISABLE
	}
}

func apigettwo(sessiond string) {
	client := &fasthttp.Client{}
	client.ReadBufferSize = 8192
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{
		Header:               fasthttp.ResponseHeader{},
		ImmediateHeaderFlush: false,
		SkipBody:             false,
	}

	req.SetRequestURI("https://i.instagram.com/api/v1/accounts/account_security_info/")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8")
	req.Header.SetMethod("POST")
	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("user-agent", "Instagram 76.0.0.15.395 Android (22/5.1.1; 240dpi; 720x1280; OnePlus; A5010; A5010; intel; tr_TR; 138226758)")
	req.Header.Add("Cookie", fmt.Sprintf("sessionid=%s", sessiond))

	err := client.Do(req, resp)
	if err != nil {
		println(err)
	}
	fmt.Println("twofamethodlarigetlendi:", string(resp.Body()))
	fmt.Println(resp)
	fasthttp.ReleaseRequest(req)

	if strings.Contains(string(resp.Body()), "form_data") {
		buffer := string(resp.Body())

		var webtwo webtwo
		json.Unmarshal([]byte(buffer), &webtwo)
		duzotp := webtwo.FormData.TwoFactorEnabled
		totp := webtwo.FormData.TotpTwoFactorEnabled
		println(duzotp, totp)
		fmt.Scanln()
		if totp == true {
			disabletwo(sessiondweb)
		}
	}

}
