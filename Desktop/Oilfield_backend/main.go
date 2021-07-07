package main

import (
    "Oilfield_backend/Timer"
    "Oilfield_backend/config"
    "Oilfield_backend/routers"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/recover"
)


func Cors(ctx iris.Context){
    ctx.Header("Access-Control-Allow-Origin","*")
    if ctx.Method() == "OPTIONS" {
        ctx.Header("Access-Control-Allow-Methods","GET,POST,PUT,DELETE,PATCH,OPTIONS")
        ctx.Header("Access-Control-Allow-Headers","Content-Type,Accept,Authorization")
        ctx.StatusCode(204)
        return
    }
    ctx.Next()
}

func main() {
    app := iris.New()
    app.Use(recover.New())
    app.UseGlobal(Cors)
    common := app.Party("/")
    {
        common.Options("*", func(ctx iris.Context) {
            ctx.Next()
        })
    }
    // 定义接口路由
    app.PartyFunc("/BackStageManage", routers.BackStageManageRouters)
    app.PartyFunc("/User", routers.UserRoutes)
    app.PartyFunc("/HomePage",routers.HomePageRoutes)
    app.PartyFunc("/AlgorithmManage",routers.AlgorithmManageRoutes)
    app.PartyFunc("/Task",routers.TaskRoutes)
    app.PartyFunc("/Warning",routers.WarningRoutes)
    app.PartyFunc("/Equip",routers.EquipRoutes)

    // 开始循环检测协程
    go Timer.CircleCheck()
    // 开始监听
    app.Run(iris.Addr(config.BackendAddress + config.Port),iris.WithPostMaxMemory(config.FileSizeLimit))

}
