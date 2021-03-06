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
	"github.com/ghts/xing_common"
	"github.com/mitchellh/go-ps"

	"fmt"
	"strings"
	"time"
)

func F전일() time.Time {
	return 전일.G값()
}

func F당일() time.Time {
	return 당일.G값()
}

func F최근_영업일_모음() []time.Time {
	return lib.F슬라이스_복사(최근_영업일_모음, nil).([]time.Time)
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
		시각.Hour(), 시각.Minute(), 시각.Second(), 시각.Nanosecond(), 시각.Location())

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
		시각.Hour(), 시각.Minute(), 시각.Second(), 시각.Nanosecond(), 시각.Location())

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

	에러_연속_발생_횟수 := 0
	ch종료 := lib.F공통_종료_채널()

	for {
		lib.F대기(13 * lib.P1초)

		if _, 에러 := (<-TrT0167_시각_조회()).G값(); 에러 == nil {
			에러_연속_발생_횟수 = 0
		} else {
			에러_연속_발생_횟수++
		}

		// 3회 연속 에러 발생하면 API 연결에 문제 있다고 판단하고, C32 API 모듈 재시작.
		if 에러_연속_발생_횟수 >= 3 {
			C32_재시작()
			에러_연속_발생_횟수 = 0
		}

		select {
		case <-ch종료:
			return
		default:
		}
	}
}

func f에러_발생(TR코드, 코드, 내용 string) bool {
	switch TR코드 {
	case xt.TR현물_정상_주문_CSPAT00600,
		xt.TR시간_조회_t0167,
		xt.TR체결_미체결_조회_t0425,
		xt.TR현물_호가_조회_t1101,
		xt.TR현물_시세_조회_t1102,
		xt.TR현물_기간별_조회_t1305,
		xt.TR현물_당일_전일_분틱_조회_t1310,
		xt.TR관리_불성실_투자유의_조회_t1404,
		xt.TR투자경고_매매정지_정리매매_조회_t1405,
		xt.TR_ETF_시간별_추이_t1902,
		xt.TR기업정보_요약_t3320,
		xt.TR재무순위_종합_t3341,
		xt.TR현물_차트_틱_t8411,
		xt.TR현물_차트_분_t8412,
		xt.TR현물_차트_일주월_t8413,
		xt.TR증시_주변_자금_추이_t8428,
		xt.TR현물_종목_조회_t8436:
		return 코드 != "00000"
	case xt.TR현물_정정_주문_CSPAT00700:
		return 코드 != "00131"
	case xt.TR현물_취소_주문_CSPAT00800:
		return 코드 != "00156"
	default:
		// 에러 출력 지우지 말 것.
		panic(lib.New에러with출력("판별 불가능한 TR코드 : '%v'\n코드 : '%v'\n내용 : '%v'", TR코드, 코드, 내용))
	}
}

func f데이터_복원_이중_응답(대기_항목 *c32_콜백_대기_항목, 수신값 *lib.S바이트_변환) (에러 error) {
	완전값 := new(xt.S이중_응답_일반형)

	if 대기_항목.대기값 != nil {
		대기값 := 대기_항목.대기값.(xt.I이중_응답)

		if 대기값.G응답1() != nil {
			완전값.M응답1 = 대기값.G응답1()
		}

		if 대기값.G응답2() != nil {
			완전값.M응답2 = 대기값.G응답2()
		}
	}

	switch 변환값 := 수신값.S해석기(xt.F바이트_변환값_해석).G해석값_단순형().(type) {
	case error:
		return 변환값
	case xt.I이중_응답:
		if 변환값.G응답1() != nil {
			완전값.M응답1 = 변환값.G응답1()
		}

		if 변환값.G응답2() != nil {
			완전값.M응답2 = 변환값.G응답2()
		}
	case xt.I이중_응답1:
		완전값.M응답1 = 변환값
	case xt.I이중_응답2:
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

func f데이터_복원_반복_조회(대기_항목 *c32_콜백_대기_항목, 수신값 *lib.S바이트_변환) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

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
		panic(lib.New에러with출력("예상하지 못한 자료형 : '%T' '%v'", 변환값, 수신값.G자료형_문자열()))
	case error:
		lib.F에러_출력(변환값.Error())
		return 변환값
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

	return nil
}

func f처리_가능한_TR코드(TR코드 string) bool {
	switch TR코드 {
	case xt.TR현물_정상_주문_CSPAT00600, xt.TR현물_정정_주문_CSPAT00700, xt.TR현물_취소_주문_CSPAT00800,
		xt.TR시간_조회_t0167, xt.TR체결_미체결_조회_t0425, xt.TR현물_호가_조회_t1101, xt.TR현물_시세_조회_t1102,
		xt.TR현물_기간별_조회_t1305, xt.TR현물_당일_전일_분틱_조회_t1310,
		xt.TR관리_불성실_투자유의_조회_t1404, xt.TR투자경고_매매정지_정리매매_조회_t1405,
		xt.TR_ETF_시간별_추이_t1902,
		xt.TR기업정보_요약_t3320, xt.TR재무순위_종합_t3341,
		xt.TR현물_차트_틱_t8411, xt.TR현물_차트_분_t8412, xt.TR현물_차트_일주월_t8413,
		xt.TR증시_주변_자금_추이_t8428, xt.TR현물_종목_조회_t8436:
		return true
	default:
		lib.F문자열_출력("예상하지 못한 TR코드 : '%v'", TR코드)
		return false
	}
}

func F접속_확인() error {
	if !접속_여부.G값() {
		return C32_재시작()
	}

	return nil
}

func C32_재시작() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	xing_C32_재실행_잠금.Lock()
	defer xing_C32_재실행_잠금.Unlock()

	if 접속_여부.G값() {
		return // 이미 재시작 되었음. 재접속 필요없음.
	}

	lib.F문자열_출력("** C32 재시작 **")

	lib.F확인(C32_종료())
	lib.F확인(f초기화_xing_C32())
	lib.F조건부_패닉(!f초기화_작동_확인(), "초기화 작동 확인 실패.")
	lib.F확인(f종목모음_설정())
	lib.F확인(f전일_당일_설정())

	fmt.Println("** C32 재시작 완료     **")

	return nil
}

func f전송_권한_획득(TR코드 string) {
	switch TR코드 {
	case "", xt.RT현물_주문_접수_SC0, xt.RT현물_주문_체결_SC1, xt.RT현물_주문_정정_SC2, xt.RT현물_주문_취소_SC3, xt.RT현물_주문_거부_SC4,
		xt.RT코스피_호가_잔량_H1, xt.RT코스피_시간외_호가_잔량_H2, xt.RT코스피_체결_S3, xt.RT코스피_예상_체결_YS3,
		xt.RT코스피_ETF_NAV_I5, xt.RT주식_VI발동해제_VI, xt.RT시간외_단일가VI발동해제_DVI, xt.RT장_운영정보_JIF:
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
