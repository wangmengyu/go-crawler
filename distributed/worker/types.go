package worker

/**
  序列化的解析器
*/
type SerializedParser struct {
	Name string      //方法名
	Args interface{} // 参数
}

//{"ParseCityList",nil}
