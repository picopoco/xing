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

type I헤더_반복값_TR데이터 interface {
	I헤더_TR데이터
	I반복값_모음_TR데이터
}

type I헤더_TR데이터 interface {
	G헤더_TR데이터() I헤더_TR데이터
}

type I반복값_모음_TR데이터 interface {
	G반복값_모음_TR데이터() I반복값_모음_TR데이터
}

type S헤더_반복값_일반형 struct {
	M헤더     I헤더_TR데이터
	M반복값_모음 I반복값_모음_TR데이터
}

func (s *S헤더_반복값_일반형) G헤더_TR데이터() I헤더_TR데이터 { return s.M헤더 }
func (s *S헤더_반복값_일반형) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

func (s *S헤더_반복값_일반형) G값(TR코드 string) interface{} {
	switch TR코드 {
	default:
		panic(lib.New에러with출력("예상하지 못한 TR코드 : '%v'", TR코드)) // 패닉 출력을 삭제하지 말 것.
	//case TR현물_시간대별_체결_조회:
	//	값 := new(S현물_시간대별_체결_응답)
	//	값.M헤더 = s.M헤더.(*S현물_시간대별_체결_응답_헤더)
	//	값.M반복값_모음 = s.M반복값_모음.(*S현물_시간대별_체결_응답_반복값_모음)
	//	return 값
	case TR현물_기간별_조회:
		값 := new(S현물_기간별_조회_응답)
		값.M헤더 = s.M헤더.(*S현물_기간별_조회_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*S현물_기간별_조회_응답_반복값_모음)
		return 값
	case TR현물_당일_전일_분틱_조회:
		값 := new(S현물_전일당일분틱조회_응답)
		값.M헤더 = s.M헤더.(*S현물_전일당일분틱조회_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*S현물_전일당일분틱조회_응답_반복값_모음)
		return 값
	case TR_ETF_시간별_추이:
		값 := new(S_ETF시간별_추이_응답)
		값.M헤더 = s.M헤더.(*S_ETF시간별_추이_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*S_ETF시간별_추이_응답_반복값_모음)
		return 값
	case TR현물_차트_틱:
		값 := new(S현물_차트_틱_응답)
		값.M헤더 = s.M헤더.(*S현물_차트_틱_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*S현물_차트_틱_응답_반복값_모음)
		return 값
	case TR현물_차트_분:
		값 := new(S현물_차트_분_응답)
		값.M헤더 = s.M헤더.(*S현물_차트_분_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*S현물_차트_분_응답_반복값_모음)
		return 값
	case TR증시_주변_자금_추이:
		값 := new(S증시_주변자금추이_응답)
		값.M헤더 = s.M헤더.(*S증시_주변자금추이_응답_헤더)
		값.M반복값_모음 = s.M반복값_모음.(*S증시_주변자금추이_응답_반복값_모음)
		return 값
	}
}

// t1101 현물 호가 조회 응답
type S현물_호가조회_응답 struct {
	M한글명         string
	M현재가         int64
	M전일대비구분      T전일대비_구분
	M전일대비등락폭     int64
	M등락율         float64
	M누적거래량       int64
	M전일종가        int64
	M매도호가_모음     []int64
	M매수호가_모음     []int64
	M매도호가수량_모음   []int64
	M매수호가수량_모음   []int64
	M직전매도대비수량_모음 []int64
	M직전매수대비수량_모음 []int64
	M매도호가수량합     int64
	M매수호가수량합     int64
	M직전매도대비수량합   int64
	M직전매수대비수량합   int64
	M수신시간        time.Time
	M예상체결가격      int64
	M예상체결수량      int64
	M예상체결전일구분    T전일대비_구분
	M예상체결전일대비    int64
	M예상체결등락율     float64
	M시간외매도잔량     int64
	M시간외매수잔량     int64
	M동시호가_구분     T동시호가_구분
	M종목코드        string
	M상한가         int64
	M하한가         int64
	M시가          int64
	M고가          int64
	M저가          int64
}

