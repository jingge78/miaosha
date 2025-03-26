package pkg

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

type Payment interface {
	Pay(Subject, OutTradeNo, TotalAmount string) string
}

type Alipay struct {
	AppId      string
	PrivateKey string
	NotifyURL  string
	ReturnURL  string
}

func NewPay() *Alipay {
	return &Alipay{
		AppId:      "9021000142690952",
		PrivateKey: "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCHOdxmStOy2g30p4AvhSZ+Ox1vI3ESDpLPtAcp+xDWdOGfa3yh1gFukIr3wrVzxacrTUR5HYyr0g3GiVUlff9dCnh24JpfqMQ8GEz8niu72rIIIBmFN8m4Ejd2qJeHeWou5Uf3TVcH6ygtzVfeWPybssa0lJfw4TPWb+d2rPP5xtt2rFktLXdsE7vI4OFOEM3tRSfuLfG242OXrAK+FxCEPQ9dT9FJklYvCLJ8RKKqLstMuBWOUzzc7otplmVWFokElH71LVEcz26mXZX1ExEvPMw47BFYZ6Y38VLJeYMslS1bqaTAXgUBD8PoKsDCxycYM11eMHidaqUxA48suYNvAgMBAAECggEALeMNjkyv/84M1EuOrRFy2Xz35QHS2bzGRuMhVzaSJSPueCmCVmyHedxku+R/rHSS4JfMt4i2dovGDuwFT76szAbEkBpxaCqdxIK+hS6rSojQxv8VieY/dk4AMizNlrQ1uwok3J+K++3paXl36sSpm7ATy61szdmtvIOmuNfBxq0c2yu+E+2nzJPsIPtX18v2nG/X4Kp+4ZCBorcPssm54RicfKmfmGA4zGxmvhUczuhHiEGIvJFbwMjP6rOwjGArfWYrlOZ1mSTXi9/W7Ji1LDXWrXEDgwmQ/RUBvBgZP1XUr2R6/PJaiFWGwCYcYkip+WJ1n4fu06+0M9EOgJ88EQKBgQDDHPbNQcJhgzI2Dw7taZJs1/K5+NrE6CvFgYzkuHp/FKY3dW/zm5rlJSUzKYCwHC/meoDaZGJ1aXN+9o73VTHsXlW+SsgWR0mZ9GlnIB7QHSuL17lJbHLnQ1tU4+PqNAIstT5yocToy22p4rpmnyosF9St0Ow+c8+z0z+oTp3B6QKBgQCxbK82ZPUwnP12tGeFarqxnJ7lQXHkdU3w//RKWox0w2W9rgpW9nSUOKUyPO4dWn46FEo2nZ3ZEsGSAlZ8iqEKxYOFeUDQVgS5O39+v/YzxqUpXYehn3xKkZDQVXnorotKK5Rj74Ba+QaGfAqe22np0IMqPU10hInZjg1dAworlwKBgAhbWjrKYT/59ZGZLYN/rRTaXvwWK5CZfR51gQpe2GhPAxuG/SeK96Ru5dv+IBPq8SZHAvPXrtvmi1rZxp/TV1MPa06+Nzm1DfL5I/aVypwRU8cmkzoQ2g8LtIK7TAzA84LktGsGgL+TzvuiyWcR1CWVU7eqJiQ6o5/JIYXc8CbZAoGAETv7cQ8xef1l6YfwnlcVt3b9QEuxIn36ijRyqF5PUnBAi8JCItxhypwN/+lHP/awWDfsVY3N7W4S+3naqNJWflNdSTPUBei1IMEUy10eLz1WgcQiDqMNUbj+Fh6XbvC1ewjsqyBymWOjLKET7wZlLV8hvpKh2XWeZlGUHrrS3BUCgYAOBx/sUXCCT0FAPy/bU//eEgnuZ8Q1ABTX77fafokqHfCBEn6tNHtGFRPowp8B0hEw7rq1dTmLkmwML2TrIzAcit0qqzGZg5/DVW83faL1G69Swxl0Xc8Vc8328+61W4mKWuiEOC9EBr3sHr0A2VEpGmdyQAQCHNqbQB5o+a39eQ==",
		NotifyURL:  "http://3ee55b16.r9.cpolar.top",
		ReturnURL:  "https://www.baidu.com",
	}
}

func (a *Alipay) Pay(Subject, OutTradeNo, TotalAmount string) string {
	var privateKey = a.PrivateKey // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New(a.AppId, privateKey, false)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var p = alipay.TradeWapPay{}
	p.NotifyURL = a.NotifyURL
	p.ReturnURL = a.ReturnURL
	p.Subject = Subject
	p.OutTradeNo = OutTradeNo
	p.TotalAmount = TotalAmount
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	return payURL
}
