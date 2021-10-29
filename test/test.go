package test

import "time"

func TestQuickStar(t *testing.T) {
	t1 := time.Now()
	t.Logf("ok, %v: %v; time: %v", res.Uid, res.Openid, time.Since(t1))
}
