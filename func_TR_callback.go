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
)

func f콜백_TR데이터_처리기(값 I콜백) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	식별번호, 대기_항목, TR코드 := f콜백_데이터_식별번호(값)

	//lib.F체크포인트(식별번호, TR코드, 값.G콜백(), 대기_항목.데이터_수신, 대기_항목.메시지_수신, 대기_항목.응답_완료, 대기_항목.에러)

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
		if 에러 = f콜백_데이터_복원(대기_항목, 값.(*S콜백_TR데이터).M데이터); 에러 != nil && 대기_항목.에러 == nil {
			switch {
			case strings.Contains(에러.Error(), "New현물_정정_주문_응답2() : 주문번호 생성 에러"),
				strings.Contains(에러.Error(), "New현물_취소_주문_응답2() : 주문번호 생성 에러"):
				return // skip
			default:
				lib.F에러_출력(에러)
			}

			lib.F체크포인트(식별번호, TR코드, 값.G콜백(), 대기_항목.데이터_수신, 대기_항목.메시지_수신, 대기_항목.응답_완료, 대기_항목.에러)
		}
	case P콜백_메시지_및_에러:
		변환값 := 값.(*S콜백_메시지_및_에러)

		//lib.F체크포인트(변환값.M코드, 변환값.M내용)

		if f에러_발생(TR코드, 변환값.M코드, 변환값.M내용) {
			대기_항목.에러 = lib.New에러("%s : %s", 변환값.M코드, 변환값.M내용)
		}

		대기_항목.메시지_수신 = true
	case P콜백_TR완료:
		대기_항목.응답_완료 = true
	case P콜백_타임아웃:
		대기_항목.에러 = lib.New에러with출력("타임아웃.")
	default:
		panic(lib.New에러with출력("예상하지 못한 경우. 콜백 구분값 : '%v', 자료형 : '%T'", 값.G콜백(), 값))
	}

	//lib.F체크포인트(식별번호, TR코드, 값.G콜백(), 대기_항목.데이터_수신, 대기_항목.메시지_수신, 대기_항목.응답_완료, 대기_항목.에러)

	// TR응답 데이터 수신 및 완료 확인이 되었는 지 확인.
	switch {
	case 대기_항목.에러 != nil && 대기_항목.메시지_수신 && 대기_항목.응답_완료:
		대기소_C32.S회신(식별번호)
	case !대기_항목.데이터_수신, !대기_항목.응답_완료, !대기_항목.메시지_수신:
		return
	default:
		대기소_C32.S회신(식별번호)
	}

	return
}

func f콜백_데이터_식별번호(값 I콜백) (식별번호 int, 대기_항목 *c32_콜백_대기_항목, TR코드 string) {
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

func f콜백_데이터_복원(대기_항목 *c32_콜백_대기_항목, 수신값 *lib.S바이트_변환) error {
	switch 대기_항목.TR코드 {
	case TR시간_조회, TR현물_호가_조회, TR현물_시세_조회, TR_ETF_시세_조회:
		대기_항목.대기값 = 수신값.S해석기(F바이트_변환값_해석).G해석값_단순형() // 단순 질의
		대기_항목.데이터_수신 = true
	case TR현물_정상_주문, TR현물_정정_주문, TR현물_취소_주문, TR기업정보_요약:
		return f데이터_복원_이중_응답(대기_항목, 수신값) // 이중 응답 질의
	case TR재무순위_종합, TR현물_기간별_조회, TR현물_당일_전일_분틱_조회,
		TR_ETF_시간별_추이, TR현물_차트_틱, TR현물_차트_분, TR현물_차트_일주월,
		TR증시_주변_자금_추이:
		return f데이터_복원_반복_조회(대기_항목, 수신값) // 반복 조회
	case TR현물_종목_조회:
		대기_항목.대기값 = 수신값.S해석기(F바이트_변환값_해석).G해석값_단순형()
		대기_항목.데이터_수신 = true
	default:
		return lib.New에러("구현되지 않은 TR코드. %v", 대기_항목.TR코드)
	}

	return nil
}

func f콜백_신호_처리기(콜백 I콜백) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	콜백_정수값, ok := 콜백.(*S콜백_정수값)
	lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", 콜백)

	정수값 := 콜백_정수값.M정수값
	신호 := T신호_C32(정수값)

	//lib.F체크포인트("콜백 신호 수신", 신호)

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
