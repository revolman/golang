// Разобрал реализацию renatofq. Интересный алгоритм.
// Но судя по всему можно было бы сделать тупее и проще.
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

type app struct {
	ts     *trackSort
	tracks []*Track
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func main() {
	app := &app{ts: defLessFuncs(tracks), tracks: tracks}
	printTracks(tracks)
	fmt.Println()
	sortBy := "Length"
	app.ts.SortBy(sortBy)
	sort.Sort(app.ts)
	printTracks(tracks)
	fmt.Println()
	sortBy = "Title"
	app.ts.SortBy(sortBy)
	sort.Sort(app.ts)
	printTracks(tracks)
}
