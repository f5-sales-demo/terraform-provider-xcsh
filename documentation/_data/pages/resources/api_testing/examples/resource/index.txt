---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_api_testing."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1221, "body_sha256": "sha256:f1b2665737fc6a69c2857ea73e329dbc1deeb63e415a585deb1aa81f08875b72", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:api_testing:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c1227be913c1052c8bc552ce9ebe6740f44bc74215b29a565808e69bd8cea2fb", "source_path": "examples/resources/xcsh_api_testing/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:api_testing:example:resource", "parent_id": "xcsh-docs:resources:api_testing:examples", "path": "documentation/resources/api_testing/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "api_testing", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-3321221100122300-1131211313201230-1111200101001123-3300032223033202-2131230212231332-2031330321223113-3200002103032101-2132332220323233", "registry_path": "docs/guides/resources--api_testing--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_testing/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_api_testing.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["api_testingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
