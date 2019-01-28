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

	"strings"
	"time"
)




func F현물_호가_조회_t1101(종목코드 string) (응답값 *S현물_호가조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	질의값 := lib.New질의값_단일_종목()
	질의값.M구분 = TR조회
	질의값.M코드 = TR현물_호가_조회
	질의값.M종목코드 = 종목코드

	i응답값 := F질의_단일TR(질의값)

	switch 값 := i응답값.(type) {
	case *S현물_호가조회_응답:
		return 값, nil
	case error:
		return nil, 값
	default:
		panic(lib.New에러("예상하지 못한 자료형 : '%T'", i응답값))
	}
}

func F현물_시세_조회_t1102(종목코드 string) (응답값 *S현물_시세조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	질의값 := lib.New질의값_단일_종목()
	질의값.M구분 = TR조회
	질의값.M코드 = TR현물_시세_조회
	질의값.M종목코드 = 종목코드

	i응답값 := F질의_단일TR(질의값)

	switch 값 := i응답값.(type) {
	case *S현물_시세조회_응답:
		return 값, nil
	case error:
		return nil, 값
	default:
		panic(lib.New에러("예상하지 못한 자료형 : '%T'", i응답값))
	}
}


func F현물_당일전일_분틱_조회_t1310(종목코드 string, 당일전일_구분 T당일전일_구분, 분틱_구분 T분틱_구분,
	종료시각 time.Time, 수량_옵션 ...int) (응답값_모음 []*S현물_전일당일분틱조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	var 수량 int
	if len(수량_옵션) == 1 {
		수량 = 수량_옵션[0]
	}

	lib.F조건부_패닉(당일전일_구분 != P당일전일구분_당일 && 당일전일_구분 != P당일전일구분_전일,
		"예상하지 못한 당일_전일 구분값 : '%v'", 당일전일_구분)

	lib.F조건부_패닉(분틱_구분 != P분틱구분_분 && 분틱_구분 != P분틱구분_틱,
		"예상하지 못한 분_틱 구분값 : '%v'", 분틱_구분)

	응답값_모음_역순 := make([]*S현물_전일당일분틱조회_응답_반복값, 0)
	연속키 := ""

	defer func() {
		수량 = len(응답값_모음_역순)
		응답값_모음 = make([]*S현물_전일당일분틱조회_응답_반복값, len(응답값_모음_역순))

		var 일자 time.Time

		if 당일전일_구분 == P당일전일구분_당일 {
			일자 = F당일()
		} else {
			일자 = F전일()
		}

		// 당일/전일 설정. 시간 기준 정렬순서 변경.
		for i, 응답값 := range 응답값_모음_역순 {
			시각 := 응답값.M시각
			응답값.M시각 = time.Date(일자.Year(), 일자.Month(), 일자.Day(),
				시각.Hour(), 시각.Minute(), 시각.Second(), 시각.Nanosecond(), 시각.Location())

			응답값_모음[수량-1-i] = 응답값
		}
	}()

	for {
		질의값 := New질의값_현물_전일당일_분틱_조회()
		질의값.M구분 = TR조회
		질의값.M코드 = TR현물_당일_전일_분틱_조회
		질의값.M종목코드 = 종목코드
		질의값.M당일전일구분 = 당일전일_구분
		질의값.M분틱구분 = 분틱_구분
		질의값.M종료시각 = 종료시각
		질의값.M연속키 = 연속키

		i응답값 := F질의_단일TR(질의값)

		switch 값 := i응답값.(type) {
		case *S현물_전일당일분틱조회_응답:
			연속키 = 값.M헤더.M연속키
			응답값_모음_역순 = append(응답값_모음_역순, 값.M반복값_모음.M배열...)
		case error:
			//체크("** 에러 발생 **", 값.Error())
			if strings.Contains(값.Error(), "원주문번호를 잘못") ||
				strings.Contains(값.Error(), "접수 대기 상태입니다") {
				//체크("** 예상된 에러 **")
				continue // 재시도
			}

			return nil, 값
		default:
			panic(lib.New에러("예상하지 못한 자료형 : '%T'", i응답값))
		}

		if 수량 > 0 && len(응답값_모음_역순) >= 수량 {
			break
		} else if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func ETF_시간별_추이_t1902(종목코드 string, 추가_옵션_모음 ...interface{}) (응답값_모음 []*S_ETF시간별_추이_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	var 수량 int
	var 시각 time.Time

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

	응답값_모음 = make([]*S_ETF시간별_추이_응답_반복값, 0)
	연속키 := ""

	defer func() {	// 순서 거꾸로 뒤집고, 종목코드 정보 추가.
		수량 := len(응답값_모음)
		응답값_모음_임시 := 응답값_모음

		응답값_모음 = make([]*S_ETF시간별_추이_응답_반복값, 수량)

		for i, 응답값 := range 응답값_모음_임시 {
			응답값.M종목코드 = 종목코드
			응답값_모음[수량-i-1] = 응답값
		}}()

	for {
		질의값 := New질의값_단일종목_연속키()
		질의값.M구분 = TR조회
		질의값.M코드 = TR_ETF_시간별_추이
		질의값.M종목코드 = 종목코드
		질의값.M연속키 = 연속키

		i응답값 := F질의_단일TR(질의값)

		lib.F메모("TR전송 제한 관련 전송 권한 모듈 에러. 3초 대기로 응급처치.")
		lib.F대기(lib.P100밀리초 * 35)

		switch 값 := i응답값.(type) {
		case *S_ETF시간별_추이_응답:
			연속키 = 값.M헤더.M연속키
			응답값_모음 = append(응답값_모음, 값.M반복값_모음.M배열...)
		case error:
			return nil, 값
		default:
			return nil, lib.New에러("예상하지 못한 자료형 : '%T'", i응답값)
		}

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

func F현물_차트_틱_t8411(종목코드 string, 시작일자, 종료일자 time.Time, 추가_인수_모음 ...interface{}) (응답값_모음 []*S현물_차트_틱_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F확인(F종목코드_검사(종목코드))
	lib.F조건부_패닉(종료일자.Before(시작일자), "시작일자가 종료일자보다 늦습니다. %v, %v", 시작일자, 종료일자)

	수량_제한 := -1
	if len(추가_인수_모음) > 0 {
		if 값, ok := 추가_인수_모음[0].(int); ok && 값 > 0 {
			수량_제한 = 값
		}
	}

	응답값_모음 = make([]*S현물_차트_틱_응답_반복값, 0)
	연속일자 := ""
	연속시간 := ""

	defer func() {
		for _, 응답값 := range 응답값_모음 {
			응답값.M종목코드 = 종목코드
		}}()

	for {
		질의값 := New질의값_현물_차트_틱()
		질의값.M구분 = TR조회
		질의값.M코드 = TR현물_차트_틱
		질의값.M종목코드 = 종목코드
		질의값.M단위 = 1
		질의값.M요청건수 = 2000
		질의값.M조회영업일수 = 0
		질의값.M시작일자 = 시작일자
		질의값.M종료일자 = 종료일자
		질의값.M연속일자 = 연속일자
		질의값.M연속시간 = 연속시간
		질의값.M압축여부 = true

		i응답값 := F질의_단일TR(질의값, lib.P30초)

		lib.F메모("TR전송 제한 관련 전송 권한 모듈 에러. 3초 대기로 응급처치.")
		lib.F대기(lib.P100밀리초 * 35)

		switch 값 := i응답값.(type) {
		case *S현물_차트_틱_응답:
			연속일자 = 값.M헤더.M연속일자
			연속시간 = 값.M헤더.M연속시간

			응답값_모음 = append(값.M반복값_모음.M배열, 응답값_모음...)

			if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
				return 응답값_모음, nil
			}
		case error:
			lib.F에러_출력(값)
			return nil, 값
		default:
			return nil, lib.New에러("예상하지 못한 자료형 : '%T'", i응답값)
		}

		if lib.F2문자열_공백제거(연속일자) == "" || lib.F2문자열_공백제거(연속시간) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func F현물_차트_분_t8412(종목코드 string, 시작일자, 종료일자 time.Time, 추가_인수_모음 ...interface{}) (응답값_모음 []*S현물_차트_분_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F확인(F종목코드_검사(종목코드))
	lib.F조건부_패닉(종료일자.Before(시작일자), "시작일자가 종료일자보다 늦습니다. %v, %v", 시작일자, 종료일자)

	수량_제한 := -1
	if len(추가_인수_모음) > 0 {
		if 값, ok := 추가_인수_모음[0].(int); ok && 값 > 0 {
			수량_제한 = 값
		}
	}

	응답값_모음 = make([]*S현물_차트_분_응답_반복값, 0)
	연속일자 := ""
	연속시간 := ""

	defer func() {
		for _, 응답값 := range 응답값_모음 {
			응답값.M종목코드 = 종목코드
		}}()

	for {
		질의값 := New질의값_현물_차트_분()
		질의값.M구분 = TR조회
		질의값.M코드 = TR현물_차트_분
		질의값.M종목코드 = 종목코드
		질의값.M단위 = 1
		질의값.M요청건수 = 2000
		질의값.M조회영업일수 = 0
		질의값.M시작일자 = 시작일자
		질의값.M종료일자 = 종료일자
		질의값.M연속일자 = 연속일자
		질의값.M연속시간 = 연속시간
		질의값.M압축여부 = true

		i응답값 := F질의_단일TR(질의값, lib.P30초)

		lib.F메모("TR전송 제한 관련 전송 권한 모듈 에러. 3초 대기로 응급처치.")
		lib.F대기(lib.P100밀리초 * 35)

		switch 값 := i응답값.(type) {
		case *S현물_차트_분_응답:
			연속일자 = 값.M헤더.M연속일자
			연속시간 = 값.M헤더.M연속시간

			응답값_모음 = append(값.M반복값_모음.M배열, 응답값_모음...)

			if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
				return 응답값_모음, nil
			}
		case error:
			lib.F에러_출력(값)
			return nil, 값
		default:
			return nil, lib.New에러("예상하지 못한 자료형 : '%T', '%v'", i응답값, len(응답값_모음))
		}

		if lib.F2문자열_공백제거(연속일자) == "" || lib.F2문자열_공백제거(연속시간) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func F현물_차트_일주월_t8413(종목코드 string, 시작일자, 종료일자 time.Time, 주기구분 T일주월_구분,
	추가_인수_모음 ...interface{}) (응답값_모음 []*S현물_차트_일주월_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F확인(F종목코드_검사(종목코드))
	lib.F조건부_패닉(종료일자.Before(시작일자), "시작일자가 종료일자보다 늦습니다. %v, %v", 시작일자, 종료일자)

	수량_제한 := -1
	if len(추가_인수_모음) > 0 {
		if 값, ok := 추가_인수_모음[0].(int); ok && 값 > 0 {
			수량_제한 = 값
		}
	}

	응답값_모음 = make([]*S현물_차트_일주월_응답_반복값, 0)
	연속일자 := ""

	defer func() {
		for _, 응답값 := range 응답값_모음 {
			응답값.M종목코드 = 종목코드
		}}()

	for {
		질의값 := New질의값_현물_차트_일주월()
		질의값.M구분 = TR조회
		질의값.M코드 = TR현물_차트_일주월
		질의값.M종목코드 = 종목코드
		질의값.M주기구분 = 주기구분
		질의값.M요청건수 = 2000 // 최대 압축 2000, 비압축 500
		질의값.M시작일자 = 시작일자
		질의값.M종료일자 = 종료일자
		질의값.M연속일자 = 연속일자
		질의값.M압축여부 = true

		i응답값 := F질의_단일TR(질의값, lib.P30초)

		lib.F메모("TR전송 제한 관련 전송 권한 모듈 에러. 3초 대기로 응급처치.")
		lib.F대기(lib.P100밀리초 * 35)

		switch 값 := i응답값.(type) {
		case *S현물_차트_일주월_응답:
			연속일자 = 값.M헤더.M연속일자

			응답값_모음 = append(값.M반복값_모음.M배열, 응답값_모음...)

			if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
				return 응답값_모음, nil
			}
		case error:
			lib.F에러_출력(값)
			return nil, 값
		default:
			return nil, lib.New에러("예상하지 못한 자료형 : '%T'", i응답값)
		}

		if lib.F2문자열_공백제거(연속일자) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

func F증시주변자금추이_t8428(시장_구분 lib.T시장구분, 추가_옵션_모음 ...interface{}) (응답값_모음 []*S증시_주변자금추이_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	lib.F조건부_패닉(
		시장_구분 != lib.P시장구분_코스피 && 시장_구분 != lib.P시장구분_코스닥,
		"예상하지 못한 시장 구분값 : '%v'", 시장_구분)

	var 수량 int
	var 일자 time.Time
	var 연속키 string

	응답값_모음 = make([]*S증시_주변자금추이_응답_반복값, 0)

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

	lib.F메모("TR응답에 연속 조회 추가 존재 여부를 포함시켜서 반복 여부 판단 조건으로 사용하는 것을 생각해 볼 것.")

	for {
		질의값 := New질의값_증시주변자금추이()
		질의값.M구분 = TR조회
		질의값.M코드 = TR증시_주변_자금_추이
		질의값.M시장구분 = 시장_구분
		질의값.M수량 = 200
		질의값.M연속키 = 연속키

		i응답값 := F질의_단일TR(질의값)

		switch 값 := i응답값.(type) {
		case *S증시_주변자금추이_응답:
			연속키 = 값.M헤더.M연속키
			응답값_모음 = append(응답값_모음, 값.M반복값_모음.M배열...)
		case error:
			return nil, 값
		default:
			panic(lib.New에러("예상하지 못한 자료형 : '%T'", i응답값))
		}

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

