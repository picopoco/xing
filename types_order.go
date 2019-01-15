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

type S현물_주문_응답_실시간_정보 struct { // 'SCn'
	M주문번호   int64
	M원_주문번호 int64
	RT코드    string
	M응답_구분  T주문_응답_구분
	M종목코드   string
	M수량     int64
	M가격     int64
	M잔량     int64
	M시각     time.Time
}

type S질의값_정상_주문 struct {
	*lib.S질의값_정상_주문
	M계좌_비밀번호 string
	M신용거래_구분 lib.T신용거래_구분
	M대출일     time.Time
}

func New질의값_정상_주문() *S질의값_정상_주문 {
	s := new(S질의값_정상_주문)
	s.S질의값_정상_주문 = lib.New질의값_정상_주문()

	return s
}

type S질의값_정정_주문 struct {
	*lib.S질의값_정정_주문
	M계좌_비밀번호 string
	M주문조건    lib.T주문조건
	M호가유형    lib.T호가유형
}

func New질의값_정정_주문() *S질의값_정정_주문 {
	s := new(S질의값_정정_주문)
	s.S질의값_정정_주문 = lib.New질의값_정정_주문()

	return s
}

type S질의값_취소_주문 struct {
	*lib.S질의값_취소_주문
	M계좌_비밀번호 string
}

func New질의값_취소_주문() *S질의값_취소_주문 {
	s := new(S질의값_취소_주문)
	s.S질의값_취소_주문 = lib.New질의값_취소_주문()

	return s
}

type I주문_응답 interface {
	I주문_응답1
	I주문_응답2
}

type I주문_응답1 interface {
	G주문_응답1() I주문_응답1
}

type I주문_응답2 interface {
	G주문_응답2() I주문_응답2
}

type S주문_응답_일반형 struct {
	M응답1 I주문_응답1
	M응답2 I주문_응답2
}

func (s *S주문_응답_일반형) G주문_응답1() I주문_응답1 { return s.M응답1 }
func (s *S주문_응답_일반형) G주문_응답2() I주문_응답2 { return s.M응답2 }

func (s *S주문_응답_일반형) G값(TR코드 string) interface{} {
	switch TR코드 {
	case TR현물_정상_주문:
		g := new(S현물_정상_주문_응답)
		g.M응답1 = s.M응답1.(*S현물_정상_주문_응답1)
		g.M응답2 = s.M응답2.(*S현물_정상_주문_응답2)
		return g
	case TR현물_정정_주문:
		g := new(S현물_정정_주문_응답)
		g.M응답1 = s.M응답1.(*S현물_정정_주문_응답1)
		g.M응답2 = s.M응답2.(*S현물_정정_주문_응답2)
		return g
	case TR현물_취소_주문:
		g := new(S현물_취소_주문_응답)
		g.M응답1 = s.M응답1.(*S현물_취소_주문_응답1)
		g.M응답2 = s.M응답2.(*S현물_취소_주문_응답2)
		return g
	default:
		panic(lib.New에러("예상하지 못한 TR코드 : '%v'", TR코드))
	}
}

type S현물_정상_주문_응답 struct {
	M응답1 *S현물_정상_주문_응답1
	M응답2 *S현물_정상_주문_응답2
}

func (s *S현물_정상_주문_응답) G주문_응답1() I주문_응답1 { return s.M응답1 }
func (s *S현물_정상_주문_응답) G주문_응답2() I주문_응답2 { return s.M응답2 }

type S현물_정상_주문_응답1 struct {
	M레코드_수량     int
	M계좌번호       string
	M계좌_비밀번호    string
	M종목코드       string
	M주문수량       int64
	M주문가격       int64
	M매매구분       lib.T매수_매도
	M호가유형       lib.T호가유형
	M프로그램_호가유형  string
	M공매도_가능     bool
	M공매도_호가구분   string
	M통신매체_코드    string
	M신용거래_구분    lib.T신용거래_구분
	M대출일        time.Time
	M회원번호       string
	M주문조건_구분    lib.T주문조건
	M전략코드       string
	M그룹ID       string
	M주문회차       int64
	M포트폴리오_번호   int64
	M트렌치_번호     int64
	M아이템_번호     int64
	M운용지시_번호    string
	M유동성_공급자_여부 bool
	M반대매매_구분    string
}

func (s *S현물_정상_주문_응답1) G주문_응답1() I주문_응답1 { return s }

