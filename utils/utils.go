package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	resty "github.com/go-resty/resty/v2"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pkg/errors"
	"github.com/s1ntaxe770r/templit/models"
	"github.com/spf13/cobra"
)

const releaseurl = `https://api.github.com/repos/s1ntaxe770r/templit/releases/latest`

// GetReleaseVersion returns the release version the user is currently running
func GetReleaseVersion() (string, error) {
	fmt.Println(color.YellowString("checking for a new release, hang tight"))
	s := spinner.New(spinner.CharSets[2], 100*time.Millisecond)
	s.Start()
	resp, err := http.Get(releaseurl)
	if err != nil {
		return "", errors.Wrapf(err, "failed to make GET request to %s", releaseurl)
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to unmarshal response body")
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", errors.Wrap(err, "failed to unmarshal json into object")
	}
	releasetag := ""
	if data["tag_name"] != nil {
		releasetag = data["tag_name"].(string)
	}
	s.Stop()
	return releasetag, nil
}

// GetTemplateDir returns template directory as a string
func GetTemplateDir() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	dir := home + "/.config/templit/templates"
	return dir
}

func CopyTemplate(src, dst string) error {
srcFile, err := os.Open(src)
if err != nil {
return err 
}
    defer srcFile.Close()

    destFile, err := os.Create(dst) // creates if file doesn't exist
    if err != nil{
	    return err
    }
    defer destFile.Close()

    _, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
    if err != nil {
	    return err
    }
    err = destFile.Sync()
    if err !=nil{
	    return err
    }
	return nil
}

func CopyRemoteTemplate(url,filename string) error{
	client := resty.New()
	color.HiYellowString("obtaining template from remote url. hang tight ⚙️")
	s := spinner.New(spinner.CharSets[2], 100*time.Millisecond)
	s.Color("yellow")
	s.Start()
	_, err := client.R().
		SetOutput(GetTemplateDir()+"/"+filename).
		Get(url)
	if err != nil {
		return  err
	}
	s.Stop()
	color.HiGreenString("done ✅")
	return nil
}

func TemplatetoTable(config []models.Template){
	t := table.NewWriter()
	tTemp := table.Table{}
	tTemp.Render()
	for _ , template := range config {
		t.AppendRow([]interface{}{color.YellowString(template.Name), color.MagentaString("->"), color.GreenString(template.Path)})
	}
	fmt.Println(t.Render())
}
