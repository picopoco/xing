# xing

이베스트투자증권의 Xing API를 호출하는 패키지
  - 실제 Xing API 호출은 xing_C32 패키지에 위임, 결과물만 전달받아서 사용.
  
#
이렇게까지 복잡하게 해야했나??
  - 64비트에서는 메모리 제한에서 자유로우므로, 대용량 데이터 처리에 유리하다.
  - 동시성 처리에 강한 Go언어는 64비트에서만 데이터 레이스(Data Race) 감지기가 제공된다.  
  - 그러나, 64비트에서는 32비트 DLL을 직접 호출할 수 없다.
  - 별도의 32비트 전용 모듈(xing_C32)에게 DLL호출을 위임한 후 결과물을 받아보는 것으로 해결.
  - 네트워크 전송 기능은 nanomsg의 Go언어 구현체인 mangos패키지로 간편하게 해결.
  
#
설치 준비물
  - Go언어 : https://golang.org/dl/
  - Rtools 패키지 (C 컴파일러) : https://cran.r-project.org/bin/windows/Rtools/index.html
  - Git 소스코드 관리 시스템 : https://git-scm.com/download/win 

#
설치법

    go get github.com/ghts/lib
    go get github.com/ghts/xing
    go get github.com/ghts/xing_C32
    
   
#   
사용법

    package main

    import "github.com/ghts/xing"

    func main() {
	    xing.F초기화()
	    defer xing.F리소스_정리()

        (... 이하 Xing API 호출 ...)
    }

TR 호출 예제 : https://github.com/ghts/xing 에서 'tr_*_test.go' 파일 참조.

문서 : https://godoc.org/github.com/ghts/xing
 

#
참고 링크.
  - xing_C32 패키지 : https://github.com/ghts/xing_C32
  - 데이터 레이스(data race) 감지기 : https://golang.org/doc/articles/race_detector.html
  - 이베스트 투자증권 : https://www.ebestsec.co.kr
  - nanomsg : https://nanomsg.org
  - mangos : https://github.com/nanomsg/mangos
 