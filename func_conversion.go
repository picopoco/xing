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
	"strings"
)

func F바이트_변환값_해석(바이트_변환값 *lib.S바이트_변환) (해석값 interface{}, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 해석값 = nil }}.S실행()

	자료형_문자열 := 바이트_변환값.G자료형_문자열()
	시작_인덱스 := strings.Index(자료형_문자열, ".") + 1
	자료형_문자열 = 자료형_문자열[시작_인덱스:]

	switch 자료형_문자열 {
	case P자료형_S현물_주문_응답_실시간_정보:
		s := new(S현물_주문_응답_실시간_정보)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_정상_주문:
		s := new(S질의값_정상_주문)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_정정_주문:
		s := new(S질의값_정정_주문)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_취소_주문:
		s := new(S질의값_취소_주문)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_현물_전일당일_분틱_조회:
		s := new(S질의값_현물_전일당일_분틱_조회)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_현물_기간별_조회:
		s := new(S질의값_현물_기간별_조회)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_단일종목_연속키:
		s := new(S질의값_단일종목_연속키)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S질의값_증시주변자금추이:
		s := new(S질의값_증시주변자금추이)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S콜백_기본형:
		s := new(S콜백_기본형)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S콜백_정수값:
		s := New콜백_정수값_기본형()
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S콜백_문자열:
		s := new(S콜백_문자열)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S콜백_TR데이터:
		s := new(S콜백_TR데이터)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S콜백_메시지_및_에러:
		s := new(S콜백_메시지_및_에러)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_정상_주문_응답:
		s := new(S현물_정상_주문_응답)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_정상_주문_응답1:
		s := new(S현물_정상_주문_응답1)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_정상_주문_응답2:
		s := new(S현물_정상_주문_응답2)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_정정_주문_응답:
		s := new(S현물_정정_주문_응답)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_정정_주문_응답1:
		s := new(S현물_정정_주문_응답1)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_정정_주문_응답2:
		s := new(S현물_정정_주문_응답2)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_취소_주문_응답:
		s := new(S현물_취소_주문_응답)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_취소_주문_응답1:
		s := new(S현물_취소_주문_응답1)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_취소_주문_응답2:
		s := new(S현물_취소_주문_응답2)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_호가조회_응답:
		s := new(S현물_호가조회_응답)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_시세조회_응답:
		s := new(S현물_시세조회_응답)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_시간대별_체결_응답:
		s := new(S현물_시간대별_체결_응답)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_시간대별_체결_응답_헤더:
		s := new(S현물_시간대별_체결_응답_헤더)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_시간대별_체결_응답_반복값:
		s := new(S현물_시간대별_체결_응답_반복값)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_시간대별_체결_응답_반복값_모음:
		s := new(S현물_시간대별_체결_응답_반복값_모음)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_기간별_조회_응답:
		s := new(S현물_기간별_조회_응답)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_기간별_조회_응답_헤더:
		s := new(S현물_기간별_조회_응답_헤더)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_기간별_조회_응답_반복값:
		s := new(S현물_기간별_조회_응답_반복값)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_기간별_조회_응답_반복값_모음:
		s := new(S현물_기간별_조회_응답_반복값_모음)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_전일당일분틱조회_응답:
		s := new(S현물_전일당일분틱조회_응답)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_전일당일분틱조회_응답_헤더:
		s := new(S현물_전일당일분틱조회_응답_헤더)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_전일당일분틱조회_응답_반복값:
		s := new(S현물_전일당일분틱조회_응답_반복값)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S현물_전일당일분틱조회_응답_반복값_모음:
		s := new(S현물_전일당일분틱조회_응답_반복값_모음)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S_ETF_현재가_조회_응답:
		s := new(S_ETF_현재가_조회_응답)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S_ETF시간별_추이_응답:
		s := new(S_ETF시간별_추이_응답)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S_ETF시간별_추이_응답_헤더:
		s := new(S_ETF시간별_추이_응답_헤더)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S_ETF시간별_추이_응답_반복값:
		s := new(S_ETF시간별_추이_응답_반복값)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S_ETF시간별_추이_응답_반복값_모음:
		s := new(S_ETF시간별_추이_응답_반복값_모음)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S증시주변자금추이_응답:
		s := new(S증시주변자금추이_응답)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S증시주변자금추이_응답_헤더:
		s := new(S증시주변자금추이_응답_헤더)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S증시주변자금추이_응답_반복값:
		s := new(S증시주변자금추이_응답_반복값)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S증시주변자금추이_응답_반복값_모음:
		s := new(S증시주변자금추이_응답_반복값_모음)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S주식종목조회_응답_반복값:
		s := new(S현물_종목조회_응답_반복값)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	case P자료형_S주식종목조회_응답_반복값_모음:
		s := new(S현물_종목조회_응답_반복값_모음)
		lib.F에러체크(바이트_변환값.G값(s))
		return s, nil
	}

	return lib.F바이트_변환값_해석(바이트_변환값)
}
