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

//订单
type OrderConfirm struct {
	XMLName            xml.Name            `xml:"OrderConfirm"`
	Orderid            string              `xml:"orderid,attr"`       //订单号
	Mailno             string              `xml:"mailno,attr"`        //运单号(如果 dealtype=2，可选)
	Dealtype           string              `xml:"dealtype,attr"`      //订单操作标识:1-订单确认 2-消单
	OrderConfirmOption *OrderConfirmOption `xml:"OrderConfirmOption"` //快件信息
}

//快件信息
type OrderConfirmOption struct {
	Weight string `xml:"weight,attr"` //订单重量，单位 KG
	//Volume         string `xml:"volume,attr"`          //托寄物的长,宽,高，以半角逗号分隔，单位 CM，精确到小数点后一位
	//ReturnTracking string `xml:"return_tracking,attr"` //签回单单号
	//ExpressType    string `xml:"express_type,attr"`    //快件产品类别(可根据需要定制扩展)1 标准快递2 顺丰特惠3 电商特惠 如果此字段为空，则以下单时的为准
	//ChildrenNos    string `xml:"children_nos,attr"`    //子单号（以半角逗号分隔）,如果此字段为空，则以下订单时为准,如果此字段不为空，则忽略下订单时的子单号，以此字段的单号为准。
}

//返回值
type OrderConfirmResponse struct {
	XMLName   xml.Name `xml:"OrderConfirmResponse"`
	Orderid   string   `xml:"orderid,attr"`    //订单号
	Mailno    string   `xml:"mailno,attr"`     //运单号，可多个单号，如子母件，以逗号分隔
	ResStatus string   `xml:"res_status,attr"` //备注1 订单号与运单不匹配2 成功
}

func (OrderConfirm *OrderConfirm) CallSF() ([]OrderConfirmResponse, error) {
	req := NewRequest("OrderConfirmService", OrderConfirm)
	resp, err := req.send()
	return resp.([]OrderConfirmResponse), err
}
