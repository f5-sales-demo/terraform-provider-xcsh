---
page_title: "gre.gre_parameters"
subcategory: ""
description: "gre.gre_parameters for xcsh_external_connector."
xcsh_docs: {"aliases": [], "body_bytes": 3658, "body_sha256": "sha256:25f8256ac022ab10bed7f5a9c5606805bb122793ba1eedbfa318bda6e86f4979", "canonical_id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters", "child_ids": ["xcsh-docs:resources:external_connector:properties:gre:gre_parameters:peer_ip_address", "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:segment", "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_inside_network", "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:site_local_network", "xcsh-docs:resources:external_connector:properties:gre:gre_parameters:tunnel_eps"], "collection_id": "xcsh-docs:resources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:external_connector:properties:gre:gre_parameters", "parent_id": "xcsh-docs:resources:external_connector:properties:gre", "path": "docs/guides/resources--external_connector--properties--gre--gre_parameters.md", "provider_name": "external_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["gre", "gre_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/external_connector/properties/gre/gre_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "gre.gre_parameters for xcsh_external_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# gre.gre_parameters

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md)
- [Property reference](resources--external_connector--reference.md)
- [gre](resources--external_connector--properties--gre.md)
- gre.gre_parameters

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

GRE configuration parameters required for GRE Connection type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tunnel_eps",
    "tunnel_mtu"),
  validators.ConflictingObjectAttributes("segment",
    "site_local_inside_network"),
  validators.ConflictingObjectAttributes("segment",
    "site_local_network"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-tunnel_source_vn": "[\"segment\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
gre_parameters {
  # Configure direct properties listed below.
}
```

## Direct properties

- [peer_ip_address](resources--external_connector--properties--gre--gre_parameters--peer_ip_address.md): complete subsection reference.

- [segment](resources--external_connector--properties--gre--gre_parameters--segment.md): complete subsection reference.

- [site_local_inside_network](resources--external_connector--properties--gre--gre_parameters--site_local_inside_network.md): complete subsection reference.

- [site_local_network](resources--external_connector--properties--gre--gre_parameters--site_local_network.md): complete subsection reference.

- [tunnel_eps](resources--external_connector--properties--gre--gre_parameters--tunnel_eps.md): complete subsection reference.

<a id="schema-gre--gre_parameters--tunnel_mtu"></a>

### tunnel_mtu property

Type: `"number"`. Optional.

Configure MTU for the GRE tunnel interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(512, 1370),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1370,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 512
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  }
}
```

## Next pages

- [gre.gre_parameters.peer_ip_address](resources--external_connector--properties--gre--gre_parameters--peer_ip_address.md)
- [gre.gre_parameters.segment](resources--external_connector--properties--gre--gre_parameters--segment.md)
- [gre.gre_parameters.site_local_inside_network](resources--external_connector--properties--gre--gre_parameters--site_local_inside_network.md)
- [gre.gre_parameters.site_local_network](resources--external_connector--properties--gre--gre_parameters--site_local_network.md)
- [gre.gre_parameters.tunnel_eps](resources--external_connector--properties--gre--gre_parameters--tunnel_eps.md)
- [gre](resources--external_connector--properties--gre.md)
- [xcsh_external_connector](../resources/external_connector.md)
