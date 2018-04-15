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

func F질의(질의값 lib.I질의값) (회신_메시지 lib.I소켓_메시지) {
	defer lib.S에러패닉_처리기{M함수with내역: func(r interface{}) { 회신_메시지, _ = lib.New소켓_메시지_에러(r) }}.S실행()

	바이트_변환값, 에러 := lib.New바이트_변환_매개체(lib.P변환형식_기본값, 질의값)
	에러체크(에러)

	호출_인수 := xt.New호출_인수_질의(바이트_변환값)

	return F질의by호출_인수(호출_인수)
}

func F질의by호출_인수(호출_인수 xt.I호출_인수) (회신_메시지 lib.I소켓_메시지) {
	defer lib.S에러패닉_처리기{M함수with내역: func(r interface{}) { 회신_메시지, _ = lib.New소켓_메시지_에러(r) }}.S실행()

	소켓_질의 := lib.New소켓_질의_단순형(lib.P주소_Xing_C함수_호출, lib.P변환형식_기본값, lib.P30초)

	return 소켓_질의.S질의(호출_인수).G응답_검사()
}

func F접속됨() (접속됨 bool, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 접속됨 = false }}.S실행()

	호출_인수 := xt.New호출_인수_기본형(xt.P함수_접속됨)
	회신_메시지 := F질의by호출_인수(호출_인수)
	접속됨 = 에러체크(회신_메시지.G해석값(0)).(bool)

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

	회신_메시지 := F질의by호출_인수(xt.New호출_인수_기본형(xt.P함수_계좌_수량))
	계좌_수량 = 에러체크(회신_메시지.G해석값(0)).(int)
	lib.F조건부_패닉(계좌_수량 == 0, "계좌 수량 0.")

	return 계좌_수량, nil
}

func F계좌_번호(인덱스 int) (계좌_번호 string, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 계좌_번호 = "" }}.S실행()

	회신_메시지 := F질의by호출_인수(xt.New호출_인수_정수값(xt.P함수_계좌_번호, 인덱스))
	계좌_번호 = 에러체크(회신_메시지.G해석값(0)).(string)

	return 계좌_번호, nil
}

func F계좌_이름(인덱스 int) (계좌_이름 string, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 계좌_이름 = "" }}.S실행()

	회신_메시지 := F질의by호출_인수(xt.New호출_인수_정수값(xt.P함수_계좌_이름, 인덱스))
	계좌_이름 = 에러체크(회신_메시지.G해석값(0)).(string)

	return 계좌_이름, nil
}

func F계좌_상세명(인덱스 int) (계좌_상세명 string, 에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러, M함수: func() { 계좌_상세명 = "" }}.S실행()

	회신_메시지 := F질의by호출_인수(xt.New호출_인수_정수값(xt.P함수_계좌_상세명, 인덱스))
	계좌_상세명 = 에러체크(회신_메시지.G해석값(0)).(string)

	return 계좌_상세명, nil
}

func F영업일_기준_전일() time.Time {
	panic("TODO")

	if 영업일_기준_전일.Equal(time.Time{}) {
		f초기화_영업일_기준_전일_당일()
	}

	return 영업일_기준_전일
}

func F영업일_기준_당일() time.Time {
	panic("TODO")

	if 영업일_기준_당일.Equal(time.Time{}) {
		f초기화_영업일_기준_전일_당일()
	}

	return 영업일_기준_당일
}

func F2당일_시각(포맷 string, 값 interface{}) (time.Time, error) {
	if strings.Contains(포맷, "2") {
		return time.Time{}, lib.New에러("포맷에 이미 날짜가 포함되어 있습니다. %v", 포맷)
	}

	시각, 에러 := lib.F2포맷된_시각(포맷, 값)
	if 에러 != nil {
		return time.Time{}, 에러
	}

	당일 := F영업일_기준_당일()

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

func xing_C32_실행_중() (실행_중 bool) {
	defer lib.S에러패닉_처리기{M함수: func() { 실행_중 = false }}.S실행()

	프로세스_모음, 에러 := ps.Processes()
	lib.F에러체크(에러)

	for _, 프로세스 := range 프로세스_모음 {
		if 실행화일명 := 프로세스.Executable(); strings.HasSuffix(xing_C32_경로, 실행화일명) {
			return true
		}
	}

	return false
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

func 바이트_모음_해석(바이트_모음 []byte, 값_포인터 interface{}) error {
	return lib.New소켓_메시지by바이트_모음(바이트_모음).G값(0, 값_포인터)
}
