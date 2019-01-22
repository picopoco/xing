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

	"strconv"
)

// TR 및 응답 종류
const (
	TR조회 lib.TR구분 = iota
	TR주문
	TR실시간_정보_구독
	TR실시간_정보_해지
	TR실시간_정보_일괄_해지
	TR접속
	TR접속됨
	TR접속_해제
	TR초기화
	TR종료

	// Xing API에서 사용되는 것들
	TR서버_이름
	TR에러_코드
	TR에러_메시지
	TR코드별_쿼터
	TR계좌_수량
	TR계좌_번호
	TR계좌_이름
	TR계좌_상세명
	//TR압축_해제
	TR소켓_테스트
	TR전일_당일
)

func TR구분_String(v lib.TR구분) string {
	switch v {
	case TR조회:
		return "TR조회"
	case TR주문:
		return "TR주문"
	case TR실시간_정보_구독:
		return "TR실시간_정보_구독"
	case TR실시간_정보_해지:
		return "TR실시간_정보_해지"
	case TR실시간_정보_일괄_해지:
		return "TR실시간_정보_일괄_해지"
	case TR접속:
		return "TR접속"
	case TR접속됨:
		return "TR접속됨"
	case TR접속_해제:
		return "TR접속_해제"
	case TR초기화:
		return "TR초기화"
	case TR종료:
		return "TR종료"
	case TR서버_이름:
		return "서버_이름"
	case TR에러_코드:
		return "에러_코드"
	case TR에러_메시지:
		return "에러_메시지"
	case TR코드별_쿼터:
		return "TR코드별_쿼터"
	case TR계좌_수량:
		return "계좌_수량"
	case TR계좌_번호:
		return "계좌_번호"
	case TR계좌_이름:
		return "계좌_이름"
	case TR계좌_상세명:
		return "계좌_상세명"
	//case TR압축_해제:
	//	//	return "압축_해제"
	case TR소켓_테스트:
		return "신호"
	case TR전일_당일:
		return "당일_전일"
	default:
		return lib.F2문자열("예상하지 못한 M값 : '%v'", v)
	}
}

const (
	// 구현된 TR코드
	TR현물_정상_주문 = "CSPAT00600"
	TR현물_정정_주문 = "CSPAT00700"
	TR현물_취소_주문 = "CSPAT00800"
	TR계좌_거래_내역 = "CDPCQ04700"
	TR시간_조회    = "t0167"
	TR현물_호가_조회 = "t1101"
	TR현물_시세_조회 = "t1102"
	//TR현물_시간대별_체결_조회  = "t1301"
	TR현물_기간별_조회      = "t1305"
	TR현물_당일_전일_분틱_조회 = "t1310"
	TR_ETF_시세_조회     = "t1901"
	TR_ETF_시간별_추이    = "t1902"
	TR현물_차트_틱        = "t8411"
	TR현물_차트_분        = "t8412"
	TR증시_주변_자금_추이    = "t8428"
	TR현물_종목_조회       = "t8436"

	// 구현된 RT코드
	RT현물_주문_접수 = "SC0"
	RT현물_주문_체결 = "SC1"
	RT현물_주문_정정 = "SC2"
	RT현물_주문_취소 = "SC3"
	RT현물_주문_거부 = "SC4"

	RT코스피_호가_잔량     = "H1_"
	RT코스피_시간외_호가_잔량 = "H2_"
	RT코스피_체결        = "S3_"
	RT코스피_예상_체결     = "YS3"
	RT코스피_ETF_NAV   = "I5_"
	RT주식_VI발동해제     = "VI_"
	RT시간외_단일가VI발동해제 = "DVI"
	RT장_운영정보        = "JIF"

	// 미구현 TR코드
	TR주식_매매일지_수수료_금일      = "t0150"
	TR주식_매매일지_수수료_날짜_지정   = "t0151"
	TR주식_잔고_2             = "t0424"
	TR주식_체결_미체결           = "t0425"
	TR종목별_증시_일정           = "t3202"
	TR해외_실시간_지수           = "t3518"
	TR해외_지수_조회            = "t3521"
	TR현물계좌_예수금_주문가능금액_총평가 = "CSPAQ12200"
	TR현물계좌_잔고내역           = "CSPAQ12300"
	TR현물계좌_주문체결내역         = "CSPAQ13700"
	TR계좌별_신용한도            = "CSPAQ00600"
	TR현물계좌_증거금률별_주문가능수량   = "CSPBQ00200"
	TR주식계좌_기간별_수익률_상세     = "FOCCQ33600"

	// 미구현 RT코드
	RT코스닥_체결         = "K3_"
	RT코스피_거래원        = "K1_"
	RT코스닥_거래원        = "OK_"
	RT코스피_기세         = "S4_"
	RT코스닥_LP호가       = "B7_"
	RT코스닥_호가잔량       = "HA_"
	RT코스닥_시간외_호가잔량   = "HB_"
	RT지수             = "IJ_"
	RT예상지수           = "YJ_"
	RT코스닥_예상_체결      = "YK3"
	RT실시간_뉴스_제목_패킷   = "NWS"
	RT업종별_투자자별_매매_현황 = "BM_"
)

