package request

import "encoding/json"

func Verify(data []any, filter []any){
	switch data[0] {
		case "notEmpty":
			//VerifyEmpty()

	}
}

func VerifyEmpty(value string, excludeZero bool , msg string){
	data := [][]any{}
	err := json.Unmarshal([]byte(value), &data)
	if err == nil {
		if len(data) < 0 {
			// TODO , 验证失败
		}
		// TODO , 验证成功
	}else{
		if (excludeZero == true){
			if value == "" {
				// TODO , 验证失败
			}
		}else{
			if (value == "") || (value == "0") {
				// TODO , 验证失败
			}
		}
	}

}