// t1102 현물 시세(현재가) 조회 응답
type S현물_시세조회_응답 struct {
	M한글명         string
	M현재가         int64
	M전일대비구분      T전일대비_구분
	M전일대비등락폭     int64
	M등락율         float64
	M누적거래량       int64
	M기준가         int64
	M가중평균        int64
	M상한가         int64
	M하한가         int64
	M전일거래량       int64
	M거래량차        int64
	M시가          int64
	M시가시간        time.Time
	M고가          int64
	M고가시간        time.Time
	M저가          int64
	M저가시간        time.Time
	M52주_최고가     int64
	M52주_최고가일    time.Time
	M52주_최저가     int64
	M52주_최저가일    time.Time
	M소진율         float64
	PER          float64
	PBRX         float64
	M상장주식수_천     int64
	M증거금율        int64
	M수량단위        int64
	M매도증권사코드_모음  []string
	M매수증권사코드_모음  []string
	M매도증권사명_모음   []string
	M매수증권사명_모음   []string
	M총매도수량_모음    []int64
	M총매수수량_모음    []int64
	M매도증감_모음     []int64
	M매수증감_모음     []int64
	M매도비율_모음     []float64
	M매수비율_모음     []float64
	M외국계_매도_합계수량 int64
	M외국계_매도_직전대비 int64
	M외국계_매도_비율   float64
	M외국계_매수_합계수량 int64
	M외국계_매수_직전대비 int64
	M외국계_매수_비율   float64
	M회전율         float64
	M종목코드        string
	M누적거래대금      int64
	M전일동시간거래량    int64
	M연중_최고가      int64
	M연중_최고가_일자   time.Time
	M연중_최저가      int64
	M연중_최저가_일자   time.Time
	M목표가         int64
	M자본금         int64
	M유동주식수       int64
	M액면가         int64
	M결산월         uint8
	M대용가         int64
	M시가총액_억      int64
	M상장일         time.Time
	M전분기명        string
	M전분기_매출액     int64
	M전분기_영업이익    int64
	M전분기_경상이익    int64
	M전분기_순이익     int64
	M전분기EPS      float64
	M전전분기명       string
	M전전분기_매출액    int64
	M전전분기_영업이익   int64
	M전전분기_경상이익   int64
	M전전분기_순이익    int64
	M전전분기EPS     float64
	M전년대비매출액     float64
	M전년대비영업이익    float64
	M전년대비경상이익    float64
	M전년대비순이익     float64
	M전년대비EPS     float64
	M락구분         string
	M관리_급등구분     string
	M정지_연장구분     string
	M투자_불성실구분    string
	M시장구분        lib.T시장구분
	T_PER        float64
	M통화ISO코드     string
	M총매도대금_모음    []int64
	M총매수대금_모음    []int64
	M총매도평단가_모음   []int64
	M총매수평단가_모음   []int64
	M외국계매도대금     int64
	M외국계매수대금     int64
	M외국계매도평단가    int64
	M외국계매수평단가    int64
	M투자주의환기      string
	M기업인수목적회사여부  bool
	M발행가격        int64
	M배분적용구분코드    string // 배분적용구분코드(1:배분발생 2:배)
	M배분적용구분      string
	M단기과열_VI발동   string
	M정적VI상한가     int64
	M정적VI하한가     int64
	M저유동성종목여부    bool
	M이상급등종목여부    bool
	M대차불가여부      bool
}

// t1301 현물 시간대별 체결 응답
type S현물_시간대별_체결_응답 struct {
	M헤더     *S현물_시간대별_체결_응답_헤더
	M반복값_모음 *S현물_시간대별_체결_응답_반복값_모음
}

