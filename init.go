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

Copyright (C) 2015-2018년 UnHa Kim (unha.kim@kuh.pe.kr)

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
)

func init() {
	f초기화_xing_C32()
	f초기화_소켓()
	f초기화_영업일_기준_전일_당일()
}

func f초기화_xing_C32() (에러 error) {
	xing_C32_실행_잠금.Lock()
	defer xing_C32_실행_잠금.Unlock()

	defer lib.S에러패닉_처리기{M에러_포인터: &에러}.S실행()

	if !lib.F인터넷에_접속됨() {
		lib.F문자열_출력("인터넷을 확인하십시오.")
		return
	} else if xing_C32_실행_중() {
		lib.F문자열_출력("xing_C32 가 이미 실행 중입니다.")
		return nil
	}

	pid, 에러 := lib.F외부_프로세스_실행(go실행파일_경로, "run", xing_C32_경로)
	lib.F조건부_패닉(pid <= 0, "예상하지 못한 PID값. %v", pid)
	lib.F에러체크(에러)

	f접속유지_실행()

	return nil
}

func f초기화_소켓() (에러 error) {
	defer lib.S에러패닉_처리기{M에러_포인터: &에러}.S실행()

	소켓SUB_콜백 = lib.New소켓SUB_단순형(lib.P주소_Xing_C함수_콜백)
	소켓SUB_실시간_정보 = lib.New소켓SUB_단순형(lib.P주소_Xing_실시간)

	return nil
}

func f초기화_영업일_기준_전일_당일() (에러 error) {
	panic("TODO")
}
