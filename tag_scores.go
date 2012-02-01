package stackongo

import (
	"os"
	"strings"
)

func (session Session) getTagScores(path string, params map[string]string) (output []TagScore, error os.Error) {
	// make the request
	response, err := session.get(path, params)

	if err != nil {
		return output, err
	}

	parsed_response, error := parseResponse(response, new(tagScoresCollection))
	collection := parsed_response.(*tagScoresCollection)

	if error != nil {
		//overload the generic error with details
		error = os.NewError(collection.Error_name + ": " + collection.Error_message)
	} else {
		output = collection.Items
	}

	return output, error

}

// TopAnswerers returns the top 30 answerers active in a single tag, for the given period 
// period should be either `all_time` or `month` 
func (session Session) TopAnswerers(tag string, period string, params map[string]string) (output []TagScore, error os.Error) {
	request_path := strings.Join([]string{"tags", tag, "top-answerers", period}, "/")
	return session.getTagScores(request_path, params)
}

// TopAskers returns the top 30 askers active in a single tag, for the given period 
// period should be either `all_time` or `month` 
func (session Session) TopAskers(tag string, period string, params map[string]string) (output []TagScore, error os.Error) {
	request_path := strings.Join([]string{"tags", tag, "top-askers", period}, "/")
	return session.getTagScores(request_path, params)
}
