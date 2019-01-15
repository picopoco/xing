package xing

import (
	"github.com/ghts/lib"
	"reflect"
	"strings"
	"testing"
)

func TestP자료형_문자열(t *testing.T) {
	lib.F테스트_같음(t, P자료형_S현물_주문_응답_실시간_정보, f자료형_문자열(S현물_주문_응답_실시간_정보{}))
	lib.F테스트_같음(t, P자료형_S질의값_정상_주문, f자료형_문자열(S질의값_정상_주문{}))
	lib.F테스트_같음(t, P자료형_S질의값_정정_주문, f자료형_문자열(S질의값_정정_주문{}))
	lib.F테스트_같음(t, P자료형_S질의값_취소_주문, f자료형_문자열(S질의값_취소_주문{}))
	lib.F테스트_같음(t, P자료형_S질의값_현물_전일당일_분틱_조회, f자료형_문자열(S질의값_현물_전일당일_분틱_조회{}))
	lib.F테스트_같음(t, P자료형_S질의값_현물_기간별_조회, f자료형_문자열(S질의값_현물_기간별_조회{}))
	lib.F테스트_같음(t, P자료형_S질의값_단일종목_연속키, f자료형_문자열(S질의값_단일종목_연속키{}))
	lib.F테스트_같음(t, P자료형_S질의값_증시주변자금추이, f자료형_문자열(S질의값_증시주변자금추이{}))

	lib.F테스트_같음(t, P자료형_S콜백_기본형, f자료형_문자열(S콜백_기본형{}))
	lib.F테스트_같음(t, P자료형_S콜백_정수값, f자료형_문자열(S콜백_정수값{}))
	lib.F테스트_같음(t, P자료형_S콜백_문자열, f자료형_문자열(S콜백_문자열{}))
	lib.F테스트_같음(t, P자료형_S콜백_TR데이터, f자료형_문자열(S콜백_TR데이터{}))
	lib.F테스트_같음(t, P자료형_S콜백_메시지_및_에러, f자료형_문자열(S콜백_메시지_및_에러{}))

	lib.F테스트_같음(t, P자료형_S현물_정상_주문_응답, f자료형_문자열(S현물_정상_주문_응답{}))
	lib.F테스트_같음(t, P자료형_S현물_정상_주문_응답1, f자료형_문자열(S현물_정상_주문_응답1{}))
	lib.F테스트_같음(t, P자료형_S현물_정상_주문_응답2, f자료형_문자열(S현물_정상_주문_응답2{}))
	lib.F테스트_같음(t, P자료형_S현물_정정_주문_응답, f자료형_문자열(S현물_정정_주문_응답{}))
	lib.F테스트_같음(t, P자료형_S현물_정정_주문_응답1, f자료형_문자열(S현물_정정_주문_응답1{}))
	lib.F테스트_같음(t, P자료형_S현물_정정_주문_응답2, f자료형_문자열(S현물_정정_주문_응답2{}))
	lib.F테스트_같음(t, P자료형_S현물_취소_주문_응답, f자료형_문자열(S현물_취소_주문_응답{}))
	lib.F테스트_같음(t, P자료형_S현물_취소_주문_응답1, f자료형_문자열(S현물_취소_주문_응답1{}))
	lib.F테스트_같음(t, P자료형_S현물_취소_주문_응답2, f자료형_문자열(S현물_취소_주문_응답2{}))

	lib.F테스트_같음(t, P자료형_S현물_호가조회_응답, f자료형_문자열(S현물_호가조회_응답{}))
	lib.F테스트_같음(t, P자료형_S현물_시세조회_응답, f자료형_문자열(S현물_시세조회_응답{}))
	lib.F테스트_같음(t, P자료형_S현물_시간대별_체결_응답, f자료형_문자열(S현물_시간대별_체결_응답{}))
	lib.F테스트_같음(t, P자료형_S현물_시간대별_체결_응답_헤더, f자료형_문자열(S현물_시간대별_체결_응답_헤더{}))
	lib.F테스트_같음(t, P자료형_S현물_시간대별_체결_응답_반복값, f자료형_문자열(S현물_시간대별_체결_응답_반복값{}))
	lib.F테스트_같음(t, P자료형_S현물_시간대별_체결_응답_반복값_모음, f자료형_문자열(S현물_시간대별_체결_응답_반복값_모음{}))
	lib.F테스트_같음(t, P자료형_S현물_기간별_조회_응답, f자료형_문자열(S현물_기간별_조회_응답{}))
	lib.F테스트_같음(t, P자료형_S현물_기간별_조회_응답_헤더, f자료형_문자열(S현물_기간별_조회_응답_헤더{}))
	lib.F테스트_같음(t, P자료형_S현물_기간별_조회_응답_반복값, f자료형_문자열(S현물_기간별_조회_응답_반복값{}))
	lib.F테스트_같음(t, P자료형_S현물_기간별_조회_응답_반복값_모음, f자료형_문자열(S현물_기간별_조회_응답_반복값_모음{}))
	lib.F테스트_같음(t, P자료형_S현물_전일당일분틱조회_응답, f자료형_문자열(S현물_전일당일분틱조회_응답{}))
	lib.F테스트_같음(t, P자료형_S현물_전일당일분틱조회_응답_헤더, f자료형_문자열(S현물_전일당일분틱조회_응답_헤더{}))
	lib.F테스트_같음(t, P자료형_S현물_전일당일분틱조회_응답_반복값, f자료형_문자열(S현물_전일당일분틱조회_응답_반복값{}))
	lib.F테스트_같음(t, P자료형_S현물_전일당일분틱조회_응답_반복값_모음, f자료형_문자열(S현물_전일당일분틱조회_응답_반복값_모음{}))
	lib.F테스트_같음(t, P자료형_S_ETF_현재가_조회_응답, f자료형_문자열(S_ETF_현재가_조회_응답{}))
	lib.F테스트_같음(t, P자료형_S_ETF시간별_추이_응답, f자료형_문자열(S_ETF시간별_추이_응답{}))
	lib.F테스트_같음(t, P자료형_S_ETF시간별_추이_응답_헤더, f자료형_문자열(S_ETF시간별_추이_응답_헤더{}))
	lib.F테스트_같음(t, P자료형_S_ETF시간별_추이_응답_반복값, f자료형_문자열(S_ETF시간별_추이_응답_반복값{}))
	lib.F테스트_같음(t, P자료형_S_ETF시간별_추이_응답_반복값_모음, f자료형_문자열(S_ETF시간별_추이_응답_반복값_모음{}))
	lib.F테스트_같음(t, P자료형_S현물_차트_틱_응답, f자료형_문자열(S현물_차트_틱_응답{}))
	lib.F테스트_같음(t, P자료형_S현물_차트_틱_응답_헤더, f자료형_문자열(S현물_차트_틱_응답_헤더{}))
	lib.F테스트_같음(t, P자료형_S현물_차트_틱_응답_반복값, f자료형_문자열(S현물_차트_틱_응답_반복값{}))
	lib.F테스트_같음(t, P자료형_S현물_차트_틱_응답_반복값_모음, f자료형_문자열(S현물_차트_틱_응답_반복값_모음{}))
	lib.F테스트_같음(t, P자료형_S증시주변자금추이_응답, f자료형_문자열(S증시주변자금추이_응답{}))
	lib.F테스트_같음(t, P자료형_S증시주변자금추이_응답_헤더, f자료형_문자열(S증시주변자금추이_응답_헤더{}))
	lib.F테스트_같음(t, P자료형_S증시주변자금추이_응답_반복값_모음, f자료형_문자열(S증시주변자금추이_응답_반복값_모음{}))
	lib.F테스트_같음(t, P자료형_S주식종목조회_응답_반복값, f자료형_문자열(S현물_종목조회_응답_반복값{}))
}

func f자료형_문자열(값 interface{}) string {
	자료형 := reflect.TypeOf(값).String()
	시작_인덱스 := strings.Index(자료형, ".") + 1

	return 자료형[시작_인덱스:]
}
