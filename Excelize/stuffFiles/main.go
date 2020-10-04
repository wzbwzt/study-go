/**
 *@Description 对excel的处理
 *@Auth wzb 2020/9/10 21:02
 **/
package stuffFiles

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

//Create
func main1() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	if err := f.SaveAs("Book1.xls"); err != nil {
		fmt.Println(err)
	}
}

type StuffFilesInfo struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StuffCode            string             `protobuf:"bytes,2,opt,name=stuff_code,json=stuffCode,proto3" json:"stuff_code,omitempty"`
	StuffModel           string             `protobuf:"bytes,4,opt,name=stuff_model,json=stuffModel,proto3" json:"stuff_model,omitempty"`
	Specifications       string             `protobuf:"bytes,5,opt,name=specifications,proto3" json:"specifications,omitempty"`
	Unit                 string             `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	Manufacturer         string             `protobuf:"bytes,7,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	ManufacturerCode     string             `protobuf:"bytes,8,opt,name=manufacturer_code,json=manufacturerCode,proto3" json:"manufacturer_code,omitempty"`
	ManufacturerPrice    float64            `protobuf:"fixed64,9,opt,name=manufacturer_price,json=manufacturerPrice,proto3" json:"manufacturer_price,omitempty"`
	ManufacturerData     string             `protobuf:"bytes,10,opt,name=manufacturer_data,json=manufacturerData,proto3" json:"manufacturer_data,omitempty"`
	Apply                string             `protobuf:"bytes,11,opt,name=apply,proto3" json:"apply,omitempty"`
	Security             bool               `protobuf:"varint,12,opt,name=security,proto3" json:"security,omitempty"`
	Place                string             `protobuf:"bytes,13,opt,name=place,proto3" json:"place,omitempty"`
	IsImport             bool               `protobuf:"varint,14,opt,name=is_import,json=isImport,proto3" json:"is_import,omitempty"`
	Img_No               string             `protobuf:"bytes,15,opt,name=img_No,json=imgNo,proto3" json:"img_No,omitempty"`
	Position             string             `protobuf:"bytes,16,opt,name=position,proto3" json:"position,omitempty"`
	Sum                  int64              `protobuf:"varint,17,opt,name=sum,proto3" json:"sum,omitempty"`
	CostPrice            float64            `protobuf:"fixed64,18,opt,name=cost_price,json=costPrice,proto3" json:"cost_price,omitempty"`
	CostFloor            int64              `protobuf:"varint,19,opt,name=cost_floor,json=costFloor,proto3" json:"cost_floor,omitempty"`
	PartsDate            string             `protobuf:"bytes,20,opt,name=parts_date,json=partsDate,proto3" json:"parts_date,omitempty"`
	CostAllPrice         float64            `protobuf:"fixed64,21,opt,name=cost_all_price,json=costAllPrice,proto3" json:"cost_all_price,omitempty"`
	QrCode               string             `protobuf:"bytes,22,opt,name=qr_code,json=qrCode,proto3" json:"qr_code,omitempty"`
	GuidePriceBase       float64            `protobuf:"fixed64,23,opt,name=guide_price_base,json=guidePriceBase,proto3" json:"guide_price_base,omitempty"`
	GuidePriceData       string             `protobuf:"bytes,24,opt,name=guide_price_data,json=guidePriceData,proto3" json:"guide_price_data,omitempty"`
	CheckDate            string             `protobuf:"bytes,25,opt,name=check_date,json=checkDate,proto3" json:"check_date,omitempty"`
	IsBom                bool               `protobuf:"varint,26,opt,name=is_bom,json=isBom,proto3" json:"is_bom,omitempty"`
	Brand                string             `protobuf:"bytes,28,opt,name=brand,proto3" json:"brand,omitempty"`
	Remark               string             `protobuf:"bytes,29,opt,name=remark,proto3" json:"remark,omitempty"`
	Misc                 string             `protobuf:"bytes,38,opt,name=misc,proto3" json:"misc,omitempty"`
	ClassifyId           int64              `protobuf:"varint,39,opt,name=classify_id,json=classifyId,proto3" json:"classify_id,omitempty"`
	ClassifyName         string             `protobuf:"bytes,44,opt,name=classify_name,json=classifyName,proto3" json:"classify_name,omitempty"`
	ChildStuffFiles      []*ChildStuffFiles `protobuf:"bytes,45,rep,name=childStuffFiles,proto3" json:"childStuffFiles,omitempty"`
	Accessories          string             `protobuf:"bytes,46,opt,name=accessories,proto3" json:"accessories,omitempty"`
}
type ChildStuffFiles struct {
	StuffFilesId         int64    `protobuf:"varint,1,opt,name=stuff_files_id,json=stuffFilesId,proto3" json:"stuff_files_id,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Specifications       string   `protobuf:"bytes,5,opt,name=specifications,proto3" json:"specifications,omitempty"`
	CostPrice            float64  `protobuf:"fixed64,6,opt,name=cost_price,json=costPrice,proto3" json:"cost_price,omitempty"`
	Type                 string   `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Unit                 string   `protobuf:"bytes,8,opt,name=unit,proto3" json:"unit,omitempty"`
}

//Read
func Init() {
	f, err := excelize.OpenFile("./stufffiles1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = f.RemoveRow("Sheet0", 1)
	if err != nil {
		panic(err)
	}
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet0")
	if err != nil {
		panic(err)
	}
	fmt.Println(rows[0])
	//var res []StuffFilesInfo
	for _,v:= range rows{
		tmp:=StuffFilesInfo{}
		tmp.Name=v[0]

	}

}