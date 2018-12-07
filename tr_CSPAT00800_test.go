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

func TestCSPAT00800현물_취소_주문(t *testing.T) {
	if !F한국증시_정규시장_거래시간임() {
		t.SkipNow()
	}

	접속됨, 에러 := F접속됨()
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, 접속됨)

	lib.F테스트_에러없음(t, F주문_응답_실시간_정보_구독())

	var 종목 = lib.New종목("069500", "KODEX 200", lib.P시장구분_ETF)

	하한가, ok := 하한가_맵[종목.G코드()]
	lib.F테스트_참임(t, ok, "하한가를 찾을 수 없음. %v", 종목.G코드())

	const 수량_정상주문 = int64(25)
	const 반복횟수 = int(수량_정상주문) - 5
	const 수량_전량_취소_주문 = 0
	const 수량_일부_취소_주문 = int64(1)
	var 가격_정상주문 = 하한가

	계좌번호_모음, 에러 := F계좌번호_모음()
	lib.F테스트_에러없음(t, 에러)
	계좌번호 := 계좌번호_모음[0]

	질의값 := New질의값_정상_주문()
	질의값.M구분 = TR주문
	질의값.M코드 = TR현물_정상_주문
	질의값.M계좌번호 = 계좌번호
	질의값.M계좌_비밀번호 = "0000" // 모의투자에서는 계좌 비밀번호를 체크하지 않음.
	질의값.M종목코드 = 종목.G코드()
	질의값.M주문수량 = 수량_정상주문
	질의값.M주문단가 = 가격_정상주문 // 시장가 주문 시 가격은 무조건 '0'을 입력해야 함.
	질의값.M매수_매도 = lib.P매수
	질의값.M호가유형 = lib.P호가유형_지정가
	질의값.M신용거래_구분 = lib.P신용거래_해당없음
	질의값.M주문조건 = lib.P주문조건_없음 // 모의투자에서는 IOC, FOK를 사용할 수 없음.
	질의값.M대출일 = time.Time{}   // 신용주문이 아닐 경우는 NewCSPAT00600InBlock1()에서 공백문자로 바꿔줌.

	정상주문_응답값, 에러 := F현물_정상주문_CSPAT00600(질의값)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, 정상주문_응답값.M응답2.M주문번호 > 0, 정상주문_응답값.M응답2.M주문번호)

	질의값_취소주문 := New질의값_취소_주문()
	질의값_취소주문.M구분 = TR주문
	질의값_취소주문.M코드 = TR현물_취소_주문
	질의값_취소주문.M원주문번호 = 정상주문_응답값.M응답2.M주문번호
	질의값_취소주문.M계좌번호 = 계좌번호
	질의값_취소주문.M계좌_비밀번호 = "0000" // 모의투자에서는 계좌 비밀번호는 체크 안 함.
	질의값_취소주문.M종목코드 = 종목.G코드()
	질의값_취소주문.M주문수량 = 수량_일부_취소_주문

	// 일부 취소주문 TR 실행
	for i := 0; i < 반복횟수 ; i++ {
		취소주문_응답값, 에러 := F현물_취소주문_CSPAT00800(질의값_취소주문)
		lib.F테스트_에러없음(t, 에러)
		lib.F테스트_참임(t, 취소주문_응답값.M응답2.M주문번호 > 0, 취소주문_응답값.M응답2.M주문번호)
	}

	// 전량 취소주문 TR 실행
	질의값_취소주문 = New질의값_취소_주문()
	질의값_취소주문.M구분 = TR주문
	질의값_취소주문.M코드 = TR현물_취소_주문
	질의값_취소주문.M원주문번호 = 정상주문_응답값.M응답2.M주문번호
	질의값_취소주문.M계좌번호 = 계좌번호
	질의값_취소주문.M계좌_비밀번호 = "" // 모의투자에서는 계좌 비밀번호는 체크 안 함.
	질의값_취소주문.M종목코드 = 종목.G코드()
	질의값_취소주문.M주문수량 = 수량_전량_취소_주문

	취소주문_응답값, 에러 := F현물_취소주문_CSPAT00800(질의값_취소주문)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, 취소주문_응답값.M응답2.M주문번호 > 0, 취소주문_응답값.M응답2.M주문번호)

	lib.F테스트_에러없음(t, F주문_응답_실시간_정보_해지())
}
