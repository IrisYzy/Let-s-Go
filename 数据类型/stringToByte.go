package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	City  string `json:"city"`
	Phone string `json:"phone"`
	State int64 `json:"state"`
	EndAt int   `json:"end_at"`
	BeginAt  int    `json:"begin_at"`
	Province string `json:"province"`
	PayTimes int `json:"pay_times"`
	PolicyId int `json:"policy_id"`
	PolicyNo  string `json:"policy_no"`
	LastPayAt int    `json:"last_payat"`
	InsuredAge  string `json:"insured_age"`
	InsuredName string `json:"insured_name"`
	PayDuration string  `json:"pay_duration"`
	RenewAmount float64 `json:"renew_amount"`
	PolicyNameC     string `json:"policy_name_c"`
	ExpirationEndAt int    `json:"expiration_end_at"`
	YearInsuranceFee string `json:"year_insurance_fee"`
}

func main() {
	var data Data
	//var dataMap map[string]string
	//dataMap := make(map[string]string)
	s := `{"city": "xx市", "phone": "18921400524", "state": 20, "end_at": 1567871999000, "begin_at": 1536336000000, "province": "江西", "pay_times": 12, 
"policy_id": 5787562, "policy_no": "6666666666666", "last_payat": 1564538371000,
"insured_age": "56", "insured_name": "袁子", "pay_duration": "month", "renew_amount": 96, 
"policy_name_c": "全民百万", "expiration_end_at": 1567871999000, "year_insurance_fee": "1152.0"}`
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data.PolicyId)
	//fmt.Println([]byte(s))
	//err := json.Unmarshal([]byte(s),&dataMap)
	//if err != nil {
	//	fmt.Println(dataMap)
	//}

}