type S현물_정상_주문_응답2 struct {
	M레코드_수량    int
	M주문번호      int64
	M주문시각      time.Time
	M주문시장_코드   T주문_시장구분
	M주문유형_코드   string
	M종목코드      string // 단축종목번호
	M관리사원_번호   string
	M주문금액      int64
	M예비_주문번호   int64
	M반대매매_일련번호 int64
	M예약_주문번호   int64
	M재사용_주문수량  int64
	M현금_주문금액   int64
	M대용_주문금액   int64
	M재사용_주문금액  int64
	M계좌명       string
	M종목명       string
}

func (s *S현물_정상_주문_응답2) G주문_응답2() I주문_응답2 { return s }

type S현물_정정_주문_응답 struct {
	M응답1 *S현물_정정_주문_응답1
	M응답2 *S현물_정정_주문_응답2
}

func (s *S현물_정정_주문_응답) G주문_응답1() I주문_응답1 { return s.M응답1 }
func (s *S현물_정정_주문_응답) G주문_응답2() I주문_응답2 { return s.M응답2 }

type S현물_정정_주문_응답1 struct {
	M레코드_수량   int
	M원_주문번호   int64
	M계좌번호     string
	M계좌_비밀번호  string
	M종목코드     string
	M주문수량     int64
	M호가유형     lib.T호가유형
	M주문조건     lib.T주문조건
	M주문가격     int64
	M통신매체_코드  string
	M전략코드     string
	M그룹ID     string
	M주문회차     int64
	M포트폴리오_번호 int64
	M바스켓_번호   int64
	M트렌치_번호   int64
	M아이템_번호   int64
}

func (s *S현물_정정_주문_응답1) G주문_응답1() I주문_응답1 { return s }

type S현물_정정_주문_응답2 struct {
	M레코드_수량    int
	M주문번호      int64
	M모_주문번호    int64
	M주문시각      time.Time
	M주문시장_코드   T주문_시장구분
	M주문유형_코드   string
	M종목코드      string // 단축종목번호
	M공매도_호가구분  string
	M공매도_가능    bool
	M신용거래_구분   lib.T신용거래_구분
	M대출일       time.Time
	M반대매매주문_구분 string
	M유동성공급자여부  bool
	M관리사원_번호   string
	M주문금액      int64
	M매매구분      lib.T매수_매도
	M예비_주문번호   int64
	M반대매매_일련번호 int64
	M예약_주문번호   int64
	M현금_주문금액   int64
	M대용_주문금액   int64
	M재사용_주문금액  int64
	M계좌명       string
	M종목명       string
}

func (s *S현물_정정_주문_응답2) G주문_응답2() I주문_응답2 { return s }

type S현물_취소_주문_응답 struct {
	M응답1 *S현물_취소_주문_응답1
	M응답2 *S현물_취소_주문_응답2
}

func (s *S현물_취소_주문_응답) G주문_응답1() I주문_응답1 { return s.M응답1 }
func (s *S현물_취소_주문_응답) G주문_응답2() I주문_응답2 { return s.M응답2 }

type S현물_취소_주문_응답1 struct {
	M레코드_수량   int
	M원_주문번호   int64
	M계좌번호     string
	M계좌_비밀번호  string
	M종목코드     string
	M주문수량     int64
	M통신매체_코드  string
	M그룹ID     string
	M전략코드     string
	M주문회차     int64
	M포트폴리오_번호 int64
	M바스켓_번호   int64
	M트렌치_번호   int64
	M아이템_번호   int64
}

func (s *S현물_취소_주문_응답1) G주문_응답1() I주문_응답1 { return s }

type S현물_취소_주문_응답2 struct {
	M레코드_수량    int
	M주문번호      int64
	M모_주문번호    int64
	M주문시각      time.Time
	M주문시장_코드   T주문_시장구분
	M주문유형_코드   string
	M종목코드      string // 단축종목번호
	M공매도_호가구분  string
	M공매도_가능    bool
	M신용거래_코드   lib.T신용거래_구분
	M대출일       time.Time
	M반대매매주문_구분 string
	M유동성공급자_여부 bool
	M관리사원_번호   string
	M예비_주문번호   int64
	M반대매매_일련번호 int64
	M예약_주문번호   int64
	M계좌명       string
	M종목명       string
}

func (s *S현물_취소_주문_응답2) G주문_응답2() I주문_응답2 { return s }
