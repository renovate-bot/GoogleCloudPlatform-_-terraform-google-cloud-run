// Copyright 2025 Google LLC
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

package v2_with_gmp

import (
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/gcloud"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/tft"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestV2WithGMP(t *testing.T) {
	runV2GMP := tft.NewTFBlueprintTest(t)

	runV2GMP.DefineVerify(func(assert *assert.Assertions) {
		// runV2GMP.DefaultVerify(assert)

		projectID := runV2GMP.GetTFSetupStringOutput("project_id")
		serviceName := runV2GMP.GetStringOutput("service_name")
		serviceLocation := runV2GMP.GetStringOutput("service_location")

		runCmd := gcloud.Run(t, "run services describe", gcloud.WithCommonArgs([]string{serviceName, "--project", projectID, "--region", serviceLocation, "--format", "json"}))

		// Verify the Cloud Run Service deployed is in ready state.
		readyCondition := utils.GetFirstMatchResult(t, runCmd.Get("status").Get("conditions").Array(), "type", "Ready")
		assert.Equal("True", readyCondition.Get("status").String(), fmt.Sprintf("Should be in ready status"))
	})
	runV2GMP.Test()
}
