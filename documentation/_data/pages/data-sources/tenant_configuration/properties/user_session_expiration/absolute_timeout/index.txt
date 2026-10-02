---
page_title: "user_session_expiration.absolute_timeout"
subcategory: ""
description: "Represents the session expiration duration."
xcsh_docs: {"aliases": ["duration", "operation timeout", "user session expiration absolute timeout"], "body_bytes": 2048, "body_sha256": "sha256:f50e6c25eafce6e0eda05661d34c3a54cf7c35f291b0b7c46af805131f8ffda6", "capabilities": ["administration"], "category": "administration", "child_ids": ["xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout:hours", "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout:minutes"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tenant_configuration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout", "parent_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration", "path": "documentation/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/index.md", "product": "distributed-cloud", "provider_name": "tenant_configuration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1102032312020023-1003102221303211-0030110002211201-0112332332221321-0330301203203132-3011113203133011-2301303021220132-1100303001111231", "registry_path": "docs/guides/data-sources--tenant_configuration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["user_session_expiration", "absolute_timeout"], "schema_version": 1, "sections": [{"aliases": ["hours"], "anchor": "section", "description": "Represents the session duration in hours.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout:hours", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["user_session_expiration", "absolute_timeout", "hours"], "syntax": "attribute", "type": "object"}, {"aliases": ["minutes"], "anchor": "section", "description": "Represents the session duration in minutes.", "document_id": "xcsh-docs:data-sources:tenant_configuration:properties:user_session_expiration:absolute_timeout:minutes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["user_session_expiration", "absolute_timeout", "minutes"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Represents the session expiration duration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tenant_configurationCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# user_session_expiration.absolute_timeout

Breadcrumbs:

- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/)
- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/)
- user_session_expiration.absolute_timeout

<a id="section"></a>

Type: `"single"`. Computed.

Represents the session expiration duration.

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

- [hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/hours/): complete subsection reference.

- [minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/minutes/): complete subsection reference.

## Next pages

- [user_session_expiration.absolute_timeout.hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/hours/)
- [user_session_expiration.absolute_timeout.minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/absolute_timeout/minutes/)
- [user_session_expiration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/properties/user_session_expiration/)
- [xcsh_tenant_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tenant_configuration/)
