package configcenter

import (
    "context"
    "fmt"
    "sync"

    "github.com/wxc/cmdb/common/blog"
    crd "github.com/wxc/cmdb/common/confregdiscover"
    "github.com/wxc/cmdb/common/errors"
    "github.com/wxc/cmdb/common/language"
    "github.com/wxc/cmdb/common/types"
)

var confC *CC

func NewConfigCenter(ctx context.Context, disc crd.ConfRegDiscvIf, confPath string, handler *CCHandler) error {
    return New(ctx, confPath, disc, handler)
}

func New(ctx context.Context, confPath string, disc crd.ConfRegDiscvIf, handler *CCHandler) error {
    confC = &CC{
        ctx:           ctx,
        disc:          disc,
        handler:       handler,
        previousProc:  new(ProcessConfig),
        previousLang:  make(map[string]language.LanguageMap),
        previousError: make(map[string]errors.ErrorCode),
    }

    // parse config only from file
    if len(confPath) != 0 {
        return LoadConfigFromLocalFile(confPath, handler)
    }

    if err := confC.run(); err != nil {
        return err
    }

    confC.sync()

    return nil
}

type ProcHandlerFunc func(previous, current ProcessConfig)

type CCHandler struct {
    OnProcessUpdate  ProcHandlerFunc
    OnExtraUpdate    ProcHandlerFunc
    OnLanguageUpdate func(previous, current map[string]language.LanguageMap)
    OnErrorUpdate    func(previous, current map[string]errors.ErrorCode)
    OnMongodbUpdate  func(previous, current ProcessConfig)
    OnRedisUpdate    func(previous, current ProcessConfig)
}

type CC struct {
    sync.Mutex
    // used to stop the config center gracefully.
    ctx             context.Context
    disc            crd.ConfRegDiscvIf
    handler         *CCHandler
    procName        string
    previousProc    *ProcessConfig
    previousExtra   *ProcessConfig
    previousMongodb *ProcessConfig
    previousRedis   *ProcessConfig
    previousLang    map[string]language.LanguageMap
    previousError   map[string]errors.ErrorCode
}

func (c *CC) run() error {
    commonConfPath := fmt.Sprintf("%s/%s", types.CC_SERVCONF_BASEPATH, types.CCConfigureCommon)
    commonConfEvent, err := c.disc.Discover(commonConfPath)
    if err != nil {
        return err
    }

    extraConfPath := fmt.Sprintf("%s/%s", types.CC_SERVCONF_BASEPATH, types.CCConfigureExtra)
    extraConfEvent, err := c.disc.Discover(extraConfPath)
    if err != nil {
        return err
    }

    mongodbConfPath := fmt.Sprintf("%s/%s", types.CC_SERVCONF_BASEPATH, types.CCConfigureMongo)
    mongodbConfEvent, err := c.disc.Discover(mongodbConfPath)
    if err != nil {
        return err
    }

    redisConfPath := fmt.Sprintf("%s/%s", types.CC_SERVCONF_BASEPATH, types.CCConfigureRedis)
    redisConfEvent, err := c.disc.Discover(redisConfPath)
    if err != nil {
        return err
    }

    langEvent, err := c.disc.Discover(types.CC_SERVLANG_BASEPATH)
    if err != nil {
        return err
    }

    errEvent, err := c.disc.Discover(types.CC_SERVERROR_BASEPATH)
    if err != nil {
        return err
    }

    go func() {
        for {
            select {
            case pEvent := <-commonConfEvent:
                c.onProcChange(pEvent)
            case pEvent := <-extraConfEvent:
                c.onExtraChange(pEvent)
            case pEvent := <-mongodbConfEvent:
                c.onMongodbChange(pEvent)
            case pEvent := <-redisConfEvent:
                c.onRedisChange(pEvent)
            case eEvent := <-errEvent:
                c.onErrorChange(eEvent)
            case langEvent := <-langEvent:
                c.onLanguageChange(langEvent)
            case <-c.ctx.Done():
                blog.Warnf("config center event watch stopped because of context done.")
                return
            }
        }
    }()
    return nil
}

func (c *CC) onProcChange(cur *crd.DiscoverEvent) {
    if cur.Err != nil {
        blog.Errorf("config center received event that %s config has changed, but got err: %v", types.CCConfigureCommon, cur.Err)
        return
    }

    //now := ParseConfigWithData(cur.Data)
    //c.Lock()
    //defer c.Unlock()
    //prev := c.previousProc
    //c.previousProc = now
    //if err := SetCommonFromByte(now.ConfigData); err != nil {
    //    blog.Errorf("add updated configuration error: %v", err)
    //    return
    //}
    //if c.handler != nil {
    //    if c.handler.OnProcessUpdate != nil {
    //        go c.handler.OnProcessUpdate(*prev, *now)
    //    }
    //}
}

func (c *CC) onExtraChange(cur *crd.DiscoverEvent) {
    if cur.Err != nil {
        blog.Errorf("config center received event that %s config has changed, but got err: %v", types.CCConfigureExtra, cur.Err)
        return
    }
}

func (c *CC) onMongodbChange(cur *crd.DiscoverEvent) {
    if cur.Err != nil {
        blog.Errorf("config center received event that %s config has changed, but got err: %v", types.CCConfigureCommon, cur.Err)
        return
    }
}

func (c *CC) onRedisChange(cur *crd.DiscoverEvent) {
    if cur.Err != nil {
        blog.Errorf("config center received event that %s config has changed, but got err: %v", types.CCConfigureCommon, cur.Err)
        return
    }
}

func (c *CC) onErrorChange(cur *crd.DiscoverEvent) {
    if cur.Err != nil {
        blog.Errorf("config center received event that *ERROR CODE* config has changed, but got err: %v", cur.Err)
        return
    }
}

func (c *CC) onLanguageChange(cur *crd.DiscoverEvent) {
    if cur.Err != nil {
        blog.Errorf("config center received event that *LANGUAGE* config has changed, but got err: %v", cur.Err)
        return
    }
    fmt.Printf("%v", string(cur.Data))
}

func (c *CC) sync() {
    panic("in sync")
}

func (c *CC) syncProc() {
    panic("in syncProc")
}

func (c *CC) syncExtra() {
    panic("in syncExtra")
}

func (c *CC) syncMongodb() {
    panic("in syncMongodb")
}

func (c *CC) syncRedis() {
    panic("in syncRedis")
}

func (c *CC) syncLang() {
    panic("in syncLang")
}

func (c *CC) syncErr() {
    panic("in syncErr")
}
