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
	"fmt"
	"github.com/ghts/lib"
	"time"
)

func init() {
	lib.TR구분_String = TR구분_String

	ch신호_C32_모음 = make([]chan T신호_C32, 2)

	for i := 0; i < len(ch신호_C32_모음); i++ {
		ch신호_C32_모음[i] = make(chan T신호_C32, 1)
	}

	메모 := `\n
- CPU 점유율이 너무 높다. 프로파일링 해서 병목 지점을 제거할 것.
- 재무정보 정확성을 위해서 교차검증이 필요하다.
- 관리 종목 조회 기능 구현할 것. (t1404, t1405)
- 체결 확인 기능 구현할 것. (t0425)
- 계좌 조회 기능 구현할 것. (CDPCQ04700, CSPAQ12200,CSPAQ12300, CSPAQ13700, FOCCQ33600)
- 업종별 조회 기능. (t8424, t4203, t1514, t8419, t8424)
- 테마별 조회 기능. (t8425, t1531)
- 선물 가격 정보 수집 기능 구현할 것. (t8415, t8416, t9943)
- 뉴스 및 공시 정보. (t3102,t3202)
- 해외 지수 (o3123, o3121, t3518, t3521) // o31xx는 해외선물 계좌 내지 모의투자 있어야 보임.
- F계좌_별명(): 원인미상의 메모리 에러가 발생함.`
	lib.F메모(메모)
}


func F초기화() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	lib.F확인(f자료형_크기_비교_확인())
	f초기화_TR전송_제한()
	f초기화_소켓()
	f초기화_Go루틴()
	lib.F확인(f초기화_xing_C32())
	lib.F조건부_패닉(!f초기화_작동_확인(), "초기화 작동 확인 실패.")
	lib.F확인(f종목모음_설정())
	lib.F확인(f전일_당일_설정())
	f접속유지_실행()

	fmt.Println("**     초기화 완료     **")

	return nil
}

func f초기화_소켓() {
	소켓REP_TR콜백 = lib.NewNano소켓REP_raw_단순형(lib.P주소_Xing_C함수_콜백)
	소켓SUB_실시간_정보 = lib.NewNano소켓SUB_단순형(lib.P주소_Xing_실시간).(lib.I소켓Raw)
}

func f초기화_xing_C32() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	xing_C32_실행_잠금.Lock()
	defer xing_C32_실행_잠금.Unlock()

	if !lib.F인터넷에_접속됨() {
		lib.F문자열_출력("인터넷을 확인하십시오.")
		return
	} else if 프로세스ID := xing_C32_실행_중(); 프로세스ID >= 0 {
		lib.F문자열_출력("xing_C32 가 이미 실행 중입니다.")
		return nil
	}

	lib.F확인(lib.F외부_프로세스_실행(xing_C32_경로))

	return nil
}

//func f초기화_xing_COM32() (에러 error) {
//	xing_COM32_실행_잠금.Lock()
//	defer xing_COM32_실행_잠금.Unlock()
//
//	if !lib.F인터넷에_접속됨() {
//		lib.F문자열_출력("인터넷을 확인하십시오.")
//		return
//	} else if 프로세스ID := xing_COM32_실행_중(); 프로세스ID >= 0 {
//		lib.F문자열_출력("xing_COM32 가 이미 실행 중입니다.")
//		return nil
//	}
//
//	_, 에러 = lib.F외부_프로세스_실행(xing_COM32_경로)
//
//	return 에러
//}

func f초기화_Go루틴() {
	고루틴_모음 := []func(chan lib.T신호) error{go_TR호출, go_TR콜백_처리, go_RT_주문처리결과}
	ch초기화 := make(chan lib.T신호, len(고루틴_모음))

	for _, 고루틴 := range 고루틴_모음 {
		go 고루틴(ch초기화)
	}

	for _ = range 고루틴_모음 {
		<-ch초기화
	}
}

