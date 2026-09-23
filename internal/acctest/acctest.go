// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

// Package acctest provides acceptance test utilities for F5 XC Terraform provider.
// Following HashiCorp's acceptance testing best practices.
package acctest

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/provider"
)

var (
	// testClient is a shared client for acceptance tests
	testClient     *client.Client
	testClientOnce sync.Once
	testClientErr  error
)

// Environment variable names for acceptance tests
// Using XCSH_* prefix for xcsh provider branding.
const (
	// EnvXCSHURL is the environment variable for the API URL
	EnvXCSHURL = "XCSH_API_URL"

	// EnvXCSHToken is the environment variable for the API token
	EnvXCSHToken = "XCSH_API_TOKEN"

	// EnvXCSHP12File is the environment variable for the P12 certificate file path
	EnvXCSHP12File = "XCSH_P12_FILE"

	// EnvXCSHP12Password is the environment variable for the P12 certificate password
	EnvXCSHP12Password = "XCSH_P12_PASSWORD" // pragma: allowlist secret

	// EnvXCSHCert is the environment variable for the PEM certificate file path
	EnvXCSHCert = "XCSH_CERT"

	// EnvXCSHKey is the environment variable for the PEM key file path
	EnvXCSHKey = "XCSH_KEY"

	// EnvXCSHTenantName is the environment variable for the tenant name
	EnvXCSHTenantName = "XCSH_TENANT_NAME"

	// EnvTFAccTest enables acceptance tests
	EnvTFAccTest = "TF_ACC"
)

// ProtoV6ProviderFactories returns the provider factories for acceptance testing
var ProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"xcsh": providerserver.NewProtocol6WithError(provider.New("test")()),
}

// ExternalProviders defines external providers used in acceptance tests
// Use this when tests require external providers like hashicorp/time
var ExternalProviders = map[string]resource.ExternalProvider{
	"time": {
		Source: "hashicorp/time",
	},
}

// AuthMethod represents the authentication method detected
type AuthMethod int

const (
	// AuthMethodNone indicates no authentication configured
	AuthMethodNone AuthMethod = iota
	// AuthMethodToken indicates API token authentication
	AuthMethodToken
	// AuthMethodP12 indicates P12 certificate authentication
	AuthMethodP12
	// AuthMethodPEM indicates PEM certificate authentication
	AuthMethodPEM
)

// DetectAuthMethod determines which authentication method is configured
func DetectAuthMethod() AuthMethod {
	// Check P12 authentication (preferred for testing)
	if os.Getenv(EnvXCSHP12File) != "" && os.Getenv(EnvXCSHP12Password) != "" {
		return AuthMethodP12
	}

	// Check PEM certificate authentication
	if os.Getenv(EnvXCSHCert) != "" && os.Getenv(EnvXCSHKey) != "" {
		return AuthMethodPEM
	}

	// Check token authentication
	if os.Getenv(EnvXCSHToken) != "" {
		return AuthMethodToken
	}

	return AuthMethodNone
}

// PreCheck validates that required environment variables are set before running tests.
// It also logs the test category as REAL_API for reporting purposes.
// When XCSH_MOCK_MODE is set, it automatically configures the environment to use
// the global mock server instead of requiring real credentials.
func PreCheck(t *testing.T) {
	t.Helper()

	// If mock mode is enabled, configure environment to use mock server
	// This must happen before any credential checks
	EnsureMockModeConfigured()

	// Log test category for reporting
	if IsMockMode() {
		LogTestCategory(t, TestCategoryMock)
		t.Logf("Using mock server mode (XCSH_MOCK_MODE=1)")
		return // Skip credential validation for mock mode
	}

	// Log test category for reporting - this is a real API test
	LogTestCategory(t, TestCategoryReal)

	// API URL is always required
	if os.Getenv(EnvXCSHURL) == "" {
		t.Fatalf("Required environment variable not set: %s", EnvXCSHURL)
	}

	// Check for at least one valid authentication method
	authMethod := DetectAuthMethod()

	switch authMethod {
	case AuthMethodP12:
		t.Logf("Using P12 certificate authentication (file: %s)", os.Getenv(EnvXCSHP12File))
	case AuthMethodPEM:
		t.Logf("Using PEM certificate authentication (cert: %s, key: %s)",
			os.Getenv(EnvXCSHCert), os.Getenv(EnvXCSHKey))
	case AuthMethodToken:
		t.Logf("Using API token authentication")
	case AuthMethodNone:
		t.Fatalf("No authentication configured. Set one of:\n"+
			"  - P12: %s and %s\n"+
			"  - PEM: %s and %s\n"+
			"  - Token: %s",
			EnvXCSHP12File, EnvXCSHP12Password,
			EnvXCSHCert, EnvXCSHKey,
			EnvXCSHToken)
	}

	precheckLiveCapability(t)
}

// SkipIfNotAccTest skips the test if TF_ACC is not set and mock mode is not enabled.
// When XCSH_MOCK_MODE is set, tests run without requiring TF_ACC.
func SkipIfNotAccTest(t *testing.T) {
	t.Helper()

	// Mock mode doesn't require TF_ACC
	if IsMockMode() {
		return
	}

	if os.Getenv(EnvTFAccTest) == "" {
		t.Skip("Acceptance tests skipped unless TF_ACC is set")
	}
}

// RandomName generates a random name with the given prefix for test resources
func RandomName(prefix string) string {
	canonicalBase := strings.TrimSuffix(TestResourcePrefix, "-")
	if prefix != canonicalBase && !strings.HasPrefix(prefix, TestResourcePrefix) {
		panic(fmt.Sprintf("acceptance-test resource prefix %q must begin with %q", prefix, TestResourcePrefix))
	}
	return fmt.Sprintf("%s-%s", prefix, acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum))
}

