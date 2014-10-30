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

//功能描述
//此接口用于查询下订单（含筛选）接口的订单情况——注：因 Internet 环境下，网络不是
//绝对可靠，用户系统下订单到顺丰后，不一定可以收到顺丰返回的数据，此接口用于在未收到返
//回数据时，查询订单当前的情况。
//订单结果查询
type OrderSearch struct {
	XMLName xml.Name `xml:"OrderSearch"`

	Orderid string `xml:"orderid,attr"` //客户订单号。

}

//返回值
type OrderSearchResponse struct {
	XMLName      xml.Name `xml:"OrderResponse"`
	Orderid      string   `xml:"orderid,attr"` //订单号
	Mailno       string   `xml:"mailno"`
	FilterResult string   `xml:"filter_result,attr"` //筛单结果：1- 人工确认，2- 可收派3-不可以收派
	Origincode   string   `xml:"origincode,attr"`    //原寄地代码
	Destcode     string   `xml:"destcode,attr"`      //目的地代码,如果可收派，此项不能为空
	Remark       string   `xml:"remark,attr"`        //1-收方超范围，2-派方超范围，3-其他原因
}

func (OrderSearch *OrderSearch) CallSF() ([]OrderSearchResponse, error) {
	req := NewRequest("OrderSearchService", OrderSearch)

	resp, err := req.send()
	return resp.([]OrderSearchResponse), err
}
