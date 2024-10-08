// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package horuser

import (
	"context"
	"fmt"
	"github.com/apache/dubbo-kubernetes/app/horus/basic/db"
	"github.com/apache/dubbo-kubernetes/app/horus/core/alert"
	"github.com/gammazero/workerpool"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	"time"
)

func (h *Horuser) RecoveryManager(ctx context.Context) error {
	go wait.UntilWithContext(ctx, h.recoveryCheck, time.Duration(h.cc.NodeRecovery.IntervalSecond)*time.Second)
	<-ctx.Done()
	return nil
}

func (h *Horuser) recoveryCheck(ctx context.Context) {
	data, err := db.GetRecoveryNodeDataInfoDate(h.cc.NodeRecovery.DayNumber)
	if err != nil {
		klog.Errorf("recovery check GetRecoveryNodeDataInfoDate err:%v", err)
		return
	}
	if len(data) == 0 {
		klog.Errorf("recovery check GetRecoveryNodeDataInfoDate zero.")
		return
	}
	wp := workerpool.New(5)
	for _, d := range data {
		d := d
		wp.Submit(func() {
			h.recoveryNodes(d)
		})

	}
	wp.StopWait()
}

func (h *Horuser) recoveryNodes(n db.NodeDataInfo) {
	addr := h.cc.PromMultiple[n.ClusterName]
	if addr == "" {
		klog.Errorf("recoveryNodes PromMultiple get addr empty.")
		klog.Infof("clusterName:%v nodeName:%v", n.ClusterName, n.NodeName)
		return
	}
	vecs, err := h.InstantQuery(addr, n.RecoveryQL, n.ClusterName, h.cc.NodeRecovery.PromQueryTimeSecond)
	if err != nil {
		klog.Errorf("recoveryNodes InstantQuery err:%v ql:%v", err, n.RecoveryQL)
		return
	}
	if len(vecs) != 1 {
		klog.Infof("Expected 1 result, but got: %d", len(vecs))
		return
	}
	if err != nil {
		klog.Errorf("recoveryNodes instantQuery err:%v ql:%v", err, n.RecoveryQL)
		return
	}
	klog.Infof("recoveryNodes check success.")
	err = h.UnCordon(n.NodeName, n.ClusterName)
	if err == nil {
		klog.Infof("Node %v is already uncordoned.", n.NodeName)
		return
	}
	res := "Success"
	if err != nil {
		res = fmt.Sprintf("failed:%v", err)
	}
	msg := fmt.Sprintf("\n【集群: %v】\n【异常节点恢复调度】\n【已恢复调度节点: %v】\n【处理结果：%v】\n【日期: %v】\n", n.ClusterName, n.NodeName, res, n.CreateTime)
	alert.DingTalkSend(h.cc.NodeRecovery.DingTalk, msg)
	alert.SlackSend(h.cc.CustomModular.Slack, msg)

	pass, err := n.RecoveryMarker()
	klog.Infof("RecoveryMarker result pass:%v err:%v", pass, err)
}
