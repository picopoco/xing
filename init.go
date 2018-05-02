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

func init() {
	ch초기화_신호_C32_모음 = make([]chan xt.T신호_C32, 2)

	for i:=0 ; i<len(ch초기화_신호_C32_모음) ; i++ {
		ch초기화_신호_C32_모음[i] = make(chan xt.T신호_C32, 1)
	}
}

func F초기화() {
	defer lib.S에러패닉_처리기{}.S실행()

	f초기화_소켓()
	f초기화_Go루틴()
	f초기화_xing_C32()	// xing_C32가 실행되면 자동으로 서버 접속까지 진행함.
	lib.F조건부_패닉(!f초기화_작동_확인(), "초기화 작동 확인 실패.")
	//
	//lib.F체크포인트()
	//
	//lib.F메모("f접속유지_실행() 보류")	//f접속유지_실행()
	//lib.F메모("f초기화_영업일_기준_전일_당일() 보류")	//f초기화_영업일_기준_전일_당일()
	//
	//lib.F문자열_출력("\n\n*** 초기화 완료 ***\n\n")
}

func f초기화_소켓() {
	소켓REP_TR콜백 = lib.NewNano소켓REP_raw_단순형(lib.P주소_Xing_C함수_콜백)
}

func f초기화_xing_C32() (에러 error) {
	xing_C32_실행_잠금.Lock()
	defer xing_C32_실행_잠금.Unlock()

	if !lib.F인터넷에_접속됨() {
		lib.F문자열_출력("인터넷을 확인하십시오.")
		return
	} else if 프로세스ID := xing_C32_실행_중(); 프로세스ID >= 0 {
		lib.F문자열_출력("xing_C32 가 이미 실행 중입니다.")
		return nil
	}

	_, 에러 = lib.F외부_프로세스_실행(xing_C32_경로)

	return 에러
}

func f초기화_Go루틴() {
	ch초기화 := make(chan lib.T신호, 1)
	go go루틴_콜백_처리(ch초기화)
	<-ch초기화
}

func f초기화_영업일_기준_전일_당일() (에러 error) {
	panic("TODO")
}

func f초기화_작동_확인() bool {
	ch확인 := make(chan lib.T신호, 1)
	ch타임아웃 := time.After(lib.P10분)

	select {
	case <-ch초기화_신호_C32_모음[xt.P신호_C32_Ready]:	// 서버 접속된 상태임.
	case <-ch타임아웃:
		return false
	}

	// C32 모듈의 소켓이 초기화 될 시간을 준다.
	// 이게 없으면 제대로 작동하지 않으며, 필수적인 부분임. 삭제하지 말 것.
	lib.F대기(lib.P10초)

	// 소켓REP_TR수신 동작 테스트
	go F소켓REP_TR수신_동작_여부_확인(ch확인)

	select {
	case <-ch확인:
	case <-ch타임아웃:
		lib.F체크포인트("F소켓REP_TR수신_동작_여부_확인() 타임아웃.")
		return false
	}

	// F접속됨() 테스트
	go F접속됨_확인(ch확인)

	select {
	case <-ch확인:
	case <-ch타임아웃:
		lib.F체크포인트("F접속됨_확인() 타임아웃.")
		return false
	}

	// F시각_조회_t0167() 테스트
	go F시각_조회_t0167_확인(ch확인)

	select {
	case <-ch확인:
	case <-ch타임아웃:
		lib.F체크포인트("F시각_조회_t0167_확인() 타임아웃.")
		return false
	}

	lib.F체크포인트("정상 작동 확인 완료.")

	return true
}

func F소켓REP_TR수신_동작_여부_확인(ch완료 chan lib.T신호) {
	defer func() { ch완료 <- lib.P신호_종료 }()

	테스트_질의값 := xt.New호출_인수_정수값(xt.P함수_신호, xt.P신호_Go_소켓REP_TR수신_테스트)
	소켓 := lib.NewNano소켓REQ_단순형(lib.P주소_Xing_C함수_호출, lib.P5초)

	for i := 0; i < 100; i++ {
		응답 := 소켓.G질의_응답_검사(lib.P변환형식_기본값, 테스트_질의값)

		if 응답.G에러() == nil && 응답.G해석값_단순형(0).(bool) {
			return
		}
	}
}

func F접속됨_확인(ch완료 chan lib.T신호) {
	defer func() { ch완료 <- lib.P신호_종료 }()

	호출_인수 := xt.New호출_인수_기본형(xt.P함수_접속됨)

	for i := 0; i < 100; i++ {
		if 응답 := F질의by호출_인수(호출_인수, lib.P10초); 응답.G에러() != nil {
			lib.F에러_출력(응답.G에러())
			continue
		} else if 접속됨, ok := 응답.G해석값_단순형(0).(bool); !ok {
			panic(lib.New에러("예상하지 못한 자료형 : '%T'", 응답.G해석값_단순형(0)))
		} else if !접속됨 {
			panic(lib.New에러("이 시점에 접속되어 있어야 함."))
		}

		return
	}
}

func F시각_조회_t0167_확인(ch완료 chan lib.T신호) {
	defer func() { ch완료 <- lib.P신호_종료 }()

	for i := 0; i < 100; i++ {
		시각, 에러 := F시각_조회_t0167()

		if 에러 != nil || 시각.Equal(time.Time{}) {
			continue
		} else if 차이 := time.Now().Sub(시각); 차이 < -1 * lib.P10분 || 차이 > lib.P10분 {
			panic(lib.New에러("서버와 시스템 시각 불일치 : 차이 '%v'분", 차이.Minutes()))
		}

		return
	}


}

func f리소스_정리() {
	lib.F패닉억제_호출(F질의by호출_인수, xt.New호출_인수_정수값(xt.P함수_신호, xt.P신호_C32_종료), lib.P10초)
	lib.F패닉억제_호출(소켓REP_TR콜백.Close)

	for {
		if 프로세스ID := xing_C32_실행_중(); 프로세스ID < 0 {
			break
		} else {
			lib.F대기(lib.P1초)
		}
	}

	lib.F공통_종료_채널_닫기()
}

