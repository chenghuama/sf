// @description 顺丰接口的sdk封装
// @link        https://github.com/chenghuama/sf
// @license     https://github.com/chenghuama/sf/blob/master/LICENSE
// @authors     chenghuama(chenghuama@vip.qq.com)
// @team yl10/忆零
package sf

import (

//"github.com/astaxie/beego"
//"strconv"
//"github.com/astaxie/beego/httplib"
)

//换货定单
type OrderChange struct {
	Orderid     string  `xml:"orderid,attr"`      //订单号
	OrigOrderid string  `xml:"orig_orderid,attr"` //原订单号
	OrigMailno  string  `xml:"orig_mailno,attr"`  //原运单号
	Template    string  `xml:"template,attr"`     //模板选择
	Order1      *Order1 `xml:"Order1"`            //订单信息
	Order2      *Order2 `xml:"Order2"`            //订单信息
}

//换货订单1
type Order1 struct {
	ExpressType      string              `xml:"express_type,attr"`   //快件产品类别1 标准快递2 顺丰特惠
	JCompany         string              `xml:"j_company,attr"`      //寄新货方公司名称
	JContact         string              `xml:"j_contact,attr"`      //寄新货方联系人
	JTel             string              `xml:"j_tel,attr"`          //寄新货方联系电话
	JMobile          string              `xml:"j_mobile,attr"`       //手机
	JAddress         string              `xml:"j_address,attr"`      //寄新货方详细地址
	JPostalCode      string              `xml:"j_postal_code,attr"`  //寄新货方邮编
	DCompany         string              `xml:"d_company,attr"`      //换货方公司名称
	DContact         string              `xml:"d_contact,attr"`      //换货方联系人
	DTel             string              `xml:"d_tel,attr"`          //换货方联系电话
	DMobile          string              `xml:"d_mobile,attr"`       //手机
	DAddress         string              `xml:"d_address,attr"`      //换货方详细地址
	DPostalCode      string              `xml:"d_postal_code,attr"`  //换货方邮编
	PayMethod        string              `xml:"pay_method,attr"`     //付款方式：1:寄方付 2:收方付默认为 1
	Custid           string              `xml:"custid,attr"`         //寄新货方月结卡号如果此字段不为空，则表示寄新货的运单为月结。
	JProvince        string              `xml:"j_province,attr"`     //寄新货方所在省份
	JCity            string              `xml:"j_city,attr"`         //寄新货方所属城市名称
	JCounty          string              `xml:"j_county,attr"`       //寄新货方所在县/区
	JShippercode     string              `xml:"j_shippercode,attr"`  //寄新货方城市区号
	DProvince        string              `xml:"d_province,attr"`     //换货方所在省份
	DCity            string              `xml:"d_city,attr"`         //换货方所属城市
	DCounty          string              `xml:"d_county,attr"`       //换货方所在县/区
	DDeliverycode    string              `xml:"d_deliverycode,attr"` //换货方城市区号
	Mailno           string              `xml:"mailno,attr"`         //运单号
	Remark           string              `xml:"remark,attr"`         //运单备注
	Sendstarttime    string              `xml:"sendstarttime,attr"`  //要求上门取件开始时间
	Order1Cargo      []*Order1Cargo      `xml:"Cargo"`               //货物信息
	BackAddedService []*BackAddedService `xml:"AddedService"`        //增值服务
}

//货物信息
type Order1Cargo struct {
	Name   string `xml:"name,attr"`   //货物名称
	Count  string `xml:"count,attr"`  //货物数量
	Unit   string `xml:"unit,attr"`   //货物单位如：个,台,本
	Weight string `xml:"weight,attr"` //货物单位重量
	Amount string `xml:"amount,attr"` //货物单价
	Remark string `xml:"remark,attr"` //收货备注，如果商品需要收派员验货则要填写这一项。
}

//换货订单2
type Order2 struct {
	ExpressType      string              `xml:"express_type,attr"`   //快件产品类别,1 标准快递2 顺丰特惠
	DCompany         string              `xml:"d_company,attr"`      //收旧货方公司名称
	DContact         string              `xml:"d_contact,attr"`      //收旧货方联系人
	DTel             string              `xml:"d_tel,attr"`          //收旧货方联系电话
	DMobile          string              `xml:"d_mobile,attr"`       //手机
	DAddress         string              `xml:"d_address,attr"`      //收旧货方详细地址
	DPostalCode      string              `xml:"d_postal_code,attr"`  //收旧货方邮编
	PayMethod        string              `xml:"pay_method,attr"`     //付款方式：1:寄方付 2:收方付
	Custid           string              `xml:"custid,attr"`         //月结卡号
	DProvince        string              `xml:"d_province,attr"`     //收旧货方所在省份
	DCity            string              `xml:"d_city,attr"`         //收旧货方所属城市名称
	DCounty          string              `xml:"d_county,attr"`       //收旧货方所在县/区
	DDeliverycode    string              `xml:"d_deliverycode,attr"` //收旧货方城市区号
	Mailno           string              `xml:"mailno,attr"`         //运单号，如果客户不指定运单号，则由顺丰平台系统生成一个运单号
	Remark           string              `xml:"remark,attr"`         //运单备注
	Order2Cargo      []*Order2Cargo      `xml:"Cargo"`               //货物信息
	BackAddedService []*BackAddedService `xml:"AddedService"`        //增值服务
}

//货物信息
type Order2Cargo struct {
	Name   string `xml:"name,attr"`   //货物名称
	Serl   string `xml:"serl,attr"`   //商品编号
	Count  string `xml:"count,attr"`  //货物数量
	Unit   string `xml:"unit,attr"`   //货物单位如：个,台,本
	Weight string `xml:"weight,attr"` //货物单位重量
	Amount string `xml:"amount,attr"` //货物单价
	Remark string `xml:"remark,attr"` //收货备注，如果商品需要收派员验货则要填写这一项。
}

//返回值
type OrderExchangeResponse struct {
	Orderid      string `xml:"orderid,attr"`       //订单号
	Mailno1      string `xml:"mailno1,attr"`       //发新货的运单号
	Mailno2      string `xml:"mailno2,attr"`       //发旧货的运单号
	Origincode   string `xml:"origincode,attr"`    //发新货的原寄地代码
	Destcode     string `xml:"destcode,attr"`      //新货的目的地代码
	Destcode2    string `xml:"destcode2,attr"`     //旧货的目的地代码
	FilterResult string `xml:"filter_result,attr"` //筛单结果：1-人工确认，2-可收派 3-不可以收派
	Remark       string `xml:"remark,attr"`        //	1-换货方超范围，2-商家超范围，3-其他原因
}

func (OrderChange *OrderChange) CallSF() ([]OrderExchangeResponse, error) {
	req := NewRequest("OrderExchangeService", OrderChange)
	resp, err := req.send()
	return resp.([]OrderExchangeResponse), err
}
