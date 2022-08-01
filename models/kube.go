package models

type KubeNameSpace struct {
	Cluster string `json:"cluster" example:"hza"`
	NameSpace string `json:"namespace" example:"example"`
}
