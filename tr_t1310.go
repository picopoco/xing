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
	"strings"
	"time"
)

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
		질의값.M종료시각 = 종료시각.Format("1504")
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)

		if strings.Contains(에러.Error(), "원주문번호를 잘못") ||
			strings.Contains(에러.Error(), "접수 대기 상태입니다") {
			continue // 재시도
		}

		lib.F확인(에러)

		값, ok := i응답값.(*S현물_전일당일분틱조회_응답)
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

// t1310 전일당일분틱조회
type S질의값_현물_전일당일_분틱_조회 struct {
	*lib.S질의값_단일_종목
	M당일전일구분 T당일전일_구분 // 0:당일, 1:전일
	M분틱구분   T분틱_구분   // 0:분, 1:틱
	M종료시각   string   // 해당 시각 이전까지의 데이터만 조회됨.
	M연속키    string   // 처음 조회시 Space. 다음 조회시 t1310OutBlock.cts_time 값 입력
}

// t1310 전일당일분틱조회 응답
type S현물_전일당일분틱조회_응답 struct {
	M헤더     *S현물_전일당일분틱조회_응답_헤더
	M반복값_모음 *S현물_전일당일분틱조회_응답_반복값_모음
}

func (s *S현물_전일당일분틱조회_응답) G헤더_TR데이터() I헤더_TR데이터 {
	return s.M헤더
}
func (s *S현물_전일당일분틱조회_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

// t1310 전일당일분틱조회 응답 헤더
type S현물_전일당일분틱조회_응답_헤더 struct {
	M연속키 string
}

func (s *S현물_전일당일분틱조회_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 {
	return s
}

// t1310 전일당일분틱조회 응답 반복값
type S현물_전일당일분틱조회_응답_반복값 struct {
	M시각      time.Time
	M현재가     int64
	M전일대비구분  T전일대비_구분
	M전일대비등락폭 int64
	M전일대비등락율 float64
	M체결수량    int64
	M체결강도    float64
	M거래량     int64
	M매도체결수량  int64
	M매도체결건수  int64
	M매수체결수량  int64
	M매수체결건수  int64
	M순체결량    int64
	M순체결건수   int64
}

type S현물_전일당일분틱조회_응답_반복값_모음 struct {
	M배열 []*S현물_전일당일분틱조회_응답_반복값
}

func (s *S현물_전일당일분틱조회_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

func NewT1310InBlock(질의값 *S질의값_현물_전일당일_분틱_조회) (g *T1310InBlock) {
	g = new(T1310InBlock)
	lib.F바이트_복사_문자열(g.Daygb[:], strconv.Itoa(int(질의값.M당일전일구분)))
	lib.F바이트_복사_문자열(g.Timegb[:], strconv.Itoa(int(질의값.M분틱구분)))
	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lib.F바이트_복사_문자열(g.Endtime[:], 질의값.M종료시각)
	lib.F바이트_복사_문자열(g.Time[:], 질의값.M연속키)

	if lib.F2문자열_공백제거(질의값.M연속키) == "" {
		lib.F바이트_복사_문자열(g.Time[:], "          ")
	}

	return g
}

func New질의값_현물_전일당일_분틱_조회() *S질의값_현물_전일당일_분틱_조회 {
	s := new(S질의값_현물_전일당일_분틱_조회)
	s.S질의값_단일_종목 = lib.New질의값_단일_종목()

	return s
}

func New현물_당일전일분틱조회_응답_헤더(b []byte) (값 *S현물_전일당일분틱조회_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT1310OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T1310OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S현물_전일당일분틱조회_응답_헤더)
	값.M연속키 = lib.F2문자열(g.Time)

	return 값, nil
}

func New현물_당일전일분틱조회_응답_반복값_모음(b []byte) (값 *S현물_전일당일분틱조회_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT1310OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT1310OutBlock1
	g_모음 := make([]*T1310OutBlock1, 수량, 수량)

	값 = new(S현물_전일당일분틱조회_응답_반복값_모음)
	값.M배열 = make([]*S현물_전일당일분틱조회_응답_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T1310OutBlock1)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		s := new(S현물_전일당일분틱조회_응답_반복값)
		s.M시각 = lib.F2금일_시각_단순형("150405", g.Chetime[:6])
		s.M현재가 = lib.F2정수64_단순형(g.Price)
		s.M전일대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Sign))
		s.M전일대비등락폭 = s.M전일대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.Change))
		s.M전일대비등락율 = s.M전일대비구분.G부호보정_실수64(lib.F2실수_단순형(g.Diff))
		s.M체결수량 = lib.F2정수64_단순형(g.Cvolume)
		s.M체결강도 = lib.F2실수_단순형(g.Chdegree)
		s.M거래량 = lib.F2정수64_단순형(g.Volume)
		s.M매도체결수량 = lib.F2정수64_단순형(g.Mdvolume)
		s.M매도체결건수 = lib.F2정수64_단순형(g.Mdchecnt)
		s.M매수체결수량 = lib.F2정수64_단순형(g.Msvolume)
		s.M매수체결건수 = lib.F2정수64_단순형(g.Mschecnt)
		s.M순체결량 = lib.F2정수64_단순형(g.Revolume)
		s.M순체결건수 = lib.F2정수64_단순형(g.Rechecnt)

		값.M배열[i] = s
	}

	return 값, nil
}
