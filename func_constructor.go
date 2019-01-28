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
	"math"
	"strconv"
	"strings"
)

func New현물_호가조회_응답(b []byte) (s *S현물_호가조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT1101OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T1101OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	s = new(S현물_호가조회_응답)
	s.M한글명 = lib.F2문자열_EUC_KR(g.Hname)
	s.M현재가 = lib.F2정수64_단순형(g.Price)
	s.M전일대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Sign))
	s.M전일대비등락폭 = lib.F2정수64_단순형(g.Change)
	s.M등락율 = lib.F2실수_소숫점_추가(g.Diff, 2)
	s.M누적거래량 = lib.F2정수64_단순형(g.Volume)
	s.M전일종가 = lib.F2정수64_단순형(g.Jnilclose)
	s.M매도호가_모음 = make([]int64, 10)
	s.M매수호가_모음 = make([]int64, 10)
	s.M매도호가수량_모음 = make([]int64, 10)
	s.M매수호가수량_모음 = make([]int64, 10)
	s.M직전매도대비수량_모음 = make([]int64, 10)
	s.M직전매수대비수량_모음 = make([]int64, 10)
	s.M매도호가_모음[0] = lib.F2정수64_단순형(g.Offerho1)
	s.M매수호가_모음[0] = lib.F2정수64_단순형(g.Bidho1)
	s.M매도호가수량_모음[0] = lib.F2정수64_단순형(g.Offerrem1)
	s.M매수호가수량_모음[0] = lib.F2정수64_단순형(g.Bidrem1)
	s.M직전매도대비수량_모음[0] = lib.F2정수64_단순형(g.Preoffercha1)
	s.M직전매수대비수량_모음[0] = lib.F2정수64_단순형(g.Prebidcha1)
	s.M매도호가_모음[1] = lib.F2정수64_단순형(g.Offerho2)
	s.M매수호가_모음[1] = lib.F2정수64_단순형(g.Bidho2)
	s.M매도호가수량_모음[1] = lib.F2정수64_단순형(g.Offerrem2)
	s.M매수호가수량_모음[1] = lib.F2정수64_단순형(g.Bidrem2)
	s.M직전매도대비수량_모음[1] = lib.F2정수64_단순형(g.Preoffercha2)
	s.M직전매수대비수량_모음[1] = lib.F2정수64_단순형(g.Prebidcha2)
	s.M매도호가_모음[2] = lib.F2정수64_단순형(g.Offerho3)
	s.M매수호가_모음[2] = lib.F2정수64_단순형(g.Bidho3)
	s.M매도호가수량_모음[2] = lib.F2정수64_단순형(g.Offerrem3)
	s.M매수호가수량_모음[2] = lib.F2정수64_단순형(g.Bidrem3)
	s.M직전매도대비수량_모음[2] = lib.F2정수64_단순형(g.Preoffercha3)
	s.M직전매수대비수량_모음[2] = lib.F2정수64_단순형(g.Prebidcha3)
	s.M매도호가_모음[3] = lib.F2정수64_단순형(g.Offerho4)
	s.M매수호가_모음[3] = lib.F2정수64_단순형(g.Bidho4)
	s.M매도호가수량_모음[3] = lib.F2정수64_단순형(g.Offerrem4)
	s.M매수호가수량_모음[3] = lib.F2정수64_단순형(g.Bidrem4)
	s.M직전매도대비수량_모음[3] = lib.F2정수64_단순형(g.Preoffercha4)
	s.M직전매수대비수량_모음[3] = lib.F2정수64_단순형(g.Prebidcha4)
	s.M매도호가_모음[4] = lib.F2정수64_단순형(g.Offerho5)
	s.M매수호가_모음[4] = lib.F2정수64_단순형(g.Bidho5)
	s.M매도호가수량_모음[4] = lib.F2정수64_단순형(g.Offerrem5)
	s.M매수호가수량_모음[4] = lib.F2정수64_단순형(g.Bidrem5)
	s.M직전매도대비수량_모음[4] = lib.F2정수64_단순형(g.Preoffercha5)
	s.M직전매수대비수량_모음[4] = lib.F2정수64_단순형(g.Prebidcha5)
	s.M매도호가_모음[5] = lib.F2정수64_단순형(g.Offerho6)
	s.M매수호가_모음[5] = lib.F2정수64_단순형(g.Bidho6)
	s.M매도호가수량_모음[5] = lib.F2정수64_단순형(g.Offerrem6)
	s.M매수호가수량_모음[5] = lib.F2정수64_단순형(g.Bidrem6)
	s.M직전매도대비수량_모음[5] = lib.F2정수64_단순형(g.Preoffercha6)
	s.M직전매수대비수량_모음[5] = lib.F2정수64_단순형(g.Prebidcha6)
	s.M매도호가_모음[6] = lib.F2정수64_단순형(g.Offerho7)
	s.M매수호가_모음[6] = lib.F2정수64_단순형(g.Bidho7)
	s.M매도호가수량_모음[6] = lib.F2정수64_단순형(g.Offerrem7)
	s.M매수호가수량_모음[6] = lib.F2정수64_단순형(g.Bidrem7)
	s.M직전매도대비수량_모음[6] = lib.F2정수64_단순형(g.Preoffercha7)
	s.M직전매수대비수량_모음[6] = lib.F2정수64_단순형(g.Prebidcha7)
	s.M매도호가_모음[7] = lib.F2정수64_단순형(g.Offerho8)
	s.M매수호가_모음[7] = lib.F2정수64_단순형(g.Bidho8)
	s.M매도호가수량_모음[7] = lib.F2정수64_단순형(g.Offerrem8)
	s.M매수호가수량_모음[7] = lib.F2정수64_단순형(g.Bidrem8)
	s.M직전매도대비수량_모음[7] = lib.F2정수64_단순형(g.Preoffercha8)
	s.M직전매수대비수량_모음[7] = lib.F2정수64_단순형(g.Prebidcha8)
	s.M매도호가_모음[8] = lib.F2정수64_단순형(g.Offerho9)
	s.M매수호가_모음[8] = lib.F2정수64_단순형(g.Bidho9)
	s.M매도호가수량_모음[8] = lib.F2정수64_단순형(g.Offerrem9)
	s.M매수호가수량_모음[8] = lib.F2정수64_단순형(g.Bidrem9)
	s.M직전매도대비수량_모음[8] = lib.F2정수64_단순형(g.Preoffercha9)
	s.M직전매수대비수량_모음[8] = lib.F2정수64_단순형(g.Prebidcha9)
	s.M매도호가_모음[9] = lib.F2정수64_단순형(g.Offerho10)
	s.M매수호가_모음[9] = lib.F2정수64_단순형(g.Bidho10)
	s.M매도호가수량_모음[9] = lib.F2정수64_단순형(g.Offerrem10)
	s.M매수호가수량_모음[9] = lib.F2정수64_단순형(g.Bidrem10)
	s.M직전매도대비수량_모음[9] = lib.F2정수64_단순형(g.Preoffercha10)
	s.M직전매수대비수량_모음[9] = lib.F2정수64_단순형(g.Prebidcha10)
	s.M매도호가수량합 = lib.F2정수64_단순형(g.Offer)
	s.M매수호가수량합 = lib.F2정수64_단순형(g.Bid)
	s.M직전매도대비수량합 = lib.F2정수64_단순형(g.Preoffercha)
	s.M직전매수대비수량합 = lib.F2정수64_단순형(g.Prebidcha)

	if 시각_문자열 := lib.F2문자열_공백제거(g.Hotime); 시각_문자열 != "" {
		s.M수신시간 = lib.F2일자별_시각_단순형(당일.G값(), "150405.999", 시각_문자열[:6]+"."+시각_문자열[6:])
	}

	s.M예상체결가격 = lib.F2정수64_단순형(g.Yeprice)
	s.M예상체결수량 = lib.F2정수64_단순형(g.Yevolume)
	s.M예상체결전일구분 = T전일대비_구분(lib.F2정수64_단순형(g.Yesign))
	s.M예상체결전일대비 = lib.F2정수64_단순형(g.Yechange)
	s.M예상체결등락율 = lib.F2실수_소숫점_추가(g.Yediff, 2)
	s.M시간외매도잔량 = lib.F2정수64_단순형(g.Tmoffer)
	s.M시간외매수잔량 = lib.F2정수64_단순형(g.Tmbid)
	s.M동시호가_구분 = T동시호가_구분(lib.F2정수64_단순형(g.Status))
	s.M종목코드 = lib.F2문자열_공백제거(g.Shcode)
	s.M상한가 = lib.F2정수64_단순형(g.Uplmtprice)
	s.M하한가 = lib.F2정수64_단순형(g.Dnlmtprice)
	s.M시가 = lib.F2정수64_단순형(g.Open)
	s.M고가 = lib.F2정수64_단순형(g.High)
	s.M저가 = lib.F2정수64_단순형(g.Low)

	return s, nil
}

