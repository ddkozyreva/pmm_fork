// Copyright (C) 2023 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package server

import (
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	pmmapitests "github.com/percona/pmm/api-tests"
	advisorClient "github.com/percona/pmm/api/advisors/v1/json/client"
	advisor "github.com/percona/pmm/api/advisors/v1/json/client/advisor_service"
	serverClient "github.com/percona/pmm/api/server/v1/json/client"
	"github.com/percona/pmm/api/server/v1/json/client/server_service"
)

func TestStartChecks(t *testing.T) {
	t.Run("with advisors enabled", func(t *testing.T) {
		toggleAdvisorChecks(t, true)
		t.Cleanup(func() { restoreSettingsDefaults(t) })

		resp, err := advisorClient.Default.AdvisorService.StartAdvisorChecks(nil)
		require.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("with advisors disabled", func(t *testing.T) {
		toggleAdvisorChecks(t, false)
		t.Cleanup(func() { restoreSettingsDefaults(t) })

		resp, err := advisorClient.Default.AdvisorService.StartAdvisorChecks(nil)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.FailedPrecondition, `Advisor checks are disabled.`)
		assert.Nil(t, resp)
	})
}

func TestGetAdvisorCheckResults(t *testing.T) {
	t.Run("with disabled Advisors", func(t *testing.T) {
		toggleAdvisorChecks(t, false)
		t.Cleanup(func() { restoreSettingsDefaults(t) })

		results, err := advisorClient.Default.AdvisorService.GetFailedChecks(nil)
		pmmapitests.AssertAPIErrorf(t, err, 400, codes.FailedPrecondition, `Advisor checks are disabled.`)
		assert.Nil(t, results)
	})

	t.Run("with enabled Advisors", func(t *testing.T) {
		toggleAdvisorChecks(t, true)
		t.Cleanup(func() { restoreSettingsDefaults(t) })

		resp, err := advisorClient.Default.AdvisorService.StartAdvisorChecks(nil)
		require.NoError(t, err)
		assert.NotNil(t, resp)

		results, err := advisorClient.Default.AdvisorService.GetFailedChecks(nil)
		require.NoError(t, err)
		assert.NotNil(t, results)
	})
}

func TestListAdvisorChecks(t *testing.T) {
	toggleAdvisorChecks(t, true)
	t.Cleanup(func() { restoreSettingsDefaults(t) })

	resp, err := advisorClient.Default.AdvisorService.ListAdvisorChecks(nil)
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Payload.Checks)
	for _, c := range resp.Payload.Checks {
		assert.NotEmpty(t, c.Name, "%+v", c)
		assert.NotEmpty(t, c.Summary, "%+v", c)
		assert.NotEmpty(t, c.Description, "%+v", c)
	}
}

func TestListAdvisors(t *testing.T) {
	toggleAdvisorChecks(t, true)
	t.Cleanup(func() { restoreSettingsDefaults(t) })

	resp, err := advisorClient.Default.AdvisorService.ListAdvisors(nil)
	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Payload.Advisors)
	for _, a := range resp.Payload.Advisors {
		assert.NotEmpty(t, a.Name, "%+v", a)
		assert.NotEmpty(t, a.Summary, "%+v", a)
		assert.NotEmpty(t, a.Description, "%+v", a)
		assert.NotEmpty(t, a.Category, "%+v", a)
		assert.NotEmpty(t, a.Comment, "%+v", a)
		assert.NotEmpty(t, a.Checks, "%+v", a)

		for _, c := range a.Checks {
			assert.NotEmpty(t, c.Name, "%+v", c)
			assert.NotEmpty(t, c.Summary, "%+v", c)
			assert.NotEmpty(t, c.Description, "%+v", c)
		}
	}
}

