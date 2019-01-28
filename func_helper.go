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
	"github.com/mitchellh/go-ps"

	"strings"
	"time"
)

func F질의(질의값 lib.I질의값, 옵션_모음 ...interface{}) (값 *lib.S바이트_변환_모음) {
	defer lib.S예외처리{M함수with내역: func(r interface{}) { 값 = lib.New바이트_변환_모음_단순형(lib.MsgPack, r) }}.S실행()

	lib.F확인(F질의값_종목코드_검사(질의값))

	소켓REQ := 소켓REQ_저장소.G소켓()
	defer 소켓REQ_저장소.S회수(소켓REQ)

	if len(옵션_모음) > 0 {
		소켓REQ.S옵션(옵션_모음...)
	}

	return 소켓REQ.G질의_응답_검사(lib.P변환형식_기본값, 질의값)
}

func F질의_단일TR(질의값 lib.I질의값, 옵션_모음 ...interface{}) (값 interface{}) {
	defer lib.S예외처리{M함수with내역: func(r interface{}) { 값 = lib.New에러(r) }}.S실행()

	i식별번호 := F질의(질의값, 옵션_모음...).G해석값_단순형(0)
	식별번호, ok := i식별번호.(int)
	lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T', '%v'\n"+
		"Xing API에서 식별번호를 부여받고, 콜백을 통해서 응답이 있는 경우에만 사용할 것.\n"+
		"그렇지 않은 경우에는 F질의()를 사용할 것.", i식별번호, i식별번호)

	ch회신 := 대기소_C32.S추가(식별번호, 질의값.TR코드())

	타임아웃 := lib.P1분

	for _, 옵션 := range 옵션_모음 {
		switch 변환값 := 옵션.(type) {
		case time.Duration:
			타임아웃 = 변환값
		}
	}

	select {
	case 값 = <-ch회신:
		return 값
	case <-time.After(타임아웃):
		return lib.New에러("타임아웃. 식별번호 : '%v'", 식별번호)
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

func F_10분_쿼터_잔여량(TR코드_모음 []string) (잔여량_모음 []int, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 잔여량_모음 = nil }}.S실행()

	for _, TR코드 := range TR코드_모음 {
		if !f처리_가능한_TR코드(TR코드) {
			return nil, lib.New에러with출력("잘못된 TR코드 : '%v'", TR코드)
		}
	}

	질의값 := lib.New질의값_문자열_모음(TR_10분_쿼터_잔여량, "" , TR코드_모음)

	lib.F메모("C32에서 원인을 알 수 없는 에러가 발생하며, nil 응답값 수신됨.")

	var 응답값 *lib.S바이트_변환_모음

	for {
		응답값 = F질의(질의값, lib.P5초)

		if 응답값 == nil {
			C32_재시작()
			continue
		}

		break
	}

	lib.F확인(응답값.G값(0, &잔여량_모음))

	return 잔여량_모음, nil
}