func New현물_시세조회_응답(b []byte) (s *S현물_시세조회_응답, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT1102OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T1102OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	s = new(S현물_시세조회_응답)
	s.M종목코드 = lib.F2문자열_공백제거(g.Shcode)
	s.M한글명 = lib.F2문자열_EUC_KR(g.Hname)
	s.M현재가 = lib.F2정수64_단순형(g.Price)
	s.M전일대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Sign))
	s.M전일대비등락폭 = lib.F2정수64_단순형(g.Change)
	s.M등락율 = lib.F2실수_소숫점_추가(g.Diff, 2)
	s.M누적거래량 = lib.F2정수64_단순형(g.Volume)
	s.M기준가 = lib.F2정수64_단순형(g.Recprice)
	s.M가중평균 = lib.F2정수64_단순형(g.Avg)
	s.M상한가 = lib.F2정수64_단순형(g.Uplmtprice)
	s.M하한가 = lib.F2정수64_단순형(g.Dnlmtprice)
	s.M전일거래량 = lib.F2정수64_단순형(g.Jnilvolume)
	s.M거래량차 = lib.F2정수64_단순형(g.Volumediff)
	s.M시가 = lib.F2정수64_단순형(g.Open)
	s.M시가시간 = lib.F2일자별_시각_단순형_공백은_초기값(당일.G값(), "150405", g.Opentime)
	s.M고가 = lib.F2정수64_단순형(g.High)
	s.M고가시간 = lib.F2일자별_시각_단순형_공백은_초기값(당일.G값(), "150405", g.Hightime)
	s.M저가 = lib.F2정수64_단순형(g.Low)
	s.M저가시간 = lib.F2일자별_시각_단순형_공백은_초기값(당일.G값(), "150405", g.Lowtime)
	s.M52주_최고가 = lib.F2정수64_단순형(g.High52w)
	s.M52주_최고가일 = lib.F2포맷된_시각_단순형_공백은_초기값("20060102", g.High52wdate)
	s.M52주_최저가 = lib.F2정수64_단순형(g.Low52w)
	s.M52주_최저가일 = lib.F2포맷된_시각_단순형_공백은_초기값("20060102", g.Low52wdate)
	s.M소진율 = lib.F2실수_소숫점_추가(g.Exhratio, 2)
	s.PER = lib.F2실수_소숫점_추가(g.Per, 2)
	s.PBRX = lib.F2실수_소숫점_추가(g.Pbrx, 2)
	s.M상장주식수_천 = lib.F2정수64_단순형(g.Listing)
	s.M증거금율 = lib.F2정수64_단순형(g.Jkrate)
	s.M수량단위 = lib.F2정수64_단순형(g.Memedan)
	s.M매도증권사코드_모음 = make([]string, 5)
	s.M매수증권사코드_모음 = make([]string, 5)
	s.M매도증권사명_모음 = make([]string, 5)
	s.M매수증권사명_모음 = make([]string, 5)
	s.M총매도수량_모음 = make([]int64, 5)
	s.M총매수수량_모음 = make([]int64, 5)
	s.M매도증감_모음 = make([]int64, 5)
	s.M매수증감_모음 = make([]int64, 5)
	s.M매도비율_모음 = make([]float64, 5)
	s.M매수비율_모음 = make([]float64, 5)
	s.M매도증권사코드_모음[0] = lib.F2문자열_공백제거(g.Offernocd1)
	s.M매수증권사코드_모음[0] = lib.F2문자열_공백제거(g.Bidnocd1)
	s.M매도증권사명_모음[0] = lib.F2문자열_EUC_KR(g.Offerno1)
	s.M매수증권사명_모음[0] = lib.F2문자열_EUC_KR(g.Bidno1)
	s.M총매도수량_모음[0] = lib.F2정수64_단순형(g.Dvol1)
	s.M총매수수량_모음[0] = lib.F2정수64_단순형(g.Svol1)
	s.M매도증감_모음[0] = lib.F2정수64_단순형(g.Dcha1)
	s.M매수증감_모음[0] = lib.F2정수64_단순형(g.Scha1)
	s.M매도비율_모음[0] = lib.F2실수_소숫점_추가(g.Ddiff1, 2)
	s.M매수비율_모음[0] = lib.F2실수_소숫점_추가(g.Sdiff1, 2)
	s.M매도증권사코드_모음[1] = lib.F2문자열_공백제거(g.Offernocd2)
	s.M매수증권사코드_모음[1] = lib.F2문자열_공백제거(g.Bidnocd2)
	s.M매도증권사명_모음[1] = lib.F2문자열_EUC_KR(g.Offerno2)
	s.M매수증권사명_모음[1] = lib.F2문자열_EUC_KR(g.Bidno2)
	s.M총매도수량_모음[1] = lib.F2정수64_단순형(g.Dvol2)
	s.M총매수수량_모음[1] = lib.F2정수64_단순형(g.Svol2)
	s.M매도증감_모음[1] = lib.F2정수64_단순형(g.Dcha2)
	s.M매수증감_모음[1] = lib.F2정수64_단순형(g.Scha2)
	s.M매도비율_모음[1] = lib.F2실수_소숫점_추가(g.Ddiff2, 2)
	s.M매수비율_모음[1] = lib.F2실수_소숫점_추가(g.Sdiff2, 2)
	s.M매도증권사코드_모음[2] = lib.F2문자열_공백제거(g.Offernocd3)
	s.M매수증권사코드_모음[2] = lib.F2문자열_공백제거(g.Bidnocd3)
	s.M매도증권사명_모음[2] = lib.F2문자열_EUC_KR(g.Offerno3)
	s.M매수증권사명_모음[2] = lib.F2문자열_EUC_KR(g.Bidno3)
	s.M총매도수량_모음[2] = lib.F2정수64_단순형(g.Dvol3)
	s.M총매수수량_모음[2] = lib.F2정수64_단순형(g.Svol3)
	s.M매도증감_모음[2] = lib.F2정수64_단순형(g.Dcha3)
	s.M매수증감_모음[2] = lib.F2정수64_단순형(g.Scha3)
	s.M매도비율_모음[2] = lib.F2실수_소숫점_추가(g.Ddiff3, 2)
	s.M매수비율_모음[2] = lib.F2실수_소숫점_추가(g.Sdiff3, 2)
	s.M매도증권사코드_모음[3] = lib.F2문자열_공백제거(g.Offernocd4)
	s.M매수증권사코드_모음[3] = lib.F2문자열_공백제거(g.Bidnocd4)
	s.M매도증권사명_모음[3] = lib.F2문자열_EUC_KR(g.Offerno4)
	s.M매수증권사명_모음[3] = lib.F2문자열_EUC_KR(g.Bidno4)
	s.M총매도수량_모음[3] = lib.F2정수64_단순형(g.Dvol4)
	s.M총매수수량_모음[3] = lib.F2정수64_단순형(g.Svol4)
	s.M매도증감_모음[3] = lib.F2정수64_단순형(g.Dcha4)
	s.M매수증감_모음[3] = lib.F2정수64_단순형(g.Scha4)
	s.M매도비율_모음[3] = lib.F2실수_소숫점_추가(g.Ddiff4, 2)
	s.M매수비율_모음[3] = lib.F2실수_소숫점_추가(g.Sdiff4, 2)
	s.M매도증권사코드_모음[4] = lib.F2문자열_공백제거(g.Offernocd5)
	s.M매수증권사코드_모음[4] = lib.F2문자열_공백제거(g.Bidnocd5)
	s.M매도증권사명_모음[4] = lib.F2문자열_EUC_KR(g.Offerno5)
	s.M매수증권사명_모음[4] = lib.F2문자열_EUC_KR(g.Bidno5)
	s.M총매도수량_모음[4] = lib.F2정수64_단순형(g.Dvol5)
	s.M총매수수량_모음[4] = lib.F2정수64_단순형(g.Svol5)
	s.M매도증감_모음[4] = lib.F2정수64_단순형(g.Dcha5)
	s.M매수증감_모음[4] = lib.F2정수64_단순형(g.Scha5)
	s.M매도비율_모음[4] = lib.F2실수_소숫점_추가(g.Ddiff5, 2)
	s.M매수비율_모음[4] = lib.F2실수_소숫점_추가(g.Sdiff5, 2)
	s.M외국계_매도_합계수량 = lib.F2정수64_단순형(g.Fwdvl)
	s.M외국계_매도_직전대비 = lib.F2정수64_단순형(g.Ftradmdcha)
	s.M외국계_매도_비율 = lib.F2실수_소숫점_추가(g.Ftradmddiff, 2)
	s.M외국계_매수_합계수량 = lib.F2정수64_단순형(g.Fwsvl)
	s.M외국계_매수_직전대비 = lib.F2정수64_단순형(g.Ftradmscha)
	s.M외국계_매수_비율 = lib.F2실수_소숫점_추가(g.Ftradmsdiff, 2)
	s.M회전율 = lib.F2실수_소숫점_추가(g.Vol, 2)
	s.M누적거래대금 = lib.F2정수64_단순형(g.Value)
	s.M전일동시간거래량 = lib.F2정수64_단순형(g.Jvolume)
	s.M연중_최고가 = lib.F2정수64_단순형(g.Highyear)
	s.M연중_최고가_일자 = lib.F2포맷된_시각_단순형("20060102", g.Highyeardate)
	s.M연중_최저가 = lib.F2정수64_단순형(g.Lowyear)
	s.M연중_최저가_일자 = lib.F2포맷된_시각_단순형("20060102", g.Lowyeardate)
	s.M목표가 = lib.F2정수64_단순형(g.Target)
	s.M자본금 = lib.F2정수64_단순형(g.Capital)
	s.M유동주식수 = lib.F2정수64_단순형(g.Abscnt)
	s.M액면가 = lib.F2정수64_단순형(g.Parprice)
	s.M결산월 = uint8(lib.F2정수64_단순형_공백은_0(g.Gsmm))
	s.M대용가 = lib.F2정수64_단순형(g.Subprice)
	s.M시가총액_억 = lib.F2정수64_단순형(g.Total)
	s.M상장일 = lib.F2포맷된_시각_단순형("20060102", g.Listdate)
	s.M전분기명 = lib.F2문자열_EUC_KR_공백제거(g.Name)
	s.M전분기_매출액 = lib.F2정수64_단순형(g.Bfsales)
	s.M전분기_영업이익 = lib.F2정수64_단순형(g.Bfoperatingincome)
	s.M전분기_경상이익 = lib.F2정수64_단순형(g.Bfordinaryincome)
	s.M전분기_순이익 = lib.F2정수64_단순형(g.Bfnetincome)
	s.M전분기EPS = lib.F2실수_소숫점_추가(g.Bfeps, 2)
	s.M전전분기명 = lib.F2문자열_EUC_KR_공백제거(g.Name2)
	s.M전전분기_매출액 = lib.F2정수64_단순형(g.Bfsales2)
	s.M전전분기_영업이익 = lib.F2정수64_단순형(g.Bfoperatingincome2)
	s.M전전분기_경상이익 = lib.F2정수64_단순형(g.Bfordinaryincome2)
	s.M전전분기_순이익 = lib.F2정수64_단순형(g.Bfnetincome2)
	s.M전전분기EPS = lib.F2실수_소숫점_추가(g.Bfeps2, 2)
	s.M전년대비매출액 = lib.F2실수_소숫점_추가(g.Salert, 2)
	s.M전년대비영업이익 = lib.F2실수_소숫점_추가(g.Opert, 2)
	s.M전년대비경상이익 = lib.F2실수_소숫점_추가(g.Ordrt, 2)
	s.M전년대비순이익 = lib.F2실수_소숫점_추가(g.Netrt, 2)
	s.M전년대비EPS = lib.F2실수_소숫점_추가(g.Epsrt, 2)
	s.M락구분 = lib.F2문자열_EUC_KR(g.Info1)
	s.M관리_급등구분 = lib.F2문자열_EUC_KR(g.Info2)
	s.M정지_연장구분 = lib.F2문자열_EUC_KR(g.Info3)
	s.M투자_불성실구분 = lib.F2문자열_EUC_KR(g.Info4)
	s.M시장구분 = f2시장구분(g.Janginfo)
	s.T_PER = lib.F2실수_소숫점_추가(g.T_per, 2)
	s.M통화ISO코드 = lib.F2문자열_공백제거(g.Tonghwa)
	s.M총매도대금_모음 = make([]int64, 5)
	s.M총매수대금_모음 = make([]int64, 5)
	s.M총매도대금_모음[0] = lib.F2정수64_단순형_공백은_0(g.Dval1)
	s.M총매수대금_모음[0] = lib.F2정수64_단순형_공백은_0(g.Sval1)
	s.M총매도대금_모음[1] = lib.F2정수64_단순형_공백은_0(g.Dval2)
	s.M총매수대금_모음[1] = lib.F2정수64_단순형_공백은_0(g.Sval2)
	s.M총매도대금_모음[2] = lib.F2정수64_단순형_공백은_0(g.Dval3)
	s.M총매수대금_모음[2] = lib.F2정수64_단순형_공백은_0(g.Sval3)
	s.M총매도대금_모음[3] = lib.F2정수64_단순형_공백은_0(g.Dval4)
	s.M총매수대금_모음[3] = lib.F2정수64_단순형_공백은_0(g.Sval4)
	s.M총매도대금_모음[4] = lib.F2정수64_단순형_공백은_0(g.Dval5)
	s.M총매수대금_모음[4] = lib.F2정수64_단순형_공백은_0(g.Sval5)
	s.M총매도평단가_모음 = make([]int64, 5)
	s.M총매수평단가_모음 = make([]int64, 5)
	s.M총매도평단가_모음[0] = lib.F2정수64_단순형_공백은_0(g.Davg1)
	s.M총매수평단가_모음[0] = lib.F2정수64_단순형_공백은_0(g.Savg1)
	s.M총매도평단가_모음[1] = lib.F2정수64_단순형_공백은_0(g.Davg2)
	s.M총매수평단가_모음[1] = lib.F2정수64_단순형_공백은_0(g.Savg2)
	s.M총매도평단가_모음[2] = lib.F2정수64_단순형_공백은_0(g.Davg3)
	s.M총매수평단가_모음[2] = lib.F2정수64_단순형_공백은_0(g.Savg3)
	s.M총매도평단가_모음[3] = lib.F2정수64_단순형_공백은_0(g.Davg4)
	s.M총매수평단가_모음[3] = lib.F2정수64_단순형_공백은_0(g.Savg4)
	s.M총매도평단가_모음[4] = lib.F2정수64_단순형_공백은_0(g.Davg5)
	s.M총매수평단가_모음[4] = lib.F2정수64_단순형_공백은_0(g.Savg5)
	s.M외국계매도대금 = lib.F2정수64_단순형(g.Ftradmdval)
	s.M외국계매수대금 = lib.F2정수64_단순형_공백은_0(g.Ftradmsval)
	s.M외국계매도평단가 = lib.F2정수64_단순형_공백은_0(g.Ftradmdavg)
	s.M외국계매도평단가 = lib.F2정수64_단순형_공백은_0(g.Ftradmdavg)
	s.M외국계매수평단가 = lib.F2정수64_단순형_공백은_0(g.Ftradmsavg)
	s.M투자주의환기 = lib.F2문자열_EUC_KR_공백제거(g.Info5)
	s.M기업인수목적회사여부 = lib.F2참거짓(g.Spac_gubun, "N", false)
	s.M발행가격 = lib.F2정수64_단순형(g.Issueprice)
	s.M배분적용구분코드 = lib.F2문자열_EUC_KR(g.Alloc_gubun)
	s.M배분적용구분 = lib.F2문자열_EUC_KR(g.Alloc_text)
	s.M단기과열_VI발동 = lib.F2문자열_EUC_KR(g.Shterm_text)
	s.M정적VI상한가 = lib.F2정수64_단순형(g.Svi_uplmtprice)
	s.M정적VI하한가 = lib.F2정수64_단순형(g.Svi_dnlmtprice)
	s.M저유동성종목여부 = lib.F2참거짓(g.Low_lqdt_gu, 1, true)
	s.M이상급등종목여부 = lib.F2참거짓(g.Abnormal_rise_gu, 1, true)

	대차불가표시_문자열 := lib.F2문자열_EUC_KR_공백제거(g.Lend_text)
	switch 대차불가표시_문자열 {
	case "":
		s.M대차불가여부 = false
	case "대차불가":
		s.M대차불가여부 = true
	default:
		panic(lib.New에러("%v '대차불가표시_문자열' 예상하지 못한 값 : '%v'", s.M종목코드, 대차불가표시_문자열))
	}

	return s, nil
}

