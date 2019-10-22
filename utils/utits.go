package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/context"
	"io/ioutil"
	"net/http"
)

func Retjson(ctx *context.Context, msg *string, code *int, data ...interface{}) {
	ret := make(map[string]interface{})
	ret["code"] = code
	ret["msg"] = msg
	if len(data) > 0 {
		d := data[0]
		switch d.(type) {
		case string:
			ret["url"] = d
		case *string:
			ret["url"] = d
		default:
			ret["data"] = d
		}
	}
	ctx.Output.JSON(ret, false, false)
}

//获取公网IP
func Get_external() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	//buf := new(bytes.Buffer)
	//buf.ReadFrom(resp.Body)
	//s := buf.String()
	return string(content)
}

//获取IP地理位置

type IPInfo struct {
	Code int `json:"code"`
	Data IP  `json:"data`
}

type IP struct {
	Country   string `json:"country"`
	CountryId string `json:"country_id"`
	Area      string `json:"area"`
	AreaId    string `json:"area_id"`
	Region    string `json:"region"`
	RegionId  string `json:"region_id"`
	City      string `json:"city"`
	CityId    string `json:"city_id"`
	Isp       string `json:"isp"`
}

func LocalIp(ip string) (string, error) {
	url := "http://ip.taobao.com/service/getIpInfo.php?ip="
	url += ip

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 502 {
		return "", errors.New("服务器连接失败")
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	var result IPInfo
	if err := json.Unmarshal(out, &result); err != nil {
		return "", err
	}
	localtion := result.Data.Country + "" + result.Data.City + "" + "" + result.Data.Isp
	return localtion, nil
}

//根据http://https://api.ip.sb/ 获取地理位置

func GetLocationIpsb(ip string) (string, error) {
	type iplocation struct {
		Ip            string  `json:"ip"`
		CountryCode   string  `json:"country_code"`
		Country       string  `json:"country"`
		RegionCode    string  `json:"region_code"`
		Region        string  `json:"region"`
		City          string  `json:"city"`
		PostalCode    string  `json:"postal_code"`
		ContinentCode string  `json:"continent_code"`
		Latitude      float64 `json:"latitude"`
		Longitude     float64 `json:"longitude"`
		Organization  string  `json:"organization"`
		Timezone      string  `json:"timezone"`
	}
	var client = &http.Client{}
	url := "https://api.ip.sb/geoip/" + ip
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result iplocation
	if err := json.Unmarshal(out, &result); err != nil {
		return "", err
	}
	return "国家:" + result.Country + " 城市:" + result.City + " 运营商:" + result.Organization, nil
}

//获取代理群
func GetDailijingling(url string) ([]string, error) {
	type DaiIP struct {
		IP   string `json:"IP"`
		Port int64  `json:"Port"`
	}

	type DaiLiIpInfo struct {
		Success string  `json:"success"`
		Code    int     `json:"code"`
		Msg     string  `json:"msg"`
		Data    []DaiIP `json:"data`
	}
	var client = &http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result DaiLiIpInfo
	//反序列话测试用
	//str := `{"code":0,"success":"true","msg":"","data":[{"IP":"111.72.110.223","Port":21108},{"IP":"122.194.249.57","Port":41787},{"IP":"112.111.97.163","Port":22356},{"IP":"117.90.6.88","Port":32981},{"IP":"163.204.217.5","Port":24473},{"IP":"222.242.184.206","Port":22758},{"IP":"122.194.249.106","Port":13346},{"IP":"116.249.181.82","Port":22245},{"IP":"124.237.65.116","Port":33251},{"IP":"1.180.165.196","Port":37049}]}`
	//bytes := []byte(str)
	if err := json.Unmarshal(out, &result); err != nil {
		return nil, err
	}
	if result.Success == "false" {
		return nil, errors.New(result.Msg)
	}
	var ips []string
	for _, v := range result.Data {
		port := fmt.Sprintf("%d", v.Port)
		ips = append(ips, v.IP+":"+port)
	}
	return ips, nil
}
