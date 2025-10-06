// 代码生成时间: 2025-10-07 02:36:21
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
)

// LogisticsTrack 结构体用于表示物流跟踪信息
type LogisticsTrack struct {
    ID         int    `json:"id"`         // 物流跟踪ID
    TrackingID string `json:"tracking_id"` // 物流跟踪号
    Status     string `json:"status"`     // 当前状态
    CreatedAt  string `json:"created_at"` // 创建时间
}

// LogisticsController 控制器处理物流跟踪相关的请求
type LogisticsController struct {
    beego.Controller
}

// GetTrack 根据物流跟踪号查询物流信息
func (c *LogisticsController) GetTrack() {
    trackingID := c.GetString("tracking_id")
    
    o := orm.NewOrm()
    var track LogisticsTrack
    err := o.QueryTable("logistics_track").Filter("tracking_id", trackingID).One(&track)
    
    if err == orm.ErrNoRows {
        c.Data["json"] = map[string]string{ "error": "物流跟踪号不存在" }
        c.ServeJSON()
        return
    }
    
    if err != nil {
        c.Data["json"] = map[string]string{ "error": "查询物流信息失败" }
        c.ServeJSON()
        return
    }
    
    c.Data["json"] = track
    c.ServeJSON()
}

func main() {
    // 初始化Beego框架
    beego.Run(":8080")
}