func NewT1305InBlock(질의값 *S질의값_현물_기간별_조회) (g *T1305InBlock) {
	g = new(T1305InBlock)
	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lib.F바이트_복사_문자열(g.Dwmcode[:], lib.F2문자열(uint8(질의값.M일주월_구분)))
	lib.F바이트_복사_문자열(g.Date[:], 질의값.M연속키)
	lib.F바이트_복사_문자열(g.Idx[:], "    ") // 정수형인데, 사용안함(Space)으로 표시됨.
	lib.F바이트_복사_정수(g.Cnt[:], 질의값.M수량)

	if lib.F2문자열_공백제거(질의값.M연속키) == "" {
		lib.F바이트_복사_문자열(g.Date[:], "       ")
	}

	return g
}

func NewT1310InBlock(질의값 *S질의값_현물_전일당일_분틱_조회) (g *T1310InBlock) {
	g = new(T1310InBlock)
	lib.F바이트_복사_문자열(g.Daygb[:], strconv.Itoa(int(질의값.M당일전일구분)))
	lib.F바이트_복사_문자열(g.Timegb[:], strconv.Itoa(int(질의값.M분틱구분)))
	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lib.F바이트_복사_문자열(g.Endtime[:], strings.Replace(질의값.M종료시각.Format("15:04"), ":", "", -1))
	lib.F바이트_복사_문자열(g.Time[:], 질의값.M연속키)

	if lib.F2문자열_공백제거(질의값.M연속키) == "" {
		lib.F바이트_복사_문자열(g.Time[:], "          ")
	}

	return g
}

