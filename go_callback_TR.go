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

func go루틴_콜백_처리(ch초기화 chan lib.T신호) {
	콜백_처리_루틴_수량 := 10

	ch종료 := lib.F공통_종료_채널()
	ch도우미_초기화 := make(chan lib.T신호, 콜백_처리_루틴_수량)
	ch도우미_종료 := make(chan error, 콜백_처리_루틴_수량)

	for i := 0; i < 콜백_처리_루틴_수량; i++ {
		go go루틴_콜백_처리_도우미(ch도우미_초기화, ch도우미_종료)
	}

	for i := 0; i < 콜백_처리_루틴_수량; i++ {
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
			go go루틴_콜백_처리_도우미(ch도우미_초기화, ch도우미_종료)
			<-ch도우미_초기화
		}
	}
}

func go루틴_콜백_처리_도우미(ch초기화 chan lib.T신호, ch에러 chan error) {
	defer lib.S에러패닉_처리기{M함수with내역: func(r interface{}) { ch에러 <- lib.New에러(r) }}.S실행()

	ch종료 := lib.F공통_종료_채널()
	ch초기화 <- lib.P신호_초기화

	for {
		select {
		case <-ch종료:
			return
		case 바이트_모음 := 에러체크(소켓SUB_콜백.Recv()).([]byte):
			회신_메시지 := lib.New소켓_메시지by바이트_모음(바이트_모음)

			값 := 회신_메시지.S해석기(xt.F바이트_변환값_해석).G해석값_단순형(0).(xt.I콜백)

			switch 값.G콜백() {
			case xt.P콜백_TR데이터, xt.P콜백_메시지_및_에러, xt.P콜백_TR완료:
				f콜백_TR데이터_처리기(값)
			case xt.P콜백_실시간_데이터:
				f콜백_실시간_데이터_처리기(값.(*xt.S콜백_실시간_데이터))
			case xt.P콜백_타임아웃:
				panic("TODO") // 변환값 := 값.(*xt.S콜백_정수값)
			case xt.P콜백_로그아웃:
				panic("TODO") // 변환값 := 값.(*xt.S콜백_기본형)
			case xt.P콜백_접속해제:
				panic("TODO") // 변환값 := 값.(*xt.S콜백_기본형)
			case xt.P콜백_링크_데이터:
				panic("TODO") // 변환값 := 값.(*xt.S콜백_기본형)
			case xt.P콜백_실시간_차트_데이터:
				panic("TODO") // //변환값 := 값.(*xt.S콜백_기본형)
			default:
				panic(lib.F2문자열("예상하지 못한 콜백 구분값 : '%v'", 값.G콜백()))
			}
		}
	}
}

func f콜백_TR데이터_처리기(값 xt.I콜백) {
	defer lib.S에러패닉_처리기{}.S실행()

	콜백_구분 := 값.G콜백()
	var 식별번호 int
	var TR코드 string
	var 대기_항목 *대기_항목_C32

	switch 콜백_구분 {
	case xt.P콜백_TR데이터:
		변환값 := 값.(*xt.S콜백_TR데이터)
		식별번호 = 변환값.M식별번호
		대기_항목 = 대기소_C32.G값(식별번호)
		TR코드 = 대기_항목.TR코드
	case xt.P콜백_메시지_및_에러:
		변환값 := 값.(*xt.S콜백_메시지_및_에러)
		식별번호 = 변환값.M식별번호
		대기_항목 = 대기소_C32.G값(식별번호)
		TR코드 = 대기_항목.TR코드
	case xt.P콜백_TR완료:
		변환값 := 값.(*xt.S콜백_정수값)
		식별번호 = 변환값.M정수값
		대기_항목 = 대기소_C32.G값(식별번호)
		TR코드 = 대기_항목.TR코드
	default:
		panic(lib.F2문자열("예상하지 못한 콜백 구분값 : '%v'", 콜백_구분))
	}

	switch TR코드 {
	case xt.TR시간_조회:
		panic("TODO")
	case xt.TR계좌_번호:
		panic("TODO")
	case xt.TR현물_정상주문:
		panic("TODO")
	case xt.TR현물_정정주문:
		panic("TODO")
	case xt.TR현물_취소주문:
		panic("TODO")
	case xt.TR계좌_거래_내역:
		panic("TODO")
	case xt.TR현물_호가_조회:
		panic("TODO")
	case xt.TR현물_시세_조회:
		panic("TODO")
	case xt.TR현물_시간대별_체결_조회:
		panic("TODO")
	case xt.TR현물_기간별_조회:
		panic("TODO")
	case xt.TR현물_당일_전일_분틱_조회:
		panic("TODO")
	case xt.TR_ETF_시세_조회:
		panic("TODO")
	case xt.TR_ETF_시간별_추이:
		panic("TODO")
	case xt.TR현물_종목_조회:
		panic("TODO")
	case xt.TR주식_매매일지_수수료_금일, xt.TR주식_매매일지_수수료_날짜_지정, xt.TR주식_잔고_2,
		xt.TR주식_체결_미체결, xt.TR종목별_증시_일정, xt.TR해외_실시간_지수, xt.TR해외_지수_조회,
		xt.TR증시_주변_자금_추이, xt.TR현물계좌_예수금_주문가능금액_총평가, xt.TR현물계좌_잔고내역,
		xt.TR현물계좌_주문체결내역, xt.TR계좌별_신용한도, xt.TR현물계좌_증거금률별_주문가능수량,
		xt.TR주식계좌_기간별_수익률_상세:
		panic(lib.F2문자열("미구현 : '%v'", TR코드))
	default:
		panic(lib.F2문자열("예상하지 못한 TR코드 : '%v'", TR코드))
	}
}
