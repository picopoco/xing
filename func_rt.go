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
)

func F실시간_정보_구독_단순TR(RT코드 string) (에러 error) {
	return F질의(lib.New질의값_기본형(lib.TR실시간_정보_구독, RT코드)).G에러()
}

func F실시간_정보_해지_단순TR(RT코드 string) (에러 error) {
	return F질의(lib.New질의값_기본형(lib.TR실시간_정보_해지, RT코드)).G에러()
}

func F실시간_정보_구독(RT코드 string, 종목코드_모음 []string) (에러 error) {
	return F질의(lib.New질의값_복수종목(lib.TR실시간_정보_구독, RT코드, 종목코드_모음)).G에러()
}

func F실시간_정보_해지(RT코드 string, 종목코드_모음 []string) (에러 error) {
	return F질의(lib.New질의값_복수종목(lib.TR실시간_정보_해지, RT코드, 종목코드_모음)).G에러()
}

func F실시간_정보_일괄_해지() (에러 error) {
	return F질의(lib.New질의값_기본형(lib.TR실시간_정보_일괄_해지, "")).G에러()
}

func F실시간_데이터_구독_ETF(종목코드_모음 []string) (에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러}.S실행()

	에러체크(F실시간_정보_구독(xt.RT코스피_호가_잔량, 종목코드_모음))
	에러체크(F실시간_정보_구독(xt.RT코스피_체결, 종목코드_모음))
	에러체크(F실시간_정보_구독(xt.RT코스피_ETF_NAV, 종목코드_모음))
	에러체크(F실시간_정보_구독(xt.RT코스피_시간외_호가_잔량, 종목코드_모음))
	에러체크(F실시간_정보_구독(xt.RT코스피_예상_체결, 종목코드_모음))

	return nil
}

func F실시간_데이터_해지_ETF(종목코드_모음 []string) (에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러}.S실행()

	에러체크(F실시간_정보_해지(xt.RT코스피_호가_잔량, 종목코드_모음))
	에러체크(F실시간_정보_해지(xt.RT코스피_체결, 종목코드_모음))
	에러체크(F실시간_정보_해지(xt.RT코스피_ETF_NAV, 종목코드_모음))
	에러체크(F실시간_정보_해지(xt.RT코스피_시간외_호가_잔량, 종목코드_모음))
	에러체크(F실시간_정보_해지(xt.RT코스피_예상_체결, 종목코드_모음))

	return nil
}

func F주문_응답_실시간_정보_구독() (에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러}.S실행()

	에러체크(F실시간_정보_구독_단순TR(xt.RT현물주문_접수))
	에러체크(F실시간_정보_구독_단순TR(xt.RT현물주문_체결))
	에러체크(F실시간_정보_구독_단순TR(xt.RT현물주문_정정))
	에러체크(F실시간_정보_구독_단순TR(xt.RT현물주문_취소))
	에러체크(F실시간_정보_구독_단순TR(xt.RT현물주문_거부))

	return nil
}

func F주문_응답_실시간_정보_해지() (에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러}.S실행()

	에러체크(F실시간_정보_해지_단순TR(xt.RT현물주문_접수))
	에러체크(F실시간_정보_해지_단순TR(xt.RT현물주문_체결))
	에러체크(F실시간_정보_해지_단순TR(xt.RT현물주문_정정))
	에러체크(F실시간_정보_해지_단순TR(xt.RT현물주문_취소))
	에러체크(F실시간_정보_해지_단순TR(xt.RT현물주문_거부))

	return nil
}
