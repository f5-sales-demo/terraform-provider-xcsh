---
page_title: "ipsec.ike_parameters.responder"
subcategory: ""
description: "ipsec.ike_parameters.responder for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1123, "body_sha256": "sha256:c5a7030fba28f2421aa438864535c1715a42c9ab8655d7406b0484c0cba235d2", "canonical_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:responder", "child_ids": [], "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters:responder", "parent_id": "xcsh-docs:resources:external_connector:properties:ipsec:ike_parameters", "path": "docs/guides/resources--external_connector--properties--ipsec--ike_parameters--responder.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ipsec", "ike_parameters", "responder"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/ipsec/ike_parameters/responder/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ipsec.ike_parameters.responder for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ipsec.ike_parameters.responder

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md)
- [Property reference](resources--external_connector--reference.md)
- [ipsec](resources--external_connector--properties--ipsec.md)
- [ipsec.ike_parameters](resources--external_connector--properties--ipsec--ike_parameters.md)
- ipsec.ike_parameters.responder

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
responder = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ipsec.ike_parameters](resources--external_connector--properties--ipsec--ike_parameters.md)
- [xcsh_external_connector](../resources/external_connector.md)
