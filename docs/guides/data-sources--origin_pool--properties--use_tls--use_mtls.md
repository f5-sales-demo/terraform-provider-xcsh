---
page_title: "use_tls.use_mtls"
subcategory: "Load Balancing"
description: "use_tls.use_mtls for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1055, "body_sha256": "sha256:afcfb9c0ff96007eeddfa642acd0895e9251b710c725ca779db80fa5e982e4dc", "canonical_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls", "child_ids": ["xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls:tls_certificates"], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:properties:use_tls:use_mtls", "parent_id": "xcsh-docs:data-sources:origin_pool:properties:use_tls", "path": "docs/guides/data-sources--origin_pool--properties--use_tls--use_mtls.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["use_tls", "use_mtls"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/properties/use_tls/use_mtls/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "use_tls.use_mtls for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_tls.use_mtls

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Property reference](data-sources--origin_pool--reference.md)
- [use_tls](data-sources--origin_pool--properties--use_tls.md)
- use_tls.use_mtls

<a id="section"></a>

Type: `"single"`. Computed.

MTLS Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

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

## Direct properties

- [tls_certificates](data-sources--origin_pool--properties--use_tls--use_mtls--tls_certificates.md): complete subsection reference.

## Next pages

- [use_tls.use_mtls.tls_certificates](data-sources--origin_pool--properties--use_tls--use_mtls--tls_certificates.md)
- [use_tls](data-sources--origin_pool--properties--use_tls.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
