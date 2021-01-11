package weixin_shop_pay

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// ProfitSharing 分账
type ProfitSharing struct {
	client *Client
}

// ReceiversAdd 添加分账接收方
func (t *ProfitSharing) ReceiversAdd(p *ProfitSharingReceiversAdd) (*ProfitSharingReceiversAddResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "/v3/ecommerce/profitsharing/receivers/add"
	resp, err := tool.PostRequest(t.client.config, urlPath, dataJsonByte)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := t.client.setErrorResponse(respData)
		if err != nil {
			return nil, err
		}
	}

	log.Println(string(respData))
	var output ProfitSharingReceiversAddResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// Apply 请求分账
func (t *ProfitSharing) Apply(p *ProfitSharingApply) (*ProfitSharingApplyResp, error) {

	var err error

	// 加密接收方姓名
	for index, receiver := range p.Receivers {
		log.Println("ReceiverName加密", receiver.ReceiverName)
		if receiver.ReceiverName != "" {
			p.Receivers[index].ReceiverName, err = tool.Encrypt(receiver.ReceiverName, t.client.config.PlatformPublicKey)
			if err != nil {
				return nil, err
			}
		}
	}

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	log.Println("分账请求参数", string(dataJsonByte))
	// 发起请求
	urlPath := "/v3/ecommerce/profitsharing/orders"
	resp, err := tool.PostRequest(t.client.config, urlPath, dataJsonByte)
	if err != nil {
		log.Println("分账错误", err.Error())
		return nil, err
	}

	log.Println("分账头信息", resp.Status)

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取分账结果错误", err.Error())
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := t.client.setErrorResponse(respData)
		if err != nil {
			return nil, err
		}
	}

	log.Println(string(respData))
	var output ProfitSharingApplyResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// Query 分账查询
func (t *ProfitSharing) Query(p *ProfitSharingQuery) (*ProfitSharingQueryResp, error) {

	// 发起请求
	urlPath := "/v3/ecommerce/profitsharing/orders"
	resp, err := tool.GetRequest(t.client.config, urlPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := t.client.setErrorResponse(respData)
		if err != nil {
			return nil, err
		}
	}

	log.Println(string(respData))
	var output ProfitSharingQueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// FinishOrder 完结分账
func (t *ProfitSharing) FinishOrder(p *ProfitSharingFinishOrder) (*ProfitSharingFinishOrderResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "/v3/ecommerce/profitsharing/finish-order"
	resp, err := tool.PostRequest(t.client.config, urlPath, dataJsonByte)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := t.client.setErrorResponse(respData)
		if err != nil {
			return nil, err
		}
	}

	log.Println(string(respData))
	var output ProfitSharingFinishOrderResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// ReturnOrders 分账回退
func (t *ProfitSharing) ReturnOrders(p *ProfitSharingReturnOrders) (*ProfitSharingReturnOrdersResp, error) {

	// 请求参数
	dataJsonByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// 发起请求
	urlPath := "/v3/ecommerce/profitsharing/returnorders"
	resp, err := tool.PostRequest(t.client.config, urlPath, dataJsonByte)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := t.client.setErrorResponse(respData)
		if err != nil {
			return nil, err
		}
	}

	log.Println(string(respData))
	var output ProfitSharingReturnOrdersResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// ReturnOrdersQuery 分账回退查询
func (t *ProfitSharing) ReturnOrdersQuery(p *ProfitSharingReturnOrdersQuery) (*ProfitSharingReturnOrdersQueryResp, error) {

	// 发起请求
	urlPath := "/v3/ecommerce/profitsharing/returnorders?sub_mchid=" + p.SubMchid + "&out_order_no=" + p.OutOrderNo + "&out_return_no=" + p.OutReturnNo
	resp, err := tool.GetRequest(t.client.config, urlPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := t.client.setErrorResponse(respData)
		if err != nil {
			return nil, err
		}
	}

	log.Println(string(respData))
	var output ProfitSharingReturnOrdersQueryResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// LeftOrderAmount 查询订单剩余待分金额
func (t *ProfitSharing) LeftOrderAmount(p *ProfitSharingLeftOrderAmount) (*ProfitSharingLeftOrderAmountResp, error) {

	// 发起请求
	urlPath := "/v3/ecommerce/profitsharing/orders/" + p.TransactionID + "/amounts"
	resp, err := tool.GetRequest(t.client.config, urlPath)
	if err != nil {
		return nil, err
	}

	// 解析返回参数
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err := t.client.setErrorResponse(respData)
		if err != nil {
			return nil, err
		}
	}

	log.Println(string(respData))
	var output ProfitSharingLeftOrderAmountResp
	err = json.Unmarshal(respData, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
