package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	data := `<?xml version="1.0" encoding="UTF-8"?>
<bpmn:definitions xmlns:bpmn="http://www.omg.org/spec/BPMN/20100524/MODEL" xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI" xmlns:di="http://www.omg.org/spec/DD/20100524/DI" xmlns:dc="http://www.omg.org/spec/DD/20100524/DC" xmlns:zeebe="http://camunda.org/schema/zeebe/1.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" id="Definitions_1" targetNamespace="http://bpmn.io/schema/bpmn" exporter="Zeebe Modeler" exporterVersion="0.1.0">
  <bpmn:process id="order-process" isExecutable="true">
    <bpmn:startEvent id="order-placed" name="Order Placed">
      <bpmn:outgoing>SequenceFlow_18tqka5</bpmn:outgoing>
    </bpmn:startEvent>
    <bpmn:endEvent id="order-delivered" name="Order Delivered">
      <bpmn:incoming>SequenceFlow_1qj94z0</bpmn:incoming>
    </bpmn:endEvent>
  </bpmn:process>
</bpmn:definitions>`
	type bpEn struct {
		XMLName         xml.Name `xml:"definitions"`
		BPMN            string   `xml:"xmlns:bpmn,attr"`
		BPMNDI          string   `xml:"bpmndi,attr"`
		DI              string   `xml:"di,attr"`
		DC              string   `xml:"dc,attr"`
		Zeebe           string   `xml:"zeebe,attr"`
		XSI             string   `xml:"xsi,attr"`
		ID              string   `xml:"id,attr"`
		TargetNamespace string   `xml:"targetNamespace,attr"`
		Exporter        string   `xml:"exporter,attr"`
		ExporterVersion string   `xml:"exporterVersion,attr"`
		Process         struct {
			ID           string `xml:"id,attr"`
			Name         string `xml:"name,attr"`
			IsExecutable bool   `xml:"isExecutable,attr"`
			StartEvent   struct {
				ID       string `xml:"id,attr"`
				Name     string `xml:"name,attr"`
				Outgoing string `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL outgoing"`
			} `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL startEvent"`
			EndEvent struct {
				ID       string `xml:"id,attr"`
				Name     string `xml:"name,attr"`
				Incoming string `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL incoming"`
			} `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL endEvent"`
		} `xml:"http://www.omg.org/spec/BPMN/20100524/MODEL process"`
	}
	var rsp bpEn
	err := xml.Unmarshal([]byte(data), &rsp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", rsp)
}
