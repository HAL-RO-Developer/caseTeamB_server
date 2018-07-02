package controller

import (
	"time"

	"github.com/HAL-RO-Developer/caseTeamB_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB_server/service"
	"github.com/gin-gonic/gin"
)

var Device = deviceimpl{}

type deviceimpl struct {
}

type deviceInfo struct {
	ChildId int      `json:"child_id"`
	Devices []string `json:"child_devices"`
}

// デバイスID発行
func (d *deviceimpl) CreateNewDevice(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	req, ok := validation.ChildrenInfoValidation(c)
	if !ok {
		return
	}
	res, ok := service.CreateDevice(name, req.ChildId)
	if !ok {
		response.BadRequest(gin.H{"error": res}, c)
		return
	}

	timer := time.NewTimer(time.Minute * 10)
	go func() {
		<-timer.C
		_, find := service.GetPin(res)
		if find {
			service.PinRemove(res)
		}
	}()
	response.Json(gin.H{"pin": res}, c)
}

// デバイス一覧取得
func (d *deviceimpl) ListDevice(c *gin.Context) {
	var userDevices []deviceInfo
	var device deviceInfo

	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	_, find := service.GetDeviceId(name)
	if !find {
		response.BadRequest(gin.H{"error": "デバイスが登録されていません。"}, c)
		return
	}
	children, _ := service.GetChildInfo(name)
	for i := 0; i < len(children); i++ {
		devices, find := service.GetDeviceIdFromChildId(name, children[i].ChildId)
		if !find {

		} else {
			device.ChildId = devices[i].ChildId
			for j := 0; j < len(devices); j++ {
				if devices[i].Mac != "" {
					device.Devices = append(device.Devices, devices[i].DeviceId)
				}
			}
			userDevices = append(userDevices, device)
		}
	}
	response.Json(gin.H{"devices": userDevices}, c)
}

// デバイスID削除
func (d *deviceimpl) DeleteDevice(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
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

	if service.ExisByMac(req.Mac) {
		response.BadRequest(gin.H{"error": "その端末は登録済みです。"}, c)
		return
	} else {
		device_id, success := service.RegistrationDevice(req.Pin, req.Mac)
		if !success {
			response.BadRequest(gin.H{"error": device_id}, c)
			return
		}
		response.Json(gin.H{"device_id": device_id}, c)
	}
}
