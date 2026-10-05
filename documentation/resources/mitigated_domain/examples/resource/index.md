---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_mitigated_domain."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1308, "body_sha256": "sha256:e41d3632d551f07c21197a86096427a4b64a8e765b323c29521070a07f12002e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:mitigated_domain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d859b6b9d3a8006d5c27b6cad66087687ae1b57c2b118afe91822b84dab15960", "source_path": "examples/resources/xcsh_mitigated_domain/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:mitigated_domain:example:resource", "parent_id": "xcsh-docs:resources:mitigated_domain:examples", "path": "documentation/resources/mitigated_domain/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "mitigated_domain", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3111222322112310-2012213112313130-2021210212310210-3332020320110021-3331100211312203-0323110130303233-0030112132130100-2100201112000321", "registry_path": "docs/guides/resources--mitigated_domain--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/mitigated_domain/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_mitigated_domain.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["mitigated_domainCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/examples/)
- [xcsh_mitigated_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/mitigated_domain/)
