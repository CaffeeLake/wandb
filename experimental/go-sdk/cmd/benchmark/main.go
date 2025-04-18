package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"

	"github.com/wandb/wandb/experimental/go-sdk/pkg/settings"
	"github.com/wandb/wandb/experimental/go-sdk/pkg/wandb"
)

type BenchOpts struct {
	host               *string
	port               *int
	numHistory         *int
	numHistoryElements *int
	teardown           *bool
	offline            *bool
	numCPUs            *int
	numWorkers         *int
}

type Bench struct {
	opts  BenchOpts
	wandb *wandb.Session
}

func NewBench(benchOpts BenchOpts) *Bench {
	return &Bench{opts: benchOpts}
}

func (b *Bench) Setup() {
	params := wandb.SessionParams{}
	if *b.opts.port != 0 {
		params.Address = fmt.Sprintf("%s:%d", *b.opts.host, *b.opts.port)
	}
	if *b.opts.offline {
		baseSettings, err := settings.New()
		if err != nil {
			panic(err)
		}
		baseSettings.FromSettings(
			&settings.Settings{
				Mode: settings.ModeOffline,
			},
		)
		params.Settings = baseSettings
	}
	var err error
	b.wandb, err = wandb.Setup(&params)
	if err != nil {
		panic(err)
	}
}

func (b *Bench) RunWorkers() {
	if *b.opts.numCPUs != 0 {
		runtime.GOMAXPROCS(*b.opts.numCPUs)
	}
	var wg sync.WaitGroup
	for i := 0; i < *b.opts.numWorkers; i++ {
		wg.Add(1)
		go func() {
			b.Worker()
			wg.Done()
		}()
	}
	wg.Wait()
}

func (b *Bench) Worker() {
	run, err := wandb.Init(&wandb.RunParams{})
	if err != nil {
		panic(err)
	}

	data := make(wandb.History)
	for i := 0; i < *b.opts.numHistoryElements; i++ {
		data[fmt.Sprintf("loss_%d", i)] = float64(100 + i)
	}

	for i := 0; i < *b.opts.numHistory; i++ {
		run.Log(data, true)
	}
	run.Finish()
}

func (b *Bench) Close() {
	if *b.opts.teardown {
		b.wandb.Close()
	}
}

func main() {
	benchOpts := BenchOpts{
		host:               flag.String("host", "localhost", "host to connect to"),
		port:               flag.Int("port", 0, "port to connect to"),
		numHistory:         flag.Int("numHistory", 1000, "number of history records to log"),
		numHistoryElements: flag.Int("numHistoryElements", 5, "number of elements in a history record"),
		teardown:           flag.Bool("close", false, "flag to close the server"),
		offline:            flag.Bool("offline", false, "use offline mode"),
		numCPUs:            flag.Int("numCPUs", 0, "number of cpus"),
		numWorkers:         flag.Int("numWorkers", 1, "number of parallel workers"),
	}
	flag.Parse()

	b := NewBench(benchOpts)
	b.Setup()
	b.RunWorkers()
	b.Close()
}
