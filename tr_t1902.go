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
	"time"
)

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

	defer func() { // 순서 거꾸로 뒤집고, 종목코드 정보 추가.
		수량 := len(응답값_모음)
		응답값_모음_임시 := 응답값_모음

		응답값_모음 = make([]*S_ETF시간별_추이_응답_반복값, 수량)

		for i, 응답값 := range 응답값_모음_임시 {
			응답값.M종목코드 = 종목코드
			응답값_모음[수량-i-1] = 응답값
		}
	}()

	for {
		질의값 := lib.New질의값_단일종목_연속키()
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

// t1902 ETF 시간별 추이 응답

type S_ETF시간별_추이_응답 struct {
	M헤더     *S_ETF시간별_추이_응답_헤더
	M반복값_모음 *S_ETF시간별_추이_응답_반복값_모음
}

func (s *S_ETF시간별_추이_응답) G헤더_TR데이터() I헤더_TR데이터 { return s.M헤더 }

func (s *S_ETF시간별_추이_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

// t1902 ETF 시간별 추이 응답 헤더

type S_ETF시간별_추이_응답_헤더 struct {
	M연속키   string
	M종목명   string
	M업종지수명 string
}

func (s *S_ETF시간별_추이_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 { return s }

// t1902 ETF 시간별 추이 응답 반복값
type S_ETF시간별_추이_응답_반복값 struct {
	M종목코드       string
	M시각         time.Time
	M현재가        int64
	M전일대비구분     T전일대비_구분
	M전일대비등락폭    int64
	M누적_거래량     float64
	M현재가_NAV_차이 float64
	NAV         float64
	NAV전일대비등락폭  float64
	M추적오차       float64
	M괴리율        float64
	M지수         float64
	M지수_전일대비등락폭 float64
	M지수_전일대비등락율 float64
}

type S_ETF시간별_추이_응답_반복값_모음 struct {
	M배열 []*S_ETF시간별_추이_응답_반복값
}

func (s *S_ETF시간별_추이_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

func NewT1902InBlock(질의값 *lib.S질의값_단일종목_연속키) (g *T1902InBlock) {
	g = new(T1902InBlock)
	lib.F바이트_복사_문자열(g.ShCode[:], 질의값.M종목코드)

	if lib.F2문자열_공백제거(질의값.M연속키) == "" {
		lib.F바이트_배열_공백문자열_채움(g.Time)
	} else {
		lib.F바이트_복사_문자열(g.Time[:], 질의값.M연속키)
	}

	return g
}

func NewETF시간별_추이_응답_헤더(b []byte) (s *S_ETF시간별_추이_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT1902OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T1902OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	s = new(S_ETF시간별_추이_응답_헤더)
	s.M연속키 = lib.F2문자열_공백제거(g.Time)
	s.M종목명 = lib.F2문자열_EUC_KR_공백제거(g.HName)
	s.M업종지수명 = lib.F2문자열_EUC_KR_공백제거(g.UpName)

	return s, nil
}

func NewETF시간별_추이_응답_반복값_모음(b []byte) (값 *S_ETF시간별_추이_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT1902OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT1902OutBlock1
	g_모음 := make([]*T1902OutBlock1, 수량, 수량)

	값 = new(S_ETF시간별_추이_응답_반복값_모음)
	값.M배열 = make([]*S_ETF시간별_추이_응답_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T1902OutBlock1)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		s := new(S_ETF시간별_추이_응답_반복값)

		if lib.F2문자열_EUC_KR(g.Time) == "장:마:감" {
			s.M시각 = lib.F2일자별_시각_단순형(당일.G값(), "15:04:05", g_모음[i+1].Time).Add(lib.P10초)
		} else {
			s.M시각 = lib.F2일자별_시각_단순형(당일.G값(), "15:04:05", g.Time)
		}

		s.M현재가 = lib.F2정수64_단순형(g.Price)
		s.M전일대비구분 = T전일대비_구분(lib.F2정수_단순형(g.Sign))
		s.M전일대비등락폭 = s.M전일대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.Change))
		s.M누적_거래량 = lib.F2실수_단순형(g.Volume)
		s.M현재가_NAV_차이 = lib.F2실수_단순형(g.NavDiff)
		s.NAV = lib.F2실수_단순형(g.Nav)
		s.NAV전일대비등락폭 = lib.F2실수_단순형(g.NavChange)
		s.M추적오차 = lib.F2실수_단순형(g.Crate)
		s.M괴리율 = lib.F2실수_단순형(g.Grate)
		s.M지수 = lib.F2실수_단순형(g.Jisu)
		s.M지수_전일대비등락폭 = lib.F2실수_단순형(g.JiChange)
		s.M지수_전일대비등락율 = lib.F2실수_단순형(g.JiRate)

		if uint8(g.X_jichange) == 160 && s.M지수_전일대비등락폭 > 0 {
			s.M지수_전일대비등락폭 = -1 * s.M지수_전일대비등락폭
		}

		값.M배열[i] = s
	}

	return 값, nil
}