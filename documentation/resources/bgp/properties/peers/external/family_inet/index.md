---
page_title: "peers.external.family_inet"
subcategory: ""
description: "Parameters for inet family."
xcsh_docs: {"aliases": ["peers external family inet"], "body_bytes": 2148, "body_sha256": "sha256:c540afee99ceba4a2843d3419916d9625b46cdacc632b99152ef93e6f51c555f", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp:properties:peers:external:family_inet:disable_spec", "xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet", "parent_id": "xcsh-docs:resources:bgp:properties:peers:external", "path": "documentation/resources/bgp/properties/peers/external/family_inet/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0113212100231110-0013013110022232-3211013200330312-2210023313122002-3331020102102121-3233203020233100-2133131030203003-1011022003130203", "registry_path": "docs/guides/resources--bgp--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "peers.external.family_inet:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "peers.external.family_inet:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "external", "family_inet"], "schema_version": 1, "sections": [{"aliases": ["disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:disable_spec", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["peers", "external", "family_inet", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable"], "anchor": "section", "description": "IPv4 Unicast.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["peers", "external", "family_inet", "enable"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/external/family_inet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Parameters for inet family.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.family_inet

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/)
- peers.external.family_inet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for family inet.

Upstream description:

Parameters for inet family.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-enable_choice": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
family_inet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/enable/): complete subsection reference.

## Next pages

- [peers.external.family_inet.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/disable_spec/)
- [peers.external.family_inet.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/enable/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
