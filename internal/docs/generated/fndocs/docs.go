// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "mdtogo"; DO NOT EDIT.
package fndocs

var FnShort = `Generate, transform, and validate configuration files.`
var FnLong = `
| Configuration Read From | Configuration Written To |
| ----------------------- | ------------------------ |
| local files or stdin    | local files or stdout    |

Functions are executables (that you can [write][Functions Developer Guide])
which accept a collection of Resource configuration as input, and emit a
collection of Resource configuration as output.

Functions can be packaged as container images, starlark scripts, or binary
executables.
`
var FnExamples = `
  # run the function defined by gcr.io/example.com/my-fn as a local container
  # against the configuration in DIR
  kpt fn run DIR/ --image gcr.io/example.com/my-fn

  # run the functions declared in files under FUNCTIONS_DIR/
  kpt fn run DIR/ --fn-path FUNCTIONS_DIR/

  # run the functions declared in files under DIR/
  kpt fn run DIR/
`

var ExportShort = `Auto-generating function pipelines for different workflow orchestrators`
var ExportLong = `
  kpt fn export DIR/ [--fn-path FUNCTIONS_DIR/] --workflow ORCHESTRATOR [--output OUTPUT_FILENAME]
  
  DIR:
    Path to a package directory.
  FUNCTIONS_DIR:
    Read functions from the directory instead of the DIR/.
  ORCHESTRATOR:
    Supported orchestrators are:
      - github-actions
      - cloud-build
      - gitlab-ci
      - jenkins
      - tekton
      - circleci
  OUTPUT_FILENAME:
    Specifies the filename of the generated pipeline. If omitted, the default
    output is stdout
`
var ExportExamples = `
  # read functions from DIR, run them against it as one step.
  # write the generated GitHub Actions pipeline to main.yaml.
  kpt fn export DIR/ --output main.yaml --workflow github-actions

  # discover functions in FUNCTIONS_DIR and run them against resource in DIR.
  # write the generated Cloud Build pipeline to stdout.
  kpt fn export DIR/ --fn-path FUNCTIONS_DIR/ --workflow cloud-build
`

var RunShort = `Locally execute one or more functions in containers`
var RunLong = `
  kpt fn run DIR [flags]

If the container exits with non-zero status code, run will fail and print the
container ` + "`" + `STDERR` + "`" + `.

  DIR:
    Path to a package directory.  Defaults to stdin if unspecified.
`
var RunExamples = `
  # read the Resources from DIR, provide them to a container my-fun as input,
  # write my-fn output back to DIR
  kpt fn run DIR/ --image gcr.io/example.com/my-fn

  # provide the my-fn with an input ConfigMap containing ` + "`" + `data: {foo: bar}` + "`" + `
  kpt fn run DIR/ --image gcr.io/example.com/my-fn:v1.0.0 -- foo=bar

  # run the functions in FUNCTIONS_DIR against the Resources in DIR
  kpt fn run DIR/ --fn-path FUNCTIONS_DIR/

  # discover functions in DIR and run them against Resource in DIR.
  # functions may be scoped to a subset of Resources -- see ` + "`" + `kpt help fn run` + "`" + `
  kpt fn run DIR/
`

var DocShort = `Locally execuate a function to get the help text`
var DocLong = `
  kpt fn doc --image=IMAGE

If the function supports --help, it will print the help text to STDOUT.
Otherwise, it will exit with non-zero exit code and print the error message to
STDERR.
`
var DocExamples = `
  kpt fn doc --image gcr.io/kpt-fn/set-namespace:unstable
`

var SinkShort = `Specify a directory as an output sink package`
var SinkLong = `
  kpt fn sink [DIR]
  
  DIR:
    Path to a package directory.  Defaults to stdout if unspecified.
`
var SinkExamples = `
  # run a function using explicit sources and sinks
  kpt fn source DIR/ |
    kpt fn run --image gcr.io/example.com/my-fn |
    kpt fn sink DIR/
`

var SourceShort = `Specify a directory as an input source package`
var SourceLong = `
  kpt fn source [DIR...]
  
  DIR:
    Path to a package directory.  Defaults to stdin if unspecified.
`
var SourceExamples = `
  # print to stdout configuration from DIR/ formatted as an input source
  kpt fn source DIR/

  # run a function using explicit sources and sinks
  kpt fn source DIR/ |
    kpt fn run --image gcr.io/example.com/my-fn |
    kpt fn sink DIR/
`
