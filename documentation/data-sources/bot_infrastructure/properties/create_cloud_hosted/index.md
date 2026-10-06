---
page_title: "create_cloud_hosted"
subcategory: ""
description: "F5 Cloud Hosted."
xcsh_docs: {"aliases": ["create cloud hosted"], "body_bytes": 1629, "body_sha256": "sha256:f506719161a511ad608cdf65d7d5b6ee1b7f467948154026deb67f3828473c3a", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted:production", "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted:testing"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted", "parent_id": "xcsh-docs:data-sources:bot_infrastructure:reference", "path": "documentation/data-sources/bot_infrastructure/properties/create_cloud_hosted/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1002221222001300-2101220322011111-0133210023323221-3223332231003001-3033003221230102-0212010220020122-0320131010300022-3310302003201121", "registry_path": "docs/guides/data-sources--bot_infrastructure--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["create_cloud_hosted"], "schema_version": 1, "sections": [{"aliases": ["create cloud hosted ip addresses"], "anchor": "schema-create_cloud_hosted--ip_addresses", "description": "Only traffic from these IP addresses is allowed to access this Bot Defense infrastructure.", "document_id": "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["create_cloud_hosted", "ip_addresses"], "syntax": "attribute", "type": "list"}, {"aliases": ["create cloud hosted production"], "anchor": "section", "description": "Production.", "document_id": "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted:production", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["create_cloud_hosted", "production"], "syntax": "attribute", "type": "object"}, {"aliases": ["create cloud hosted testing"], "anchor": "section", "description": "Testing", "document_id": "xcsh-docs:data-sources:bot_infrastructure:properties:create_cloud_hosted:testing", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["create_cloud_hosted", "testing"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_infrastructure/properties/create_cloud_hosted/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "F5 Cloud Hosted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# create_cloud_hosted

Breadcrumbs:

- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/)
- create_cloud_hosted

<a id="section"></a>

Type: `"single"`. Computed.

F5 Cloud Hosted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type_choice": "[\"production\",\"testing\"]"
}
```

## Direct properties

<a id="schema-create_cloud_hosted--ip_addresses"></a>

### ip_addresses property

Type: `["list", "string"]`. Computed.

Only traffic from these IP addresses is allowed to access this Bot Defense infrastructure.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

- [production](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/create_cloud_hosted/production/): complete subsection reference.

- [testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/create_cloud_hosted/testing/): complete subsection reference.
