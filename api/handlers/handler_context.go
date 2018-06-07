package handlers

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

type Context struct {
	DB *gorm.DB
}

type HandlerWithContext func(context *Context, w http.ResponseWriter, r *http.Request)
