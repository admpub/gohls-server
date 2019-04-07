package cmd

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/admpub/gohls-server/internal/api"
	"github.com/admpub/gohls-server/internal/config"
	"github.com/admpub/gohls-server/internal/hls"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	listen string
)

var reNumber = regexp.MustCompile(`^[0-9]+$`)

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&listen, "listen", "l", "127.0.0.1:8080", "The address to listen on (default is 127.0.0.1:8080)")
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs ths server",
	Long:  `Runs ths server`,
	Run: func(cmd *cobra.Command, args []string) {
		init_hls()

		config, err := config.GetConfig(cfgFile)
		if err != nil {
			log.Fatalf("Error reading config: %v", err)
		}

		api.Setup(config)

		if reNumber.MatchString(listen) {
			listen = `:` + listen
		}

		// Dump information to user
		fmt.Printf("Path to ffmpeg executable: %v\n", hls.FFMPEGPath)
		fmt.Printf("Path to ffprobe executable: %v\n", hls.FFProbePath)
		fmt.Printf("HLS data directory: %v/\n", hls.HomeDir)
		fmt.Printf("Visit http://%v/\n", listen)

		if herr := http.ListenAndServe(listen, nil); herr != nil {
			fmt.Printf("Error listening %v", herr)
		}

	},
}
