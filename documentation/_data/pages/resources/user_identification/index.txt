---
page_title: "xcsh_user_identification"
subcategory: ""
description: "xcsh_user_identification for xcsh_user_identification."
xcsh_docs: {"aliases": [], "body_bytes": 1595, "body_sha256": "sha256:6e5979c0ff33dafa42e4c342d31a908d9a713345c10dfc139f013f69bb85256d", "child_ids": ["xcsh-docs:resources:user_identification:reference", "xcsh-docs:resources:user_identification:examples", "xcsh-docs:resources:user_identification:import", "xcsh-docs:resources:user_identification:timeouts"], "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:resources:user_identification:fundamentals", "parent_id": null, "path": "documentation/resources/user_identification/index.md", "provider_name": "user_identification", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_user_identification for xcsh_user_identification.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["user_identificationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_user_identification

Breadcrumbs:

- xcsh_user_identification

Manages user\_identification creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UserIdentification Resource Example
# Manages user_identification creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UserIdentification configuration
resource "xcsh_user_identification" "example" {
  name      = "example-user-identification"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/lifecycle/timeouts/)
