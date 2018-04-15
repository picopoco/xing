/* Copyright (C) 2015-2018 김운하(UnHa Kim)  unha.kim@kuh.pe.kr

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2018년 UnHa Kim (unha.kim@kuh.pe.kr)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package xing

import (
	"github.com/ghts/lib"

	"testing"
	"time"
)

func TestGo루틴_실시간_정보_중계(t *testing.T) {
	if !lib.F한국증시_정규시장_거래시간임() {
		t.SkipNow()
	}

	ch실시간_데이터 := make(chan lib.I소켓_메시지, 100)

	종목_모음, 에러 := lib.F종목모음_코스피()
	lib.F테스트_에러없음(t, 에러)

	종목코드_모음 := lib.F종목코드_추출(종목_모음, 100)

	F실시간_데이터_구독_ETF(ch실시간_데이터, 종목코드_모음)
	defer F실시간_데이터_해지_ETF(종목코드_모음)

	ch초기화 := make(chan lib.T신호)
	go Go루틴_실시간_정보_중계(ch초기화)
	lib.F테스트_같음(t, <-ch초기화, lib.P신호_초기화)

	ch대기시간_초과 := time.After(lib.P1분 * 2)

	var 테스트_수량 int
	if lib.F한국증시_정규시장_거래시간임() {
		테스트_수량 = 10
	} else {
		테스트_수량 = 1
	}

	수신_수량 := 0

	for {
		select {
		case <-ch실시간_데이터:
			수신_수량++
			if 수신_수량 > 테스트_수량 {
				return
			}
		case <-ch대기시간_초과:
			lib.F문자열_출력("대기시간 초과. '%v'", 수신_수량)
			t.FailNow()
		}
	}
}
