package model

type Information_receipts struct {
	Id                         int     `gorm:"column:id"`
	ParentId                   int     `gorm:"column:parent_id"`
	UserId                     int     `gorm:"column:user_id"`     // คนสร้างใบรับข้อมูล
	EmployeeId                 int     `gorm:"column:employee_id"` // ชื่อคนขาย
	CustomerId                 int     `gorm:"column:customer_id"`
	CustomerDeliveryAddressId  int     `gorm:"column:customer_delivery_address_id"`
	CustomerBillAddressId      int     `gorm:"column:customer_bill_address_id"`
	WarehouseId                int     `gorm:"column:warehouse_id"`
	Code                       string  `gorm:"column:code"`
	DocDate                    string  `gorm:"column:doc_date"`         // วันที่เอกสาร
	ProductionDate             string  `gorm:"column:production_date"`  // วันที่สั่งผลิต
	AppointmentDate            string  `gorm:"column:appointment_date"` // วันนัดรับสินค้า
	TypeTransfer               string  `gorm:"column:type_transfer"`    // วิธีการจัดส่ง
	SourceInformationReceiptId int     `gorm:"column:source_information_receipt_id"`
	SumTotalPrice              float64 `gorm:"column:sum_total_price"`
	SumDiscount                float64 `gorm:"column:sum_discount"` // ส่วนลด
	CreatedAt                  string  `gorm:"column:created_at"`
	UpdatedAt                  string  `gorm:"column:updated_at"`
	DeletedAt                  string  `gorm:"column:deleted_at"`
	SalesOrderCodeOld          string  `gorm:"column:sales_order_code_old"` // รหัสใบสั่งขายเก่าก่อนที่จะลบ
	QuotationCodeOld           string  `gorm:"column:quotation_code_old"`   // รหัสใบเสนอราคาเก่าก่อนที่จะลบ
	StatusVat                  int     `gorm:"column:status_vat"`
	IsInformationReceipt       int     `gorm:"column:is_information_receipt"` // เป็นใบรับข้อมูล
	RequestReceiptId           int     `gorm:"column:request_receipt_id"`
	SellRequestId              int     `gorm:"column:sell_request_id"`
	MemberRefId                int     `gorm:"column:member_ref_id"`
	IsRefApprove               int     `gorm:"column:is_ref_approve"`
	IsLineLiff                 int     `gorm:"column:is_line_liff"`
	AproveLineLiffCenter       int     `gorm:"column:aprove_line_liff_center"`
	OrderGroupCode             string  `gorm:"column:order_group_code"`
	Longitude                  string  `gorm:"column:longitude"`
	Latitude                   string  `gorm:"column:latitude"`
	Weight                     float64 `gorm:"column:weight"`
	DistanceKm                 float64 `gorm:"column:distance_km"`
	StatusConfirmLineLiff      int     `gorm:"column:status_confirm_line_liff"`
	CancelUserId               int     `gorm:"column:cancel_user_id"`
	CancelReason               string  `gorm:"column:cancel_reason"`
	PaidBeforeDatetime         string  `gorm:"column:paid_before_datetime"` // วันที่จ่ายเงินก่อนหมดอายุ
}
