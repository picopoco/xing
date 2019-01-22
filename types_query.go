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

// t1305 기간별 주가
type S질의값_현물_기간별_조회 struct {
	*lib.S질의값_단일_종목
	M일주월_구분 T일주월_구분
	M수량     int
	M연속키    string
}

func New질의값_현물_기간별_조회() *S질의값_현물_기간별_조회 {
	s := new(S질의값_현물_기간별_조회)
	s.S질의값_단일_종목 = lib.New질의값_단일_종목()

	return s
}

// t1310 전일당일분틱조회
type S질의값_현물_전일당일_분틱_조회 struct {
	*lib.S질의값_단일_종목
	M당일전일구분 T당일전일_구분  // 0:당일, 1:전일
	M분틱구분   T분틱_구분    // 0:분, 1:틱
	M종료시각   time.Time // 해당 시각 이전까지의 데이터만 조회됨.
	M연속키    string    // 처음 조회시 Space. 다음 조회시 t1310OutBlock.cts_time 값 입력
}

func New질의값_현물_전일당일_분틱_조회() *S질의값_현물_전일당일_분틱_조회 {
	s := new(S질의값_현물_전일당일_분틱_조회)
	s.S질의값_단일_종목 = lib.New질의값_단일_종목()

	return s
}

type S질의값_단일종목_연속키 struct {
	*lib.S질의값_단일_종목
	M연속키 string
}

func New질의값_단일종목_연속키() *S질의값_단일종목_연속키 {
	s := new(S질의값_단일종목_연속키)
	s.S질의값_단일_종목 = lib.New질의값_단일_종목()

	return s
}

// t8411 현물 차트 틱
type S질의값_현물_차트_틱 struct {
	*lib.S질의값_단일_종목
	M단위     int // n틱
	M요청건수   int // 최대 압축 2000, 비압축 500
	M조회영업일수 int // 0 : 미사용, 1 >= 사용
	M시작일자   time.Time
	M종료일자   time.Time
	M연속일자   string
	M연속시간   string
	M압축여부   bool
}

func New질의값_현물_차트_틱() *S질의값_현물_차트_틱 {
	s := new(S질의값_현물_차트_틱)
	s.S질의값_단일_종목 = lib.New질의값_단일_종목()

	return s
}

// t8412 현물 차트 분
type S질의값_현물_차트_분 struct {
	*lib.S질의값_단일_종목
	M단위     int // n분
	M요청건수   int // 최대 압축 2000, 비압축 500
	M조회영업일수 int // 0 : 미사용, 1 >= 사용
	M시작일자   time.Time
	M종료일자   time.Time
	M연속일자   string
	M연속시간   string
	M압축여부   bool
}

func New질의값_현물_차트_분() *S질의값_현물_차트_분 {
	s := new(S질의값_현물_차트_분)
	s.S질의값_단일_종목 = lib.New질의값_단일_종목()

	return s
}

// t8428 증시주변자금추이
type S질의값_증시주변자금추이 struct {
	*lib.S질의값_기본형
	//M시작_일자 time.Time  // 게시판 답변 : 해당 필드(시작,종료)의 일자는 사용하지 않습니다.
	//M종료_일자 time.Time  // 게시판 답변 : 해당 필드(시작,종료)의 일자는 사용하지 않습니다.
	//M구분 uint8.// 게시판 답변 : 해당 구분값은 의미가 없습니다.
	M시장구분 lib.T시장구분
	M수량   int
	M연속키  string
}

func New질의값_증시주변자금추이() *S질의값_증시주변자금추이 {
	s := new(S질의값_증시주변자금추이)
	s.S질의값_기본형 = new(lib.S질의값_기본형)

	return s
}
