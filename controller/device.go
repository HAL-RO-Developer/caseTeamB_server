package controller

import (
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var Device = deviceimpl{}

type deviceimpl struct {
	DeviceId string `json:"device_id"`
	ChildId  int    `json:"child_id"`
}

type deviceInfo struct {
	DeviceId string `json:"device_id"`
	Alive    bool   `json:"device_alive"`
}

// デバイスID発行
func (d *deviceimpl) CreateNewDevice(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	req, ok := validation.ChildrenInfoValidation(c)
	if !ok {
		return
	}
	pin, ok := service.CreateDevice(name, req.ChildId)
	if !ok {
		response.BadRequest(gin.H{"error": pin}, c)
		return
	}
	response.Json(gin.H{"pin": pin}, c)
}

// デバイス一覧取得
func (d *deviceimpl) ListDevice(c *gin.Context) {
	var userDevices []deviceInfo
	var device deviceInfo
	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	devices, find := service.GetDeviceId(name)
	if find {
		for i := 0; i < len(devices); i++ {
			device.DeviceId = devices[i].DeviceId
			device.Alive = false
			if devices[i].Mac != "" {
				device.Alive = true
			}
			userDevices = append(userDevices, device)
		}
		response.Json(gin.H{"devices": userDevices}, c)
		return
	}
	response.BadRequest(gin.H{"error": "デバイスが登録されていません。"}, c)
	return
}

// デバイスID削除
func (d *deviceimpl) DeleteDevice(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}

	deviceId := c.Param("device_id")

	if service.DeleteDeviceId(name, deviceId) {
		response.Json(gin.H{"success": "デバイスIDを削除しました。"}, c)
		return
	}
	response.BadRequest(gin.H{"error": "デバイスIDが見つかりません。"}, c)
}

// デバイスとIDの紐付け
func (d *deviceimpl) DeviceRegistration(c *gin.Context) {
	req, ok := validation.DeviceRegistrationCheck(c)
	if !ok {
		return
	}
	if !service.ExisByPin(req.Pin) {
		response.BadRequest(gin.H{"error": "pinが見つかりません。"}, c)
		return
	}
	if service.ExisByMac(req.Mac) {
		response.BadRequest(gin.H{"error": "その端末は登録済みです。"}, c)
		return
	} else {
		device_id, err := service.RegistrationButton(req.Pin, req.Mac)
		if err != nil {
			response.BadRequest(gin.H{"error": "データベースエラー"}, c)
			return
		}
		response.Json(gin.H{"device_id": device_id}, c)
	}
}
