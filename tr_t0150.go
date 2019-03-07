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

func F현물_당일_매매일지_수수료_t0150(계좌번호 string) (응답값 *S현물_당일_매매일지_수수료_응답_t0150, 에러 error) {
	panic("TODO")
}

type S현물_당일_매매일지_수수료_응답_t0150 struct {
	M헤더     *S현물_당일_매매일지_수수료_응답_헤더_t0150
	M반복값_모음 []*S현물_당일_매매일지_수수료_응답_반복값_t0150
}

type S현물_당일_매매일지_수수료_응답_헤더_t0150 struct {
	M매도_수량   int64
	M매도_약정금액 int64
	M매도_수수료  int64
	M매도_거래세  int64
	M매도_농특세  int64
	M매도_제비용합 int64
	M매도_정산금액 int64
	M매수_수량   int64
	M매수_약정금액 int64
	M매수_수수료  int64
	M매수_제비용합 int64
	M매수_정산금액 int64
	M합계_수량   int64
	M합계_약정금액 int64
	M합계_수수료  int64
	M합계_거래세  int64
	M합계_농특세  int64
	M합계_제비용합 int64
	M합계_정산금액 int64
	CTS_매매구분 string
	CTS_종목번호 string
	CTS_단가   string
	CTS_매체   string
}

type S현물_당일_매매일지_수수료_응답_반복값_t0150 struct {
	M매매구분 string
	M종목코드 string
	M수량   int64
	M단가   int64
	M약정금액 int64
	M수수료  int64
	M거래세  int64
	M농특세  int64
	M정산금액 int64
	M매체   string
}