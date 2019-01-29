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

func F기간별_주가_조회_t1305(종목코드 string, 일주월_구분 T일주월_구분, 추가_옵션_모음 ...interface{}) (
	응답값_모음 []*S현물_기간별_조회_응답_반복값, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	var 수량 int
	var 일자 time.Time

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

	lib.F조건부_패닉(일주월_구분 != P일주월_일 && 일주월_구분 != P일주월_주 &&
		일주월_구분 != P일주월_월, "예상하지 못한 일주월 구분값 : '%v'", 일주월_구분)

	연속키 := ""
	응답값_모음 = make([]*S현물_기간별_조회_응답_반복값, 0)

	lib.F메모("TR응답에 연속 조회 추가 존재 여부를 포함시켜서 반복 여부 판단 조건으로 사용하는 것을 생각해 볼 것.")

	for {
		질의값 := New질의값_현물_기간별_조회()
		질의값.M구분 = TR조회
		질의값.M코드 = TR현물_기간별_조회
		질의값.M종목코드 = 종목코드
		질의값.M일주월_구분 = 일주월_구분
		질의값.M수량 = 200
		질의값.M연속키 = 연속키

		i응답값 := F질의_단일TR(질의값)

		switch 값 := i응답값.(type) {
		case *S현물_기간별_조회_응답:
			연속키 = 값.M헤더.M연속키
			응답값_모음 = append(응답값_모음, 값.M반복값_모음.M배열...)

			lib.F조건부_패닉(값.M헤더.M수량 != int64(len(값.M반복값_모음.M배열)),
				"반복값 수량 불일치. '%v', '%v'", 값.M헤더.M수량, len(값.M반복값_모음.M배열))
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
		} else if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

// t1305 기간별 주가
type S질의값_현물_기간별_조회 struct {
	*lib.S질의값_단일_종목
	M일주월_구분 T일주월_구분
	M수량     int
	M연속키    string
}

// t1305 현물 기간별 조회 응답
type S현물_기간별_조회_응답 struct {
	M헤더     *S현물_기간별_조회_응답_헤더
	M반복값_모음 *S현물_기간별_조회_응답_반복값_모음
}

func (s *S현물_기간별_조회_응답) G헤더_TR데이터() I헤더_TR데이터 { return s.M헤더 }
func (s *S현물_기간별_조회_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

// t1305 기간별 주가 조회 응답 헤더. 추가 질의값 생성에 사용.
type S현물_기간별_조회_응답_헤더 struct {
	M수량  int64
	M연속키 string
}

func (s *S현물_기간별_조회_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 { return s }

// t1305 기간별 주가 조회 응답 반복값
type S현물_기간별_조회_응답_반복값 struct {
	M종목코드      string
	M일자        time.Time
	M시가        int64
	M고가        int64
	M저가        int64
	M종가        int64
	M전일대비구분    T전일대비_구분
	M전일대비등락폭   int64
	M전일대비등락율   float64
	M시가대비구분    T전일대비_구분
	M시가대비등락폭   int64
	M시가대비등락율   float64
	M고가대비구분    T전일대비_구분
	M고가대비등락폭   int64
	M고가대비등락율   float64
	M저가대비구분    T전일대비_구분
	M저가대비등락폭   int64
	M저가대비등락율   float64
	M누적거래량     int64
	M누적거래대금_백만 int64
	M거래_증가율    float64
	M체결강도      float64
	M소진율       float64
	M회전율       float64
	M외국인_순매수   int64
	M기관_순매수    int64
	M개인_순매수    int64
	M시가총액_백만   int64
}

type S현물_기간별_조회_응답_반복값_모음 struct {
	M배열 []*S현물_기간별_조회_응답_반복값
}

func (s *S현물_기간별_조회_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

func New질의값_현물_기간별_조회() *S질의값_현물_기간별_조회 {
	s := new(S질의값_현물_기간별_조회)
	s.S질의값_단일_종목 = lib.New질의값_단일_종목()

	return s
}

func NewT1305InBlock(질의값 *S질의값_현물_기간별_조회) (g *T1305InBlock) {
	g = new(T1305InBlock)
	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lib.F바이트_복사_문자열(g.Dwmcode[:], lib.F2문자열(uint8(질의값.M일주월_구분)))
	lib.F바이트_복사_문자열(g.Date[:], 질의값.M연속키)
	lib.F바이트_복사_문자열(g.Idx[:], "    ") // 정수형인데, 사용안함(Space)으로 표시됨.
	lib.F바이트_복사_정수(g.Cnt[:], 질의값.M수량)

	if lib.F2문자열_공백제거(질의값.M연속키) == "" {
		lib.F바이트_복사_문자열(g.Date[:], "       ")
	}

	return g
}

func New현물_기간별_조회_응답_헤더(b []byte) (값 *S현물_기간별_조회_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT1305OutBlock, "예상하지 못한 길이 : '%v", len(b))

	g := new(T1305OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S현물_기간별_조회_응답_헤더)
	값.M수량 = lib.F2정수64_단순형(g.Cnt)
	값.M연속키 = lib.F2문자열_공백제거(g.Date)

	return 값, nil
}

func New현물_기간별_조회_응답_반복값_모음(b []byte) (값 *S현물_기간별_조회_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT1305OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT1305OutBlock1
	g_모음 := make([]*T1305OutBlock1, 수량, 수량)

	값 = new(S현물_기간별_조회_응답_반복값_모음)
	값.M배열 = make([]*S현물_기간별_조회_응답_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T1305OutBlock1)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		일자_문자열_원본 := lib.F2문자열(g.Date)
		버퍼 := new(bytes.Buffer)
		버퍼.WriteString(일자_문자열_원본[0:4])
		버퍼.WriteString("/")
		버퍼.WriteString(일자_문자열_원본[4:6])
		버퍼.WriteString("/")
		버퍼.WriteString(일자_문자열_원본[6:])
		일자_문자열 := 버퍼.String()

		s := new(S현물_기간별_조회_응답_반복값)
		s.M종목코드 = lib.F2문자열(g.Shcode)
		s.M일자 = lib.F2포맷된_일자_단순형("2006/01/02", 일자_문자열)
		s.M시가 = lib.F2정수64_단순형(g.Open)
		s.M고가 = lib.F2정수64_단순형(g.High)
		s.M저가 = lib.F2정수64_단순형(g.Low)
		s.M종가 = lib.F2정수64_단순형(g.Close)

		if 전일대비_구분값, 에러 := lib.F2정수64(g.Sign); 에러 == nil {
			s.M전일대비구분 = T전일대비_구분(전일대비_구분값)
		} else if lib.F2문자열_공백제거(g.Sign) == "" &&
			lib.F2정수64_단순형(g.Change) == 0 && lib.F2실수_단순형(g.Diff) == 0.0 {
			s.M전일대비구분 = P구분_보합
		} else {
			lib.F문자열_출력("일자 : '%v', 잘못된 전일구분. '%v'", s.M일자, lib.F2문자열(g.Sign))
			s.M전일대비구분 = T전일대비_구분(0)
		}

		s.M전일대비등락폭 = s.M전일대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.Change))
		s.M전일대비등락율 = s.M전일대비구분.G부호보정_실수64(lib.F2실수_단순형(g.Diff))
		s.M시가대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.O_sign))
		s.M시가대비등락폭 = s.M시가대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.O_change))
		s.M시가대비등락율 = s.M시가대비구분.G부호보정_실수64(lib.F2실수_단순형(g.O_diff))
		s.M고가대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.H_sign))
		s.M고가대비등락폭 = s.M고가대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.H_change))
		s.M고가대비등락율 = s.M고가대비구분.G부호보정_실수64(lib.F2실수_단순형(g.H_diff))
		s.M저가대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.L_sign))
		s.M저가대비등락폭 = s.M저가대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.L_change))
		s.M저가대비등락율 = s.M저가대비구분.G부호보정_실수64(lib.F2실수_단순형(g.L_diff))
		s.M누적거래량 = lib.F2정수64_단순형(g.Volume)
		s.M누적거래대금_백만 = lib.F2정수64_단순형(g.Value)
		s.M거래_증가율 = lib.F2실수_단순형(g.Diff_vol)
		s.M체결강도 = lib.F2실수_단순형(g.Chdegree)
		s.M소진율 = lib.F2실수_단순형_공백은_0(g.Sojinrate)
		s.M회전율 = lib.F2실수_단순형(g.Changerate)
		s.M외국인_순매수 = lib.F2정수64_단순형_공백은_0(g.Fpvolume)
		s.M기관_순매수 = lib.F2정수64_단순형_공백은_0(g.Covolume)
		s.M개인_순매수 = lib.F2정수64_단순형_공백은_0(g.Ppvolume)
		s.M시가총액_백만 = lib.F2정수64_단순형(g.Marketcap)

		값.M배열[i] = s
	}

	return 값, nil
}
