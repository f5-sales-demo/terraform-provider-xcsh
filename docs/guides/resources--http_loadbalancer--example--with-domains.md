---
page_title: "With domains"
subcategory: "Load Balancing"
description: "With domains for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1117, "body_sha256": "sha256:cf716084fb074fd8ab7957f441acedc017019fe1653f3c9400f890cbc7f9b665", "canonical_id": "xcsh-docs:resources:http_loadbalancer:example:with-domains", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:1a4378104ebc071c5ead72e6964e63ea12d8aa818fd0da9d4fb4731e46b0ab68", "source_path": "examples/resources/xcsh_http_loadbalancer/with-domains.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:http_loadbalancer:example:with-domains", "parent_id": "xcsh-docs:resources:http_loadbalancer:examples", "path": "docs/guides/resources--http_loadbalancer--example--with-domains.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["with-domains"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/examples/with-domains/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With domains for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# With domains

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Examples](resources--http_loadbalancer--examples.md)
- With domains

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/with-domains.tf`; digest `sha256:1a4378104ebc071c5ead72e6964e63ea12d8aa818fd0da9d4fb4731e46b0ab68`.

```terraform
# WithDomains — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "system"

  labels = {
    environment = "test"
  }

  domains = [
    "app.example.com",
    "api.example.com"
  ]

  http {
    port = 80
  }

  advertise_on_public_default_vip = {}
}
```

## Next pages

- [Examples](resources--http_loadbalancer--examples.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
