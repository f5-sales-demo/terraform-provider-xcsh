---
page_title: "blocked_services.web_user_interface"
subcategory: ""
description: "blocked_services.web_user_interface for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 970, "body_sha256": "sha256:6e2e23307c7478b4577fbc67715ecccc84a90bb34406b77305558051b269ae32", "canonical_id": "xcsh-docs:resources:fleet:properties:blocked_services:web_user_interface", "child_ids": [], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:blocked_services:web_user_interface", "parent_id": "xcsh-docs:resources:fleet:properties:blocked_services", "path": "docs/guides/resources--fleet--properties--blocked_services--web_user_interface.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["blocked_services", "web_user_interface"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/blocked_services/web_user_interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "blocked_services.web_user_interface for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# blocked_services.web_user_interface

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
- [blocked_services](resources--fleet--properties--blocked_services.md)
- blocked_services.web_user_interface

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
web_user_interface = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [blocked_services](resources--fleet--properties--blocked_services.md)
- [xcsh_fleet](../resources/fleet.md)
