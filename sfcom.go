// @description 顺丰接口的sdk封装
// @link        https://github.com/chenghuama/sf
// @license     https://github.com/chenghuama/sf/blob/master/LICENSE
// @authors     chenghuama(chenghuama@vip.qq.com)
// @team yl10/忆零
package sf

import (
	"cominterface/com"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"html"

	"strings"
)

var (
	incode    = beego.AppConfig.String("incode")
	checkword = beego.AppConfig.String("checkword")
	sfurl     = beego.AppConfig.String("sfurl")
)

const (
	ConstService_OrderService         = "OrderService"         //下订单
	ConstService_OrderConfirmService  = "OrderConfirmService"  //订单发货确认
	ConstService_OrderReverseService  = "OrderReverseService"  //下退货订单（含筛选）
	ConstService_OrderExchangeService = "OrderExchangeService" //下换货订单（含筛选）
	ConstService_OrderSearchService   = "OrderSearchService"   //订单结果查询
	ConstService_OrderFilterService   = "OrderFilterService"   //订单筛选
	ConstService_RouteService         = "RouteService"         //路由查询

)

type Request struct {
	XMLName xml.Name `xml:"Request"`
	Service string   `xml:"service,attr"`
	Lang    string   `xml:"lang,attr"`
	Head    string   `xml:"Head"`

	Body struct {
		XMLName  xml.Name    `xml:Body`
		Bodydata interface{} `xml:"Body"`
	}
}

//正常返回类型
type Response struct {
	XMLName xml.Name `xml:"Response"`
	Service string   `xml:"service,attr"`
	Head    string   `xml:"Head"`
	Body    struct {
		XMLName  xml.Name    `xml:Body`
		Bodydata interface{} `xml:"Body"`
	}
	ERROR struct {
		Code string `xml:"code,attr"`
	} `xml:"ERROR"`
}

type SoaResponse struct {
	XMLNname xml.Name `xml:"Envelope"`
	Body     struct {
		SfexpressServiceResponse struct {
			XMLName xml.Name `xml:"sfexpressServiceResponse"`
			Return  string   `xml:"return"`
		}
	}
}
type Sf struct {
	Service  string
	Request  *Request
	Response *Response
}

//type SOAPEnvelope struct {
//	//Envelope xml.Name `xml:"soap:Envelope"`
//	NSsoap string `xml:"xmlns:soap,attr"`
//	NSxsi  string `xml:"xmlns:xsi,attr"`
//	NSxsd  string `xml:"xmlns:xsd,attr"`
//	Body   SOAbody
//}

//构造xml
func sfrequesttoxml(s string) string {
	xmlstr := `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
  <soap:Body>
    <sfexpressService xmlns="http://service.expressservice.integration.sf.com/">
      <arg0 xmlns="">` + "%s" + `</arg0>
    </sfexpressService>
  </soap:Body>
</soap:Envelope>`
	return fmt.Sprintf(xmlstr, s)

}

//创建请求，传入服务名称和请求实例
func NewRequest(service string, body interface{}) *Request {
	req := Request{}
	req.Head = incode + "," + checkword
	req.Lang = "zh-CN"
	req.Service = service
	req.Body.Bodydata = body
	return &req
}

//创建请求，创建
func NewResponse(body interface{}) *Response {
	res := Response{}
	res.Body.Bodydata = body
	return &res
}

func getbodyData(strxml string) string {

	strxml = strings.Replace(strxml, "<?xml version='1.0' encoding='UTF-8'?>", "", 1)
	strxml = strings.Replace(strxml, `<?xml version="1.0" encoding="UTF-8"?>`, "", 1)

	doc, err := com.LoadByXml(strxml)

	if err != nil {
		fmt.Println("aaa" + err.Error())
		return ""
	}
	fmt.Println(strxml)
	return doc.Node("Body").ToString()
}

type SFrequest interface {
	CallSF()
	send()
}

