package supervisord

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"github.com/kolo/xmlrpc"
	"os"
	"time"
)

type Supervisor struct {
	Host string `toml:"host"`
}

type ProcessInfo struct {
	Name          string `xmlrpc:"name"`
	Group         string `xmlrpc:"group"`
	Description   string `xmlrpc:"description"`
	Start         int64  `xmlrpc:"start"`
	Stop          int64  `xmlrpc:"stop"`
	Now           int64  `xmlrpc:"now"`
	State         int16  `xmlrpc:"state"`
	Statename     string `xmlrpc:"statename"`
	Spawnerr      string `xmlrpc:"spawnerr"`
	Exitstatus    int64  `xmlrpc:"exitstatus"`
	StdoutLogfile string `xmlrpc:"stdout_logfile"`
	StderrLogfile string `xmlrpc:"stderr_logfile"`
	Pid           int64  `xmlrpc:"pid"`
}

type processTags map[string]string

type processFields map[string]interface{}

var (
	defaultHost           = "http://127.0.0.1:9001/RPC2"
	defaultHostname       = "hostname"
	getAllProcessInfoCall = "supervisor.getAllProcessInfo"

	sampleConfig = `
  ## Default supervisor RPC host
  # host = "http://127.0.0.1:9001/RPC2"
`
)

// SampleConfig returns sample configuration message
func (s *Supervisor) SampleConfig() string {
	return sampleConfig
}

// Description returns the plugin description
func (s *Supervisor) Description() string {
	return "Accepts syslog messages following RFC5424 format with transports as per RFC5426, RFC5425, or RFC6587"
}

func (s *Supervisor) Gather(a telegraf.Accumulator) (err error) {
	client, err := xmlrpc.NewClient(s.Host, nil)
	if err != nil {
		return err
	}
	defer client.Close()

	programs := make([]ProcessInfo, 0)
	err = client.Call(getAllProcessInfoCall, nil, &programs)
	if err != nil {
		return err
	}

	for _, program := range programs {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = defaultHostname
		}
		tags := processTags{
			"host":    hostname,
			"program": program.Name,
		}
		fields := program.toMap()
		t := time.Unix(program.Now, 0)
		a.AddFields("supervisor", fields, tags, t)
	}

	return nil
}

func (p *ProcessInfo) toMap() processFields {
	return processFields{
		"name":           p.Name,
		"group":          p.Group,
		"description":    p.Description,
		"start":          p.Start,
		"stop":           p.Stop,
		"now":            p.Now,
		"state":          p.State,
		"statename":      p.Statename,
		"spawnerr":       p.Spawnerr,
		"exitstatus":     p.Exitstatus,
		"stdout_logfile": p.StdoutLogfile,
		"stderr_logfile": p.StderrLogfile,
		"pid":            p.Pid,
	}
}

func init() {
	inputs.Add("supervisor", func() telegraf.Input {
		return &Supervisor{
			Host: defaultHost,
		}
	})
}
