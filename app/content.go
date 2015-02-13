package main

import (
    "io/ioutil"
    "net/http"
    "encoding/json"
    "fmt"
    "errors"
)

type Content struct {
	text string
	url string
}

func generateTweetContent() (string, error) {
	contents, err := callAPI()
	if( err != nil ) {
		return "", err
	}

	for _, content := range contents {
		// TODO check already tweeted

		return content.text+" "+content.url, nil
	}
	

	return "", errors.New("No tweet content found")
}

func callAPI() ([]Content, error) {
	content, err := getContent(DATA_SOURCE)
	if err != nil {
		fmt.Println("Error while calling API")
		return nil, err
	} else {
	    // Fill the record with the data from the JSON
	    var jsonObj map[string]interface{}
	    err = json.Unmarshal(content, &jsonObj)
	    if err != nil {
	    	fmt.Println("An error occurred while converting our JSON to an object")
	    	return nil, err
	    }

	    contents := make([]Content, 0)

	    if val, ok := jsonObj["results"]; ok {

	    	if results, ok := val.(map[string]interface{}); ok {

			    if val, ok := results["collection1"]; ok {
			    	
			    	if collection, ok := val.([]interface{}); ok {
			    		
			    		for _, item := range collection {

			    			if property, ok := item.(map[string]interface{}); ok {

			    				if property1, ok := property["property1"]; ok {

			    					if contentData, ok := property1.(map[string]interface{}); ok {

			    						if text, ok := contentData["text"].(string); ok {

			    							if href, ok := contentData["href"].(string); ok {

			    								contents = append(contents, Content{text: text, url: href})

			    							} else {
			    								fmt.Println("Error mapping href as string")
			    								return nil, errors.New("json mapping error")
			    							}

			    						} else {
		    								fmt.Println("Error mapping text as string")
		    								return nil, errors.New("json mapping error")
		    							}

			    					} else {
			    						fmt.Println("Error mapping property1 as json obj")
			    						return nil, errors.New("json mapping error")
			    					}

			    				} else {
			    					fmt.Println("Missing mappings property1")
			    					return nil, errors.New("json mapping error")
			    				}

			    			} else {
			    				fmt.Println("Error mappings property1")
			    				return nil, errors.New("json mapping error")
			    			}

			    		}

			    	} else {
			    		fmt.Println("Error mappings collection1 as array")
			    		return nil, errors.New("json mapping error")
			    	}

			    } else {
			    	fmt.Println("Error mappings collection1")
			    	return nil, errors.New("json mapping error")
			    }

			} else {
				fmt.Println("Error mappings results as json obj")
			    return nil, errors.New("json mapping error")
			}

			return contents, nil

		} else {
			fmt.Println("No field results in json")
			return nil, errors.New("json mapping error")
		}
	}
}

func getContent(url string) ([]byte, error) {
    // Build the request
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
      return nil, err
    }
    // Send the request via a client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
      return nil, err
    }
    // Defer the closing of the body
    defer resp.Body.Close()
    // Read the content into a byte array
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return nil, err
    }
    // At this point we're done - simply return the bytes
    return body, nil
}