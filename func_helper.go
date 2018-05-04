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
	"github.com/mitchellh/go-ps"

	"strings"
	"time"
)

func F질의(질의값 lib.I질의값, 옵션_모음 ...interface{}) (값 *lib.S바이트_변환_모음) {
	소켓REQ := 소켓REQ_저장소.G소켓()
	defer 소켓REQ_저장소.S회수(소켓REQ)

	if len(옵션_모음) > 0 {
		소켓REQ.S옵션(옵션_모음...)
	}

	return 소켓REQ.G질의_응답_검사(lib.P변환형식_기본값, 질의값)
}

func F접속됨() (접속됨 bool, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 접속됨 = false }}.S실행()

	질의값 := lib.New질의값_기본형(lib.TR접속됨, "")
	접속됨 = 에러체크(F질의(질의값, lib.P10초).G해석값(0)).(bool)

	return 접속됨, nil
}

func F계좌번호_모음() (계좌번호_모음 []string, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 계좌번호_모음 = nil }}.S실행()

	계좌_수량 := 에러체크(F계좌_수량()).(int)
	계좌번호_모음 = make([]string, 계좌_수량)

	for i := 0; i < 계좌_수량; i++ {
		계좌번호_모음[i] = 에러체크(F계좌_번호(0)).(string)
	}

	return 계좌번호_모음, nil
}

func F계좌_수량() (계좌_수량 int, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 계좌_수량 = 0 }}.S실행()

	회신_메시지 := F질의(lib.New질의값_기본형(lib.TR계좌_수량, ""))
	계좌_수량 = 에러체크(회신_메시지.G해석값(0)).(int)
	lib.F조건부_패닉(계좌_수량 == 0, "계좌 수량 0.")

	return 계좌_수량, nil
}

func F계좌_번호(인덱스 int) (계좌_번호 string, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 계좌_번호 = "" }}.S실행()

	회신_메시지 := F질의(lib.New질의값_정수(lib.TR계좌_번호, "", 인덱스))
	계좌_번호 = 에러체크(회신_메시지.G해석값(0)).(string)

	return 계좌_번호, nil
}

func F계좌_이름(계좌_번호 string) (계좌_이름 string, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 계좌_이름 = "" }}.S실행()

	회신_메시지 := F질의(lib.New질의값_문자열(lib.TR계좌_이름, "", 계좌_번호))
	계좌_이름 = 에러체크(회신_메시지.G해석값(0)).(string)

	return 계좌_이름, nil
}

func F계좌_상세명(계좌_번호  string) (계좌_상세명 string, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 계좌_상세명 = "" }}.S실행()

	회신_메시지 := F질의(lib.New질의값_문자열(lib.TR계좌_상세명, "", 계좌_번호))
	계좌_상세명 = 에러체크(회신_메시지.G해석값(0)).(string)

	return 계좌_상세명, nil
}

func f전일_당일_설정() (에러 error) {
	전일_당일_설정_잠금.Lock()
	defer 전일_당일_설정_잠금.Unlock()

	전일_당일_설정_일자값 := 전일_당일_설정_일자.G값()
	금일 := f금일()

	if 전일_당일_설정_일자값.Equal(금일) {
		return
	}

	defer 전일_당일_설정_일자.S값(lib.F2일자(time.Now()))

	질의값 := xt.New질의값_현물_기간별_조회()
	질의값.TR구분 = lib.TR조회
	질의값.TR코드 = xt.TR현물_기간별_조회
	질의값.M종목코드 = "069500"
	질의값.M일주월_구분 = xt.P일주월_일
	질의값.M연속키 = ""
	질의값.M수량 = 2

	응답값 := F질의_단일TR_단순형(질의값).(*xt.S현물_기간별_조회_응답)

	lib.F조건부_패닉(응답값.M헤더.M수량 != 2, "예상하지 못한 수량 : '%v'", 응답값.M헤더.M수량)
	lib.F조건부_패닉(len(응답값.M반복값_모음.M배열) != 2, "예상하지 못한 수량 : '%v'", len(응답값.M반복값_모음.M배열))

	if 응답값.M반복값_모음.M배열[0].M일자.After(응답값.M반복값_모음.M배열[1].M일자) {
		당일.S값(응답값.M반복값_모음.M배열[0].M일자)
		전일.S값(응답값.M반복값_모음.M배열[1].M일자)
	} else {
		당일.S값(응답값.M반복값_모음.M배열[1].M일자)
		전일.S값(응답값.M반복값_모음.M배열[0].M일자)
	}

	return nil
}