const (
	P주문_응답_신규_주문 T주문_응답_구분 = iota
	P주문_응답_정정_주문
	P주문_응답_취소_주문
	P주문_응답_체결_확인
)

type T주문_응답_구분 uint8

func (v T주문_응답_구분) String() string {
	switch v {
	case P주문_응답_신규_주문:
		return "신규 주문"
	case P주문_응답_정정_주문:
		return "정정 주문"
	case P주문_응답_취소_주문:
		return "취소 주문"
	case P주문_응답_체결_확인:
		return "체결 확인"
	default:
		return lib.F2문자열("예상하지 못한 값 : '%v'", uint8(v))
	}
}

const (
	P매도 = T매수_매도("1")
	P매수 = T매수_매도("2")
)

type T매수_매도 string

func (p T매수_매도) String() string {
	switch p {
	case P매도:
		return "매도"
	case P매수:
		return "매수"
	default:
		return lib.F2문자열("예상하지 못한 매수 매도 구분. '%v'", string(p))
	}
}

const (
	P지정가        = T호가유형("00")
	P시장가        = T호가유형("03")
	P조건부_지정가    = T호가유형("05")
	P최유리_지정가    = T호가유형("06")
	P최우선_지정가    = T호가유형("07")
	P시간외종가_장개시전 = T호가유형("61")
	P시간외종가      = T호가유형("81")
	P시간외단일가     = T호가유형("82")
)

type T호가유형 string

func (p T호가유형) String() string {
	switch p {
	case P지정가:
		return "지정가"
	case P시장가:
		return "시장가"
	case P조건부_지정가:
		return "조건부 지정가"
	case P최유리_지정가:
		return "최유리 지정가"
	case P최우선_지정가:
		return "최우선_지정가"
	case P시간외종가_장개시전:
		return "장개시전"
	case P시간외종가:
		return "시간외 종가"
	case P시간외단일가:
		return "시간외 단일가"
	}

	panic(lib.F2문자열("예상하지 못한 호가유형. '%v'", string(p)))

	return ""
}

const (
	P신용거래_아님  = T신용거래_구분("000")
	P유통융자신규   = T신용거래_구분("001")
	P자기융자신규   = T신용거래_구분("003")
	P유통대주신규   = T신용거래_구분("005")
	P자기대주신규   = T신용거래_구분("007")
	P유통융자상환   = T신용거래_구분("101")
	P자기융자상환   = T신용거래_구분("103")
	P유통대주상환   = T신용거래_구분("105")
	P자기대주상환   = T신용거래_구분("107")
	P예탁담보대출상환 = T신용거래_구분("180")
)

type T신용거래_구분 string

