// @description 顺丰接口的sdk封装
// @link        https://github.com/chenghuama/sf
// @license     https://github.com/chenghuama/sf/blob/master/LICENSE
// @authors     chenghuama(chenghuama@vip.qq.com)
// @team yl10/忆零
package sf

import ()

//退货订单
type OrderBack struct {
	Orderid       string       `xml:"orderid,attr"`        //订单号
	OrigOrderid   string       `xml:"orig_orderid,attr"`   //原订单号
	OrigMailno    string       `xml:"orig_mailno,attr"`    //原运单号
	ExpressType   string       `xml:"express_type,attr"`   //快件产品类别(可根据需要定制扩展)1 标准快递2 顺丰特惠
	JCompany      string       `xml:"j_company,attr"`      //退货方公司名称
	JContact      string       `xml:"j_contact,attr"`      //退货方联系人
	JTel          string       `xml:"j_tel,attr"`          //退货方联系电话
	JMobile       string       `xml:"j_mobile,attr"`       //手机
	JAddress      string       `xml:"j_address,attr"`      //退货方详细地址
	JPostalCode   string       `xml:"j_postal_code,attr"`  //退货方邮编
	DCompany      string       `xml:"d_company,attr"`      //商家（收货方）公司名称
	DContact      string       `xml:"d_contact,attr"`      //商家（收货方）联系人
	DTel          string       `xml:"d_tel,attr"`          //商家（收货方）联系电话
	DMobile       string       `xml:"d_mobile,attr"`       //手机
	DAddress      string       `xml:"d_address,attr"`      //商家（收货方）详细地址
	DPostalCode   string       `xml:"d_postal_code,attr"`  //商家（收货方）邮编
	ShopName      string       `xml:"shop_name,attr"`      //商家店铺名称
	PayMethod     string       `xml:"pay_method,attr"`     //付款方式：1:寄方付 2:收方付
	Custid        string       `xml:"custid,attr"`         //月结卡号
	JProvince     string       `xml:"j_province,attr"`     //退货方所在省份
	JCity         string       `xml:"j_city,attr"`         //退货方所属城市名称
	JCounty       string       `xml:"j_county,attr"`       //退货方所在县/区
	JShippercode  string       `xml:"j_shippercode,attr"`  //退货方城市区号
	DProvince     string       `xml:"d_province,attr"`     //商家（收货方）所在省份
	DCity         string       `xml:"d_city,attr"`         //商家（收货方）所属城市名称
	DCounty       string       `xml:"d_county,attr"`       //商家（收货方）所在县/区
	DDeliverycode string       `xml:"d_deliverycode,attr"` //商家（收货方）城市区号
	Mailno        string       `xml:"mailno,attr"`         //运单号，如果客户不指定运单号，则由顺丰平台系统生成一个运单号
	Remark        string       `xml:"remark,attr"`         //运单备注
	Template      string       `xml:"template,attr"`       //模板选择
	Sendstarttime string       `xml:"sendstarttime,attr"`  //要求上门取件开始时间
	BackCargo     []*BackCargo `xml:"Cargo"`               //货物信息
}

//货物信息
type BackCargo struct {
	Name             string `xml:"name,attr"`              //货物名称
	Serl             string `xml:"serl,attr"`              //商品编号
	Count            string `xml:"count,attr"`             //货物数量
	Unit             string `xml:"unit,attr"`              //货物单位如：个,台,本
	Weight           string `xml:"weight,attr"`            //货物单位重量
	Amount           string `xml:"amount,attr"`            //货物单价
	VisualInspection string `xml:"visual_inspection,attr"` //外观，如果商品需要收派员验货则要填写这一项。
	CheckRemark      string `xml:"check_remark,attr"`      //收货备注，如果商品需要收派员验货则要填写这一项。
}

//退单增值
type BackAddedService struct {
	Name   string `xml:"name,attr"`  //货物名称
	value  string `xml:"value,attr"` //增值服务值 1
	value1 string `xml:"name,attr"`  //增值服务值 2
}

//返回值
type OrderReverseResponse struct {
	Orderid      string `xml:"orderid,attr"`       //订单号
	Mailno       string `xml:"mailno,attr"`        //运单号，可多个单号，如子母件，以逗号分隔
	FilterResult string `xml:"filter_result,attr"` //筛单结果：1-人工确认，2-可收派 3-不可以收派
	Remark       string `xml:"remark,attr"`        //1-退货方超范围，2-商家超范围，3-其他原因
}

func (OrderBack *OrderBack) CallSF() ([]OrderReverseResponse, error) {
	req := NewRequest("OrderReverseResponse", &OrderBack)
	resp, err := req.send()
	return resp.([]OrderReverseResponse), err
}