func f금일() time.Time {
	return lib.F2일자(time.Now())
}

func F전일() time.Time {
	if !전일_당일_설정_일자.G값().Equal(f금일()) {
		for f전일_당일_설정() != nil {
			lib.F대기(lib.P1초)
		}
	}

	return 전일.G값()
}

func F당일() time.Time {
	if !전일_당일_설정_일자.G값().Equal(f금일()) {
		for f전일_당일_설정() != nil {
			lib.F대기(lib.P1초)
		}
	}

	return 당일.G값()
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
	return 에러체크(F2당일_시각(포맷, 값)).(time.Time)
}

func F2당일_시각_단순형_공백은_초기값(포맷 string, 값 interface{}) time.Time {
	if lib.F2문자열_공백제거(값) == "" {
		return time.Time{}
	}

	return F2당일_시각_단순형(포맷, 값)
}

func etf종목_여부(종목코드 string) bool {
	종목모음_ETF, 에러 := lib.F종목모음_ETF()
	lib.F에러체크(에러)

	for _, 종목 := range 종목모음_ETF {
		if 종목코드 == 종목.G코드() {
			return true
		}
	}

	종목, 에러 := lib.F종목by코드(종목코드)
	lib.F에러체크(에러)

	switch {
	case strings.Contains(종목.G이름(), " ETN"),
		strings.Contains(종목.G이름(), " ETF"):
		return true
	}

	return false
}

func xing_C32_실행_중() (프로세스ID int) {
	defer lib.S에러패닉_처리기{M함수: func() { 프로세스ID = -1 }}.S실행()

	프로세스_모음, 에러 := ps.Processes()
	lib.F에러체크(에러)

	for _, 프로세스 := range 프로세스_모음 {
		if 실행화일명 := 프로세스.Executable(); strings.HasSuffix(xing_C32_경로, 실행화일명) {
			return 프로세스.Pid()
		}
	}

	return -1
}

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
	정기점검 := time.NewTicker(lib.P10초)

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

func f콜백_데이터_식별번호(값 xt.I콜백) (식별번호 int, 대기_항목 *대기_항목_C32, TR코드 string) {
	switch 변환값 := 값.(type) {
	case *xt.S콜백_TR데이터:
		식별번호 = 변환값.M식별번호
	case *xt.S콜백_메시지_및_에러:
		식별번호 = 변환값.M식별번호
	case *xt.S콜백_정수값:
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
	case xt.TR현물_정상주문,
		xt.TR시간_조회,
		xt.TR현물_호가_조회,
		xt.TR현물_시세_조회,
		xt.TR현물_기간별_조회,
		xt.TR현물_당일_전일_분틱_조회,
		xt.TR_ETF_시간별_추이,
		xt.TR증시_주변_자금_추이,
		xt.TR현물_종목_조회:
		return 코드 != "00000"
	case xt.TR현물_정정주문:
		return 코드 != "00131"
	case xt.TR현물_취소주문:
		return 코드 != "00156"
	default:
		panic(lib.New에러("판별 불가능한 TR코드 : '%v'\n코드 : '%v'\n내용 : '%v'", TR코드, 코드, 내용))
	}
}

func f데이터_복원(대기_항목 *대기_항목_C32, 수신값 *lib.S바이트_변환) error {
	switch 대기_항목.TR코드 {
	case xt.TR시간_조회, xt.TR현물_호가_조회, xt.TR현물_시세_조회, xt.TR_ETF_시세_조회:
		대기_항목.대기값 = 수신값.S해석기(xt.F바이트_변환값_해석).G해석값_단순형() // 단순 질의
		대기_항목.데이터_수신 = true
	case xt.TR현물_정상주문, xt.TR현물_정정주문, xt.TR현물_취소주문:
		f데이터_복원_현물_주문(대기_항목, 수신값) // 주문 질의
	case xt.TR현물_시간대별_체결_조회, xt.TR현물_기간별_조회, xt.TR현물_당일_전일_분틱_조회,
		xt.TR_ETF_시간별_추이, xt.TR증시_주변_자금_추이:
		f데이터_복원_반복_조회(대기_항목, 수신값) // 반복 조회
	case xt.TR현물_종목_조회:
		대기_항목.대기값 = 수신값.S해석기(xt.F바이트_변환값_해석).G해석값_단순형()
		대기_항목.데이터_수신 = true
	default:
		return lib.New에러("구현되지 않은 TR코드. %v", 대기_항목.TR코드)
	}

	return nil
}

