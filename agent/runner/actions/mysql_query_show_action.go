// Copyright (C) 2023 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package actions

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/percona/pmm/agent/tlshelpers"
	"github.com/percona/pmm/api/agentpb"
	"github.com/percona/pmm/utils/sqlrows"
)

type mysqlQueryShowAction struct {
	id      string
	timeout time.Duration
	params  *agentpb.StartActionRequest_MySQLQueryShowParams
}

// NewMySQLQueryShowAction creates MySQL SHOW query Action.
func NewMySQLQueryShowAction(id string, timeout time.Duration, params *agentpb.StartActionRequest_MySQLQueryShowParams) Action {
	return &mysqlQueryShowAction{
		id:      id,
		timeout: timeout,
		params:  params,
	}
}

// ID returns an Action ID.
func (a *mysqlQueryShowAction) ID() string {
	return a.id
}

// Timeout returns Action timeout.
func (a *mysqlQueryShowAction) Timeout() time.Duration {
	return a.timeout
}

// Type returns an Action type.
func (a *mysqlQueryShowAction) Type() string {
	return "mysql-query-show"
}

// DSN returns a DSN for the Action.
func (a *mysqlQueryShowAction) DSN() string {
	return a.params.Dsn
}

// Run runs an Action and returns output and error.
func (a *mysqlQueryShowAction) Run(ctx context.Context) ([]byte, error) {
	db, err := mysqlOpen(a.params.Dsn, a.params.TlsFiles, a.params.TlsSkipVerify)
	if err != nil {
		return nil, err
	}
	defer db.Close() //nolint:errcheck
	defer tlshelpers.DeregisterMySQLCerts()

	// use prepared statement to force binary protocol usage that returns correct types
	stmt, err := db.PrepareContext(ctx, "SHOW /* pmm-agent */ "+a.params.Query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer stmt.Close() //nolint:errcheck

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	columns, dataRows, err := sqlrows.ReadRows(rows)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return agentpb.MarshalActionQuerySQLResult(columns, dataRows)
}

func (a *mysqlQueryShowAction) sealed() {}
