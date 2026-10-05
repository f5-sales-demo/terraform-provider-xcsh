---
page_title: "sumo_logic_receiver.url"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["sumo logic receiver url"], "body_bytes": 2325, "body_sha256": "sha256:061245a343b9439e916e8e064e4d3854540f86118cf42126287f04a1eda163c8", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url:blindfold_secret_info", "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url:clear_secret_info"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url", "parent_id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver", "path": "documentation/resources/global_log_receiver/properties/sumo_logic_receiver/url/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1211232133011100-3311220123122001-2231031332221030-2311333131201230-3030311013102130-1320012132103022-0212200323121201-1213232003112030", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "sumo_logic_receiver.url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "sumo_logic_receiver.url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["sumo_logic_receiver", "url"], "schema_version": 1, "sections": [{"aliases": ["sumo logic receiver url blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-sumo_logic_receiver--url--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "sumo_logic_receiver.url.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url:blindfold_secret_info", "type": "requires"}], "schema_path": ["sumo_logic_receiver", "url", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["sumo logic receiver url clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-sumo_logic_receiver--url--clear_secret_info--url", "enforcement": "provider-schema", "group": "sumo_logic_receiver.url.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:sumo_logic_receiver:url:clear_secret_info", "type": "requires"}], "schema_path": ["sumo_logic_receiver", "url", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/sumo_logic_receiver/url/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# sumo_logic_receiver.url

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- [sumo_logic_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/)
- sumo_logic_receiver.url

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
url {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/url/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/url/clear_secret_info/): complete subsection reference.

## Next pages

- [sumo_logic_receiver.url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/url/blindfold_secret_info/)
- [sumo_logic_receiver.url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/url/clear_secret_info/)
- [sumo_logic_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
