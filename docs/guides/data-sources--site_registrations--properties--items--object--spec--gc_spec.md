---
page_title: "items.object.spec.gc_spec"
subcategory: ""
description: "items.object.spec.gc_spec for xcsh_site_registrations."
xcsh_docs: {"aliases": [], "body_bytes": 2990, "body_sha256": "sha256:8d4afe652fada17ea864c3a615713422dca484d2eb07ae6e7d8114cc213cae3e", "canonical_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec", "child_ids": ["xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra", "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:passport", "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:site"], "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec", "path": "docs/guides/data-sources--site_registrations--properties--items--object--spec--gc_spec.md", "provider_name": "site_registrations", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/object/spec/gc_spec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.spec.gc_spec for xcsh_site_registrations.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec

Breadcrumbs:

- [xcsh_site_registrations](../data-sources/site_registrations.md)
- [Property reference](data-sources--site_registrations--reference.md)
- [items](data-sources--site_registrations--properties--items.md)
- [items.object](data-sources--site_registrations--properties--items--object.md)
- [items.object.spec](data-sources--site_registrations--properties--items--object--spec.md)
- items.object.spec.gc_spec

<a id="section"></a>

Type: `"single"`. Computed.

Global Specification.

## Direct properties

<a id="schema-items--object--spec--gc_spec--connected_regions"></a>

### connected_regions property

Type: `["list", "string"]`. Computed.

Optional. REs in selected region to which CEs connect, contains primary and backup RE region info.

- [infra](data-sources--site_registrations--properties--items--object--spec--gc_spec--infra.md): complete subsection reference.

- [passport](data-sources--site_registrations--properties--items--object--spec--gc_spec--passport.md): complete subsection reference.

<a id="schema-items--object--spec--gc_spec--role"></a>

### role property

Type: `["list", "string"]`. Computed.

Role of registered node. Used by system to determine what roles should be enforced.

- [site](data-sources--site_registrations--properties--items--object--spec--gc_spec--site.md): complete subsection reference.

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

- [items.object.spec.gc_spec.infra](data-sources--site_registrations--properties--items--object--spec--gc_spec--infra.md)
- [items.object.spec.gc_spec.passport](data-sources--site_registrations--properties--items--object--spec--gc_spec--passport.md)
- [items.object.spec.gc_spec.site](data-sources--site_registrations--properties--items--object--spec--gc_spec--site.md)
- [items.object.spec](data-sources--site_registrations--properties--items--object--spec.md)
- [xcsh_site_registrations](../data-sources/site_registrations.md)
