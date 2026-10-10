---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_mitigated_domain."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1074, "body_sha256": "sha256:1b1503797070ac0e10c6b9b242f1d3f6941cb8b7800980b4c47e540b18c9cef2", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:mitigated_domain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d859b6b9d3a8006d5c27b6cad66087687ae1b57c2b118afe91822b84dab15960", "source_path": "examples/resources/xcsh_mitigated_domain/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:mitigated_domain:example:resource", "parent_id": "xcsh-docs:resources:mitigated_domain:examples", "path": "documentation/resources/mitigated_domain/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "mitigated_domain", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3111222322112310-2012213112313130-2021210212310210-3332020320110021-3331100211312203-0323110130303233-0030112132130100-2100201112000321", "registry_path": "docs/guides/resources--mitigated_domain--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/mitigated_domain/examples/resource/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Resource for xcsh_mitigated_domain.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["mitigated_domainCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_mitigated_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_mitigated_domain/resource.tf`; digest `sha256:d859b6b9d3a8006d5c27b6cad66087687ae1b57c2b118afe91822b84dab15960`.

```terraform
# MitigatedDomain Resource Example
# Manages Mitigated Domain in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic MitigatedDomain configuration
resource "xcsh_mitigated_domain" "example" {
  name      = "example-mitigated-domain"
  namespace = "staging"

  mitigated_domain = "example-value"
}
```