func New현물_당일전일분틱조회_응답_헤더(b []byte) (값 *S현물_전일당일분틱조회_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT1310OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T1310OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S현물_전일당일분틱조회_응답_헤더)
	값.M연속키 = lib.F2문자열(g.Time)

	return 값, nil
}

func New현물_당일전일분틱조회_응답_반복값_모음(b []byte) (값 *S현물_전일당일분틱조회_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT1310OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT1310OutBlock1
	g_모음 := make([]*T1310OutBlock1, 수량, 수량)

	값 = new(S현물_전일당일분틱조회_응답_반복값_모음)
	값.M배열 = make([]*S현물_전일당일분틱조회_응답_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T1310OutBlock1)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		s := new(S현물_전일당일분틱조회_응답_반복값)
		s.M시각 = lib.F2금일_시각_단순형("150405", g.Chetime[:6])
		s.M현재가 = lib.F2정수64_단순형(g.Price)
		s.M전일대비구분 = T전일대비_구분(lib.F2정수64_단순형(g.Sign))
		s.M전일대비등락폭 = s.M전일대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.Change))
		s.M전일대비등락율 = s.M전일대비구분.G부호보정_실수64(lib.F2실수_단순형(g.Diff))
		s.M체결수량 = lib.F2정수64_단순형(g.Cvolume)
		s.M체결강도 = lib.F2실수_단순형(g.Chdegree)
		s.M거래량 = lib.F2정수64_단순형(g.Volume)
		s.M매도체결수량 = lib.F2정수64_단순형(g.Mdvolume)
		s.M매도체결건수 = lib.F2정수64_단순형(g.Mdchecnt)
		s.M매수체결수량 = lib.F2정수64_단순형(g.Msvolume)
		s.M매수체결건수 = lib.F2정수64_단순형(g.Mschecnt)
		s.M순체결량 = lib.F2정수64_단순형(g.Revolume)
		s.M순체결건수 = lib.F2정수64_단순형(g.Rechecnt)

		값.M배열[i] = s
	}

	return 값, nil
}

