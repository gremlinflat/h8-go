package repository

import (
	. "C2-simple-rest-gin/model"
)

var DataBooks []Book

func init() {
	DataBooks = []Book{
		{ID: 1, Title: "Golang", Author: "GG Martin", Desc: "buku golang"},
		{ID: 2, Title: "Hitler Mati di Garut", Author: "Xang Aby", Desc: "Autobiografi Haji Hitler"},
	}
}
