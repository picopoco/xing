/* Copyright (C) 2015-2019 김운하(UnHa Kim)  unha.kim@kuh.pe.kr

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

Copyright (C) 2015-2019년 UnHa Kim (unha.kim@kuh.pe.kr)

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
	"math"
	"strings"
	"testing"
)

func TestF재무_순위_종합_t3341(t *testing.T) {
	lib.F메모("재무 순위 t3341 테스트 및 디버깅 할 것.")
	t.SkipNow()

	t.Parallel()

	접속됨, 에러 := F접속됨()
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, 접속됨)

	시장구분_모음 := []lib.T시장구분{lib.P시장구분_전체, lib.P시장구분_코스피, lib.P시장구분_코스닥}
	시장구분 := 시장구분_모음[lib.F임의_범위_이내_정수값(0,2)]

	재무순위_구분_모음 := []T재무순위_구분{P재무순위_매출액증가율,
		P재무순위_영업이익증가율, P재무순위_세전계속이익증가율, P재무순위_부채비율,
		P재무순위_유보율, P재무순위_EPS, P재무순위_BPS, P재무순위_ROE,
		P재무순위_PER, P재무순위_PBR, P재무순위_PEG}
	재무순위_구분 := 재무순위_구분_모음[lib.F임의_범위_이내_정수값(0,len(재무순위_구분_모음))]

	const 수량 = 100
	값_모음, 에러 := F재무_순위_종합_t3341(시장구분, 재무순위_구분, 수량)
	lib.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		lib.F테스트_참임(t, 값.M순위 > 0 && 값.M순위 <= 수량)
		lib.F테스트_참임(t, F종목코드_존재함(값.M종목코드))
		lib.F테스트_참임(t, strings.TrimSpace(값.M기업명) != "")
		lib.F테스트_참임(t, math.Abs(값.M매출액_증가율) < 1000)
		lib.F테스트_참임(t, math.Abs(값.M영업이익_증가율) < 1000)
		lib.F테스트_참임(t, math.Abs(값.M경상이익_증가율) < 1000)
		lib.F테스트_참임(t, 값.M부채비율 >= 0 && 값.M부채비율 < 3000)
		lib.F테스트_참임(t, 값.M유보율 >= 0 && 값.M유보율 < 1000)
		lib.F테스트_참임(t, 값.EPS * 값.PER >= 0)
		lib.F테스트_참임(t, 값.EPS * 값.ROE >= 0)
		lib.F테스트_참임(t, 값.BPS > 0)
		lib.F테스트_참임(t, math.Abs(값.ROE) < 100)
		lib.F테스트_참임(t, math.Abs(값.PER) < 100)
		lib.F테스트_참임(t, math.Abs(값.PBR) < 100)
		lib.F테스트_참임(t, math.Abs(값.PEG) < 10)
	}
}