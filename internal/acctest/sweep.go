// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

const (
	TestResourcePrefix = "tf-acc-test-"
	SweeperTimeout     = 5 * time.Minute
)

var sharedClient *client.Client

func GetSharedClient() (*client.Client, error) {
	if sharedClient != nil {
		return sharedClient, nil
	}
	apiURL := strings.TrimRight(os.Getenv(EnvXCSHURL), "/")
	if apiURL == "" {
		return nil, fmt.Errorf("%s must be set for sweepers", EnvXCSHURL)
	}
	if strings.HasSuffix(strings.ToLower(apiURL), "/api") {
		apiURL = apiURL[:len(apiURL)-4]
	}
	var err error
	switch DetectAuthMethod() {
	case AuthMethodP12:
		sharedClient, err = client.NewClientWithP12(apiURL, os.Getenv(EnvXCSHP12File), os.Getenv(EnvXCSHP12Password))
	case AuthMethodPEM:
		sharedClient, err = client.NewClientWithCert(apiURL, os.Getenv(EnvXCSHCert), os.Getenv(EnvXCSHKey), "")
	case AuthMethodToken:
		sharedClient = client.NewClient(apiURL, os.Getenv(EnvXCSHToken))
	default:
		return nil, fmt.Errorf("no authentication configured for sweepers")
	}
	if err != nil {
		return nil, fmt.Errorf("create sweeper client: %w", err)
	}
	return sharedClient, nil
}

func isTestResource(name string) bool { return strings.HasPrefix(name, TestResourcePrefix) }

// Only resources with generated list operations can be prefix-swept. All
// release-surface resources remain covered by per-test tracked cleanup.
func init() {
	resource.AddTestSweepers("xcsh_http_loadbalancer", &resource.Sweeper{Name: "xcsh_http_loadbalancer", F: sweepHTTPLoadbalancers})
	resource.AddTestSweepers("xcsh_origin_pool", &resource.Sweeper{Name: "xcsh_origin_pool", F: sweepOriginPools, Dependencies: []string{"xcsh_http_loadbalancer"}})
}

func sweepHTTPLoadbalancers(namespace string) error {
	return sweepListed(namespace, "http load balancer", func(ctx context.Context, c *client.Client, ns string) (*client.ListResponse, error) {
		return c.ListHTTPLoadBalancers(ctx, ns)
	}, func(ctx context.Context, c *client.Client, ns, name string) error {
		return c.DeleteHTTPLoadBalancer(ctx, ns, name)
	})
}

func sweepOriginPools(namespace string) error {
	return sweepListed(namespace, "origin pool", func(ctx context.Context, c *client.Client, ns string) (*client.ListResponse, error) {
		return c.ListOriginPools(ctx, ns)
	}, func(ctx context.Context, c *client.Client, ns, name string) error {
		return c.DeleteOriginPool(ctx, ns, name)
	})
}

func sweepListed(namespace, label string, list func(context.Context, *client.Client, string) (*client.ListResponse, error), deleteResource ResourceDeleter) error {
	if namespace == "" {
		namespace = "system"
	}
	c, err := GetSharedClient()
	if err != nil {
		return fmt.Errorf("get client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), SweeperTimeout)
	defer cancel()
	response, err := list(ctx, c, namespace)
	if err != nil {
		return fmt.Errorf("list %ss in %s: %w", label, namespace, err)
	}
	var failures []string
	for _, item := range response.Items {
		name := item.Metadata.Name
		if !isTestResource(name) {
			continue
		}
		WaitBeforeCleanup()
		if err := deleteResource(ctx, c, namespace, name); err != nil && !isNotFoundError(err) {
			failures = append(failures, fmt.Sprintf("%s: %v", name, err))
			continue
		}
		log.Printf("[INFO] swept %s %s/%s", label, namespace, name)
	}
	if len(failures) > 0 {
		return fmt.Errorf("sweep %ss failed: %s", label, strings.Join(failures, "; "))
	}
	return nil
}

func isNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	message := strings.ToLower(err.Error())
	return strings.Contains(message, "404") || strings.Contains(message, "not found") || strings.Contains(message, "not_found")
}
