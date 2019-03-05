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

func New시각_조회_응답_t0167(시각 time.Time, 에러 error) *S시각_조회_응답_t0167 {
	s := new(S시각_조회_응답_t0167)
	s.M시각 = 시각
	s.M에러 = 에러

	return s
}

type S시각_조회_응답_t0167 struct {
	M시각 time.Time
	M에러 error
}

func (s S시각_조회_응답_t0167) G값() (time.Time, error) {
	return s.M시각, s.M에러
}

// 가장 간단한 질의. 접속 유지 및 질의 기능 테스트 용도로 적합함.
func F시각_조회_t0167() (ch응답 chan *S시각_조회_응답_t0167) {

	F접속_확인()

	ch응답 = make(chan *S시각_조회_응답_t0167, 1)

	ch질의 <- lib.New작업(f시각_조회_작업, ch응답)

	return ch응답
}

func f시각_조회_작업(인수 interface{}) {
	var 에러 error
	var ch응답 chan *S시각_조회_응답_t0167

	defer lib.S예외처리{M에러: &에러, M함수: func() {
		if ch응답 != nil {
			ch응답 <- New시각_조회_응답_t0167(time.Time{}, 에러)
		}
	}}.S실행()

	ch응답 = 인수.(chan *S시각_조회_응답_t0167)

	질의값 := lib.S질의값_기본형{M구분: TR조회, M코드: TR시간_조회_t0167}
	i응답값, 에러 := F질의_단일TR(질의값)
	lib.F확인(에러)

	값, ok := i응답값.(time.Time)
	lib.F조건부_패닉(!ok, "예상하지 못한 자료형 : '%T", i응답값)

	ch응답 <- New시각_조회_응답_t0167(값, nil)
}

func New시간_조회_응답(b []byte) (값 time.Time, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = time.Time{} }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT0167OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T0167OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	날짜_문자열 := lib.F2문자열(g.Date)
	시간_문자열 := lib.F2문자열(g.Time)

	return lib.F2포맷된_시각("20060102150405.99999999", 날짜_문자열+시간_문자열[:6]+"."+시간_문자열[7:])
}
