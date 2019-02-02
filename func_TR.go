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
	"time"
)

func F질의(질의값 lib.I질의값, 옵션_모음 ...interface{}) (값 *lib.S바이트_변환_모음) {
	var 에러 error

	defer lib.S예외처리{M에러: &에러, M함수: func() {
		값 = lib.New바이트_변환_모음_단순형(lib.MsgPack, 에러)
	}}.S실행()

	lib.F확인(F질의값_종목코드_검사(질의값))
	f전송_권한_획득(질의값.TR코드())

	소켓REQ := 소켓REQ_저장소.G소켓()
	defer 소켓REQ_저장소.S회수(소켓REQ)

	if len(옵션_모음) > 0 {
		소켓REQ.S옵션(옵션_모음...)
	}

	return 소켓REQ.G질의_응답_검사(lib.P변환형식_기본값, 질의값)
}

func F질의_단일TR(질의값 lib.I질의값, 옵션_모음 ...interface{}) (값 interface{}, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = 에러 }}.S실행()

	타임아웃 := lib.P1분

	for _, 옵션 := range 옵션_모음 {
		switch 변환값 := 옵션.(type) {
		case time.Duration:
			타임아웃 = 변환값
		}
	}

	i식별번호 := F질의(질의값, 옵션_모음...).G해석값_단순형(0)
	식별번호, ok := i식별번호.(int)

	lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T', '%v'\n"+
		"Xing API에서 식별번호를 부여받고, 콜백을 통해서 응답이 있는 경우에만 사용할 것.\n"+
		"그렇지 않은 경우에는 F질의()를 사용할 것.", i식별번호, i식별번호)

	ch회신 := 대기소_C32.S추가(식별번호, 질의값.TR코드())

	select {
	case 값 := <-ch회신:
		switch 변환값 := 값.(type) {
		case error:
			return nil, 변환값
		default:
			return 값, nil
		}
	case <-time.After(타임아웃):
		return nil, lib.New에러("타임아웃. '%v' '%v'", 질의값.TR코드(), 식별번호)
	}
}

func F접속됨() (접속됨 bool, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 접속됨 = false }}.S실행()

	질의값 := lib.New질의값_기본형(TR접속됨, "")
	접속됨 = lib.F확인(F질의(질의값, lib.P10초).G해석값(0)).(bool)

	return 접속됨, nil
}

func F계좌번호_모음() (계좌번호_모음 []string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌번호_모음 = nil }}.S실행()

	계좌_수량 := lib.F확인(F계좌_수량()).(int)
	계좌번호_모음 = make([]string, 계좌_수량)

	for i := 0; i < 계좌_수량; i++ {
		계좌번호_모음[i] = lib.F확인(F계좌_번호(0)).(string)
	}

	return 계좌번호_모음, nil
}

func F계좌_수량() (계좌_수량 int, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_수량 = 0 }}.S실행()

	회신_메시지 := F질의(lib.New질의값_기본형(TR계좌_수량, ""))
	계좌_수량 = lib.F확인(회신_메시지.G해석값(0)).(int)
	lib.F조건부_패닉(계좌_수량 == 0, "계좌 수량 0.")

	return 계좌_수량, nil
}

func F계좌_번호(인덱스 int) (계좌_번호 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_번호 = "" }}.S실행()

	회신_메시지 := F질의(lib.New질의값_정수(TR계좌_번호, "", 인덱스))
	계좌_번호 = lib.F확인(회신_메시지.G해석값(0)).(string)

	return 계좌_번호, nil
}

func F계좌_이름(계좌_번호 string) (계좌_이름 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_이름 = "" }}.S실행()

	회신_메시지 := F질의(lib.New질의값_문자열(TR계좌_이름, "", 계좌_번호))
	계좌_이름 = lib.F확인(회신_메시지.G해석값(0)).(string)

	return 계좌_이름, nil
}

func F계좌_상세명(계좌_번호 string) (계좌_상세명 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_상세명 = "" }}.S실행()

	회신_메시지 := F질의(lib.New질의값_문자열(TR계좌_상세명, "", 계좌_번호))
	계좌_상세명 = lib.F확인(회신_메시지.G해석값(0)).(string)

	return 계좌_상세명, nil
}

