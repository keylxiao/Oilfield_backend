package Timer

import (
    "Oilfield_backend/config"
    "Oilfield_backend/models"
    "Oilfield_backend/utils"
    "context"
    "fmt"
    "net/http"
    "os/exec"
    "sync"
    "time"
)

// CircleCheck 循环检测任务
func CircleCheck() {
    equip := time.NewTicker(config.CycleEquipTimeDuration)
    for {
        select {
        case <-equip.C:
            go CheckOnline()
        }
    }
}

// CheckOnline 并发检查摄像头是否在线
func CheckOnline() {
    var wg sync.WaitGroup
    db := utils.DBOpen()
    var ipcs []models.Ipc
    var nvrs []models.Nvr
    db.Find(&ipcs)
    db.Find(&nvrs)

    for k := range ipcs {
        go ChangeOnline("ipc", ipcs[k].Ip, &wg)
        wg.Add(1)
    }
    for k := range nvrs {
        go ChangeOnline("nvr", nvrs[k].Ip, &wg)
        wg.Add(1)
    }

    wg.Wait()
}

// ChangeOnline 检测某一摄像头是否在线, 并改变在线状态
func ChangeOnline(equipType, ip string, wg *sync.WaitGroup) {
    timer := time.NewTicker(config.GetEquipTime)
    ch := make(chan int)
    url := fmt.Sprintf("http://%s", ip)
    go func(out chan int) {
        _, err := http.Get(url)
        if err == nil {
            close(out)
        }
    }(ch)
    for {
        select {
        case <-timer.C:
            models.UpdateOnline(equipType, ip, 1)
            wg.Done()
            return
        case <-ch:
            models.UpdateOnline(equipType, ip, 0)
            wg.Done()
            return
        }
    }
}

func RtspToRtmp(ip, rtsp, rtmp string, ch1 chan int,channel chan string) {
    ctx, cancel := context.WithCancel(context.Background())
    timer := time.NewTicker(config.RtmpTime)
    red := utils.RedisOpen()
    cmd := exec.CommandContext(ctx, "ffmpeg", "-i", rtsp, "-vcodec", "copy", "-acodec", "copy", "-f", "flv", rtmp)
    err := cmd.Start()
    if err != nil {
        close(ch1)
    }

    for{
        select {
        case <-timer.C:
            red.Set(ip,0,0)
            goto end
        case v := <- channel :
            // 待销毁地址是该协程管理的
            if v == rtmp {
                a, _ := red.Get(ip).Int()
                // 如果只有这一个
                if a == 1 {
                    red.Set(ip,0,0)
                    goto end
                } else {
                    // 不止这一个
                    red.Set(ip, a-1, 0)
                }
            } else {
                // 把拿出的数据再传回去
                channel <- v
            }
        }
    }
end:
    cancel()
}