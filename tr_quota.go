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

func f전송_권한_획득(TR코드 string) {
	switch TR코드 {
	case "", RT현물_주문_접수_SC0, RT현물_주문_체결_SC1, RT현물_주문_정정_SC2, RT현물_주문_취소_SC3, RT현물_주문_거부_SC4,
		RT코스피_호가_잔량_H1, RT코스피_시간외_호가_잔량_H2, RT코스피_체결_S3, RT코스피_예상_체결_YS3,
		RT코스피_ETF_NAV_I5, RT주식_VI발동해제_VI, RT시간외_단일가VI발동해제_DVI, RT장_운영정보_JIF:
		return
	}

	F접속_확인()
	f10분당_전송_제한_확인(TR코드)
	f초당_전송_제한_확인(TR코드)
}

func f10분당_전송_제한_확인(TR코드 string) lib.I전송_권한 {
	전송_권한, 존재함 := tr코드별_10분당_전송_제한[TR코드]

	switch {
	case !존재함:
		return nil // 해당 TR코드 관련 제한이 존재하지 않음.
	case 전송_권한.TR코드() != TR코드:
		panic("예상하지 못한 경우.")
	}

	return 전송_권한.G획득()
}

func f초당_전송_제한_확인(TR코드 string) lib.I전송_권한 {
	전송_권한, 존재함 := tr코드별_초당_전송_제한[TR코드]

	switch {
	case !존재함:
		panic(lib.New에러("전송제한을 찾을 수 없음 : '%v'", TR코드))
	case 전송_권한.TR코드() != TR코드:
		panic("예상하지 못한 경우.")
	case 전송_권한.G남은_수량() > 100:
		panic("전송 한도가 너무 큼. 1초당 한도와 10분당 한도를 혼동한 듯함.")
	}

	return 전송_권한.G획득()
}
