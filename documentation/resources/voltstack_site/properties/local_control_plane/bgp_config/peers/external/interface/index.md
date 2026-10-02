---
page_title: "local_control_plane.bgp_config.peers.external.interface"
subcategory: ""
description: "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name."
xcsh_docs: {"aliases": ["local control plane bgp config peers external interface"], "body_bytes": 6798, "body_sha256": "sha256:84926abb6db0c5777dbfea94d4fc973882eaa5f81b546797bc904dec536e580e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:interface", "parent_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external", "path": "documentation/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/interface/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2132121323021033-1301300230310202-3310300120210001-3233032120232003-3202113202032330-3221223313212132-0031011002213312-0210211120012302", "registry_path": "docs/guides/resources--voltstack_site--reference--group-009.md", "relationships": [{"anchor": "schema-local_control_plane--bgp_config--peers--external--interface--name", "enforcement": "provider-schema", "group": "local_control_plane.bgp_config.peers.external.interface:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:interface", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "interface"], "schema_version": 1, "sections": [{"aliases": ["name"], "anchor": "schema-local_control_plane--bgp_config--peers--external--interface--name", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then name will hold the referred object's(e.g. Route's) name.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:interface", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "interface", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-local_control_plane--bgp_config--peers--external--interface--namespace", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then namespace will hold the referred object's(e.g. Route's) namespace.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:interface", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "interface", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["tenant"], "anchor": "schema-local_control_plane--bgp_config--peers--external--interface--tenant", "description": "When a configuration object(e.g. Virtual_host) refers to another(e.g route) then tenant will hold the referred object's(e.g. Route's) tenant.", "document_id": "xcsh-docs:resources:voltstack_site:properties:local_control_plane:bgp_config:peers:external:interface", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers", "external", "interface", "tenant"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/interface/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config.peers.external.interface

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/)
- [local_control_plane.bgp_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/)
- [local_control_plane.bgp_config.peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/)
- [local_control_plane.bgp_config.peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/)
- local_control_plane.bgp_config.peers.external.interface

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
interface {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-local_control_plane--bgp_config--peers--external--interface--name"></a>

### name property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="schema-local_control_plane--bgp_config--peers--external--interface--namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="schema-local_control_plane--bgp_config--peers--external--interface--tenant"></a>

### tenant property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

## Next pages

- [local_control_plane.bgp_config.peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/local_control_plane/bgp_config/peers/external/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
