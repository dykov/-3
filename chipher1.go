package main
import(
	"fmt"
	"unicode/utf8"
)

/*
close
function_needs_result_type
*/

type ColumnStruct struct{
    KeyChar rune
    Num int
    SliceOfChar []rune
}

func Round( A , B int ) int {
    if( A % B == 0 ){
        return A / B
    }
    return A / B + 1
}

func Sort( sl []ColumnStruct ) {
    var key rune
    var bufSlice []rune
    var j int
    for i := 1 ; i < len(sl) ; i++ {
        key = sl[i].KeyChar
        bufSlice = sl[i].SliceOfChar
        j = i-1
        for (j >= 0 && sl[j].KeyChar > key){
            sl[j+1].KeyChar = sl[j].KeyChar
            sl[j+1].SliceOfChar = sl[j].SliceOfChar
            j = j-1
        }
        sl[j+1].KeyChar = key
        sl[j+1].SliceOfChar = bufSlice
    }
}

func Print( SliceOfStructs []ColumnStruct ){
    fmt.Println("--------------------")
    for _,st := range SliceOfStructs {
        fmt.Printf("%c : ", st.KeyChar )
        for _,vv := range st.SliceOfChar {
            fmt.Print( string(vv) )
        }
        fmt.Println("")
    }
    fmt.Println("--------------------")
}


func main(){
	var message , key string
    fmt.Print("Введите Ваш ключ: ")
	fmt.Scan( &key )
    fmt.Println("[Введено" , utf8.RuneCountInString(key) , "символов]" ) // кол-во рун-символов

	fmt.Print("Введите Ваше сообщение: ")
	fmt.Scan( &message )
	fmt.Println("[Введено" , utf8.RuneCountInString(message) , "символов]" ) // кол-во рун-символов

    numColumns := utf8.RuneCountInString(key)
    fmt.Println("N столбцов =" , numColumns )

    numLines := Round( int(utf8.RuneCountInString(message)) , numColumns ) + 2
    fmt.Println("N строк = [" , utf8.RuneCountInString(message) , " / " , numColumns , "] + 2 =" , numLines  )

    if( len(message) % len(key) != 0 ){
        for len(message) % len(key) != 0{
            message += "_"
        }
        fmt.Println("Ваше сообщение: ",message)
    }

    /* даем каждому символу номер по алфавиту */
    runeKey := []rune( key ) // переводим строку в слайс рун, чтоб перебрать посимвольно
    runeMessage := []rune( message )
    SliceOfStructs := make( []ColumnStruct, 0 ) // слайс для хранения всех столбцов
    for _,c := range runeKey {
        SliceOfStructs = append( SliceOfStructs , []ColumnStruct{ ColumnStruct{ KeyChar : c } }... )
    }

    for numofchar := 0 ; numofchar < len(runeMessage) ; {
        for numofcolumn := 0 ; numofcolumn < len(runeKey) ; numofcolumn++ {
            for numofline := 0; numofline < numLines-2; numofline++ {
                SliceOfStructs[numofcolumn].SliceOfChar = append( SliceOfStructs[numofcolumn].SliceOfChar , []rune{ runeMessage[numofchar] }... )
                numofchar++
            }
        }
    }

    Print( SliceOfStructs )

    Sort( SliceOfStructs )

    Print( SliceOfStructs )

    for i := 0 ; i < numLines-2 ; i++{
        for _,column := range SliceOfStructs {
            fmt.Printf("%c", column.SliceOfChar[i])
        }
    }



}
