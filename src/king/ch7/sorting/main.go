package main

import (
	"time"
	"text/tabwriter"
	"os"
	"fmt"
)

type Track struct {
	Title string
	Artist string
	Album string
	Year string
	Length time.Duration
}
	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s"),}
}
	func length(s string) time.Duration {
		d, err := time.ParseDuration(s)
		if err != nil {
			panic(s)
		}

		return d
}

	func printTracks(tracks []*Track) {
		const format  = "%v\t%v\t%v\t%v\t\n"
		tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
		fmt.Fprintf(tw, format, "Title", "Artist", "Alum", "Year", "Length")
		fmt.Fprintf(tw, format, "-----", "------", "------", "------")
		for _,t := range tracks {
			fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
		}
		tw.Flush()
	}

