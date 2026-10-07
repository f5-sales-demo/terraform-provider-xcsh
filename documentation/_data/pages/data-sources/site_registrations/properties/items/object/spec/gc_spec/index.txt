---
page_title: "items.object.spec.gc_spec"
subcategory: ""
description: "Global Specification."
xcsh_docs: {"aliases": ["items object spec gc spec"], "body_bytes": 3165, "body_sha256": "sha256:22174666d7d76f0b5d897a1578bcade26cc2904a885050f0a4f396c4ab5c0f94", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra", "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:passport", "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:site"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec", "path": "documentation/data-sources/site_registrations/properties/items/object/spec/gc_spec/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1030101233031222-1002233113211320-3232321032333213-0231300013231132-3032101123211102-2022211203311030-2013312033010111-0010302013023311", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec"], "schema_version": 1, "sections": [{"aliases": ["items object spec gc spec connected regions"], "anchor": "schema-items--object--spec--gc_spec--connected_regions", "description": "Optional. REs in selected region to which CEs connect, contains primary and backup RE region info.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "connected_regions"], "syntax": "attribute", "type": "list"}, {"aliases": ["items object spec gc spec infra"], "anchor": "section", "description": "InfraMetadata stores information about instance infrastructure.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object spec gc spec passport"], "anchor": "section", "description": "Passport stores information about identification and node configuration provided by CE during registration. It can be manually updated by user during approval.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:passport", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "passport"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object spec gc spec role"], "anchor": "schema-items--object--spec--gc_spec--role", "description": "Role of registered node. Used by system to determine what roles should be enforced.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "role"], "syntax": "attribute", "type": "list"}, {"aliases": ["items object spec gc spec site"], "anchor": "section", "description": "Site for this registration, assigned after registration is assigned to site.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:site", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "site"], "syntax": "attribute", "type": "object"}, {"aliases": ["items object spec gc spec token"], "anchor": "schema-items--object--spec--gc_spec--token", "description": "Token is used for machine and tenant identification.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "token"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec tunnel type"], "anchor": "schema-items--object--spec--gc_spec--tunnel_type", "description": "Tunnel encapsulation to be used between sites Tunnel can operate in both IPsec and SSL, with IPsec being preferred over SSL. Tunnel is of type IPsec Tunnel is of type SSL. Possible values are `SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL`, `SITE_TO_SITE_TUNNEL_IPSEC`, `SITE_TO_SITE_TUNNEL_SSL`. Defaults to", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["SITE_TO_SITE_TUNNEL_IPSEC", "SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL", "SITE_TO_SITE_TUNNEL_SSL"], "version": 1}], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "tunnel_type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/object/spec/gc_spec/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Global Specification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/)
- items.object.spec.gc_spec

<a id="section"></a>

Type: `"single"`. Computed.

Global Specification.

## Direct properties

<a id="schema-items--object--spec--gc_spec--connected_regions"></a>

### connected_regions property

Type: `["list", "string"]`. Computed.

Optional. REs in selected region to which CEs connect, contains primary and backup RE region info.

- [infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/): complete subsection reference.

- [passport](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/passport/): complete subsection reference.

<a id="schema-items--object--spec--gc_spec--role"></a>

### role property

Type: `["list", "string"]`. Computed.

Role of registered node. Used by system to determine what roles should be enforced.

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/site/): complete subsection reference.

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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_TO_SITE_TUNNEL_IPSEC","SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL","SITE_TO_SITE_TUNNEL_SSL"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL",
    "SITE_TO_SITE_TUNNEL_IPSEC",
    "SITE_TO_SITE_TUNNEL_SSL"),
}
```
