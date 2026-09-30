---
page_title: "vn_config.global_network_list.global_network_connections"
subcategory: ""
description: "vn_config.global_network_list.global_network_connections for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 2690, "body_sha256": "sha256:425a97bea90a09fc431975100864053c4791c493c6ed68abfddbb435fb608dfd", "canonical_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:sli_to_global_dr", "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:slo_to_global_dr"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list", "path": "docs/guides/resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["vn_config", "global_network_list", "global_network_connections"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "vn_config.global_network_list.global_network_connections for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# vn_config.global_network_list.global_network_connections

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
- [Property reference](resources--aws_tgw_site--reference.md)
- [vn_config](resources--aws_tgw_site--properties--vn_config.md)
- [vn_config.global_network_list](resources--aws_tgw_site--properties--vn_config--global_network_list.md)
- vn_config.global_network_list.global_network_connections

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

- [sli_to_global_dr](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--sli_to_global_dr.md): complete subsection reference.

- [slo_to_global_dr](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--slo_to_global_dr.md): complete subsection reference.

## Next pages

- [vn_config.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--sli_to_global_dr.md)
- [vn_config.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_tgw_site--properties--vn_config--global_network_list--global_network_connections--slo_to_global_dr.md)
- [vn_config.global_network_list](resources--aws_tgw_site--properties--vn_config--global_network_list.md)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md)
