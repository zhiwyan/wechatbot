package utils

import (
	"fmt"
	"testing"
)

// GetPddUri的单元测试
func TestGetPddUri(t *testing.T) {
	link := "https://mobile.yangkeduo.com/goods2.html?_wv=41729&_wvx=10&_x_share_id=KYuDsIneU5CkYWoSPWQBUSswZpdMLwwk&goods_id=108459872240&page_from=26&_oc_trace_mark=199&pxq_secret_key=CP2PEEM6RUWMJCVDCUGGFAHZNUBBSHWPI6FJXYLU6IVCPTB7PJ4Q&_oak_share_snapshot_num=880&_oak_share_detail_id=4214407887&_oak_share_time=1715093704&refer_share_id=ve6k0zh2k1bw81sjv2k4rvu99a72b59l&refer_share_uin=TLU6TMJLOMU7OMB6KVILMM36II_GEXDA&refer_share_channel=message"
	fmt.Println(GetPddUri(link))
}
