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
	"github.com/ghts/xing_types"

	"time"
)

// 가장 간단한 질의. 접속 유지 및 질의 기능 테스트 용도로 적합함.
func F시각_조회_t0167() (시각 time.Time, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 시각 = time.Time{} }}.S실행_No출력()

	질의값 := lib.S질의값_기본형{TR구분: xt.TR조회, TR코드: xt.TR시간_조회}
	시각 = 에러체크(F질의_단일TR(질의값, lib.P20초)).(time.Time)

	return 시각, nil
}

func F현물_정상주문_CSPAT00600(질의값 *xt.S질의값_정상주문) (응답값 *xt.S현물_정상주문_응답, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 응답값 = nil }}.S실행()

	응답값 = 에러체크(F질의_단일TR(질의값)).(*xt.S현물_정상주문_응답)

	return 응답값, nil
}

func F현물_정정주문_CSPAT00700(질의값 *xt.S질의값_정정주문) (응답값 *xt.S현물_정정주문_응답, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 응답값 = nil }}.S실행()

	응답값 = 에러체크(F질의_단일TR(질의값)).(*xt.S현물_정정주문_응답)

	return 응답값, nil
}

func F현물_취소주문_CSPAT00800(질의값 *xt.S질의값_취소주문) (응답값 *xt.S현물_취소주문_응답, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 응답값 = nil }}.S실행()

	응답값 = 에러체크(F질의_단일TR(질의값)).(*xt.S현물_취소주문_응답)

	return 응답값, nil
}

func F현물_호가_조회_t1101(종목코드 string) (응답값 *xt.S현물_호가조회_응답, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 응답값 = nil }}.S실행()

	질의값 := lib.New질의값_단일종목()
	질의값.TR구분 = xt.TR조회
	질의값.TR코드 = xt.TR현물_호가_조회
	질의값.M종목코드 = 종목코드

	응답값 = 에러체크(F질의_단일TR(질의값)).(*xt.S현물_호가조회_응답)

	return 응답값, nil
}

func F현물_시세_조회_t1102(종목코드 string) (응답값 *xt.S현물_시세조회_응답, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 응답값 = nil }}.S실행()

	질의값 := lib.New질의값_단일종목()
	질의값.TR구분 = xt.TR조회
	질의값.TR코드 = xt.TR현물_시세_조회
	질의값.M종목코드 = 종목코드

	응답값 = 에러체크(F질의_단일TR(질의값)).(*xt.S현물_시세조회_응답)

	return 응답값, nil
}

