---
page_title: "params.ipsec.ipsec_psk"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["params ipsec ipsec psk"], "body_bytes": 2256, "body_sha256": "sha256:773d3b2fc29094b8a3b5a39c6f93f1c1ed2804eb663947215e705b6729b7586a", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk:blindfold_secret_info", "xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk", "parent_id": "xcsh-docs:resources:tunnel:properties:params:ipsec", "path": "documentation/resources/tunnel/properties/params/ipsec/ipsec_psk/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0120100121011131-0222212233122303-0033310323102030-0111122220023121-0012322203220132-3213012300321002-0300020001321011-2000312230323110", "registry_path": "docs/guides/resources--tunnel--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "params.ipsec.ipsec_psk:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "params.ipsec.ipsec_psk:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["params", "ipsec", "ipsec_psk"], "schema_version": 1, "sections": [{"aliases": ["params ipsec ipsec psk blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-params--ipsec--ipsec_psk--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "params.ipsec.ipsec_psk.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk:blindfold_secret_info", "type": "requires"}], "schema_path": ["params", "ipsec", "ipsec_psk", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["params ipsec ipsec psk clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-params--ipsec--ipsec_psk--clear_secret_info--url", "enforcement": "provider-schema", "group": "params.ipsec.ipsec_psk.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk:clear_secret_info", "type": "requires"}], "schema_path": ["params", "ipsec", "ipsec_psk", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/params/ipsec/ipsec_psk/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# params.ipsec.ipsec_psk

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/)
- [params.ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/)
- params.ipsec.ipsec_psk

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
ipsec_psk {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/ipsec_psk/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/ipsec_psk/clear_secret_info/): complete subsection reference.

## Next pages

- [params.ipsec.ipsec_psk.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/ipsec_psk/blindfold_secret_info/)
- [params.ipsec.ipsec_psk.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/ipsec_psk/clear_secret_info/)
- [params.ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/params/ipsec/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/)
