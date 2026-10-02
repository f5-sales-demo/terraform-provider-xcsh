---
page_title: "user_session_expiration.absolute_timeout.hours"
subcategory: ""
description: "Represents the session duration in hours."
xcsh_docs: {"aliases": ["duration", "operation timeout", "user session expiration absolute timeout hours"], "body_bytes": 2843, "body_sha256": "sha256:0f72eedc5c7660c4b94a541a367a0da9f9ed0c801be77451de55588acd4dc401", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:hours", "parent_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout", "path": "documentation/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/hours/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0330301320200331-2313011321301203-3313122002101001-1302223220323120-2213020111132231-1312031033032000-2333023000333322-1021102201003003", "registry_path": "docs/guides/resources--tenant_configuration--reference--group-001.md", "relationships": [{"anchor": "schema-user_session_expiration--absolute_timeout--hours--duration", "enforcement": "provider-schema", "group": "user_session_expiration.absolute_timeout.hours:RequiredObjectAttributes:duration", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:hours", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["user_session_expiration", "absolute_timeout", "hours"], "schema_version": 1, "sections": [{"aliases": ["duration"], "anchor": "schema-user_session_expiration--absolute_timeout--hours--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:resources:tenant_configuration:properties:user_session_expiration:absolute_timeout:hours", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_session_expiration", "absolute_timeout", "hours", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/hours/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Represents the session duration in hours.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_session_expiration.absolute_timeout.hours

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/)
- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/)
- [user_session_expiration.absolute_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/)
- user_session_expiration.absolute_timeout.hours

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Represents the session duration in hours.

Provider validators and defaults (from schema source):

```go
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
hours {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-user_session_expiration--absolute_timeout--hours--duration"></a>

### duration property

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 720),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 720,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "720"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "720"
  }
}
```

## Next pages

- [user_session_expiration.absolute_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/properties/user_session_expiration/absolute_timeout/)
- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tenant_configuration/)
