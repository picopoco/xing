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
	"github.com/go-mangos/mangos"
)

func go_TR콜백_처리(ch초기화 chan lib.T신호) {
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

func go루틴_콜백_처리_도우미(ch초기화 chan lib.T신호, ch도우미_종료 chan error) (에러 error) {
	var 수신_메시지 *mangos.Message // 최대한 재활용 해야 성능 문제를 걱정할 필요가 없어진다.

	defer func() { ch도우미_종료 <- 에러 }()
	defer lib.S예외처리{
		M에러: &에러,
		M함수with내역: func(r interface{}) {
			if 수신_메시지 != nil {
				소켓REP_TR콜백.S회신_Raw(수신_메시지, lib.JSON, lib.New에러(r))
			}
		}}.S실행()

	var 콜백값 I콜백
	var ok bool
	var 수신값 *lib.S바이트_변환_모음
	ch종료 := lib.F공통_종료_채널()

	ch초기화 <- lib.P신호_초기화

	for {
		select {
		case <-ch종료:
			return
		default:
			수신_메시지, 에러 = 소켓REP_TR콜백.G수신_Raw()
			if 에러 != nil {
				select {
				case <-ch종료:
					에러 = nil
					return
				default:
					lib.F에러_출력(lib.New에러(에러))
					continue
				}
			}

			수신값 = lib.New바이트_변환_모음from바이트_배열_단순형(수신_메시지.Body)
			lib.F조건부_패닉(수신값.G수량() != 1, "메시지 길이 : 예상값 1, 실제값 %v.", 수신값.G수량())

			i값 := 수신값.S해석기(F바이트_변환값_해석).G해석값_단순형(0)


			콜백값, ok = i값.(I콜백)
			lib.F조건부_패닉(!ok, "'I콜백'형이 아님 : '%T'", i값)

			변환_형식 := 수신값.G변환_형식(0)

			//lib.F체크포인트(값.G콜백())

			switch 콜백값.G콜백() {
			case P콜백_TR데이터, P콜백_메시지_및_에러, P콜백_TR완료, P콜백_타임아웃:
				에러 = f콜백_TR데이터_처리기(콜백값)
			case P콜백_신호:
				에러 = f콜백_신호_처리기(콜백값)
			case P콜백_링크_데이터, P콜백_실시간_차트_데이터:
				panic("TODO") // 변환값 := 값.(*S콜백_기본형)
			default:
				panic(lib.New에러("예상하지 못한 콜백 구분값 : '%v'", 콜백값.G콜백()))
			}

			소켓REP_TR콜백.S회신_Raw(수신_메시지, 변환_형식, lib.P신호_OK)
		}
	}
}

func f콜백_TR데이터_처리기(값 I콜백) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행_No출력()

	식별번호, 대기_항목, TR코드 := f콜백_데이터_식별번호(값)

	if 식별번호 == 0 || 대기_항목 == nil || TR코드 == "" {
		if 값.G콜백() == P콜백_타임아웃 {
			lib.F체크포인트()
			return nil
		}

		panic(lib.New에러("대기 항목 없음. '%v' '%v' '%v' '%v'", 값.G콜백(), 식별번호, 대기_항목, TR코드))
	}

	lib.F조건부_패닉(식별번호 == 0 || 대기_항목 == nil || TR코드 == "", "대기항목 없음.")
	lib.F조건부_패닉(!f처리_가능한_TR코드(TR코드), "처리 불가 TR코드 : '%v'", TR코드)
	lib.F조건부_패닉(대기_항목 == nil, "TR 식별번호 '%v' : nil 대기항목.", 식별번호)

	대기_항목.Lock()
	defer 대기_항목.Unlock()

	switch 값.G콜백() {
	case P콜백_TR데이터:
		if 에러 := f데이터_복원(대기_항목, 값.(*S콜백_TR데이터).M데이터); 에러 != nil && 대기_항목.에러 == nil {
			대기_항목.에러 = 에러
		}
	case P콜백_메시지_및_에러:
		변환값 := 값.(*S콜백_메시지_및_에러)

		if f에러_발생(TR코드, 변환값.M코드, 변환값.M내용) {
			lib.F문자열_출력("에러 발생 : %v\n%v : %v", 대기_항목.식별번호, 변환값.M코드, 변환값.M내용)
			대기_항목.에러 = lib.New에러("%s : %s", 변환값.M코드, 변환값.M내용)
		} else {
			대기_항목.메시지_수신 = true
		}
	case P콜백_TR완료:
		대기_항목.응답_완료 = true
	case P콜백_타임아웃:
		대기_항목.에러 = lib.New에러("타임아웃.")
		//lib.F체크포인트("타임아웃.")
	default:
		lib.F체크포인트()
		panic(lib.New에러("예상하지 못한 경우. 콜백 구분값 : '%v', 자료형 : '%T'", 값.G콜백(), 값))
	}

	// TR응답 데이터 수신 및 완료 확인이 되었는 지 확인.
	switch {
	case !대기_항목.데이터_수신, !대기_항목.응답_완료, !대기_항목.메시지_수신:
		//lib.F체크포인트(값.G콜백(), 대기_항목.식별번호, "추가 수신 필요.")
		return
	case 대기_항목.에러 != nil:
		lib.F체크포인트(값.G콜백(), 대기_항목.식별번호, "에러 회신")
		대기소_C32.S회신(식별번호)
	default:
		//lib.F체크포인트(값.G콜백(), 대기_항목.식별번호, "TR 완료")
		대기소_C32.S회신(식별번호)
	}

	return
}

func f콜백_신호_처리기(콜백 I콜백) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	콜백_정수값, ok := 콜백.(*S콜백_정수값)
	lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", 콜백)

	정수값 := 콜백_정수값.M정수값
	신호 := T신호_C32(정수값)

	switch 신호 {
	case P신호_C32_READY, P신호_C32_종료:
		select {
		case ch신호_C32_모음[정수값] <- 신호:
		default:
		}
	default:
		return lib.New에러with출력("예상하지 못한 신호 : '%v'", 신호)
	}

	return nil
}
