// 代码生成时间: 2025-10-10 01:56:35
package main

import (
    "beego/logs"
    "beego/orm"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "strings"
)

// Device represents a device in the IoT system
type Device struct {
    Id      int    `orm:"auto"`
    Name    string `orm:"size(100)"`
    Type    string `orm:"size(100)"`
    Status  string `orm:"size(100)"`
    Data    string `orm:"size(1024)"`
}

// DeviceController handles the device operations
type DeviceController struct {
    beego.Controller
}

// AddDevice adds a new device to the system
func (this *DeviceController) AddDevice() {
    var d Device
    if err := json.Unmarshal(this.Ctx.Input.RequestBody, &d); err != nil {
        this.Data["json"] = map[string]string{"error": "Invalid JSON format"}
        this.ServeJSON()
        return
    }
    if _, err := orm.Insert(&d); err != nil {
        this.Data["json"] = map[string]string{"error": "Failed to add device"}
    } else {
        this.Data["json"] = map[string]string{"success": "Device added successfully"}
    }
    this.ServeJSON()
}

// GetDevice retrieves a device by its ID
func (this *DeviceController) GetDevice() {
    var d Device
    id := this.Ctx.Input.Param(":id")
    if err := orm.QueryTable(Device{}).Filter("id", id).One(&d); err != nil {
        this.Data["json"] = map[string]string{"error": "Device not found"}
    } else {
        this.Data["json"] = map[string]Device{
            "device": d,
        }
    }
    this.ServeJSON()
}

// UpdateDevice updates an existing device
func (this *DeviceController) UpdateDevice() {
    var d Device
    id := this.Ctx.Input.Param(":id\)
    if err := json.Unmarshal(this.Ctx.Input.RequestBody, &d); err != nil {
        this.Data["json"] = map[string]string{"error": "Invalid JSON format"}
        this.ServeJSON()
        return
    }
    d.Id, _ = strconv.Atoi(id)
    if _, err := orm.Update(&d); err != nil {
        this.Data["json"] = map[string]string{"error": "Failed to update device"}
    } else {
        this.Data["json"] = map[string]string{"success": "Device updated successfully"}
    }
    this.ServeJSON()
}

// DeleteDevice deletes a device by its ID
func (this *DeviceController) DeleteDevice() {
    id := this.Ctx.Input.Param(":id\)
    d := Device{Id: int(id)}
    if _, err := orm.Delete(&d); err != nil {
        this.Data["json"] = map[string]string{"error": "Failed to delete device"}
    } else {
        this.Data["json"] = map[string]string{"success": "Device deleted successfully"}
    }
    this.ServeJSON()
}

func main() {
    beego.Router("/device/add", &DeviceController{}, "post:AddDevice")
    beego.Router("/device/:id", &DeviceController{}, "get:GetDevice")
    beego.Router("/device/:id", &DeviceController{}, "put:UpdateDevice")
    beego.Router("/device/:id", &DeviceController{}, "delete:DeleteDevice\)
    beego.Run()
}