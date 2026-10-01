---
page_title: "params.ipsec.ipsec_psk"
subcategory: ""
description: "params.ipsec.ipsec_psk for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1980, "body_sha256": "sha256:22b88cae5b318e36df88b1d1039c3100e42c55e822b43565e0a1a4512b116757", "child_ids": ["xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk:blindfold_secret_info", "xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:properties:params:ipsec:ipsec_psk", "parent_id": "xcsh-docs:data-sources:tunnel:properties:params:ipsec", "path": "documentation/data-sources/tunnel/properties/params/ipsec/ipsec_psk/index.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["params", "ipsec", "ipsec_psk"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/properties/params/ipsec/ipsec_psk/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "params.ipsec.ipsec_psk for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# params.ipsec.ipsec_psk

Breadcrumbs:

- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/params/)
- [params.ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/params/ipsec/)
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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/params/ipsec/ipsec_psk/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/params/ipsec/ipsec_psk/clear_secret_info/): complete subsection reference.

## Next pages

- [params.ipsec.ipsec_psk.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/params/ipsec/ipsec_psk/blindfold_secret_info/)
- [params.ipsec.ipsec_psk.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/params/ipsec/ipsec_psk/clear_secret_info/)
- [params.ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/params/ipsec/)
- [xcsh_tunnel](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/)