func NewT1902InBlock(질의값 *S질의값_단일종목_연속키) (g *T1902InBlock) {
	g = new(T1902InBlock)
	lib.F바이트_복사_문자열(g.ShCode[:], 질의값.M종목코드)

	if lib.F2문자열_공백제거(질의값.M연속키) == "" {
		lib.F바이트_배열_공백문자열_채움(g.Time)
	} else {
		lib.F바이트_복사_문자열(g.Time[:], 질의값.M연속키)
	}

	return g
}

func NewETF시간별_추이_응답_헤더(b []byte) (s *S_ETF시간별_추이_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT1902OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T1902OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	s = new(S_ETF시간별_추이_응답_헤더)
	s.M연속키 = lib.F2문자열_공백제거(g.Time)
	s.M종목명 = lib.F2문자열_EUC_KR_공백제거(g.HName)
	s.M업종지수명 = lib.F2문자열_EUC_KR_공백제거(g.UpName)

	return s, nil
}

func NewETF시간별_추이_응답_반복값_모음(b []byte) (값 *S_ETF시간별_추이_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT1902OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT1902OutBlock1
	g_모음 := make([]*T1902OutBlock1, 수량, 수량)

	값 = new(S_ETF시간별_추이_응답_반복값_모음)
	값.M배열 = make([]*S_ETF시간별_추이_응답_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T1902OutBlock1)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		s := new(S_ETF시간별_추이_응답_반복값)

		if lib.F2문자열_EUC_KR(g.Time) == "장:마:감" {
			s.M시각 = lib.F2일자별_시각_단순형(당일.G값(), "15:04:05", g_모음[i+1].Time).Add(lib.P10초)
		} else {
			s.M시각 = lib.F2일자별_시각_단순형(당일.G값(), "15:04:05", g.Time)
		}

		s.M현재가 = lib.F2정수64_단순형(g.Price)
		s.M전일대비구분 = T전일대비_구분(lib.F2정수_단순형(g.Sign))
		s.M전일대비등락폭 = s.M전일대비구분.G부호보정_정수64(lib.F2정수64_단순형(g.Change))
		s.M누적_거래량 = lib.F2실수_단순형(g.Volume)
		s.M현재가_NAV_차이 = lib.F2실수_단순형(g.NavDiff)
		s.NAV = lib.F2실수_단순형(g.Nav)
		s.NAV전일대비등락폭 = lib.F2실수_단순형(g.NavChange)
		s.M추적오차 = lib.F2실수_단순형(g.Crate)
		s.M괴리율 = lib.F2실수_단순형(g.Grate)
		s.M지수 = lib.F2실수_단순형(g.Jisu)
		s.M지수_전일대비등락폭 = lib.F2실수_단순형(g.JiChange)
		s.M지수_전일대비등락율 = lib.F2실수_단순형(g.JiRate)

		if uint8(g.X_jichange) == 160 && s.M지수_전일대비등락폭 > 0 {
			s.M지수_전일대비등락폭 = -1 * s.M지수_전일대비등락폭
		}

		값.M배열[i] = s
	}

	return 값, nil
}

