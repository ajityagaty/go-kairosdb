// Copyright 2016 Ajit Yagaty
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "fmt"
    "errors"
    "github.com/dan-compton/go-kairosdb/kairosdb"
	log "github.com/opsee/logrus"
	"github.com/hashicorp/go-multierror"
)

const (
    KdbQuery = "api/v1/datapoints/query"
)

// This is the type that implements the Client interface.
type KDBClient struct {
    kdbAddress string
}

// Returns a new http client
// TODO: reuse transports
func New(addr string) Client {
    return &KDBClient{
        kdbAddress: addr,
    }
}

// Queries the kdb server via the QueryMetricsRequest protobuf generate type
func (s *KDBClient) Query(in *kairosdb.QueryMetricsRequest) (*kairosdb.QueryMetricsResponse, error) {
    b, err := json.Marshal(in)
    if err != nil {
        return nil, err
    }

    hreq, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", s.kdbAddress, KdbQuery), bytes.NewBuffer(b))
    hreq.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(hreq)
    if err != nil {
        log.WithError(err).Fatalf("couldn't query kairosdb server")
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    gr := &kairosdb.QueryMetricsResponse{}
    err = json.Unmarshal(body, &gr)
    if err != nil {
        return nil, err
    }

    // handle errors field
    if len(gr.Errors) > 0 {
        var errs error
        for _, e := range gr.Errors {
            errs = multierror.Append(errs, errors.New(e))
        }
        return nil, errs
    }

    return gr, nil
}

