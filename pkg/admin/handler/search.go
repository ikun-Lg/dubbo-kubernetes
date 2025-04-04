/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package handler

import (
	"net/http"

	"github.com/apache/dubbo-kubernetes/pkg/admin/model"
	"github.com/apache/dubbo-kubernetes/pkg/admin/service"
	core_runtime "github.com/apache/dubbo-kubernetes/pkg/core/runtime"
	"github.com/gin-gonic/gin"
)

// Search API Definition: https://app.apifox.com/project/3732499
// 全局搜索
func BannerGlobalSearch(rt core_runtime.Runtime) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 参考 API 定义 request 参数
		req := model.NewSearchReq()
		if err := c.ShouldBindQuery(req); err != nil {
			c.JSON(http.StatusBadRequest, model.NewErrorResp(err.Error()))
			return
		}

		// 根据 request 分流调用，如服务未实现继续实现

		var res *model.SearchRes
		switch req.SearchType {
		case "ip":
			instances, _ := service.BannerSearchIp(rt, req)
			res = convertInstancesToSearchRes(instances)
		case "instanceName":
			instances, _ := service.BannerSearchInstances(rt, req)
			res = convertInstancesToSearchRes(instances)
		case "appName":
			applications, _ := service.BannerSearchApplications(rt, req)
			res = convertApplicationsToSearchRes(applications)
		case "serviceName":
			sreq := &model.ServiceSearchReq{
				ServiceName: "",
				Keywords:    req.Keywords,
				PageReq:     req.PageReq,
			}
			services, _ := service.BannerSearchServices(rt, sreq)
			res = convertServicesToSearchRes(services)
		default:
			c.JSON(http.StatusBadRequest, model.NewErrorResp("invalid search type"))
			return
		}

		c.JSON(http.StatusOK, model.NewSuccessResp(model.NewPageData().WithData(res).WithTotal(len(res.Candidates)).WithPageSize(req.PageSize).WithCurPage(req.PageOffset)))
	}
}

func convertInstancesToSearchRes(pagedInstances *model.SearchPaginationResult) *model.SearchRes {
	instances := pagedInstances.List.([]*model.SearchInstanceResp)
	res := &model.SearchRes{}
	if len(instances) == 0 {
		res.Find = false
		return res
	}
	for _, ins := range instances {
		res.Candidates = append(res.Candidates, ins.Name)
	}
	res.Find = true
	return res
}

func convertApplicationsToSearchRes(apps []*model.ApplicationSearchResp) *model.SearchRes {
	res := &model.SearchRes{}
	if len(apps) == 0 {
		res.Find = false
		return res
	}
	for _, app := range apps {
		res.Candidates = append(res.Candidates, app.AppName)
	}
	res.Find = true
	return res
}

func convertServicesToSearchRes(pagedServices *model.SearchPaginationResult) *model.SearchRes {
	services := pagedServices.List.([]*model.ServiceSearchResp)
	res := &model.SearchRes{}
	if len(services) == 0 {
		res.Find = false
		return res
	}
	for _, s := range services {
		res.Candidates = append(res.Candidates, s.ServiceName)
	}
	res.Find = true
	return res
}
