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

	응답값_모음 = make([]*S_ETF시간별_추이_응답_반복값, 0)
	연속키 := ""

	defer func() { // 순서 거꾸로 뒤집고, 종목코드 정보 및 누락된 시각 데이터 추가.
		nil시각 := time.Time{}
		수량 := len(응답값_모음)
		응답값_모음_임시 := 응답값_모음

		응답값_모음 = make([]*S_ETF시간별_추이_응답_반복값, 수량)

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
		질의값.M구분 = TR조회
		질의값.M코드 = TR_ETF_시간별_추이_t1902
		질의값.M종목코드 = 종목코드
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*S_ETF시간별_추이_응답)
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
	M누적_거래량     int64
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

func NewETF시간별_추이_응답_반복값_모음(b []byte) (값_모음 *S_ETF시간별_추이_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeT1902OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT1902OutBlock1
	g_모음 := make([]*T1902OutBlock1, 수량, 수량)

	값_모음 = new(S_ETF시간별_추이_응답_반복값_모음)
	값_모음.M배열 = make([]*S_ETF시간별_추이_응답_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T1902OutBlock1)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		값 := new(S_ETF시간별_추이_응답_반복값)

		if 값.M시각, 에러 = lib.F2일자별_시각(당일.G값(), "15:04:05", g.Time); 에러 != nil {
			값.M시각 = time.Time{} // ETF_시간별_추이_t1902() 에서 수정
		}
		값.M현재가 = lib.F2정수64_단순형(g.Price)
		값.M전일대비구분 = T전일대비_구분(lib.F2정수_단순형(g.Sign))
		값.M전일대비등락폭 = 값.M전일대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.Change))
		값.M누적_거래량 = lib.F2정수64_단순형(g.Volume)
		값.M현재가_NAV_차이 = lib.F2실수_소숫점_추가_단순형(g.NavDiff, 2)
		값.NAV = lib.F2실수_소숫점_추가_단순형(g.Nav, 2)
		값.NAV전일대비등락폭 = lib.F2실수_소숫점_추가_단순형(g.NavChange, 2)
		값.M추적오차 = lib.F2실수_소숫점_추가_단순형(g.Crate, 2)
		값.M괴리율 = lib.F2실수_소숫점_추가_단순형(g.Grate, 2)
		값.M지수 = lib.F2실수_소숫점_추가_단순형(g.Jisu, 2)
		값.M지수_전일대비등락폭 = lib.F2실수_소숫점_추가_단순형(g.JiChange, 2)
		값.M지수_전일대비등락율 = lib.F2실수_소숫점_추가_단순형(g.JiRate, 2)

		if uint8(g.X_jichange) == 160 && 값.M지수_전일대비등락폭 > 0 {
			값.M지수_전일대비등락폭 = -1 * 값.M지수_전일대비등락폭
		}

		값_모음.M배열[i] = 값
	}

	return 값_모음, nil
}