func f초기화_작동_확인() (작동_여부 bool) {
	defer lib.S예외처리{M함수: func() { 작동_여부 = false}}.S실행()

	ch확인 := make(chan lib.T신호, 1)
	ch타임아웃 := time.After(lib.P10분)

	select {
	case <-ch신호_C32_모음[P신호_C32_READY]: // 서버 접속된 상태임.
	case <-ch타임아웃:
		lib.F체크포인트("C32 초기화 타임아웃")
		return false
	}

	// C32 모듈의 소켓이 초기화 될 시간을 준다.
	// 이게 없으면 제대로 작동하지 않으며, 필수적인 부분임. 삭제하지 말 것.
	lib.F대기(lib.P10초)

	// 소켓REP_TR수신 동작 테스트
	go tr수신_소켓_동작_확인(ch확인)

	select {
	case <-ch확인:
	case <-ch타임아웃:
		lib.F체크포인트("F소켓REP_TR수신_동작_여부_확인() 타임아웃.")
		return false
	}

	// F접속됨() 테스트
	go f접속_확인(ch확인)

	select {
	case <-ch확인:
	case <-ch타임아웃:
		lib.F체크포인트("F접속됨_확인() 타임아웃.")
		return false
	}

	// F시각_조회_t0167() 테스트
	go tr동작_확인(ch확인)

	select {
	case <-ch확인:
	case <-ch타임아웃:
		lib.F체크포인트("F시각_조회_t0167_확인() 타임아웃.")
		return false
	}

	fmt.Println("**     C32 동작 확인 완료     **")

	return true
}

func tr수신_소켓_동작_확인(ch완료 chan lib.T신호) {
	defer func() { ch완료 <- lib.P신호_종료 }()

	for i := 0; i < 100; i++ {
		if 응답 := F질의(lib.New질의값_기본형(TR소켓_테스트, ""), lib.P5초); 응답.G에러() == nil {
			return
		}
	}
}

func f접속_확인(ch완료 chan lib.T신호) {
	defer func() { ch완료 <- lib.P신호_종료 }()

	for i := 0; i < 100; i++ {
		if 접속됨, 에러 := F접속됨(); 에러 != nil {
			lib.F에러_출력(에러)
			continue
		} else if !접속됨 {
			panic(lib.New에러("이 시점에 접속되어 있어야 함."))
		}

		return
	}
}

func tr동작_확인(ch완료 chan lib.T신호) {
	defer func() { ch완료 <- lib.P신호_종료 }()

	for i := 0; i < 100; i++ {
		시각, 에러 := (<-F시각_조회_t0167()).G값()

		if 에러 != nil || 시각.Equal(time.Time{}) {
			continue
		} else if 차이 := time.Now().Sub(시각); 차이 < -1*lib.P10분 || 차이 > lib.P10분 {
			panic(lib.New에러("서버와 시스템 시각 불일치 : 차이 '%v'분", 차이.Minutes()))
		}

		return
	}
}

func f전일_당일_설정() (에러 error) {
	lib.S예외처리{M에러: &에러}.S실행()

	const 수량 = 30

	질의값_기간별_조회 := New질의값_현물_기간별_조회()
	질의값_기간별_조회.M구분 = TR조회
	질의값_기간별_조회.M코드 = TR현물_기간별_조회
	질의값_기간별_조회.M종목코드 = "069500"
	질의값_기간별_조회.M일주월_구분 = P일주월_일
	질의값_기간별_조회.M연속키 = ""
	질의값_기간별_조회.M수량 = 수량

	i응답값, 에러 := F질의_단일TR(질의값_기간별_조회)
	lib.F확인(에러)

	switch 응답값 := i응답값.(type) {
	case *S현물_기간별_조회_응답:
		lib.F조건부_패닉(응답값.M헤더.M수량 != int64(수량), "예상하지 못한 수량 : '%v' '%v'", 응답값.M헤더.M수량, 수량)
		lib.F조건부_패닉(len(응답값.M반복값_모음.M배열) != 수량, "예상하지 못한 수량 : '%v' '%v'", len(응답값.M반복값_모음.M배열), 수량)
		lib.F조건부_패닉(응답값.M반복값_모음.M배열[0].M일자.Before(응답값.M반복값_모음.M배열[1].M일자), "예상하지 못한 순서")

		당일 = lib.New안전한_시각(응답값.M반복값_모음.M배열[0].M일자)
		전일 = lib.New안전한_시각(응답값.M반복값_모음.M배열[1].M일자)

		최근_영업일_모음 = make([]time.Time, 수량, 수량)

		for i, 값 := range 응답값.M반복값_모음.M배열 {
			최근_영업일_모음[i] = lib.F2일자(값.M일자)
		}

		return nil
	default:
		panic(lib.New에러("예상하지 못한 자료형 : '%T'", i응답값))
	}
}

