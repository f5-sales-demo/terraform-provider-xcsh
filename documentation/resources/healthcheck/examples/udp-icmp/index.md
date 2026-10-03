---
page_title: "Udp icmp"
subcategory: "Monitoring"
description: "Udp icmp for xcsh_healthcheck."
xcsh_docs: {"aliases": ["udp-icmp"], "body_bytes": 1318, "body_sha256": "sha256:8e31a316cd0cfa0a73e8c6f8df85a29a94141630ca1146b910cb3405119d4cba", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:a9f4239096a1306a82a2b1a05588ca6a7a9bf652a5283a4195352053169e6fb7", "source_path": "examples/resources/xcsh_healthcheck/udp-icmp.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:healthcheck:example:udp-icmp", "parent_id": "xcsh-docs:resources:healthcheck:examples", "path": "documentation/resources/healthcheck/examples/udp-icmp/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0202330033310230-2122330132020231-2232010012202201-3210330020023002-1131302202211123-3103323223131111-0000203121030110-0033112011013330", "registry_path": "docs/guides/resources--healthcheck--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["udp-icmp"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/examples/udp-icmp/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Udp icmp for xcsh_healthcheck.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["healthcheckCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Udp icmp

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/examples/)
- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