func (p T신용거래_구분) String() string {
	switch p {
	case P신용거래_아님:
		return "신용거래 아님"
	case P유통융자신규:
		return "유통융자신규"
	case P자기융자신규:
		return "자기융자신규"
	case P유통대주신규:
		return "유통대주신규"
	case P자기대주신규:
		return "자기대주신규"
	case P유통융자상환:
		return "유통융자상환"
	case P자기융자상환:
		return "자기융자상환"
	case P유통대주상환:
		return "유통대주상환"
	case P자기대주상환:
		return "자기대주상환"
	case P예탁담보대출상환:
		return "예탁담보대출상환"
	}

	panic(lib.F2문자열("예상하지 못한 신용거래 구분. '%v'", string(p)))

	return ""
}

const (
	P주문조건_없음  = T주문조건("0")
	P주문조건_IOC = T주문조건("1")
	P주문조건_FOK = T주문조건("2")
)

type T주문조건 string

func (p T주문조건) String() string {
	switch p {
	case P주문조건_없음:
		return "주문조건 없음"
	case P주문조건_IOC:
		return "IOC"
	case P주문조건_FOK:
		return "FOK"
	}

	panic(lib.F2문자열("예상하지 못한 주문조건. '%v'", p))

	return ""
}

const (
	P동시호가_아님  = T동시호가_구분(0)
	P동시호가_장중  = T동시호가_구분(1)
	P동시호가_시간외 = T동시호가_구분(2)
	P동시호가_동시  = T동시호가_구분(3)
)

type T동시호가_구분 uint8

func (s T동시호가_구분) String() string {
	switch s {
	case P동시호가_아님:
		return "동시호가 아님"
	case P동시호가_장중:
		return "장중"
	case P동시호가_시간외:
		return "시간외"
	case P동시호가_동시:
		return "동시호가"
	}

	return lib.F2문자열("예상하지 못한 동시호가 구분. '%v'", int(s))

	return ""
}

const (
	P구분_상한 T전일대비_구분 = iota + 1
	P구분_상승
	P구분_보합
	P구분_하한
	P구분_하락
)

type T전일대비_구분 uint8

func (s T전일대비_구분) G부호보정_정수64(등락폭 int64) int64 {
	switch s {
	case P구분_상한, P구분_상승:
		if 등락폭 < 0 {
			등락폭 = 등락폭 * -1
		}
	//case P구분_보합:
	//	return "보합"
	case P구분_하한, P구분_하락:
		if 등락폭 > 0 {
			등락폭 = 등락폭 * -1
		}
	}

	return 등락폭
}

func (s T전일대비_구분) G부호보정_실수64(등락율 float64) float64 {
	switch s {
	case P구분_상한, P구분_상승:
		if 등락율 < 0.0 {
			등락율 = 등락율 * -1
		}
	//case P구분_보합:
	//	return "보합"
	case P구분_하한, P구분_하락:
		if 등락율 > 0.0 {
			등락율 = 등락율 * -1
		}
	}

	return 등락율
}

func (s T전일대비_구분) String() string {
	switch s {
	case P구분_상한:
		return "상한"
	case P구분_상승:
		return "상승"
	case P구분_보합:
		return "보합"
	case P구분_하한:
		return "하한"
	case P구분_하락:
		return "하락"
	}

	return strconv.Itoa(int(uint8(s)))
}

const (
	P당일전일구분_당일 = T당일전일_구분(0)
	P당일전일구분_전일 = T당일전일_구분(1)
)

type T당일전일_구분 uint8

func (s T당일전일_구분) String() string {
	switch s {
	case P당일전일구분_당일:
		return "당일"
	case P당일전일구분_전일:
		return "전일"
	default:
		return lib.F2문자열("예상하지 못한 당일전일 구분. '%v'", int(s))
	}
}

const (
	P분틱구분_분 = T분틱_구분(0)
	P분틱구분_틱 = T분틱_구분(1)
)

type T분틱_구분 uint8

