---
page_title: "xcsh_malicious_user_mitigation"
subcategory: ""
description: "xcsh_malicious_user_mitigation for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": [], "body_bytes": 1477, "body_sha256": "sha256:9568dfdcb815994dc1e55f684e10c5ff99d92be4afcee6900b15bfe3b2997bb1", "canonical_id": "xcsh-docs:resources:malicious_user_mitigation:fundamentals", "child_ids": ["xcsh-docs:resources:malicious_user_mitigation:reference", "xcsh-docs:resources:malicious_user_mitigation:examples", "xcsh-docs:resources:malicious_user_mitigation:import", "xcsh-docs:resources:malicious_user_mitigation:timeouts"], "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:resources:malicious_user_mitigation:fundamentals", "parent_id": null, "path": "docs/resources/malicious_user_mitigation.md", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_malicious_user_mitigation for xcsh_malicious_user_mitigation.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_malicious_user_mitigation

Breadcrumbs:

- xcsh_malicious_user_mitigation

Manages malicious\_user\_mitigation creates a new object in the storage backend for
metadata.namespace in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# MaliciousUserMitigation Resource Example
# Manages malicious_user_mitigation creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic MaliciousUserMitigation configuration
resource "xcsh_malicious_user_mitigation" "example" {
  name      = "example-malicious-user-mitigation"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--malicious_user_mitigation--reference.md)
- [Examples](../guides/resources--malicious_user_mitigation--examples.md)
- [Import](../guides/resources--malicious_user_mitigation--import.md)
- [Timeouts](../guides/resources--malicious_user_mitigation--timeouts.md)
