// 代码生成时间: 2025-10-26 08:06:42
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/astaxie/beego"
)

// Node represents a blockchain node
type Node struct {
    ID string
}

// NodeManager manages the blockchain nodes
type NodeManager struct {
    nodes map[string]Node
}

// NewNodeManager creates a new NodeManager instance
func NewNodeManager() *NodeManager {
    return &NodeManager{
        nodes: make(map[string]Node),
    }
}

// AddNode adds a new node to the NodeManager
func (nm *NodeManager) AddNode(node Node) error {
    if _, exists := nm.nodes[node.ID]; exists {
        return fmt.Errorf("node with ID %s already exists", node.ID)
    }
    nm.nodes[node.ID] = node
    return nil
}

// RemoveNode removes a node from the NodeManager
func (nm *NodeManager) RemoveNode(nodeID string) error {
    if _, exists := nm.nodes[nodeID]; !exists {
        return fmt.Errorf("node with ID %s not found", nodeID)
    }
    delete(nm.nodes, nodeID)
    return nil
}

// GetNode retrieves a node by its ID
func (nm *NodeManager) GetNode(nodeID string) (Node, error) {
    node, exists := nm.nodes[nodeID]
    if !exists {
        return Node{}, fmt.Errorf("node with ID %s not found", nodeID)
    }
    return node, nil
}

// StartServer starts the Beego HTTP server with the node management routes
func StartServer() {
    nm := NewNodeManager()
    beego.Router("/node/add", &NodeManagerController{nm: nm}, "post:AddNode")
    beego.Router("/node/remove/:nodeID", &NodeManagerController{nm: nm}, "delete:RemoveNode")
    beego.Router("/node/get/:nodeID", &NodeManagerController{nm: nm}, "get:GetNode")
    beego.Run()
}

// NodeManagerController handles HTTP requests for node management
type NodeManagerController struct {
    nm *NodeManager
}

// AddNode adds a new node to the blockchain
func (ctrl *NodeManagerController) AddNode() {
    var node Node
    if err := beego.ParseJSON(beego.Ctx.Input, &node); err != nil {
        beego.Error(err)
        beego.ServeJSON(&beego.Controller{Ctx: beego.Ctx}, map[string]string{"error": "invalid JSON"})
        return
    }
    if err := ctrl.nm.AddNode(node); err != nil {
        beego.Error(err)
        beego.ServeJSON(&beego.Controller{Ctx: beego.Ctx}, map[string]string{"error": err.Error()})
        return
    }
    beego.ServeJSON(&beego.Controller{Ctx: beego.Ctx}, map[string]string{"message": "node added successfully"})
}

// RemoveNode removes a node from the blockchain
func (ctrl *NodeManagerController) RemoveNode() {
    nodeID := beego.Ctx.Input.Param(":nodeID")
    if err := ctrl.nm.RemoveNode(nodeID); err != nil {
        beego.Error(err)
        beego.ServeJSON(&beego.Controller{Ctx: beego.Ctx}, map[string]string{"error": err.Error()})
        return
    }
    beego.ServeJSON(&beego.Controller{Ctx: beego.Ctx}, map[string]string{"message": "node removed successfully"})
}

// GetNode retrieves a node by its ID
func (ctrl *NodeManagerController) GetNode() {
    nodeID := beego.Ctx.Input.Param(":nodeID")
    node, err := ctrl.nm.GetNode(nodeID)
    if err != nil {
        beego.Error(err)
        beego.ServeJSON(&beego.Controller{Ctx: beego.Ctx}, map[string]string{"error": err.Error()})
        return
    }
    beego.ServeJSON(&beego.Controller{Ctx: beego.Ctx}, map[string]Node{ "node": node })
}

func main() {
    StartServer()
}
