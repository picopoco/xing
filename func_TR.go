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

	"strings"
	"time"
)

func TrCSPAT00600_현물_정상주문(질의값 *xt.CSPAT00600_현물_정상_주문_질의값) (응답값 *xt.CSPAT00600_현물_정상_주문_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	F접속_확인()

	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	응답값, ok := i응답값.(*xt.CSPAT00600_현물_정상_주문_응답)
	lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

	return 응답값, nil
}

func TrCSPAT00700_현물_정정주문_(질의값 *xt.CSPAT00700_현물_정정_주문_질의값) (응답값 *xt.CSPAT00700_현물_정정_주문_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	F접속_확인()

	for i := 0; i < 10; i++ { // 최대 10번 재시도
		i응답값, 에러 := F질의_단일TR(질의값)

		if 에러 != nil && (strings.Contains(에러.Error(), "원주문번호를 잘못") ||
			strings.Contains(에러.Error(), "접수 대기 상태입니다")) {
			continue
		}

		lib.F확인(에러)

		응답값, ok := i응답값.(*xt.CSPAT00700_현물_정정_주문_응답)
		lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

		return 응답값, nil
	}

	return nil, lib.New에러("정정 주문 TR 실행 실패.")
}

func TrCSPAT00800_현물_취소주문(질의값 *xt.CSPAT00800_현물_취소_주문_질의값) (응답값 *xt.CSPAT00800_현물_취소_주문_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	F접속_확인()

	for i := 0; i < 10; i++ { // 최대 10번 재시도
		i응답값, 에러 := F질의_단일TR(질의값)

		if 에러 != nil && (strings.Contains(에러.Error(), "원주문번호를 잘못") ||
			strings.Contains(에러.Error(), "접수 대기 상태")) {
			continue // 재시도
		}

		lib.F확인(에러)

		응답값, ok := i응답값.(*xt.CSPAT00800_현물_취소_주문_응답)
		lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

		return 응답값, nil
	}

	return nil, lib.New에러("취소 주문 TR 실행 실패.")
}

func TrT0167_시각_조회() (ch응답 chan *xt.T0167_시각_조회_응답) {

	F접속_확인()

	ch응답 = make(chan *xt.T0167_시각_조회_응답, 1)

	ch질의 <- lib.New작업(f시각_조회_작업, ch응답)

	return ch응답
}

func f시각_조회_작업(인수 interface{}) {
	var 에러 error
	var ch응답 chan *xt.T0167_시각_조회_응답

	defer lib.S예외처리{M에러: &에러, M함수: func() {
		if ch응답 != nil {
			응답값 := new(xt.T0167_시각_조회_응답)
			응답값.M시각 = time.Time{}
			응답값.M에러 = 에러

			ch응답 <- 응답값
		}
	}}.S실행()

	ch응답 = 인수.(chan *xt.T0167_시각_조회_응답)

	질의값 := lib.S질의값_기본형{M구분: xt.TR조회, M코드: xt.TR시간_조회_t0167}
	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	시각값, ok := i응답값.(time.Time)
	lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T", i응답값)

	응답값 := new(xt.T0167_시각_조회_응답)
	응답값.M시각 = 시각값
	응답값.M에러 = nil

	ch응답 <- 응답값
}

