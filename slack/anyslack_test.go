package slack_test
import (
	"slack"
	"testing"
)

func TestGoSlack(t *testing.T) {
	resp, err := slack.GoSlack(0)
	if err != nil{
		  t.Fatal(err)
	}
	if resp.StatusCode != 200{
		t.Errorf("StatusCode is %d", resp.StatusCode)
	}
}