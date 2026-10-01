---
page_title: "gre.gre_parameters.peer_ip_address"
subcategory: ""
description: "gre.gre_parameters.peer_ip_address for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 2088, "body_sha256": "sha256:ddf940cbf6192916e4cf10e16c48a7e5d1267699b3ad860744db9552c9511d1d", "canonical_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:peer_ip_address", "child_ids": [], "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:peer_ip_address", "parent_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters", "path": "docs/guides/resources--external_connector--properties--gre--gre_parameters--peer_ip_address.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["gre", "gre_parameters", "peer_ip_address"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/gre/gre_parameters/peer_ip_address/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gre.gre_parameters.peer_ip_address for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gre.gre_parameters.peer_ip_address

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md)
- [Property reference](resources--external_connector--reference.md)
- [gre](resources--external_connector--properties--gre.md)
- [gre.gre_parameters](resources--external_connector--properties--gre--gre_parameters.md)
- gre.gre_parameters.peer_ip_address

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IPv4 Address. IPv4 Address in dot-decimal notation.

Upstream description:

IPv4 Address in dot-decimal notation.

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
peer_ip_address {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-gre--gre_parameters--peer_ip_address--addr"></a>

### addr property

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

## Next pages

- [gre.gre_parameters](resources--external_connector--properties--gre--gre_parameters.md)
- [xcsh_external_connector](../resources/external_connector.md)
