// @description 顺丰接口的sdk封装
// @link        https://github.com/chenghuama/sf
// @license     https://github.com/chenghuama/sf/blob/master/LICENSE
// @authors     chenghuama(chenghuama@vip.qq.com)
// @team yl10/忆零
package sf

import (
	"encoding/xml"
	//"github.com/astaxie/beego"
	//"strconv"
	//"github.com/astaxie/beego/httplib"
)

//订单筛选
type OrderFilter struct {
	XMLName           xml.Name             `xml:"OrderFilter"`
	FilterType        string               `xml:"filter_type,attr"`  //筛单类别：1-自动筛单（系统根据地址库进行判断，并返回结果），2-可人工筛单（系统首先根据地址库判断，如果无法自动判断是否收派，系统将生成需要人工判断的任务，后续由人工处理，处理结束后，顺丰可主动推送给客户系统
	Orderid           string               `xml:"orderid,attr"`      //客户订单号。如果filter_type=2，则必须提供
	DAddress          string               `xml:"d_address,attr"`    //到件方详细地址
	OrderFilterOption []*OrderFilterOption `xml:"OrderFilterOption"` //订单筛选选项
}

//订单筛选选项
type OrderFilterOption struct {
	JTel      string `xml:"j_tel,attr"`      //寄件方电话
	Country   string `xml:"country,attr"`    //寄件人所在国家
	Province  string `xml:"province,attr"`   //寄件人所在省份
	City      string `xml:"city,attr"`       //寄件方所属城市名称
	County    string `xml:"county,attr"`     //寄件人所在县/区
	DCountry  string `xml:"d_country,attr"`  //到件方国家
	DProvince string `xml:"d_province,attr"` //到件方所在省份
	DCity     string `xml:"d_city,attr"`     //到件方所属城市名称
	DCounty   string `xml:"d_county,attr"`   //到件方所在县/区
	JAddress  string `xml:"j_address,attr"`  //寄件方详细地址
	DTel      string `xml:"d_tel,attr"`      //到件方电话
	JCustid   string `xml:"j_custid,attr"`   //寄方客户编码
}

//返回值
type OrderFilterResponse struct {
	XMLName      xml.Name `xml:"OrderFilterResponse"`
	Orderid      string   `xml:"orderid,attr"`       //订单号
	FilterResult string   `xml:"filter_result,attr"` //筛单结果：1- 人工确认，2- 可收派3-不可以收派
	Origincode   string   `xml:"origincode,attr"`    //原寄地代码
	Destcode     string   `xml:"destcode,attr"`      //目的地代码,如果可收派，此项不能为空
	Remark       string   `xml:"remark,attr"`        //1-收方超范围，2-派方超范围，3-其他原因
}

func (OrderFilter *OrderFilter) CallSF() ([]OrderFilterResponse, error) {
	req := NewRequest("OrderFilterService", OrderFilter)

	resp, err := req.send()
	return resp.([]OrderFilterResponse), err
}
