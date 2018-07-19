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

import "github.com/ghts/lib"

func go_RT_주문처리결과(ch초기화 chan lib.T신호) {
	ch종료 := lib.F공통_종료_채널()
	ch초기화 <- lib.P신호_초기화

	var 수신값 *lib.S바이트_변환_모음
	var 주문_처리_결과 *S현물_주문_응답_실시간_정보
	var 에러 error

	for {
		select {
		case <-ch종료:
			return
		default:
			수신값, 에러 = 소켓SUB_주문처리.G수신()
			if 에러 != nil {
				select {
				case <-ch종료:
					에러 = nil
					return
				default:
					lib.F에러_출력(lib.New에러(에러))
					continue
				}
			}

			lib.F조건부_패닉(수신값.G수량() != 1, "메시지 길이 : 예상값 1, 실제값 %v.", 수신값.G수량())

			if 수신값.G자료형_문자열(0) != P자료형_S현물_주문_응답_실시간_정보 {
				continue
			}

			주문_처리_결과 = new(S현물_주문_응답_실시간_정보)
			if 에러 = 수신값.G값(0, 주문_처리_결과); 에러 != nil {
				lib.F에러_출력(에러)
				continue
			}

			//체크(주문_처리_결과.RT코드, 주문_처리_결과.M응답_구분, 주문_처리_결과.M주문번호, 주문_처리_결과.M원_주문번호, 주문_처리_결과.M수량, 주문_처리_결과.M잔량)
		}
	}
}
