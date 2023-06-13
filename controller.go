package appdrest

import (
	"net/url"
)

// Controller represents the AppDynamics Controller
type Controller struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Account  string `json:"account"`
	Protocol string `json:"protocol"`
	BaseURL  *url.URL
}
