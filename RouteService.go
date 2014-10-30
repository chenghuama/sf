// @description 顺丰接口的sdk封装
// @link        https://github.com/chenghuama/sf
// @license     https://github.com/chenghuama/sf/blob/master/LICENSE
// @authors     chenghuama(chenghuama@vip.qq.com)
// @team yl10/忆零

package sf

import (
	"encoding/xml"
	//	"github.com/astaxie/beego"
	//"strconv"
	//"github.com/astaxie/beego/httplib"
)

const (
	//路由查询方式
	ConstTracking_type_Orderno = 1 //按照订单号
	ConstTracking_type_mailno  = 2 //按照运单号

)

//路由查询
type RouteRequest struct {
	XMLName        xml.Name `xml:"RouteRequest"`
	TrackingType   int64    `xml:"tracking_type,attr"`   //查询类别, tracking_type 字段说明：1-根据运单号查询，order节点中 track_number 将被当作运单号处理，2-根据订单号查询，order 节点中 track_number 将被当作订单号处理-自动筛单（系统根据地址库进行判断，并返回结果），2-可人工筛单（系统首先根据地址库判断，如果无法自动判断是否收派，系统将生成需要人工判断的任务，后续由人工处理，处理结束后，顺丰可主动推送给客户系统
	TrackingNumber string   `xml:"tracking_number,attr"` //查询号, 如果tracking_type=1，则此值为运单号。如果 tracking_type=2，则此值为订单号如果有多个单号，以逗号分隔
	MethodType     string   `xml:"method_type,attr"`     //查询方法选择1-标准查询2-定制查询
}

//返回值
type RouteResponse struct {
	XMLName xml.Name `xml:"RouteResponse"`
	Orderid string   `xml:"orderid,attr"` //订单号
	Mailno  string   `xml:"mailno,attr"`  //运单号
	Route   []*Route `xml:"Route"`        //路由
}

//返回值
type Route struct {
	AcceptTime    string `xml:"accept_time,attr"`    //路由发生的时间
	AcceptAddress string `xml:"accept_address,attr"` //路由发生的地点
	Remark        string `xml:"remark,attr"`         //具体描述
	Opcode        string `xml:"opcode,attr"`         //操作码
}

func (RouteRequest *RouteRequest) CallSF() ([]RouteResponse, error) {
	req := NewRequest("RouteService", RouteRequest)
	resp, err := req.send()
	return resp.([]RouteResponse), err
}
