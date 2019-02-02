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
	"testing"
	"time"
)

func TestF현물_차트_일주월_t8413(t *testing.T) {
	t.Parallel()

	접속됨, 에러 := F접속됨()
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, 접속됨)

	const 종목코드 = "069500" // 코덱스200
	var 이전_일자 time.Time

	시작일 := 당일.G값().AddDate(-1, 0, 0)

	값_모음, 에러 := F현물_차트_일주월_t8413(종목코드, 시작일, 당일.G값(), P일주월_일, 2300)
	lib.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		lib.F테스트_참임(t, 값.M일자.After(이전_일자) || 값.M일자.Equal(이전_일자))
		lib.F테스트_참임(t, 값.M일자.Equal(시작일) || 값.M일자.After(시작일), 값.M일자, 시작일)
		이전_일자 = 값.M일자

		F테스트_현물_차트_일주월_응답_반복값_t8413(t, 값, 종목코드)
	}
}