func C32_종료() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	lib.F체크포인트()

	// 종료 신호 전송
	func() {
		defer lib.S예외처리{M출력_숨김: true}.S실행()

		소켓REQ := 소켓REQ_저장소.G소켓()
		defer 소켓REQ_저장소.S회수(소켓REQ)

		소켓REQ.S옵션(lib.P1초)
		소켓REQ.S송신(lib.MsgPack, lib.New질의값_기본형(TR종료, ""))
	}()

	select {
	case <-ch신호_C32_모음[P신호_C32_종료]:
	case <-time.After(lib.P1초):
	}

	lib.F체크포인트()

	// 강제 종료
	for {
		if 프로세스ID := xing_C32_실행_중(); 프로세스ID < 0 {
			lib.F체크포인트()
			return
		} else {
			lib.F프로세스_종료by프로세스ID(프로세스ID)
			lib.F대기(lib.P3초)
		}
	}
}

func F리소스_정리() {
	C32_종료()
	lib.F공통_종료_채널_닫기()
	lib.F패닉억제_호출(소켓REP_TR콜백.Close)
}

func f초기화_TR전송_제한() {
	코드별_10분당_TR전송_제한 := make(map[string]int)

	코드별_10분당_TR전송_제한[TR현물_기간별_조회] = 200
	코드별_10분당_TR전송_제한[TR_ETF_시간별_추이] = 200
	코드별_10분당_TR전송_제한[TR기업정보_요약] = 200
	코드별_10분당_TR전송_제한[TR재무순위_종합] = 200
	코드별_10분당_TR전송_제한[TR현물_차트_틱] = 200
	코드별_10분당_TR전송_제한[TR현물_차트_분] = 200
	코드별_10분당_TR전송_제한[TR현물_차트_일주월] = 200
	코드별_10분당_TR전송_제한[TR증시_주변_자금_추이] = 200

	for TR코드, 매10분당_제한_횟수 := range 코드별_10분당_TR전송_제한 {
		tr코드별_10분당_전송_제한[TR코드] = lib.New전송_권한(TR코드, 매10분당_제한_횟수, lib.P10분)
	}

	코드별_초당_TR전송_제한 := make(map[string]int)
	코드별_초당_TR전송_제한[TR현물_정상_주문] = 30
	코드별_초당_TR전송_제한[TR현물_정정_주문] = 30
	코드별_초당_TR전송_제한[TR현물_취소_주문] = 30
	코드별_초당_TR전송_제한[TR시간_조회] = 5
	코드별_초당_TR전송_제한[TR현물_호가_조회] = 5
	코드별_초당_TR전송_제한[TR현물_시세_조회] = 5
	코드별_초당_TR전송_제한[TR현물_기간별_조회] = 1
	코드별_초당_TR전송_제한[TR현물_당일_전일_분틱_조회] = 1
	코드별_초당_TR전송_제한[TR_ETF_시간별_추이] = 1
	코드별_초당_TR전송_제한[TR기업정보_요약] = 1
	코드별_초당_TR전송_제한[TR재무순위_종합] = 1
	코드별_초당_TR전송_제한[TR현물_차트_틱] = 1
	코드별_초당_TR전송_제한[TR현물_차트_분] = 1
	코드별_초당_TR전송_제한[TR현물_차트_일주월] = 1
	코드별_초당_TR전송_제한[TR증시_주변_자금_추이] = 1
	코드별_초당_TR전송_제한[TR현물_종목_조회] = 2

	for TR코드, 초당_제한_횟수 := range 코드별_초당_TR전송_제한 {
		tr코드별_초당_전송_제한[TR코드] = lib.New전송_권한(TR코드, 초당_제한_횟수, lib.P1초)
	}
}
