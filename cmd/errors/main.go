package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/webhippie/errors/pkg/version"
	"gopkg.in/urfave/cli.v2"
)

const (
	// FormatHeader defined the format header name.
	FormatHeader = "X-Format"

	// CodeHeader defines the code header name.
	CodeHeader = "X-Code"

	// ContentType defines the content type header name.
	ContentType = "Content-Type"
)

type payload struct {
	Status int
	Error  string
}

func main() {
	app := &cli.App{
		Name:    "errors",
		Version: version.Version.String(),
		Usage:   "display proper error documents",
		Authors: authors(),
		Flags:   flags(),
		Before:  before(),
		Action:  action(),
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print the current version if this",
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func authors() []*cli.Author {
	return []*cli.Author{
		{
			Name:  "Thomas Boerger",
			Email: "thomas@webhippie.de",
		},
	}

}

func flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "address",
			Value:   "0.0.0.0:8080",
			Usage:   "address to bind the server",
			EnvVars: []string{"ERRORS_ADDRESS"},
		},
		&cli.StringFlag{
			Name:    "assets",
			Value:   "assets/",
			Usage:   "folder with required assets",
			EnvVars: []string{"ERRORS_ASSETS"},
		},
	}
}

func before() cli.BeforeFunc {
	return func(c *cli.Context) error {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		return nil
	}
}

func action() cli.ActionFunc {
	return func(c *cli.Context) error {
		http.HandleFunc("/", handler(c.String("assets")))
		http.HandleFunc("/healthz", healthz())
		http.HandleFunc("/readyz", readyz())

		http.Handle("/metrics", promhttp.Handler())

		var gr run.Group

		{
			server := &http.Server{
				Addr:         c.String("address"),
				ReadTimeout:  5 * time.Second,
				WriteTimeout: 10 * time.Second,
			}

			gr.Add(func() error {
				log.Info().
					Str("addr", c.String("address")).
					Msg("starting server")

				return server.ListenAndServe()
			}, func(reason error) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				if err := server.Shutdown(ctx); err != nil {
					log.Error().
						Err(err).
						Msg("failed to shutdown gracefully")

					return
				}

				log.Info().
					Err(reason).
					Msg("http shutdown gracefully")
			})
		}

		{
			stop := make(chan os.Signal, 1)

			gr.Add(func() error {
				signal.Notify(stop, os.Interrupt)

				<-stop

				return nil
			}, func(err error) {
				close(stop)
			})
		}

		return gr.Run()
	}
}

func handler(assets string) func(http.ResponseWriter, *http.Request) {
	templates := templates(assets)
	errors := errors(assets)

	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		defer func(start time.Time, proto string) {
			duration := time.Since(start).Seconds() * 1e3

			requestCounter.WithLabelValues(proto).Inc()
			requestDuration.WithLabelValues(proto).Observe(duration)
		}(start, fmt.Sprintf("%d.%d", r.ProtoMajor, r.ProtoMinor))

		status := status(r)
		format := format(r)
		ext := ext(format)

		w.Header().Set(ContentType, format)
		w.WriteHeader(status)

		msg, ok := errors[status]

		if !ok {
			msg = http.StatusText(status)
		}

		payload := payload{
			Status: status,
			Error:  msg,
		}

		if err := templates.ExecuteTemplate(w, ext, payload); err != nil {
			log.Error().
				Err(err).
				Msg("failed to execute template")

			io.WriteString(w, http.StatusText(status))
		}
	}
}

func healthz() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, http.StatusText(http.StatusOK))
	}
}

func readyz() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, http.StatusText(http.StatusOK))
	}
}

func status(r *http.Request) int {
	status, err := strconv.Atoi(r.Header.Get(CodeHeader))

	if err != nil {
		status = 404

		log.Info().
			Msg("code not specified, using default")
	}

	return status
}

func format(r *http.Request) string {
	format := r.Header.Get(FormatHeader)

	if format == "" {
		format = "text/html"

		log.Info().
			Msg("format not specified, using default")
	}

	return format
}

func ext(format string) string {
	mediaType, _, _ := mime.ParseMediaType(format)
	cext, err := mime.ExtensionsByType(mediaType)

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to parse media type extension")

		return "html"
	}

	if len(cext) == 0 {
		log.Info().
			Msg("could not detect media type extension")

		return "html"
	}

	return strings.TrimPrefix(
		cext[0],
		".",
	)
}

func templates(assets string) *template.Template {
	tpls := template.New(
		"",
	)

	if stat, err := os.Stat(assets); err == nil && stat.IsDir() {
		files := []string{}

		filepath.Walk(assets, func(filename string, f os.FileInfo, err error) error {
			if f.IsDir() {
				return nil
			}

			if !strings.HasPrefix(path.Base(filename), "template") {
				return nil
			}

			files = append(
				files,
				filename,
			)

			return nil
		})

		for _, name := range files {
			file, errRead := ioutil.ReadFile(name)

			if errRead != nil {
				log.Error().
					Err(errRead).
					Str("file", name).
					Msg("failed to read template")

				continue
			}

			base := strings.TrimPrefix(
				path.Ext(
					name,
				),
				".",
			)

			_, errParse := tpls.New(
				base,
			).Parse(
				string(file),
			)

			if errParse != nil {
				log.Error().
					Err(errParse).
					Str("file", name).
					Msg("failed to parse template")

				continue
			}
		}
	} else {
		log.Error().
			Str("dir", assets).
			Msg("assets dir doesn't exist")
	}

	return tpls
}

func errors(assets string) map[int]string {
	result := make(map[int]string)
	name := path.Join(assets, "errors.json")

	if _, err := os.Stat(name); os.IsNotExist(err) {
		log.Error().
			Err(err).
			Str("file", name).
			Msg("errors file doesn't exist")

		return result
	}

	content, err := ioutil.ReadFile(name)

	if err != nil {
		log.Error().
			Err(err).
			Str("file", name).
			Msg("failed to load errors file")

		return result
	}

	if err := json.Unmarshal(content, &result); err != nil {
		log.Error().
			Err(err).
			Str("file", name).
			Msg("failed to parse errors file")

		return result
	}

	return result
}
