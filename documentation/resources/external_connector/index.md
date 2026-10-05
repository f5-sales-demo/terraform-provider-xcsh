---
page_title: "xcsh_external_connector"
subcategory: ""
description: "Manages a External Connector resource in F5 Distributed Cloud for external_connector configuration specification. configuration."
xcsh_docs: {"aliases": ["external connector"], "body_bytes": 1687, "body_sha256": "sha256:958fd961b28237f9aa9445b2086e193e9408284bd0cfef5906be80e110185c9e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:external_connector:reference", "xcsh-docs:resources:external_connector:examples", "xcsh-docs:resources:external_connector:import", "xcsh-docs:resources:external_connector:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/external_connector/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222", "registry_path": "docs/resources/external_connector.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Manages a External Connector resource in F5 Distributed Cloud for external_connector configuration specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_external_connector

Breadcrumbs:

- xcsh_external_connector

Manages a External Connector resource in F5 Distributed Cloud for external\_connector configuration
specification. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ExternalConnector Resource Example
# Manages a External Connector resource in F5 Distributed Cloud for external_connector configuration specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ExternalConnector configuration
resource "xcsh_external_connector" "example" {
  name      = "example-external-connector"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/lifecycle/timeouts/)
