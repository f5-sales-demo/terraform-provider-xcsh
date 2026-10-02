---
page_title: "ethernet_interface.dhcp_server.dhcp_networks.pools"
subcategory: ""
description: "List of non overlapping IP address ranges."
xcsh_docs: {"aliases": ["ethernet interface dhcp server dhcp networks pools"], "body_bytes": 5031, "body_sha256": "sha256:ff472cf69e8cdf10cbede673195dd5aa8d7840b2c1aa07e826503d6a8ac648c0", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:pools", "parent_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks", "path": "documentation/resources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/pools/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0022313112313210-0123113032330113-1133331031313132-0123232333232223-0221003112011231-0100203021231322-1113212223121130-3203121100201013", "registry_path": "docs/guides/resources--network_interface--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_networks", "pools"], "schema_version": 1, "sections": [{"aliases": ["end ip"], "anchor": "schema-ethernet_interface--dhcp_server--dhcp_networks--pools--end_ip", "description": "Ending IP of the pool range. In case of address allocator, offset is derived based on network prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:pools", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_networks", "pools", "end_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["exclude"], "anchor": "schema-ethernet_interface--dhcp_server--dhcp_networks--pools--exclude", "description": "Exclude this address range from DHCP allocation.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:pools", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_networks", "pools", "exclude"], "syntax": "attribute", "type": "bool"}, {"aliases": ["start ip"], "anchor": "schema-ethernet_interface--dhcp_server--dhcp_networks--pools--start_ip", "description": "Starting IP of the pool range. In case of address allocator, offset is derived based on network prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.", "document_id": "xcsh-docs:resources:network_interface:properties:ethernet_interface:dhcp_server:dhcp_networks:pools", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ethernet_interface", "dhcp_server", "dhcp_networks", "pools", "start_ip"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/pools/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of non overlapping IP address ranges.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ethernet_interface.dhcp_server.dhcp_networks.pools

Breadcrumbs:

- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/)
- [ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/)
- [ethernet_interface.dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/dhcp_server/)
- [ethernet_interface.dhcp_server.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/)
- ethernet_interface.dhcp_server.dhcp_networks.pools

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of non overlapping IP address ranges.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ethernet_interface--dhcp_server--dhcp_networks--pools--end_ip"></a>

### end_ip property

Type: `"string"`. Optional.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

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

<a id="schema-ethernet_interface--dhcp_server--dhcp_networks--pools--exclude"></a>

### exclude property

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="schema-ethernet_interface--dhcp_server--dhcp_networks--pools--start_ip"></a>

### start_ip property

Type: `"string"`. Optional.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

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

- [ethernet_interface.dhcp_server.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/properties/ethernet_interface/dhcp_server/dhcp_networks/)
- [xcsh_network_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_interface/)
