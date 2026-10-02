---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_external_connector."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1357, "body_sha256": "sha256:a9550cfdb184b6bc089927ae4acaa762aadc0373119d3c4c3013dba273076354", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8660d99e20bf0369dc735e047075e9bae757c75b8856770f0f1f3584ea3dc187", "source_path": "examples/data-sources/xcsh_external_connector/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:external_connector:example:data-source", "parent_id": "xcsh-docs:data-sources:external_connector:examples", "path": "documentation/data-sources/external_connector/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0213100032310001-2023013202121123-2101132021002000-1201312213320131-0203100102021030-0200133133320113-0102210020102300-1113312122112103", "registry_path": "docs/guides/data-sources--external_connector--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_external_connector.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/examples/)
- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
