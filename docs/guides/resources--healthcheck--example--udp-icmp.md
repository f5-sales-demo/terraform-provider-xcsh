---
page_title: "Udp icmp"
subcategory: "Monitoring"
description: "Udp icmp for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 1013, "body_sha256": "sha256:360ec6121ec8c46d2706f734a09ea46c2432529ad917cf419588b0d476cb37bb", "canonical_id": "xcsh-docs:resources:healthcheck:example:udp-icmp", "child_ids": [], "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:a9f4239096a1306a82a2b1a05588ca6a7a9bf652a5283a4195352053169e6fb7", "source_path": "examples/resources/xcsh_healthcheck/udp-icmp.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:healthcheck:example:udp-icmp", "parent_id": "xcsh-docs:resources:healthcheck:examples", "path": "docs/guides/resources--healthcheck--example--udp-icmp.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["udp-icmp"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/examples/udp-icmp/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Udp icmp for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Udp icmp

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md)
- [Examples](resources--healthcheck--examples.md)
- Udp icmp

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/udp-icmp.tf`; digest `sha256:a9f4239096a1306a82a2b1a05588ca6a7a9bf652a5283a4195352053169e6fb7`.

```terraform
# UdpIcmp — Acceptance-test-derived Configuration
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

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  udp_icmp_health_check = {}
}
```

## Next pages

- [Examples](resources--healthcheck--examples.md)
- [xcsh_healthcheck](../resources/healthcheck.md)
