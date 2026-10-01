---
page_title: "params.ipsec.ipsec_psk"
subcategory: ""
description: "params.ipsec.ipsec_psk for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1758, "body_sha256": "sha256:b9873423aa45e84a37df3cdde6956f028e529e9603b2ac058938329fd194de56", "canonical_id": "xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk", "child_ids": ["xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk:blindfold_secret_info", "xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk:clear_secret_info"], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:properties:params:ipsec:ipsec_psk", "parent_id": "xcsh-docs:resources:tunnel:properties:params:ipsec", "path": "docs/guides/resources--tunnel--properties--params--ipsec--ipsec_psk.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["params", "ipsec", "ipsec_psk"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/properties/params/ipsec/ipsec_psk/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "params.ipsec.ipsec_psk for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# params.ipsec.ipsec_psk

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md)
- [Property reference](resources--tunnel--reference.md)
- [params](resources--tunnel--properties--params.md)
- [params.ipsec](resources--tunnel--properties--params--ipsec.md)
- params.ipsec.ipsec_psk

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
ipsec_psk {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](resources--tunnel--properties--params--ipsec--ipsec_psk--blindfold_secret_info.md): complete subsection reference.

- [clear_secret_info](resources--tunnel--properties--params--ipsec--ipsec_psk--clear_secret_info.md): complete subsection reference.

## Next pages

- [params.ipsec.ipsec_psk.blindfold_secret_info](resources--tunnel--properties--params--ipsec--ipsec_psk--blindfold_secret_info.md)
- [params.ipsec.ipsec_psk.clear_secret_info](resources--tunnel--properties--params--ipsec--ipsec_psk--clear_secret_info.md)
- [params.ipsec](resources--tunnel--properties--params--ipsec.md)
- [xcsh_tunnel](../resources/tunnel.md)
