---
page_title: "xcsh_external_connector"
subcategory: ""
description: "Manages a External Connector resource in F5 Distributed Cloud for external_connector configuration specification. configuration."
xcsh_docs: {"aliases": ["external connector"], "body_bytes": 1687, "body_sha256": "sha256:958fd961b28237f9aa9445b2086e193e9408284bd0cfef5906be80e110185c9e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:external_connector:reference", "xcsh-docs:resources:external_connector:examples", "xcsh-docs:resources:external_connector:import", "xcsh-docs:resources:external_connector:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/external_connector/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222", "registry_path": "docs/resources/external_connector.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a External Connector resource in F5 Distributed Cloud for external_connector configuration specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
