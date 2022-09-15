package main

import (
	"errors"
	"github.com/evanw/esbuild/pkg/api"
)

func Build(out string, minify bool, files ...string) error {
	for _, f := range files {
		r := api.Build(api.BuildOptions{
			EntryPoints:       []string{f},
			Outdir:            out,
			Bundle:            true,
			Write:             true,
			Sourcemap:         api.SourceMapNone,
			MinifyWhitespace:  minify,
			MinifyIdentifiers: minify,
			MinifySyntax:      minify,
		})

		if len(r.Errors) > 0 {
			return errors.New(r.Errors[0].Text)
		}
	}

	return nil
}