func NewT8411InBlock(질의값 *S질의값_현물_차트_틱) (g *T8411InBlock) {
	g = new(T8411InBlock)
	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lib.F바이트_복사_정수(g.Ncnt[:], 질의값.M단위)
	lib.F바이트_복사_정수(g.Qrycnt[:], 질의값.M요청건수)
	lib.F바이트_복사_정수(g.Nday[:], 질의값.M조회영업일수)
	lib.F바이트_복사_문자열(g.Sdate[:], 질의값.M시작일자.Format("20060102"))
	lib.F바이트_복사_문자열(g.Edate[:], 질의값.M종료일자.Format("20060102"))
	lib.F바이트_복사_문자열(g.Cts_date[:], 질의값.M연속일자)
	lib.F바이트_복사_문자열(g.Cts_time[:], 질의값.M연속시간)
	lib.F바이트_복사_문자열(g.Comp_yn[:], lib.F조건부_문자열(질의값.M압축여부, "Y", "N"))

	return g
}

func New현물_차트_틱_응답_헤더(b []byte) (값 *S현물_차트_틱_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT8411OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T8411OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S현물_차트_틱_응답_헤더)
	값.M종목코드 = lib.F2문자열(g.Shcode)
	값.M전일시가 = lib.F2정수64_단순형(g.Jisiga)
	값.M전일고가 = lib.F2정수64_단순형(g.Jihigh)
	값.M전일저가 = lib.F2정수64_단순형(g.Jilow)
	값.M전일종가 = lib.F2정수64_단순형(g.Jiclose)
	값.M전일거래량 = lib.F2정수64_단순형(g.Jivolume)
	값.M당일시가 = lib.F2정수64_단순형(g.Disiga)
	값.M당일고가 = lib.F2정수64_단순형(g.Dihigh)
	값.M당일저가 = lib.F2정수64_단순형(g.Dilow)
	값.M당일종가 = lib.F2정수64_단순형(g.Diclose)
	값.M상한가 = lib.F2정수64_단순형(g.Highend)
	값.M하한가 = lib.F2정수64_단순형(g.Lowend)
	값.M연속일자 = lib.F2문자열(g.Cts_date)
	값.M연속시간 = lib.F2문자열(g.Cts_time)
	값.M장시작시간 = lib.F2일자별_시각_단순형(당일.G값(), "150405", g.S_time)
	값.M장종료시간 = lib.F2일자별_시각_단순형(당일.G값(), "150405", g.E_time)
	값.M동시호가처리시간 = lib.F2정수_단순형(g.Dshmin)
	값.M수량 = lib.F2정수64_단순형(g.Rec_count)

	return 값, nil
}

func New현물_차트_틱_응답_반복값_모음(b []byte) (값 *S현물_차트_틱_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT8411OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT8411OutBlock1
	g_모음 := make([]*T8411OutBlock1, 수량, 수량)

	값 = new(S현물_차트_틱_응답_반복값_모음)
	값.M배열 = make([]*S현물_차트_틱_응답_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T8411OutBlock1)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		날짜_문자열 := lib.F2문자열_공백제거(g.Date)
		시각_문자열 := lib.F2문자열_공백제거(g.Time[:6])

		s := new(S현물_차트_틱_응답_반복값)
		s.M일자_시각 = lib.F2포맷된_시각_단순형("20060102 150405", 날짜_문자열+" "+시각_문자열)
		s.M시가 = lib.F2정수64_단순형(g.Open)
		s.M고가 = lib.F2정수64_단순형(g.High)
		s.M저가 = lib.F2정수64_단순형(g.Low)
		s.M종가 = lib.F2정수64_단순형(g.Close)
		s.M거래량 = lib.F2정수64_단순형(g.Vol)
		s.M수정구분 = lib.F2정수64_단순형_공백은_0(g.Jongchk)
		s.M수정비율 = lib.F2실수_단순형_공백은_0(g.Rate)
		s.M수정주가반영항목 = lib.F2정수64_단순형_공백은_0(g.Pricechk)

		값.M배열[i] = s
	}

	return 값, nil
}

