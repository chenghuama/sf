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
//"time"
//"github.com/astaxie/beego/httplib"
)

//路由推送
type WaybillRoute struct {
	Id            string `xml:"id,attr"`            //路由编号，每一个 id 代表一条不同的路由
	Mailno        string `xml:"mailno,attr"`        //运单号
	Orderid       string `xml:"orderid,attr"`       //订单号
	AcceptTime    string `xml:"acceptTime,attr"`    //路由产生时间
	AcceptAddress string `xml:"acceptAddress,attr"` //路由发生城市
	Remark        string `xml:"remark,attr"`        //路由说明
	OpCode        string `xml:"opCode,attr"`        //操作码
}

type WaybillRouteRequest struct {
	Service string         `xml:"service,attr"`
	Lang    string         `xml:"lang,attr"`
	Routes  []WaybillRoute `xml:"Body>WaybillRoute"`
}

//返回值
type WaybillRouteResponse struct {
	Id      string `xml:"id,attr"`       //成功接收的路由编号，如果有多个路由编 号 ， 以 逗 号 分 隔
	IdError string `xml:"id_error,attr"` //未成功接收的路由编号，如果有多个路由 编 号 ， 以 逗 号 分 隔
}

func CallRoutePush(WaybillRoute *WaybillRoute) interface{} {
	req := NewRequest("RoutePushService", &WaybillRoute)

	str, err := req.send()

	if err != nil {
		return ""
	}
	return str
}

func RouteError(msg string) []byte {
	str := `<Response service=”RoutePushService”>
<Head>ERR</Head>
<ERROR code="NNNN">` + msg + "</ERROR></Response>"
	return []byte(str)
}
