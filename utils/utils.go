package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const releaseurl = `https://api.github.com/repos/s1ntaxe770r/templit/releases/latest`

// GetReleaseVersion returns the release version the user is currently running
func GetReleaseVersion() (string, error) {
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
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return  err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	err = out.Sync()
	return err
}
//
//func CopyRemoteTemplate(url string) error{
//
//}
