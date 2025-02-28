/*
 *  Copyright (c) 2018 Uber Technologies, Inc.
 *
 *     Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package astro

import (
	multierror "github.com/hashicorp/go-multierror"

	"github.com/romanwozniak/astro/astro/conf"
)

// Option is an option for the c that allows for changing of options or
// dependency injection for testing.
type Option func(*Project) error

func (c *Project) applyOptions(opts ...Option) (errs error) {
	for _, opt := range opts {
		if err := opt(c); err != nil {
			errs = multierror.Append(errs, err)
		}
	}
	return errs
}

// WithConfig allows you to pass project config.
func WithConfig(config conf.Project) Option {
	return func(c *Project) error {
		if err := config.Validate(); err != nil {
			return err
		}
		c.config = &config
		return nil
	}
}
