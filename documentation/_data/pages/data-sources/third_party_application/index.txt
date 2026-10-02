---
page_title: "xcsh_third_party_application"
subcategory: ""
description: "Manages a Third Party Application resource in F5 Distributed Cloud for third party application specification. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["third party application"], "body_bytes": 1503, "body_sha256": "sha256:f5be6d3b3eb7b3d9731485a57149f431e55eeed23bc034764a00f297708145e6", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:third_party_application:reference", "xcsh-docs:data-sources:third_party_application:examples"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/third_party_application/index.md", "product": "distributed-cloud", "provider_name": "third_party_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2320222231123022-0323332222212132-3300113320120123-0222011302331223-3222001221131320-2123313213312301-2022200110110030-3312032011311320", "registry_path": "docs/data-sources/third_party_application.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Third Party Application resource in F5 Distributed Cloud for third party application specification. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_third_party_application

Breadcrumbs:

- xcsh_third_party_application

Manages a Third Party Application resource in F5 Distributed Cloud for third party application
specification. configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/examples/)
