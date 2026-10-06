---
page_title: "create_cloud_hosted"
subcategory: ""
description: "F5 Cloud Hosted."
xcsh_docs: {"aliases": ["create cloud hosted"], "body_bytes": 1912, "body_sha256": "sha256:b66fcf4a538bb24545b4881454c9f210342aa2b21173646c41a68e72af6a3b2c", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted", "parent_id": "xcsh-docs:resources:bot_infrastructure:reference", "path": "documentation/resources/bot_infrastructure/properties/create_cloud_hosted/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2330310010102100-2003222100001321-3123220220113030-3220331102211303-0232000201002002-2321233202221132-3101030133100233-1003211122120022", "registry_path": "docs/guides/resources--bot_infrastructure--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "create_cloud_hosted:ConflictingObjectAttributes:production,testing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "create_cloud_hosted:ConflictingObjectAttributes:production,testing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["create_cloud_hosted"], "schema_version": 1, "sections": [{"aliases": ["create cloud hosted ip addresses"], "anchor": "schema-create_cloud_hosted--ip_addresses", "description": "Only traffic from these IP addresses is allowed to access this Bot Defense infrastructure.", "document_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["create_cloud_hosted", "ip_addresses"], "syntax": "attribute", "type": "list"}, {"aliases": ["create cloud hosted production"], "anchor": "section", "description": "Production.", "document_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-create_cloud_hosted--production--region_1", "enforcement": "provider-schema", "group": "create_cloud_hosted.production:RequiredObjectAttributes:region_1,region_2", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "type": "requires"}, {"anchor": "schema-create_cloud_hosted--production--region_2", "enforcement": "provider-schema", "group": "create_cloud_hosted.production:RequiredObjectAttributes:region_1,region_2", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "type": "requires"}], "schema_path": ["create_cloud_hosted", "production"], "syntax": "block", "type": "object"}, {"aliases": ["create cloud hosted testing"], "anchor": "section", "description": "Testing", "document_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-create_cloud_hosted--testing--region_1", "enforcement": "provider-schema", "group": "create_cloud_hosted.testing:RequiredObjectAttributes:region_1", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing", "type": "requires"}], "schema_path": ["create_cloud_hosted", "testing"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_infrastructure/properties/create_cloud_hosted/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "F5 Cloud Hosted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# create_cloud_hosted

Breadcrumbs:

- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/)
- create_cloud_hosted

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

F5 Cloud Hosted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("production",
    "testing")}
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
  "x-ves-oneof-field-type_choice": "[\"production\",\"testing\"]"
}
```

Terraform syntax:

```terraform
create_cloud_hosted {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-create_cloud_hosted--ip_addresses"></a>

### ip_addresses property

Type: `["list", "string"]`. Optional.

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

- [production](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/production/): complete subsection reference.

- [testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/testing/): complete subsection reference.
