package api_gateway

import (
	"github.com/cucumber/godog"
)

func initJwtServiceFallback(ctx *godog.ScenarioContext) {
	s, err := CreateScenarioWithRawAPIResource("istio-jwt-service-fallback.yaml", "istio-jwt-service-fallback")
	if err != nil {
		t.Fatalf("could not initialize scenario err=%s", err)
	}

	scenario := istioJwtManifestScenario{s}

	ctx.Step(`ServiceFallback: There is a httpbin service$`, scenario.thereIsAHttpbinService)
	ctx.Step(`ServiceFallback: There is an endpoint secured with JWT on path "([^"]*)" with service definition$`, scenario.thereIsAnEndpointWithServiceDefinition)
	ctx.Step(`ServiceFallback: There is an endpoint secured with JWT on path "([^"]*)"$`, scenario.thereIsAnJwtSecuredPath)
	ctx.Step(`ServiceFallback: The APIRule with service on root level is applied$`, scenario.theAPIRuleIsApplied)
	ctx.Step(`ServiceFallback: Calling the "([^"]*)" endpoint with a valid "([^"]*)" token should result in status between (\d+) and (\d+)$`, scenario.callingTheEndpointWithValidTokenShouldResultInStatusBetween)
}

func (s *istioJwtManifestScenario) thereIsAnEndpointWithServiceDefinition(path string) {
	s.manifestTemplate["jwtSecuredPathWithService"] = path
}
