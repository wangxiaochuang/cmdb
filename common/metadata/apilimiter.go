package metadata

type LimiterRule struct {
    RuleName string `json:"rulename"`
    AppCode  string `json:"appcode"`
    User     string `json:"user"`
    IP       string `json:"ip"`
    Method   string `json:"method"`
    Url      string `json:"url"`
    Limit    int64  `json:"limit"`
    TTL      int64  `json:"ttl"`
    DenyAll  bool   `json:"denyall"`
}
