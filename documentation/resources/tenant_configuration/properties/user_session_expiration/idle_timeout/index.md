---
page_title: "user_session_expiration.idle_timeout"
subcategory: ""
description: "Represents the cookie expiration duration."
xcsh_docs: {"aliases": ["duration", "user session expiration idle timeout"], "body_bytes": 1630, "body_sha256": "sha256:2e76168e0b40359a86c94ea801b549d75ff31944972d399e3b95ce8d884b30ff", "capabilities": ["administration"], "category": "administration", "child_ids": ["xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:hours", "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:minutes"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout", "parent_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration", "path": "documentation/resources/tenant_configuration/properties/user_session_expiration/idle_timeout/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2013211033012201-1303322322012201-0030031200023130-3301130101231210-0113013331113100-0310232130230220-1232000031122322-1301202130122312", "registry_path": "docs/guides/resources--tenant_configuration--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "user_session_expiration.idle_timeout:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:hours", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "user_session_expiration.idle_timeout:ConflictingObjectAttributes:hours,minutes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:minutes", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["user_session_expiration", "idle_timeout"], "schema_version": 1, "sections": [{"aliases": ["cookie expiration hours", "duration", "idle duration hours", "session idle timeout hours", "user session expiration idle timeout hours"], "anchor": "section", "description": "Specifies tenant user idle session expiration duration in hours (allowed range: 1 to 720 hours).", "document_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:hours", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-user_session_expiration--idle_timeout--hours--duration", "enforcement": "provider-schema", "group": "user_session_expiration.idle_timeout.hours:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:hours", "type": "requires"}], "schema_path": ["user_session_expiration", "idle_timeout", "hours"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "user session expiration idle timeout minutes"], "anchor": "section", "description": "Represents the cookie duration in minutes.", "document_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:minutes", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-user_session_expiration--idle_timeout--minutes--duration", "enforcement": "provider-schema", "group": "user_session_expiration.idle_timeout.minutes:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:idle_timeout:minutes", "type": "requires"}], "schema_path": ["user_session_expiration", "idle_timeout", "minutes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/properties/user_session_expiration/idle_timeout/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Represents the cookie expiration duration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
EnumExtractionComplete: false
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
