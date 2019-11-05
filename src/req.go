package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"strings"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
	//query      = flag.String("query", "Bella Wolfine", "Search term")
	maxResults = flag.Int64("max-results", 5, "Max YouTube results")
)

const developerKey = "Your Developer Key"

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main() {
	args := getCmdArgs()
	commandParsing(args)

}

// Print the ID and title of each result in a list as well as a name that
// identifies the list. For example, print the word section name "Videos"
// above a list of video search results, followed by the video ID and title
// of each matching video.
func printIDs(sectionName string, matches map[string]string) {
	d1 := ""
	//i := 1
	for id, title := range matches {
		d1 += fmt.Sprintf("%v\n%v\n", id, title)
		//i++
	}
	//d1 += fmt.Sprintf("\n\n")



    d2 := []byte(d1)
    err := ioutil.WriteFile("/tmp/vidOut", d2, 0644)
    check(err)

}

func getCmdArgs() []string{
	return os.Args[1:]
}

func commandParsing(args []string){
	if(args[0] == "search" || args[0] == "Search"){
		flag.Parse()
		stringRes := strings.Join(args[1:]," ")
		query := flag.String("query", stringRes, "Search term")
		client := &http.Client{
			Transport: &transport.APIKey{Key: developerKey},
		}
	
		service, err := youtube.New(client)
		if err != nil {
			log.Fatalf("Error creating new YouTube client: %v", err)
		}
	
		// Make the API call to YouTube.
		call := service.Search.List("id,snippet").
			Q(*query).
			MaxResults(*maxResults)
		response, err := call.Do()
		//handleError(err, "")
	
		// Group video, channel, and playlist results in separate lists.
		videos := make(map[string]string)
		channels := make(map[string]string)
		playlists := make(map[string]string)
	
		// Iterate through each item and add it to the correct list.
		for _, item := range response.Items {
			switch item.Id.Kind {
			case "youtube#video":
				videos[item.Id.VideoId] = item.Snippet.Title
			case "youtube#channel":
				channels[item.Id.ChannelId] = item.Snippet.Title
			case "youtube#playlist":
				playlists[item.Id.PlaylistId] = item.Snippet.Title
			}
		}
	
		printIDs("Videos", videos)
	} else {
		d1 := []byte("Error: invalid query")
		err := ioutil.WriteFile("/tmp/vidOut", d1, 0644)
    	check(err)
	}
}