func NewT8412InBlock(질의값 *S질의값_현물_차트_분) (g *T8412InBlock) {
	g = new(T8412InBlock)

	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lib.F바이트_복사_정수(g.Ncnt[:], 질의값.M단위)
	lib.F바이트_복사_정수(g.Qrycnt[:], 질의값.M요청건수)
	lib.F바이트_복사_정수(g.Nday[:], 질의값.M조회영업일수)
	lib.F바이트_복사_문자열(g.Sdate[:], 질의값.M시작일자.Format("20060102"))
	lib.F바이트_복사_문자열(g.Edate[:], 질의값.M종료일자.Format("20060102"))
	lib.F바이트_복사_문자열(g.Cts_date[:], 질의값.M연속일자)
	lib.F바이트_복사_문자열(g.Cts_time[:], 질의값.M연속시간)
	lib.F바이트_복사_문자열(g.Comp_yn[:], lib.F조건부_문자열(질의값.M압축여부, "Y", "N"))

	return g
}

func New현물_차트_분_응답_헤더(b []byte) (값 *S현물_차트_분_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT8412OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T8412OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S현물_차트_분_응답_헤더)
	값.M종목코드 = lib.F2문자열(g.Shcode)
	값.M전일시가 = lib.F2정수64_단순형(g.Jisiga)
	값.M전일고가 = lib.F2정수64_단순형(g.Jihigh)
	값.M전일저가 = lib.F2정수64_단순형(g.Jilow)
	값.M전일종가 = lib.F2정수64_단순형(g.Jiclose)
	값.M전일거래량 = lib.F2정수64_단순형(g.Jivolume)
	값.M당일시가 = lib.F2정수64_단순형(g.Disiga)
	값.M당일고가 = lib.F2정수64_단순형(g.Dihigh)
	값.M당일저가 = lib.F2정수64_단순형(g.Dilow)
	값.M당일종가 = lib.F2정수64_단순형(g.Diclose)
	값.M상한가 = lib.F2정수64_단순형(g.Highend)
	값.M하한가 = lib.F2정수64_단순형(g.Lowend)
	값.M연속일자 = lib.F2문자열(g.Cts_date)
	값.M연속시간 = lib.F2문자열(g.Cts_time)
	값.M장시작시간 = lib.F2일자별_시각_단순형(당일.G값(), "150405", g.S_time)
	값.M장종료시간 = lib.F2일자별_시각_단순형(당일.G값(), "150405", g.E_time)
	값.M동시호가처리시간 = lib.F2정수_단순형(g.Dshmin)
	값.M수량 = lib.F2정수64_단순형(g.Rec_count)

	return 값, nil
}

func New현물_차트_분_응답_반복값_모음(b []byte) (값 *S현물_차트_분_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT8412OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT8412OutBlock1
	g_모음 := make([]*T8412OutBlock1, 수량, 수량)

	값 = new(S현물_차트_분_응답_반복값_모음)
	값.M배열 = make([]*S현물_차트_분_응답_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T8412OutBlock1)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		날짜_문자열 := lib.F2문자열_공백제거(g.Date)
		시각_문자열 := lib.F2문자열_공백제거(g.Time[:6])

		s := new(S현물_차트_분_응답_반복값)
		s.M일자_시각 = lib.F2포맷된_시각_단순형("20060102 150405", 날짜_문자열+" "+시각_문자열)
		s.M시가 = lib.F2정수64_단순형(g.Open)
		s.M고가 = lib.F2정수64_단순형(g.High)
		s.M저가 = lib.F2정수64_단순형(g.Low)
		s.M종가 = lib.F2정수64_단순형(g.Close)
		s.M거래량 = lib.F2정수64_단순형(g.Vol)
		s.M거래대금_백만 = lib.F2정수64_단순형(g.Value)
		s.M수정구분 = lib.F2정수64_단순형_공백은_0(g.Jongchk)
		s.M수정비율 = lib.F2실수_단순형_공백은_0(g.Rate)
		s.M종가등락구분 = T전일대비_구분(lib.F2정수_단순형(g.Sign))

		값.M배열[i] = s
	}

	return 값, nil
}

func NewT8413InBlock(질의값 *S질의값_현물_차트_일주월) (g *T8413InBlock) {
	g = new(T8413InBlock)

	lib.F바이트_복사_문자열(g.Shcode[:], 질의값.M종목코드)
	lib.F바이트_복사_문자열(g.Gubun[:], strconv.Itoa(int(uint8(질의값.M주기구분)+1)))
	lib.F바이트_복사_정수(g.Qrycnt[:], 질의값.M요청건수)
	lib.F바이트_복사_문자열(g.Sdate[:], 질의값.M시작일자.Format("20060102"))
	lib.F바이트_복사_문자열(g.Edate[:], 질의값.M종료일자.Format("20060102"))
	lib.F바이트_복사_문자열(g.Cts_date[:], 질의값.M연속일자)
	lib.F바이트_복사_문자열(g.Comp_yn[:], lib.F조건부_문자열(질의값.M압축여부, "Y", "N"))

	return g
}

func New현물_차트_일주월_응답_헤더(b []byte) (값 *S현물_차트_일주월_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT8413OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T8413OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S현물_차트_일주월_응답_헤더)
	값.M종목코드 = lib.F2문자열(g.Shcode)
	값.M전일시가 = lib.F2정수64_단순형(g.Jisiga)
	값.M전일고가 = lib.F2정수64_단순형(g.Jihigh)
	값.M전일저가 = lib.F2정수64_단순형(g.Jilow)
	값.M전일종가 = lib.F2정수64_단순형(g.Jiclose)
	값.M전일거래량 = lib.F2정수64_단순형(g.Jivolume)
	값.M당일시가 = lib.F2정수64_단순형(g.Disiga)
	값.M당일고가 = lib.F2정수64_단순형(g.Dihigh)
	값.M당일저가 = lib.F2정수64_단순형(g.Dilow)
	값.M당일종가 = lib.F2정수64_단순형(g.Diclose)
	값.M상한가 = lib.F2정수64_단순형(g.Highend)
	값.M하한가 = lib.F2정수64_단순형(g.Lowend)
	값.M연속일자 = lib.F2문자열(g.Cts_date)
	값.M장시작시간 = lib.F2일자별_시각_단순형(당일.G값(), "150405", g.S_time)
	값.M장종료시간 = lib.F2일자별_시각_단순형(당일.G값(), "150405", g.E_time)
	값.M동시호가처리시간 = lib.F2정수_단순형(g.Dshmin)
	값.M수량 = lib.F2정수64_단순형(g.Rec_count)

	return 값, nil
}

