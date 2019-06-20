package main

type Response struct {
	Id             string `json:"id" form:"id" query:"id"`
	Node           string `json:"node"`
	Address        string `json:"address"`
	ServiceId      string `json:"service_id"`
	ServiceAddress string `json:"service_address"`
	ServicePort    int    `json:"service_port"`
}
