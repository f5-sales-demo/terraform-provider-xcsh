---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_irule."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1212, "body_sha256": "sha256:1139d326e46cd4f509a283f40fd24b3e2f7bf4830df65ae9b36b00258bee300a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:irule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:42cb09305e24144e92e6cea8fc9300e3823e0b03a09b749cb4b4a6120a2dca4b", "source_path": "examples/resources/xcsh_irule/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:irule:example:resource", "parent_id": "xcsh-docs:resources:irule:examples", "path": "documentation/resources/irule/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "irule", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0121033310322201-2113232211001202-2212002230230323-1023202203302203-0131202322320320-0022210011030332-2011130333023221-3113320313120331", "registry_path": "docs/guides/resources--irule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/irule/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_irule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["iruleCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_irule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_irule/resource.tf`; digest `sha256:42cb09305e24144e92e6cea8fc9300e3823e0b03a09b749cb4b4a6120a2dca4b`.

```terraform
# Irule Resource Example
# Manages iRule in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Irule configuration
resource "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"

  description_spec = "example-value"
  irule            = "example-value"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/examples/)
- [xcsh_irule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/)