func (s T분틱_구분) String() string {
	switch s {
	case P분틱구분_분:
		return "분"
	case P분틱구분_틱:
		return "틱"
	}

	panic(lib.F2문자열("예상하지 못한 분틱 구분. '%v'", s))

	return ""
}

// XingAPI 에러코드
const (
	P에러_소켓_생성_실패          = -1
	P에러_서버_연결_실패          = -2
	P에러_잘못된_서버_주소         = -3
	P에러_서버_연결시간_초과        = -4
	P에러_이미_서버에_연결_중       = -5
	P에러_사용불가_TR           = -6
	P에러_로그인_필요            = -7
	P에러_시세전용_모드에서_사용불가    = -8
	P에러_잘못된_계좌번호          = -9
	P에러_잘못된_패킷_크기         = -10
	P에러_잘못된_데이터_길이        = -11
	P에러_존재하지_않는_계좌        = -12
	P에러_Request_ID_부족     = -13
	P에러_소켓_미생성            = -14
	P에러_암호화_생성_실패         = -15
	P에러_데이터_전송_실패         = -16
	P에러_암호화_RTN_처리_실패     = -17
	P에러_공인인증_파일_없음        = -18
	P에러_공인인증_Function_없음  = -19
	P에러_메모리_부족            = -20
	P에러_TR쿼터_초과           = -21
	P에러_TR_함수_미적용         = -22
	P에러_TR정보_없음           = -23
	P에러_계좌위치_미지정          = -24
	P에러_계좌_없음             = -25
	P에러_파일_읽기_실패          = -26 // (종목 검색 조회 시, 파일이 없는 경우)
	P에러_실시간_종목검색_쿼터_초과    = -27
	P에러_API_HTS_종목_연동키_오류 = -28 // 등록 키에 대한 정보를 찾을 수 없습니다
)

const (
	P서버_실거래 T서버_구분 = iota
	P서버_모의투자
	P서버_XingACE
)

type T서버_구분 int

func (p T서버_구분) String() string {
	switch p {
	case P서버_실거래:
		return "hts.ebestsec.co.kr"
	case P서버_모의투자:
		return "demo.ebestsec.co.kr"
	case P서버_XingACE:
		return "127.0.0.1"
	}

	panic(lib.F2문자열("예상하지 못한 서버 구분값. %v", p))

	return ""
}

const (
	P시장상태_장전동시호가개시   = T시장상태(11)
	P시장상태_장시작        = T시장상태(21)
	P시장상태_장개시10초전    = T시장상태(22)
	P시장상태_장개시1분전     = T시장상태(23)
	P시장상태_장개시5분전     = T시장상태(24)
	P시장상태_장개시10분전    = T시장상태(25)
	P시장상태_장후동시호가개시   = T시장상태(31)
	P시장상태_장마감        = T시장상태(41)
	P시장상태_장마감10초전    = T시장상태(42)
	P시장상태_장마감1분전     = T시장상태(43)
	P시장상태_장마감5분전     = T시장상태(44)
	P시장상태_시간외종가매매개시  = T시장상태(51)
	P시장상태_시간외종가매매종료  = T시장상태(52)
	P시장상태_시간외단일가매매개시 = T시장상태(53)
	P시장상태_시간외단일가매매종료 = T시장상태(54)
)

type T시장상태 uint8

func (p T시장상태) String() string {
	switch p {
	case P시장상태_장전동시호가개시:
		return "장전동시호가개시"
	case P시장상태_장시작:
		return "장시작"
	case P시장상태_장개시10초전:
		return "장개시10초전"
	case P시장상태_장개시1분전:
		return "장개시1분전"
	case P시장상태_장개시5분전:
		return "장개시5분전"
	case P시장상태_장개시10분전:
		return "장개시10분전"
	case P시장상태_장후동시호가개시:
		return "장후동시호가개시"
	case P시장상태_장마감:
		return "장마감"
	case P시장상태_장마감10초전:
		return "장마감10초전"
	case P시장상태_장마감1분전:
		return "장마감1분전"
	case P시장상태_장마감5분전:
		return "장마감5분전"
	case P시장상태_시간외종가매매개시:
		return "시간외종가매매개시"
	case P시장상태_시간외종가매매종료:
		return "시간외종가매매종료"
	case P시장상태_시간외단일가매매개시:
		return "시간외단일가매매개시"
	case P시장상태_시간외단일가매매종료:
		return "시간외단일가매매종료"
	}

	panic(lib.F2문자열("예상하지 못한 시장상태. %v", p))

	return ""
}

