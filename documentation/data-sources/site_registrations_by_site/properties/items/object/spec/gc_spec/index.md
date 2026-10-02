---
page_title: "items.object.spec.gc_spec"
subcategory: ""
description: "Global Specification."
xcsh_docs: {"aliases": ["items object spec gc spec"], "body_bytes": 3746, "body_sha256": "sha256:74ee2eec1ec6ff66f31dbdd9df0a856e260221e156460a3db008841f2837878a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:passport", "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec", "path": "documentation/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3121202112033110-3121103020022113-0230130011300322-2020223020011123-3020210211300200-0122212301213210-0013211323122211-0311212021313013", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec"], "schema_version": 1, "sections": [{"aliases": ["connected regions"], "anchor": "schema-items--object--spec--gc_spec--connected_regions", "description": "Optional. REs in selected region to which CEs connect, contains primary and backup RE region info.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "connected_regions"], "syntax": "attribute", "type": "list"}, {"aliases": ["infra"], "anchor": "section", "description": "InfraMetadata stores information about instance infrastructure.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra"], "syntax": "attribute", "type": "object"}, {"aliases": ["passport"], "anchor": "section", "description": "Passport stores information about identification and node configuration provided by CE during registration. It can be manually updated by user during approval.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:passport", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "passport"], "syntax": "attribute", "type": "object"}, {"aliases": ["role"], "anchor": "schema-items--object--spec--gc_spec--role", "description": "Role of registered node. Used by system to determine what roles should be enforced.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "role"], "syntax": "attribute", "type": "list"}, {"aliases": ["site"], "anchor": "section", "description": "Site for this registration, assigned after registration is assigned to site.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "site"], "syntax": "attribute", "type": "object"}, {"aliases": ["token"], "anchor": "schema-items--object--spec--gc_spec--token", "description": "Token is used for machine and tenant identification.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "token"], "syntax": "attribute", "type": "string"}, {"aliases": ["tunnel type"], "anchor": "schema-items--object--spec--gc_spec--tunnel_type", "description": "Tunnel encapsulation to be used between sites Tunnel can operate in both IPsec and SSL, with IPsec being preferred over SSL. Tunnel is of type IPsec Tunnel is of type SSL. Possible values are `SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL`, `SITE_TO_SITE_TUNNEL_IPSEC`, `SITE_TO_SITE_TUNNEL_SSL`. Defaults to", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "tunnel_type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Global Specification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/)
- items.object.spec.gc_spec

<a id="section"></a>

Type: `"single"`. Computed.

Global Specification.

## Direct properties

<a id="schema-items--object--spec--gc_spec--connected_regions"></a>

### connected_regions property

Type: `["list", "string"]`. Computed.

Optional. REs in selected region to which CEs connect, contains primary and backup RE region info.

- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/): complete subsection reference.

- [passport](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/passport/): complete subsection reference.

<a id="schema-items--object--spec--gc_spec--role"></a>

### role property

Type: `["list", "string"]`. Computed.

Role of registered node. Used by system to determine what roles should be enforced.

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/site/): complete subsection reference.

<a id="schema-items--object--spec--gc_spec--token"></a>

### token property

Type: `"string"`. Computed.

Token is used for machine and tenant identification.

<a id="schema-items--object--spec--gc_spec--tunnel_type"></a>

### tunnel_type property

Type: `"string"`. Computed.

\[Enum:
SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL|SITE\_TO\_SITE\_TUNNEL\_IPSEC|SITE\_TO\_SITE\_TUNNEL\_SSL\]
Tunnel encapsulation to be used between sites Tunnel can operate in both IPsec and SSL, with IPsec
being preferred over SSL. Tunnel is of type IPsec Tunnel is of type SSL. Possible values are
\`SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL\`, \`SITE\_TO\_SITE\_TUNNEL\_IPSEC\`,
\`SITE\_TO\_SITE\_TUNNEL\_SSL\`. Defaults to \`SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL",
    "SITE_TO_SITE_TUNNEL_IPSEC",
    "SITE_TO_SITE_TUNNEL_SSL"),
}
```

## Next pages

- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/)
- [items.object.spec.gc_spec.passport](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/passport/)
- [items.object.spec.gc_spec.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/site/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