// TestNamespace returns the namespace for tests (default: "default")
func TestNamespace() string {
	if ns := os.Getenv("XCSH_DEFAULT_NAMESPACE"); ns != "" {
		return ns
	}
	return "default"
}

// ConfigCompose composes multiple Terraform configurations
func ConfigCompose(configs ...string) string {
	var sb strings.Builder
	for _, config := range configs {
		sb.WriteString(config)
		sb.WriteString("\n")
	}
	return sb.String()
}

// ProviderConfig returns the provider configuration for tests.
// Note: The terraform-plugin-testing framework with ProtoV6ProviderFactories
// handles provider registration automatically. No required_providers block
// is needed - the framework injects the provider via reattach configuration.
func ProviderConfig() string {
	return `
provider "xcsh" {
  # Configuration from environment variables
}
`
}

// CheckResourceExists returns a resource.TestCheckFunc that verifies a resource exists
func CheckResourceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource ID is not set: %s", resourceName)
		}

		return nil
	}
}

// CheckResourceDestroyed returns a resource.TestCheckFunc that verifies a resource is destroyed
func CheckResourceDestroyed(resourceType string) resource.TestCheckFunc {
	// Use the registry-based API verification when available
	// This delegates to CheckResourceDestroyedWithAPIVerification which:
	// - Performs actual API calls to verify resource deletion
	// - Includes retry logic for async deletion
	// - Falls back to state-only check for unregistered types
	return CheckResourceDestroyedWithAPIVerification(resourceType)
}

// CheckResourceAttr is a convenience wrapper around resource.TestCheckResourceAttr
func CheckResourceAttr(name, key, value string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttr(name, key, value)
}

// CheckResourceAttrSet is a convenience wrapper around resource.TestCheckResourceAttrSet
func CheckResourceAttrSet(name, key string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttrSet(name, key)
}

// CheckResourceAttrPair is a convenience wrapper around resource.TestCheckResourceAttrPair
func CheckResourceAttrPair(nameFirst, keyFirst, nameSecond, keySecond string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttrPair(nameFirst, keyFirst, nameSecond, keySecond)
}

// ImportStateVerify returns the import state verify settings
func ImportStateVerify(resourceName string) resource.TestStep {
	return resource.TestStep{
		ResourceName:      resourceName,
		ImportState:       true,
		ImportStateVerify: true,
	}
}

// ImportStateVerifyIgnore returns import state verify with ignored attributes
func ImportStateVerifyIgnore(resourceName string, ignoreFields ...string) resource.TestStep {
	return resource.TestStep{
		ResourceName:            resourceName,
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateVerifyIgnore: ignoreFields,
	}
}

// TestResource provides a base structure for resource tests
type TestResource struct {
	Name         string
	ResourceType string
	Namespace    string
}

// NewTestResource creates a new test resource with a random name
func NewTestResource(resourceType string) *TestResource {
	return &TestResource{
		Name:         RandomName("tf-test"),
		ResourceType: resourceType,
		Namespace:    TestNamespace(),
	}
}

// FullResourceName returns the full Terraform resource name
func (r *TestResource) FullResourceName() string {
	return fmt.Sprintf("xcsh_%s.test", r.ResourceType)
}

// IDAttribute returns the attribute path for the ID
func (r *TestResource) IDAttribute() string {
	return fmt.Sprintf("%s.id", r.FullResourceName())
}

// ErrorContains checks if an error contains a specific substring
func ErrorContains(t *testing.T, err error, substring string) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected error containing %q, got nil", substring)
	}

	if !strings.Contains(err.Error(), substring) {
		t.Fatalf("expected error containing %q, got: %s", substring, err.Error())
	}
}

// DefaultTestTimeout is the default timeout for test operations
const DefaultTestTimeout = 10 * time.Minute

// ContextWithTimeout returns a context with the standard test timeout
func ContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), DefaultTestTimeout)
}

// TestCheckFuncCompose composes multiple TestCheckFuncs into one
func TestCheckFuncCompose(funcs ...resource.TestCheckFunc) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(funcs...)
}

// GetTestClient returns a shared API client for acceptance tests.
// The client is created once and reused across all tests.
func GetTestClient() (*client.Client, error) {
	testClientOnce.Do(func() {
		apiURL := os.Getenv(EnvXCSHURL)
		if apiURL == "" {
			apiURL = "https://console.ves.volterra.io"
		}

		// Normalize URL (remove trailing slashes and /api suffix)
		apiURL = strings.TrimRight(apiURL, "/")
		if strings.HasSuffix(strings.ToLower(apiURL), "/api") {
			apiURL = apiURL[:len(apiURL)-4]
		}
		apiURL = strings.TrimRight(apiURL, "/")

		switch DetectAuthMethod() {
		case AuthMethodP12:
			testClient, testClientErr = client.NewClientWithP12(
				apiURL,
				os.Getenv(EnvXCSHP12File),
				os.Getenv(EnvXCSHP12Password),
			)
		case AuthMethodPEM:
			testClient, testClientErr = client.NewClientWithCert(
				apiURL,
				os.Getenv(EnvXCSHCert),
				os.Getenv(EnvXCSHKey),
				"", // CA cert optional
			)
		case AuthMethodToken:
			testClient = client.NewClient(apiURL, os.Getenv(EnvXCSHToken))
		default:
			testClientErr = fmt.Errorf("no authentication method configured")
		}
	})

	return testClient, testClientErr
}
