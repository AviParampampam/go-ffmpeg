package ffmpeg

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/thanhpk/randstr"
)

var mu sync.Mutex

type Worker struct {
	cmd           *exec.Cmd
	cmdBackup     exec.Cmd
	cmdString     string
	murderPlanned bool
}

// Start starts the specified worker but does not wait for it to complete.
func (w *Worker) Start() error {
	return w.cmd.Start()
}

// Run starts the specified worker and waits for it to complete.
func (w *Worker) Run() error {
	return w.cmd.Run()
}

func (w *Worker) Stop() error {
	w.murderPlanned = true

	if err := w.cmd.Process.Signal(os.Kill); err != nil {
		return err
	}
	return w.cmd.Wait()
}

func (f *FFmpeg) RunOnceWorker(files ...OptionIO) error {
	// Генерация уникального имени для воркера.
	var workername string
	for {
		_workername := "ONCE_" + randstr.String(4)

		mu.Lock()
		_, exists := f.workers[_workername]
		mu.Unlock()

		if !exists {
			workername = _workername
			break
		}
	}

	mu.Lock()
	w, err := f.SetWorker(workername, files...)
	if err != nil {
		return err
	}
	mu.Unlock()

	err = w.Run()

	mu.Lock()
	f.DelWorker(workername)
	mu.Unlock()
	return err
}

func (w *Worker) IsActive() bool {
	return w.cmd.ProcessState == nil
}

// FFmpeg process checker..
func (w *Worker) Cron(timeout time.Duration) {
	for {
		w.cmd.Process.Wait()

		if w.murderPlanned {
			break
		}

		fmt.Fprintf(os.Stderr, "[WARN] the worker stopped the process unexpectedly")

		w.cmd = &(w.cmdBackup)
		if err := w.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "error when trying to restart the worker: %s\n", err)
		}

		time.Sleep(timeout)
	}
}

// Format files to string.
func (ff *FFmpeg) ftos(files []OptionIO) string {
	cmd := filepath.FromSlash(ff.BinPath) + " -loglevel error"

	for _, f := range files {
		cmd += " " + f.String()
	}

	return cmd
}

// Format string to exec.Cmd.
func (ff *FFmpeg) stoc(s string) *exec.Cmd {
	splits := strings.Split(s, " ")
	cmd := exec.Command(splits[0], splits[1:]...)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	if ff.Report != (report{}) {
		_ = os.MkdirAll(filepath.Dir(ff.Report.File), 0755)
		cmd.Env = append(cmd.Env, fmt.Sprintf("FFREPORT=file=%s:level=%d", ff.Report.File, ff.Report.LogLevel))
	}
	return cmd
}

func (ff *FFmpeg) SetWorker(key string, files ...OptionIO) (*Worker, error) {
	if _, find := ff.GetWorker(key); find {
		return nil, errors.New("worker with the same name already exists")
	}

	s := ff.ftos(files)
	// fmt.Println(s)
	cmd := ff.stoc(s)

	w := Worker{cmd: cmd, cmdBackup: *cmd, cmdString: s, murderPlanned: false}

	ff.workers[key] = &w

	return &w, nil
}

func (ff *FFmpeg) GetWorker(key string) (*Worker, bool) {
	w, find := ff.workers[key]
	return w, find
}

func (ff *FFmpeg) DelWorker(key string) {
	w, find := ff.GetWorker(key)
	if !find {
		return
	}

	if w.cmd.Process != nil {
		w.Stop()
	}

	delete(ff.workers, key)
}
