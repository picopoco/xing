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
)

func F재무_순위_종합_t3341(시장구분 lib.T시장구분, 재무순위_구분 T재무순위_구분,
	추가_인수_모음 ...interface{}) (응답값_모음 []*S재무순위_응답_반복값_t3341, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	switch 시장구분 {
	case lib.P시장구분_전체,
		lib.P시장구분_코스피,
		lib.P시장구분_코스닥: // OK
	default:
		panic(lib.New에러("잘못된 시장구분값 : '%s' '%d'", 시장구분, 시장구분))
	}

	switch 재무순위_구분 {
	case P재무순위_매출액증가율,
		P재무순위_영업이익증가율,
		P재무순위_세전계속이익증가율,
		P재무순위_부채비율,
		P재무순위_유보율,
		P재무순위_EPS,
		P재무순위_BPS,
		P재무순위_ROE,
		P재무순위_PER,
		P재무순위_PBR,
		P재무순위_PEG:
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

	응답값_모음 = make([]*S재무순위_응답_반복값_t3341, 0)
	연속키 := ""

	for {
		질의값 := New질의값_재무순위_t3341()
		질의값.M시장구분 = 시장구분
		질의값.M재무순위_구분 = 재무순위_구분
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		값, ok := i응답값.(*S재무순위_응답_t3341)
		lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키
		응답값_모음 = append(응답값_모음, 값.M반복값_모음.M배열...)

		if 수량_제한 > 0 && len(응답값_모음) > 수량_제한 {
			return 응답값_모음, nil
		}
	}

	return 응답값_모음, nil
}

// HTS 3303 화면
type S질의값_재무순위_t3341 struct {
	*lib.S질의값_기본형
	M시장구분 lib.T시장구분
	M재무순위_구분 T재무순위_구분
	M연속키 string
}

type S재무순위_응답_t3341 struct {
	M헤더     *S재무순위_응답_헤더_t3341
	M반복값_모음 *S재무순위_응답_반복값_모음_t3341
}

type S재무순위_응답_헤더_t3341 struct {
	M수량 int
	M연속키 string
}

func (s *S재무순위_응답_헤더_t3341) G헤더_TR데이터() I헤더_TR데이터 { return s }

// HTS 3303화면과 동일합니다.
// long으로 들어오는 데이터를 소수점 2째자리로 변경하셔야 합니다.
type S재무순위_응답_반복값_t3341 struct {
	M순위 int
	M종목코드 string
	M기업명 string
	M매출액_증가율 float64
	M영업이익_증가율 float64
	M경상이익_증가율 float64
	M부채비율 float64
	M유보율 float64
	EPS float64
	BPS float64
	ROE float64
	PER float64
	PBR float64
	PEG float64
}

type S재무순위_응답_반복값_모음_t3341 struct {
	M배열 []*S재무순위_응답_반복값_t3341
}

func (s *S재무순위_응답_반복값_모음_t3341) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

func New질의값_재무순위_t3341() *S질의값_재무순위_t3341 {
	s := new(S질의값_재무순위_t3341)
	s.S질의값_기본형 = lib.New질의값_기본형(lib.TR조회, TR재무순위_종합)

	return s
}

func NewT3341InBlock(질의값 *S질의값_재무순위_t3341) (g *T3341InBlock) {
	var xing시장구분 string
	var xing재무순위_구분 string

	switch 질의값.M시장구분 {
	case lib.P시장구분_전체:
		xing시장구분 = "0"
	case lib.P시장구분_코스피:
		xing시장구분 = "1"
	case lib.P시장구분_코스닥:
		xing시장구분 = "2"
	default:
		panic(lib.New에러("잘못된 시장구분값 : '%s' '%d'", 질의값.M시장구분, 질의값.M시장구분))
	}

	switch 질의값.M재무순위_구분 {
	case P재무순위_매출액증가율,
		P재무순위_영업이익증가율,
		P재무순위_세전계속이익증가율,
		P재무순위_부채비율,
		P재무순위_유보율,
		P재무순위_EPS,
		P재무순위_BPS,
		P재무순위_ROE,
		P재무순위_PER,
		P재무순위_PBR,
		P재무순위_PEG:
		xing재무순위_구분 = 질의값.M재무순위_구분.T3341()
	default:
		panic(lib.New에러("잘못된 재무순위 구분값 : '%s' '%s'", string(질의값.M재무순위_구분), 질의값.M재무순위_구분.String()))
	}

	g = new(T3341InBlock)
	lib.F바이트_복사_문자열(g.Gubun[:], xing시장구분)
	lib.F바이트_복사_정수(g.Gubun1[:], xing재무순위_구분)
	lib.F바이트_복사_정수(g.Gubun2[:], "1")
	lib.F바이트_복사_정수(g.Idx[:], 질의값.M연속키)

	return g
}

func New재무순위_응답_헤더(b []byte) (값 *S재무순위_응답_헤더_t3341, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT3341OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(T3341OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S재무순위_응답_헤더_t3341)
	값.M수량 = lib.F2정수_단순형(g.Cnt)
	값.M연속키 = lib.F2문자열(g.Idx)

	return 값, nil
}

func New재무순위_응답_반복값(b []byte) (값 *S재무순위_응답_반복값_t3341, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT3341OutBlock1, "예상하지 못한 길이 : '%v", len(b))

	g := new(T3341OutBlock1)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S재무순위_응답_반복값_t3341)
	값.M종목코드 = lib.F2문자열(g.Shcode)
	값.M순위 = lib.F2정수_단순형(g.Rank)
	값.M기업명 = lib.F2문자열(g.Hname)
	값.M매출액_증가율 = lib.F2실수_소숫점_추가(g.Salesgrowth, 2)
	값.M영업이익_증가율 = lib.F2실수_소숫점_추가(g.Operatingincomegrowt, 2)
	값.M경상이익_증가율 = lib.F2실수_소숫점_추가(g.Ordinaryincomegrowth, 2)
	값.M부채비율 = lib.F2실수_소숫점_추가(g.Liabilitytoequity, 2)
	값.M유보율 = lib.F2실수_소숫점_추가(g.Enterpriseratio, 2)
	값.EPS = lib.F2실수_소숫점_추가(g.Eps, 2)
	값.BPS = lib.F2실수_소숫점_추가(g.Bps, 2)
	값.ROE = lib.F2실수_소숫점_추가(g.Roe, 2)
	값.PER = lib.F2실수_소숫점_추가(g.Per, 2)
	값.PBR = lib.F2실수_소숫점_추가(g.Pbr, 2)
	값.PEG = lib.F2실수_소숫점_추가(g.Peg, 2)

	return 값, nil
}