const (
	P주문_시장구분_비상장    = T주문_시장구분(0)
	P주문_시장구분_코스피    = T주문_시장구분(10)
	P주문_시장구분_채권     = T주문_시장구분(11)
	P주문_시장구분_장외시장   = T주문_시장구분(19)
	P주문_시장구분_코스닥    = T주문_시장구분(20)
	P주문_시장구분_코넥스    = T주문_시장구분(23)
	P주문_시장구분_프리보드   = T주문_시장구분(30)
	P주문_시장구분_동경거래소  = T주문_시장구분(61)
	P주문_시장구분_JASDAQ = T주문_시장구분(62)
)

type T주문_시장구분 uint8

func (p T주문_시장구분) String() string {
	switch p {
	case P주문_시장구분_비상장:
		return "비상장"
	case P주문_시장구분_코스피:
		return "코스피"
	case P주문_시장구분_채권:
		return "채권"
	case P주문_시장구분_장외시장:
		return "장외시장"
	case P주문_시장구분_코스닥:
		return "코스닥"
	case P주문_시장구분_코넥스:
		return "코넥스"
	case P주문_시장구분_프리보드:
		return "프리보드"
	case P주문_시장구분_동경거래소:
		return "동경거래소"
	case P주문_시장구분_JASDAQ:
		return "JASDAQ"
	}

	panic(lib.F2문자열("예상하지 못한 주문_시장구분. '%v'", p))

	return ""
}

const (
	P증권그룹_주식           T증권그룹 = T증권그룹(1)
	P증권그룹_예탁증서         T증권그룹 = T증권그룹(3)
	P증권그룹_증권투자회사_뮤추얼펀드 T증권그룹 = T증권그룹(4)
	P증권그룹_Reits종목      T증권그룹 = T증권그룹(6)
	P증권그룹_상장지수펀드_ETF   T증권그룹 = T증권그룹(8)
	P증권그룹_선박투자회사       T증권그룹 = T증권그룹(10)
	P증권그룹_인프라투융자회사     T증권그룹 = T증권그룹(12)
	P증권그룹_해외ETF        T증권그룹 = T증권그룹(13)
	P증권그룹_해외원주         T증권그룹 = T증권그룹(14)
	P증권그룹_ETN          T증권그룹 = T증권그룹(15)
)

type T증권그룹 uint8

func (p T증권그룹) String() string {
	switch p {
	case P증권그룹_주식:
		return "주식"
	case P증권그룹_예탁증서:
		return "예탁증서"
	case P증권그룹_증권투자회사_뮤추얼펀드:
		return "증권투자회사_뮤추얼펀드"
	case P증권그룹_Reits종목:
		return "Reits종목"
	case P증권그룹_상장지수펀드_ETF:
		return "상장지수펀드_ETF"
	case P증권그룹_선박투자회사:
		return "선박투자회사"
	case P증권그룹_인프라투융자회사:
		return "인프라투융자회사"
	case P증권그룹_해외ETF:
		return "해외ETF"
	case P증권그룹_해외원주:
		return "해외원주"
	case P증권그룹_ETN:
		return "ETN"
	}

	panic(lib.F2문자열("예상하지 못한 증권그룹 값. %v", p))

	return ""
}

func (p T증권그룹) XingCode() string {
	코드 := strconv.Itoa(int(uint8(p)))

	if len(코드) < 2 {
		코드 = "0" + 코드
	}

	return 코드
}

const (
	P일주월_일 T일주월_구분 = iota + 1
	P일주월_주
	P일주월_월
)

