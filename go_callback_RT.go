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

func go루틴_실시간_데이터_처리(ch초기화 chan lib.T신호) {
	실시간_데이터_처리_도우미_수량 := 10

	ch종료 := lib.F공통_종료_채널()
	ch도우미_초기화 := make(chan lib.T신호, 실시간_데이터_처리_도우미_수량)
	ch도우미_종료 := make(chan error, 실시간_데이터_처리_도우미_수량)

	for i := 0; i < 실시간_데이터_처리_도우미_수량; i++ {
		go go루틴_실시간_데이터_처리_도우미(ch도우미_초기화, ch도우미_종료)
	}

	for i := 0; i < 실시간_데이터_처리_도우미_수량; i++ {
		<-ch도우미_초기화
	}

	ch초기화 <- lib.P신호_초기화

	for {
		select {
		case <-ch종료:
			return
		case 에러 := <-ch도우미_종료:
			select {
			case <-ch종료:
				return
			default:
			}

			lib.F에러_출력(에러)
			go go루틴_실시간_데이터_처리_도우미(ch도우미_초기화, ch도우미_종료)
			<-ch도우미_초기화
		}
	}
}

func go루틴_실시간_데이터_처리_도우미(ch초기화 chan lib.T신호, ch에러 chan error) {
	defer lib.S에러패닉_처리기{M함수with내역: func(r interface{}) { ch에러 <- lib.New에러(r) }}.S실행()

	ch종료 := lib.F공통_종료_채널()
	ch초기화 <- lib.P신호_초기화

	for {
		select {
		case <-ch종료:
			return
		case 바이트_모음 := 에러체크(소켓SUB_실시간_정보.Recv()).([]byte):
			회신_메시지 := lib.New소켓_메시지by바이트_모음(바이트_모음)

			값 := new(*xt.S콜백_실시간_데이터)
			에러체크(회신_메시지.G값(0, 값))

			f콜백_실시간_데이터_처리기(값)
		}
	}
}

func f콜백_실시간_데이터_처리기(값 *xt.S콜백_실시간_데이터) {
	switch 값.RT코드 {
	case xt.RT현물주문_접수:
		panic("TODO")
	case xt.RT현물주문_체결:
		panic("TODO")
	case xt.RT현물주문_정정:
		panic("TODO")
	case xt.RT현물주문_취소:
		panic("TODO")
	case xt.RT현물주문_거부:
		panic("TODO")
	case xt.RT코스피_호가_잔량:
		panic("TODO")
	case xt.RT코스피_시간외_호가_잔량:
		panic("TODO")
	case xt.RT코스피_체결:
		panic("TODO")
	case xt.RT코스피_예상_체결:
		panic("TODO")
	case xt.RT코스피_ETF_NAV:
		panic("TODO")
	case xt.RT주식_VI발동해제:
		panic("TODO")
	case xt.RT시간외_단일가VI발동해제:
		panic("TODO")
	case xt.RT장_운영정보:
		panic("TODO")
	case xt.RT코스닥_체결, xt.RT코스피_거래원, xt.RT코스닥_거래원,
		xt.RT코스피_기세, xt.RT코스닥_LP호가, xt.RT코스닥_호가잔량,
		xt.RT코스닥_시간외_호가잔량, xt.RT지수, xt.RT예상지수,
		xt.RT코스닥_예상_체결, xt.RT실시간_뉴스_제목_패킷, xt.RT업종별_투자자별_매매_현황:
		panic(lib.F2문자열("미구현 : '%v'", 값.RT코드))
	default:
		panic(lib.F2문자열("예상하지 못한 TR코드 : '%v'", 값.RT코드))
	}
}
