package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

var (
	flags = struct {
		AppendURI, Permanent *bool
		Dest, ListenAddr     *string
	}{
		flag.Bool("append-uri", true, "append URI?"),
		flag.Bool("permanent", false, "use 301 permanent redirect"),
		flag.String("dest", "", "destination URL"),
		flag.String("listen-addr", "0.0.0.0:8000", "network address to liste on"),
	}

	statusCode int
)

func redirect(w http.ResponseWriter, r *http.Request) {
	dest := *flags.Dest
	if *flags.AppendURI {
		dest += r.RequestURI
	}

	http.Redirect(w, r, dest, statusCode)
}

func main() {
	flag.Parse()

	*flags.Dest = strings.Trim(*flags.Dest, "/")
	if *flags.Dest == "" {
		log.Fatalln("please pass a destination URL")
	}

	if *flags.Permanent {
		statusCode = 301
	} else {
		statusCode = 302
	}

	log.Printf("listening on %s\n", *flags.ListenAddr)
	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(*flags.ListenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