func New현물_차트_일주월_응답_반복값_모음(b []byte) (값 *S현물_차트_일주월_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT8413OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT8413OutBlock1
	g_모음 := make([]*T8413OutBlock1, 수량, 수량)

	값 = new(S현물_차트_일주월_응답_반복값_모음)
	값.M배열 = make([]*S현물_차트_일주월_응답_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T8413OutBlock1)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		s := new(S현물_차트_일주월_응답_반복값)
		s.M일자 = lib.F2포맷된_시각_단순형("20060102", lib.F2문자열_공백제거(g.Date))
		s.M시가 = lib.F2정수64_단순형(g.Open)
		s.M고가 = lib.F2정수64_단순형(g.High)
		s.M저가 = lib.F2정수64_단순형(g.Low)
		s.M종가 = lib.F2정수64_단순형(g.Close)
		s.M거래량 = lib.F2정수64_단순형(g.Vol)
		s.M거래대금_백만 = lib.F2정수64_단순형(g.Value)
		s.M수정구분 = lib.F2정수64_단순형_공백은_0(g.Jongchk)
		s.M수정비율 = lib.F2실수_단순형_공백은_0(g.Rate)
		s.M수정주가반영항목 = lib.F2정수64_단순형_공백은_0(g.Pricechk)
		s.M수정비율반영거래대금 = lib.F2정수64_단순형_공백은_0(g.Ratevalue)
		s.M종가등락구분 = T전일대비_구분(lib.F2정수_단순형(g.Sign))

		값.M배열[i] = s
	}

	return 값, nil
}

func NewT8428InBlock(질의값 *S질의값_증시주변자금추이) (g *T8428InBlock) {
	시장구분_문자열 := ""
	switch 질의값.M시장구분 {
	case lib.P시장구분_코스피:
		시장구분_문자열 = "001"
	case lib.P시장구분_코스닥:
		시장구분_문자열 = "301"
	default:
		panic(lib.New에러("예상하지 못한 시장구분 값 : '%v'", 질의값.M시장구분))
	}

	g = new(T8428InBlock)
	lib.F바이트_복사_문자열(g.KeyDate[:], 질의값.M연속키)
	lib.F바이트_복사_문자열(g.Upcode[:], 시장구분_문자열)
	lib.F바이트_복사_정수(g.Cnt[:], 질의값.M수량)

	if lib.F2문자열_공백제거(질의값.M연속키) == "" {
		lib.F바이트_복사_문자열(g.KeyDate[:], "        ")
	}

	return g
}

func New증시주변자금추이_응답_헤더(b []byte) (값 *S증시_주변자금추이_응답_헤더, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	lib.F조건부_패닉(len(b) != SizeT8413OutBlock,
		"예상하지 못한 길이 : '%v", len(b))

	g := new(T8428OutBlock)
	lib.F확인(binary.Read(bytes.NewBuffer(b), binary.BigEndian, g))

	값 = new(S증시_주변자금추이_응답_헤더)
	값.M연속키 = lib.F2문자열(g.Date)
	값.M인덱스 = lib.F2정수64_단순형(g.Idx)

	return 값, nil
}

func New증시주변자금추이_응답_반복값_모음(b []byte) (값 *S증시_주변자금추이_응답_반복값_모음, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값 = nil }}.S실행()

	나머지 := len(b) % SizeT8428OutBlock1
	lib.F조건부_패닉(나머지 != 0, "예상하지 못한 길이. '%v' '%v'", len(b), 나머지)

	버퍼 := bytes.NewBuffer(b)
	수량 := len(b) / SizeT8428OutBlock1
	g_모음 := make([]*T8428OutBlock1, 수량, 수량)

	값 = new(S증시_주변자금추이_응답_반복값_모음)
	값.M배열 = make([]*S증시_주변자금추이_응답_반복값, 수량, 수량)

	for i, g := range g_모음 {
		g = new(T8428OutBlock1)
		lib.F확인(binary.Read(버퍼, binary.BigEndian, g))

		s := new(S증시_주변자금추이_응답_반복값)
		s.M일자 = lib.F2포맷된_시각_단순형("20060102", lib.F2문자열(g.Date))
		s.M지수 = lib.F2실수_단순형(g.Jisu)
		s.M전일대비_구분 = T전일대비_구분(lib.F2정수64_단순형(g.Sign))
		s.M전일대비_등락폭 = s.M전일대비_구분.G부호보정_실수64(lib.F2실수_단순형(g.Change))
		s.M전일대비_등락율 = s.M전일대비_구분.G부호보정_실수64(lib.F2실수_단순형(g.Diff))
		s.M거래량 = lib.F2정수64_단순형(g.Volume)
		s.M고객예탁금_억 = lib.F2정수64_단순형(g.Custmoney)
		s.M예탁증감_억 = lib.F2정수64_단순형(g.Yecha)

		if strings.Contains(strings.ToLower(lib.F2문자열(g.Vol)), "inf") {
			s.M회전율 = math.Inf(1)
		} else {
			s.M회전율 = lib.F2실수_단순형(g.Vol)
		}

		s.M미수금_억 = lib.F2정수64_단순형(g.Outmoney)
		s.M신용잔고_억 = lib.F2정수64_단순형(g.Trjango)
		s.M선물예수금_억 = lib.F2정수64_단순형(g.Futymoney)
		s.M주식형_억 = lib.F2정수64_단순형(g.Stkmoney)
		s.M혼합형_주식_억 = lib.F2정수64_단순형(g.Mstkmoney)
		s.M혼합형_채권_억 = lib.F2정수64_단순형(g.Mbndmoney)
		s.M채권형_억 = lib.F2정수64_단순형(g.Bndmoney)
		s.MMF_억 = lib.F2정수64_단순형(g.Mmfmoney)

		값.M배열[i] = s
	}

	return 값, nil
}

