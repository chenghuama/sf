// @description 顺丰接口的sdk封装
// @link        https://github.com/chenghuama/sf
// @license     https://github.com/chenghuama/sf/blob/master/LICENSE
// @authors     chenghuama(chenghuama@vip.qq.com)
// @team yl10/忆零
package sf

import (
	"fmt"
	//"github.com/astaxie/beego"
	"testing"
	//"time"
	"encoding/xml"
	//"github.com/astaxie/beego/httplib"
)

var (
	testorderid string
	testmailno  string
)

func TestMain(t *testing.T) {
	testorderid = "XJFS_071100233"
}

//下单接口
func TestOrderService(t *testing.T) {

	orderxml := `<Order orderid='` + testorderid + `'

              express_type='1'

              j_company='xx公司'

              j_contact='客服'

              j_tel='025-10106699'

              j_mobile='13800138000'

              j_province='北京'

              j_city='北京市'

              j_county='海淀区'

              j_address='北京市海淀区科学园科健路328号'

              d_company='顺丰速运'

              d_contact='小顺'

              d_tel='0755-33992159'

              d_mobile='15602930913'

              d_province='广东省'

              d_city='深圳市'

              d_county='福田区'

              d_address='广东省深圳市福田区新洲十一街万基商务大厦10楼'

              parcel_quantity='1'

              pay_method='1'

              custid='0216838753'

              cargo_total_weight='2.18'

              sendstarttime='2014-07-11 12:07:11'

              remark='' >

       <Cargo name='服装' count='1' unit='台' weight='2.36' amount='2000' currency='CNY' source_area='中国'></Cargo>

       <Cargo name='手机' count='1' unit='台' weight='2.36' amount='2000' currency='CNY' source_area='中国'></Cargo>

       <AddedService name='COD' value='2000' value1='0216838753' />

       <AddedService name='INSURE' value='2000' />

</Order>
`
	//fmt.Println(sfrequesttoxml(sfxml))
	o := Order{}
	xml.Unmarshal([]byte(orderxml), &o)

	resp, err := o.CallSF()
	if err != nil {
		fmt.Println(err)
	} else {
		testmailno = resp[0].Mailno
	}

	fmt.Printf("下单接口测试正常，返回运单号为:/%s", testmailno)
}

//确定订单接口
func TestConfirm(t *testing.T) {
	o := OrderConfirm{}
	o.Orderid = testorderid
	o.Mailno = testmailno
	o.Dealtype = "1"

	op1 := OrderConfirmOption{}

	op1.Weight = "3.56"
	op1.Volume = "43200,1,1"
	o.OrderConfirmOption = []*OrderConfirmOption{&op1}

	resp, err := o.CallSF()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
	fmt.Println("确认发货接口测试正常")
}

//筛单接口
func TestOrderFilterService(t *testing.T) {
	o := OrderFilter{}
	o.DAddress = "上海市浦东开发区张东路玉兰香苑小区4其117幢1001室"
	o.FilterType = "1"
	o.Orderid = testorderid
	fo := OrderFilterOption{}
	fo.JAddress = "广东省深圳市福田区新洲十一街万基商务大厦10楼"
	fo.DTel = "15964349164"
	fo.JCustid = "0216838753"
	o.OrderFilterOption = []*OrderFilterOption{&fo}

	resp, err := o.CallSF()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	fmt.Println("筛单接口测试正常")
}

//订单结果确认

func TestOrderSearchService(t *testing.T) {
	o := OrderSearch{}
	o.Orderid = testorderid
	resp, err := o.CallSF()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	fmt.Println("订单确认接口测试正常")
}

//路由查询
func TestRouteService(t *testing.T) {
	o := RouteRequest{}
	o.MethodType = "1"
	o.TrackingType = 1
	switch o.MethodType {
	case "1":
		o.TrackingNumber = testmailno
	case "2":
		o.TrackingNumber = testorderid
	}

	resp, err := o.CallSF()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
		fmt.Println("路由查询接口测试通过")
	}
}
