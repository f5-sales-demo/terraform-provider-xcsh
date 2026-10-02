---
page_title: "user_session_expiration.idle_timeout"
subcategory: ""
description: "Represents the cookie expiration duration."
xcsh_docs: {"aliases": ["duration", "operation timeout", "user session expiration idle timeout"], "body_bytes": 2015, "body_sha256": "sha256:4200adeffb4bab8b3252bf39c92298f3de60e2a73e699be323da2daf212ac3dc", "capabilities": ["administration"], "category": "administration", "child_ids": ["xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:idle_timeout:hours", "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:idle_timeout:minutes"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:idle_timeout", "parent_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration", "path": "documentation/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3320311220321303-1100332303331012-2022102131012110-3101320200312110-1300300211032100-0112202203112013-0130102010220232-2323232122022033", "registry_path": "docs/guides/data-sources--tenant_configuration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_session_expiration", "idle_timeout"], "schema_version": 1, "sections": [{"aliases": ["hours"], "anchor": "section", "description": "Represents the cookie duration in hours.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:idle_timeout:hours", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["user_session_expiration", "idle_timeout", "hours"], "syntax": "attribute", "type": "object"}, {"aliases": ["minutes"], "anchor": "section", "description": "Represents the cookie duration in minutes.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:idle_timeout:minutes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["user_session_expiration", "idle_timeout", "minutes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Represents the cookie expiration duration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_session_expiration.idle_timeout

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/)
- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/)
- user_session_expiration.idle_timeout

<a id="section"></a>

Type: `"single"`. Computed.

Represents the cookie expiration duration.

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

## Direct properties

- [hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/hours/): complete subsection reference.

- [minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/minutes/): complete subsection reference.

## Next pages

- [user_session_expiration.idle_timeout.hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/hours/)
- [user_session_expiration.idle_timeout.minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/idle_timeout/minutes/)
- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/)
- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/)
