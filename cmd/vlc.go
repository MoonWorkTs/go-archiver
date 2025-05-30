package cmd

import (
	"archiver/lib/vlc"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	filePath := args[0]

	r, err := os.Open(filePath)

	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)

	if err != nil {
		handleErr(err)
	}

	packed := vlc.Encode(string(data))

	if err := os.WriteFile(packedFileName(filePath), []byte(packed), 0644); err != nil {
		handleErr(err)
	}
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)
	fileExt := filepath.Ext(fileName)

	return strings.TrimSuffix(fileName, fileExt) + "." + packedExtension + fileExt
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
