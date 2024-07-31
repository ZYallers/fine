package route

import (
	"bytes"
	"fmt"
	"gitlab.sys.hxsapp.net/hxs/fine/os/fcfg"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type dump struct {
	rows []row
}

type row struct {
	path        string
	method      string
	handlerName string
}

var Dumper = &dump{rows: make([]row, 0)}

func (d *dump) Append(path, method, handlerName string) {
	row := row{path: path, method: method, handlerName: handlerName}
	d.rows = append(d.rows, row)
}

func (d *dump) Dump() {
	if len(d.rows) == 0 {
		return
	}
	buffer := bytes.NewBuffer(nil)
	table := tablewriter.NewWriter(buffer)
	headers := []string{"NO", "APP_NAME", "SERVER_ADDRESS", "PATH", "METHOD", "HANDLER"}
	table.SetHeader(headers)
	table.SetRowLine(true)
	table.SetBorder(false)
	table.SetCenterSeparator("|")
	var appName = fcfg.GetString("app.name")
	var address = fcfg.GetString("app.server.addr")
	for i, row := range d.rows {
		table.Append([]string{
			strconv.Itoa(i + 1),
			appName,
			address,
			row.path,
			strings.ToUpper(row.method),
			row.handlerName,
		})
	}
	table.Render()
	fmt.Printf("\n%s\n", buffer.String())
}
