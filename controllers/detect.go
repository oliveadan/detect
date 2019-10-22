package controllers

import (
	"crypto/tls"
	"detect/utils"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type DetectController struct {
	beego.Controller
}

type dict struct {
	SiteName   string
	IP         string
	Loc        string
	StatusCode string
	Status     string
	Location   string
	Warn       string
	Err        string
	Info       string
}

func (d *DetectController) DoNow() {
	var code int
	var msg string
	var dicts []dict
	data := make(map[string]interface{})
	defer utils.Retjson(d.Ctx, &msg, &code, &data)
	pertimes, _ := d.GetInt("pertimes", 0)
	iptype, _ := d.GetInt("iptype", 0)
	proxyapi := d.GetString("proxyapi")
	//获取检测网址
	s := beego.AppConfig.String("sitespath")
	ss, e := ioutil.ReadFile(s)
	if e != nil {
		beego.Error("read sitespath error", e)
		msg = "读取检测网址失败"
		return
	}
	//获取关键字
	k := beego.AppConfig.String("keywordspath")
	ks, e1 := ioutil.ReadFile(k)
	if e1 != nil {
		beego.Error("read keywords error", e1)
		msg = "读取检测网址失败"
		return
	}
	sites := strings.Fields(string(ss))
	keywords := strings.Fields(string(ks))
	//本地IP检测
	if iptype == 1 {
		var client = &http.Client{
			Timeout: time.Second * 10,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				if len(via) >= 1 {
					return errors.New("stopped after 1 redirects")
				}
				return nil
			},
		}

		for _, v := range sites {
			//需要检测网站的名字,这么错是防止下面重命名
			vv := v
			//开始测试
			for i := 1; i <= pertimes; i++ {
				var dict dict
				dict.SiteName = vv
				if !strings.HasPrefix(v, "http://") && !strings.HasPrefix(v, "https://") {
					v = "http://" + v
				}
				resp, clineterr := client.Get(strings.TrimSpace(v))
				//获取公共IP
				dict.IP = utils.Get_external()
				//获取IP地理位置
				ip, err := utils.GetLocationIpsb(strings.TrimSpace(utils.Get_external()))
				if err != nil {
					beego.Error("get localIP error", err)
					dict.Loc = "获取地理位置失败"
				} else {
					dict.Loc = ip
				}
				if clineterr != nil {
					serror := fmt.Sprintf("%s", clineterr)
					contains := strings.Contains(serror, "stopped after 1 redirects")
					if !contains {
						dict.Location = v
						dict.Status = "0"
						dict.Err = serror
						dicts = append(dicts, dict)
						continue
					}
				}
				//判断状态 status: 0：异常；1:正常；2：跳转
				stautscode := strconv.Itoa(resp.StatusCode)
				dict.StatusCode = stautscode
				if resp.StatusCode == 200 {
					dict.Status = "1"
					dict.Location = v
				} else if resp.StatusCode == 301 {
					url, _ := resp.Location()
					locationhttp := strings.TrimLeft(url.String(), "https://")
					locationhttp1 := strings.TrimRight(locationhttp, "/")
					http := strings.TrimLeft(v, "http://")
					if http == locationhttp1 {
						dict.Status = "1"
						dict.Location = url.String()
					} else {
						dict.Status = "2"
						dict.Location = url.String()
					}
				} else if resp.StatusCode == 302 || resp.StatusCode == 303 || resp.StatusCode == 307 {
					url, _ := resp.Location()
					dict.Status = "2"
					dict.Location = url.String()
				}
				//关键字检测
				body, _ := ioutil.ReadAll(resp.Body)
				for _, v := range keywords {
					contains := strings.Contains(string(body), strings.TrimSpace(v))
					if !contains {
						dict.Warn = "缺失关键字" + v
					}
				}
				resp.Body.Close()
				dicts = append(dicts, dict)
			}
		}
		//代理IP群测试
	} else {
		// 获取代理IP群
		dailijingling, e := utils.GetDailijingling(strings.TrimSpace(proxyapi))
		if e != nil {
			beego.Error("get dailiqun error", e)
			msg = "获取代理群失败"
			return
		}
		//测试用
		//var dailijingling = []string{"39.108.170.173:80", "139.219.8.96:8080","49.51.75.109:80","221.7.255.168:8080","149.129.75.181:8080","117.90.137.50:9000","111.29.3.190:80","111.29.3.225:8080","123.163.97.127:9999","111.29.3.188:80","114.244.214.46:8060","111.29.3.221:8080","115.218.208.165:9000","183.166.71.174:9999","222.139.6.236:8060","123.163.122.251:9999","222.139.6.236:8060"}
		for _, v := range sites {
			//需要检测网站的名字,这么做是防止下面重命名
			vv := v
			for _, j := range dailijingling {
				proxy, _ := url.Parse("http://" + j)
				tr := &http.Transport{
					Proxy:           http.ProxyURL(proxy),
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				}
				var client = &http.Client{
					Transport: tr,
					Timeout:   time.Second * 10,
					CheckRedirect: func(req *http.Request, via []*http.Request) error {
						if len(via) >= 1 {
							return errors.New("stopped after 1 redirects")
						}
						return nil
					},
				}
				//开始测试
				for i := 1; i <= pertimes; i++ {
					var dict dict
					dict.SiteName = vv
					if !strings.HasPrefix(v, "http://") && !strings.HasPrefix(v, "https://") {
						v = "http://" + v
					}
					resp, clineterr := client.Get(strings.TrimSpace(v))
					//获取公共IP
					split := strings.Split(j, ":")
					dict.IP = split[0]
					//获取IP地理位置
					ip, err := utils.GetLocationIpsb(strings.TrimSpace(split[0]))
					if err != nil {
						beego.Error("get localIP error", err)
						dict.Loc = "获取地理位置失败"
					} else {
						dict.Loc = ip
					}
					if clineterr != nil {
						serror := fmt.Sprintf("%s", clineterr)
						contains := strings.Contains(serror, "stopped after 1 redirects")
						if !contains {
							dict.Location = v
							dict.Status = "0"
							dict.Err = serror
							dicts = append(dicts, dict)
							continue
						}
					}
					//判断状态 status: 0：异常；1:正常；2：跳转
					stautscode := strconv.Itoa(resp.StatusCode)
					dict.StatusCode = stautscode
					if resp.StatusCode == 200 {
						dict.Status = "1"
						dict.Location = v
					} else if resp.StatusCode == 301 {
						url, _ := resp.Location()
						locationhttp := strings.TrimLeft(url.String(), "https://")
						locationhttp1 := strings.TrimRight(locationhttp, "/")
						http := strings.TrimLeft(v, "http://")
						if http == locationhttp1 {
							dict.Status = "1"
							dict.Location = url.String()
						} else {
							dict.Status = "2"
							dict.Location = url.String()
						}
					} else if resp.StatusCode == 302 || resp.StatusCode == 303 || resp.StatusCode == 307 {
						url, _ := resp.Location()
						dict.Status = "2"
						dict.Location = url.String()
					} else {
						dict.Location = v
						dict.Status = "0"
						dict.Err = resp.Status
					}
					if dict.Err == "" {
						//关键字检测
						body, _ := ioutil.ReadAll(resp.Body)
						for _, v := range keywords {
							contains := strings.Contains(string(body), strings.TrimSpace(v))
							if !contains {
								dict.Warn = "缺失关键字" + v
							}
						}
						resp.Body.Close()
					}
					dicts = append(dicts, dict)
				}
			}
		}

	}
	code = 1
	msg = "检测成功"
	data["dict"] = dicts
	/*for _, v := range dicts {
		fmt.Println()
		fmt.Printf("%+v", v)
	}*/
}
