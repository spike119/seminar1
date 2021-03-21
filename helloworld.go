/*
패키지는 두 종류가 있습니다.
바로 main 패키지와 그 외 패키지입니다.
main 패키지를 제외한 모든 패키지는 개발자가 직접 작성하든 Go언어에서 제공하는 패키지든 똑같이 취급됩니다.
반면에, main 패키지는 컴파일러에 의해 특별히 인식됩니다.
컴파일러는 main 패키지를 공유 패키지(라이브러리)가 아닌 실행(executable) 프로그램으로 만듭니다.
자연스럽게 main 패키지 안에 main() 함수는 프로그램의 시작점(Entry point)이자 종료점이 됩니다.
따라서 패키지를 공유 라이브러리로 만들 때에는 main 패키지나 main 함수를 사용하면 안됩니다.
*/

package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
