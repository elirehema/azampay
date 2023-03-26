package datas

type Properties struct {
	PropertyOne string `json:"property1" binding:"required"`
	PropertyTwo string `json:"property2" binding:"required"`
}
