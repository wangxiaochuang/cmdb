package flowctrl

import "github.com/juju/ratelimit"

type RateLimiter interface {
    // TryAccept returns true if a token is taken immediately. Otherwise,
    // it returns false.
    TryAccept() bool

    // Accept will wait and not return unless a token becomes available.
    Accept()

    // QPS returns QPS of this rate limiter
    QPS() int64

    // Burst returns the burst of this rate limiter
    Burst() int64
}

func NewRateLimiter(qps, burst int64) RateLimiter {
    limiter := ratelimit.NewBucketWithRate(float64(qps), burst)
    return &tokenBucket{
        limiter: limiter,
        qps:     qps,
        burst:   burst,
    }
}

type tokenBucket struct {
    limiter *ratelimit.Bucket
    qps     int64
    burst   int64
}

func (t *tokenBucket) TryAccept() bool {
    return t.limiter.TakeAvailable(1) == 1
}

func (t *tokenBucket) Accept() {
    t.limiter.Wait(1)
}

func (t *tokenBucket) QPS() int64 {
    return t.qps
}

func (t *tokenBucket) Burst() int64 {
    return t.burst
}

func NewMockRateLimiter() RateLimiter {
    return &mockRatelimiter{}
}

type mockRatelimiter struct{}

func (*mockRatelimiter) TryAccept() bool {
    return true
}

func (*mockRatelimiter) Accept() {

}

func (*mockRatelimiter) QPS() int64 {
    return 0
}

func (*mockRatelimiter) Burst() int64 {
    return 0
}
