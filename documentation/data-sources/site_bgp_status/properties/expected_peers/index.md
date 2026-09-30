---
page_title: "expected_peers"
subcategory: ""
description: "expected_peers for xcsh_site_bgp_status."
xcsh_docs: {"aliases": [], "body_bytes": 1335, "body_sha256": "sha256:eb6da0407c608561ade76b3241c3e19e0c6f0162d83cbe1322cce5d36b4d3d67", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_bgp_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "parent_id": "xcsh-docs:data-sources:site_bgp_status:reference", "path": "documentation/data-sources/site_bgp_status/properties/expected_peers/index.md", "provider_name": "site_bgp_status", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["expected_peers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_bgp_status/properties/expected_peers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "expected_peers for xcsh_site_bgp_status.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# expected_peers

Breadcrumbs:

- [xcsh_site_bgp_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/)
- expected_peers

<a id="section"></a>

Type: `"map"`. Required.

## Direct properties

<a id="schema-expected_peers--expected_imported_routes"></a>

### expected_imported_routes property

Type: `["set", "string"]`. Required.

Exact prefixes that must be imported from this remote peer on the expected node.

<a id="schema-expected_peers--mac"></a>

### mac property

Type: `"string"`. Required.

<a id="schema-expected_peers--node"></a>

### node property

Type: `"string"`. Required.

<a id="schema-expected_peers--peer_address"></a>

### peer_address property

Type: `"string"`. Required.

<a id="schema-expected_peers--role"></a>

### role property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.OneOf("slo",
    "sli")}
```

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/properties/)
- [xcsh_site_bgp_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_bgp_status/)
