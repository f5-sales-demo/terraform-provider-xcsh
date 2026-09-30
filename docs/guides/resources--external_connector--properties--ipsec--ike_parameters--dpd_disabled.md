---
page_title: "ipsec.ike_parameters.dpd_disabled"
subcategory: ""
description: "ipsec.ike_parameters.dpd_disabled for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1033, "body_sha256": "sha256:f192f7bfc55fbbad01722672ee181e04eee1f38b10e947878883f63cb6d07dbf", "canonical_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:dpd_disabled", "parent_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "path": "docs/guides/resources--external_connector--properties--ipsec--ike_parameters--dpd_disabled.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ipsec", "ike_parameters", "dpd_disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/ipsec/ike_parameters/dpd_disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ipsec.ike_parameters.dpd_disabled for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ipsec.ike_parameters.dpd_disabled

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md)
- [Property reference](resources--external_connector--reference.md)
- [ipsec](resources--external_connector--properties--ipsec.md)
- [ipsec.ike_parameters](resources--external_connector--properties--ipsec--ike_parameters.md)
- ipsec.ike_parameters.dpd_disabled

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
dpd_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ipsec.ike_parameters](resources--external_connector--properties--ipsec--ike_parameters.md)
- [xcsh_external_connector](../resources/external_connector.md)