type T일주월_구분 uint8

func (p T일주월_구분) String() string {
	switch p {
	case P일주월_일:
		return "일"
	case P일주월_주:
		return "주"
	case P일주월_월:
		return "월"
	}

	panic(lib.F2문자열("예상하지 못한 일주월 구분. '%v'", p))

	return ""
}

const (
	VI해제 VI발동해제 = iota
	VI정적발동
	VI동적발동
)

type VI발동해제 uint8

func (p VI발동해제) String() string {
	switch p {
	case VI해제:
		return "일"
	case VI정적발동:
		return "주"
	case VI동적발동:
		return "월"
	}

	panic(lib.F2문자열("예상하지 못한 VI발동해제 구분. '%v'", p))

	return ""
}

const (
	P시장구분_코스피         = T시장구분("1")
	P시장구분_코스닥         = T시장구분("2")
	P시장구분_선물_옵션       = T시장구분("5")
	P시장구분_CME야간선물     = T시장구분("7")
	P시장구분_EUREX야간선물옵션 = T시장구분("8")
	P시장구분_미국주식        = T시장구분("9")
	P시장구분_중국주식_오전     = T시장구분("A")
	P시장구분_중국주식_오후     = T시장구분("B")
	P시장구분_홍콩주식_오전     = T시장구분("C")
	P시장구분_홍콩주식_오후     = T시장구분("D")
)

type T시장구분 string

func (p T시장구분) String() string {
	switch p {
	case P시장구분_코스피:
		return "코스피"
	case P시장구분_코스닥:
		return "코스닥"
	case P시장구분_선물_옵션:
		return "선물_옵션"
	case P시장구분_CME야간선물:
		return "CME야간선물"
	case P시장구분_EUREX야간선물옵션:
		return "EUREX야간선물옵션"
	case P시장구분_미국주식:
		return "미국주식"
	case P시장구분_중국주식_오전:
		return "중국주식 오전"
	case P시장구분_중국주식_오후:
		return "중국주식 오후"
	case P시장구분_홍콩주식_오전:
		return "홍콩주식 오전"
	case P시장구분_홍콩주식_오후:
		return "홍콩주식 오후"
	}

	panic(lib.F2문자열("예상하지 못한 시장구분. '%v'", p))

	return ""
}

