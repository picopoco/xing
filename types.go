/* Copyright (C) 2015-2018 김운하(UnHa Kim)  unha.kim@kuh.pe.kr

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

Copyright (C) 2015~2017년 UnHa Kim (unha.kim@kuh.pe.kr)

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

	"sync"
)

type 대기_항목_C32 struct {
	ch회신 chan interface{}
	TR코드 string
	대기값  interface{}
}

func new대기_TR_저장소_C32() *대기_TR_저장소_C32 {
	s := new(대기_TR_저장소_C32)
	s.저장소 = make(map[int]*대기_항목_C32)
}

// xing_C32  응답을 기다리는 TR 저장.
type 대기_TR_저장소_C32 struct {
	sync.RWMutex
	저장소 map[int]*대기_항목_C32
}

func (s *대기_TR_저장소_C32) G값(식별번호 int) *대기_항목_C32 {
	s.RLock()
	값 := s.저장소[식별번호]
	s.RUnlock()

	return 값
}

func (s *대기_TR_저장소_C32) S추가(식별번호 int, TR코드 string) chan interface{} {
	대기_항목 := new(대기_항목_C32)
	대기_항목.ch회신 = make(chan interface{}, 1)
	대기_항목.TR코드 = TR코드

	s.Lock()
	s.저장소[식별번호] = 대기_항목
	s.Unlock()

	return 대기_항목.ch회신
}

func (s *대기_TR_저장소_C32) S회신(식별번호 int) {
	대기_항목 := s.G값(식별번호)
	대기_항목.ch회신 <- 대기_항목.대기값

	s.Lock()
	delete(s.저장소, 식별번호)
	s.Unlock()
}

// (xing_C32가 아닌) 모듈로부터 소켓으로 들어오는 TR질의 저장
type s소켓_메시지_대기_저장소 struct {
	sync.Mutex
	저장소 map[lib.I소켓_메시지](chan lib.I소켓_메시지)
}

func (s *s소켓_메시지_대기_저장소) S추가(메시지 lib.I소켓_메시지, ch수신 chan lib.I소켓_메시지) {
	s.Lock()
	defer s.Unlock()
	s.저장소[메시지] = ch수신
}

func (s *s소켓_메시지_대기_저장소) S재전송() {
	s.Lock()
	defer s.Unlock()

	for 메시지, ch수신 := range s.저장소 {
		s.s재전송_도우미(메시지, ch수신)
	}
}

func (s *s소켓_메시지_대기_저장소) s재전송_도우미(메시지 lib.I소켓_메시지, ch수신 chan lib.I소켓_메시지) {
	// 채널이 이미 닫힌 경우 송신할 때 패닉이 발생함.
	// 그럴 경우에는 해당 메시지를 대기목록에서 삭제함.
	defer lib.S에러패닉_처리기{M함수: func() { delete(s.저장소, 메시지) }}.S실행()

	select {
	case ch수신 <- 메시지:
		// 중계 성공한 메시지는 대기열에서 삭제.
		delete(s.저장소, 메시지)
	default:
		// 중계 실패. 저장소에 그대로 두고 추후 재전송 시도.
	}
}

func new대기_중_데이터_저장소() *s소켓_메시지_대기_저장소 {
	s := new(s소켓_메시지_대기_저장소)
	s.저장소 = make(map[lib.I소켓_메시지](chan lib.I소켓_메시지))

	return s
}
