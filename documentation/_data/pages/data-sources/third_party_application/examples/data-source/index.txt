---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_third_party_application."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1420, "body_sha256": "sha256:5e9a6355869a41e86bcc7924a231765e600f52fb63535f9c2e165658ea2f7c95", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e3e758943d32635744f0c37b6cc645f6b182639afd8477f0071a1bfaa5279e53", "source_path": "examples/data-sources/xcsh_third_party_application/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:third_party_application:example:data-source", "parent_id": "xcsh-docs:data-sources:third_party_application:examples", "path": "documentation/data-sources/third_party_application/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0200101300201233-2010323323200232-1302213031311002-0003000231113013-2132310012222121-2002232120011103-0022033012313231-1122120330222102", "registry_path": "docs/guides/data-sources--third_party_application--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_third_party_application.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_third_party_application/data-source.tf`; digest `sha256:e3e758943d32635744f0c37b6cc645f6b182639afd8477f0071a1bfaa5279e53`.

```terraform
# ThirdPartyApplication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ThirdPartyApplication by name
data "xcsh_third_party_application" "example" {
  name      = "example-third-party-application"
  namespace = "staging"
}

output "third_party_application_id" {
  value = data.xcsh_third_party_application.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/examples/)
- [xcsh_third_party_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/)
