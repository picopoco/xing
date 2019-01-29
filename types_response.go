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
)

type I이중_응답 interface {
	I이중_응답1
	I이중_응답2
}

type I이중_응답1 interface {
	G응답1() I이중_응답1
}

type I이중_응답2 interface {
	G응답2() I이중_응답2
}

type S이중_응답_일반형 struct {
	M응답1 I이중_응답1
	M응답2 I이중_응답2
}

func (s *S이중_응답_일반형) G응답1() I이중_응답1 { return s.M응답1 }
func (s *S이중_응답_일반형) G응답2() I이중_응답2 { return s.M응답2 }

func (s *S이중_응답_일반형) G값(TR코드 string) interface{} {
	switch TR코드 {
	case TR현물_정상_주문:
		g := new(S현물_정상_주문_응답)
		g.M응답1 = s.M응답1.(*S현물_정상_주문_응답1)
		g.M응답2 = s.M응답2.(*S현물_정상_주문_응답2)
		return g
	case TR현물_정정_주문:
		g := new(S현물_정정_주문_응답)
		g.M응답1 = s.M응답1.(*S현물_정정_주문_응답1)
		g.M응답2 = s.M응답2.(*S현물_정정_주문_응답2)
		return g
	case TR현물_취소_주문:
		g := new(S현물_취소_주문_응답)
		g.M응답1 = s.M응답1.(*S현물_취소_주문_응답1)
		g.M응답2 = s.M응답2.(*S현물_취소_주문_응답2)
		return g
	case TR기업정보_요약:
		g := new(S기업정보_요약_응답)
		g.M응답1 = s.M응답1.(*S기업정보_요약_응답1)
		g.M응답2 = s.M응답2.(*S기업정보_요약_응답2)
		return g
	default:
		panic(lib.New에러("예상하지 못한 TR코드 : '%v'", TR코드))
	}
}

type I헤더_반복값_TR데이터 interface {
	I헤더_TR데이터
	I반복값_모음_TR데이터
}

type I헤더_TR데이터 interface {
	G헤더_TR데이터() I헤더_TR데이터
}

type I반복값_모음_TR데이터 interface {
	G반복값_모음_TR데이터() I반복값_모음_TR데이터
}

type S헤더_반복값_일반형 struct {
	M헤더     I헤더_TR데이터
	M반복값_모음 I반복값_모음_TR데이터
}

func (s *S헤더_반복값_일반형) G헤더_TR데이터() I헤더_TR데이터 { return s.M헤더 }
func (s *S헤더_반복값_일반형) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

func (s *S헤더_반복값_일반형) G값(TR코드 string) interface{} {
	switch TR코드 {
	default:
		panic(lib.New에러with출력("예상하지 못한 TR코드 : '%v'", TR코드)) // 패닉 출력을 삭제하지 말 것.
	//case TR현물_시간대별_체결_조회:
	//	값 := new(S현물_시간대별_체결_응답)
	//	값.M헤더 = s.M헤더.(*S현물_시간대별_체결_응답_헤더)
	//	값.M반복값_모음 = s.M반복값_모음.(*S현물_시간대별_체결_응답_반복값_모음)
	//	return 값
	case TR현물_기간별_조회:
		값 := new(S현물_기간별_조회_응답)
		값.M헤더 = s.M헤더.(*S현물_기간별_조회_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*S현물_기간별_조회_응답_반복값_모음)
		return 값
	case TR현물_당일_전일_분틱_조회:
		값 := new(S현물_전일당일분틱조회_응답)
		값.M헤더 = s.M헤더.(*S현물_전일당일분틱조회_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*S현물_전일당일분틱조회_응답_반복값_모음)
		return 값
	case TR_ETF_시간별_추이:
		값 := new(S_ETF시간별_추이_응답)
		값.M헤더 = s.M헤더.(*S_ETF시간별_추이_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*S_ETF시간별_추이_응답_반복값_모음)
		return 값
	case TR현물_차트_틱:
		값 := new(S현물_차트_틱_응답)
		값.M헤더 = s.M헤더.(*S현물_차트_틱_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*S현물_차트_틱_응답_반복값_모음)
		return 값
	case TR현물_차트_분:
		값 := new(S현물_차트_분_응답)
		값.M헤더 = s.M헤더.(*S현물_차트_분_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*S현물_차트_분_응답_반복값_모음)
		return 값
	case TR현물_차트_일주월:
		값 := new(S현물_차트_일주월_응답)
		값.M헤더 = s.M헤더.(*S현물_차트_일주월_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*S현물_차트_일주월_응답_반복값_모음)
		return 값
	case TR증시_주변_자금_추이:
		값 := new(S증시_주변자금추이_응답)
		값.M헤더 = s.M헤더.(*S증시_주변자금추이_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*S증시_주변자금추이_응답_반복값_모음)
		return 값
	}
}
