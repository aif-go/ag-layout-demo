package main

import (
	"ag-layout-demo/internal"

	"github.com/aif-go/ag-core/ag/ag_app"
	"github.com/aif-go/ag-core/ag/ag_log"

	"github.com/aif-go/ag-core/fxs"

	"github.com/aif-go/ag-core/ag/ag_service"
	hserver "github.com/aif-go/ag-core/contribute/aghertz/server"
	kserver "github.com/aif-go/ag-core/contribute/agkitex/server"

	"go.uber.org/fx"

	// _ "ag-layout-demo/api/logs/logserver"
	_ "go.uber.org/automaxprocs"
)

func main() {
	var fxopts []fx.Option

	fxopts = append(
		fxopts,
		mainFx,
		fx.Invoke(func(s *ag_app.App) {}),
	)

	fxapp := fx.New(
		fxopts...,
	)

	fxapp.Run()
}

var mainFx = fx.Module("main",
	/** conf **/
	// 初始化配置
	fxs.FxAgConfModule,

	// nacosconf
	// agnacos.FxNacosConfigMode,
	// agnacos.FxEnableNacosRemoteConfigMode,

	// 日志模块
	ag_log.FxAglogMode,

	// 根APP
	fxs.FxAppMode,

	// http服务
	hserver.FxAgHertzServerModule,

	// grpc服务
	kserver.FxKitexServerBaseModule,

	// 服务基础模块
	ag_service.FxAgServiceMode,

	// 数据库模块
	// agdb.FxAgDbModule,
	// gormdb.FxAicGromdbModule,

	// 内部组件
	internal.FxInternalModule,
)
