---
page_title: "local_control_plane.bgp_config"
subcategory: ""
description: "BGP configuration parameters."
xcsh_docs: {"aliases": ["local control plane bgp config"], "body_bytes": 2402, "body_sha256": "sha256:462aad8bfda0a91a42174fc20a0dd9c9f3bfd80814fc800d0ccca6dc7bcd09fc", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config", "parent_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane", "path": "documentation/data-sources/voltstack_site/properties/local_control_plane/bgp_config/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1003123310232102-0232300200310320-0030002202230022-2333121312030312-2131102203000003-3110131330033301-0303301213230312-2133123031033100", "registry_path": "docs/guides/data-sources--voltstack_site--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_control_plane", "bgp_config"], "schema_version": 1, "sections": [{"aliases": ["asn"], "anchor": "schema-local_control_plane--bgp_config--asn", "description": "Autonomous System Number.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "asn"], "syntax": "attribute", "type": "number"}, {"aliases": ["peers"], "anchor": "section", "description": "BGP parameters for peer.", "document_id": "xcsh-docs:data-sources:voltstack_site:properties:local_control_plane:bgp_config:peers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["local_control_plane", "bgp_config", "peers"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/properties/local_control_plane/bgp_config/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "BGP configuration parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_control_plane.bgp_config

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/)
- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/)
- local_control_plane.bgp_config

<a id="section"></a>

Type: `"single"`. Computed.

BGP Configuration. BGP configuration parameters.

Upstream description:

BGP configuration parameters.

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

<a id="schema-local_control_plane--bgp_config--asn"></a>

### asn property

Type: `"number"`. Computed.

ASN. Autonomous System Number.

Upstream description:

Autonomous System Number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/): complete subsection reference.

## Next pages

- [local_control_plane.bgp_config.peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/bgp_config/peers/)
- [local_control_plane](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/properties/local_control_plane/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
