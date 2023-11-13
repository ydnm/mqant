package log

//go:generate optiongen --option_with_struct_name=false
func OptionsOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"Debug":       false,
		"ProcessID":   "",
		"LogDir":      "",
		"LogFileName": func(logdir, prefix, processID, suffix string) string { return "" },
		"BiDir":       "",
		"BIFileName":  func(logdir, prefix, processID, suffix string) string { return "" },
		"BiSetting":   map[string]interface{}{},
		"LogSetting":  map[string]interface{}{},
	}
}
