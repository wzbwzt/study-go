package instore

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type InStoreInfo struct {
	BillNo               string              `protobuf:"bytes,1,opt,name=BillNo,proto3" json:"BillNo,omitempty"`//单据编号
	HandleBy             string              `protobuf:"bytes,2,opt,name=HandleBy,proto3" json:"HandleBy,omitempty"` //经办人name
	HandlerById          string              `protobuf:"bytes,20,opt,name=handler_by_id,json=handlerById,proto3" json:"handler_by_id,omitempty"`//经办人id
	InData               string              `protobuf:"bytes,3,opt,name=InData,proto3" json:"InData,omitempty"`//入库时间
	OriginalBill         string              `protobuf:"bytes,4,opt,name=OriginalBill,proto3" json:"OriginalBill,omitempty"`//原始单据
	HasTax               bool                `protobuf:"varint,5,opt,name=HasTax,proto3" json:"HasTax,omitempty"`//是否含税
	SupplyBy             string              `protobuf:"bytes,6,opt,name=SupplyBy,proto3" json:"SupplyBy,omitempty"`//供货单位
	MadeBy               string              `protobuf:"bytes,7,opt,name=MadeBy,proto3" json:"MadeBy,omitempty"` //制单人
	InReason             string              `protobuf:"bytes,8,opt,name=InReason,proto3" json:"InReason,omitempty"`//入库原因
	Valuation            float64             `protobuf:"fixed64,9,opt,name=Valuation,proto3" json:"Valuation,omitempty"`//暂估价格
	IsChanged            bool                `protobuf:"varint,45,opt,name=IsChanged,proto3" json:"IsChanged,omitempty"`//
	Misc                 string              `protobuf:"bytes,11,opt,name=Misc,proto3" json:"Misc,omitempty"`//冗余（暂存放备注）
	StuffBills           []*InStoreStuffInfo `protobuf:"bytes,10,rep,name=StuffBills,proto3" json:"StuffBills,omitempty"`//
	Accessories          string              `protobuf:"bytes,46,opt,name=accessories,proto3" json:"accessories,omitempty"`//附件
	DeptID               int64               `protobuf:"varint,12,opt,name=DeptID,proto3" json:"DeptID,omitempty"`//
	StoreID              int64               `protobuf:"varint,13,opt,name=StoreID,proto3" json:"StoreID,omitempty"`
	RegionID             int64               `protobuf:"varint,14,opt,name=RegionID,proto3" json:"RegionID,omitempty"`
	ShelvesID            int64               `protobuf:"varint,15,opt,name=ShelvesID,proto3" json:"ShelvesID,omitempty"`
	DeptName             string              `protobuf:"bytes,16,opt,name=DeptName,proto3" json:"DeptName,omitempty"`
	StoreName            string              `protobuf:"bytes,17,opt,name=StoreName,proto3" json:"StoreName,omitempty"`
	RegionName           string              `protobuf:"bytes,18,opt,name=RegionName,proto3" json:"RegionName,omitempty"`
	ShelvesName          string              `protobuf:"bytes,19,opt,name=ShelvesName,proto3" json:"ShelvesName,omitempty"`

	PayStatus string //付款状态
	HasReimbursement bool //报销情况

}

type BatchImport struct {
	InStoreInfo
	PayStatus string //付款状态
	HasReimbursement bool //报销情况
}

