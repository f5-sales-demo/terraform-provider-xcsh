---
page_title: "origin_servers.private_name"
subcategory: "Load Balancing"
description: "Specify origin server with private or public DNS name and site information."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "origin servers private name", "upstream servers"], "body_bytes": 6162, "body_sha256": "sha256:9ebfde5cd600d18b7466c6d0a6d704a6aa0e5bb3354ef323f5d08ac15d5964b3", "capabilities": ["load-balancing", "load-balancing.backend-servers"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:inside_network", "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:outside_network", "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:segment", "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:site_locator", "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:snat_pool"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name", "parent_id": "xcsh-docs:resources:origin_pool:properties:origin_servers", "path": "documentation/resources/origin_pool/properties/origin_servers/private_name/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323", "registry_path": "docs/guides/resources--origin_pool--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.private_name:ConflictingObjectAttributes:inside_network,outside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.private_name:ConflictingObjectAttributes:inside_network,segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.private_name:ConflictingObjectAttributes:inside_network,outside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:outside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.private_name:ConflictingObjectAttributes:outside_network,segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:outside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.private_name:ConflictingObjectAttributes:inside_network,segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.private_name:ConflictingObjectAttributes:outside_network,segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:segment", "type": "conflicts"}, {"anchor": "schema-origin_servers--private_name--dns_name", "enforcement": "provider-schema", "group": "origin_servers.private_name:RequiredObjectAttributes:dns_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["origin_servers", "private_name"], "schema_version": 1, "sections": [{"aliases": ["dns name"], "anchor": "schema-origin_servers--private_name--dns_name", "description": "DNS Name", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "private_name", "dns_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:inside_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "private_name", "inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["outside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:outside_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "private_name", "outside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["refresh interval"], "anchor": "schema-origin_servers--private_name--refresh_interval", "description": "Interval for DNS refresh in seconds. Max value is 7 days as per https://datatracker.ietf.org/doc/HTML/rfc8767.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["origin_servers", "private_name", "refresh_interval"], "syntax": "attribute", "type": "number"}, {"aliases": ["segment"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:segment", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-origin_servers--private_name--segment--name", "enforcement": "provider-schema", "group": "origin_servers.private_name.segment:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:segment", "type": "requires"}], "schema_path": ["origin_servers", "private_name", "segment"], "syntax": "block", "type": "object"}, {"aliases": ["site locator"], "anchor": "section", "description": "This message defines a reference to a site or virtual site object.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:site_locator", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.private_name.site_locator:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:site_locator:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.private_name.site_locator:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:site_locator:virtual_site", "type": "conflicts"}], "schema_path": ["origin_servers", "private_name", "site_locator"], "syntax": "block", "type": "object"}, {"aliases": ["snat pool"], "anchor": "section", "description": "SNAT Pool configuration.", "document_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:snat_pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.private_name.snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:snat_pool:no_snat_pool", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "origin_servers.private_name.snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:origin_servers:private_name:snat_pool:snat_pool", "type": "conflicts"}], "schema_path": ["origin_servers", "private_name", "snat_pool"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/origin_servers/private_name/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Specify origin server with private or public DNS name and site information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# origin_servers.private_name

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- origin_servers.private_name

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with private or public DNS name and site information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name"),
  validators.ConflictingObjectAttributes("inside_network",
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]"
}
```

Terraform syntax:

```terraform
private_name {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-origin_servers--private_name--dns_name"></a>

### dns_name property

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/inside_network/): complete subsection reference.

- [outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/outside_network/): complete subsection reference.

<a id="schema-origin_servers--private_name--refresh_interval"></a>

### refresh_interval property

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

- [segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/segment/): complete subsection reference.

- [site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/site_locator/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/snat_pool/): complete subsection reference.

## Next pages

- [origin_servers.private_name.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/inside_network/)
- [origin_servers.private_name.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/outside_network/)
- [origin_servers.private_name.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/segment/)
- [origin_servers.private_name.site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/site_locator/)
- [origin_servers.private_name.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/private_name/snat_pool/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/origin_servers/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
