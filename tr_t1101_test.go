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

func TestF현물_호가_조회_t1101(t *testing.T) {
	t.Parallel()

	접속됨, 에러 := F접속됨()
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, 접속됨)

	종목코드 := "069500" // 코덱스200

	값, 에러 := F현물_호가_조회_t1101(종목코드)
	lib.F테스트_에러없음(t, 에러)

	lib.F테스트_다름(t, 값.M한글명, "")
	lib.F테스트_참임(t, 값.M현재가 >= 0)
	lib.F테스트_같음(t, 값.M전일대비구분, P구분_상한, P구분_상승, P구분_보합, P구분_하한, P구분_하락)
	lib.F테스트_참임(t, 값.M전일대비등락폭 >= 0)

	switch 값.M전일대비구분 { // 등락율 확인
	case P구분_상한, P구분_상승:
		lib.F테스트_참임(t, 값.M등락율 >= 0)
	case P구분_하한, P구분_하락:
		lib.F테스트_참임(t, 값.M등락율 <= 0)
	case P구분_보합:
		lib.F테스트_같음(t, 값.M등락율, 0)
	}

	lib.F테스트_참임(t, 값.M누적거래량 >= 0)
	lib.F테스트_참임(t, 값.M전일종가 >= 0)
	lib.F테스트_같음(t, len(값.M매도호가_모음), 10)
	lib.F테스트_같음(t, len(값.M매수호가_모음), 10)
	lib.F테스트_같음(t, len(값.M매도호가수량_모음), 10)
	lib.F테스트_같음(t, len(값.M매수호가수량_모음), 10)
	lib.F테스트_같음(t, len(값.M직전매도대비수량_모음), 10)
	lib.F테스트_같음(t, len(값.M직전매수대비수량_모음), 10)

	var 매도호가수량합, 매수호가수량합 int64
	// 게시판 문의 답변 내용 : (직전매도, 직전매수) 대비수량합과 1~10의 합계는 일치하지 않습니다.

	for i := 0; i < 10; i++ {
		lib.F테스트_참임(t, 값.M매도호가_모음[i] >= 0, 값.M매도호가_모음[i])
		lib.F테스트_참임(t, 값.M매수호가_모음[i] >= 0, 값.M매수호가_모음[i])
		lib.F테스트_참임(t, 값.M매도호가수량_모음[i] >= 0, 값.M매도호가수량_모음[i])
		lib.F테스트_참임(t, 값.M매수호가수량_모음[i] >= 0, 값.M매수호가수량_모음[i])

		// (-) 값이 나오더라..
		//lib.F테스트_참임(t, 값.M직전매도대비수량_모음[i] >= 0, 값.M직전매도대비수량_모음[i])
		//lib.F테스트_참임(t, 값.M직전매수대비수량_모음[i] >= 0, 값.M직전매수대비수량_모음[i])

		매도호가수량합 = 매도호가수량합 + 값.M매도호가수량_모음[i]
		매수호가수량합 = 매수호가수량합 + 값.M매수호가수량_모음[i]
	}

	if lib.F한국증시_정규시장_거래시간임() {
		lib.F테스트_참임(t, 값.M수신시간.After(time.Now().Add(-1*lib.P10분)),
			값.M수신시간.Format("15:04:06"), time.Now().Add(-1 * lib.P10분).Format("15:04:06"))
		lib.F테스트_참임(t, 값.M수신시간.Before(time.Now().Add(lib.P10분)),
			time.Now().Add(lib.P10분).Format("15:04:06"), 값.M수신시간.Format("15:04:06"))
	}

	if 값.M예상체결가격 != 0 {
		lib.F테스트_참임(t, float64(값.M예상체결가격) >= float64(값.M현재가)*0.7)
		lib.F테스트_참임(t, float64(값.M예상체결가격) <= float64(값.M현재가)*1.3)
		lib.F테스트_참임(t, 값.M예상체결수량 >= 0)
	}

	lib.F테스트_같음(t, 값.M예상체결전일구분, P구분_상한, P구분_상승, P구분_보합, P구분_하한, P구분_하락)
	lib.F테스트_참임(t, 값.M예상체결전일대비 >= 0)

	switch 값.M예상체결전일구분 { // 예상 체결 등락율 확인
	case P구분_상한, P구분_상승:
		lib.F테스트_참임(t, 값.M예상체결등락율 >= 0)
	case P구분_하한, P구분_하락:
		lib.F테스트_참임(t, 값.M예상체결등락율 <= 0)
	case P구분_보합:
		lib.F테스트_같음(t, 값.M예상체결등락율, 0)
	}

	lib.F테스트_참임(t, 값.M시간외매도잔량 >= 0)
	lib.F테스트_참임(t, 값.M시간외매수잔량 >= 0)
	lib.F테스트_같음(t, 값.M동시호가_구분, P동시호가_아님, P동시호가_장중, P동시호가_시간외, P동시호가_동시)
	lib.F테스트_같음(t, len(값.M종목코드), 6)
	lib.F테스트_참임(t, 값.M상한가 >= 값.M현재가)
	lib.F테스트_참임(t, 값.M상한가 >= 값.M하한가)
	lib.F테스트_참임(t, 값.M상한가 >= 값.M시가)
	lib.F테스트_참임(t, 값.M상한가 >= 값.M고가)
	lib.F테스트_참임(t, 값.M상한가 >= 값.M저가)
	lib.F테스트_참임(t, 값.M하한가 <= 값.M현재가)
	lib.F테스트_참임(t, 값.M하한가 <= 값.M시가 || 값.M시가 == 0)
	lib.F테스트_참임(t, 값.M하한가 <= 값.M고가 || 값.M고가 == 0)
	lib.F테스트_참임(t, 값.M하한가 <= 값.M저가 || 값.M저가 == 0)
	lib.F테스트_참임(t, 값.M고가 >= 값.M현재가 || 값.M고가 == 0)
	lib.F테스트_참임(t, 값.M고가 >= 값.M시가 || 값.M고가 == 0)
	lib.F테스트_참임(t, 값.M고가 >= 값.M저가 || 값.M고가 == 0)
	lib.F테스트_참임(t, 값.M저가 <= 값.M현재가)
}
