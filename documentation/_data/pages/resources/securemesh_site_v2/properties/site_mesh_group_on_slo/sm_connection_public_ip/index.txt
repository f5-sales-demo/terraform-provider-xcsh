---
page_title: "site_mesh_group_on_slo.sm_connection_public_ip"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["site mesh group on slo sm connection public ip"], "body_bytes": 1075, "body_sha256": "sha256:3e16f6d51d34d06b3f42009c85e22cc78797be27edfb9feccff5cfb7020f1231", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo:sm_connection_public_ip", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:site_mesh_group_on_slo", "path": "documentation/resources/securemesh_site_v2/properties/site_mesh_group_on_slo/sm_connection_public_ip/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2231331302323330-0213212301333112-3033310021003321-0132133122112030-3031132203021330-2213303033233103-3022122303110220-1100123022011033", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_mesh_group_on_slo", "sm_connection_public_ip"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/site_mesh_group_on_slo/sm_connection_public_ip/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_mesh_group_on_slo.sm_connection_public_ip

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [site_mesh_group_on_slo](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/site_mesh_group_on_slo/)
- site_mesh_group_on_slo.sm_connection_public_ip

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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
sm_connection_public_ip = {}
```

This is an empty object or choice marker. It has no direct properties.
