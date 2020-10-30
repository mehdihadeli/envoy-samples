/*
Copyright 2018 Turbine Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package file

import (
	"bytes"
	"errors"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	fsnotify "gopkg.in/fsnotify.v1"

	"github.com/turbinelabs/api"
	"github.com/turbinelabs/codec"
	tbnos "github.com/turbinelabs/nonstdlib/os"
	"github.com/turbinelabs/rotor/updater"
	"github.com/turbinelabs/test/assert"
	"github.com/turbinelabs/test/check"
	"github.com/turbinelabs/test/matcher"
	"github.com/turbinelabs/test/tempfile"
)

const (
	YamlInput = `
- cluster: c1
  instances:
  - host: c1h1
    port: 8000
    metadata:
    - key: c1h1m1
      value: c1h1v1
    - key: c1h1m2
      value: c1h1v2
  - host: c1h2
    port: 8001
    metadata:
      - key: c1h2m1
        value: c1h2v1
- cluster: c2
  instances:
  - host: c2h1
    port: 8000
    metadata:
    - key: c2h1m1
      value: c2h1v1
`

	YamlInputWithDuplicateCluster = `
- cluster: c1
  instances:
  - host: c1h1
    port: 8000
- cluster: c1
  instances:
  - host: c1h2
    port: 8001
`

	SimpleYamlInput = `
- cluster: c1
  instances:
  - host: c1h1
    port: 8000
    metadata:
    - key: c1h1m1
      value: c1h1v1
`

	JsonInput = `
[
  {
    "cluster": "c1",
    "instances": [
      {
        "host": "c1h1",
        "port": 8000,
        "metadata": [
          { "key": "c1h1m1", "value": "c1h1v1" },
          { "key": "c1h1m2", "value": "c1h1v2" }
        ]
      },
      {
        "host": "c1h2",
        "port": 8001,
        "metadata": [
          { "key": "c1h2m1", "value": "c1h2v1" }
        ]
      }
    ]
  },
  {
    "cluster": "c2",
    "instances": [
      {
        "host": "c2h1",
        "port": 8000,
        "metadata": [
          { "key": "c2h1m1", "value": "c2h1v1" }
        ]
      }
    ]
  }
]`
)

var (
	simpleTestClusters = []api.Cluster{
		{
			Name: "c1",
			Instances: []api.Instance{
				{
					Host: "c1h1",
					Port: 8000,
					Metadata: api.Metadata{
						{Key: "c1h1m1", Value: "c1h1v1"},
					},
				},
			},
		},
	}

	expectedClusters = []api.Cluster{
		{
			Name: "c1",
			Instances: []api.Instance{
				{
					Host: "c1h1",
					Port: 8000,
					Metadata: []api.Metadatum{
						{
							Key:   "c1h1m1",
							Value: "c1h1v1",
						},
						{
							Key:   "c1h1m2",
							Value: "c1h1v2",
						},
					},
				},
				{
					Host: "c1h2",
					Port: 8001,
					Metadata: []api.Metadatum{
						{
							Key:   "c1h2m1",
							Value: "c1h2v1",
						},
					},
				},
			},
		},
		{
			Name: "c2",
			Instances: []api.Instance{
				{
					Host: "c2h1",
					Port: 8000,
					Metadata: []api.Metadatum{
						{
							Key:   "c2h1m1",
							Value: "c2h1v1",
						},
					},
				},
			},
		},
	}
)

func makeYamlFileCollector() *fileCollector {
	return &fileCollector{parser: mkParser(codec.NewYaml()), os: tbnos.New()}
}

func makeJSONFileCollector() *fileCollector {
	return &fileCollector{parser: mkParser(codec.NewJson()), os: tbnos.New()}
}

func makeFileCollectorAndMock(
	t *testing.T,
) (*fileCollector, *gomock.Controller, *updater.MockUpdater) {
	fileCollector := makeYamlFileCollector()

	ctrl := gomock.NewController(assert.Tracing(t))
	mockUpdater := updater.NewMockUpdater(ctrl)

	fileCollector.updater = mockUpdater

	return fileCollector, ctrl, mockUpdater
}

func TestFileCollectorParseYaml(t *testing.T) {
	fileCollector := makeYamlFileCollector()

	clusters, err := fileCollector.parser(bytes.NewBufferString(YamlInput))
	assert.Nil(t, err)
	assert.HasSameElements(t, clusters, expectedClusters)
}

func TestFileCollectorParseJson(t *testing.T) {
	fileCollector := makeJSONFileCollector()

	clusters, err := fileCollector.parser(bytes.NewBufferString(JsonInput))
	assert.Nil(t, err)
	assert.HasSameElements(t, clusters, expectedClusters)
}

func TestFileCollectorJsonYamlEqual(t *testing.T) {
	yamlCollector := makeYamlFileCollector()
	jsonCollector := makeJSONFileCollector()

	yamlClusters, err := yamlCollector.parser(bytes.NewBufferString(YamlInput))
	assert.Nil(t, err)

	jsonClusters, err := jsonCollector.parser(bytes.NewBufferString(JsonInput))
	assert.Nil(t, err)

	assert.HasSameElements(t, yamlClusters, jsonClusters)
}

func TestFileCollectorParseDuplicateClusters(t *testing.T) {
	yamlCollector := makeYamlFileCollector()

	clusters, err := yamlCollector.parser(bytes.NewBufferString(YamlInputWithDuplicateCluster))
	assert.NonNil(t, err)
	assert.Nil(t, clusters)
}

func TestFileCollectorParseError(t *testing.T) {
	yamlCollector := makeYamlFileCollector()

	clusters, err := yamlCollector.parser(bytes.NewBufferString("nope nope nope"))
	assert.Nil(t, clusters)
	assert.NonNil(t, err)
}

func testReload(t *testing.T, collector *fileCollector, data string) error {
	tempFile, cleanup := tempfile.Write(t, data, "filecollector-reload")
	defer cleanup()
	collector.file = tempFile

	return collector.reload()
}

func TestFileCollectorReload(t *testing.T) {
	yamlCollector, ctrl, mockUpdater := makeFileCollectorAndMock(t)
	defer ctrl.Finish()

	mockUpdater.EXPECT().Replace(simpleTestClusters)

	error := testReload(t, yamlCollector, SimpleYamlInput)
	assert.Nil(t, error)
}

func TestFileCollectorReloadParseError(t *testing.T) {
	yamlCollector, ctrl, _ := makeFileCollectorAndMock(t)
	defer ctrl.Finish()

	error := testReload(t, yamlCollector, "nope nope nope")
	assert.ErrorContains(t, error, "cannot unmarshal")
}

func TestFileCollectorReloadReadFileError(t *testing.T) {
	yamlCollector := makeYamlFileCollector()

	tempFile, cleanup := tempfile.Make(t, "filecollector-reload")
	cleanup()

	yamlCollector.file = tempFile

	err := yamlCollector.reload()
	assert.True(t, os.IsNotExist(err))
}

func TestFileCollectorStartWatcher(t *testing.T) {
	collector := &fileCollector{parser: mkParser(codec.NewYaml()), os: tbnos.New()}

	tempDir := tempfile.TempDir(t, "filecollector-watcher")
	defer tempDir.Cleanup()

	collector.file = tempDir.Make(t)

	events, errors, cleanup, err := collector.startWatcher()
	assert.Nil(t, err)
	defer cleanup.Close()

	assert.ChannelEmpty(t, events)
	assert.ChannelEmpty(t, errors)

	newFile := tempDir.Write(t, "line\n")
	assert.Nil(t, os.Rename(newFile, collector.file))

	<-events
}

func TestFileEventLoopErrorExit(t *testing.T) {
	testFileEventLoop(t, true)
}

func TestFileEventLoopSignalExit(t *testing.T) {
	testFileEventLoop(t, false)
}

func testFileEventLoop(t *testing.T, exitOnErr bool) {
	yamlCollector, ctrl, mockUpdater := makeFileCollectorAndMock(t)
	defer ctrl.Finish()

	mockUpdater.EXPECT().Replace(simpleTestClusters)

	tempFile, cleanup := tempfile.Write(t, SimpleYamlInput, "filecollector-eventloop")
	defer cleanup()

	yamlCollector.file = tempFile

	events := make(chan fsnotify.Event)
	errs := make(chan error)
	signals := make(chan os.Signal)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	var eventLoopResult error
	go func() {
		eventLoopResult = yamlCollector.eventLoop(events, errs, signals)
		waitGroup.Done()
	}()

	events <- fsnotify.Event{
		Name: "wrong file",
		Op:   fsnotify.Create,
	}
	events <- fsnotify.Event{
		Name: tempFile,
		Op:   fsnotify.Remove, // ignored op
	}
	events <- fsnotify.Event{
		Name: tempFile,
		Op:   fsnotify.Create,
	}
	if exitOnErr {
		errs <- errors.New("fail")
	} else {
		signals <- os.Interrupt
	}

	waitGroup.Wait()

	if exitOnErr {
		assert.ErrorContains(t, eventLoopResult, "fail")
	} else {
		assert.Nil(t, eventLoopResult)
	}
}

func TestFileCollectorRun(t *testing.T) {
	testFileCollectorRun(
		t,
		func(tempDir tempfile.Dir, file string) string { return file },
		func(file string, newFile string) {
			assertNilOrDie(t, os.Rename(newFile, file))
		},
	)
}

func assertNilOrDie(t *testing.T, i interface{}) {
	if !check.IsNil(i) {
		assert.Tracing(t).Fatalf("got (%T) %v, want <nil>", i, i)
	}
}

func testFileCollectorRun(
	t *testing.T,
	createWatchedPath func(d tempfile.Dir, f string) string,
	updateWatchedPath func(f string, newF string),
) {
	collector, ctrl, mockUpdater := makeFileCollectorAndMock(t)
	defer ctrl.Finish()

	syncCh := make(chan struct{}, 1)
	sync := func(_ []api.Cluster) { syncCh <- struct{}{} }

	gomock.InOrder(
		mockUpdater.EXPECT().Replace(matcher.SameElements{simpleTestClusters}).Do(sync),
		mockUpdater.EXPECT().Replace(matcher.SameElements{expectedClusters}).Do(sync).MinTimes(1),
		mockUpdater.EXPECT().Close(),
	)

	// Need to put temp file in a directory that can be watched.
	tempDir := tempfile.TempDir(t, "filecollector-run")
	defer tempDir.Cleanup()

	// Use a second directory for the files before they are place in the watch directory.
	scratchDir := tempfile.TempDir(t, "filecollector-run-scratch")
	defer scratchDir.Cleanup()

	originalFile := scratchDir.Write(t, SimpleYamlInput, "create")
	updatedFile := scratchDir.Write(t, YamlInput, "update")

	collector.file = createWatchedPath(tempDir, originalFile)

	result := make(chan error, 1)
	go func() {
		result <- collector.Run()
	}()

	<-syncCh // wait for reload

	updateWatchedPath(collector.file, updatedFile)

	timer := time.NewTimer(100 * time.Millisecond)
	for {
		select {
		case <-syncCh:
			// detected changed file and updated
			updater.StopLoop()
			assert.Nil(t, <-result)
			return
		case <-timer.C:
			// handle case where we updated the file before the watcher was started
			// by updating the file with a comment
			updateWatchedPath(collector.file, tempDir.Write(t, YamlInput))
			timer.Reset(100 * time.Millisecond)
		}
	}
}

func TestFileCollectorRunErrors(t *testing.T) {
	collector, ctrl, mockUpdater := makeFileCollectorAndMock(t)
	defer ctrl.Finish()

	mockUpdater.EXPECT().Close()

	// Need to put temp file in a director that can be watched.
	tempDir := tempfile.TempDir(t, "filecollector-run")
	defer tempDir.Cleanup()

	collector.file = tempDir.Write(t, "this is not yaml, my dude")
	assert.NonNil(t, collector.Run())
}
