---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_protocol_inspection."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1317, "body_sha256": "sha256:cb67bb0ce2976cbb6f67ac91d46efc9b5bdd82b7ee22472212baa956c9680dbc", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:protocol_inspection:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0edd9213402a7bd99b02d8b9c2f3a4ee63eb0305d083b63f69f415300d8acfa9", "source_path": "examples/resources/xcsh_protocol_inspection/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:protocol_inspection:example:resource", "parent_id": "xcsh-docs:resources:protocol_inspection:examples", "path": "documentation/resources/protocol_inspection/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "protocol_inspection", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0131201103312101-2302203022021031-0112010133131322-0222112131332302-2122303022102221-3133020111213221-3000100102231301-3131212011310100", "registry_path": "docs/guides/resources--protocol_inspection--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protocol_inspection/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_protocol_inspection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_inspectionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/examples/)
- [xcsh_protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protocol_inspection/)
