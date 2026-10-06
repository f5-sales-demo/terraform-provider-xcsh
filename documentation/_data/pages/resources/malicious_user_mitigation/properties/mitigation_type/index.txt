---
page_title: "mitigation_type"
subcategory: ""
description: "Settings that specify the actions to be taken when malicious users are determined to be at different threat levels. User's activity is monitored and continuously analyzed for malicious behavior. From this analysis, a threat-level is assigned to each user. The settings defined in malicious user mitigation specify what"
xcsh_docs: {"aliases": ["mitigation type"], "body_bytes": 1562, "body_sha256": "sha256:81db53208b9c648623dd6790148ce1fe96086482aaaf216348d7a4c114dd59f4", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:reference", "path": "documentation/resources/malicious_user_mitigation/properties/mitigation_type/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0120200313002001-2220011100300123-2211210022323122-3031231321312001-2101212221230203-1133011303122233-1121112112303313-3011200201333230", "registry_path": "docs/guides/resources--malicious_user_mitigation--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "mitigation_type:RequiredObjectAttributes:rules", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["mitigation_type"], "schema_version": 1, "sections": [{"aliases": ["mitigation type rules"], "anchor": "section", "description": "Define the threat levels and the corresponding mitigation actions to be taken.", "document_id": "xcsh-docs:resources:malicious_user_mitigation:properties:mitigation_type:rules", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["mitigation_type", "rules"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/properties/mitigation_type/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Settings that specify the actions to be taken when malicious users are determined to be at different threat levels. User's activity is monitored and continuously analyzed for malicious behavior. From this analysis, a threat-level is assigned to each user. The settings defined in malicious user mitigation specify what", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# mitigation_type

Breadcrumbs:

- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/)
- mitigation_type

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Settings that specify the actions to be taken when malicious users are determined to be at different
threat levels. User's activity is monitored and continuously analyzed for malicious behavior. From
this analysis, a threat-level is assigned to each user. Server applies default when omitted.

Additional upstream details:

The settings defined in malicious user mitigation specify what mitigation actions to take for user
determined to be at different threat levels.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
```

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
mitigation_type {
  # Configure direct properties listed below.
}
```

## Direct properties

- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/properties/mitigation_type/rules/): complete subsection reference.
