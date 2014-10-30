// @description 顺丰接口的sdk封装
// @link        https://github.com/chenghuama/sf
// @license     https://github.com/chenghuama/sf/blob/master/LICENSE
// @authors     chenghuama(chenghuama@vip.qq.com)
// @team yl10/忆零

package sf

import (
	//	"encoding/xml"
	//	"github.com/astaxie/beego"
	//"strconv"
	//"github.com/astaxie/beego/httplib"
	"encoding/xml"
)

//订单
type Order struct {
	XMLName     xml.Name `xml:"Order"`
	Orderid     string   `xml:"orderid,attr"`
	ExpressType string   `xml:"express_type,attr"` //快件产品类别
	JCompany    string   `xml:"j_company,attr"`    //寄件方公司名称
	JContact    string   `xml:"j_contact,attr"`    //寄件方联系人
	JTel        string   `xml:"j_tel,attr"`        //寄件方联系电话
	JMobile     string   `xml:"j_mobile,attr"`     //寄件方手机
	JProvince   string   `xml:"j_province,attr"`   //寄件方所在省份字段
	JCity       string   `xml:"j_city,attr"`       //寄件方所属城市名称
	//JCountry             string         `xml:"j_country,attr"`               //寄方国家
	//	JCounty              string         `xml:"j_county,attr"`                //寄件人所在县/区
	//	JShippercode         string         `xml:"j_shippercode,attr"`           //寄件方代码
	//	JPostCode            string         `xml:"j_post_code,attr"`             //寄方邮编

	JAddress  string `xml:"j_address,attr"`  //寄件方详细地址
	DCompany  string `xml:"d_company,attr"`  //到件方公司名称
	DContact  string `xml:"d_contact,attr"`  //到件方联系人
	DTel      string `xml:"d_tel,attr"`      //到件方联系电话
	DMobile   string `xml:"d_mobile,attr"`   //到件方手机
	DAddress  string `xml:"d_address,attr"`  //到件方详细地址
	DProvince string `xml:"d_province,attr"` //到件方所在省份
	DCity     string `xml:"d_city,attr"`     //到件方所属城市
	//DeclaredValue         string `xml:"declared_value,attr"`          //托寄物声明
	//DeclaredValueCurrency string `xml:"declared_value_currency,attr"` //托寄物声明价值

	ParcelQuantity string `xml:"parcel_quantity,attr"` //包裹数
	PayMethod      string `xml:"pay_method,attr"`      //付款方式
	Custid         string `xml:"custid,attr"`          //月结卡号
	//Template             string         `xml:"template,attr"`                //模板选择
	//DCountry string `xml:"d_country,attr"` //到方国家
	DCounty string `xml:"d_county,attr"` //到件人所在县/区
	//DDeliverycode    string `xml:"d_deliverycode,attr"`     //到件方代码
	//CargoTotalWeight string `xml:"cargo_total_weight,attr"` //订单货物总重量
	Sendstarttime string `xml:"sendstarttime,attr"` //要求上门取件开始时间
	//Mailno           string `xml:"mailno,attr"`             //运单号
	//ReturnTracking       string         `xml:"return_tracking,attr"`         //签回单单号
	Remark string `xml:"remark,attr"` //备注
	//NeedReturnTrackingNo string         `xml:"need_return_tracking_no,attr"` //是否需要签回单号，1：需要
	//IsDocall             string         `xml:"is_docall,attr"`               //是否下 call 1-下 call
	//IsGenBillNo          string         `xml:"is_gen_bill_no,attr"`          //是否申请运单号   1- 申请运单号
	//IsGenEletricPic      string         `xml:"is_gen_eletric_pic,attr"`      //是否生成电子运单图片1- 生成
	//WaybillSize          string         `xml:"waybill_size,attr"`            //图片格式：如果需要顺丰系统推送图片，则要传这个值。1：A4 2：A5
	//CargoLength          string         `xml:"cargo_length,attr"`            //长
	//CargoWidth           string         `xml:"cargo_width,attr"`             //宽
	//CargoHeight          string         `xml:"cargo_height,attr"`            //高
	Cargo        []*Cargo        `xml:"Cargo"`        //货物信息
	AddedService []*AddedService `xml:"AddedService"` //增值服务
}

//货物信息
type Cargo struct {
	Name  string `xml:"name,attr"`  //货物名称
	Count string `xml:"count,attr"` //货物数量
	Unit  string `xml:"unit,attr"`  //货物单位
	//Weight     string `xml:"weight,attr"`      //货物单位重量
	//Currency   string `xml:"currency,attr"`    //货物单价的币别,与 Order 的 declared_value_currency属性相同
	//SourceArea string `xml:"source_area,attr"` //原产地国别
}

//增值服务
type AddedService struct {
	Name   string `xml:"name,attr"`   //增值服务名
	Value  string `xml:"value,attr"`  //增值服务值 1
	Value1 string `xml:"value1,attr"` //增值服务值 2
	Value2 string `xml:"value2,attr"` //增值服务值 3
	Value3 string `xml:"value3,attr"` //增值服务值 4
	Value4 string `xml:"value4,attr"` //增值服务值 5
}

//返回值
type OrderResponse struct {
	XMLName          xml.Name `xml:"OrderResponse"`
	Orderid          string   `xml:"orderid,attr"`            //订单号
	Mailno           string   `xml:"mailno,attr"`             //运单号，可多个单号，如子母件，以逗号分隔
	ReturnTrackingNo string   `xml:"return_tracking_no,attr"` //签单返还运单号
	AgentMailno      string   `xml:"agent_mailno,attr"`       //代理运单号，可多个单号，如子母件，以逗号分隔
	Origincode       string   `xml:"origincode,attr"`         //原寄地代码
	Destcode         string   `xml:"destcode,attr"`           //目的地代码
	FilterResult     string   `xml:"filter_result,attr"`      //单结果：1-人工确认，2-可收派 3-不可以收派
	Remark           string   `xml:"remark,attr"`             //1-收方超范围，2-派方超范围，3-其他原因
}

func (order *Order) CallSF() ([]OrderResponse, error) {
	req := NewRequest("OrderService", order)
	resp, err := req.send()
	return resp.([]OrderResponse), err

}
