---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_protocol_inspection."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1074, "body_sha256": "sha256:0b67ff4f5b8c1bb7cc7a90315914f389ceeb7fe87fd9eff9e1fcfd6bbf16780c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protocol_inspection:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0edd9213402a7bd99b02d8b9c2f3a4ee63eb0305d083b63f69f415300d8acfa9", "source_path": "examples/resources/xcsh_protocol_inspection/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:protocol_inspection:example:resource", "parent_id": "xcsh-docs:resources:protocol_inspection:examples", "path": "documentation/resources/protocol_inspection/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0131201103312101-2302203022021031-0112010133131322-0222112131332302-2122303022102221-3133020111213221-3000100102231301-3131212011310100", "registry_path": "docs/guides/resources--protocol_inspection--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_inspection/examples/resource/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Resource for xcsh_protocol_inspection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_protocol_inspection/resource.tf`; digest `sha256:0edd9213402a7bd99b02d8b9c2f3a4ee63eb0305d083b63f69f415300d8acfa9`.

```terraform
# ProtocolInspection Resource Example
# Manages Protocol Inspection Specification in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtocolInspection configuration
resource "xcsh_protocol_inspection" "example" {
  name      = "example-protocol-inspection"
  namespace = "staging"
}
```