func TrT0425_현물_체결_미체결_조회(계좌번호, 비밀번호, 종목코드 string, 체결_구분 lib.T체결_구분,
	매도_매수_구분 lib.T매도_매수_구분) (응답값_모음 []*xt.T0425_현물_체결_미체결_조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F확인(F종목코드_검사(종목코드))

	응답값_모음 = make([]*xt.T0425_현물_체결_미체결_조회_응답_반복값, 0)
	연속키 := ""

	F접속_확인()

	for {
		질의값 := new(xt.T0425_현물_체결_미체결_조회_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR체결_미체결_조회_t0425)
		질의값.M계좌번호 = 계좌번호
		질의값.M비밀번호 = 비밀번호
		질의값.M종목코드 = 종목코드
		질의값.M체결구분 = 체결_구분
		질의값.M매도_매수_구분 = 매도_매수_구분
		질의값.M정렬순서 = xt.P주문번호_순
		//질의값.M정렬순서 = P주문번호_역순
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		if i응답값 == nil {
			break
		}

		값, ok := i응답값.(*xt.T0425_현물_체결_미체결_조회_응답)
		lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키
		응답값_모음 = append(값.M반복값_모음, 응답값_모음...)

		if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1101_현물_호가_조회(종목코드 string) (응답값 *xt.T1101_현물_호가_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	F접속_확인()

	질의값 := lib.New질의값_단일_종목()
	질의값.M구분 = xt.TR조회
	질의값.M코드 = xt.TR현물_호가_조회_t1101
	질의값.M종목코드 = 종목코드

	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	응답값, ok := i응답값.(*xt.T1101_현물_호가_조회_응답)
	lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

	return 응답값, nil
}

func TrT1102_현물_시세_조회(종목코드 string) (응답값 *xt.T1102_현물_시세_조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	F접속_확인()

	질의값 := lib.New질의값_단일_종목()
	질의값.M구분 = xt.TR조회
	질의값.M코드 = xt.TR현물_시세_조회_t1102
	질의값.M종목코드 = 종목코드

	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	응답값, ok := i응답값.(*xt.T1102_현물_시세_조회_응답)
	lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

	return 응답값, nil
}

func TrT1305_기간별_주가_조회(종목코드 string, 일주월_구분 xt.T일주월_구분, 추가_옵션_모음 ...interface{}) (
	응답값_모음 []*xt.T1305_현물_기간별_조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	F접속_확인()

	var 수량 int
	var 일자 time.Time

	for _, 추가_옵션 := range 추가_옵션_모음 {
		switch 변환값 := 추가_옵션.(type) {
		case int:
			수량 = 변환값
		case time.Time:
			일자 = 변환값
		default:
			panic(lib.New에러("예상하지 못한 옵션값 : '%T' '%v'", 추가_옵션, 추가_옵션))
		}
	}

	lib.F조건부_패닉(일주월_구분 != xt.P일주월_일 && 일주월_구분 != xt.P일주월_주 &&
		일주월_구분 != xt.P일주월_월, "예상하지 못한 일주월 구분값 : '%v'", 일주월_구분)

	연속키 := ""
	응답값_모음 = make([]*xt.T1305_현물_기간별_조회_응답_반복값, 0)

	lib.F메모("TR응답에 연속 조회 추가 존재 여부를 포함시켜서 반복 여부 판단 조건으로 사용하는 것을 생각해 볼 것.")

	defer func() { // 순서 거꾸로 뒤집기.
		수량 := len(응답값_모음)
		응답값_모음_임시 := 응답값_모음

		응답값_모음 = make([]*xt.T1305_현물_기간별_조회_응답_반복값, 수량)

		for i, 응답값 := range 응답값_모음_임시 {
			응답값_모음[수량-i-1] = 응답값
		}
	}()

	for {
		질의값 := xt.NewT1305_현물_기간별_조회_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR현물_기간별_조회_t1305
		질의값.M종목코드 = 종목코드
		질의값.M일주월_구분 = 일주월_구분
		질의값.M수량 = 200
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		값, ok := i응답값.(*xt.T1305_현물_기간별_조회_응답)
		lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 값.M반복값_모음.M배열...)

		lib.F조건부_패닉(값.M헤더.M수량 != int64(len(값.M반복값_모음.M배열)),
			"반복값 수량 불일치. '%v', '%v'", 값.M헤더.M수량, len(값.M반복값_모음.M배열))

		if !일자.Equal(time.Time{}) {
			원하는_일자까지_검색 := false
			for _, 응답값 := range 응답값_모음 {
				if 응답값.M일자.Equal(일자) || 응답값.M일자.Before(일자) {
					원하는_일자까지_검색 = true
					break
				}
			}

			if 원하는_일자까지_검색 {
				break
			}
		}

		if 수량 > 0 && len(응답값_모음) >= 수량 {
			break
		} else if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1310_현물_당일전일_분틱_조회(종목코드 string, 당일전일_구분 xt.T당일전일_구분, 분틱_구분 xt.T분틱_구분,
	종료시각 time.Time, 수량_옵션 ...int) (응답값_모음 []*xt.T1310_현물_전일당일분틱조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	F접속_확인()

	var 수량 int
	if len(수량_옵션) == 1 {
		수량 = 수량_옵션[0]
	}

	lib.F조건부_패닉(당일전일_구분 != xt.P당일전일구분_당일 && 당일전일_구분 != xt.P당일전일구분_전일,
		"예상하지 못한 당일_전일 구분값 : '%v'", 당일전일_구분)

	lib.F조건부_패닉(분틱_구분 != xt.P분틱구분_분 && 분틱_구분 != xt.P분틱구분_틱,
		"예상하지 못한 분_틱 구분값 : '%v'", 분틱_구분)

	응답값_모음_역순 := make([]*xt.T1310_현물_전일당일분틱조회_응답_반복값, 0)
	연속키 := ""

	defer func() {
		일자 := lib.F조건부_시간(당일전일_구분 == xt.P당일전일구분_당일, F당일(), F전일())
		수량 = len(응답값_모음_역순)
		응답값_모음 = make([]*xt.T1310_현물_전일당일분틱조회_응답_반복값, len(응답값_모음_역순))

		// 종목코드, 당일/전일 설정. 시간 기준 정렬순서 변경.
		for i, 응답값 := range 응답값_모음_역순 {
			응답값.M종목코드 = 종목코드

			시각 := 응답값.M시각
			응답값.M시각 = time.Date(일자.Year(), 일자.Month(), 일자.Day(),
				시각.Hour(), 시각.Minute(), 시각.Second(), 시각.Nanosecond(), 시각.Location())

			응답값_모음[수량-1-i] = 응답값
		}
	}()

	for {
		질의값 := xt.NewT1310_현물_전일당일_분틱_조회_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR현물_당일_전일_분틱_조회_t1310
		질의값.M종목코드 = 종목코드
		질의값.M당일전일구분 = 당일전일_구분
		질의값.M분틱구분 = 분틱_구분
		질의값.M종료시각 = 종료시각.Format("1504")
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*xt.T1310_현물_전일당일분틱조회_응답)
		lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키
		응답값_모음_역순 = append(응답값_모음_역순, 값.M반복값_모음.M배열...)

		if 수량 > 0 && len(응답값_모음_역순) >= 수량 {
			break
		} else if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1404_관리종목_조회(시장_구분 lib.T시장구분, 관리_질의_구분 xt.T관리_질의_구분) (응답값_모음 []*xt.T1404_관리종목_조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	응답값_모음 = make([]*xt.T1404_관리종목_조회_응답_반복값, 0)
	연속키 := ""

	lib.F확인(F접속_확인())

	for {
		질의값 := new(xt.T1404_관리종목_조회_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR관리_불성실_투자유의_조회_t1404)
		질의값.M시장_구분 = 시장_구분
		질의값.M관리_질의_구분 = 관리_질의_구분
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*xt.T1404_관리종목_조회_응답)
		lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키

		응답값_모음 = append(값.M반복값_모음.M배열, 응답값_모음...)

		if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1405_투자경고_조회(시장_구분 lib.T시장구분, 투자경고_질의_구분 xt.T투자경고_질의_구분) (응답값_모음 []*xt.T1405_투자경고_조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	응답값_모음 = make([]*xt.T1405_투자경고_조회_응답_반복값, 0)
	연속키 := ""

	lib.F확인(F접속_확인())

	for {
		질의값 := new(xt.T1405_투자경고_조회_질의값)
		질의값.S질의값_기본형 = lib.New질의값_기본형(xt.TR조회, xt.TR투자경고_매매정지_정리매매_조회_t1405)
		질의값.M시장_구분 = 시장_구분
		질의값.M투자경고_질의_구분 = 투자경고_질의_구분
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*xt.T1405_투자경고_조회_응답)
		lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

		응답값_모음 = append(값.M반복값_모음.M배열, 응답값_모음...)

		if 연속키 = 값.M헤더.M연속키; lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT1902_ETF_시간별_추이(종목코드 string, 추가_옵션_모음 ...interface{}) (응답값_모음 []*xt.T1902_ETF시간별_추이_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	var 수량 int
	var 시각 time.Time

	F접속_확인()

	for _, 추가_옵션 := range 추가_옵션_모음 {
		switch 변환값 := 추가_옵션.(type) {
		case int:
			수량 = 변환값
		case time.Time:
			시각 = 변환값
		default:
			panic(lib.New에러("예상하지 못한 옵션값 : '%T' '%v'", 추가_옵션, 추가_옵션))
		}
	}

	응답값_모음 = make([]*xt.T1902_ETF시간별_추이_응답_반복값, 0)
	연속키 := ""

	defer func() { // 순서 거꾸로 뒤집고, 종목코드 정보 및 누락된 시각 데이터 추가.
		nil시각 := time.Time{}
		수량 := len(응답값_모음)
		응답값_모음_임시 := 응답값_모음

		응답값_모음 = make([]*xt.T1902_ETF시간별_추이_응답_반복값, 수량)

		for i, 응답값 := range 응답값_모음_임시 {
			if 응답값.M시각.Equal(nil시각) && i != 0 && !응답값_모음_임시[i-1].M시각.Equal(nil시각) {
				응답값.M시각 = 응답값_모음_임시[i-1].M시각.Add(-1 * lib.P10초)
			}

			응답값.M종목코드 = 종목코드
			응답값_모음[수량-i-1] = 응답값
		}

		for i, 응답값 := range 응답값_모음 {
			if 응답값.M시각.Equal(nil시각) && i != 0 && !응답값_모음_임시[i-1].M시각.Equal(nil시각) {
				응답값.M시각 = 응답값_모음[i-1].M시각.Add(lib.P10초)
			}
		}
	}()

	for {
		질의값 := lib.New질의값_단일종목_연속키()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR_ETF_시간별_추이_t1902
		질의값.M종목코드 = 종목코드
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*xt.T1902_ETF시간별_추이_응답)
		lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 값.M반복값_모음.M배열...)

		if !시각.Equal(time.Time{}) {
			원하는_일자까지_검색 := false
			for _, 응답값 := range 응답값_모음 {
				if 응답값.M시각.Equal(시각) || 응답값.M시각.Before(시각) {
					원하는_일자까지_검색 = true
					break
				}
			}

			if 원하는_일자까지_검색 {
				break
			}
		}

		if 수량 > 0 && len(응답값_모음) >= 수량 {
			break
		} else if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

// HTS 3302 화면. t3320 은 참고자료로서 정보의 정확성이나 완전성은 보장하기는 어렵습니다. 숫자 엉망이다.
func TrT3320_F기업정보_요약(종목코드 string) (응답값 *xt.T3320_기업정보_요약_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	F접속_확인()

	질의값 := lib.New질의값_단일_종목()
	질의값.M구분 = xt.TR조회
	질의값.M코드 = xt.TR기업정보_요약_t3320
	질의값.M종목코드 = 종목코드

	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
	// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
	lib.F대기(lib.P3초)

	응답값, ok := i응답값.(*xt.T3320_기업정보_요약_응답)
	lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

	응답값.M종목코드 = 종목코드
	return 응답값, nil
}

// HTS 3303 화면
func TrT3341_재무_순위_종합(시장구분 lib.T시장구분, 재무순위_구분 xt.T재무순위_구분,
	추가_인수_모음 ...interface{}) (응답값_모음 []*xt.T3341_재무순위_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	F접속_확인()

	switch 시장구분 {
	case lib.P시장구분_전체,
		lib.P시장구분_코스피,
		lib.P시장구분_코스닥: // OK
	default:
		panic(lib.New에러("잘못된 시장구분값 : '%s' '%d'", 시장구분, 시장구분))
	}

	switch 재무순위_구분 {
	case xt.P재무순위_매출액증가율,
		xt.P재무순위_영업이익증가율,
		xt.P재무순위_세전계속이익증가율,
		xt.P재무순위_부채비율,
		xt.P재무순위_유보율,
		xt.P재무순위_EPS,
		xt.P재무순위_BPS,
		xt.P재무순위_ROE,
		xt.P재무순위_PER,
		xt.P재무순위_PBR,
		xt.P재무순위_PEG:
		// OK
	default:
		panic(lib.New에러("잘못된 재무순위 구분값 : '%s' '%s'", string(재무순위_구분), 재무순위_구분.String()))
	}

	수량_제한 := -1
	if len(추가_인수_모음) > 0 {
		if 값, ok := 추가_인수_모음[0].(int); ok && 값 > 0 {
			수량_제한 = 값
		}
	}

	응답값_모음 = make([]*xt.T3341_재무순위_응답_반복값, 0)
	연속키 := ""

	for {
		질의값 := xt.NewT3341_재무순위_질의값()
		질의값.M시장구분 = 시장구분
		질의값.M재무순위_구분 = 재무순위_구분
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		값, ok := i응답값.(*xt.T3341_재무순위_응답)
		lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 값.M반복값_모음.M배열...)

		if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
			return 응답값_모음, nil
		}
	}

	return 응답값_모음, nil
}

func TrT8411_현물_차트_틱(종목코드 string, 시작일자, 종료일자 time.Time, 추가_인수_모음 ...interface{}) (응답값_모음 []*xt.T8411_현물_차트_틱_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F확인(F종목코드_검사(종목코드))
	lib.F조건부_패닉(종료일자.Before(시작일자), "시작일자가 종료일자보다 늦습니다. %v, %v", 시작일자, 종료일자)

	수량_제한 := -1
	if len(추가_인수_모음) > 0 {
		if 값, ok := 추가_인수_모음[0].(int); ok && 값 > 0 {
			수량_제한 = 값
		}
	}

	응답값_모음 = make([]*xt.T8411_현물_차트_틱_응답_반복값, 0)
	연속일자 := ""
	연속시간 := ""

	defer func() {
		for _, 응답값 := range 응답값_모음 {
			응답값.M종목코드 = 종목코드
		}
	}()

	F접속_확인()

	for {
		질의값 := xt.NewT8411_현물_차트_틱_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR현물_차트_틱_t8411
		질의값.M종목코드 = 종목코드
		질의값.M단위 = 1
		질의값.M요청건수 = 2000
		질의값.M조회영업일수 = 0
		질의값.M시작일자 = 시작일자.Format("20060102")
		질의값.M종료일자 = 종료일자.Format("20060102")
		질의값.M연속일자 = 연속일자
		질의값.M연속시간 = 연속시간
		질의값.M압축여부 = true

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*xt.T8411_현물_차트_틱_응답)
		lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

		연속일자 = 값.M헤더.M연속일자
		연속시간 = 값.M헤더.M연속시간

		응답값_모음 = append(값.M반복값_모음.M배열, 응답값_모음...)

		if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
			return 응답값_모음, nil
		}

		if lib.F2문자열_공백제거(연속일자) == "" || lib.F2문자열_공백제거(연속시간) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT8412_현물_차트_분(종목코드 string, 시작일자, 종료일자 time.Time, 추가_인수_모음 ...interface{}) (응답값_모음 []*xt.T8412_현물_차트_분_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F확인(F종목코드_검사(종목코드))
	lib.F조건부_패닉(종료일자.Before(시작일자), "시작일자가 종료일자보다 늦습니다. %v, %v", 시작일자, 종료일자)

	수량_제한 := -1
	if len(추가_인수_모음) > 0 {
		if 값, ok := 추가_인수_모음[0].(int); ok && 값 > 0 {
			수량_제한 = 값
		}
	}

	응답값_모음 = make([]*xt.T8412_현물_차트_분_응답_반복값, 0)
	연속일자 := ""
	연속시간 := ""

	defer func() {
		for _, 응답값 := range 응답값_모음 {
			응답값.M종목코드 = 종목코드
		}
	}()

	F접속_확인()

	for {
		질의값 := xt.NewT8412_현물_차트_분_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR현물_차트_분_t8412
		질의값.M종목코드 = 종목코드
		질의값.M단위 = 0 // 0:30초, 1: 1분, 2: 2분, ....., n: n분
		질의값.M요청건수 = 2000
		질의값.M조회영업일수 = 0
		질의값.M시작일자 = 시작일자.Format("20060102")
		질의값.M종료일자 = 종료일자.Format("20060102")
		질의값.M연속일자 = 연속일자
		질의값.M연속시간 = 연속시간
		질의값.M압축여부 = true

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*xt.T8412_현물_차트_분_응답)
		lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T', '%v'", i응답값, len(응답값_모음))

		연속일자 = 값.M헤더.M연속일자
		연속시간 = 값.M헤더.M연속시간

		응답값_모음 = append(값.M반복값_모음.M배열, 응답값_모음...)

		if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
			return 응답값_모음, nil
		}

		if lib.F2문자열_공백제거(연속일자) == "" || lib.F2문자열_공백제거(연속시간) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT8413_현물_차트_일주월(종목코드 string, 시작일자, 종료일자 time.Time, 주기구분 xt.T일주월_구분,
	추가_인수_모음 ...interface{}) (응답값_모음 []*xt.T8413_현물_차트_일주월_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F확인(F종목코드_검사(종목코드))
	lib.F조건부_패닉(종료일자.Before(시작일자), "시작일자가 종료일자보다 늦습니다. %v, %v", 시작일자, 종료일자)

	수량_제한 := -1
	if len(추가_인수_모음) > 0 {
		if 값, ok := 추가_인수_모음[0].(int); ok && 값 > 0 {
			수량_제한 = 값
		}
	}

	응답값_모음 = make([]*xt.T8413_현물_차트_일주월_응답_반복값, 0)
	연속일자 := ""

	defer func() {
		for _, 응답값 := range 응답값_모음 {
			응답값.M종목코드 = 종목코드
		}
	}()

	F접속_확인()

	for {
		질의값 := xt.NewT8413_현물_차트_일주월_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR현물_차트_일주월_t8413
		질의값.M종목코드 = 종목코드
		질의값.M주기구분 = 주기구분
		질의값.M요청건수 = 2000 // 최대 압축 2000, 비압축 500
		질의값.M시작일자 = 시작일자.Format("20060102")
		질의값.M종료일자 = 종료일자.Format("20060102")
		질의값.M연속일자 = 연속일자
		질의값.M압축여부 = true

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*xt.T8413_현물_차트_일주월_응답)
		lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

		연속일자 = 값.M헤더.M연속일자

		응답값_모음 = append(값.M반복값_모음.M배열, 응답값_모음...)

		if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
			return 응답값_모음, nil
		}

		if lib.F2문자열_공백제거(연속일자) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

// HTS 1503 화면
func TrT8428_증시주변자금추이(시장_구분 lib.T시장구분, 추가_옵션_모음 ...interface{}) (응답값_모음 []*xt.T8428_증시_주변자금추이_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F조건부_패닉(
		시장_구분 != lib.P시장구분_코스피 && 시장_구분 != lib.P시장구분_코스닥,
		"예상하지 못한 시장 구분값 : '%v'", 시장_구분)

	var 수량 int
	var 일자 time.Time
	var 연속키 string

	응답값_모음 = make([]*xt.T8428_증시_주변자금추이_응답_반복값, 0)

	for _, 추가_옵션 := range 추가_옵션_모음 {
		switch 변환값 := 추가_옵션.(type) {
		case int:
			수량 = 변환값
		case time.Time:
			일자 = 변환값
		default:
			panic(lib.New에러("예상하지 못한 옵션값 : '%T' '%v'", 추가_옵션, 추가_옵션))
		}
	}

	F접속_확인()

	for {
		질의값 := xt.NewT8428_증시주변자금추이_질의값()
		질의값.M구분 = xt.TR조회
		질의값.M코드 = xt.TR증시_주변_자금_추이_t8428
		질의값.M시장구분 = 시장_구분
		질의값.M수량 = 200
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		값, ok := i응답값.(*xt.T8428_증시_주변자금추이_응답)
		lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 값.M반복값_모음.M배열...)

		if !일자.Equal(time.Time{}) {
			원하는_일자까지_검색 := false
			for _, 응답값 := range 응답값_모음 {
				if 응답값.M일자.Equal(일자) || 응답값.M일자.Before(일자) {
					원하는_일자까지_검색 = true
					break
				}
			}

			if 원하는_일자까지_검색 {
				break
			}
		}

		if 수량 > 0 && len(응답값_모음) >= 수량 {
			break
		} else if len(lib.F정규식_검색(연속키, []string{"[0-9]*"})) < 8 {
			break
		}
	}

	return 응답값_모음, nil
}

func TrT8436_주식종목_조회(시장_구분 lib.T시장구분) (응답값_모음 []*xt.T8436_현물_종목조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	var 시장구분_문자열 string

	switch 시장_구분 {
	case lib.P시장구분_전체:
		시장구분_문자열 = "0"
	case lib.P시장구분_코스피:
		시장구분_문자열 = "1"
	case lib.P시장구분_코스닥:
		시장구분_문자열 = "2"
	default:
		panic(lib.New에러("예상하지 못한 시장 구분값 : '%v'", 시장_구분))
	}

	F접속_확인()

	질의값 := lib.New질의값_문자열(xt.TR조회, xt.TR현물_종목_조회_t8436, 시장구분_문자열)
	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	값, ok := i응답값.(*xt.S현물_종목조회_응답_반복값_모음)
	lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

	return 값.M배열, nil
}

func F질의(질의값 lib.I질의값, 옵션_모음 ...interface{}) (값 *lib.S바이트_변환_모음) {
	var 에러 error

	defer lib.S예외처리{M에러: &에러, M함수: func() {
		값 = lib.New바이트_변환_모음_단순형(lib.MsgPack, 에러)
	}}.S실행()

	lib.F확인(F질의값_종목코드_검사(질의값))

	switch 질의값.TR구분() {
	case xt.TR조회, xt.TR주문:
		f전송_권한_획득(질의값.TR코드())
	}

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

	질의값 := lib.New질의값_기본형(xt.TR접속됨, "")
	접속됨 = lib.F확인(F질의(질의값, lib.P10초).G해석값(0)).(bool)

	return 접속됨, nil
}

func F계좌번호_모음() (응답값 []string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌번호_모음 = nil }}.S실행()

	if len(계좌번호_모음) != 0 {
		return 계좌번호_모음, nil
	}

	질의값 := lib.New질의값_기본형(xt.TR계좌번호_모음, "")

	계좌번호_모음 = make([]string, 0)
	lib.F확인(F질의(질의값, lib.P10초).G값(0, &계좌번호_모음))

	return 계좌번호_모음, nil
}

func F계좌_수량() (계좌_수량 int, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_수량 = 0 }}.S실행()

	회신_메시지 := F질의(lib.New질의값_기본형(xt.TR계좌_수량, ""))
	계좌_수량 = lib.F확인(회신_메시지.G해석값(0)).(int)
	lib.F조건부_패닉(계좌_수량 == 0, "계좌 수량 0.")

	return 계좌_수량, nil
}

