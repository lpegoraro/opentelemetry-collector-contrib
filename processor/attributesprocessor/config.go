// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package attributesprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"

import (
	"errors"
	"go.opentelemetry.io/collector/config"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/attraction"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/processor/filterconfig"
)

// Config specifies the set of attributes to be inserted, updated, upserted and
// deleted and the properties to include/exclude a span from being processed.
// This processor handles all forms of modifications to attributes within a span, log, or metric.
// Prior to any actions being applied, each span is compared against
// the include properties and then the exclude properties if they are specified.
// This determines if a span is to be processed or not.
// The list of actions is applied in order specified in the configuration.
type Config struct {
	config.ProcessorSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct

	filterconfig.MatchConfig `mapstructure:",squash"`

	// Specifies the list of attributes to act on.
	// The set of actions are {INSERT, UPDATE, UPSERT, DELETE, HASH, EXTRACT}.
	// This is a required field.
	attraction.Settings `mapstructure:",squash"`
}

var _ config.Processor = (*Config)(nil)

// Validate checks if the processor configuration is valid
func (cfg *Config) Validate() error {
	return nil
}

func (cfg *Config) actionBuilder(key string, value interface{}, regexPattern, fromAttribute, fromContext, convertedType, action string) (newAction attraction.ActionKeyValue, err error) {
	if value == nil || fromAttribute == "" || fromContext == "" {
		return attraction.ActionKeyValue{}, errors.New("failed to create action, one value must be attributed")
	}
	return attraction.ActionKeyValue{
		Key:           key,
		Value:         value,
		RegexPattern:  regexPattern,
		FromAttribute: fromAttribute,
		FromContext:   fromContext,
		ConvertedType: convertedType,
		Action:        attraction.Action(action),
	}, nil
}

func (cfg *Config) AddInsertAction(key string, value interface{}, regexPattern, attributeMatcher, contextMatcher, convertedType string) (ok bool) {
	newAction, err := cfg.actionBuilder(key, value, regexPattern, attributeMatcher, contextMatcher, convertedType, "insert")
	if err != nil {
		return false
	}
	cfg.Actions = cfg.AppendAction(newAction)
	return true
}

func (cfg *Config) AppendAction(newAction attraction.ActionKeyValue) []attraction.ActionKeyValue {
	if cfg.Actions == nil {
		cfg.Actions = []attraction.ActionKeyValue{}
	}
	return append(cfg.Actions, newAction)
}

func (cfg *Config) AddDeleteAction(key string, value interface{}, regexPattern, attributeMatcher, contextMatcher, convertedType string) (ok bool) {
	newAction, err := cfg.actionBuilder(key, value, regexPattern, attributeMatcher, contextMatcher, convertedType, "delete")
	if err != nil {
		return false
	}
	cfg.Actions = append(cfg.Actions, newAction)
	return true
}

func (cfg *Config) AddUpdateAction(key string, value interface{}, regexPattern, attributeMatcher, contextMatcher, convertedType string) (ok bool) {
	newAction, err := cfg.actionBuilder(key, value, regexPattern, attributeMatcher, contextMatcher, convertedType, "update")
	if err != nil {
		return false
	}
	cfg.Actions = append(cfg.Actions, newAction)
	return true
}

func (cfg *Config) AddConvertAction(key string, value interface{}, regexPattern, attributeMatcher, contextMatcher, convertedType string) (ok bool) {
	newAction, err := cfg.actionBuilder(key, value, regexPattern, attributeMatcher, contextMatcher, convertedType, "update")
	if err != nil {
		return false
	}
	cfg.Actions = append(cfg.Actions, newAction)
	return true
}

func (cfg *Config) AddHashAction(key string, value interface{}, regexPattern, attributeMatcher, contextMatcher, convertedType string) (ok bool) {
	newAction, err := cfg.actionBuilder(key, value, regexPattern, attributeMatcher, contextMatcher, convertedType, "update")
	if err != nil {
		return false
	}
	cfg.Actions = append(cfg.Actions, newAction)
	return true
}
