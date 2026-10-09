---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_external_connector."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1111, "body_sha256": "sha256:7b8f4fbf6f2144e77d86290f29f3f089fe914a391b4a215b6d98aa200570d5f4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8660d99e20bf0369dc735e047075e9bae757c75b8856770f0f1f3584ea3dc187", "source_path": "examples/data-sources/xcsh_external_connector/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:external_connector:example:data-source", "parent_id": "xcsh-docs:data-sources:external_connector:examples", "path": "documentation/data-sources/external_connector/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0213100032310001-2023013202121123-2101132021002000-1201312213320131-0203100102021030-0200133133320113-0102210020102300-1113312122112103", "registry_path": "docs/guides/data-sources--external_connector--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_external_connector.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["external_connectorCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_external_connector/data-source.tf`; digest `sha256:8660d99e20bf0369dc735e047075e9bae757c75b8856770f0f1f3584ea3dc187`.

```terraform
# ExternalConnector Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ExternalConnector by name
data "xcsh_external_connector" "example" {
  name      = "example-external-connector"
  namespace = "staging"
}

output "external_connector_id" {
  value = data.xcsh_external_connector.example.id
}
```
