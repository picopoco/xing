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

type S기업정보_요약_응답 struct {
	M종목코드 string
	M응답1  *S기업정보_요약_응답1
	M응답2  *S기업정보_요약_응답2
}

func (s *S기업정보_요약_응답) G응답1() I이중_응답1 { return s.M응답1 }
func (s *S기업정보_요약_응답) G응답2() I이중_응답2 { return s.M응답2 }

type S기업정보_요약_응답1 struct {
	M업종구분명  string
	M시장구분   string
	M시장구분명  string
	M한글기업명  string
	M본사주소   string
	M본사전화번호 string
	M최근결산년도 string
	M결산월    string
	M최근결산년월 string
	M주당액면가  int64
	M주식수    int64
	M홈페이지   string
	M그룹명    string
	M외국인    float64
	M주담전화   string
	M자본금    float64
	M시가총액   float64
	M배당금    float64
	M배당수익율  float64
	M현재가    int64
	M전일종가   int64
}

func (s *S기업정보_요약_응답1) G응답1() I이중_응답1 { return s }

type S기업정보_요약_응답2 struct {
	M기업코드    string
	M결산년월    string
	M결산구분    string
	PER      float64
	EPS      float64
	PBR      float64
	ROA      float64
	ROE      float64
	EBITDA   float64
	EVEBITDA float64
	M액면가     float64
	SPS      float64
	CPS      float64
	BPS      float64
	T_PER    float64
	T_EPS    float64
	PEG      float64
	T_PEG    float64
	M최근분기년도  string
}

func (s *S기업정보_요약_응답2) G응답2() I이중_응답2 { return s }

func F기업정보_요약_t3320(종목코드 string) (응답값 *S기업정보_요약_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값 = nil }}.S실행()

	질의값 := lib.New질의값_단일_종목()
	질의값.M구분 = TR조회
	질의값.M코드 = TR기업정보_요약
	질의값.M종목코드 = 종목코드

	i응답값 := F질의_단일TR(질의값)

	switch 값 := i응답값.(type) {
	case *S기업정보_요약_응답:
		값.M종목코드 = 종목코드
		return 값, nil
	case error:
		return nil, 값
	default:
		panic(lib.New에러with출력("예상하지 못한 자료형 : '%T'", i응답값))
	}
}

func NewS기업정보_요약_응답1(b []byte) (값 *S기업정보_요약_응답1, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT3320OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T3320OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S기업정보_요약_응답1)
	값.M업종구분명 = lib.F2문자열(g.Upgubunnm)
	값.M시장구분 = lib.F2문자열(g.Sijangcd)
	값.M시장구분명 = lib.F2문자열(g.Marketnm)
	값.M한글기업명 = lib.F2문자열(g.Company)
	값.M본사주소 = lib.F2문자열(g.Baddress)
	값.M본사전화번호 = lib.F2문자열(g.Btelno)
	값.M최근결산년도 = lib.F2문자열(g.Gsyyyy)
	값.M결산월 = lib.F2문자열(g.Gsmm)
	값.M최근결산년월 = lib.F2문자열(g.Gsym)
	값.M주당액면가 = lib.F2정수64_단순형(g.Lstprice)
	값.M주식수 = lib.F2정수64_단순형(g.Gstock)
	값.M홈페이지 = lib.F2문자열(g.Homeurl)
	값.M그룹명 = lib.F2문자열(g.Grdnm)
	값.M외국인 = lib.F2실수_단순형(g.Foreignratio)
	값.M주담전화 = lib.F2문자열(g.Irtel)
	값.M자본금 = lib.F2실수_단순형(g.Capital)
	값.M시가총액 = lib.F2실수_단순형(g.Sigavalue)
	값.M배당금 = lib.F2실수_단순형(g.Cashsis)
	값.M배당수익율 = lib.F2실수_단순형(g.Cashrate)
	값.M현재가 = lib.F2정수64_단순형(g.Price)
	값.M전일종가 = lib.F2정수64_단순형(g.Jnilclose)

	return 값, nil
}

func NewS기업정보_요약_응답2(b []byte) (값 *S기업정보_요약_응답2, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT3320OutBlock1,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T3320OutBlock1)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S기업정보_요약_응답2)
	값.M기업코드 = lib.F2문자열(g.Gicode)
	값.M결산년월 = lib.F2문자열(g.Gsym)
	값.M결산구분 = lib.F2문자열(g.Gsgb)
	값.PER = lib.F2실수_단순형(g.Per)
	값.EPS = lib.F2실수_단순형(g.Eps)
	값.PBR = lib.F2실수_단순형(g.Pbr)
	값.ROA = lib.F2실수_단순형(g.Roa)
	값.ROE = lib.F2실수_단순형(g.Roe)
	값.EBITDA = lib.F2실수_단순형(g.Ebitda)
	값.EVEBITDA = lib.F2실수_단순형(g.Evebitda)
	값.M액면가 = lib.F2실수_단순형(g.Par)
	값.SPS = lib.F2실수_단순형(g.Sps)
	값.CPS = lib.F2실수_단순형(g.Cps)
	값.BPS = lib.F2실수_단순형(g.Bps)
	값.T_PER = lib.F2실수_단순형(g.Tper)
	값.T_EPS = lib.F2실수_단순형(g.Teps)
	값.PEG = lib.F2실수_단순형(g.Peg)
	값.T_PEG = lib.F2실수_단순형(g.Tpeg)
	값.M최근분기년도 = lib.F2문자열(g.Tgsym)

	return 값, nil
}