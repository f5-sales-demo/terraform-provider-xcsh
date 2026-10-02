---
page_title: "create_cloud_hosted"
subcategory: ""
description: "F5 Cloud Hosted."
xcsh_docs: {"aliases": ["create cloud hosted"], "body_bytes": 2486, "body_sha256": "sha256:25c3fc8ac2fd7b6ca27794c785cb74a5d6917192a82143d40ac5dadc163a1c7f", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted", "parent_id": "xcsh-docs:resources:bot_infrastructure:reference", "path": "documentation/resources/bot_infrastructure/properties/create_cloud_hosted/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2330310010102100-2003222100001321-3123220220113030-3220331102211303-0232000201002002-2321233202221132-3101030133100233-1003211122120022", "registry_path": "docs/guides/resources--bot_infrastructure--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "create_cloud_hosted:ConflictingObjectAttributes:production,testing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "create_cloud_hosted:ConflictingObjectAttributes:production,testing", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["create_cloud_hosted"], "schema_version": 1, "sections": [{"aliases": ["ip addresses"], "anchor": "schema-create_cloud_hosted--ip_addresses", "description": "Only traffic from these IP addresses is allowed to access this Bot Defense infrastructure.", "document_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["create_cloud_hosted", "ip_addresses"], "syntax": "attribute", "type": "list"}, {"aliases": ["production"], "anchor": "section", "description": "Production.", "document_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-create_cloud_hosted--production--region_1", "enforcement": "provider-schema", "group": "create_cloud_hosted.production:RequiredObjectAttributes:region_1,region_2", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "type": "requires"}, {"anchor": "schema-create_cloud_hosted--production--region_2", "enforcement": "provider-schema", "group": "create_cloud_hosted.production:RequiredObjectAttributes:region_1,region_2", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:production", "type": "requires"}], "schema_path": ["create_cloud_hosted", "production"], "syntax": "block", "type": "object"}, {"aliases": ["testing"], "anchor": "section", "description": "Testing", "document_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-create_cloud_hosted--testing--region_1", "enforcement": "provider-schema", "group": "create_cloud_hosted.testing:RequiredObjectAttributes:region_1", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_infrastructure:properties:create_cloud_hosted:testing", "type": "requires"}], "schema_path": ["create_cloud_hosted", "testing"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_infrastructure/properties/create_cloud_hosted/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "F5 Cloud Hosted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [create_cloud_hosted.production](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/production/)
- [create_cloud_hosted.testing](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/create_cloud_hosted/testing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/)
- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/)
