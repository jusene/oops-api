package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.hho-inc.com/devops/oops-api/models"
	"gitlab.hho-inc.com/devops/oops-api/utils"
)

// @summary 列出 Deployment
// @description 列出 Deployment
// @tags kubernetes
// @produce json
// @accept json
// @param cluster query string true "集群 ack or hza"
// @param app query string true "应用名"
// @param proj query string true "项目名"
// @param env query string true "环境名"
// @success 200 {object} models.Empty
// @failure 404 {object} models.Empty
// @router /kubernetes/deployments [get]
func ListKubeDeployments(ctx *gin.Context) {
	cluster, ok := ctx.GetQuery("cluster")
	if  !ok {
		CheckError(ctx, 404, "未获取cluster")
		return
	}
	app, ok := ctx.GetQuery("app")
	if  !ok {
		CheckError(ctx, 404, "未获取app")
		return
	}
	proj, ok := ctx.GetQuery("proj")
	if  !ok {
		CheckError(ctx, 404, "未获取proj")
		return
	}
	env, ok := ctx.GetQuery("env")
	if  !ok {
		CheckError(ctx, 404, "未获取env")
		return
	}
	kubeClient := utils.NewKubeClient(cluster)
	kube := utils.Kube(kubeClient)
	deployment, err := kube.ListDeployment(app, proj, env)
	if err != nil {
		fmt.Println("err", err)
		CheckError(ctx, 500, err.Error())
		return
	}
	ctx.JSON(200, deployment)
}

// @summary 列出 Pods
// @description 列出 Pods
// @tags kubernetes
// @produce json
// @accept json
// @param cluster query string true "集群 ack or hza"
// @param app query string true "应用名"
// @param proj query string true "项目名"
// @param env query string true "环境名"
// @success 200 {object} models.Empty
// @failure 404 {object} models.Empty
// @router /kubernetes/pods [get]
func ListKubePods(ctx *gin.Context) {
	cluster, ok := ctx.GetQuery("cluster")
	if  !ok {
		CheckError(ctx, 404, "未获取cluster")
		return
	}
	app, ok := ctx.GetQuery("app")
	if  !ok {
		CheckError(ctx, 404, "未获取app")
		return
	}
	proj, ok := ctx.GetQuery("proj")
	if  !ok {
		CheckError(ctx, 404, "未获取proj")
		return
	}
	env, ok := ctx.GetQuery("env")
	if  !ok {
		CheckError(ctx, 404, "未获取env")
		return
	}
	kubeClient := utils.NewKubeClient(cluster)
	kube := utils.Kube(kubeClient)
	pods, err := kube.ListPods(app, proj, env)
	if err != nil {
		fmt.Println("err", err)
		CheckError(ctx, 500, err.Error())
		return
	}
	ctx.JSON(200, pods)
}

// @summary 列出 Service
// @description 列出 Services
// @tags kubernetes
// @produce json
// @accept json
// @param cluster query string true "集群 ack or hza"
// @param app query string true "应用名"
// @param proj query string true "项目名"
// @param env query string true "环境名"
// @success 200 {object} models.Empty
// @failure 404 {object} models.Empty
// @router /kubernetes/services [get]
func ListKubeServices(ctx *gin.Context) {
	cluster, ok := ctx.GetQuery("cluster")
	if  !ok {
		CheckError(ctx, 404, "未获取cluster")
		return
	}
	app, ok := ctx.GetQuery("app")
	if  !ok {
		CheckError(ctx, 404, "未获取app")
		return
	}
	proj, ok := ctx.GetQuery("proj")
	if  !ok {
		CheckError(ctx, 404, "未获取proj")
		return
	}
	env, ok := ctx.GetQuery("env")
	if  !ok {
		CheckError(ctx, 404, "未获取env")
		return
	}
	kubeClient := utils.NewKubeClient(cluster)
	kube := utils.Kube(kubeClient)
	service, err := kube.ListServices(app, proj, env)
	if err != nil {
		fmt.Println("err", err)
		CheckError(ctx, 500, err.Error())
		return
	}
	ctx.JSON(200, service)
}

