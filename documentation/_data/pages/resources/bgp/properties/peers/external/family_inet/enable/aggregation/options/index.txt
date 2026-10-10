---
page_title: "peers.external.family_inet.enable.aggregation.options"
subcategory: ""
description: "Configuration parameter for options"
xcsh_docs: {"aliases": ["peers external family inet enable aggregation options"], "body_bytes": 2165, "body_sha256": "sha256:b958a757504023a65ef6e6570d43fcf1df78f9a63b0ab3420276b2165953e868", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable:aggregation:options:summary_only"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable:aggregation:options", "parent_id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable:aggregation", "path": "documentation/resources/bgp/properties/peers/external/family_inet/enable/aggregation/options/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3220221002330213-2220220311222320-2232032232101021-1011011200301021-2210323313020322-3322000210321201-3301033001112220-1330001221223233", "registry_path": "docs/guides/resources--bgp--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["peers", "external", "family_inet", "enable", "aggregation", "options"], "schema_version": 1, "sections": [{"aliases": ["peers external family inet enable aggregation options summary only"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:bgp:properties:peers:external:family_inet:enable:aggregation:options:summary_only", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["peers", "external", "family_inet", "enable", "aggregation", "options", "summary_only"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/peers/external/family_inet/enable/aggregation/options/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Configuration parameter for options", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["bgpCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# peers.external.family_inet.enable.aggregation.options

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [peers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/)
- [peers.external](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/)
- [peers.external.family_inet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/)
- [peers.external.family_inet.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/enable/)
- [peers.external.family_inet.enable.aggregation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/enable/aggregation/)
- peers.external.family_inet.enable.aggregation.options

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Aggregation OPTIONS. Configuration parameter for options

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [summary_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/peers/external/family_inet/enable/aggregation/options/summary_only/): complete subsection reference.
