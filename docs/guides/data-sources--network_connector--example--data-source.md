---
page_title: "Data source"
subcategory: "Networking"
description: "Data source for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1138, "body_sha256": "sha256:599fb7beb237b35a98e73815cb279dab9e42eca0d11721a48d1bcd19a8f8b2b1", "canonical_id": "xcsh-docs:data-sources:network_connector:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:842679078281a092ba5ea7fed221d6b2c452c242a5d7cc3a8d19f42931b95be3", "source_path": "examples/data-sources/xcsh_network_connector/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_connector:example:data-source", "parent_id": "xcsh-docs:data-sources:network_connector:examples", "path": "docs/guides/data-sources--network_connector--example--data-source.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md)
- [Examples](data-sources--network_connector--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_connector/data-source.tf`; digest `sha256:842679078281a092ba5ea7fed221d6b2c452c242a5d7cc3a8d19f42931b95be3`.

```terraform
# NetworkConnector Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkConnector by name
data "xcsh_network_connector" "example" {
  name      = "example-network-connector"
  namespace = "staging"
}

output "network_connector_id" {
  value = data.xcsh_network_connector.example.id
}
```

## Next pages

- [Examples](data-sources--network_connector--examples.md)
- [xcsh_network_connector](../data-sources/network_connector.md)
