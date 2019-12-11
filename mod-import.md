import text package in go.mod
import 
//convert html from gbk=>utf-8
	transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
	
import "golang.org/x/net/html"	
//determine charset of html
func determineEncoding(r io.Reader) encoding.Encoding{
	bytes,err := bufio.NewReader(r).Peek(1024)
	if err!=nil {
		panic(err)
	}
	e,_,_ :=charset.DetermineEncoding(bytes,"")
	return e
}


	