const (
	P자료형_S현물_주문_응답_실시간_정보 = "S현물_주문_응답_실시간_정보"

	P자료형_S질의값_정상_주문         = "S질의값_정상_주문"
	P자료형_S질의값_정정_주문         = "S질의값_정정_주문"
	P자료형_S질의값_취소_주문         = "S질의값_취소_주문"
	P자료형_S질의값_현물_전일당일_분틱_조회 = "S질의값_현물_전일당일_분틱_조회"
	P자료형_S질의값_현물_기간별_조회     = "S질의값_현물_기간별_조회"
	P자료형_S질의값_단일종목_연속키      = "S질의값_단일종목_연속키"
	P자료형_S질의값_현물_차트_틱       = "S질의값_현물_차트_틱"
	P자료형_S질의값_현물_차트_분       = "S질의값_현물_차트_분"
	P자료형_S질의값_증시주변자금추이      = "S질의값_증시주변자금추이"

	P자료형_S콜백_기본형      = "S콜백_기본형"
	P자료형_S콜백_정수값      = "S콜백_정수값"
	P자료형_S콜백_문자열      = "S콜백_문자열"
	P자료형_S콜백_TR데이터    = "S콜백_TR데이터"
	P자료형_S콜백_메시지_및_에러 = "S콜백_메시지_및_에러"

	P자료형_S현물_정상_주문_응답  = "S현물_정상_주문_응답"
	P자료형_S현물_정상_주문_응답1 = "S현물_정상_주문_응답1"
	P자료형_S현물_정상_주문_응답2 = "S현물_정상_주문_응답2"
	P자료형_S현물_정정_주문_응답  = "S현물_정정_주문_응답"
	P자료형_S현물_정정_주문_응답1 = "S현물_정정_주문_응답1"
	P자료형_S현물_정정_주문_응답2 = "S현물_정정_주문_응답2"
	P자료형_S현물_취소_주문_응답  = "S현물_취소_주문_응답"
	P자료형_S현물_취소_주문_응답1 = "S현물_취소_주문_응답1"
	P자료형_S현물_취소_주문_응답2 = "S현물_취소_주문_응답2"

	P자료형_S현물_호가조회_응답            = "S현물_호가조회_응답"
	P자료형_S현물_시세조회_응답            = "S현물_시세조회_응답"
	P자료형_S현물_시간대별_체결_응답         = "S현물_시간대별_체결_응답"
	P자료형_S현물_시간대별_체결_응답_헤더      = "S현물_시간대별_체결_응답_헤더"
	P자료형_S현물_시간대별_체결_응답_반복값     = "S현물_시간대별_체결_응답_반복값"
	P자료형_S현물_시간대별_체결_응답_반복값_모음  = "S현물_시간대별_체결_응답_반복값_모음"
	P자료형_S현물_기간별_조회_응답          = "S현물_기간별_조회_응답"
	P자료형_S현물_기간별_조회_응답_헤더       = "S현물_기간별_조회_응답_헤더"
	P자료형_S현물_기간별_조회_응답_반복값      = "S현물_기간별_조회_응답_반복값"
	P자료형_S현물_기간별_조회_응답_반복값_모음   = "S현물_기간별_조회_응답_반복값_모음"
	P자료형_S현물_전일당일분틱조회_응답        = "S현물_전일당일분틱조회_응답"
	P자료형_S현물_전일당일분틱조회_응답_헤더     = "S현물_전일당일분틱조회_응답_헤더"
	P자료형_S현물_전일당일분틱조회_응답_반복값    = "S현물_전일당일분틱조회_응답_반복값"
	P자료형_S현물_전일당일분틱조회_응답_반복값_모음 = "S현물_전일당일분틱조회_응답_반복값_모음"
	P자료형_S_ETF_현재가_조회_응답        = "S_ETF_현재가_조회_응답"
	P자료형_S_ETF시간별_추이_응답         = "S_ETF시간별_추이_응답"
	P자료형_S_ETF시간별_추이_응답_헤더      = "S_ETF시간별_추이_응답_헤더"
	P자료형_S_ETF시간별_추이_응답_반복값     = "S_ETF시간별_추이_응답_반복값"
	P자료형_S_ETF시간별_추이_응답_반복값_모음  = "S_ETF시간별_추이_응답_반복값_모음"
	P자료형_S현물_차트_틱_응답            = "S현물_차트_틱_응답"
	P자료형_S현물_차트_틱_응답_헤더         = "S현물_차트_틱_응답_헤더"
	P자료형_S현물_차트_틱_응답_반복값        = "S현물_차트_틱_응답_반복값"
	P자료형_S현물_차트_틱_응답_반복값_모음     = "S현물_차트_틱_응답_반복값_모음"
	P자료형_S현물_차트_분_응답            = "S현물_차트_분_응답"
	P자료형_S현물_차트_분_응답_헤더         = "S현물_차트_분_응답_헤더"
	P자료형_S현물_차트_분_응답_반복값        = "S현물_차트_분_응답_반복값"
	P자료형_S현물_차트_분_응답_반복값_모음     = "S현물_차트_분_응답_반복값_모음"
	P자료형_S증시_주변자금추이_응답          = "S증시_주변자금추이_응답"
	P자료형_S증시_주변자금추이_응답_헤더       = "S증시_주변자금추이_응답_헤더"
	P자료형_S증시_주변자금추이_응답_반복값      = "S증시_주변자금추이_응답_반복값"
	P자료형_S증시_주변자금추이_응답_반복값_모음   = "S증시_주변자금추이_응답_반복값_모음"
	P자료형_S현물_종목조회_응답_반복값        = "S현물_종목조회_응답_반복값"
	P자료형_S현물_종목조회_응답_반복값_모음     = "S현물_종목조회_응답_반복값_모음"
)

