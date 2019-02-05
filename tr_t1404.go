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

	"time"
)

func F관리종목_조회_t1404(시장_구분 lib.T시장구분, 관리_질의_구분 T관리_질의_구분) (응답값_모음 []*S관리종목_조회_응답_반복값_t1404, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 응답값_모음 = nil }}.S실행()

	응답값_모음 = make([]*S관리종목_조회_응답_반복값_t1404, 0)
	연속키 := ""

	lib.F확인(F접속_확인())

	for {
		질의값 := new(S질의값_관리종목_조회_t1404)
		질의값.S질의값_기본형 = lib.New질의값_기본형(TR조회, TR관리_불성실_투자유의_조회_t1404)
		질의값.M시장_구분 = 시장_구분
		질의값.M관리_질의_구분 = 관리_질의_구분
		질의값.M연속키 = 연속키

		i응답값, 에러 := F질의_단일TR(질의값)
		lib.F확인(에러)

		// TR전송 제한이 걸리면, 타임아웃이 되면서 데이터 수집에 오히려 방해가 됨.
		// TR전송 제한 소모 속도를 늦추어서, 타임아웃이 되지 않게 하는 것이 오히려 도움이 됨.
		lib.F대기(lib.P3초)

		값, ok := i응답값.(*S관리종목_조회_응답_t1404)
		lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T'", i응답값)

		연속키 = 값.M헤더.M연속키

		응답값_모음 = append(값.M반복값_모음.M배열, 응답값_모음...)

		if lib.F2문자열_공백제거(연속키) == "" {
			break
		}
	}

	return 응답값_모음, nil
}

type S질의값_관리종목_조회_t1404 struct {
	*lib.S질의값_기본형
	M시장_구분    lib.T시장구분
	M관리_질의_구분 T관리_질의_구분
	M연속키      string
}

type S관리종목_조회_응답_t1404 struct {
	M헤더     *S관리종목_조회_응답_헤더_t1404
	M반복값_모음 *S관리종목_조회_응답_반복값_모음_t1404
}

func (s *S관리종목_조회_응답_t1404) G헤더_TR데이터() I헤더_TR데이터 {
	return s.M헤더
}

func (s *S관리종목_조회_응답_t1404) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

type S관리종목_조회_응답_헤더_t1404 struct {
	M연속키 string
}

func (s *S관리종목_조회_응답_헤더_t1404) G헤더_TR데이터() I헤더_TR데이터 {
	return s
}

type S관리종목_조회_응답_반복값_모음_t1404 struct {
	M배열 []*S관리종목_조회_응답_반복값_t1404
}

func (s *S관리종목_조회_응답_반복값_모음_t1404) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

type S관리종목_조회_응답_반복값_t1404 struct {
	M종목코드       string
	M종목명        string
	M현재가        int64
	M전일대비구분     T전일대비_구분
	M전일대비_등락폭   int64
	M전일대비_등락율   float64
	M거래량        int64
	M지정일_주가     int64
	M지정일_대비_등락폭 int64
	M지정일_대비_등락율 float64
	M사유         T관리종목_지정_사유_구분
	M지정일        time.Time
	M해제일        time.Time
}

func NewT1404InBlock(질의값 *S질의값_관리종목_조회_t1404) (g *T1404InBlock) {
	g = new(T1404InBlock)
	lib.F바이트_복사_문자열(g.Gubun[:], strconv.Itoa(int(질의값.M시장_구분)))
	lib.F바이트_복사_문자열(g.Jongchk[:], strconv.Itoa(int(질의값.M관리_질의_구분)))
	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M연속키)

	return g
}

func New관리종목_조회_응답_헤더_t1404(b []byte) (값 *S관리종목_조회_응답_헤더_t1404, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT1404OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T1404OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S관리종목_조회_응답_헤더_t1404)
	값.M연속키 = lib.F2문자열(g.Shcode)

	return 값, nil
}

func New관리종목_조회_응답_반복값_모음_t1404(b []byte) (값_모음 *S관리종목_조회_응답_반복값_모음_t1404, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	나머지 := len(b) % SizeT1404OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT1404OutBlock1
	g_모음 := make([]*T1404OutBlock1, 수량, 수량)

	값_모음 = new(S관리종목_조회_응답_반복값_모음_t1404)
	값_모음.M배열 = make([]*S관리종목_조회_응답_반복값_t1404, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T1404OutBlock1)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		값 := new(S관리종목_조회_응답_반복값_t1404)
		값.M종목코드 = lib.F2문자열(g.Shcode)
		값.M종목명 = lib.F2문자열_EUC_KR_공백제거(g.Hname)
		값.M현재가 = lib.F2정수64_단순형(g.Price)
		값.M전일대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Sign))
		값.M전일대비_등락폭 = 값.M전일대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.Change))
		값.M전일대비_등락율 = 값.M전일대비구분.G부호보정_실수64(lib.F2실수_소숫점_추가_단순형(g.Diff, 2))
		값.M거래량 = lib.F2정수64_단순형(g.Volume)
		값.M지정일 = lib.F2포맷된_일자_단순형("20060102", g.Date)
		값.M지정일_주가 = lib.F2정수64_단순형(g.Tprice)
		값.M지정일_대비_등락폭 = lib.F2정수64_단순형(g.Tchange)
		값.M지정일_대비_등락율 = lib.F2실수_소숫점_추가_단순형(g.Tdiff, 2)
		값.M사유 = T관리종목_지정_사유_구분(lib.F2정수64_단순형_공백은_0(g.Reason))
		값.M해제일 = lib.F2포맷된_일자_단순형_공백은_초기값("20060102", g.Edate)

		값_모음.M배열[i] = 값
	}

	return 값_모음, nil
}
