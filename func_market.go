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
	"time"
)

func F종목코드_모음_전체() []string {
	종목코드_모음 := make([]string, len(종목모음_전체), len(종목모음_전체))

	for i, 종목 := range 종목모음_전체 {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_KOSPI() []string {
	종목코드_모음 := make([]string, len(종목모음_코스피), len(종목모음_코스피))

	for i, 종목 := range 종목모음_코스피 {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_KOSDAQ() []string {
	종목코드_모음 := make([]string, len(종목모음_코스닥), len(종목모음_코스닥))

	for i, 종목 := range 종목모음_코스닥 {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_ETF() []string {
	종목코드_모음 := make([]string, len(종목모음_ETF), len(종목모음_ETF))

	for i, 종목 := range 종목모음_ETF {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_ETN() []string {
	종목코드_모음 := make([]string, len(종목모음_ETN), len(종목모음_ETN))

	for i, 종목 := range 종목모음_ETN {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F종목코드_모음_ETF_ETN() []string {
	종목코드_모음 := make([]string, len(종목모음_ETF_ETN), len(종목모음_ETF_ETN))

	for i, 종목 := range 종목모음_ETF_ETN {
		종목코드_모음[i] = 종목.G코드()
	}

	return 종목코드_모음
}

func F질의값_종목코드_검사(질의값_원본 lib.I질의값) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	switch 질의값 := 질의값_원본.(type) {
	case lib.I종목코드:
		lib.F조건부_패닉(!F종목코드_존재함(질의값.G종목코드()),
			"존재하지 않는 종목코드 : '%v'", 질의값.G종목코드())
	case lib.I종목코드_모음:
		종목코드_모음 := 질의값.G종목코드_모음()

		for _, 종목코드 := range 종목코드_모음 {
			lib.F조건부_패닉(!F종목코드_존재함(종목코드), "존재하지 않는 종목코드 : '%v'", 종목코드)
		}
	}

	return nil
}

func F종목코드_존재함(종목코드 string) bool {
	_, 존재함 := 종목맵_전체[종목코드]

	return 존재함
}

func F종목코드_검사(종목코드 string) error {
	if !F종목코드_존재함(종목코드) {
		return lib.New에러("존재하지 않는 종목코드 : '%s'.", 종목코드)
	}

	return nil
}

func f종목모음_설정() (에러 error) {
	종목모음_설정_잠금.Lock()
	defer 종목모음_설정_잠금.Unlock()

	defer lib.S예외처리{
		M에러: &에러,
		M함수: func() {
			종목모음_코스피 = make([]*lib.S종목, 0)
			종목모음_코스닥 = make([]*lib.S종목, 0)
			종목모음_ETF = make([]*lib.S종목, 0)
			종목모음_ETN = make([]*lib.S종목, 0)
			종목모음_ETF_ETN = make([]*lib.S종목, 0)
			종목모음_전체 = make([]*lib.S종목, 0)
			종목맵_전체 = make(map[string]*lib.S종목)
			기준가_맵 = make(map[string]int64)
			하한가_맵 = make(map[string]int64)
			종목모음_설정일 = lib.New안전한_시각(time.Time{})
		}}.S실행()

	if len(종목모음_코스피) > 0 &&
		len(종목모음_코스닥) > 0 &&
		len(종목모음_ETF) > 0 &&
		len(종목모음_ETN) > 0 &&
		len(종목모음_ETF_ETN) > 0 &&
		len(종목모음_전체) > 0 &&
		len(종목맵_전체) > 0 &&
		len(기준가_맵) > 0 &&
		len(하한가_맵) > 0 &&
		종목모음_설정일.G값().Equal(lib.F금일()) {
		return nil
	}

	종목_정보_모음, 에러 := TrT8436_주식종목_조회(lib.P시장구분_전체)
	lib.F확인(에러)

	종목모음_코스피 = make([]*lib.S종목, 0)
	종목모음_코스닥 = make([]*lib.S종목, 0)
	종목모음_ETF = make([]*lib.S종목, 0)
	종목모음_ETN = make([]*lib.S종목, 0)
	종목모음_ETF_ETN = make([]*lib.S종목, 0)
	종목모음_전체 = make([]*lib.S종목, 0)
	종목맵_전체 = make(map[string]*lib.S종목)
	기준가_맵 = make(map[string]int64)
	하한가_맵 = make(map[string]int64)

	for _, s := range 종목_정보_모음 {
		종목 := lib.New종목(s.M종목코드, s.M종목명, s.M시장구분)

		기준가_맵[s.M종목코드] = s.M기준가
		하한가_맵[s.M종목코드] = s.M하한가
		종목맵_전체[종목.G코드()] = 종목
		종목모음_전체 = append(종목모음_전체, 종목)

		switch s.M시장구분 {
		case lib.P시장구분_코스피:
			종목모음_코스피 = append(종목모음_코스피, 종목)
		case lib.P시장구분_코스닥:
			종목모음_코스닥 = append(종목모음_코스닥, 종목)
		case lib.P시장구분_ETF:
			종목모음_ETF = append(종목모음_ETF, 종목)
			종목모음_ETF_ETN = append(종목모음_ETF_ETN, 종목)
		case lib.P시장구분_ETN:
			종목모음_ETN = append(종목모음_ETN, 종목)
			종목모음_ETF_ETN = append(종목모음_ETF_ETN, 종목)
		}
	}

	종목모음_설정일 = lib.New안전한_시각(lib.F금일())

	return nil
}

func f한국증시_거래시간_도우미(시작_시간, 시작_분, 종료_시간, 종료_분 int) bool {
	값 := 당일.G값()
	지금 := time.Now()
	로케일 := 지금.Location()

	시작_시각 := time.Date(값.Year(), 값.Month(), 값.Day(), 시작_시간, 시작_분, 0, 0, 로케일)
	종료_시각 := time.Date(값.Year(), 값.Month(), 값.Day(), 종료_시간, 종료_분, 0, 0, 로케일)

	if 지금.After(시작_시각) && 지금.Before(종료_시각) {
		return true
	}

	return false
}

func F한국증시_정규시장_거래시간임() bool {
	return f한국증시_거래시간_도우미(9, 0, 15, 30)
}

func F한국증시_정규경쟁대량매매_거래시간임() bool {
	return f한국증시_거래시간_도우미(9, 0, 15, 00)
}

func F한국증시_동시호가_시간임() bool {
	return f한국증시_거래시간_도우미(8, 0, 9, 0) ||
		f한국증시_거래시간_도우미(15, 20, 15, 30)
}

func F한국증시_시간외_종가매매_시간임() bool {
	return f한국증시_거래시간_도우미(15, 40, 16, 0)
}

func F한국증시_시간외_단일가매매_시간임() bool {
	return f한국증시_거래시간_도우미(16, 0, 18, 0)
}

func F한국증시_시간외_대량바스켓매매_거래시간임() bool {
	return f한국증시_거래시간_도우미(15, 40, 18, 0)
}

func F종목by코드(종목코드 string) (종목 *lib.S종목, 에러 error) {
	if 종목, ok := 종목맵_전체[종목코드]; !ok {
		return nil, lib.New에러("해당 종목코드가 존재하지 않습니다. '%v'", 종목코드)
	} else {
		return 종목, nil
	}
}

func F임의_종목() *lib.S종목 {
	return f임의_종목_추출(종목모음_전체)
}

func F임의_종목_코스피_주식() *lib.S종목 {
	return f임의_종목_추출(종목모음_코스피)
}

func F임의_종목_코스닥_주식() *lib.S종목 {
	return f임의_종목_추출(종목모음_코스닥)
}

func F임의_종목_ETF() *lib.S종목 {
	return f임의_종목_추출(종목모음_ETF)
}

func f임의_종목_추출(종목_모음 []*lib.S종목) *lib.S종목 {
	return 종목_모음[lib.F임의_범위_이내_정수값(0, len(종목_모음))].G복제본()
}

func ETF종목_여부(종목_코드 string) bool {
	종목, 에러 := F종목by코드(종목_코드)

	if 에러 != nil && 종목.G시장구분() == lib.P시장구분_ETF {
		return true
	}

	return false
}

func ETN종목_여부(종목_코드 string) bool {
	종목, 에러 := F종목by코드(종목_코드)

	if 에러 != nil && 종목.G시장구분() == lib.P시장구분_ETN {
		return true
	}

	return false
}

func F최소_호가단위by종목코드(종목코드 string) (값 int64, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = 0 }}.S실행()

	종목 := lib.F확인(F종목by코드(종목코드)).(*lib.S종목)

	return F최소_호가단위by종목(종목)
}

func F최소_호가단위by종목(종목 *lib.S종목) (값 int64, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = 0 }}.S실행()

	return F최소_호가단위by시장구분_기준가(종목.G시장구분(), 기준가_맵[종목.G코드()])
}

func F최소_호가단위by시장구분_기준가(시장구분 lib.T시장구분, 기준가 int64) (값 int64, 에러 error) {
	switch 시장구분 {
	case lib.P시장구분_코스피:
		switch {
		case 기준가 < 1000:
			return 1, nil
		case 기준가 >= 1000 && 기준가 < 5000:
			return 5, nil
		case 기준가 >= 5000 && 기준가 < 10000:
			return 10, nil
		case 기준가 >= 10000 && 기준가 < 50000:
			return 50, nil
		case 기준가 >= 50000 && 기준가 < 100000:
			return 100, nil
		case 기준가 >= 100000 && 기준가 < 500000:
			return 500, nil
		case 기준가 >= 500000:
			return 1000, nil
		default:
			panic(lib.New에러with출력("예상하지 못한 경우. %v", 기준가))
		}
	case lib.P시장구분_코스닥:
		switch {
		case 기준가 < 1000:
			return 1, nil
		case 기준가 >= 1000 && 기준가 < 5000:
			return 5, nil
		case 기준가 >= 5000 && 기준가 < 10000:
			return 10, nil
		case 기준가 >= 10000 && 기준가 < 50000:
			return 50, nil
		case 기준가 >= 50000:
			return 100, nil
		default:
			panic(lib.New에러with출력("예상하지 못한 경우. %v", 기준가))
		}
	case lib.P시장구분_코넥스:
		switch {
		case 기준가 < 5000:
			return 5, nil
		case 기준가 >= 5000 && 기준가 < 10000:
			return 10, nil
		case 기준가 >= 10000 && 기준가 < 50000:
			return 50, nil
		case 기준가 >= 50000 && 기준가 < 100000:
			return 100, nil
		case 기준가 >= 100000 && 기준가 < 500000:
			return 500, nil
		case 기준가 >= 500000:
			return 1000, nil
		default:
			panic(lib.New에러with출력("예상하지 못한 경우. %v", 기준가))
		}
	case lib.P시장구분_ETF:
		return 5, nil
	}

	return 0, lib.New에러with출력("예상하지 못한 시장구분. %v", 시장구분)
}
