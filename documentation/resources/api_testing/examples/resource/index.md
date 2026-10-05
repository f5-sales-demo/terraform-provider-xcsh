---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_api_testing."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1221, "body_sha256": "sha256:f1b2665737fc6a69c2857ea73e329dbc1deeb63e415a585deb1aa81f08875b72", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c1227be913c1052c8bc552ce9ebe6740f44bc74215b29a565808e69bd8cea2fb", "source_path": "examples/resources/xcsh_api_testing/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:api_testing:example:resource", "parent_id": "xcsh-docs:resources:api_testing:examples", "path": "documentation/resources/api_testing/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3321221100122300-1131211313201230-1111200101001123-3300032223033202-2131230212231332-2031330321223113-3200002103032101-2132332220323233", "registry_path": "docs/guides/resources--api_testing--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_api_testing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["api_testingCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_api_testing/resource.tf`; digest `sha256:c1227be913c1052c8bc552ce9ebe6740f44bc74215b29a565808e69bd8cea2fb`.

```terraform
# APITesting Resource Example
# Manages a API Testing resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APITesting configuration
resource "xcsh_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/examples/)
- [xcsh_api_testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_testing/)
