package main

//{"month": "4", "num": 571, "link": "", "year": "2009", "news": "", "safe_title": "Can't Sleep", "transcript": "[[Someone is in bed, presumably trying to sleep. The top of each panel is a thought bubble showing sheep leaping over a fence.]]\n1 ... 2 ...\n<<baaa>>\n[[Two sheep are jumping from left to right.]]\n\n... 1,306 ... 1,307 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow.]]\n\n... 32,767 ... -32,768 ...\n<<baaa>> <<baaa>> <<baaa>> <<baaa>> <<baaa>>\n[[A whole flock of sheep is jumping over the fence from right to left. The would-be sleeper is sitting up.]]\nSleeper: ?\n\n... -32,767 ... -32,766 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow over his head.]]\n\n{{Title text: If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.}}", "alt": "If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.", "img": "https://imgs.xkcd.com/comics/cant_sleep.png", "title": "Can't Sleep", "day": "20"}

/*
{
   "month": "4",
   "num": 571,
   "link": "",
   "year": "2009",
   "news": "",
   "safe_title": "Can't Sleep",
   "transcript": "[[Someone is in bed, presumably trying to sleep. The top of each panel is a thought bubble showing sheep leaping over a fence.]]\n1 ... 2 ...\n\u003c\u003cbaaa\u003e\u003e\n[[Two sheep are jumping from left to right.]]\n\n... 1,306 ... 1,307 ...\n\u003c\u003cbaaa\u003e\u003e\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow.]]\n\n... 32,767 ... -32,768 ...\n\u003c\u003cbaaa\u003e\u003e \u003c\u003cbaaa\u003e\u003e \u003c\u003cbaaa\u003e\u003e \u003c\u003cbaaa\u003e\u003e \u003c\u003cbaaa\u003e\u003e\n[[A whole flock of sheep is jumping over the fence from right to left. The would-be sleeper is sitting up.]]\nSleeper: ?\n\n... -32,767 ... -32,766 ...\n\u003c\u003cbaaa\u003e\u003e\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow over his head.]]\n\n{{Title text: If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.}}",
   "alt": "If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.",
   "img": "https://imgs.xkcd.com/comics/cant_sleep.png",
   "title": "Can't Sleep",
   "day": "20"
}
*/
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type XKCDResp struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func SearchSingle() {
	resp, err := http.Get("https://xkcd.com/571/info.0.json")
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return
	}

	var result XKCDResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		fmt.Println(err)
		return
	}

	datafmt, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		fmt.Printf("JSON marshaling failed:%s", err)
		return
	}
	fmt.Printf("%s\n", datafmt)

	resp.Body.Close()
}

func main() {
	SearchSingle()
}
