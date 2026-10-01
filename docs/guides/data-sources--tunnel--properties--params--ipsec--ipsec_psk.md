---
page_title: "params.ipsec.ipsec_psk"
subcategory: ""
description: "params.ipsec.ipsec_psk for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1482, "body_sha256": "sha256:b4e41313871c0be1ed3d87852802b54c579c6c070af2d63852dd37c6c8cb99d9", "canonical_id": "xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk:blindfold_secret_info", "xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk", "parent_id": "xcsh-docs:data-sources:tunnel:properties:params:ipsec", "path": "docs/guides/data-sources--tunnel--properties--params--ipsec--ipsec_psk.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["params", "ipsec", "ipsec_psk"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/params/ipsec/ipsec_psk/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "params.ipsec.ipsec_psk for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# params.ipsec.ipsec_psk

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md)
- [Property reference](data-sources--tunnel--reference.md)
- [params](data-sources--tunnel--properties--params.md)
- [params.ipsec](data-sources--tunnel--properties--params--ipsec.md)
- params.ipsec.ipsec_psk

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

- [blindfold_secret_info](data-sources--tunnel--properties--params--ipsec--ipsec_psk--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](data-sources--tunnel--properties--params--ipsec--ipsec_psk--clear_secret_info.md): complete subsection reference.

## Next pages

- [params.ipsec.ipsec_psk.blindfold_secret_info](data-sources--tunnel--properties--params--ipsec--ipsec_psk--blindfold_secret_info.md)
- [params.ipsec.ipsec_psk.clear_secret_info](data-sources--tunnel--properties--params--ipsec--ipsec_psk--clear_secret_info.md)
- [params.ipsec](data-sources--tunnel--properties--params--ipsec.md)
- [xcsh_tunnel](../data-sources/tunnel.md)
