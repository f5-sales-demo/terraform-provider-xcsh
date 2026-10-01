---
page_title: "private_key"
subcategory: "Security"
description: "private_key for xcsh_certificate."
xcsh_docs: {"aliases": [], "body_bytes": 1301, "body_sha256": "sha256:eaf30d5a052ea7a3ff6cbcf9ba0d9cde682528210ff51ce625832f1c20ec66d3", "canonical_id": "xcsh-docs:data-sources:certificate:properties:private_key", "child_ids": ["xcsh-docs:data-sources:certificate:properties:private_key:blindfold_secret_info", "xcsh-docs:data-sources:certificate:properties:private_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:certificate:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certificate:properties:private_key", "parent_id": "xcsh-docs:data-sources:certificate:reference", "path": "docs/guides/data-sources--certificate--properties--private_key.md", "provider_name": "certificate", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certificate/properties/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "private_key for xcsh_certificate.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["certificateCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_key

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md)
- [Property reference](data-sources--certificate--reference.md)
- private_key

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

- [blindfold_secret_info](data-sources--certificate--properties--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--certificate--properties--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [private_key.blindfold_secret_info](data-sources--certificate--properties--private_key--blindfold_secret_info.md)
- [private_key.clear_secret_info](data-sources--certificate--properties--private_key--clear_secret_info.md)
- [Property reference](data-sources--certificate--reference.md)
- [xcsh_certificate](../data-sources/certificate.md)