func F계좌_번호(인덱스 int) (계좌_번호 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_번호 = "" }}.S실행()

	회신_메시지 := F질의(lib.New질의값_정수(xt.TR계좌번호_모음, "", 인덱스))
	계좌_번호 = lib.F확인(회신_메시지.G해석값(0)).([]string)[인덱스]

	return 계좌_번호, nil
}

func F계좌_이름(계좌_번호 string) (계좌_이름 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_이름 = "" }}.S실행()

	회신_메시지 := F질의(lib.New질의값_문자열(xt.TR계좌_이름, "", 계좌_번호))
	계좌_이름 = lib.F확인(회신_메시지.G해석값(0)).(string)

	return 계좌_이름, nil
}

func F계좌_상세명(계좌_번호 string) (계좌_상세명 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_상세명 = "" }}.S실행()

	회신_메시지 := F질의(lib.New질의값_문자열(xt.TR계좌_상세명, "", 계좌_번호))
	계좌_상세명 = lib.F확인(회신_메시지.G해석값(0)).(string)

	return 계좌_상세명, nil
}

func F계좌_별명(계좌_번호 string) (계좌_별명 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_별명 = "" }}.S실행()

	회신_메시지 := F질의(lib.New질의값_문자열(xt.TR계좌_별명, "", 계좌_번호))
	계좌_별명 = lib.F확인(회신_메시지.G해석값(0)).(string)

	return 계좌_별명, nil
}
