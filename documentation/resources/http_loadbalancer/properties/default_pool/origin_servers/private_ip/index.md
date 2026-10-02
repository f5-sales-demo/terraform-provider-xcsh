---
page_title: "default_pool.origin_servers.private_ip"
subcategory: "Load Balancing"
description: "Specify origin server with private or public IP address and site information."
xcsh_docs: {"aliases": ["default pool origin servers private ip"], "body_bytes": 4949, "body_sha256": "sha256:ed8bc0dcb7062753a2c15a50e1daed9aae6d03001b650ebc28c7c6ad542f194a", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:inside_network", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:outside_network", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:segment", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator", "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:snat_pool"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers", "path": "documentation/resources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2100232131321013-0211212332213002-3003202300233021-1101200100013113-0321100032112010-1001031320030113-1122123230322022-0031130000310021", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-016.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.origin_servers.private_ip:ConflictingObjectAttributes:inside_network,outside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.origin_servers.private_ip:ConflictingObjectAttributes:inside_network,segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:inside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.origin_servers.private_ip:ConflictingObjectAttributes:inside_network,outside_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:outside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.origin_servers.private_ip:ConflictingObjectAttributes:outside_network,segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:outside_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.origin_servers.private_ip:ConflictingObjectAttributes:inside_network,segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:segment", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.origin_servers.private_ip:ConflictingObjectAttributes:outside_network,segment", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:segment", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["default_pool", "origin_servers", "private_ip"], "schema_version": 1, "sections": [{"aliases": ["inside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:inside_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_ip", "inside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["ip"], "anchor": "schema-default_pool--origin_servers--private_ip--ip", "description": "Exclusive with Private IPv4 address.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_ip", "ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["outside network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:outside_network", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_pool", "origin_servers", "private_ip", "outside_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["segment"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:segment", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-default_pool--origin_servers--private_ip--segment--name", "enforcement": "provider-schema", "group": "default_pool.origin_servers.private_ip.segment:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:segment", "type": "requires"}], "schema_path": ["default_pool", "origin_servers", "private_ip", "segment"], "syntax": "block", "type": "object"}, {"aliases": ["site locator"], "anchor": "section", "description": "This message defines a reference to a site or virtual site object.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.origin_servers.private_ip.site_locator:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.origin_servers.private_ip.site_locator:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:site_locator:virtual_site", "type": "conflicts"}], "schema_path": ["default_pool", "origin_servers", "private_ip", "site_locator"], "syntax": "block", "type": "object"}, {"aliases": ["snat pool"], "anchor": "section", "description": "SNAT Pool configuration.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:snat_pool", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.origin_servers.private_ip.snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:snat_pool:no_snat_pool", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "default_pool.origin_servers.private_ip.snat_pool:ConflictingObjectAttributes:no_snat_pool,snat_pool", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:default_pool:origin_servers:private_ip:snat_pool:snat_pool", "type": "conflicts"}], "schema_path": ["default_pool", "origin_servers", "private_ip", "snat_pool"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specify origin server with private or public IP address and site information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# default_pool.origin_servers.private_ip

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [default_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/)
- [default_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/)
- default_pool.origin_servers.private_ip

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

- [inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/inside_network/): complete subsection reference.

<a id="schema-default_pool--origin_servers--private_ip--ip"></a>

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

- [outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/outside_network/): complete subsection reference.

- [segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/segment/): complete subsection reference.

- [site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/site_locator/): complete subsection reference.

- [snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/snat_pool/): complete subsection reference.

## Next pages

- [default_pool.origin_servers.private_ip.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/inside_network/)
- [default_pool.origin_servers.private_ip.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/outside_network/)
- [default_pool.origin_servers.private_ip.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/segment/)
- [default_pool.origin_servers.private_ip.site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/site_locator/)
- [default_pool.origin_servers.private_ip.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/private_ip/snat_pool/)
- [default_pool.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/default_pool/origin_servers/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
