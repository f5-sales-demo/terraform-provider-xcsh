---
page_title: "xcsh_authentication"
subcategory: ""
description: "Manages a Authentication resource in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["authentication", "credential setup", "credentials"], "body_bytes": 1327, "body_sha256": "sha256:55b1fb0aa19b6037105f4b5e40048d13c05a1f70ef4c565af3c3f7055edb3205", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:authentication:reference", "xcsh-docs:data-sources:authentication:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:authentication:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/authentication/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3330131022001332-3012131003233201-0203110311220323-2121221200223033-3213320331001200-2020031013302300-0110311123332221-2120122213100103", "registry_path": "docs/data-sources/authentication.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authentication/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Authentication resource in F5 Distributed Cloud.", "tasks": ["authentication", "configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["authenticationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_authentication

Breadcrumbs:

- xcsh_authentication

Manages a Authentication resource in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Authentication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Authentication by name
data "xcsh_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"
}

output "authentication_id" {
  value = data.xcsh_authentication.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/examples/)
