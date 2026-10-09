---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_external_connector."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1116, "body_sha256": "sha256:7a36fba1d36e2080b51041861b7e246785f7ff221e9e9f2fc68b190a54fc651e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e6d206bd3e4355ffab92fe542ac16d79135373b63491d27d87ac9997db3c291a", "source_path": "examples/resources/xcsh_external_connector/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:external_connector:example:resource", "parent_id": "xcsh-docs:resources:external_connector:examples", "path": "documentation/resources/external_connector/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1311210323213000-2020122012123210-0121232212103321-2100123111010301-2020231230132333-0000023203230331-3033131002100022-0013331232230331", "registry_path": "docs/guides/resources--external_connector--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/examples/resource/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Resource for xcsh_external_connector.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["external_connectorCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/external_connector/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_external_connector/resource.tf`; digest `sha256:e6d206bd3e4355ffab92fe542ac16d79135373b63491d27d87ac9997db3c291a`.

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
