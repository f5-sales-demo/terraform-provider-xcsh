---
page_title: "params.ipsec.ipsec_psk"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["params ipsec ipsec psk"], "body_bytes": 1393, "body_sha256": "sha256:766eccbd5982afe9459c1fc952a8bdbb57563371c2c9e7bd40ec54f325a5f1f1", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk:blindfold_secret_info", "xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk:clear_secret_info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk", "parent_id": "xcsh-docs:data-sources:tunnel:properties:params:ipsec", "path": "documentation/data-sources/tunnel/properties/params/ipsec/ipsec_psk/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1022210110031320-3302000230230122-0221122220020322-3311033123210001-1101002303020330-3121102321331322-3013000332221322-1123330312311021", "registry_path": "docs/guides/data-sources--tunnel--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["params", "ipsec", "ipsec_psk"], "schema_version": 1, "sections": [{"aliases": ["params ipsec ipsec psk blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["params", "ipsec", "ipsec_psk", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["params ipsec ipsec psk clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk:clear_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["params", "ipsec", "ipsec_psk", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/params/ipsec/ipsec_psk/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["tunnelCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# params.ipsec.ipsec_psk

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/params/)
- [params.ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/params/ipsec/)
- params.ipsec.ipsec_psk

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/params/ipsec/ipsec_psk/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/params/ipsec/ipsec_psk/clear_secret_info/): complete subsection reference.
