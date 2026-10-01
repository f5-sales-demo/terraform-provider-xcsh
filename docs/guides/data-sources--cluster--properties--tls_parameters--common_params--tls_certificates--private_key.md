---
page_title: "tls_parameters.common_params.tls_certificates.private_key"
subcategory: ""
description: "tls_parameters.common_params.tls_certificates.private_key for xcsh_cluster."
xcsh_docs: {"aliases": [], "body_bytes": 2030, "body_sha256": "sha256:8e3f0aacb058aad070d464592d62b8fd227e405cf2074892cac4e9e513cab4ff", "canonical_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key", "child_ids": ["xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params:tls_certificates:private_key", "parent_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params:tls_certificates", "path": "docs/guides/data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key.md", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters", "common_params", "tls_certificates", "private_key"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/tls_parameters/common_params/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.common_params.tls_certificates.private_key for xcsh_cluster.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.tls_certificates.private_key

Breadcrumbs:

- [xcsh_cluster](../data-sources/cluster.md)
- [Property reference](data-sources--cluster--reference.md)
- [tls_parameters](data-sources--cluster--properties--tls_parameters.md)
- [tls_parameters.common_params](data-sources--cluster--properties--tls_parameters--common_params.md)
- [tls_parameters.common_params.tls_certificates](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates.md)
- tls_parameters.common_params.tls_certificates.private_key

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

- [blindfold_secret_info](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md): complete subsection reference.

## Next pages

- [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key--blindfold_secret_info.md)
- [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates--private_key--clear_secret_info.md)
- [tls_parameters.common_params.tls_certificates](data-sources--cluster--properties--tls_parameters--common_params--tls_certificates.md)
- [xcsh_cluster](../data-sources/cluster.md)
