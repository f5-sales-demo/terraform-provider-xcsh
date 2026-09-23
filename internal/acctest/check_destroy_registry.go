// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
)

type ResourceVerifier func(context.Context, *client.Client, string, string) error
type ResourceDeleter func(context.Context, *client.Client, string, string) error

// Only release-surface resources with conventional read/delete lifecycles
// belong here. Command-style resources have focused lifecycle tests instead.
var resourceVerifierRegistry = map[string]ResourceVerifier{
	"xcsh_bgp": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetBGP(ctx, ns, name)
		return err
	},
	"xcsh_dns_zone": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetDNSZone(ctx, ns, name)
		return err
	},
	"xcsh_external_connector": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetExternalConnector(ctx, ns, name)
		return err
	},
	"xcsh_http_loadbalancer": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetHTTPLoadBalancer(ctx, ns, name)
		return err
	},
	"xcsh_origin_pool": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetOriginPool(ctx, ns, name)
		return err
	},
	"xcsh_securemesh_site_v2": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetSecuremeshSiteV2(ctx, ns, name)
		return err
	},
	"xcsh_token": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetToken(ctx, ns, name)
		return err
	},
	"xcsh_virtual_site": func(ctx context.Context, c *client.Client, ns, name string) error {
		_, err := c.GetVirtualSite(ctx, ns, name)
		return err
	},
}

var resourceDeleterRegistry = map[string]ResourceDeleter{
	"xcsh_bgp": func(ctx context.Context, c *client.Client, ns, name string) error { return c.DeleteBGP(ctx, ns, name) },
	"xcsh_dns_zone": func(ctx context.Context, c *client.Client, ns, name string) error {
		return c.DeleteDNSZone(ctx, ns, name)
	},
	"xcsh_external_connector": func(ctx context.Context, c *client.Client, ns, name string) error {
		return c.DeleteExternalConnector(ctx, ns, name)
	},
	"xcsh_http_loadbalancer": func(ctx context.Context, c *client.Client, ns, name string) error {
		return c.DeleteHTTPLoadBalancer(ctx, ns, name)
	},
	"xcsh_origin_pool": func(ctx context.Context, c *client.Client, ns, name string) error {
		return c.DeleteOriginPool(ctx, ns, name)
	},
	"xcsh_securemesh_site_v2": func(ctx context.Context, c *client.Client, ns, name string) error {
		return c.DeleteSecuremeshSiteV2(ctx, ns, name)
	},
	"xcsh_token": func(ctx context.Context, c *client.Client, ns, name string) error {
		return c.DeleteToken(ctx, ns, name)
	},
	"xcsh_virtual_site": func(ctx context.Context, c *client.Client, ns, name string) error {
		return c.DeleteVirtualSite(ctx, ns, name)
	},
}

func CheckResourceDestroyedWithAPIVerification(resourceType string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		verifier, ok := resourceVerifierRegistry[resourceType]
		if !ok {
			return fmt.Errorf("no API verifier registered for release-surface resource type %q", resourceType)
		}
		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}
		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceType {
				continue
			}
			name := rs.Primary.Attributes["name"]
			if name == "" {
				name = rs.Primary.ID
			}
			ns := rs.Primary.Attributes["namespace"]
			if ns == "" {
				ns = "system"
			}
			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			err := waitForResourceDisappearance(ctx, c, verifier, ns, name, 6, 5*time.Second)
			cancel()
			if err != nil {
				return fmt.Errorf("verify deletion of %s %s/%s: %w", resourceType, ns, name, err)
			}
		}
		return nil
	}
}

func CheckResourceDisappears(resourceType, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found in state: %s", resourceName)
		}
		deleter, ok := resourceDeleterRegistry[resourceType]
		if !ok {
			return fmt.Errorf("no deleter registered for release-surface resource type %q", resourceType)
		}
		verifier, ok := resourceVerifierRegistry[resourceType]
		if !ok {
			return fmt.Errorf("no verifier registered for release-surface resource type %q", resourceType)
		}
		c, err := GetTestClient()
		if err != nil {
			return fmt.Errorf("failed to get test client: %w", err)
		}
		name, ns := rs.Primary.Attributes["name"], rs.Primary.Attributes["namespace"]
		if ns == "" {
			ns = "system"
		}
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		if err := retryIdempotentDelete(ctx, c, deleter, ns, name, 3, 5*time.Second); err != nil && !isNotFoundError(err) {
			return fmt.Errorf("delete %s %s/%s: %w", resourceType, ns, name, err)
		}
		if err := waitForResourceDisappearance(ctx, c, verifier, ns, name, 6, 3*time.Second); err != nil {
			return fmt.Errorf("confirm deletion of %s %s/%s: %w", resourceType, ns, name, err)
		}
		return nil
	}
}

func retryIdempotentDelete(ctx context.Context, c *client.Client, deleter ResourceDeleter, ns, name string, maxAttempts int, baseWait time.Duration) error {
	var lastErr error
	for attempt := 0; attempt < maxAttempts; attempt++ {
		lastErr = deleter(ctx, c, ns, name)
		if lastErr == nil {
			return nil
		}
		var apiErr *xcsherrors.XCSHError
		if !errors.As(lastErr, &apiErr) || !apiErr.IsRetryable() || attempt == maxAttempts-1 {
			return lastErr
		}
		if err := waitContext(ctx, baseWait*time.Duration(attempt+1)); err != nil {
			return err
		}
	}
	return lastErr
}

func waitForResourceDisappearance(ctx context.Context, c *client.Client, verifier ResourceVerifier, ns, name string, maxAttempts int, wait time.Duration) error {
	var lastErr error
	for attempt := 0; attempt < maxAttempts; attempt++ {
		lastErr = verifier(ctx, c, ns, name)
		if isNotFoundError(lastErr) {
			return nil
		}
		if lastErr != nil {
			var apiErr *xcsherrors.XCSHError
			if !errors.As(lastErr, &apiErr) || !apiErr.IsRetryable() {
				return lastErr
			}
		}
		if attempt < maxAttempts-1 {
			if err := waitContext(ctx, wait); err != nil {
				return err
			}
		}
	}
	if lastErr != nil {
		return lastErr
	}
	return errors.New("resource remained visible after deletion")
}

func waitContext(ctx context.Context, duration time.Duration) error {
	timer := time.NewTimer(duration)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func RegisterResourceVerifier(resourceType string, verifier ResourceVerifier) {
	resourceVerifierRegistry[resourceType] = verifier
}
func RegisterResourceDeleter(resourceType string, deleter ResourceDeleter) {
	resourceDeleterRegistry[resourceType] = deleter
}
func GetRegisteredResourceTypes() []string { return registryKeys(resourceVerifierRegistry) }
func GetRegisteredDeleterTypes() []string  { return registryKeys(resourceDeleterRegistry) }
func GetRegistrySize() int                 { return len(resourceVerifierRegistry) }
func GetDeleterRegistrySize() int          { return len(resourceDeleterRegistry) }
func registryKeys[T any](registry map[string]T) []string {
	keys := make([]string, 0, len(registry))
	for key := range registry {
		keys = append(keys, key)
	}
	return keys
}
