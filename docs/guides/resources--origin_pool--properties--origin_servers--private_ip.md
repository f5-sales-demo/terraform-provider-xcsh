---
page_title: "origin_servers.private_ip"
subcategory: "Load Balancing"
description: "origin_servers.private_ip for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 3686, "body_sha256": "sha256:015715e16338c8df708b18e96c7c113b103a2c578d7b5d7d075800dbc50b7f30", "canonical_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip", "child_ids": ["xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip:inside_network", "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip:outside_network", "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip:segment", "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip:site_locator", "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip:snat_pool"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_ip", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers", "path": "docs/guides/resources--origin_pool--properties--origin_servers--private_ip.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["origin_servers", "private_ip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/private_ip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "origin_servers.private_ip for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.private_ip

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md)
- [Property reference](resources--origin_pool--reference.md)
- [origin_servers](resources--origin_pool--properties--origin_servers.md)
- origin_servers.private_ip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with private or public IP address and site information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "segment"),
  validators.ConflictingObjectAttributes("outside_network",
    "segment")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]",
  "x-ves-oneof-field-private_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
private_ip {
  # Configure direct properties listed below.
}
```

## Direct properties

- [inside_network](resources--origin_pool--properties--origin_servers--private_ip--inside_network.md): complete subsection reference.

<a id="schema-origin_servers--private_ip--ip"></a>

### ip property

Type: `"string"`. Optional.

IP. Exclusive with \[\] Private IPv4 address.

Upstream description:

Exclusive with \[\] Private IPv4 address.

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

- [outside_network](resources--origin_pool--properties--origin_servers--private_ip--outside_network.md): complete subsection reference.

- [segment](resources--origin_pool--properties--origin_servers--private_ip--segment.md): complete subsection reference.

- [site_locator](resources--origin_pool--properties--origin_servers--private_ip--site_locator.md): complete subsection reference.

- [snat_pool](resources--origin_pool--properties--origin_servers--private_ip--snat_pool.md): complete subsection reference.

## Next pages

- [origin_servers.private_ip.inside_network](resources--origin_pool--properties--origin_servers--private_ip--inside_network.md)
- [origin_servers.private_ip.outside_network](resources--origin_pool--properties--origin_servers--private_ip--outside_network.md)
- [origin_servers.private_ip.segment](resources--origin_pool--properties--origin_servers--private_ip--segment.md)
- [origin_servers.private_ip.site_locator](resources--origin_pool--properties--origin_servers--private_ip--site_locator.md)
- [origin_servers.private_ip.snat_pool](resources--origin_pool--properties--origin_servers--private_ip--snat_pool.md)
- [origin_servers](resources--origin_pool--properties--origin_servers.md)
- [xcsh_origin_pool](../resources/origin_pool.md)
