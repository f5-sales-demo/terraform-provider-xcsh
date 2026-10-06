---
page_title: "local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface"
subcategory: ""
description: "Interface reference on this node."
xcsh_docs: {"aliases": ["local vrf slo config static v6 routes static routes node interface list interface"], "body_bytes": 7015, "body_sha256": "sha256:9ba543b6af7c8ecb65858eda6ced090811d538b4ccfee19483733ef3bc0ddd39", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes:static_routes:node_interface:list:interface", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes:static_routes:node_interface:list", "path": "documentation/resources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/static_routes/node_interface/list/interface/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2122222222000100-3331202322230202-3133030301110313-0302103313201321-0021332232102111-0023322113201033-3220300222010312-1101323202233330", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_vrf", "slo_config", "static_v6_routes", "static_routes", "node_interface", "list", "interface"], "schema_version": 1, "sections": [{"aliases": ["local vrf slo config static v6 routes static routes node interface list interface kind"], "anchor": "schema-local_vrf--slo_config--static_v6_routes--static_routes--node_interface--list--interface--kind", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g. \"route\")", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes:static_routes:node_interface:list:interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "static_v6_routes", "static_routes", "node_interface", "list", "interface", "kind"], "syntax": "attribute", "type": "string"}, {"aliases": ["local vrf slo config static v6 routes static routes node interface list interface name"], "anchor": "schema-local_vrf--slo_config--static_v6_routes--static_routes--node_interface--list--interface--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes:static_routes:node_interface:list:interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "static_v6_routes", "static_routes", "node_interface", "list", "interface", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["local vrf slo config static v6 routes static routes node interface list interface namespace"], "anchor": "schema-local_vrf--slo_config--static_v6_routes--static_routes--node_interface--list--interface--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes:static_routes:node_interface:list:interface", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "static_v6_routes", "static_routes", "node_interface", "list", "interface", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["local vrf slo config static v6 routes static routes node interface list interface tenant"], "anchor": "schema-local_vrf--slo_config--static_v6_routes--static_routes--node_interface--list--interface--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes:static_routes:node_interface:list:interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "static_v6_routes", "static_routes", "node_interface", "list", "interface", "tenant"], "syntax": "attribute", "type": "string"}, {"aliases": ["local vrf slo config static v6 routes static routes node interface list interface uid"], "anchor": "schema-local_vrf--slo_config--static_v6_routes--static_routes--node_interface--list--interface--uid", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then uid will hold the referred object's(e.g. Route's) uid.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_v6_routes:static_routes:node_interface:list:interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_vrf", "slo_config", "static_v6_routes", "static_routes", "node_interface", "list", "interface", "uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/static_routes/node_interface/list/interface/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Interface reference on this node.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [local_vrf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/)
- [local_vrf.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/)
- [local_vrf.slo_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/)
- [local_vrf.slo_config.static_v6_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/static_routes/)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/static_routes/node_interface/)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/static_v6_routes/static_routes/node_interface/list/)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-local_vrf--slo_config--static_v6_routes--static_routes--node_interface--list--interface--kind"></a>

### kind property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-local_vrf--slo_config--static_v6_routes--static_routes--node_interface--list--interface--name"></a>

### name property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-local_vrf--slo_config--static_v6_routes--static_routes--node_interface--list--interface--namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-local_vrf--slo_config--static_v6_routes--static_routes--node_interface--list--interface--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-local_vrf--slo_config--static_v6_routes--static_routes--node_interface--list--interface--uid"></a>

### uid property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
