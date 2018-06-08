// Functions to benchmark speed of api response. 
package main

func Benchmark(){
	Write_Result_File("BENCHMARK\n")
	Write_Result_File("\n")

	get_data, response_time_get := Get_Data_Word("Vexed")
	Write_Result_File("GET Access Time: "+response_time_get+"\n")
	_ = get_data


	json_data := "{\"word\":\"test\",\"type\":\"test\",\"definition\":\"test\"}"
	response_add, response_time_add := Add_Data(json_data)
	_ = response_add
	Write_Result_File("Post Time: "+response_time_add+"\n")

	response_remove, response_time_remove := Remove_Data("test")
	_ = response_remove
	Write_Result_File("Delete Time: "+response_time_remove+"\n")	

}