type InStoreStuffInfo struct {
	StuffFilesID         int64    `protobuf:"varint,1,opt,name=StuffFilesID,proto3" json:"StuffFilesID,omitempty"`//物料id
	StuffCount           int64    `protobuf:"varint,2,opt,name=StuffCount,proto3" json:"StuffCount,omitempty"`//物料数量(可能为复数)
	Price                float64  `protobuf:"fixed64,3,opt,name=Price,proto3" json:"Price,omitempty"`//单价
	AllPrice             float64  `protobuf:"fixed64,4,opt,name=AllPrice,proto3" json:"AllPrice,omitempty"`//总价
	Freight              float64  `protobuf:"fixed64,5,opt,name=Freight,proto3" json:"Freight,omitempty"`//运费
	Name                 string   `protobuf:"bytes,7,opt,name=Name,proto3" json:"Name,omitempty"`//物料名字
	Code                 string   `protobuf:"bytes,8,opt,name=Code,proto3" json:"Code,omitempty"`//物料代码
	Specifications       string   `protobuf:"bytes,9,opt,name=Specifications,proto3" json:"Specifications,omitempty"`//规格
	Type                 string   `protobuf:"bytes,10,opt,name=Type,proto3" json:"Type,omitempty"`//型号
	Unit                 string   `protobuf:"bytes,11,opt,name=Unit,proto3" json:"Unit,omitempty"`//单位
	ChildRemark          string   `protobuf:"bytes,12,opt,name=ChildRemark,proto3" json:"ChildRemark,omitempty"`//备注
	CostPrice            float64  `protobuf:"fixed64,13,opt,name=CostPrice,proto3" json:"CostPrice,omitempty"`//成本单价
	StuffQrCodes         string   `protobuf:"bytes,14,opt,name=StuffQrCodes,proto3" json:"StuffQrCodes,omitempty"`
}
type InStore struct {
	CorpID       int64          `gorm:"not null"`       //企业id
	BillNo       string         `gorm:"index;not null"` //单据编号
	HandleBy     string         `gorm:"index"`          //经办人
	HandleByID   string         `gorm:"index"`          //经办人ID
	InDate       string         `gorm:"not null"`       //入库日期
	OriginalBill string         `gorm:"index;not null"` //原始单据
	HasTax       bool           `gorm:"not null"`       //是否含税
	SupplyBy     string         `gorm:"index;not null"` //供货单位
	MadeBy       string         `gorm:"not null"`       //制单人
	InReason     string         `gorm:"index;not null"` //入库原因
	Valuation    float64        `gorm:"not null"`       //暂估价值
	Misc         string         `gorm:"size:512"`
	StuffBills   []*InStuffInfo `gorm:"-"` //忽略这个字段（材料清单）
	Accessories  string         `gorm:"type:json"`
	//w$:入库单据的付款和报销情况
	PayStatus         string  `gorm:"default:'1'"` //付款状态（1未知（默认），2不需付款，3待付款，4已付部分款，5已付款）
	PayDate           string  //付款时间
	HasInvoice        bool    //是否有发票
	InvoiceNum        string  //发票编号
	PayNum            float32 //付款金额
	HasReimbursement  bool    //是否报销
	ReimbursementDate string  //报销时间

	InStoreType int32 //入库单据类型；1：入库单；2：借件入库类型
	//<-----------------可根据以下字段查询货位里的材料的入库记录---------------->

	DeptID      int64  `gorm:"not null"`
	DeptName    string `gorm:"not null"`
	StoreID     int64  `gorm:"not null"`
	StoreName   string `gorm:"not null"`
	RegionID    int64  `gorm:"not null"`
	RegionName  string `gorm:"not null"`
	ShelvesID   int64  `gorm:"not null"`
	ShelvesName string `gorm:"not null"`
}

type InStuffInfo struct {
	StuffFilesID   int64   //物料ID
	StuffCount     int64   //入库数量
	Price          float64 //采购单价
	AllPrice       float64 //入库总价
	Freight        float64 //运输费用
	ChildRemark    string  //备注
	Name           string  //物料名字
	Code           string  //物料代码
	Specifications string  //规格
	Type           string  //型号
	Unit           string  //单位
	CostPrice      float64 //成本单价
	StuffQrCodes   string
}