func (s *S현물_시간대별_체결_응답) G헤더_TR데이터() I헤더_TR데이터 {
	return s.M헤더
}
func (s *S현물_시간대별_체결_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

// t1301 현물 시간대별 체결 응답 헤더
type S현물_시간대별_체결_응답_헤더 struct {
	M연속키 string
}

func (s *S현물_시간대별_체결_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 {
	return s
}

// t1301 현물 시간대별 체결 응답 반복값
type S현물_시간대별_체결_응답_반복값 struct {
	M시각      time.Time
	M현재가     int64
	M전일대비구분  T전일대비_구분
	M전일대비등락폭 int64
	M전일대비등락율 float64
	M체결수량    int64
	M체결강도    float64
	M거래량     int64
	M매도체결수량  int64
	M매도체결건수  int64
	M매수체결수량  int64
	M매수체결건수  int64
	M순체결량    int64
	M순체결건수   int64
}

type S현물_시간대별_체결_응답_반복값_모음 struct {
	M배열 []*S현물_시간대별_체결_응답_반복값
}

func (s *S현물_시간대별_체결_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
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

// t1310 전일당일분틱조회 응답
type S현물_전일당일분틱조회_응답 struct {
	M헤더     *S현물_전일당일분틱조회_응답_헤더
	M반복값_모음 *S현물_전일당일분틱조회_응답_반복값_모음
}

func (s *S현물_전일당일분틱조회_응답) G헤더_TR데이터() I헤더_TR데이터 {
	return s.M헤더
}
func (s *S현물_전일당일분틱조회_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

// t1310 전일당일분틱조회 응답 헤더
type S현물_전일당일분틱조회_응답_헤더 struct {
	M연속키 string
}

func (s *S현물_전일당일분틱조회_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 {
	return s
}

// t1310 전일당일분틱조회 응답 반복값
type S현물_전일당일분틱조회_응답_반복값 struct {
	M시각      time.Time
	M현재가     int64
	M전일대비구분  T전일대비_구분
	M전일대비등락폭 int64
	M전일대비등락율 float64
	M체결수량    int64
	M체결강도    float64
	M거래량     int64
	M매도체결수량  int64
	M매도체결건수  int64
	M매수체결수량  int64
	M매수체결건수  int64
	M순체결량    int64
	M순체결건수   int64
}

type S현물_전일당일분틱조회_응답_반복값_모음 struct {
	M배열 []*S현물_전일당일분틱조회_응답_반복값
}

func (s *S현물_전일당일분틱조회_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

// t1901 ETF 현재가 조회 응답
type S_ETF_현재가_조회_응답 struct {
	M종목코드           string
	M명칭             string
	M현재가            int64
	M전일대비구분         T전일대비_구분
	M전일대비등락폭        int64
	M전일대비등락율        float64
	M누적_거래량         int64
	M기준가            int64
	M가중평균           int64
	M상한가            int64
	M하한가            int64
	M전일_거래량         float64
	M거래량차           int64
	M시가             int64
	M시가시각           time.Time
	M고가             int64
	M고가시각           time.Time
	M저가             int64
	M저가시각           time.Time
	M52주_최고가        int64
	M52주_최고가일       time.Time
	M52주_최저가        int64
	M52주_최저가일       time.Time
	M소진율            float64
	M외국인_보유수량       float64
	PER             float64
	M상장주식수_천        int64
	M증거금율           int64
	M증거율            float64
	M누적_거래대금        int64
	M연중_최고가         int64
	M연중_최고일자        time.Time
	M연중_최저가         int64
	M연중_최저일자        time.Time
	M업종명            string
	M업종코드           string
	M업종_현재가         float64
	M업종_전일대비구분      T전일대비_구분
	M업종_전일대비등락폭     int64
	M업종_전일대비등락율     float64
	M선물_최근_월물명      string
	M선물_최근_월물코드     string
	M선물_현재가         float64
	M선물_전일대비구분      T전일대비_구분
	M선물_전일대비등락폭     int64
	M선물_전일대비등락율     float64
	NAV             float64
	NAV_전일대비구분      T전일대비_구분
	NAV_전일대비등락폭     float64
	NAV_전일대비등락율     float64
	M추적_오차율         float64
	M괴리율            float64
	M대용가            int64
	M매도_증권사_코드      []string
	M매수_증권사_코드      []string
	M총매도수량          []int64
	M총매수수량          []int64
	M매도증감           []int64
	M매수증감           []int64
	M매도비율           []float64
	M매수비율           []float64
	M외국계_매도_합계_수량   int64
	M외국계_매도_직전_대비   T전일대비_구분
	M외국계_매도_비율      float64
	M외국계_매수_합계_수량   int64
	M외국계_매수_직전_대비   T전일대비_구분
	M외국계_매수_비율      float64
	M참고지수명          string
	M참고지수코드         string
	M참고지수현재가        float64
	M전일NAV          float64
	M전일NAV_전일대비구분   T전일대비_구분
	M전일NAV_전일대비등락폭  float64
	M전일NAV_전일대비등락율  float64
	M순자산총액_억        int64
	M스프레드           float64
	M레버리지           int64
	M과세구분           uint8
	M운용사            string
	M유동성공급자         []string
	M복제방법           string
	M상품유형           string
	VI발동해제          string
	ETN상품분류         string
	ETN만기일          time.Time
	ETN지급일          time.Time
	ETN최종거래일        time.Time
	ETN발행시장참가자      string
	ETN만기상환가격결정_시작일 time.Time
	ETN만기상환가격결정_종료일 time.Time
	ETN유동성공급자_보유수량  int64
}

// t1902 ETF 시간별 추이 응답
type S_ETF시간별_추이_응답 struct {
	M헤더     *S_ETF시간별_추이_응답_헤더
	M반복값_모음 *S_ETF시간별_추이_응답_반복값_모음
}

func (s *S_ETF시간별_추이_응답) G헤더_TR데이터() I헤더_TR데이터 { return s.M헤더 }
func (s *S_ETF시간별_추이_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

// t1902 ETF 시간별 추이 응답 헤더
type S_ETF시간별_추이_응답_헤더 struct {
	M연속키   string
	M종목명   string
	M업종지수명 string
}

func (s *S_ETF시간별_추이_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 { return s }

// t1902 ETF 시간별 추이 응답 반복값
type S_ETF시간별_추이_응답_반복값 struct {
	M시각         time.Time
	M현재가        int64
	M전일대비구분     T전일대비_구분
	M전일대비등락폭    int64
	M누적_거래량     float64
	M현재가_NAV_차이 float64
	NAV         float64
	NAV전일대비등락폭  float64
	M추적오차       float64
	M괴리율        float64
	M지수         float64
	M지수_전일대비등락폭 float64
	M지수_전일대비등락율 float64
}

type S_ETF시간별_추이_응답_반복값_모음 struct {
	M배열 []*S_ETF시간별_추이_응답_반복값
}

func (s *S_ETF시간별_추이_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

// t8411 현물 차트 틱 응답
type S현물_차트_틱_응답 struct {
	M헤더     *S현물_차트_틱_응답_헤더
	M반복값_모음 *S현물_차트_틱_응답_반복값_모음
}

func (s *S현물_차트_틱_응답) G헤더_TR데이터() I헤더_TR데이터 { return s.M헤더 }
func (s *S현물_차트_틱_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

// t8411 현물 차트 틱 응답 헤더
type S현물_차트_틱_응답_헤더 struct {
	M종목코드     string
	M전일시가     int64
	M전일고가     int64
	M전일저가     int64
	M전일종가     int64
	M전일거래량    int64
	M당일시가     int64
	M당일고가     int64
	M당일저가     int64
	M당일종가     int64
	M상한가      int64
	M하한가      int64
	M연속일자     string
	M연속시간     string
	M장시작시간    time.Time
	M장종료시간    time.Time
	M동시호가처리시간 int
	M수량       int64
}

func (s *S현물_차트_틱_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 { return s }

// t8411 현물 차트 틱 응답 반복값
type S현물_차트_틱_응답_반복값 struct {
	M일자_시각    time.Time
	M시가       int64
	M고가       int64
	M저가       int64
	M종가       int64
	M거래량      int64
	M수정구분     T수정구분
	M수정비율     float64
	M수정주가반영항목 int64
}

type S현물_차트_틱_응답_반복값_모음 struct {
	M배열 []*S현물_차트_틱_응답_반복값
}

func (s *S현물_차트_틱_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

// t8412 현물 차트 분 응답
type S현물_차트_분_응답 struct {
	M헤더     *S현물_차트_분_응답_헤더
	M반복값_모음 *S현물_차트_분_응답_반복값_모음
}

func (s *S현물_차트_분_응답) G헤더_TR데이터() I헤더_TR데이터 { return s.M헤더 }
func (s *S현물_차트_분_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

// t8412 현물 차트 틱 응답 헤더
type S현물_차트_분_응답_헤더 struct {
	M종목코드     string
	M전일시가     int64
	M전일고가     int64
	M전일저가     int64
	M전일종가     int64
	M전일거래량    int64
	M당일시가     int64
	M당일고가     int64
	M당일저가     int64
	M당일종가     int64
	M상한가      int64
	M하한가      int64
	M연속일자     string
	M연속시간     string
	M장시작시간    time.Time
	M장종료시간    time.Time
	M동시호가처리시간 int
	M수량       int64
}

func (s *S현물_차트_분_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 { return s }

// t8412 현물 차트 틱 응답 반복값
type S현물_차트_분_응답_반복값 struct {
	M일자_시각  time.Time
	M시가     int64
	M고가     int64
	M저가     int64
	M종가     int64
	M거래량    int64
	M거래대금   int64
	M수정구분   T수정구분
	M수정비율   float64
	M종가등락구분 T전일대비_구분
}

type S현물_차트_분_응답_반복값_모음 struct {
	M배열 []*S현물_차트_분_응답_반복값
}

func (s *S현물_차트_분_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

// t8428 증시주변자금추이 응답
type S증시_주변자금추이_응답 struct {
	M헤더     *S증시_주변자금추이_응답_헤더
	M반복값_모음 *S증시_주변자금추이_응답_반복값_모음
}

func (s *S증시_주변자금추이_응답) G헤더_TR데이터() I헤더_TR데이터 {
	return s.M헤더
}
func (s *S증시_주변자금추이_응답) G반복값_TR데이터() I반복값_모음_TR데이터 {
	return s.M반복값_모음
}

// t8428 증시주변자금추이 응답 헤더
type S증시_주변자금추이_응답_헤더 struct {
	M연속키 string
	M인덱스 int64
}

func (s *S증시_주변자금추이_응답_헤더) G헤더_TR데이터() I헤더_TR데이터 { return s }

// t8428 증시주변자금추이 응답 반복값
type S증시_주변자금추이_응답_반복값 struct {
	M일자       time.Time
	M지수       float64
	M전일대비_구분  T전일대비_구분
	M전일대비_등락폭 float64
	M전일대비_등락율 float64
	M거래량      int64
	M고객예탁금_억  int64
	M예탁증감_억   int64
	M회전율      float64
	M미수금_억    int64
	M신용잔고_억   int64
	M선물예수금_억  int64
	M주식형_억    int64
	M혼합형_주식_억 int64
	M혼합형_채권_억 int64
	M채권형_억    int64
	MMF_억     int64
}

type S증시_주변자금추이_응답_반복값_모음 struct {
	M배열 []*S증시_주변자금추이_응답_반복값
}

func (s *S증시_주변자금추이_응답_반복값_모음) G반복값_모음_TR데이터() I반복값_모음_TR데이터 {
	return s
}

// t8436 현물 종목조회 응답
type S현물_종목조회_응답_반복값 struct {
	M종목명        string
	M종목코드       string
	M시장구분       lib.T시장구분
	M주문수량단위     int
	M상한가        int64
	M하한가        int64
	M전일가        int64
	M기준가        int64
	M증권그룹       T증권그룹
	M기업인수목적회사여부 bool
}

type S현물_종목조회_응답_반복값_모음 struct {
	M배열 []*S현물_종목조회_응답_반복값
}