func (req *Request) send() (responsedata interface{}, err error) {

	httpreq := httplib.Post(sfurl)
	httpreq.Header("Content-Type", "application/soap+xml; charset=utf-8")
	resp := Response{}
	body, err := xml.Marshal(req)

	if err != nil {
		return nil, err
	}

	//fmt.Println(sfrequesttoxml(html.EscapeString(string(body))))
	httpreq.Body([]byte(sfrequesttoxml(html.EscapeString(string(body)))))
	data, err1 := httpreq.Bytes()
	if err1 != nil {
		return nil, err1
	}

	res := SoaResponse{}
	if err := xml.Unmarshal(data, &res); err != nil {
		fmt.Println(err)
		return nil, err
	}
	//获取return里的值
	resopnxml := res.Body.SfexpressServiceResponse.Return
	//fmt.Println(resopnxml)
	if err := xml.Unmarshal([]byte(resopnxml), &resp); err != nil {
		return nil, err
	}

	//解析值，如果head为ERR也视为失败

	if resp.Head == "ERR" {
		return nil, errors.New(resp.ERROR.Code)
	}
	dataxml := []byte(getbodyData(resopnxml))
	switch req.Body.Bodydata.(type) {

	case Order, *Order:
		var resbody struct {
			XMLName  xml.Name        `xml:"Body"`
			Bodydata []OrderResponse `xml:"OrderResponse"`
		}

		if err := xml.Unmarshal(dataxml, &resbody); err != nil {
			return nil, err
		} else {
			return resbody.Bodydata, nil
		}

	case OrderChange, *OrderChange:
		var resbody struct {
			XMLName  xml.Name                `xml:"Body"`
			Bodydata []OrderExchangeResponse `xml:"OrderExchangeResponse"`
		}

		if err := xml.Unmarshal(dataxml, &resbody); err != nil {
			return nil, err
		} else {
			return resbody.Bodydata, nil
		}
	case OrderConfirm, *OrderConfirm:
		var resbody struct {
			XMLName  xml.Name               `xml:"Body"`
			Bodydata []OrderConfirmResponse `xml:"OrderConfirmResponse"`
		}

		if err := xml.Unmarshal(dataxml, &resbody); err != nil {
			return nil, err
		} else {
			return resbody.Bodydata, nil
		}
	case OrderFilter, *OrderFilter:
		var resbody struct {
			XMLName  xml.Name              `xml:"Body"`
			Bodydata []OrderFilterResponse `xml:"OrderFilterResponse"`
		}
		fmt.Println(string(dataxml))
		if err := xml.Unmarshal(dataxml, &resbody); err != nil {
			return nil, err
		} else {
			return resbody.Bodydata, nil
		}
	case RouteRequest, *RouteRequest:
		var resbody struct {
			XMLName  xml.Name        `xml:"Body"`
			Bodydata []RouteResponse `xml:"RouteResponse"`
		}

		fmt.Println(string(dataxml))
		if err := xml.Unmarshal(dataxml, &resbody); err != nil {
			return nil, err
		} else {
			return resbody.Bodydata, nil
		}
	case OrderBack, *OrderBack:
		var resbody struct {
			XMLName  xml.Name               `xml:"Body"`
			Bodydata []OrderReverseResponse `xml:"OrderReverseResponse"`
		}

		if err := xml.Unmarshal(dataxml, &resbody); err != nil {
			return nil, err
		} else {
			return resbody.Bodydata, nil
		}
	case OrderSearch, *OrderSearch:
		var resbody struct {
			XMLName  xml.Name              `xml:"Body"`
			Bodydata []OrderSearchResponse `xml:"OrderResponse"`
		}

		if err := xml.Unmarshal(dataxml, &resbody); err != nil {
			return nil, err
		} else {
			return resbody.Bodydata, nil
		}

	default:
		return nil, errors.New("未定义的请求类型！")
	}

}
