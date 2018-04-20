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
		default:
			바이트_모음 := 에러체크(소켓SUB_콜백.Recv()).([]byte)
			회신_메시지 := lib.New소켓_메시지by바이트_모음(바이트_모음)
			회신_메시지.S해석기(xt.F바이트_변환값_해석)
			값 := 회신_메시지.G해석값_단순형(0).(xt.I콜백)

			if 값.G콜백() != xt.P콜백_소켓PUB_확인 {
				lib.F체크포인트(값.G콜백(), 값)
			}

			switch 값.G콜백() {
			case xt.P콜백_TR데이터, xt.P콜백_메시지_및_에러, xt.P콜백_TR완료:
				f콜백_TR데이터_처리기(값)
			case xt.P콜백_타임아웃:
				panic("TODO") // 변환값 := 값.(*xt.S콜백_정수값)
			case xt.P콜백_접속해제:
				panic("TODO") // 변환값 := 값.(*xt.S콜백_기본형)
			case xt.P콜백_링크_데이터:
				panic("TODO") // 변환값 := 값.(*xt.S콜백_기본형)
			case xt.P콜백_실시간_차트_데이터:
				panic("TODO") // //변환값 := 값.(*xt.S콜백_기본형)
			case xt.P콜백_소켓PUB_확인:
				호출_인수 := xt.New호출_인수_기본형(xt.P함수_소켓PUB_확인)
				F질의by호출_인수No검사(호출_인수, lib.P3초)

				select {
				case ch초기화_소켓PUB_확인 <- lib.P신호_종료:
				default:
				}
			default:
				panic(lib.New에러("예상하지 못한 콜백 구분값 : '%v'", 값.G콜백()))
			}
		}
	}
}

func f콜백_TR데이터_처리기(값 xt.I콜백) {
	defer lib.S에러패닉_처리기{}.S실행()

	식별번호, 대기_항목, TR코드 := f콜백_데이터_식별번호(값)
	lib.F조건부_패닉(!f처리_가능한_TR코드(TR코드), "처리 불가 TR코드 : '%v'", TR코드)
	lib.F조건부_패닉(대기_항목 == nil, "TR 식별번호 '%v' : nil 대기항목.", 식별번호)

	대기_항목.Lock()
	defer 대기_항목.Unlock()

	lib.F체크포인트(식별번호, 값.G콜백(), TR코드, 대기_항목)

	switch 값.G콜백() {
	case xt.P콜백_TR데이터:
		에러체크(f데이터_복원(대기_항목, 값.(*xt.S콜백_TR데이터).M데이터))

		lib.F체크포인트(대기_항목.식별번호, 대기_항목.데이터_수신, 대기_항목.응답_완료, 대기_항목.코드, 대기_항목.메시지)
	case xt.P콜백_메시지_및_에러:
		lib.F체크포인트()
		변환값 := 값.(*xt.S콜백_메시지_및_에러)

		대기_항목.코드 = 변환값.M코드
		대기_항목.메시지 = 변환값.M내용

		if f에러_발생(TR코드, 변환값.M코드, 변환값.M내용) {
			lib.F체크포인트(대기_항목.식별번호, 변환값.M코드, 변환값.M내용)
			대기_항목.에러_발생 = true
		}
		lib.F체크포인트(대기_항목.식별번호, 대기_항목.데이터_수신, 대기_항목.응답_완료, 대기_항목.코드, 대기_항목.메시지)
	case xt.P콜백_TR완료:
		lib.F체크포인트()
		대기_항목.응답_완료 = true

		lib.F체크포인트(대기_항목.식별번호, 대기_항목.데이터_수신, 대기_항목.응답_완료, 대기_항목.코드, 대기_항목.메시지)
	default:
		panic(lib.New에러("예상하지 못한 경우. 콜백 구분값 : '%v', 자료형 : '%T'", 값.G콜백(), 값))
	}

	lib.F체크포인트(대기_항목.식별번호, 대기_항목.데이터_수신, 대기_항목.응답_완료, 대기_항목.코드, 대기_항목.메시지)
	// TR응답 데이터 수신 및 완료 확인이 되었는 지 확인.
	switch {
	case !대기_항목.데이터_수신, !대기_항목.응답_완료:
		lib.F체크포인트()
		return // 추가 데이터 수신이 필요하거나, 응답 완료 확인이 필요함.
	case 대기_항목.에러_발생:
		lib.F체크포인트()
		대기소_C32.S회신(식별번호)
		return
	}

	lib.F체크포인트(대기_항목.식별번호, 대기_항목.데이터_수신, 대기_항목.응답_완료, 대기_항목.코드, 대기_항목.메시지)

	대기_항목.S회신()

	lib.F체크포인트(대기_항목.식별번호, 대기_항목.데이터_수신, 대기_항목.응답_완료, 대기_항목.코드, 대기_항목.메시지)
}
