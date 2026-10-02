---
page_title: "routes.request_headers_to_add.secret_value"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["routes request headers to add secret value"], "body_bytes": 2196, "body_sha256": "sha256:d468a3205ede74813db31431cb55ce3f5d6e1ace0a7a88a3a5a766699d735991", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:route:properties:routes:request_headers_to_add:secret_value:blindfold_secret_info", "xcsh-docs:data-sources:route:properties:routes:request_headers_to_add:secret_value:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:route:properties:routes:request_headers_to_add:secret_value", "parent_id": "xcsh-docs:data-sources:route:properties:routes:request_headers_to_add", "path": "documentation/data-sources/route/properties/routes/request_headers_to_add/secret_value/index.md", "product": "distributed-cloud", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0211000021002133-2200333200112330-1003321112000133-3112000303313301-1022100211101110-1233121020213122-0231032100231211-3133302022010000", "registry_path": "docs/guides/data-sources--route--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["routes", "request_headers_to_add", "secret_value"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:route:properties:routes:request_headers_to_add:secret_value:blindfold_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "request_headers_to_add", "secret_value", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:route:properties:routes:request_headers_to_add:secret_value:clear_secret_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["routes", "request_headers_to_add", "secret_value", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/properties/routes/request_headers_to_add/secret_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.request_headers_to_add.secret_value

Breadcrumbs:

- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/)
- [routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/)
- [routes.request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/request_headers_to_add/)
- routes.request_headers_to_add.secret_value

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/request_headers_to_add/secret_value/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/request_headers_to_add/secret_value/clear_secret_info/): complete subsection reference.

## Next pages

- [routes.request_headers_to_add.secret_value.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/request_headers_to_add/secret_value/blindfold_secret_info/)
- [routes.request_headers_to_add.secret_value.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/request_headers_to_add/secret_value/clear_secret_info/)
- [routes.request_headers_to_add](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/properties/routes/request_headers_to_add/)
- [xcsh_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/route/)