type T콜백 uint8

const (
	P콜백_TR데이터 = iota
	P콜백_메시지_및_에러
	P콜백_TR완료
	P콜백_타임아웃
	P콜백_링크_데이터
	P콜백_실시간_차트_데이터
	P콜백_신호
)

func (p T콜백) String() string {
	switch p {
	case P콜백_TR데이터:
		return "데이터"
	case P콜백_메시지_및_에러:
		return "메시지 및 에러"
	case P콜백_TR완료:
		return "TR완료"
	case P콜백_타임아웃:
		return "타임아웃"
	case P콜백_링크_데이터:
		return "링크_데이터"
	case P콜백_실시간_차트_데이터:
		return "실시간_차트_데이터"
	case P콜백_신호:
		return "신호"
	default:
		return lib.F2문자열("예상하지 못한 콜백값 : '%v'", p)
	}
}

type T신호_C32 uint8

const (
	P신호_C32_READY = iota
	P신호_C32_종료
)

func (p T신호_C32) String() string {
	switch p {
	case P신호_C32_READY:
		return "C32 READY"
	case P신호_C32_종료:
		return "C32 종료"
	default:
		return lib.F2문자열("예상하지 못한 T신호_C32 값 : '%v'", p)
	}
}

type T수정구분 uint32

const (
	P수정구분_없음 = T수정구분(0x0)
	P수정구분_권리락 = T수정구분(0x1)
	P수정구분_배당락 = T수정구분(0x2)
	P수정구분_액면분할 = T수정구분(0x4)
	P수정구분_액면병합 = T수정구분(0x8)
	P수정구분_주식병합 = T수정구분(0x10)
	P수정구분_기업분할 = T수정구분(0x20)
	P수정구분_관리종목 = T수정구분(0x80)
	P수정구분_투자경고 = T수정구분(0x100)
	P수정구분_거래정지 = T수정구분(0x200)
	P수정구분_기준가조정 = T수정구분(0x1000)
	P수정구분_우선주 = T수정구분(0x4000)
	P수정구분_CB발동예고 = T수정구분(0x8000)
	P수정구분_증거금50_100 = T수정구분(0x200000)
	P수정구분_증거금50 = T수정구분(0x00400000)
	P수정구분_증거금100 = T수정구분(0x00800000)
	P수정구분_정리매매 = T수정구분(0x01000000)
	P수정구분_투자유의 = T수정구분(0x04000000)
	P수정구분_불성실공시 = T수정구분(0x80000000)
)


func (p T수정구분) String() string {
	switch p {
	case P수정구분_없음:
		return "없음"
	case P수정구분_권리락:
		return "권리락"
	case P수정구분_배당락:
		return "배당락"
	case P수정구분_액면분할:
		return "액면분할"
	case P수정구분_액면병합:
		return "액면병합"
	case P수정구분_주식병합:
		return "주식병합"
	case P수정구분_기업분할:
		return "기업분할"
	case P수정구분_관리종목:
		return "관리종목"
	case P수정구분_투자경고:
		return "투자경고"
	case P수정구분_거래정지:
		return "거래정지"
	case P수정구분_기준가조정:
		return "기준가 조정"
	case P수정구분_우선주:
		return "우선주"
	case P수정구분_CB발동예고:
		return "CB발동 예고"
	case P수정구분_증거금50_100:
		return "증거금 50/100"
	case P수정구분_증거금50:
		return "증거금 50"
	case P수정구분_증거금100:
		return "증거금 100"
	case P수정구분_정리매매:
		return "정리매매"
	case P수정구분_투자유의:
		return "투자유의"
	case P수정구분_불성실공시:
		return "불성실 공시"
	default:
		return lib.F2문자열("예상하지 못한 값 : '%v'", p)
	}
}
