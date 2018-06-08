// Functions to access api on jakobmorrison.net.

package main

import "gopkg.in/resty.v1"
import "fmt"
import "os"
import "strconv"
import "strings"


// POST new data to database
func Add_Data(data_json string) (string, string){ 
	// POST api call		
	resp, err := resty.R().
      SetHeader("Content-Type", "application/json").
      SetBody(data_json).
      Post("http://jakobmorrison.net/definitions/api/")
    // Error Handling
	if err != nil {
		s := err.Error()
    	Write_Debug_Log(s)
	}
	// Write response results to log file (debug.txt)
	debug_line := resp.ReceivedAt().String()+ " Command: POST, Status: "+strconv.Itoa(resp.StatusCode())+", Time: "+resp.Time().String()+", Response: "+resp.String()+"\n"
    Write_Debug_Log(debug_line)

    // Return Response Time
    return resp.String(), resp.Time().String()

}

// DELETE data from database
func Remove_Data(word string) (string, string){
	// DELETE api call	
	resp, err := resty.R().
      Delete("http://jakobmorrison.net/definitions/api/word="+word)
    // Error Handling
	if err != nil {
    	Write_Debug_Log(err.Error())
	}
	// Write response results to log file (debug.txt)
	debug_line := resp.ReceivedAt().String()+ " Command: DELETE, Status: "+strconv.Itoa(resp.StatusCode())+", Time: "+resp.Time().String()+", Response: "+resp.String()+"\n"
    Write_Debug_Log(debug_line)

    // Return Response Time
    return resp.String(), resp.Time().String()

}

// GET data by id from database
func Get_Data_Id(id string) (string, string){
	// GET api call
	resp, err := resty.R().
      Get("http://jakobmorrison.net/definitions/api/id="+id)
    // Error Handling
	if err != nil {
    	Write_Debug_Log(err.Error())
	}
	// Write response results to log file (debug.txt)
	debug_line := resp.ReceivedAt().String()+ " Command: GET, Status: "+strconv.Itoa(resp.StatusCode())+", Time: "+resp.Time().String()+", Response: "+resp.String()+"\n"
	Write_Debug_Log(debug_line)

	// Return data
	json_data := resp.String()
	json_data = strings.Replace(json_data, "[", "", -1)
	json_data = strings.Replace(json_data, "]", "", -1)
    return json_data, resp.Time().String()
}

// GET data by word from database
func Get_Data_Word(word string) (string, string){
	// GET api call
	resp, err := resty.R().
      Get("http://jakobmorrison.net/definitions/api/word="+word)
    // Error Handling
	if err != nil {
    	Write_Debug_Log(err.Error())
	}
	// Write response results to log file (debug.txt)
	debug_line := resp.ReceivedAt().String()+ " Command: GET, Status: "+strconv.Itoa(resp.StatusCode())+", Time: "+resp.Time().String()+", Response: "+resp.String()+"\n"
	Write_Debug_Log(debug_line)

	// Return data
	json_data := resp.String()
	json_data = strings.Replace(json_data, "[", "", -1)
	json_data = strings.Replace(json_data, "]", "", -1)
    return json_data, resp.Time().String()
}


// Write to debug log
func Write_Debug_Log(log_line string){
    file, err := os.OpenFile("debug.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    fmt.Fprintf(file, log_line)
}

func Write_Result_File(log_line string){
    file, err := os.OpenFile("result.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    fmt.Fprintf(file, log_line)
}


// Main function. Will run either unit test or benchmark. 
func main() {
	fmt.Println("Running Unit Test and Benchmark. Please See result.txt.")	
	Test()	
	Benchmark()
}
