---
page_title: "Resource"
subcategory: "Networking"
description: "Resource for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1256, "body_sha256": "sha256:8c6d09b162846389cf36c7130a939832fdb2c2145d8ed559354b815eda6c0994", "child_ids": [], "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:79950822068b0bbf291a51dad0cff745ae6eacc14d286990d529310031a03379", "source_path": "examples/resources/xcsh_network_connector/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:network_connector:example:resource", "parent_id": "xcsh-docs:resources:network_connector:examples", "path": "documentation/resources/network_connector/examples/resource/index.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_connector/resource.tf`; digest `sha256:79950822068b0bbf291a51dad0cff745ae6eacc14d286990d529310031a03379`.

```terraform
# NetworkConnector Resource Example
# Manages a Network Connector resource in F5 Distributed Cloud for network connector is created by users in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkConnector configuration
resource "xcsh_network_connector" "example" {
  name      = "example-network-connector"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/examples/)
- [xcsh_network_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/)
