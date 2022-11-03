package delivery

import (
	"be12/mentutor/features/forum"
)

type Request struct {
	// ID      uint   `json:"id"`
	Caption string `json:"caption" form:"caption"`
	Images  string `json:"images" form:"images"`
}

func ToDomain(i interface{}) forum.Status {
	switch i.(type) {
	case Request:
		cnv := i.(Request)
		return forum.Status{Caption: cnv.Caption, Images: cnv.Images}
	}
	return forum.Status{}
}
