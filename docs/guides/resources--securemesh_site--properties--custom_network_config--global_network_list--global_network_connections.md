---
page_title: "custom_network_config.global_network_list.global_network_connections"
subcategory: ""
description: "custom_network_config.global_network_list.global_network_connections for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 2993, "body_sha256": "sha256:bbaf58f94f7cabc00ae79c4e31e9a526277e822a7a9ba9e01ee796831c9962e8", "canonical_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:global_network_list:global_network_connections", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:custom_network_config:global_network_list:global_network_connections:sli_to_global_dr", "xcsh-docs:resources:securemesh_site:properties:custom_network_config:global_network_list:global_network_connections:slo_to_global_dr"], "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:global_network_list:global_network_connections", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:global_network_list", "path": "docs/guides/resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["custom_network_config", "global_network_list", "global_network_connections"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/global_network_list/global_network_connections/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "custom_network_config.global_network_list.global_network_connections for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.global_network_list.global_network_connections

Breadcrumbs:

- [xcsh_securemesh_site](../resources/securemesh_site.md)
- [Property reference](resources--securemesh_site--reference.md)
- [custom_network_config](resources--securemesh_site--properties--custom_network_config.md)
- [custom_network_config.global_network_list](resources--securemesh_site--properties--custom_network_config--global_network_list.md)
- custom_network_config.global_network_list.global_network_connections

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
global_network_connections {
  # Configure direct properties listed below.
}
```

## Direct properties

- [sli_to_global_dr](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--sli_to_global_dr.md): complete subsection reference.

- [slo_to_global_dr](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--slo_to_global_dr.md): complete subsection reference.

## Next pages

- [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--sli_to_global_dr.md)
- [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr](resources--securemesh_site--properties--custom_network_config--global_network_list--global_network_connections--slo_to_global_dr.md)
- [custom_network_config.global_network_list](resources--securemesh_site--properties--custom_network_config--global_network_list.md)
- [xcsh_securemesh_site](../resources/securemesh_site.md)
