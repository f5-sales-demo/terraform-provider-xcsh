---
page_title: "user_session_expiration.absolute_timeout"
subcategory: ""
description: "Represents the session expiration duration."
xcsh_docs: {"aliases": ["duration", "operation timeout", "user session expiration absolute timeout"], "body_bytes": 2308, "body_sha256": "sha256:7f292e969b351c2c7c3b8766e93043dbffdb6722062d53b10b2329ad0b122bf4", "capabilities": ["administration"], "category": "administration", "child_ids": ["xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:hours", "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:minutes"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout", "parent_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration", "path": "documentation/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0131112302222330-0332200030123212-0100031133203323-0333330313033111-2211333123210323-0231123233100022-3202031211002022-2310113122310211", "registry_path": "docs/guides/resources--tenant_configuration--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "user_session_expiration.absolute_timeout:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:hours", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "user_session_expiration.absolute_timeout:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:minutes", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["user_session_expiration", "absolute_timeout"], "schema_version": 1, "sections": [{"aliases": ["hours"], "anchor": "section", "description": "Represents the session duration in hours.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:hours", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-user_session_expiration--absolute_timeout--hours--duration", "enforcement": "provider-schema", "group": "user_session_expiration.absolute_timeout.hours:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:hours", "type": "requires"}], "schema_path": ["user_session_expiration", "absolute_timeout", "hours"], "syntax": "block", "type": "object"}, {"aliases": ["minutes"], "anchor": "section", "description": "Represents the session duration in minutes.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:minutes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-user_session_expiration--absolute_timeout--minutes--duration", "enforcement": "provider-schema", "group": "user_session_expiration.absolute_timeout.minutes:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:minutes", "type": "requires"}], "schema_path": ["user_session_expiration", "absolute_timeout", "minutes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Represents the session expiration duration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_session_expiration.absolute_timeout

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/)
- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/)
- user_session_expiration.absolute_timeout

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Represents the session expiration duration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("hours",
    "minutes")}
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
  "x-ves-oneof-field-unit_of_time": "[\"hours\",\"minutes\"]"
}
```

Terraform syntax:

```terraform
absolute_timeout {
  # Configure direct properties listed below.
}
```

## Direct properties

- [hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/hours/): complete subsection reference.

- [minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/minutes/): complete subsection reference.

## Next pages

- [user_session_expiration.absolute_timeout.hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/hours/)
- [user_session_expiration.absolute_timeout.minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/minutes/)
- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/)
- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
