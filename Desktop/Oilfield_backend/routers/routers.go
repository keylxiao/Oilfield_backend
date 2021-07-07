package routers

import (
    "Oilfield_backend/controllers"
    "github.com/kataras/iris/v12"
)

// BackStageManageRouters 后台管理路由组
func BackStageManageRouters(manage iris.Party) {
    // 获取主机监控信息
    manage.Get("/HostMonitor", controllers.GetHostMonitor)
    // 运行日志

}

// UserRoutes 用户路由组
func UserRoutes(user iris.Party) {
    // 用户登录轮询
    user.Get("/UserOnlineCheck",controllers.GetUserOnlineCheck)
    // 用户下线, 清空缓存
    user.Get("/UserOffline",controllers.GetUserOffline)
    // 查看登录情况
    user.Get("/UserLoginLog", controllers.GetUserLoginLog)
    // 查看用户信息
    user.Get("/UserInfo", controllers.GetUserInfo)
    // 查看所有用户信息
    user.Get("/AllUserInfo", controllers.GetAllUserInfo)
    // 用户注册
    user.Post("/UserRegister", controllers.PostUserRegister)
    // 用户登录
    user.Post("/UserLogin", controllers.PostUserLogin)
    // 修改用户权限
    user.Put("/UpdateUserStatus", controllers.PutUpdateUserStatus)
    // 用户信息修改
    user.Put("/UpdateUserInfo", controllers.PutUpdateUserInfo)
    // 删除用户
    user.Delete("/DeleteUser", controllers.DeleteUser)
}

// HomePageRoutes 首页路由
func HomePageRoutes(head iris.Party) {
    // 首页总览信息
    head.Get("/Overview", controllers.GetOverview)
}

// AlgorithmManageRoutes 算法管理路由
func AlgorithmManageRoutes(alg iris.Party) {
    // 创建一条算法
    alg.Post("/CreateAlgorithm", controllers.PostCreateAlgorithm)
    // 上传算法文件
    alg.Post("/CreateAlgorithmFile", controllers.PostCreateAlgorithmFile)
    // 查看所有算法
    alg.Get("/AllAlgorithm", controllers.GetAllAlgorithm)
    // 更新算法信息
    alg.Put("/UpdateAlgorithm", controllers.PutUpdateAlgorithm)
    // 删除算法
    alg.Delete("/DeleteAlgorithm", controllers.DeleteAlgorithm)
}

// TaskRoutes 任务路由
func TaskRoutes(check iris.Party) {
    // RTSP转流http-flv
    check.Get("/RtspToRtmp",controllers.GetRtspToRtmp)
    // 用户停止一个http-flv
    check.Get("/StopRtmp",controllers.GetStopRtmp)
    // 查看所有任务
    check.Get("/AllTask", controllers.GetAllTask)
    // 条件查看任务
    check.Get("/SomeTask", controllers.GetSomeTask)
    // 创建并添加根任务
    check.Post("/CreateRootTask", controllers.PostCreateRootTask)
    // 创建并添加子任务
    check.Post("/CreateSonTask",controllers.PostCreateSonTask)
    // 更新根任务
    check.Put("/UpdateRootTask", controllers.PutUpdateRootTask)
    // 更新子任务
    check.Put("/UpdateSonTask",controllers.PutUpdateSonTask)
    // 删除根任务
    check.Delete("/DeleteRootTask", controllers.DeleteRootTask)
    // 删除子任务
    check.Delete("/DeleteSonTask",controllers.DeleteSonTask)

}

// WarningRoutes 告警路由
func WarningRoutes(warn iris.Party) {
    // 查看所有告警
    warn.Get("/AllWarning", controllers.GetAllWarning)
    // 告警数据下载
    warn.Get("/DownloadFile", controllers.GetDownloadFile)
    // 告警视频流传输
    warn.Get("/FileStream", controllers.GetFileStream)
    // 获取当前告警配置
    warn.Get("Settings",controllers.GetSettings)
    // 创建并添加一条告警
    warn.Post("/CreateWarning", controllers.PostCreateWarning)
    // 条件查看告警
    warn.Post("/SomeWarning", controllers.PostSomeWarning)
    // 获取告警图片
    warn.Post("/WarningPhotos", controllers.PostWarningPhotos)
    // 告警数据缓存
    warn.Post("/CreateWarningFile", controllers.PostCreateWarningFile)
    // 更新告警设置
    warn.Put("/UpdateSettings", controllers.PutUpdateSettings)
}

// EquipRoutes 设备路由
func EquipRoutes(equip iris.Party) {
    // 添加Ipc设备
    equip.Post("/CreateEquipIpc", controllers.PostCreateEquipIpc)
    // 添加Nvr设备
    equip.Post("/CreateEquipNvr", controllers.PostCreateEquipNvr)
    // 查看所有设备
    equip.Get("/AllEquip", controllers.GetAllEquip)
    // 条件查看设备
    equip.Post("/SomeEquip", controllers.PostSomeEquip)
    // 更新设备
    equip.Put("/UpdateEquip", controllers.PutUpdateEquip)
    // 删除设备
    equip.Delete("/DeleteEquip", controllers.DeleteEquip)
}