//Read
func Init() {
	f, err := excelize.OpenFile("C:\\Users\\LM-LL\\Desktop\\inStore1(old).xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = f.RemoveRow("Sheet4", 1)
	if err != nil {
		panic(err)
	}
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet4")
	if err != nil {
		panic(err)
	}
	//fmt.Println(rows[0])
	var res []*BatchImport
    rowLine:=0
	for _,v:= range rows{
         rowLine++
		parse, err := time.Parse("2006-01-02", v[3])
		if err != nil {
			panic(err)
		}
		formatinDate := parse.Format(time.RFC3339)
	  var isTax bool
		if v[7]=="是"{
			isTax=true
		}else {
			isTax=false
		}
		var payStatus string
		if v[9]=="未知" {
			payStatus="1"
		}
		var  hasInvoice bool
		if v[10]=="未报销" {
			hasInvoice=false
		}else {
			hasInvoice=true
		}
		tmp:=&BatchImport{
			InStoreInfo: InStoreInfo{
				BillNo:v[0],
				OriginalBill:v[1],
				HandleBy: v[2],
				InData:formatinDate,
				InReason: v[4],
				SupplyBy:v[5],
				HasTax: isTax,
				Misc:v[11],
				MadeBy: v[12],
			},
			PayStatus:payStatus ,
			HasReimbursement:hasInvoice ,
		}
		//处理材料清单
		childStuffs:= v[6]
		split := strings.Split(childStuffs, ";")
		for _,vv:= range split{
			vv = strings.TrimSpace(vv)
			if len(vv)!=0 {
				vtmp,err := parseChildStuff(vv)
				if err!=nil {
					fmt.Printf("Failed: %s;data:%s;Line:%d\n",err,v,rowLine)
					return
				}
				tmp.StuffBills=append(tmp.StuffBills,&vtmp)
			}
			continue
		}
		res=append(res,tmp)
	}
	fmt.Println(rowLine)
	var totalCount int32
	for _,re:= range res{
		if re.StuffBills!=nil {
			for _,vv:=range re.StuffBills{
				indent, err := json.MarshalIndent(*vv, "", "\t")
				if err != nil {
					panic(err)
				}
				fmt.Println(string(indent))
				totalCount+=int32(vv.StuffCount)
			}
		}
		//fmt.Println(*re)
	}
	fmt.Println(totalCount)
}


func parseChildStuff(stuffOne string)(vtmp InStoreStuffInfo,err error){
	defer func() {
		if e:= recover();e!=nil {
			err=e.(error)
		}
	}()
	compile := regexp.MustCompile("\\d+\\.?\\d*|\\D*")
	details := strings.Split(stuffOne, "|")
	for k,v:=range details{
		space := strings.TrimSpace(v)
		details[k]=space
	}
	vtmp.Code=details[0]
	vtmp.Name=details[1]
	if len(details)==3 {
		all := compile.FindAllStringSubmatch(details[2], 2)
		//count, err := strconv.ParseFloat(all[0][0], 32)
		count, err := strconv.Atoi(all[0][0])
		if err != nil {
				return vtmp, err
			}
			vtmp.StuffCount=int64(count)
		if len(all) > 1 {
			unit:= all[1][0]
			vtmp.Unit=unit
		}

	}
	if len(details)==4 {
		//数量+单位可能在第三位或者在第四位
		all := compile.FindAllStringSubmatch(details[3], 2)
		//count, err := strconv.ParseFloat(all[0][0], 32)
		count, err := strconv.Atoi(all[0][0])
 		if err != nil {
			all=compile.FindAllStringSubmatch(details[2],2)
			//atoi, err := strconv.ParseFloat(all[0][0], 32)
			atoi, err := strconv.Atoi(all[0][0])
			if err != nil {
				return vtmp,err
			}
			vtmp.StuffCount=int64(atoi)
			if len(all) > 1 {
				vtmp.Unit=all[1][0]
			}
			vtmp.ChildRemark=details[3]
 		}else {
			if len(all)>1 {
				unit:= all[1][0]
				vtmp.Unit=unit
			}
			vtmp.StuffCount=int64(count)
			vtmp.Type=details[2]
		}

	}
	if len(details)==5 {
		vtmp.Type=details[2]
		all := compile.FindAllStringSubmatch(details[4], 2)
		//count, err := strconv.ParseFloat(all[0][0], 32)
		count, err := strconv.Atoi(all[0][0])
		if err != nil {
			all = compile.FindAllStringSubmatch(details[3], 2)
			//atoi, err := strconv.ParseFloat(all[0][0], 32)
			atoi, err := strconv.Atoi(all[0][0])
			if err != nil {
				return vtmp,err
			}
			vtmp.StuffCount=int64(atoi)
			if len(all) > 1 {
				vtmp.Unit=all[1][0]
			}
			vtmp.Specifications = details[3]
		}else {
			if len(all)>1 {
				unit:= all[1][0]
				vtmp.Unit=unit
			}
			vtmp.StuffCount=int64(count)
			vtmp.Specifications=details[3]
		}
	}
	if len(details)==6 {
		vtmp.Type=details[2]
		vtmp.Specifications=details[3]
		vtmp.ChildRemark=details[5]
		all := compile.FindAllStringSubmatch(details[4], 2)
		//count, err := strconv.ParseFloat(all[0][0], 32)
		count, err := strconv.Atoi(all[0][0])
		if err != nil {
			return vtmp, err
		}
		if len(all)>1 {
			unit:= all[1][0]
			vtmp.Unit=unit
		}
		vtmp.StuffCount=int64(count)
	}
	 return vtmp, nil
}