func TestChangeAdvisorChecks(t *testing.T) {
	toggleAdvisorChecks(t, true)
	t.Cleanup(func() { restoreSettingsDefaults(t) })

	t.Run("enable disable", func(t *testing.T) {
		t.Run("enable disable", func(t *testing.T) {
			resp, err := advisorClient.Default.AdvisorService.ListAdvisorChecks(nil)
			require.NoError(t, err)
			require.NotEmpty(t, resp.Payload.Checks)

			var check *advisor.ListAdvisorChecksOKBodyChecksItems0

			// enable ⥁ disable loop, it checks current state of first returned check and changes its state,
			// then in second iteration it returns state to its origin.
			for i := 0; i < 2; i++ {
				check = resp.Payload.Checks[0]
				params := &advisor.ChangeAdvisorChecksParams{
					Body: advisor.ChangeAdvisorChecksBody{
						Params: []*advisor.ChangeAdvisorChecksParamsBodyParamsItems0{
							{
								Name:   check.Name,
								Enable: pointer.ToBool(!check.Enabled),
							},
						},
					},
					Context: pmmapitests.Context,
				}

				_, err = advisorClient.Default.AdvisorService.ChangeAdvisorChecks(params)
				require.NoError(t, err)

				resp, err = advisorClient.Default.AdvisorService.ListAdvisorChecks(nil)
				require.NoError(t, err)
				require.NotEmpty(t, resp.Payload.Checks)

				for _, c := range resp.Payload.Checks {
					if c.Name == check.Name {
						assert.NotEqual(t, check.Enabled, c.Enabled)
						break
					}
				}
			}
		})

		t.Run("unrecognized interval is ignored", func(t *testing.T) {
			t.Cleanup(func() { restoreCheckIntervalDefaults(t) })

			resp, err := advisorClient.Default.AdvisorService.ListAdvisorChecks(nil)
			require.NoError(t, err)
			require.NotEmpty(t, resp.Payload.Checks)

			check := resp.Payload.Checks[0]
			interval := *check.Interval
			params := &advisor.ChangeAdvisorChecksParams{
				Body: advisor.ChangeAdvisorChecksBody{
					Params: []*advisor.ChangeAdvisorChecksParamsBodyParamsItems0{
						{
							Name:     check.Name,
							Interval: pointer.ToString("unknown_interval"),
						},
					},
				},
				Context: pmmapitests.Context,
			}

			_, err = advisorClient.Default.AdvisorService.ChangeAdvisorChecks(params)
			require.NoError(t, err)

			resp, err = advisorClient.Default.AdvisorService.ListAdvisorChecks(nil)
			require.NoError(t, err)
			require.NotEmpty(t, resp.Payload.Checks)

			var found bool
			for _, c := range resp.Payload.Checks {
				if c.Name == check.Name {
					found = true
					assert.Equal(t, interval, *c.Interval)
				}
			}

			assert.True(t, found, "required check wasn't found")
		})

		t.Run("change interval normal", func(t *testing.T) {
			t.Cleanup(func() { restoreSettingsDefaults(t) })

			resp, err := advisorClient.Default.AdvisorService.ListAdvisorChecks(nil)
			require.NoError(t, err)
			require.NotEmpty(t, resp.Payload.Checks)

			// convert all checks to RARE interval
			pp := make([]*advisor.ChangeAdvisorChecksParamsBodyParamsItems0, len(resp.Payload.Checks))
			for i, check := range resp.Payload.Checks {
				pp[i] = &advisor.ChangeAdvisorChecksParamsBodyParamsItems0{
					Name:     check.Name,
					Interval: pointer.ToString(advisor.ChangeAdvisorChecksParamsBodyParamsItems0IntervalADVISORCHECKINTERVALRARE),
				}
			}

			params := &advisor.ChangeAdvisorChecksParams{
				Body:    advisor.ChangeAdvisorChecksBody{Params: pp},
				Context: pmmapitests.Context,
			}
			_, err = advisorClient.Default.AdvisorService.ChangeAdvisorChecks(params)
			require.NoError(t, err)

			resp, err = advisorClient.Default.AdvisorService.ListAdvisorChecks(nil)
			require.NoError(t, err)
			require.NotEmpty(t, resp.Payload.Checks)

			for _, check := range resp.Payload.Checks {
				assert.Equal(t, "ADVISOR_CHECK_INTERVAL_RARE", *check.Interval)
			}

			t.Run("intervals should be preserved on restart", func(t *testing.T) {
				resp, err := advisorClient.Default.AdvisorService.ListAdvisorChecks(nil)
				require.NoError(t, err)
				require.NotEmpty(t, resp.Payload.Checks)
				assert.Equal(t, "ADVISOR_CHECK_INTERVAL_RARE", *resp.Payload.Checks[0].Interval)
			})
		})
	})
}

func toggleAdvisorChecks(t *testing.T, enable bool) {
	t.Helper()

	res, err := serverClient.Default.ServerService.ChangeSettings(&server_service.ChangeSettingsParams{
		Body: server_service.ChangeSettingsBody{
			EnableAdvisor: pointer.ToBool(enable),
		},
		Context: pmmapitests.Context,
	})
	require.NoError(t, err)
	require.Equal(t, enable, res.Payload.Settings.AdvisorEnabled)

	if enable {
		// It takes some time to load check files
		time.Sleep(time.Second)
	}
}