// @summary 列出 Nodes
// @description 列出 Nodes
// @tags kubernetes
// @produce json
// @accept json
// @param cluster query string true "集群 ack or hza"
// @success 200 {object} models.Empty
// @failure 404 {object} models.Empty
// @router /kubernetes/nodes [get]
func ListKubeNodes(ctx *gin.Context) {
	cluster, ok := ctx.GetQuery("cluster")
	if  !ok {
		CheckError(ctx, 404, "未获取cluster")
		return
	}

	kubeClient := utils.NewKubeClient(cluster)
	kube := utils.Kube(kubeClient)
	nodes, err := kube.ListNodes()
	if err != nil {
		fmt.Println("err", err)
		CheckError(ctx, 500, err.Error())
		return
	}
	ctx.JSON(200, nodes)
}

// @summary 列出指定namespace的pods
// @description 列出指定namespace的pods
// @tags kubernetes
// @produce json
// @accept json
// @param cluster query string true "集群 ack or hza"
// @param namespace path string true "namespace"
// @success 200 {object} models.Empty
// @failure 404 {object} models.Empty
// @router /kubernetes/pods/{namespace} [get]
func ListKubeNameSpacePods(ctx *gin.Context) {
	cluster, ok := ctx.GetQuery("cluster")
	if  !ok {
		CheckError(ctx, 404, "未获取cluster")
		return
	}

	namespace := ctx.Param("namespace")

	kubeClient := utils.NewKubeClient(cluster)
	kube := utils.Kube(kubeClient)
	nodes, err := kube.ListNameSpacePods(namespace)
	if err != nil {
		fmt.Println("err", err)
		CheckError(ctx, 500, err.Error())
		return
	}
	ctx.JSON(200, nodes)
}

// @summary 列出指定namespace的deployments
// @description 列出指定namespace的deployments
// @tags kubernetes
// @produce json
// @accept json
// @param cluster query string true "集群 ack or hza"
// @param namespace path string true "namespace"
// @success 200 {object} models.Empty
// @failure 404 {object} models.Empty
// @router /kubernetes/deployments/{namespace} [get]
func ListKubeNameSpaceDeployments(ctx *gin.Context) {
	cluster, ok := ctx.GetQuery("cluster")
	if  !ok {
		CheckError(ctx, 404, "未获取cluster")
		return
	}

	namespace := ctx.Param("namespace")

	kubeClient := utils.NewKubeClient(cluster)
	kube := utils.Kube(kubeClient)
	deployments, err := kube.ListNameSpaceDeployments(namespace)
	if err != nil {
		fmt.Println("err", err)
		CheckError(ctx, 500, err.Error())
		return
	}
	ctx.JSON(200, deployments)
}

// @summary 列出指定namespace的services
// @description 列出指定namespace的services
// @tags kubernetes
// @produce json
// @accept json
// @param cluster query string true "集群 ack or hza"
// @param namespace path string true "namespace"
// @success 200 {object} models.Empty
// @failure 404 {object} models.Empty
// @router /kubernetes/services/{namespace} [get]
func ListKubeNameSpaceServices(ctx *gin.Context) {
	cluster, ok := ctx.GetQuery("cluster")
	if  !ok {
		CheckError(ctx, 404, "未获取cluster")
		return
	}

	namespace := ctx.Param("namespace")

	kubeClient := utils.NewKubeClient(cluster)
	kube := utils.Kube(kubeClient)
	services, err := kube.ListNameSpaceServices(namespace)
	if err != nil {
		fmt.Println("err", err)
		CheckError(ctx, 500, err.Error())
		return
	}
	ctx.JSON(200, services)
}

// @summary 创建namespace
// @description 创建namespace
// @tags kubernetes
// @produce json
// @accept json
// @param body body models.KubeNameSpace true "namespace"
// @success 200 {object} models.Empty
// @failure 404 {object} models.Empty
// @router /kubernetes/namespace [post]
func CreateKubeNameSpace(ctx *gin.Context) {
	var clusterNamespace models.KubeNameSpace

	if err := ctx.ShouldBindJSON(&clusterNamespace); err != nil {
		CheckError(ctx, 400, err.Error())
		return
	}

	kubeClient := utils.NewKubeClient(clusterNamespace.Cluster)
	kube := utils.Kube(kubeClient)
	result, err := kube.CreateNameSpace(clusterNamespace.NameSpace)
	if err != nil {
		CheckError(ctx, 500, err.Error())
		return
	}

	ctx.JSON(200, result)
}

