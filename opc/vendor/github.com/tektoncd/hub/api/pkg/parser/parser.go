// Copyright © 2020 The Tekton Authors.
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

package parser

import (
	"fmt"
	"sync"

	"github.com/tektoncd/hub/api/pkg/git"
	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

type Parser interface {
	Parse() ([]Resource, error)
}

func registerSchema() {
	tektonSchemes := runtime.NewSchemeBuilder(v1beta1.AddToScheme, v1.AddToScheme)
	if err := tektonSchemes.AddToScheme(scheme.Scheme); err != nil {
		fmt.Printf("Failed to apply all stored functions to the scheme %v", err)
	}
}

var once sync.Once

func ForCatalog(logger *zap.SugaredLogger, repo git.Repo, contextPath string) *CatalogParser {
	once.Do(registerSchema)
	return &CatalogParser{
		logger:      logger.With("component", "parser"),
		repo:        repo,
		contextPath: contextPath,
	}
}
