// Functions to test access_api functionality. 
package main

import "os"
import "strconv"
import "encoding/json"


// Main test function
func Test(){
	err := os.Remove("result.txt")
    if err != nil {
        panic(err)
    }

	Write_Result_File("TESTING FUNCTIONS\n")
	Write_Result_File("\n")

	// Get unit test
	utest_get10_res := utest_get10()
	utest_get_nonexistent_res := utest_get_nonexistent()
	utest_get_by_word_res := utest_get_by_word()
	utest_get_by_blank_word_res := utest_get_by_blank_word()

	// Add unit test
	utest_post_exist_res := utest_post_exist()
	utest_post_new_word_res := utest_post_new_word()

	// Delete Unit test
	utest_delete_word_res := utest_delete_word()


	Write_Result_File("\n")
	Write_Result_File("TESTING: END Results\n")
	Write_Result_File("--------------\n")
	Write_Result_File("utest_get10: "+utest_get10_res+"\n")
	Write_Result_File("utest_get_nonexistent_res: "+utest_get_nonexistent_res+"\n")
	Write_Result_File("utest_get_by_word_res: "+utest_get_by_word_res+"\n")
	Write_Result_File("utest_get_by_blank_word: "+utest_get_by_blank_word_res+"\n")
	Write_Result_File("utest_post_exist_res: "+utest_post_exist_res+"\n")
	Write_Result_File("utest_post_new_word_res: "+utest_post_new_word_res+"\n")
	Write_Result_File("utest_delete_word_res: "+utest_delete_word_res+"\n")
	Write_Result_File("\n")
}

// Test to grab the first ten entries
func utest_get10()(string){
	utest_get10_pf := "pass"
	Write_Result_File("Get 10 responses \n")
	Write_Result_File("--------------\n")
	for i := 1; i < 11; i++ {
		json_data, response_time := Get_Data_Id(strconv.Itoa(i))
		_ = response_time
		if json_data != ""{
			Write_Result_File("PASS: "+json_data+"\n")
		}else{
			Write_Result_File("FAIL: Id = "+strconv.Itoa(i)+"\n")
			if utest_get10_pf == "pass"{
				utest_get10_pf = "fail"
			}
		}
	}
	return utest_get10_pf
}

// Test to handle no existent ids. 
func utest_get_nonexistent()(string){
	utest_get_nonexistent_pf := "pass"
	Write_Result_File("\n")
	Write_Result_File("Get Non Existent ID \n")
	Write_Result_File("--------------\n")
	json_data, response_time := Get_Data_Id("-1")
	_ = response_time
	if json_data == ""{
		Write_Result_File("PASS: NO Data Returned. \n")
	}else{
			Write_Result_File("Fail (Should have returned nothing): "+json_data+"\n")
			if utest_get_nonexistent_pf == "pass"{
				utest_get_nonexistent_pf = "fail"
			}
		}
	return utest_get_nonexistent_pf
}

// Get data by word 
func utest_get_by_word()(string){
	utest_get_by_word_pf := "pass"
	Write_Result_File("\n")
	Write_Result_File("Get data by word \n")
	Write_Result_File("--------------\n")
	json_data, response_time := Get_Data_Word("Vexed")
	_ = response_time
	if json_data != ""{
		Write_Result_File("Pass: "+json_data+"\n")
	}else{
			Write_Result_File("FAIL: word = Vexed\n")
			if utest_get_by_word_pf == "pass"{
				utest_get_by_word_pf = "fail"
			}
		}
	return utest_get_by_word_pf
}

// Get data by blank 
func utest_get_by_blank_word()(string){
	utest_get_by_blank_word_pf := "pass"
	Write_Result_File("\n")
	Write_Result_File("Get data by blank word \n")
	Write_Result_File("--------------\n")
	json_data, response_time := Get_Data_Word("")
	_ = response_time
	if json_data == ""{
		Write_Result_File("PASS: NO Data Returned. \n")
	}else{
			Write_Result_File("Fail (Should have returned nothing): "+json_data+"\n")
			if utest_get_by_blank_word_pf == "pass"{
				utest_get_by_blank_word_pf = "fail"
			}
		}
	return utest_get_by_blank_word_pf
}


//Add items that already exist
func utest_post_exist()(string){
	utest_post_exist_pf := "pass"
	Write_Result_File("\n")
	Write_Result_File("Add data that exist \n")
	Write_Result_File("--------------\n")

	json_data, response_time_get := Get_Data_Id("1")
	_ = response_time_get
	response_add, response_time_add := Add_Data(json_data)
	_ = response_time_add

	data := make(map[string]interface{})
	json.Unmarshal([]byte(response_add), &data)

    val := data["word"].([]interface{})[0]
    val_str := val.(string)

	if val_str == "words with this word already exists."{
		Write_Result_File("PASS: words with this word already exists.. \n")
	}else{
		Write_Result_File("Fail")
		if utest_post_exist_pf == "pass"{
			utest_post_exist_pf = "fail"
		}	
	}	
	return utest_post_exist_pf
}


//Add new word
func utest_post_new_word()(string){
	utest_post_new_word_pf := "pass"
	Write_Result_File("\n")
	Write_Result_File("Add new word \n")
	Write_Result_File("--------------\n")

	json_data := "{\"word\":\"test\",\"type\":\"test\",\"definition\":\"test\"}"
	response_add, response_time_add := Add_Data(json_data)
	_ = response_time_add


	if response_add != ""{
		Write_Result_File("PASS: "+response_add+"\n")
	}else{
		Write_Result_File("Fail: Data not added.")
		if utest_post_new_word_pf == "pass"{
			utest_post_new_word_pf = "fail"
		}	
	}	
	return utest_post_new_word_pf
}

func utest_delete_word()(string){
	utest_delete_word_pf := "pass"
	Write_Result_File("\n")
	Write_Result_File("Delete Word \n")
	Write_Result_File("--------------\n")

	response_remove, response_time_remove := Remove_Data("test")
	_ = response_time_remove

	if response_remove == ""{
		Write_Result_File("PASS: Deleted. \n")
	}else{
			Write_Result_File("Fail (Should have returned nothing): "+response_remove+"\n")
			if utest_delete_word_pf == "pass"{
				utest_delete_word_pf = "fail"
			}
		}
	return utest_delete_word_pf

}
