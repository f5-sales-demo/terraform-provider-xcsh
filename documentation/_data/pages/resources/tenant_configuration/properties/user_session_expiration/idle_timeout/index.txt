---
page_title: "user_session_expiration.idle_timeout"
subcategory: ""
description: "Represents the cookie expiration duration."
xcsh_docs: {"aliases": ["duration", "operation timeout", "user session expiration idle timeout"], "body_bytes": 2271, "body_sha256": "sha256:ce31779039c24464a739593e16614f0d515a7b165526a55a4798f95994412d91", "capabilities": ["administration"], "category": "administration", "child_ids": ["xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:hours", "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:minutes"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout", "parent_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration", "path": "documentation/resources/tenant_configuration/properties/user_session_expiration/idle_timeout/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2013211033012201-1303322322012201-0030031200023130-3301130101231210-0113013331113100-0310232130230220-1232000031122322-1301202130122312", "registry_path": "docs/guides/resources--tenant_configuration--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "user_session_expiration.idle_timeout:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:hours", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "user_session_expiration.idle_timeout:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:minutes", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["user_session_expiration", "idle_timeout"], "schema_version": 1, "sections": [{"aliases": ["hours"], "anchor": "section", "description": "Represents the cookie duration in hours.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:hours", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-user_session_expiration--idle_timeout--hours--duration", "enforcement": "provider-schema", "group": "user_session_expiration.idle_timeout.hours:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:hours", "type": "requires"}], "schema_path": ["user_session_expiration", "idle_timeout", "hours"], "syntax": "block", "type": "object"}, {"aliases": ["minutes"], "anchor": "section", "description": "Represents the cookie duration in minutes.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:minutes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-user_session_expiration--idle_timeout--minutes--duration", "enforcement": "provider-schema", "group": "user_session_expiration.idle_timeout.minutes:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:minutes", "type": "requires"}], "schema_path": ["user_session_expiration", "idle_timeout", "minutes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/properties/user_session_expiration/idle_timeout/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Represents the cookie expiration duration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_session_expiration.idle_timeout

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/)
- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/)
- user_session_expiration.idle_timeout

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Represents the cookie expiration duration.

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
idle_timeout {
  # Configure direct properties listed below.
}
```

## Direct properties

- [hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/idle_timeout/hours/): complete subsection reference.

- [minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/idle_timeout/minutes/): complete subsection reference.

## Next pages

- [user_session_expiration.idle_timeout.hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/idle_timeout/hours/)
- [user_session_expiration.idle_timeout.minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/idle_timeout/minutes/)
- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/)
- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
