---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_alert_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1051, "body_sha256": "sha256:75435176dfc6c4f9139e4a31ded976de327d55c650d6a6f6e300128bd480351e", "canonical_id": "xcsh-docs:resources:alert_receiver:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7520b9716f7e91da22e687315c0f06dd552d7dfa1c32467bdbbdf9c32f5c2efb", "source_path": "examples/resources/xcsh_alert_receiver/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:alert_receiver:example:resource", "parent_id": "xcsh-docs:resources:alert_receiver:examples", "path": "docs/guides/resources--alert_receiver--example--resource.md", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_alert_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_alert_receiver](../resources/alert_receiver.md)
- [Examples](resources--alert_receiver--examples.md)
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

- [Examples](resources--alert_receiver--examples.md)
- [xcsh_alert_receiver](../resources/alert_receiver.md)
