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
	"bytes"
	"encoding/binary"
	"github.com/ghts/lib"
	"strconv"
	"testing"
	"time"
)

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
		}
	}()

	F접속_확인()

	for {
		질의값 := New질의값_현물_차트_일주월()
		질의값.M구분 = TR조회
		질의값.M코드 = TR현물_차트_일주월_t8413
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

		값, ok := i응답값.(*S현물_차트_일주월_응답)
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

// t8413 현물 차트 일주월
type S질의값_현물_차트_일주월_t8413 struct {
	*lib.S질의값_단일_종목
	M주기구분 T일주월_구분
	M요청건수 int // 최대 압축 2000, 비압축 500
	M시작일자 string
	M종료일자 string
	M연속일자 string
	M압축여부 bool
}

// t8413 현물 차트 일주월 응답
type S현물_차트_일주월_응답 struct {
	M헤더     *S현물_차트_일주월_응답_헤더
	M반복값_모음 *S현물_차트_일주월_응답_반복값_모음
}

func (s *S현물_차트_일주월_응답) G헤더_TR데이터() I헤더_TR데이터 { return s.M헤더 }
func (s *S현물_차트_일주월_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

// t8413 현물 차트 일주월 응답 헤더
type S현물_차트_일주월_응답_헤더 struct {
	M종목코드     string
	M전일시가     int64
	M전일고가     int64
	M전일저가     int64
	M전일종가     int64
	M전일거래량    int64
	M당일시가     int64
	M당일고가     int64
	M당일저가     int64
	M당일종가     int64
	M상한가      int64
	M하한가      int64
	M연속일자     string
	M장시작시간    time.Time
	M장종료시간    time.Time
	M동시호가처리시간 int
	M수량       int64
}

func (s *S현물_차트_일주월_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 { return s }

// t8413 현물 차트 일주월 응답 반복값
type S현물_차트_일주월_응답_반복값 struct {
	M종목코드       string
	M일자         time.Time
	M시가         int64
	M고가         int64
	M저가         int64
	M종가         int64
	M거래량        int64
	M거래대금_백만    int64
	M수정구분       int64
	M수정비율       float64
	M수정주가반영항목   int64
	M수정비율반영거래대금 int64
	M종가등락구분     T전일대비_구분
}

func (s *S현물_차트_일주월_응답_반복값) G수정구분_모음() ([]T수정구분, error) {
	return f2수정구분_모음(s.M수정구분)
}

// t8413 현물 차트 일주월 응답 반복값 모음
type S현물_차트_일주월_응답_반복값_모음 struct {
	M배열 []*S현물_차트_일주월_응답_반복값
}

func (s *S현물_차트_일주월_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

func NewT8413InBlock(질의값 *S질의값_현물_차트_일주월_t8413) (g *T8413InBlock) {
	g = new(T8413InBlock)

	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lib.F바이트_복사_문자열(g.Gubun[:], strconv.Itoa(int(uint8(질의값.M주기구분)+1)))
	lib.F바이트_복사_정수(g.Qrycnt[:], 질의값.M요청건수)
	lib.F바이트_복사_문자열(g.Sdate[:], 질의값.M시작일자)
	lib.F바이트_복사_문자열(g.Edate[:], 질의값.M종료일자)
	lib.F바이트_복사_문자열(g.Cts_date[:], 질의값.M연속일자)
	lib.F바이트_복사_문자열(g.Comp_yn[:], lib.F조건부_문자열(질의값.M압축여부, "Y", "N"))

	f속성값_초기화(g)

	return g
}

func New현물_차트_일주월_응답_헤더(b []byte) (값 *S현물_차트_일주월_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT8413OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T8413OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S현물_차트_일주월_응답_헤더)
	값.M종목코드 = lib.F2문자열(g.Shcode)
	값.M전일시가 = lib.F2정수64_단순형(g.Jisiga)
	값.M전일고가 = lib.F2정수64_단순형(g.Jihigh)
	값.M전일저가 = lib.F2정수64_단순형(g.Jilow)
	값.M전일종가 = lib.F2정수64_단순형(g.Jiclose)
	값.M전일거래량 = lib.F2정수64_단순형(g.Jivolume)
	값.M당일시가 = lib.F2정수64_단순형(g.Disiga)
	값.M당일고가 = lib.F2정수64_단순형(g.Dihigh)
	값.M당일저가 = lib.F2정수64_단순형(g.Dilow)
	값.M당일종가 = lib.F2정수64_단순형(g.Diclose)
	값.M상한가 = lib.F2정수64_단순형(g.Highend)
	값.M하한가 = lib.F2정수64_단순형(g.Lowend)
	값.M연속일자 = lib.F2문자열(g.Cts_date)
	값.M장시작시간 = lib.F2일자별_시각_단순형(당일.G값(), "150405", g.S_time)
	값.M장종료시간 = lib.F2일자별_시각_단순형(당일.G값(), "150405", g.E_time)
	값.M동시호가처리시간 = lib.F2정수_단순형(g.Dshmin)
	값.M수량 = lib.F2정수64_단순형(g.Rec_count)

	return 값, nil
}

func New현물_차트_일주월_응답_반복값_모음(b []byte) (값 *S현물_차트_일주월_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT8413OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT8413OutBlock1
	g_모음 := make([]*T8413OutBlock1, 수량, 수량)

	값 = new(S현물_차트_일주월_응답_반복값_모음)
	값.M배열 = make([]*S현물_차트_일주월_응답_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T8413OutBlock1)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		s := new(S현물_차트_일주월_응답_반복값)
		s.M일자 = lib.F2포맷된_시각_단순형("20060102", lib.F2문자열_공백제거(g.Date))
		s.M시가 = lib.F2정수64_단순형(g.Open)
		s.M고가 = lib.F2정수64_단순형(g.High)
		s.M저가 = lib.F2정수64_단순형(g.Low)
		s.M종가 = lib.F2정수64_단순형(g.Close)
		s.M거래량 = lib.F2정수64_단순형(g.Vol)
		s.M거래대금_백만 = lib.F2정수64_단순형(g.Value)
		s.M수정구분 = lib.F2정수64_단순형_공백은_0(g.Jongchk)
		s.M수정비율 = lib.F2실수_소숫점_추가_단순형_공백은_0(g.Rate, 2)
		s.M수정주가반영항목 = lib.F2정수64_단순형_공백은_0(g.Pricechk)
		s.M수정비율반영거래대금 = lib.F2정수64_단순형_공백은_0(g.Ratevalue)
		s.M종가등락구분 = T전일대비_구분(lib.F2정수_단순형(g.Sign))

		값.M배열[i] = s
	}

	return 값, nil
}

func New질의값_현물_차트_일주월() *S질의값_현물_차트_일주월_t8413 {
	s := new(S질의값_현물_차트_일주월_t8413)
	s.S질의값_단일_종목 = lib.New질의값_단일_종목()

	return s
}

func F테스트_현물_차트_일주월_응답_반복값_t8413(t *testing.T, 값 *S현물_차트_일주월_응답_반복값, 종목코드 string) {
	lib.F테스트_같음(t, 값.M종목코드, 종목코드)
	lib.F테스트_참임(t, 값.M일자.Equal(lib.F금일()) || 값.M일자.Before(lib.F금일()))
	lib.F테스트_참임(t, 값.M고가 >= 값.M시가)
	lib.F테스트_참임(t, 값.M고가 >= 값.M종가)
	lib.F테스트_참임(t, 값.M저가 <= 값.M시가)
	lib.F테스트_참임(t, 값.M저가 <= 값.M종가)
	lib.F테스트_참임(t, 값.M거래량 >= 0, 값.M종목코드, 값.M일자, 값.M거래량)
	lib.F테스트_참임(t, 값.M거래대금_백만 >= 0, 값.M일자, 값.M거래량, 값.M거래대금_백만)
	//lib.F테스트_에러없음(t, lib.F마지막_에러값(값.G수정구분_모음()))	// 수정구분 해석에 에러가 많음.
	lib.F테스트_같음(t, 값.M종가등락구분, P구분_상한, P구분_상승, P구분_보합, P구분_하한, P구분_하락)
}
