---
page_title: "user_session_expiration.absolute_timeout.minutes"
subcategory: ""
description: "Represents the session duration in minutes."
xcsh_docs: {"aliases": ["duration", "user session expiration absolute timeout minutes"], "body_bytes": 2480, "body_sha256": "sha256:592365d1900b6151521a782395b48f1af79b2dc7565bdeb75a382d874681a19e", "capabilities": ["administration"], "category": "administration", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout:minutes", "parent_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout", "path": "documentation/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/minutes/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3002300121203211-3212221113130232-2201323110230322-0010312332030331-1122300312313303-3033032132121232-1212102200022020-0210212230021113", "registry_path": "docs/guides/data-sources--tenant_configuration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_session_expiration", "absolute_timeout", "minutes"], "schema_version": 1, "sections": [{"aliases": ["duration", "user session expiration absolute timeout minutes duration"], "anchor": "schema-user_session_expiration--absolute_timeout--minutes--duration", "description": "Configuration parameter for duration", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout:minutes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["user_session_expiration", "absolute_timeout", "minutes", "duration"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/minutes/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Represents the session duration in minutes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_session_expiration.absolute_timeout.minutes

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/)
- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/)
- [user_session_expiration.absolute_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/)
- user_session_expiration.absolute_timeout.minutes

<a id="section"></a>

Type: `"single"`. Computed.

Represents the session duration in minutes.

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

## Direct properties

<a id="schema-user_session_expiration--absolute_timeout--minutes--duration"></a>

### duration property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Next pages

- [user_session_expiration.absolute_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/)
- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/)
