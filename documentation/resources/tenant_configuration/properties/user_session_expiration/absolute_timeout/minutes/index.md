---
page_title: "user_session_expiration.absolute_timeout.minutes"
subcategory: ""
description: "Represents the session duration in minutes."
xcsh_docs: {"aliases": ["duration", "user session expiration absolute timeout minutes"], "body_bytes": 2539, "body_sha256": "sha256:61551f8fff66f9ed2bc6f6393ba45408f728e116ab8d88d1e140263aae5d27c5", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:minutes", "parent_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout", "path": "documentation/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/minutes/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2121112223201202-1330003023120300-2323300211223301-3001203311001012-1003200020001023-0202313310203112-0320132201113220-0012211331121033", "registry_path": "docs/guides/resources--tenant_configuration--reference--group-001.md", "relationships": [{"anchor": "schema-user_session_expiration--absolute_timeout--minutes--duration", "enforcement": "provider-schema", "group": "user_session_expiration.absolute_timeout.minutes:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:minutes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["user_session_expiration", "absolute_timeout", "minutes"], "schema_version": 1, "sections": [{"aliases": ["duration", "user session expiration absolute timeout minutes duration"], "anchor": "schema-user_session_expiration--absolute_timeout--minutes--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:minutes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_session_expiration", "absolute_timeout", "minutes", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/minutes/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Represents the session duration in minutes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_session_expiration.absolute_timeout.minutes

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/)
- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/)
- [user_session_expiration.absolute_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/)
- user_session_expiration.absolute_timeout.minutes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Represents the session duration in minutes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
minutes {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-user_session_expiration--absolute_timeout--minutes--duration"></a>

### duration property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(5, 43200),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 43200,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 5
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "43200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "43200"
  }
}
```