func f데이터_복원_현물_주문(대기_항목 *대기_항목_C32, 수신값 *lib.S바이트_변환) {
	완전값 := new(xt.S주문_응답_일반형)

	if 대기_항목.대기값 != nil {
		대기값 := 대기_항목.대기값.(xt.I주문_응답)

		if 대기값.G주문_응답1() != nil {
			완전값.M응답1 = 대기값.G주문_응답1()
		}

		if 대기값.G주문_응답2() != nil {
			완전값.M응답2 = 대기값.G주문_응답2()
		}
	}

	switch 변환값 := 수신값.G해석값_단순형().(type) {
	default:
		panic(lib.New에러("예상하지 못한 자료형 문자열 : '%v'", 수신값.G자료형_문자열()))
	case xt.I주문_응답:
		if 변환값.G주문_응답1() != nil {
			완전값.M응답1 = 변환값.G주문_응답1()
		}

		if 변환값.G주문_응답2() != nil {
			완전값.M응답2 = 변환값.G주문_응답2()
		}
	case xt.I주문_응답1:
		완전값.M응답1 = 변환값
	case xt.I주문_응답2:
		완전값.M응답2 = 변환값
	}

	대기_항목.대기값 = 완전값

	if 완전값.M응답1 != nil && 완전값.M응답2 != nil {
		대기_항목.데이터_수신 = true
	} else {
		대기_항목.데이터_수신 = false
	}
}

func f데이터_복원_반복_조회(대기_항목 *대기_항목_C32, 수신값 *lib.S바이트_변환) {
	완전값 := new(xt.S헤더_반복값_일반형)

	if 대기_항목.대기값 != nil {
		대기값 := 대기_항목.대기값.(xt.I헤더_반복값_TR데이터)

		if 대기값.G헤더_TR데이터() != nil {
			완전값.M헤더 = 대기값.G헤더_TR데이터()
		}

		if 대기값.G반복값_모음_TR데이터() != nil {
			완전값.M반복값_모음 = 대기값.G반복값_모음_TR데이터()
		}
	}

	switch 변환값 := 수신값.S해석기(xt.F바이트_변환값_해석).G해석값_단순형().(type) {
	default:
		panic(lib.New에러("예상하지 못한 자료형 : '%v' '%v'", 변환값, 수신값.G자료형_문자열()))
	case xt.I헤더_반복값_TR데이터:
		if 변환값.G헤더_TR데이터() != nil {
			완전값.M헤더 = 변환값.G헤더_TR데이터()
		}

		if 변환값.G반복값_모음_TR데이터() != nil {
			완전값.M반복값_모음 = 변환값.G반복값_모음_TR데이터()
		}
	case xt.I헤더_TR데이터:
		완전값.M헤더 = 변환값
	case xt.I반복값_모음_TR데이터:
		완전값.M반복값_모음 = 변환값
	}

	대기_항목.대기값 = 완전값

	if 완전값.M헤더 != nil && 완전값.M반복값_모음 != nil {
		대기_항목.데이터_수신 = true
	} else {
		대기_항목.데이터_수신 = false
	}
}

func f처리_가능한_TR코드(TR코드 string) bool {
	switch TR코드 {
	case xt.TR시간_조회, xt.TR계좌_번호, xt.TR현물_정상주문, xt.TR현물_정정주문, xt.TR현물_취소주문,
		xt.TR계좌_거래_내역, xt.TR현물_호가_조회, xt.TR현물_시세_조회, xt.TR현물_시간대별_체결_조회,
		xt.TR현물_기간별_조회, xt.TR현물_당일_전일_분틱_조회, xt.TR_ETF_시세_조회, xt.TR_ETF_시간별_추이,
		xt.TR현물_종목_조회:
		return true
	case xt.TR주식_매매일지_수수료_금일, xt.TR주식_매매일지_수수료_날짜_지정, xt.TR주식_잔고_2,
		xt.TR주식_체결_미체결, xt.TR종목별_증시_일정, xt.TR해외_실시간_지수, xt.TR해외_지수_조회,
		xt.TR증시_주변_자금_추이, xt.TR현물계좌_예수금_주문가능금액_총평가, xt.TR현물계좌_잔고내역,
		xt.TR현물계좌_주문체결내역, xt.TR계좌별_신용한도, xt.TR현물계좌_증거금률별_주문가능수량,
		xt.TR주식계좌_기간별_수익률_상세:
		return false
	default:
		return false
	}
}
