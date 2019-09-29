package uid

var sf = &Snowflake{
	hasInit: false,
}

//取得UID
func GetUid() (int64, error) {
	return sf.NextUid()
}

//开启生成UID服务，需要在程序开始时调用
func StartUidServe() error {
	return sf.StartServe()
}