func f전일_당일_설정() (에러 error) {
	const 수량 = 30

	질의값_기간별_조회 := New질의값_현물_기간별_조회()
	질의값_기간별_조회.M구분 = TR조회
	질의값_기간별_조회.M코드 = TR현물_기간별_조회
	질의값_기간별_조회.M종목코드 = "069500"
	질의값_기간별_조회.M일주월_구분 = P일주월_일
	질의값_기간별_조회.M연속키 = ""
	질의값_기간별_조회.M수량 = 수량

	i응답값 := F질의_단일TR(질의값_기간별_조회)

	switch 응답값 := i응답값.(type) {
	case error:
		return 응답값
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

func f전일_당일_전달() (에러 error) {
	질의값_전일_당일 := lib.New질의값_바이트_변환_모음(TR전일_당일, "", lib.F금일(), F전일(), F당일())
	바이트_변환_모음 := F질의(질의값_전일_당일)
	신호 := 바이트_변환_모음.G해석값_단순형(0).(lib.T신호)
	lib.F조건부_패닉(신호 != lib.P신호_OK, "예상하지 못한 응답값 : '%v'", 신호)

	return nil
}

func F전일() time.Time {
	return 전일.G값()
}

func F당일() time.Time {
	return 당일.G값()
}

func F최근_영업일_모음() []time.Time {
	return  lib.F슬라이스_복사(최근_영업일_모음, nil).([]time.Time)
}

func F2전일_시각(포맷 string, 값 interface{}) (time.Time, error) {
	if strings.Contains(포맷, "2") {
		return time.Time{}, lib.New에러("포맷에 이미 날짜가 포함되어 있습니다. %v", 포맷)
	}

	시각, 에러 := lib.F2포맷된_시각(포맷, 값)
	if 에러 != nil {
		return time.Time{}, 에러
	}

	전일 := F전일()

	전일_시각 := time.Date(전일.Year(), 전일.Month(), 전일.Day(),
		시각.Hour(), 시각.Minute(), 시각.Second(), 0, 시각.Location())

	return 전일_시각, nil
}

func F2전일_시각_단순형(포맷 string, 값 interface{}) time.Time {
	return lib.F확인(F2전일_시각(포맷, 값)).(time.Time)
}

func F2당일_시각(포맷 string, 값 interface{}) (time.Time, error) {
	if strings.Contains(포맷, "2") {
		return time.Time{}, lib.New에러("포맷에 이미 날짜가 포함되어 있습니다. %v", 포맷)
	}

	시각, 에러 := lib.F2포맷된_시각(포맷, 값)
	if 에러 != nil {
		return time.Time{}, 에러
	}

	당일 := F당일()

	당일_시각 := time.Date(당일.Year(), 당일.Month(), 당일.Day(),
		시각.Hour(), 시각.Minute(), 시각.Second(), 0, 시각.Location())

	return 당일_시각, nil
}

func F2당일_시각_단순형(포맷 string, 값 interface{}) time.Time {
	return lib.F확인(F2당일_시각(포맷, 값)).(time.Time)
}

func xing_C32_실행_중() (프로세스ID int) {
	defer lib.S예외처리{M함수: func() { 프로세스ID = -1 }}.S실행()

	프로세스_모음, 에러 := ps.Processes()
	lib.F확인(에러)

	for _, 프로세스 := range 프로세스_모음 {
		if 실행화일명 := 프로세스.Executable(); strings.HasSuffix(xing_C32_경로, 실행화일명) {
			return 프로세스.Pid()
		}
	}

	return -1
}

//func xing_COM32_실행_중() (프로세스ID int) {
//	defer lib.S예외처리{M함수: func() { 프로세스ID = -1 }}.S실행()
//
//	프로세스_모음, 에러 := ps.Processes()
//	lib.F에러체크(에러)
//
//	for _, 프로세스 := range 프로세스_모음 {
//		if 실행화일명 := 프로세스.Executable(); strings.HasSuffix(xing_COM32_경로, 실행화일명) {
//			return 프로세스.Pid()
//		}
//	}
//
//	return -1
//}

func f접속유지_실행() {
	if !lib.F인터넷에_접속됨() {
		return
	}

	if 접속유지_실행중.G값() {
		return
	} else if 에러 := 접속유지_실행중.S값(true); 에러 != nil {
		return
	}

	go f접속유지_도우미()
}

func f접속유지_도우미() {
	defer 접속유지_실행중.S값(false)

	ch종료 := lib.F공통_종료_채널()
	정기점검 := time.NewTicker(13 * lib.P1초)

	for {
		select {
		case <-정기점검.C:
			F시각_조회_t0167() // t0167 활용
		case <-ch종료:
			정기점검.Stop()
			return
		}
	}
}

func f콜백_데이터_식별번호(값 I콜백) (식별번호 int, 대기_항목 *대기_항목_C32, TR코드 string) {
	switch 변환값 := 값.(type) {
	case *S콜백_TR데이터:
		식별번호 = 변환값.M식별번호
	case *S콜백_메시지_및_에러:
		식별번호 = 변환값.M식별번호
	case *S콜백_정수값:
		식별번호 = 변환값.M정수값
	default:
		panic(lib.New에러("예상하지 못한 경우. 콜백 구분 : '%v', 자료형 : '%T'", 값.G콜백(), 값))
	}

	대기_항목 = 대기소_C32.G값(식별번호)

	if 대기_항목 != nil {
		TR코드 = 대기_항목.TR코드
	}

	return 식별번호, 대기_항목, TR코드
}

func f에러_발생(TR코드, 코드, 내용 string) bool {
	switch TR코드 {
	case TR현물_정상_주문,
		TR시간_조회,
		TR현물_호가_조회,
		TR현물_시세_조회,
		TR현물_기간별_조회,
		TR현물_당일_전일_분틱_조회,
		TR_ETF_시간별_추이,
		TR현물_차트_틱,
		TR현물_차트_분,
		TR현물_차트_일주월,
		TR증시_주변_자금_추이,
		TR현물_종목_조회:
		return 코드 != "00000"
	case TR현물_정정_주문:
		return 코드 != "00131"
	case TR현물_취소_주문:
		return 코드 != "00156"
	default:
		// 에러 출력 지우지 말 것.
		panic(lib.New에러with출력("판별 불가능한 TR코드 : '%v'\n코드 : '%v'\n내용 : '%v'", TR코드, 코드, 내용))
	}
}

func f데이터_복원(대기_항목 *대기_항목_C32, 수신값 *lib.S바이트_변환) error {
	switch 대기_항목.TR코드 {
	case TR시간_조회, TR현물_호가_조회, TR현물_시세_조회, TR_ETF_시세_조회:
		대기_항목.대기값 = 수신값.S해석기(F바이트_변환값_해석).G해석값_단순형() // 단순 질의
		대기_항목.데이터_수신 = true
	case TR현물_정상_주문, TR현물_정정_주문, TR현물_취소_주문, TR기업정보_요약:
		return f데이터_복원_이중_응답(대기_항목, 수신값) // 이중 응답 질의
	case TR현물_기간별_조회, TR현물_당일_전일_분틱_조회, TR_ETF_시간별_추이,
		TR현물_차트_틱, TR현물_차트_분, TR현물_차트_일주월, TR증시_주변_자금_추이:
		return f데이터_복원_반복_조회(대기_항목, 수신값) // 반복 조회
	case TR현물_종목_조회:
		대기_항목.대기값 = 수신값.S해석기(F바이트_변환값_해석).G해석값_단순형()
		대기_항목.데이터_수신 = true
	default:
		return lib.New에러("구현되지 않은 TR코드. %v", 대기_항목.TR코드)
	}

	return nil
}

func f데이터_복원_이중_응답(대기_항목 *대기_항목_C32, 수신값 *lib.S바이트_변환) (에러 error) {
	완전값 := new(S이중_응답_일반형)

	if 대기_항목.대기값 != nil {
		대기값 := 대기_항목.대기값.(I이중_응답)

		if 대기값.G응답1() != nil {
			완전값.M응답1 = 대기값.G응답1()
		}

		if 대기값.G응답2() != nil {
			완전값.M응답2 = 대기값.G응답2()
		}
	}

	switch 변환값 := 수신값.S해석기(F바이트_변환값_해석).G해석값_단순형().(type) {
	case error:
		return 변환값
	case I이중_응답:
		if 변환값.G응답1() != nil {
			완전값.M응답1 = 변환값.G응답1()
		}

		if 변환값.G응답2() != nil {
			완전값.M응답2 = 변환값.G응답2()
		}
	case I이중_응답1:
		완전값.M응답1 = 변환값
	case I이중_응답2:
		완전값.M응답2 = 변환값
	default:
		panic(lib.New에러with출력("예상하지 못한 자료형 문자열 : '%v'", 수신값.G자료형_문자열()))
	}


	대기_항목.대기값 = 완전값

	if 완전값.M응답1 != nil && 완전값.M응답2 != nil {
		대기_항목.데이터_수신 = true
	} else {
		대기_항목.데이터_수신 = false
	}

	return nil
}

func f데이터_복원_반복_조회(대기_항목 *대기_항목_C32, 수신값 *lib.S바이트_변환) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	완전값 := new(S헤더_반복값_일반형)

	if 대기_항목.대기값 != nil {
		대기값 := 대기_항목.대기값.(I헤더_반복값_TR데이터)

		if 대기값.G헤더_TR데이터() != nil {
			완전값.M헤더 = 대기값.G헤더_TR데이터()
		}

		if 대기값.G반복값_모음_TR데이터() != nil {
			완전값.M반복값_모음 = 대기값.G반복값_모음_TR데이터()
		}
	}

	switch 변환값 := 수신값.S해석기(F바이트_변환값_해석).G해석값_단순형().(type) {
	default:
		panic(lib.New에러with출력("예상하지 못한 자료형 : '%v' '%v'", 변환값, 수신값.G자료형_문자열()))
	case error:
		lib.F에러_출력(변환값.Error())
		return 변환값
	case I헤더_반복값_TR데이터:
		if 변환값.G헤더_TR데이터() != nil {
			완전값.M헤더 = 변환값.G헤더_TR데이터()
		}

		if 변환값.G반복값_모음_TR데이터() != nil {
			완전값.M반복값_모음 = 변환값.G반복값_모음_TR데이터()
		}
	case I헤더_TR데이터:
		완전값.M헤더 = 변환값
	case I반복값_모음_TR데이터:
		완전값.M반복값_모음 = 변환값
	}

	대기_항목.대기값 = 완전값

	if 완전값.M헤더 != nil && 완전값.M반복값_모음 != nil {
		대기_항목.데이터_수신 = true
	} else {
		대기_항목.데이터_수신 = false
	}

	//lib.F체크포인트(대기_항목.식별번호, 대기_항목.TR코드, 대기_항목.데이터_수신, 대기_항목.메시지_수신, 대기_항목.응답_완료, 대기_항목.에러)

	return nil
}

func f처리_가능한_TR코드(TR코드 string) bool {
	switch TR코드 {
	case TR시간_조회, TR현물_정상_주문, TR현물_정정_주문, TR현물_취소_주문,
		TR계좌_거래_내역, TR현물_호가_조회, TR현물_시세_조회, //TR현물_시간대별_체결_조회,
		TR현물_기간별_조회, TR현물_당일_전일_분틱_조회, TR_ETF_시세_조회, TR_ETF_시간별_추이,
		TR현물_차트_틱, TR현물_차트_분, TR현물_차트_일주월, TR증시_주변_자금_추이, TR현물_종목_조회:
		return true
	case TR주식_매매일지_수수료_금일, TR주식_매매일지_수수료_날짜_지정, TR주식_잔고_2,
		TR주식_체결_미체결, TR종목별_증시_일정, TR해외_실시간_지수, TR해외_지수_조회,
		TR현물계좌_예수금_주문가능금액_총평가, TR현물계좌_잔고내역, TR현물계좌_주문체결내역,
		TR계좌별_신용한도, TR현물계좌_증거금률별_주문가능수량, TR주식계좌_기간별_수익률_상세:
		lib.F문자열_출력("미구현 TR코드 : '%v'", TR코드)
		return false
	default:
		lib.F문자열_출력("예상하지 못한 TR코드 : '%v'", TR코드)
		return false
	}
}

func C32_재시작() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	lib.F메모("해결하지 못하고 있는 에러. '다른 프로세스가 파일을 사용 중이기 때문에 프로세스가 액세스 할 수 없습니다.'")

	lib.F문자열_출력("** C32 재시작 **")

	lib.F확인(C32_종료())
	lib.F확인(f초기화_xing_C32())

	lib.F체크포인트("")

	lib.F조건부_패닉(!f초기화_작동_확인(), "초기화 작동 확인 실패.")

	lib.F체크포인트()

	lib.F확인(f전일_당일_전달())

	fmt.Println("** C32 재시작 완료     **")

	return nil
}