func F기간별_주가_조회_t1305(종목코드 string, 일주월_구분 xt.T일주월_구분) (
	응답값_모음 []*xt.S현물_기간별_조회_응답_반복값, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F조건부_패닉(일주월_구분 != xt.P일주월_일 && 일주월_구분 != xt.P일주월_주 &&
		일주월_구분 != xt.P일주월_월, "예상하지 못한 일주월 구분값 : '%v'", 일주월_구분)

	연속키 := ""
	응답값_모음 = make([]*xt.S현물_기간별_조회_응답_반복값, 0)

	lib.F메모("TR응답에 연속 조회 추가 존재 여부를 포함시켜서 반복 여부 판단 조건으로 사용하는 것을 생각해 볼 것.")

	for {
		질의값 := xt.New질의값_현물_기간별_조회()
		질의값.TR구분 = xt.TR조회
		질의값.TR코드 = xt.TR현물_기간별_조회
		질의값.M종목코드 = 종목코드
		질의값.M일주월_구분 = 일주월_구분
		질의값.M수량 = 200
		질의값.M연속키 = 연속키

		응답값 := 에러체크(F질의_단일TR(질의값)).(*xt.S현물_기간별_조회_응답)
		연속키 = 응답값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 응답값.M반복값_모음.M배열...)

		lib.F조건부_패닉(응답값.M헤더.M수량 != int64(len(응답값.M반복값_모음.M배열)),
			"반복값 수량 불일치. '%v', '%v'", 응답값.M헤더.M수량, len(응답값.M반복값_모음.M배열))

		if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func F현물_당일전일_분틱_조회_t1310(종목코드 string, 당일전일_구분 xt.T당일전일_구분, 분틱_구분 xt.T분틱_구분,
	종료시각 time.Time) (응답값_모음 []*xt.S현물_전일당일분틱조회_응답_반복값, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F조건부_패닉(당일전일_구분 != xt.P당일전일구분_당일 && 당일전일_구분 != xt.P당일전일구분_전일,
		"예상하지 못한 당일_전일 구분값 : '%v'", 당일전일_구분)

	lib.F조건부_패닉(분틱_구분 != xt.P분틱구분_분 && 분틱_구분 != xt.P분틱구분_틱,
		"예상하지 못한 분_틱 구분값 : '%v'", 분틱_구분)

	응답값_모음 = make([]*xt.S현물_전일당일분틱조회_응답_반복값, 0)
	연속키 := ""

	lib.F메모("TR응답에 연속 조회 추가 존재 여부를 포함시켜서 반복 여부 판단 조건으로 사용하는 것을 생각해 볼 것.")

	for {
		질의값 := xt.New질의값_현물_전일당일_분틱_조회()
		질의값.TR구분 = xt.TR조회
		질의값.TR코드 = xt.TR현물_당일_전일_분틱_조회
		질의값.M종목코드 = 종목코드
		질의값.M당일전일구분 = 당일전일_구분
		질의값.M분틱구분 = 분틱_구분
		질의값.M종료시각 = 종료시각
		질의값.M연속키 = 연속키

		응답값 := 에러체크(F질의_단일TR(질의값)).(*xt.S현물_전일당일분틱조회_응답)
		연속키 = 응답값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 응답값.M반복값_모음.M배열...)

		if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func ETF_시간별_추이_t1902(종목코드 string) (응답값_모음 []*xt.S_ETF시간별_추이_응답_반복값, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	응답값_모음 = make([]*xt.S_ETF시간별_추이_응답_반복값, 0)
	연속키 := ""

	lib.F메모("TR응답에 연속 조회 추가 존재 여부를 포함시켜서 반복 여부 판단 조건으로 사용하는 것을 생각해 볼 것.")

	for {
		질의값 := xt.New질의값_단일종목_연속키()
		질의값.TR구분 = xt.TR조회
		질의값.TR코드 = xt.TR_ETF_시간별_추이
		질의값.M종목코드 = 종목코드
		질의값.M연속키 = 연속키

		응답값 := 에러체크(F질의_단일TR(질의값)).(*xt.S_ETF시간별_추이_응답)
		연속키 = 응답값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 응답값.M반복값_모음.M배열...)

		if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func F증시주변자금추이_t8428(시장_구분 lib.T시장구분) (응답값_모음 []*xt.S증시주변자금추이_응답_반복값, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F조건부_패닉(
		시장_구분 != lib.P시장구분_코스피 && 시장_구분 != lib.P시장구분_코스닥,
		"예상하지 못한 시장 구분값 : '%v'", 시장_구분)

	연속키 := ""
	응답값_모음 = make([]*xt.S증시주변자금추이_응답_반복값, 0)

	lib.F메모("TR응답에 연속 조회 추가 존재 여부를 포함시켜서 반복 여부 판단 조건으로 사용하는 것을 생각해 볼 것.")

	for {
		질의값 := xt.New질의값_증시주변자금추이()
		질의값.TR구분 = xt.TR조회
		질의값.TR코드 = xt.TR증시_주변_자금_추이
		질의값.M시장구분 = 시장_구분
		질의값.M수량 = 200
		질의값.M연속키 = 연속키

		응답값 := 에러체크(F질의_단일TR(질의값)).(*xt.S증시주변자금추이_응답)
		연속키 = 응답값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 응답값.M반복값_모음.M배열...)

		if len(lib.F정규식_검색(연속키, []string{"[0-9]*"})) < 8 {
			break
		}
	}

	return 응답값_모음, nil
}

func F주식종목조회_t8436(시장_구분 lib.T시장구분) (응답값_모음 []*xt.S현물_종목조회_응답_반복값, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	var 시장구분_문자열 string

	switch 시장_구분 {
	case lib.P시장구분_전체:
		시장구분_문자열 = "0"
	case lib.P시장구분_코스피:
		시장구분_문자열 = "1"
	case lib.P시장구분_코스닥:
		시장구분_문자열 = "2"
	default:
		panic(lib.New에러("예상하지 못한 시장 구분값 : '%v'", 시장_구분))
	}

	질의값 := lib.New질의값_문자열(xt.TR조회, xt.TR현물_종목_조회, 시장구분_문자열)
	응답값_모음 = 에러체크(F질의_단일TR(질의값)).(*xt.S현물_종목조회_응답_반복값_모음).M배열

	return 응답값_모음, nil
}
