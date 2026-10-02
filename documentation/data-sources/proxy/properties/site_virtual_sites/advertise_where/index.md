---
page_title: "site_virtual_sites.advertise_where"
subcategory: ""
description: "Where should this load balancer be available."
xcsh_docs: {"aliases": ["site virtual sites advertise where"], "body_bytes": 3957, "body_sha256": "sha256:8da99c83d82bfa6a24b8965f7bbe3942edb03ba97c40e2560c16f5e651a09f3f", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:site", "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:use_default_port", "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:virtual_site"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where", "parent_id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites", "path": "documentation/data-sources/proxy/properties/site_virtual_sites/advertise_where/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1101311121310312-0200013102222131-0303021112013023-0200203203003231-1221230311001231-0230200233121202-1303302123221200-0111202023011311", "registry_path": "docs/guides/data-sources--proxy--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["site_virtual_sites", "advertise_where"], "schema_version": 1, "sections": [{"aliases": ["port"], "anchor": "schema-site_virtual_sites--advertise_where--port", "description": "Exclusive with TCP port to Listen.", "document_id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_virtual_sites", "advertise_where", "port"], "syntax": "attribute", "type": "number"}, {"aliases": ["site"], "anchor": "section", "description": "This defines a reference to a CE site along with network type and an optional IP address where a load balancer could be advertised.", "document_id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_virtual_sites", "advertise_where", "site"], "syntax": "attribute", "type": "object"}, {"aliases": ["use default port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:use_default_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_virtual_sites", "advertise_where", "use_default_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual site"], "anchor": "section", "description": "This defines a reference to a customer site virtual site along with network type where a load balancer could be advertised.", "document_id": "xcsh-docs:data-sources:proxy:properties:site_virtual_sites:advertise_where:virtual_site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["site_virtual_sites", "advertise_where", "virtual_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/proxy/properties/site_virtual_sites/advertise_where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Where should this load balancer be available.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_virtual_sites.advertise_where

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/)
- [site_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/site_virtual_sites/)
- site_virtual_sites.advertise_where

<a id="section"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-site_virtual_sites--advertise_where--port"></a>

### port property

Type: `"number"`. Computed.

Exclusive with \[use\_default\_port\] TCP port to Listen.

Upstream description:

Exclusive with \[use\_default\_port\] TCP port to Listen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/site_virtual_sites/advertise_where/site/): complete subsection reference.

- [use_default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/site_virtual_sites/advertise_where/use_default_port/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/site_virtual_sites/advertise_where/virtual_site/): complete subsection reference.

## Next pages

- [site_virtual_sites.advertise_where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/site_virtual_sites/advertise_where/site/)
- [site_virtual_sites.advertise_where.use_default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/site_virtual_sites/advertise_where/use_default_port/)
- [site_virtual_sites.advertise_where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/site_virtual_sites/advertise_where/virtual_site/)
- [site_virtual_sites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/properties/site_virtual_sites/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/proxy/)
