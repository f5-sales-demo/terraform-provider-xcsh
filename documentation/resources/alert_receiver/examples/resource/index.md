---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_alert_receiver."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1257, "body_sha256": "sha256:ee45c0c7c8f6ac5bef156b955650632964d7e2a27765bb0293d8b50af327205e", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7520b9716f7e91da22e687315c0f06dd552d7dfa1c32467bdbbdf9c32f5c2efb", "source_path": "examples/resources/xcsh_alert_receiver/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:alert_receiver:example:resource", "parent_id": "xcsh-docs:resources:alert_receiver:examples", "path": "documentation/resources/alert_receiver/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0302211312132003-2302010323020101-0131303320113013-1001122332121003-1012202002120331-1200011020020323-3122133110000031-2221032233231212", "registry_path": "docs/guides/resources--alert_receiver--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_alert_receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_alert_receiver/resource.tf`; digest `sha256:7520b9716f7e91da22e687315c0f06dd552d7dfa1c32467bdbbdf9c32f5c2efb`.

```terraform
# AlertReceiver Resource Example
# Manages new Alert Receiver object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertReceiver configuration
resource "xcsh_alert_receiver" "example" {
  name      = "example-alert-receiver"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/examples/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
